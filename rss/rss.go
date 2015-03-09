package rss

import (
    "encoding/xml"
)

type Rss struct {
    XMLName xml.Name `xml:"rss"`
    Items []Item `xml:"channel>item"`
}

type Item struct {
    Id string `xml:"guid"`
    Title string `xml:"title"`
    Link string `xml:"link"`
    Published string `xml:"pubDate"`
}

func Load(b []byte) (*Rss, error) {
    rss := new(Rss)
    err := xml.Unmarshal(b, rss)
    return rss, err
}

