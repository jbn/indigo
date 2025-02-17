package bsky

import (
	"context"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// schema: app.bsky.notification.listNotifications

func init() {
}

type NotificationListNotifications_Notification struct {
	Author        *ActorDefs_ProfileView             `json:"author" cborgen:"author"`
	Cid           string                             `json:"cid" cborgen:"cid"`
	IndexedAt     string                             `json:"indexedAt" cborgen:"indexedAt"`
	IsRead        bool                               `json:"isRead" cborgen:"isRead"`
	Labels        []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	Reason        string                             `json:"reason" cborgen:"reason"`
	ReasonSubject *string                            `json:"reasonSubject,omitempty" cborgen:"reasonSubject,omitempty"`
	Record        *util.LexiconTypeDecoder           `json:"record" cborgen:"record"`
	Uri           string                             `json:"uri" cborgen:"uri"`
}

type NotificationListNotifications_Output struct {
	Cursor        *string                                       `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Notifications []*NotificationListNotifications_Notification `json:"notifications" cborgen:"notifications"`
}

func NotificationListNotifications(ctx context.Context, c *xrpc.Client, cursor string, limit int64) (*NotificationListNotifications_Output, error) {
	var out NotificationListNotifications_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.notification.listNotifications", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
