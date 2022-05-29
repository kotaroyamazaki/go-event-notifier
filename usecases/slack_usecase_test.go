package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/KotaroYamazaki/go-event-notifier/entites"
	"github.com/KotaroYamazaki/go-event-notifier/pkg/slack"
)

func Test_slackUsecase_Notify(t *testing.T) {
	type args struct {
		ctx    context.Context
		appLog *entites.AppLog
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "AppEngine のログでの正常系",
			args: args{
				ctx: context.Background(),
				appLog: &entites.AppLog{
					InsertID: "test-insert-id",
					JsonPayload: entites.JsonPayload{
						Message:     "test-message",
						FirebaseUID: "test-firebase-uid",
						Caller:      "test-caller",
						AppVersion:  "v1.0.0",
					},
					Resource: entites.Resource{
						Type: entites.ResourceTypeAppEngine,
						Labels: entites.Labels{
							ProjectID:    "project-id",
							RevisionName: "app engine",
							ServiceName:  "test-service",
						},
					},
					TextPayload: "test-text-payload",
					Timestamp:   time.Now().String(),
					Severity:    "ERROR",
					Trace:       "test-trace",
				},
			},
			wantErr: false,
		},
		{
			name: "該当しないリソースタイプの場合はエラーが返る",
			args: args{
				ctx: context.Background(),
				appLog: &entites.AppLog{
					InsertID: "test-insert-id",
					JsonPayload: entites.JsonPayload{
						Message:     "test-message",
						FirebaseUID: "test-firebase-uid",
						Caller:      "test-caller",
						AppVersion:  "v1.0.0",
					},
					Resource: entites.Resource{
						Type: "hoge",
						Labels: entites.Labels{
							ProjectID:    "project-id",
							RevisionName: "app engine",
							ServiceName:  "test-service",
						},
					},
					TextPayload: "test-text-payload",
					Timestamp:   time.Now().String(),
					Severity:    "ERROR",
					Trace:       "test-trace",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &slackUsecase{
				slackClient: slack.NewMockSlackClient(),
			}
			if err := uc.Notify(tt.args.ctx, tt.args.appLog); (err != nil) != tt.wantErr {
				t.Errorf("slackUsecase.Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
