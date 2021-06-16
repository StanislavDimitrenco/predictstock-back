package providers

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/telegram"
)

// Boot all needed dependencies
func Boot(ctx context.Context) context.Context {
	ctx = database.Boot(ctx)
	telegramBot := telegram.NewBot()
	ctx = context.WithValue(ctx, "telegramBot", telegramBot)

	return ctx
}
