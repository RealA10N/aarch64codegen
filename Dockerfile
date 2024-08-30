# This container is used for testing.
# It is required because some tests compare the generated code with the output
# of the gcc assembler (aarch64-linux-gnu-as), so we use a container to ensure
# the required environment and binaries are present to run the tests.

FROM golang:1.23

WORKDIR /aarch64codegen

COPY . .

RUN go mod download

CMD ["go", "test", "-v", "./..."]