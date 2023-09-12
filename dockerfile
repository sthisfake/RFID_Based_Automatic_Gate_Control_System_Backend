# FROM golang:alpine AS build

# RUN apk add build-base

# WORKDIR /app

# COPY . ./

# # Install dependencies
# RUN go mod download && \
#   # Build the app
#   CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main && \
#   # Make the final output executable
#   chmod +x ./main

# FROM alpine:latest

# # Install os packages
# RUN apk --no-cache add bash

# WORKDIR /app

# COPY --from=build /app/main .
# COPY ./pb_data /app/pb_data

# CMD ["./main" ,"serve", "--http=0.0.0.0:8000"]

# EXPOSE 8000

FROM golang:latest

WORKDIR /src

# copy all your .go source files
# (or use a .dockerignore and COPY . .)
COPY *.go .

# remove any previously initialized go.mod and go.sum files
# (this is in case the container data wasn't destroyed)
RUN rm -f go.mod rm go.sum

# initialize Go modules
RUN go mod init app

# fetch dependencies
RUN go mod tidy

# build (switch to 1 to use the CGO SQLite)
RUN CGO_ENABLED=0 go build -o /pocketbase

# export listener port
EXPOSE 8090

# run
# (ps. don't forget to mount the pb_data as volume in /pb_data)
CMD ["/pocketbase", "serve", "--http=0.0.0.0:8090"]