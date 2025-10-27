# Study Assist Bot (Go Migration)

## Overview

The Study Assist Bot is a Telegram bot designed to help users learn and test their knowledge on various topics. This project is a **migration from Python (aiogram) to Go** using the [go-telegram/bot](https://github.com/go-telegram/bot) library.

## Current Status

✅ **Phase 1: Basic Bot Setup** - COMPLETED
- Project structure created
- Bot initialization with go-telegram/bot
- Configuration management with environment variables
- Graceful shutdown handling

⏳ **Phase 2: Core Functionality** - IN PROGRESS
- Command handlers
- Testing system
- User statistics
- Multilingual support

## Quick Start

### Prerequisites

- Go 1.21 or higher
- A Telegram bot token (get from [@BotFather](https://t.me/botfather))

### Installation & Running

1. **Clone and navigate to the project:**
   ```bash
   cd src
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Configure environment:**
   ```bash
   cp .env.example .env
   # Edit .env and add your TELEGRAM_BOT_TOKEN
   ```

4. **Run the bot:**
   ```bash
   # Option 1: Direct run
   go run ./cmd/bot

   # Option 2: Build and run
   go build -o bot ./cmd/bot
   ./bot
   ```

5. **Stop the bot:**
   Press `Ctrl+C` for graceful shutdown

### Running the Bot

To start the bot, run the following command:

```
go run cmd/bot/main.go
```

### Testing

The project includes unit tests and integration tests to ensure functionality. To run the tests, use:

```
go test ./...
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch and create a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- Thanks to the developers of the go-telegram/bot library for providing a powerful tool for building Telegram bots in Go.
- Special thanks to the community for their support and contributions.
