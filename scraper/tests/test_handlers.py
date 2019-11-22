# Tests for handlers module
from aiohttp import web

import handlers.handlers as handlers
import scrape.scrape as scrape

original_retrieve_raw = scrape.retrieve_raw


def return_none(url):
    return None


def mock_retrieve(url):
    return """
    <html><body>
    <div><h1>My Title</h1></div>
    <div class="aticle-body"><p> some text  </p></div>
    </body></html>
    """


async def mock_client(handler, aiohttp_client, params={}):
    app = web.Application()
    # tests on root as the handler is interchanged for the tests
    app.router.add_get("/", handler)
    client = await aiohttp_client(app)
    resp = await client.get("/", params=params)
    return resp


async def test_handle_article_successful(aiohttp_client):
    # replaces current function in scrape package to mock function mock_retrieve
    scrape.retrieve_raw = mock_retrieve

    resp = await mock_client(
        handlers.handle_article,
        aiohttp_client,
        params={"URL": "https://realpython.com/blog/"},
    )
    # reset current functions to originals for future tests in suite
    scrape.retrieve_raw = original_retrieve_raw

    assert resp.status == 200
    json_response = await resp.json()
    assert "meta" in json_response
    assert json_response["meta"]["url"] == "https://realpython.com/blog/"


async def test_handle_article_unsuccessful_no_url(aiohttp_client):
    resp = await mock_client(handlers.handle_article, aiohttp_client, params={})

    assert resp.status == 404
    json_response = await resp.json()


async def test_handle_article_unsuccessful_no_content(aiohttp_client):
    # replaces current function in scrape package to mock function mock_retrieve
    scrape.retrieve_raw = return_none
    resp = await mock_client(
        handlers.handle_article,
        aiohttp_client,
        params={"URL": "https://realpython.com/blog/"},
    )
    # reset current functions to originals for future tests in suite
    scrape.retrieve_raw = original_retrieve_raw

    assert resp.status == 500
    json_response = await resp.json()


if __name__ == "__main__":
    sys.exit(pytest.main(["--ignore", "external"]))
