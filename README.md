### Development Set Up

1. Run `go install` to install dependencies.

2. Ensure your postgres database is up and running then source database variables, defaults to default postgres user and database.
   1. `export DB_NAME=postgres`
   2. `export DB_USER=postgres`

3. Start the client
```bash
go run warehousemanager-client/cmd/client/main.go --server-grpc-port 2000 --client-port 5000
```

4. On a different terminal, start the server:
```bash
go run warehousemanager-server/cmd/server/main.go --grpc-port 2000 --server-port 8000
```

5. Open a different terminal and send your csv file over for processing.
```bash
export FILE_PATH=/path/to/csv/file.csv
curl --location --request POST 'localhost:5000/upload' --form 'file=@'"\"${FILE_PATH}\""
```

6. Generate Cargo
```bash
curl --location --request GET 'localhost:8000/cargo?page_id=1&limit=5'
```


### Suggested Improvements
1. Unit and Functional Tests on the CargoList Interface to ensure all the logic is sound
2. Using db pools to increase number of concurrent db operations; we can also increase max_connections on our db to accomodate this.
3. There should be a `MaxCargoSize` to ensure more even order distribution
4. Ideally I would have this in different repositories to decouple their deployments
5. Create more advanced ErrorHandling and Logging Interfaces to facilitate different errors and logs
6. Have more definitive fields on the Order and Cargo schemas such as isComplete, archivedAt in order to capture more information.
7. Make this idempotent by using external_id as the primaryKey
8. End-to-end tests to ensure that the data served in the client matches what the server receives