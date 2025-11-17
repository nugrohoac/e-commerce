## Prerequisites
Make sure you have installed all of the following prerequisites on your development machine:
* go - [Download & Install](https://go.dev/dl/)
* golang-migrate - [Install golang-migrate](https://github.com/golang-migrate/migrate)
* MySQL
* mockery - [Install golang-migrate](https://github.com/vektra/mockery)


## Integration Test
This is will running integration test and unit test.
#### 1. Create file migrate
```bash
$ migrate create -ext sql -dir destination/directory name_migration
# Example
$ migrate create -ext sql -dir migrations create_table_applicant_order
```

Each migration has an up and down migration.
```bash
20251116111716_create_table_order.up.sql
20251116111716_create_table_order.down.sql
```
* write up ddl at file.up.sql
* write reverse of ddl at file.down.sql

#### 2. Create file migrate
Make sure you have modified configuration at
* driver     = "defauult mysql"
* host       = "default is localhost if running local"
* dbname     = "make sure you have been created database for testing"
* sslMode    = "deafult disable"
* userName   = "default root"
* password   = ""
* searchPath = "default is public"
* port       = "default is 5432"


#### Running Integration Test
```bash
$ make test
```

## Unit Test
This is just running unit test without integration test. Make sure mocks is up-to-date.
* generate or update mock base on name of [interface](src/business/contract.go)
#### generate or update mocks
```bash
$ mockery -name=name-of-interface
$ mockery --dir=source/directory --name=nameInterface --output=destination/directory
```
#### Running Integration Test
```bash
$ make unittest
```

## Running Apps
#### Running Integration Test
* download all dependencies base on go.mod
```bash
$ go mod vendor
```
* running command
```bash
$ make run
```
