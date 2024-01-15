# Tectonic Sync

Tectonic Sync is a cronjob designed to update player names in our database using the Wise Old Man name change API endpoint. It utilizes the Wise Old Man API for retrieving player name change information.

## Installation

To run this project, you need to have Go installed.

```bash
# Clone the repository
git clone https://github.com/yourusername/tectonic-sync.git

# Change into the project directory
cd tectonic-sync

# Run the application
go run main.go
```

## Dependencies

* Wise Old Man API
* Postgres database

## Usage

This cron job can be executed using a cron scheduler. The application itself does not internally handle anything cron-related. You can use a tool like Railway for hosting and scheduled runs.

```bash
# Example command to run the cron job
go run main.go --db=DATABASE_URL --group=WOM_GROUP_ID --verbose
```

## Configuration

The application accepts configuration through either terminal flags or environment variables:

> [!WARNING]
> You need to be running a Postgres database.

### Flags
* `--db=DB_URL`: Database connection URL.
* `--group=WOM_GROUP`: Group ID for Wise Old Man API.
* `--verbose`: Flag for verbose output.

### Environment variables
* `DATABASE_URL`: Database connection URL.
* `GROUP_ID`: Wise Old Man group id to query for namechanges.

### SQL
Batch query:
```sql
UPDATE rsn SET rsn = $1 WHERE wom_id = $2;
```
Our table structure:
```sql
DROP TABLE IF EXISTS "rsn";
CREATE TABLE "public"."rsn" (
    "rsn" character varying(32) NOT NULL,
    "wom_id" character varying(32) NOT NULL,
    "user_id" character varying(32) NOT NULL,
    "guild_id" character varying(32) NOT NULL,
    CONSTRAINT "rsn_pkey" PRIMARY KEY ("wom_id", "guild_id")
) WITH (oids = false);
```

## Output

The output of the cron job indicates the number of players updated and provides details about which players were affected.
Logging

Logs are outputted to both stdout and stderr.

## Contributing

Contributions to this project are welcome. If you have suggestions, enhancements, or bug fixes, feel free to submit a pull request.
