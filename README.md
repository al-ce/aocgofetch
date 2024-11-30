# aocgofetch

Fetches Advent of Code puzzle inputs and prints to standard out.

![screenshot](https://github.com/user-attachments/assets/78a7cd4a-e4b2-4b19-826c-f6c868fae363)

## Installation

Install with go >= 1.17

```bash
go install github.com/al-ce/aocgofetch@latest
```

Or clone this repository and build with `go build -o <path-to-binary>`.

This project uses [joho/gotodenv](https://github.com/joho/godotenv) as a dependency.

## Usage

### Authenticate

Authenticate yourself at `adventofcode.com` and get the value for the `session` cookie from your browser.

- Firefox: `Developer Tools > Storage > Cookies`
- Chrome: `Developer Tools > Application > Cookies`

Add that value to a `.env` file at the root of your working directory, like your directory where you're writing your puzzle-solvers.

### Check your `.gitignore`

***WARNING:*** make sure you add the `.env` to your `.gitignore` if you're backing up your files to a remote repository!

***Reminder:*** please respect the author and do not include the puzzle text or the puzzle input in your repository.

### Run the program

Give the year (2015-present) and day (1-25) as arguments.
The puzzle input will print to `stdout`:

```bash
❯ ./aocgofetch 2015 4
ckczppom
```
