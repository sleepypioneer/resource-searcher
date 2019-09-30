import logging
from requests import get
from requests.exceptions import RequestException
from contextlib import closing

def retrieve_raw(url):
    """
    Attempts to get the content at `url` by making an HTTP GET request.
    If the content-type of response is some kind of HTML/XML, return the
    text content, otherwise return None.
    """
    try:
        with closing(get(url, stream=True)) as resp:
            if is_HTML_response(resp):
                return resp.content
            else:
                logging.info('Error during requests to {0} : URL does not return a HTML response'.format(url))
                return None

    except RequestException as e:
        logging.info('Error during requests to {0} : {1}'.format(url, str(e)))
        return None


def is_HTML_response(resp):
    """
    Returns True if the response seems to be HTML, False otherwise.
    """
    content_type = resp.headers['Content-Type'].lower()
    return (resp.status_code == 200 
            and content_type is not None 
            and content_type.find('html') > -1)
