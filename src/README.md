# README for Study Assist Bot in Go

## Overview

The Study Assist Bot is a Telegram bot designed to help users learn and test their knowledge on various topics. This bot is built using Go and the go-telegram/bot library, providing a robust and efficient solution for conducting tests, tracking user statistics, and offering multilingual support.

## Project Structure

The project is organized into several directories, each serving a specific purpose:

- **cmd/bot**: Contains the entry point of the application.
- **internal**: Holds the core logic of the bot, including configuration, controllers, models, repository, services, utilities, and views.
- **pkg**: Contains reusable packages, such as the Telegram bot interface.
- **locales**: Contains localization files for multilingual support.
- **scripts**: Includes scripts for database migration.
- **go.mod** and **go.sum**: Manage dependencies for the Go module.

## Features

- **Section Management**: Users can interact with different sections of content.
- **Testing Functionality**: Conduct tests by sections and topics, track user progress, and provide feedback.
- **User Statistics**: Track user performance and statistics related to their interactions with the bot.
- **Explanations**: Provide detailed explanations for questions to enhance learning.
- **Multilingual Support**: Support for multiple languages, allowing users to interact in their preferred language.

## Getting Started

### Prerequisites

- Go 1.18 or higher
- A Telegram bot token (create a bot using [BotFather](https://core.telegram.org/bots#botfather))
- A database (e.g., SQLite, PostgreSQL) for storing user data and statistics

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd study-assist-bot-go
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Set up your environment variables or configuration file for the bot token and database connection.

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