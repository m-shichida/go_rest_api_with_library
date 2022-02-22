golang 勉強用 REST API サーバー

- echo
- sqlx

## migration

used golang-migrate.

**create**

```
migrate create -ext sql -dir db/migrations -seq file_name
```

**up**

```
migrate -database $DATABASE_URL -path db/migrations up
```

**down**

```
migrate -database $DATABASE_URL -path db/migrations down
```
