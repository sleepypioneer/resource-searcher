from bs4 import BeautifulSoup


def get_article_title(html_string):
    """
    Returns a string of the article title taken form the
    first <h1> tag in the parsed HTML string passed to it.
    """
    return html_string.find("h1").text.strip()


def get_article_content(html_string):
    """
    Returns a string with all the text from the
     <p> tags in the div with class `article-body`
    from the parsed HTML string passed to it.
    """
    article_body = html_string.find("div", class_="article-body")
    article_content = set()
    for p in html_string.select("p"):
        article_content.add(str(p.text.strip()))
    return article_content


def extract_article(raw_content):
    """
    Extracts the article content and body from the raw_content of the website
    """
    html = BeautifulSoup(raw_content, "html.parser")

    title = get_article_title(html)
    content = get_article_content(html)
    return title, content
