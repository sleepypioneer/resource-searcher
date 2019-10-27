FROM python:3.7-alpine

COPY requirements.txt /
RUN pip install -r requirements.txt
COPY . /scraper
WORKDIR /scraper
EXPOSE 8001
CMD python main.py