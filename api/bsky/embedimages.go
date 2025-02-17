package bsky

import (
	"github.com/bluesky-social/indigo/lex/util"
)

// schema: app.bsky.embed.images

func init() {
	util.RegisterType("app.bsky.embed.images#main", &EmbedImages{})
}

// RECORDTYPE: EmbedImages
type EmbedImages struct {
	LexiconTypeID string               `json:"$type,const=app.bsky.embed.images" cborgen:"$type,const=app.bsky.embed.images"`
	Images        []*EmbedImages_Image `json:"images" cborgen:"images"`
}

type EmbedImages_Image struct {
	Alt   string        `json:"alt" cborgen:"alt"`
	Image *util.LexBlob `json:"image" cborgen:"image"`
}

// RECORDTYPE: EmbedImages_View
type EmbedImages_View struct {
	LexiconTypeID string                   `json:"$type,const=app.bsky.embed.images" cborgen:"$type,const=app.bsky.embed.images"`
	Images        []*EmbedImages_ViewImage `json:"images" cborgen:"images"`
}

type EmbedImages_ViewImage struct {
	Alt      string `json:"alt" cborgen:"alt"`
	Fullsize string `json:"fullsize" cborgen:"fullsize"`
	Thumb    string `json:"thumb" cborgen:"thumb"`
}
