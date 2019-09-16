# 🤖 Resource Scraper

Given a URL this service scrapes the title (`<h1>`) and content (all `<p>`) combined. It returns them as a `json` that looks like:

{
    title: "article title",
    content: "article content"
}

## ⚙️ Tech stack

Python 3.7
PyTest
Pipenv

## 💻 Local Development

You'll need [pipenv](https://pipenv.readthedocs.io/en/latest/) installed.

To run the program use the command:

```sh
pipenv run python main.py
```

### 🐋 Run in Docker

You can run the scraper in a Docker container by running the following command in the root directory:

```sh
docker-compose up --build
```
*The `--build` tag is required only on the first run and when changes have been made.*

## ✔️ Tests

All tests can be run with:

```sh
pipenv python -m pytest tests/
```
