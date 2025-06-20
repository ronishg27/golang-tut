# base image
FROM  golang:1.24.2-alpine3.20

# select working directory
WORKDIR /app

# copying all files from the current directory to the working directory
COPY . .

# downloading dependencies
RUN go mod download

# building the application

RUN go build -o myapi .

# exposing the port on which the application will run
EXPOSE 8080

# running the application
CMD [ "./myapi" ]
