FROM golang:1.20
WORKDIR /app
RUN apt-get update \
    && apt-get install -y libc6 \
    && rm -rf /var/lib/apt/lists/*
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
CMD air
