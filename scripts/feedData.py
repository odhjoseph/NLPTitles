import feedparser
import json
from datetime import date
import schedule
import time
import feedMod

rssFeeds = ['https://www.bloomberg.com/view/rss/topics/economics.rss',
            'https://www.brookings.edu/blog/ben-bernanke/feed/',
            'cnbc.com/id/10000666/device/rss',
            'https://www.schaeffersresearch.com/feeds/news+and+analysis',
            'https://optionalpha.com/feed',
            'http://feeds.marketwatch.com/marketwatch/topstories/'
            ]


def writeDict(rssFeed):
    feeds = dict()
    searchDict = dict()
    for feed in rssFeed:
        rssInfo = feedparser.parse(feed)
        item = rssInfo["items"]
        if item:
            try:
                for article in item:
                    title = article["title"]
                    summary = article["summary"]
                    link = article["link"]
                    if feed in searchDict:
                        searchDict[feed].append(title)
                    else:
                        searchDict.update({feed: [title]})

                    feeds.update({title: summary})
            except:  # reminder
                print("Unformatted Token: " + str(item))
        else:
            print(rssInfo)

    today = date.today().isoformat() + ".json"
    with open("jsonFeeds/" + today, "wb") as f:
        f.write(json.dumps(feeds).encode("utf-8"))
    with open('titleDisplay/' + "search" + today, "wb") as f:
        f.write(json.dumps(searchDict).encode("utf-8"))

writeDict(rssFeeds)

# schedule.every().day.at('11:55').do(writeDict)

# while True:
#    schedule.run_pending()
#    time.sleep(1)
