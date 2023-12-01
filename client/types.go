package client

import (
	"go.uber.org/zap"
)

// Client represents a MTProto client to Telegram.
type TGoCall struct {
	Log *zap.Logger // immutable

	// Client config.
	ApiID   int    // immutable
	ApiHash string // immutable
}

type GroupCall struct {
}
