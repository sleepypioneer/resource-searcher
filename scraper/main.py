from scrape import retrieve_raw, logger
from extract import extract_article

def main():
    content = retrieve_raw('https://realpython.com/python-web-scraping-practical-introduction/')
    if content is not None:
        title, article_length = extract_article(content)
        logger('Article is called "{}", it contains {} paragraphs.'.format(title, article_length))
    logger("done")

if __name__ == "__main__":
    main()