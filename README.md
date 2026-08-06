# dsqlsea: A Dynamic SQL Compiler (temporary hard fork of sqlc)

![go](https://github.com/theoriginalstove/dsqlsea/workflows/go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/theoriginalstove/dsqlsea)](https://goreportcard.com/report/github.com/theoriginalstove/dsqlsea)

dsqlsea (CLI: `dsql`) generates **type-safe code** from SQL. It is based directly on
sqlc; see the upstream [overview](https://docs.sqlc.dev) and
[playground](https://play.sqlc.dev/) for the original motivation. This fork was created to add support for dynamic queries (such as those used for filtering or end user chosed list ordering)

## Overview

- [Documentation](https://docs.sqlc.dev) (upstream sqlc docs)
- [Installation](./docs/overview/install.md)
- [Community](https://discord.gg/EcXzGe5SEs) (upstream sqlc community)

## Supported languages

- [sqlc-gen-go](https://github.com/sqlc-dev/sqlc-gen-go)
- [sqlc-gen-kotlin](https://github.com/sqlc-dev/sqlc-gen-kotlin)
- [sqlc-gen-python](https://github.com/sqlc-dev/sqlc-gen-python)
- [sqlc-gen-typescript](https://github.com/sqlc-dev/sqlc-gen-typescript)

Additional languages can be added via upstream [plugins](https://docs.sqlc.dev/en/latest/reference/language-support.html#community-language-support).
