# CLI

```sh
Usage:
  dsql [command]

Available Commands:
  compile     Statically check SQL for syntax and type errors
  completion  Generate the autocompletion script for the specified shell
  createdb    Create an ephemeral database
  diff        Compare the generated files to the existing files
  generate    Generate source code from SQL
  help        Help about any command
  init        Create an empty sqlc.yaml settings file
  version     Print the dsql version number
  vet         Vet examines queries

Flags:
  -f, --file string    specify an alternate config file (default: sqlc.yaml)
  -h, --help           help for dsql
      --no-database    disable database connections (default: false)

Use "dsql [command] --help" for more information about a command.
```
