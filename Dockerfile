FROM golang:1.21.6

WORKDIR /app/auth_service

COPY . .
RUN go mod download

RUN go build -o /goauth
CMD [ "/goauth" ]