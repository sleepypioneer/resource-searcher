from scrape import retrieve_raw

def test_retieve_raw_HTML(): 
    raw_html = retrieve_raw('https://realpython.com/blog/')
    assert raw_html != None

def test_retrieve_raw_not_HTML():
    no_html = retrieve_raw('https://realpython.com/blog/nope-not-gonna-find-it')
    assert no_html == None