# Stage 1: Builder Stage
FROM golang:1.23-bullseye AS builder

# Add non-root user
RUN useradd -u 1001 nonroot

# Working Directory
WORKDIR /app

# Install air and goose dependencies
RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy only go.mod for efficient dependency installation
COPY go.mod go.sum ./

# Use cache mounts for faster dependency installation
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# Copy the entire application source code
COPY . .

# Build the app
RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o backend-boilerplate-template ./cmd

# Execute the command for migration and running the app
#ENTRYPOINT [ "/bin/ash", "-c", "source .env && goose -dir database/migrations up && air" ]
CMD [ "air" ]

# Use a minimal scratch image as the base
## Stage 3: Final Stage gcr.io/distroless/static-debian11 | scratch
FROM gcr.io/distroless/static

# Working Directory
WORKDIR /

# Copy the env file
COPY .env /

# Copy /etc/passwd for non-root user information
COPY --from=builder /etc/passwd /etc/passwd

# Copy the compiled application binary
COPY --from=builder /app/backend-boilerplate-template /backend-boilerplate-template
COPY --from=builder /app/queries /queries
COPY --from=builder /app/public/templates /public/templates

# Use the non-root user
USER nonroot

# Expose the application port
EXPOSE 8081

# Define the command to run the application
ENTRYPOINT ["./backend-boilerplate-template"]
