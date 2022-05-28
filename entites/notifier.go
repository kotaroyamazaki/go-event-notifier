package entites

import "context"

type Notifier interface {
	Notify(ctx context.Context, appLog *AppLog) error
}
