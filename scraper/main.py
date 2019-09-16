from scrape import retrieve_raw, logger
from extract import extract_article

def main():
    content = retrieve_raw('https://realpython.com/blog/')
    if content is not None:
        title, article_length = extract_article(content)
        logger('Article is called {}, it\'s length is {} characters'.format(title, article_length))
    logger("done")

if __name__ == "__main__":
    main()