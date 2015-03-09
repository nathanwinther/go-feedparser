package feedparser

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "github.com/nathanwinther/go-feedparser/atom"
    "github.com/nathanwinther/go-feedparser/rss"
)

type Item struct {
    Title string
    Url string
    Published int64
}

func Load(feed string) ([]*Item, error) {
    r, err := http.Get(feed)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return nil, err
    }

    v1, err := atom.Load(b)
    if err == nil {
        items := make([]*Item, len(v1.Items))
        for idx, item := range v1.Items {
            items[idx] = &Item {
                item.Title,
                item.ParseLink(),
                parseUTCTime(item.Updated, item.Published),
            }
        }
        return items, nil
    }

    v2, err := rss.Load(b)
    if err == nil {
        items := make([]*Item, len(v2.Items))
        for idx, item := range v2.Items {
            items[idx] = &Item {
                item.Title,
                item.Link,
                parseUTCTime(item.Published),
            }
        }
        return items, nil
    }

    return nil, errors.New(fmt.Sprintf("Could not load feed: %s", feed))
}

func parseUTCTime(args ...string) int64 {
    for _, s := range args {
        if s == "" {
            continue
        }
        t, err := time.Parse(time.RFC3339, s)
        if err == nil {
            return t.Unix()
        }
        t, err = time.Parse(time.RFC1123Z, s)
        if err == nil {
            return t.Unix()
        }
        t, err = time.Parse(time.RFC1123, s)
        if err == nil {
            return t.Unix()
        }
    }

    return time.Now().Unix()
}

