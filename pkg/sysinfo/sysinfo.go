/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package sysinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	secrets "github.com/BlueSageSolutions/secrets/cmd"
)

type SystemInfo struct {
	SystemName          sql.NullString `json:"system_name"`
	URL                 sql.NullString `json:"url"`
	Username            sql.NullString `json:"user_name"`
	Password            sql.NullString `json:"password"`
	EnabledYN           bool           `json:"enabled_yn"`
	Notes               sql.NullString `json:"notes"`
	SystemProperties    sql.NullString `json:"system_properties"`
	ConfigData          sql.NullString `json:"config_data"`
	ValidationClassName sql.NullString `json:"validation_class_name"`
	AwsSystemType       sql.NullString `json:"aws_system_type"`
}

func Sluggify(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}

func Unsluggify(input string) string {
	return strings.ToUpper(strings.ReplaceAll(input, "-", " "))
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

func Normalize(systemInfo SystemInfo) (string, error) {
	normal := make(map[string]string, 0)

	normal["system_name"] = systemInfo.SystemName.String
	normal["url"] = systemInfo.URL.String
	normal["user_name"] = systemInfo.Username.String
	normal["password"] = systemInfo.Password.String
	normal["enabled_yn"] = strconv.FormatBool(systemInfo.EnabledYN)
	normal["notes"] = systemInfo.Notes.String
	normal["system_properties"] = systemInfo.SystemProperties.String
	normal["config_data"] = systemInfo.ConfigData.String
	normal["validation_class_name"] = systemInfo.ValidationClassName.String
	normal["aws_system_type"] = systemInfo.AwsSystemType.String
	normalized, err := json.Marshal(normal)
	if err != nil {
		return "", err
	}
	return string(normalized), err
}

func Denormalize(system, url, username, password, notes, systemProperties, configData, validationClass, awsSystemType string, enabled bool) string {

	denormal := make(map[string]string, 0)
	denormal["system_name"] = Unsluggify(system)
	denormal["url"] = url
	denormal["user_name"] = username
	denormal["password"] = password
	denormal["enabled_yn"] = strconv.FormatBool(enabled)
	denormal["notes"] = notes
	denormal["system_properties"] = systemProperties
	denormal["config_data"] = configData
	denormal["validation_class_name"] = validationClass
	denormal["aws_system_type"] = awsSystemType
	denormalized, err := json.Marshal(denormal)
	if err != nil {
		return ""
	}
	return string(denormalized)
}

func SetSystemInfo(profile, region, sluggifiedName, client, environment string, systemInfo string) error {

	err := secrets.SetSecret(profile, region, "", "secrets", client, environment, sluggifiedName, "system-info", systemInfo)
	if err != nil {
		return fmt.Errorf("failed to put parameter in Parameter Store: %v", err)
	}

	fmt.Printf("Stored parameter for system: %s\n", sluggifiedName)
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
			system_name, 
			url, 
			user_name, 
			password, 
			enabled_yn, 
			notes, 
			system_properties, 
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
			&systemInfo.SystemName,
			&systemInfo.URL,
			&systemInfo.Username,
			&systemInfo.Password,
			&enabledYNBit,
			&systemInfo.Notes,
			&systemInfo.SystemProperties,
			&systemInfo.ConfigData,
			&systemInfo.ValidationClassName,
			&systemInfo.AwsSystemType,
		); err != nil {
			return err
		}

		systemInfo.EnabledYN = len(enabledYNBit) > 0 && enabledYNBit[0] == 1

		sluggifiedName := Sluggify(systemInfo.SystemName.String)
		parameterValue, err := Normalize(systemInfo)
		if err != nil {
			return fmt.Errorf("failed to marshal system info to JSON: %v", err)
		}
		err = SetSystemInfo(profile, region, sluggifiedName, client, environment, parameterValue)
		if err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
