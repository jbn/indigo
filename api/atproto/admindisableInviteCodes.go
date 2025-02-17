package atproto

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// schema: com.atproto.admin.disableInviteCodes

func init() {
}

type AdminDisableInviteCodes_Input struct {
	Accounts []string `json:"accounts,omitempty" cborgen:"accounts,omitempty"`
	Codes    []string `json:"codes,omitempty" cborgen:"codes,omitempty"`
}

func AdminDisableInviteCodes(ctx context.Context, c *xrpc.Client, input *AdminDisableInviteCodes_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.admin.disableInviteCodes", nil, input, nil); err != nil {
		return err
	}

	return nil
}
