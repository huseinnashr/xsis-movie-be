# xsis-movie-be

## Get Started
### 1. Start dev environment
11. Download docker-desktop with docker-compose cli https://www.docker.com/products/docker-desktop/
12. Run docker-compose up -d

### 2. Setup dev tool
21. Install `make` cli https://formulae.brew.sh/formula/make
22. Run `make setup` to install goose and other dev tool

### 3. Migrate DB with Goose
31. Set these env
```sh
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/postgres"
```
32. Run goose -dir migrations up
```sh
goose -dir migrations up
```

### 4. Start service
41. Run `make start-local`


## Try the API
### 1. List Movie
```sh
curl --location 'localhost:8080/movies?page_size=2&page_token=eyJpZCI6MX0'
```

### 2. Create Movie
```sh
curl --location 'localhost:8080/movies' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Pengabdi Setan 2 Comunion",
    "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
    "rating": 7.2,
    "image": "image.png"
}'
```

### 3. Get Movie
```sh
curl --location 'localhost:8080/movies/1'
```

### 4. Update Movie
```sh
curl --location --request PATCH 'localhost:8080/movies/1' \
--header 'Content-Type: application/json' \
--data '{
    "rating": 7,
    "image": "new-image.png"
}'
```

### 5. Delete Movie
```sh
curl --location --request DELETE 'localhost:8080/movies/4'
```

## Test
### 1. Create Movie
![List Movie Page 1](./files/screenshots/1.%20List%20Page%201.png)
![List Movie Page 1](./files/screenshots/1.%20List%20Page%202.png)

### 2. List Movie
![Create Movie](./files/screenshots/2.%20Create%20Movie.png)

### 3. Get Movie
![Get Movie](./files/screenshots/3.%20Get%20Movie.png)

### 4. Update Movie
![Update Movie](./files/screenshots/4.%20Update%20Movie.png)

### 5. Delete Movie
![Delete Movie](./files/screenshots/5.%20Delete%20Movie.png)
