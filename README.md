# cards-api

These APIs are developed with go1.19 and mysql8.0

Clone the source

    git clone https://github.com/shadow300893/cards-api

Make sure that cloned directory is accessible in GOPATH

Setup dependencies

    go get -u github.com/go-chi/chi
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/google/uuid
    go get -u github.com/joho/godotenv

Setup mysql database structure by importing db.sql file

Add the database config details in .env file for following params:
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=

Run the app

    go build .
    go run .

And visit

## API ENDPOINTS

### Create a Deck
- Path : `http://localhost:8000/decks/create`
- Method: `GET`
- Response: `201`
- QueryParams: `?shuffled=true/false and ?cards=KS,AC`

### Open a Deck
- Path : `http://localhost:8000/decks/{id}`
- Method: `GET`
- Response: `200`

### Draw cards
- Path : `http://localhost:8000/decks/{id}/cards/draw`
- Method: `GET`
- Response: `200`
- QueryParams: `?count=2`