# Tests for extract module
import extract.extract as extract
from bs4 import BeautifulSoup


def test_get_article_title():
    html = "<div><h1>My Title</h1></div>"
    html_string = BeautifulSoup(html, "html.parser")
    extracted_title = extract.get_article_title(html_string)
    assert extracted_title == "My Title"


def test_get_article_content():
    html = """
    <html><body>
    <div class="aticle-body"><p> some text </p> </div>
    </body></html>
    """
    html_string = BeautifulSoup(html, "html.parser")
    content = extract.get_article_content(html_string)
    assert content == {"some text"}


def test_extract_article():
    html = """
    <html><body>
    <div><h1>My Title</h1></div>
    <div class="aticle-body"><p> some text </p></div>
    </body></html>
    """
    title, content_length = extract.extract_article(html)
    assert title == "My Title"
