## StraySafe

#### Running the backend
The backend, authentication service and the database are containerized as services in the docker compose file. Simply 
run `docker compose up -d` to build and deploy the containers. Modify the environment variables per your need, especially the JWT_TOKEN.
To run the migrations

```shell
goose postgres "postgres://database:password@localost:5432/straysafe" up
```