package atom

import (
    "encoding/xml"
)

type Atom struct {
    XMLName xml.Name `xml:"feed,http://www.w3.org/2005/Atom"`
    Items []Item `xml:"entry"`
}

type Item struct {
    Id string `xml:"id"`
    Title string `xml:"title"`
    Links []Link `xml:"link"`
    Published string `xml:"published"`
    Updated string `xml:"updated"`
}

type Link struct {
    Rel string `xml:"rel,attr"`
    Type string `xml:"type,attr"`
    Href string `xml:"href,attr"`
}

func (item *Item) ParseLink() string {
    // Daring Fireball Style
    for _, v := range item.Links {
        if v.Rel == "shorturl" {
            return v.Href
        }
    }
    for _, v := range item.Links {
        if v.Rel == "alternate" && v.Type == "text/html" {
            return v.Href
        }
    }
    return ""
}

func Load(b []byte) (*Atom, error) {
    atom := new(Atom)
    err := xml.Unmarshal(b, atom)
    return atom, err
}

