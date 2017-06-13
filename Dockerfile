FROM golang:1.8

RUN wget -q https://github.com/Masterminds/glide/releases/download/0.10.1/glide-0.10.1-linux-amd64.tar.gz && \
    tar xvfz glide-0.10.1-linux-amd64.tar.gz -C /usr/local/bin --strip-components=1 linux-amd64/glide && \
    rm glide-0.10.1-linux-amd64.tar.gz

ENV DATABASE_URL $DATABASE_URL
ENV PORT ${PORT:-8080}
EXPOSE $PORT

WORKDIR /go/src/github.com/fsufitch/r9kd
COPY . .

RUN glide update --update-vendored
RUN go get github.com/fsufitch/r9kd
RUN go get github.com/fsufitch/r9kd/db/r9kd_migrate

CMD r9kd_migrate postgres $DATABASE_URL up && r9kd $PORT
