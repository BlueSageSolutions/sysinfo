# sysinfo: A tool for managing `sys_global_external_system`

The `sysinfo` tool provide get/set access to the `sysinfo` records in AWS Parameter Store. It also provides an import feature to copy data from the `sys_global_external_system` table into AWS Parameter Store.

## Installation

Download the appropriate binary from this repository's [releases](https://github.com/BlueSageSolutions/sysinfo/releases).

## Usage

```sh
$ sysinfo --help
Manage the system info records in parameter store

Usage:
  sysinfo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Retrieve system info record(s) from parameter store
  help        Help about any command
  import      Import system info records into parameter store
  set         Create a system info record in parameter store
  version     Version of sysinfo

Flags:
  -h, --help     help for sysinfo
  -t, --toggle   Help message for toggle

Use "sysinfo [command] --help" for more information about a command.
```

### Import

The `import` command will connect to the appropriately configured database for the provided `client` and `environment` arguments. If the `system` is provided, then only the records for that system will be imported. Otherwise, all records in the `sys_global_external_system` table will be imported into  AWS Parameter Store.

The process of `import` will sluggify the `system` name so as to be compatible with AWS Parameter Store.

```sh
$ sysinfo import --help
Import system info records into parameter store

Usage:
  sysinfo import [flags]

Flags:
      --client string        Name of client account. Use the client code always!
      --environment string   Environment
  -h, --help                 help for import
      --profile string       Profile name used for secrets account
      --region string        Region used for secrets account (default "us-east-1")
      --system string        System name

```

### Set

The `set` command will create an entry in AWS Parameter Store for the provided system information.

It will format the supplied variables as JSON in a consistent manner.

```sh
$ sysinfo set --help
Create a system info record in parameter store

Usage:
  sysinfo set [flags]

Flags:
      --aws-system-type string     Aws System Type
      --client string              Name of client account. Use the client code always!
      --config-data string         Config Data
      --enabled                    Enabled
      --environment string         Environment
  -h, --help                       help for set
      --notes string               Notes
      --password string            Password
      --profile string             Profile name used for secrets account
      --region string              Region used for secrets account (default "us-east-1")
      --system string              System name
      --system-properties string   System Properties
      --url string                 URL
      --username string            Username
      --validation-class string    Validation Class
```

### Get

The `get` command will retrieve a `system-info` record for a given system.

```sh
$ sysinfo get --help
Retrieve system info record(s) from parameter store

Usage:
  sysinfo get [flags]

Flags:
      --client string        Name of client account. Use the client code always!
      --environment string   Environment
  -h, --help                 help for get
      --profile string       Profile name used for secrets account
      --region string        Region used for secrets account (default "us-east-1")
      --system string        System name
```