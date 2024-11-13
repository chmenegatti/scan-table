# Scan Table

## Introduction

This is a simple program that scans a database and returns all tables with at least one column that is passed as an argument.

This is a Command Line Interface (CLI) program that takes a database connection string in yaml file and column name as arguments.

## Installation

To install the program, clone the repository and run the following command:

```bash
git clone

cd scan-table

go mod tidy

```
## Usage

At the root of the project, you have to create a yaml file with the following format:

```yaml
database:
  host: localhost
  port: 3306
  user: dbuser
  password: password
  dbname: dbname

```

Then you can run the program with the following command:

```bash
go run main.go -n column_name
```

## Example

```bash
go run main.go -n fileset
```
This command shoul returns all tables with a column named fileset.

```bash
Tabelas que contêm pelo menos um dos campos especificados:
data_protection_hosts
fileset_per_clusters
shares
```
