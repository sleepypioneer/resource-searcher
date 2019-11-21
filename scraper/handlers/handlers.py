import logging
from aiohttp import web
import scrape.scrape as scrape
import extract.extract as extract


async def handle_article(request):
    try:
        URL = request.rel_url.query["URL"]
        logging.info(URL)
    except Exception as e:
        logging.info("Error retrieving URL parameter: {}".format(str(e)))
        response = web.json_response({})
        response.set_status(404)
        return response

    content = scrape.retrieve_raw(URL)
    if content is None:
        response = web.json_response({})
        response.set_status(500)
        return response

    article_title, article_content = extract.extract_article(content)
    logging.info(
        f'Article is called "{article_title}", it contains {len(article_content)} paragraphs.'
    )
    response = web.json_response(
        {
            "meta": {"url": URL,},
            "document": {"title": article_title, "content": list(article_content),},
        }
    )
    response.set_status(200)
    return response
