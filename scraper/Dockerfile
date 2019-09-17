FROM python:3.7-alpine

COPY requirements.txt /
RUN pip install -r requirements.txt
COPY . /scraper
WORKDIR /src/resource-searcher/scraper

CMD URL='https://realpython.com/python-web-scraping-practical-introduction/' python /scraper/main.py

