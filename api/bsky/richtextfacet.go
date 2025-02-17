package bsky

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/bluesky-social/indigo/lex/util"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// schema: app.bsky.richtext.facet

func init() {
}

type RichtextFacet struct {
	Features []*RichtextFacet_Features_Elem `json:"features" cborgen:"features"`
	Index    *RichtextFacet_ByteSlice       `json:"index" cborgen:"index"`
}

type RichtextFacet_ByteSlice struct {
	ByteEnd   int64 `json:"byteEnd" cborgen:"byteEnd"`
	ByteStart int64 `json:"byteStart" cborgen:"byteStart"`
}

type RichtextFacet_Features_Elem struct {
	RichtextFacet_Mention *RichtextFacet_Mention
	RichtextFacet_Link    *RichtextFacet_Link
}

func (t *RichtextFacet_Features_Elem) MarshalJSON() ([]byte, error) {
	if t.RichtextFacet_Mention != nil {
		t.RichtextFacet_Mention.LexiconTypeID = "app.bsky.richtext.facet#mention"
		return json.Marshal(t.RichtextFacet_Mention)
	}
	if t.RichtextFacet_Link != nil {
		t.RichtextFacet_Link.LexiconTypeID = "app.bsky.richtext.facet#link"
		return json.Marshal(t.RichtextFacet_Link)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *RichtextFacet_Features_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.richtext.facet#mention":
		t.RichtextFacet_Mention = new(RichtextFacet_Mention)
		return json.Unmarshal(b, t.RichtextFacet_Mention)
	case "app.bsky.richtext.facet#link":
		t.RichtextFacet_Link = new(RichtextFacet_Link)
		return json.Unmarshal(b, t.RichtextFacet_Link)

	default:
		return nil
	}
}

func (t *RichtextFacet_Features_Elem) MarshalCBOR(w io.Writer) error {

	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if t.RichtextFacet_Mention != nil {
		return t.RichtextFacet_Mention.MarshalCBOR(w)
	}
	if t.RichtextFacet_Link != nil {
		return t.RichtextFacet_Link.MarshalCBOR(w)
	}
	return fmt.Errorf("cannot cbor marshal empty enum")
}
func (t *RichtextFacet_Features_Elem) UnmarshalCBOR(r io.Reader) error {
	typ, b, err := util.CborTypeExtractReader(r)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.richtext.facet#mention":
		t.RichtextFacet_Mention = new(RichtextFacet_Mention)
		return t.RichtextFacet_Mention.UnmarshalCBOR(bytes.NewReader(b))
	case "app.bsky.richtext.facet#link":
		t.RichtextFacet_Link = new(RichtextFacet_Link)
		return t.RichtextFacet_Link.UnmarshalCBOR(bytes.NewReader(b))

	default:
		return nil
	}
}

// RECORDTYPE: RichtextFacet_Link
type RichtextFacet_Link struct {
	LexiconTypeID string `json:"$type,const=app.bsky.richtext.facet" cborgen:"$type,const=app.bsky.richtext.facet"`
	Uri           string `json:"uri" cborgen:"uri"`
}

// RECORDTYPE: RichtextFacet_Mention
type RichtextFacet_Mention struct {
	LexiconTypeID string `json:"$type,const=app.bsky.richtext.facet" cborgen:"$type,const=app.bsky.richtext.facet"`
	Did           string `json:"did" cborgen:"did"`
}
