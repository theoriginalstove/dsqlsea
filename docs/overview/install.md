# Installing dsql

dsql is distributed as a single binary with zero dependencies.

## go install

Installing recent versions of dsql requires Go 1.21+.

```
go install github.com/theoriginalstove/dsqlsea/cmd/dsql@latest
```

For upstream sqlc, use `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`.

## Docker

This fork does not publish an official Docker image yet. For upstream sqlc:

```
docker pull sqlc/sqlc
```

Run `sqlc` using `docker run`:

```
docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
```

## Downloads

Pre-built binaries are not yet published for dsqlsea. See [downloads.sqlc.dev](https://downloads.sqlc.dev/)
for upstream sqlc releases.
