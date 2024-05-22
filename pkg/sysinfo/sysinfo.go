/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package sysinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	secrets "github.com/BlueSageSolutions/secrets/cmd"
)

type SystemInfo struct {
	Version             sql.NullInt64  `json:"version"`
	EffectiveDateTime   sql.NullTime   `json:"effective_date_time"`
	ExpirationDateTime  sql.NullTime   `json:"expiration_date_time"`
	SystemName          sql.NullString `json:"system_name"`
	URL                 sql.NullString `json:"url"`
	UserName            sql.NullString `json:"user_name"`
	Password            sql.NullString `json:"password"`
	EnabledYN           bool           `json:"enabled_yn"`
	Notes               sql.NullString `json:"notes"`
	SystemProperties    sql.NullString `json:"system_properties"`
	SecretManagerID     sql.NullString `json:"secret_manager_id"`
	ConfigData          sql.NullString `json:"config_data"`
	ValidationClassName sql.NullString `json:"validation_class_name"`
	AwsSystemType       sql.NullString `json:"aws_system_type"`
}

func Sluggify(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}

func GetDatabaseConnection(client, environment string) (*sql.DB, error) {
	ctx := &secrets.CommandContext{Namespace: "secrets", Client: client, Environment: environment}
	credentials, err := secrets.GetDatabaseCredentials(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve credentials: %v", err)
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s)/", credentials.Username, credentials.Password, credentials.Hostname)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}

func SetSystemInfo(profile, region, sluggifiedName, client, environment string, systemInfo SystemInfo) error {
	parameterValue, err := json.Marshal(systemInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal system info to JSON: %v", err)
	}

	err = secrets.SetSecret(profile, region, "", "secrets", client, environment, sluggifiedName, "system-info", string(parameterValue))
	if err != nil {
		return fmt.Errorf("failed to put parameter in Parameter Store: %v", err)
	}

	fmt.Printf("Stored parameter for system: %s\n", systemInfo.SystemName.String)
	return nil
}

func GetSchema(profile, region, client, environment string) (string, error) {
	schema, err := secrets.GetSecret("", profile, region, "secrets", client, environment, "database", "schema-name", "", true, true)
	if err != nil {
		return "", err
	}
	if len(schema) == 0 {
		return "", fmt.Errorf("no schema for %s in %s", client, environment)
	}
	return schema, err
}

func GetSQL(profile, region, client, environment, system string) (string, error) {
	schema, err := GetSchema(profile, region, client, environment)
	if err != nil {
		return "", err
	}

	query := `
		SELECT 
			version, 
			effective_date_time, 
			expiration_date_time, 
			system_name, 
			url, 
			user_name, 
			password, 
			enabled_yn, 
			notes, 
			system_properties, 
			secret_manager_id, 
			config_data, 
			validation_class_name, 
			aws_system_type 
		FROM {{schema}}.sys_global_external_system
	`
	query = strings.Replace(query, "{{schema}}", schema, 1)
	if len(system) > 0 {
		query = fmt.Sprintf("%s WHERE system_name LIKE '%%%s%%'", query, system)
	}
	return query, nil
}

func MigrateSystemInfo(profile, region, client, environment, system string) error {
	db, err := GetDatabaseConnection(client, environment)
	if err != nil {
		return err
	}
	defer db.Close()
	query, err := GetSQL(profile, region, client, environment, system)
	if err != nil {
		return err
	}
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var systemInfo SystemInfo
		var enabledYNBit []byte

		if err := rows.Scan(
			&systemInfo.Version,
			&systemInfo.EffectiveDateTime,
			&systemInfo.ExpirationDateTime,
			&systemInfo.SystemName,
			&systemInfo.URL,
			&systemInfo.UserName,
			&systemInfo.Password,
			&enabledYNBit,
			&systemInfo.Notes,
			&systemInfo.SystemProperties,
			&systemInfo.SecretManagerID,
			&systemInfo.ConfigData,
			&systemInfo.ValidationClassName,
			&systemInfo.AwsSystemType,
		); err != nil {
			return err
		}

		systemInfo.EnabledYN = len(enabledYNBit) > 0 && enabledYNBit[0] == 1

		sluggifiedName := Sluggify(systemInfo.SystemName.String)

		err = SetSystemInfo(profile, region, sluggifiedName, client, environment, systemInfo)
		if err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
