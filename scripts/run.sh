docker container stop rest-server newsletter-db
docker container rm rest-server newsletter-db

# build and run img. The steps are defined in  Dockerfile.multistage file
docker build -t rest_api:multistage -f Dockerfile.multistage .

DB_USER="${POSTGRES_USER:=charlesonyewuenyi}" 
DB_PASSWORD="${POSTGRES_PASSWORD:=password}" 
DB_NAME="${POSTGRES_DB:=rest_api}" 
DB_PORT="${POSTGRES_PORT:=5432}"


docker network create -d bridge newsletter-db

#  expose port 8080 to port 8080 on the host: [host_port]:[container_port]
docker run -d --rm \
    --network newsletter-db \
    --name rest-server \
    -p 8080:8080 \
    -e POSTGRES_USER=${DB_USER} \
    -e POSTGRES_PASSWORD=${DB_PASSWORD} \
    -e PGHOST=newsletter-db \
    -e PGPORT=${DB_PORT} \
    -e POSTGRES_DB=${DB_NAME} \
    rest_api:multistage 



# -v "$PWD:/var/lib/postgresql/data" \

# Allow to skip Docker if a dockerized Postgres database is already running
if [[ -z "${SKIP_DOCKER}" ]]; then
        docker run \
        --name=newsletter-db \
        --hostname=newsletter-db \
        -e POSTGRES_USER=${DB_USER} \
        -e POSTGRES_PASSWORD=${DB_PASSWORD} \
        -e POSTGRES_DB=${DB_NAME} \
        -p "${DB_PORT}":5432 \
        -d postgres \
        postgres -N 1000
fi

export PGPASSWORD="${DB_PASSWORD}"


