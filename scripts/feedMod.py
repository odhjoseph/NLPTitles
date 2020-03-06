from bs4 import BeautifulSoup as soup
from urllib.request import urlopen as uReq

def articleReader(url):
    uClient = uReq(url)
    page_html = uClient.read()
    uClient.close()
    page_soup = soup(page_html, "html.parser")
    if not page_soup:
        page_soup.replace('\n', " ")
    return page_soup.prettify()


def webInfo(html, subString):
    indexOfSub = [i for i in range(len(html)) if html.startswith(subString, i)]
    return indexOfSub


def relevantInfo(html):
    info = set()
    indexesStart = (webInfo(html, "<p>"))
    indexesEnd = (webInfo(html, "</p>"))
    for i in range(min(len(indexesStart), len(indexesEnd))):
        info.add(html[indexesStart[i]:indexesEnd[i]].replace("\n", " "))
    return info
