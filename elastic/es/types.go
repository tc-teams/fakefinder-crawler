package es

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawler"
)

// Hit is a structure used for serializing/deserializing data in Elasticsearch.
type Data struct {
	Version string              `json:"version,omitempty"`
	News    crawler.RelatedNews `json:"fields"`
	Message string              `json:"message,omitempty"`
	Time    timestamp.Timestamp `json:"time,omitempty"`
}


var (
	Fields = "fields.Title"
	Index  = "fakefinder"
)
