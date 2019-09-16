from scrape import retrieve_raw, logger

def main():
    content = retrieve_raw('https://realpython.com/blog/')
    if content is not None:
        logger("Length of content is {} characters".format(len(content)))
    logger("done")

if __name__ == "__main__":
    main()