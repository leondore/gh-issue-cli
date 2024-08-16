# GitHub Issue Manager

This is a simple issue manager for GitHub repositories. It allows you to read, create, update and close issues. This is exercise 4.11 from the book: The Go Porgramming Language.

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
GITHUB_USER=<your_github_username>
GITHUB_TOKEN=<your_github_token>
EDITOR=<your_editor>
```

Then run the following commands:

```bash
./scripts/build.sh
```
