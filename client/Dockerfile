FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify
COPY . .


EXPOSE 6969
EXPOSE 420

CMD [ "sh", "/usr/src/app/client/create_cert.sh" ]