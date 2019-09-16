from bs4 import BeautifulSoup

def get_article_title(html_string):
    # H1
    return "title"

def get_article_body(html_string):
# p in article-body
    return "body of article"

def extract_article(raw_content):
    html = BeautifulSoup(raw_content, 'html.parser')

    title = get_article_title(html)
    body = get_article_body(html)

    return title, len(body)