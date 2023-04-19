# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY main.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk --no-cache --update --verbose add grep bash tcpdump && \
    rm -rf /var/cache/apk/* /tmp/* /sbin/halt /sbin/poweroff /sbin/reboot

WORKDIR /var/tmp
RUN echo '#!/bin/sh' > cleanup.sh
RUN echo 'ls -tp trace*| tail -n +24 | xargs -I{} sh -c "echo \"Removing {}\"; rm {}"' >> cleanup.sh && chmod a+x cleanup.sh

# Need to put it inside the tcpdump code
#WORKDIR /dump
#CMD [ "-i", "any", "-U", "-K", "-w", "/dump/trace_%Y-%m-%d_%H%M%S", "-G", "3600", "-z", "/var/tmp/cleanup.sh" ]

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /usr/local/bin/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT [ "main" ]
