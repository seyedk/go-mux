export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD=password123
export APP_DB_NAME=postgres

docker run -it -p 5432:5432 --name MY_PQ -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_PASSWORD=password123  -d postgres 