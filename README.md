# RSS Scraper

A Go-based RSS scraper designed to fetch and process RSS feeds concurrently. This project uses Go's concurrency features to efficiently scrape multiple RSS feeds and handle them in parallel.

## Features

- Concurrently fetches and processes RSS feeds
- Periodically scrapes feeds based on a defined interval
- Marks feeds as fetched in a PostgreSQL database
- Logs the details of the posts found in each feed

## Installation

### Prerequisites

- Go 1.18+ (for Go development)
- PostgreSQL (for the database)

### Clone the Repository

```bash
git clone https://github.com/YourUsername/RSS-Scraper-go.git
cd RSS-Scraper-go
```

### Setup

1. **Install Dependencies**

   Run the following command to download the required Go modules:

   ```bash
   go mod tidy
   ```

2. **Configure the Database**

   Set up a PostgreSQL database and configure the connection settings. Ensure you have a table for feeds and the necessary fields.

## Usage

### Running the Scraper

Build and run the RSS scraper using Go:

```bash
go run main.go
```

Ensure that the `config.yml` file is correctly set up and accessible by the application.

## Code Structure

- **`main.go`**: Entry point of the application. Contains the `startScraping` function and the main application logic.
- **`internal/`**: Contains internal packages for database access and authentication.
  - **`auth/`**: Handles authentication-related functionality.
  - **`database/`**: Contains database queries and models.
- **`config.yml`**: Example configuration file for setting up database and scraping intervals.

