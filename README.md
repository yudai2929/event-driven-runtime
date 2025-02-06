# Event-Driven Runtime

## Overview
**Event-Driven Runtime** is an ongoing project that implements a function execution runtime similar to AWS Lambda or Cloud Functions using the Go language. This project is being developed as a practice exercise rather than for real-world deployment. It provides the capability to execute compiled Go binaries and aims to support the following trigger features.

## Features
- **Binary Execution**: Executes compiled Go binaries.
- **Event Triggers**: Processes HTTP requests and custom events as triggers.
- **Message Queue Triggers**: Supports message processing based on Pub/Sub and queue-based mechanisms.
- **File System Watcher Triggers**: Detects file changes and executes functions.
- **Cron Triggers**: Executes scheduled tasks.
- **Scalability Enhancements**: Auto-scales instances under high load conditions.

## Getting Started
### Prerequisites
- Go 1.20+
- Docker (optional)

### Installation
```sh
git clone https://github.com/yudai2929/event-driven-runtime.git
cd event-driven-runtime
go build -o runtime
```

### Usage
#### Execute a Binary
```sh
./runtime --exec /path/to/binary
```

#### Using HTTP Trigger
```sh
./runtime --http :8080
```

#### Using Cron Trigger
```sh
./runtime --cron "*/5 * * * *" --exec /path/to/binary
```

## Roadmap
- [ ] WebAssembly support
- [ ] Integration with API Gateway
- [ ] Multi-runtime support (Python, Node.js, etc.)

## License
MIT License
```

