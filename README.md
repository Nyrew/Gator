
# Gator

Gator is a Go-based CLI tool for interacting with an RSS feed aggregator. It allows users to add, follow, and unfollow feeds, and scrape posts from them at configurable intervals. It is designed to run in a long-running loop, continuously fetching new posts from RSS feeds.

## Prerequisites

To run this program, you'll need the following software installed:

- **Go** (Version 1.18 or higher)  
  Install Go from [the official Go website](https://go.dev/dl/).
  
- **PostgreSQL** (for the database)  
  Install PostgreSQL from [the official PostgreSQL website](https://www.postgresql.org/download/).

## Installing Gator

### 1. Install the `gator` CLI

To install the `gator` CLI tool, you can use `go install`:

- **bash**
    ```bash
    go install github.com/your-github-username/gator@latest
    ```

This will download and compile the binary for gator, and place it in your Go bin directory, typically located at `$GOPATH/bin` or `$HOME/go/bin`. After installation, you can run the gator command from your terminal.

### 2. Set Up the Config File
Before running gator, you need to configure it by setting up a config file. Here’s an example of a simple configuration file:

```yaml
currentUserName: "user123"
database:
  user: "your-postgres-username"
  password: "your-postgres-password"
  dbname: "gator"
  host: "localhost"
  port: 5432
```

Place this configuration file at `~/.gator/config.yaml` or any location you prefer and make sure to specify its path when running gator.

### 3. Set Up PostgreSQL Database
The program requires a PostgreSQL database. Follow these steps to set it up:

1. **Create a new PostgreSQL database**:
    ```bash
    createdb gator
    ```

2. **Run the migrations** to set up the necessary tables and schema. These migrations include creating tables for users, feeds, and feed-follow relationships:
    ```bash
    psql -d gator -f migrations/
    ```

Make sure the database is running and that your connection details are correct in the config file.

---

## Running the Program

### 1. Running in Development Mode
If you are in development mode, you can run the program using the following command:

```bash
go run .
```

This will run the application directly from the source code, allowing you to develop and test without needing to build a binary.

### 2. Running in Production Mode
For production use, build the gator CLI tool:

```bash
go build
```

Then, run the program as a standalone binary:

```bash
./gator <command> <arguments>
```

Alternatively, you can install it globally using `go install` and run it directly from anywhere:

```bash
gator <command> <arguments>
```

---

## Available Commands

Here are some of the main commands you can use:

### `addfeed`
Adds a new feed to the system.

```bash
gator addfeed "Feed Name" "http://example.com/rss"
```

### `follow`
Follows a feed for the current user.

```bash
gator follow "http://example.com/rss"
```

### `unfollow`
Unfollows a feed for the current user.

```bash
gator unfollow "http://example.com/rss"
```

### `agg`
Aggregates feeds, fetching posts at regular intervals.

```bash
gator agg 1m
```

This command fetches posts from the feeds every minute. You can specify different time intervals, such as `1s`, `30s`, `1h`, etc.

---

## Contributing

If you'd like to contribute to gator, please fork the repository, create a new branch, and submit a pull request. Here are some ways you can help:

- Report bugs or suggest improvements
- Write additional tests
- Add more functionality or commands
- Improve documentation

---

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

## Contact

For any questions or issues, please open an issue in the GitHub repository, or contact me directly at [your-email@example.com].

---

### Key Points:
1. **Installation**: The instructions guide users to install `gator` using `go install` and configure the system with a YAML configuration file for the database connection and current username.
2. **Database Setup**: It explains setting up PostgreSQL, creating the database, and running migrations.
3. **Running the Program**: It distinguishes between development mode (using `go run .`) and production mode (using a built binary or `go install`).
4. **Commands**: The `README` lists the main available commands (`addfeed`, `follow`, `unfollow`, `agg`) and how to use them.
5. **Contributing**: It encourages contributions and provides a contact email.

After updating the `README.md` file, remember to push the changes to your GitHub repository:

```bash
git add README.md
git commit -m "Add README.md"
git push origin main
```
