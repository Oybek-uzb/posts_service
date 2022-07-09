# IMAN test-task Posts Service

This service contains main business logic of the test-task.

Here posts which are fetched from https://gorest.co.in/public/v1/posts are stored in Postgres DB. So before running this service in any machine there must be installed Postgres.
Also, there are migrations that have to be migrated. The posts fetched in many goroutines from 50 different pages. Requests can be sent from API Gateway or Posts CRUD Service to this service. Both gateway and postsCrud requests are responded via gRPC connections.

Recommended PostgresURL =  "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

Running on port ":8082"