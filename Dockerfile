FROM golang:1.17

ENV GO111MODULE=off

COPY . /release/app

#RUN apt-get update && apt-get install curl -y
#RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
#RUN mv migrate.linux-amd64 /usr/bin/migrate
#
#RUN migrate -path /release/app/migrations -database postgresql://identity:identitypass@localhost:5432/identity?sslmode=disable up

#COPY ./scripts/wait-for-it.sh /usr/wait-for-it.sh
#RUN chmod +x /usr/wait-for-it.sh

WORKDIR /release/app
RUN make build

CMD /release/app/bin/identity
