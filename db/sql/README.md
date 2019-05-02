# PostgreSQL test database setup

Use this setup to quickly integrate a database with the tests in the project.

## Usage

1. Create a new database called __<db_name>__.
2. Then, at the root of the project execute the following:

```

psql -U <db_user> -p <db_port> -h <db_host> --dbname=<db_name> --command="\i ./db/sql/deploy.sql"

```

## License

[MIT](https://choosealicense.com/licenses/mit/)
