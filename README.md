# Prime-Based Backoff Algorithm

This project implements a backoff algorithm inspired by the prime number sequence from the movie *Contact*. It uses prime numbers as retry intervals and includes integration tests for load testing.

> **Disclaimer**: This project may contain traces of late-night inspiration and questionable life choices. Use at your own risk. 😅

## Features

- Prime-based backoff intervals (2s, 3s, 5s, 7s, 11s, ...).
- Integration tests using Vegeta for HTTP load testing.
- Docker and Docker Compose support for local development and testing.
- GitHub Actions for automated testing.

## Getting Started

### Prerequisites

- Go 1.20+
- Docker and Docker Compose

### Running Locally

1. Clone the repository:
   ```
   git clone https://github.com/ericmanlol/prime-backoff.git
   cd prime-backoff
   ```

2. Build and run the application:
   ```
   make build
   make run
   ```

3. Run unit tests:
   ```
   make test
   ```

4. Run integration tests:
   ```
   make integration-test
   ```

### Running with Docker

1. Build and run the application with Docker:
   ```
   make docker-build
   make docker-run
   ```

2. Run integration tests with Docker Compose:
   ```
   make integration-test
   ```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.