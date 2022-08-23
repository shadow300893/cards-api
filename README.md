# cards-api

Clone the source

    git clone https://github.com/shadow300893/cards-api

Setup dependencies

    go get -u github.com/go-chi/chi
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/google/uuid
    go get -u github.com/joho/godotenv

Setup mysql database structure by importing db.sql file

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