# aocgofetch

Fetches Advent of Code puzzle inputs and prints to standard out.

## Installation

Clone and build with `go build -o <path-to-binary>`

## Usage

### Authentication

Authenticate yourself at `adventofcode.com` and get the value for the `session` cookie from your browser.

- Firefox: `Developer Tools > Storage > Cookies`
- Chrome: `Developer Tools > Application > Cookies`

Add that value to a `.env` file at the root of your working directory, like your directory where you're writing your puzzle-solvers.

### Hey! Listen!

***WARNING:*** make sure you add the `.env` to your `.gitignore` if you're backing up your files to a remote repository!

***Reminder:*** please respect the author and do not include the puzzle text or the puzzle input to your repository.

### CLI Usage

Give the year (2015-present) and day (1-25) as arguments to the program.
The puzzle input will print to `stdout`:

```bash
❯ ./aocgofetch 2015 4
ckczppom
```
