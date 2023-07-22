# Exchange Currencies Rate
A go utility/library to fetch currencies using [exchangeratesapi.io](https://exchangeratesapi.io/)
Please register on this site to get an API Key


## Dependencies 

- Database
  - [sqlx](https://github.com/jmoiron/sqlx) to connect to db 
  - [goose](https://github.com/pressly/goose) for migration
  - [lib-pq](https://github.com/lib/pq) as database driver

- Please install goose globally to run migrations using
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Run 
- copy .env.example file to .env and replace API_KEY with your own which can be obtained by the linked website above

- Start the db using docker, please make sure docker is allowed to create /var/db/data directory, create the directory /var/db/data using sudo
```
sudo mkdir -p /var/db/data 
sudo chown -R $USER:$USER /var/db/data
docker compose up -d
```

- Migrate the Database
```
make migrate-db
```

- Run the command
```
make run
```

This should fetch the details from the API and save them to DB and display the results from DB
