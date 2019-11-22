import time
import json
import os
import sys
import time
import asyncio
import logging
from http.server import BaseHTTPRequestHandler, HTTPServer
from aiohttp import web

from handlers.handlers import handle_article

HOST_NAME = "0.0.0.0"  # This will map to avialable port in docker
PORT_NUMBER = 8001


async def log_startup(app):
    """
    Log when server starts
    """
    logging.info("Server started at {}".format(time.time()))


def main():
    logging.basicConfig(
        format="%(asctime)s: %(levelname)s %(message)s", level=logging.DEBUG
    )

    app = web.Application(middlewares=[])
    [app.on_startup.append(i) for i in [log_startup,]]

    app.router.add_routes(
        [web.get("/article", handle_article),]
    )

    web.run_app(app, port=PORT_NUMBER)


if __name__ == "__main__":
    main()
