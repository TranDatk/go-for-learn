# Taskfile Guide for Go + PostgreSQL + Air

This document explains how to use the Taskfile configured for a Go project that uses PostgreSQL, Goose for migrations, and Air for live reloading during development.

The Taskfile automates common commands such as running the server, building the application, and managing database migrations.

## Prerequisites

Before using the Taskfile, ensure that the following tools are installed:

- Go 1.20+
- Task (taskfile.dev)
- Air (live reload tool)
- Goose (database migration tool)
- PostgreSQL

Your project must also contain a `.env` file providing the database connection string:

```text
DB_ADDRESS=postgres://username:password@localhost:5432/project?sslmode=disable
```

---

# Task Commands Overview

Below is an explanation of all available Taskfile commands.

---

## Development Commands

### task dev
Runs your Go application using Air, enabling automatic reload whenever you modify code.

```bash
task dev
```

### task run
Runs your application without Air. Useful for production-like execution.

```bash
task run
```

### task build
Builds your application into a binary located in the `./bin` directory.

```bash
task build
```


---

## Migration Commands

These commands use Goose to manage PostgreSQL migrations.  
The migration directory path used is:

```bash
./cmd/migrate/migrations
```


### task migrate:up
Applies all available migrations to the database.

```bash
task migrate:up
```

### task migrate:down
Rolls back the most recent migration.

```bash
task migrate:down
```

### task migrate:reset
Drops all migrations and applies them again from the beginning.

```bash
task migrate:reset
```


### task migrate:status
Shows the current migration status, including version and applied files.

```bash
task migrate:status
```

### task migrate:create
Creates a new migration file.  
Usage example:

```bash
task migrate:create name=migration_name
```

This generates an SQL migration file inside the migration directory.

---

## Go Utility Commands

### task fmt
Formats your Go code using `go fmt`.

```bash
task fmt
```

### task tidy
Updates and cleans dependencies in `go.mod` and `go.sum`.

```bash
task tidy
```

### task vet
Runs static analysis on your Go code.

```bash
task vet
```

### task test
Executes all Go tests within your project.

```bash
task test
```

### task clean
Removes the `bin` directory and other generated artifacts.

```bash
task clean
```

---

# Notes

1. The Taskfile automatically loads variables from `.env` using the `dotenv` setting.
2. All migration commands rely on the `DB_ADDRESS` value from `.env`.
3. Air requires a configuration file `.air.toml` if custom settings are needed.

