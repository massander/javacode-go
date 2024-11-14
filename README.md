# wallet-api

## How to run

Launch postgres

    docker compose up postgres

Exec into container and run migrations (for now so, fix later)
   docker exec -it postgres psql -U postgres

Run app
   
    docker compose up wallet-api
