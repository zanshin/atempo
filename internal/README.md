The `internal` directory contains all the logic and database pacakges for Atempo.

- app: The entry points into the application.
- config: Reads atempo configuration file and returns `Config` struct
- database: Database creation, connection, and queries
- logger: Simplistic logging helper
- model: Entities (structs) to match database tables or views
