# GoShort

GoShort is a web application built with Golang that facilitates the creation of short URLs for long URLs, providing users with a simple way to manage and share URLs efficiently.

## Features
- **Create Short URLs**: Quickly generate short URLs for any long URL.
- **Redirection**: Redirect to the original URL using the short URL.
- **Simple Frontend**: A basic HTML frontend for user interaction.

## Contributing
Contributions to URL Shortener are welcome! To contribute, follow these steps:

1. Fork the URL Shortener repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Make your changes to the codebase.
4. Ensure that your changes are thoroughly tested and functional.
5. Submit a pull request with your changes, providing a clear description of the modifications.

## Prerequisites
Before getting started, ensure you have Go installed on your system. If you don't already have it, you can download and install it by following the instructions at [https://go.dev/dl](https://go.dev/dl).

## Running the Project

1. Clone the repository.

```bash
git clone https://github.com/vhsenna/go-short.git
cd go-short
```

2. Initialize Go modules.

```bash
go mod tidy
```

3. Run the project.

```bash
go run ./cmd/web
```

4. Open your browser and navigate to `http://localhost:8080` to use the application.

## Upcoming Features

- [ ] **Custom Short URLs**: Allow users to create custom short URLs.
- [ ] **Analytics**: Track the number of clicks on each short URL.
- [ ] **User Authentication**: Enable users to manage their short URLs.
- [ ] **Expiration Dates**: Set expiration dates for short URLs.
