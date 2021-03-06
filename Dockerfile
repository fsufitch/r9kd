FROM golang:1.8
ARG POSTGRES_URL
ARG R9KD_PORT
ARG ADMIN_KEY

RUN wget -q https://github.com/Masterminds/glide/releases/download/0.10.1/glide-0.10.1-linux-amd64.tar.gz && \
    tar xvfz glide-0.10.1-linux-amd64.tar.gz -C /usr/local/bin --strip-components=1 linux-amd64/glide && \
    rm glide-0.10.1-linux-amd64.tar.gz

# Get source
WORKDIR /go/src/github.com/fsufitch/r9kd
COPY . .

# Configure environment
RUN . dev_scripts/env.sh
ENV DATABASE_URL $POSTGRES_URL
ENV PORT $R9KD_PORT
ENV BOOTSTRAP_ADMIN_KEY $ADMIN_KEY
EXPOSE $R9KD_PORT

# Get dependencies and install
RUN go get github.com/fsufitch/r9kd

# Entry point
CMD r9kd
