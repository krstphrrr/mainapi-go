FROM golang:latest AS builder
WORKDIR /app
# COPY ./src/ .

COPY ./src/go.mod ./src/go.sum ./ 
RUN go mod download

#  dev packages
# RUN go get -u github.com/gorilla/mux
# RUN go install github.com/gorilla/mux@latest \ 
#   && cp $GOPATH/bin/package /usr/local/bin
# RUN go install
# CMD ["tail", "-f", "/dev/null"]

COPY ./src/ .
RUN CGO_ENABLED=0 GOOS=linux && go build -o /app/main 
EXPOSE 3000
# CMD ["tail", "-f", "/dev/null"]
CMD ["./main"]
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/main .
# EXPOSE 3000
# CMD ["tail", "-f", "/dev/null"]
# # CMD ["./main"]