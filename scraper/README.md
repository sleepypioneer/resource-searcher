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

## ✔️ Tests

All tests can be run with:

```sh
pipenv python -m pytest tests/
```
