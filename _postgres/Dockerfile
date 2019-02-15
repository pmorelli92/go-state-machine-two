FROM postgres:11.2-alpine
COPY _postgres/01_init.sql /docker-entrypoint-initdb.d/01.sql

ENV POSTGRES_PASSWORD="db_password"
ENV POSTGRES_USER="db_user"
ENV POSTGRES_DB="gostate"

## Local build, use from root folder of the repository
# eval $(minikube docker-env) &&  docker build -t postgres-go-state:local -f _postgres/Dockerfile .