import os
from scrape import retrieve_raw, logger
from extract import extract_article

def main():
    URL = os.environ['URL']
    content = retrieve_raw(URL)
    if content is not None:
        title, article_length = extract_article(content)
        logger('Article is called "{}", it contains {} paragraphs.'.format(title, article_length))
    logger("done")

if __name__ == "__main__":
    main()