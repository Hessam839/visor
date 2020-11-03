# Building the binary of the App
FROM golang AS build

# `boilerplate` should be replaced with your project name
WORKDIR /visor/src

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
#RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
RUN go build -o visor .

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

RUN apk add --update \
        bash \
        tzdata \
    && rm -rf /var/cache/apk/*

RUN cp /usr/share/zoneinfo/Asia/Tehran /etc/localtime

#RUN echo "Asia/Tehran" > /etc/timezone

#RUN date

WORKDIR /visor

# Create the `public` dir and copy all the assets into it
#RUN mkdir ./public
#COPY ./public ./public

# `boilerplate` should be replaced here as well
COPY --from=build /go/visor .

# Exposes port 3000 because our program listens on that port
EXPOSE 80

CMD ["./visor"]
