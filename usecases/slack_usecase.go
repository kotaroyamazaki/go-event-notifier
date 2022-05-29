package usecases

import (
	"context"
	"fmt"

	"github.com/KotaroYamazaki/go-event-notifier/entites"
	"github.com/KotaroYamazaki/go-event-notifier/pkg/slack"
)

type SlackUsecase interface {
	Notify(ctx context.Context, appLog *entites.AppLog) error
}

type slackUsecase struct {
	slackClient slack.SlackClient
}

func NewSlackUsecase(slackClient slack.SlackClient) *slackUsecase {
	return &slackUsecase{
		slackClient: slackClient,
	}
}

func (uc *slackUsecase) Notify(ctx context.Context, appLog *entites.AppLog) error {
	attachment := slack.Attachment{
		Title:     appLog.Severity + ": " + appLog.JsonPayload.Message,
		TitleLink: appLog.MakeLoggingURL(),
		Color:     slack.AttachmentColorError,
	}
	attachment.
		AddShortField("ProjectID", fmt.Sprintf("`%s`", appLog.Resource.Labels.ProjectID)).
		AddShortField("Resouce Type", fmt.Sprintf("`%s`", appLog.Resource.Type))

	switch appLog.Resource.Type {
	case entites.ResourceTypeCloudRun:
		attachment.
			AddShortField("AppVersion", fmt.Sprintf("`%s`", appLog.JsonPayload.AppVersion)).
			AddShortField("Firebase UID", fmt.Sprintf("`%s`", appLog.JsonPayload.FirebaseUID)).
			AddShortField("Timestamp", fmt.Sprintf("`%s`", appLog.Timestamp)).
			AddField("Caller", fmt.Sprintf("`%s`", appLog.JsonPayload.Caller)).
			AddField("Trace", fmt.Sprintf("`%s`", appLog.Trace)).
			AddField("Message", fmt.Sprintf("```%s```", appLog.JsonPayload.Message))
	case entites.ResourceTypeCloudFunction:
		attachment.
			AddShortField("FunctionName", fmt.Sprintf("`%s`", appLog.Resource.Labels.FunctionName)).
			AddShortField("Timestamp", fmt.Sprintf("`%s`", appLog.Timestamp)).
			AddField("Trace", fmt.Sprintf("`%s`", appLog.Trace)).
			AddField("Message", fmt.Sprintf("```%s```", appLog.TextPayload))
	case entites.ResourceTypeAppEngine:
		attachment.
			AddShortField("Module ID", fmt.Sprintf("`%s`", appLog.Resource.Labels.ModuleID)).
			AddShortField("Version ID", fmt.Sprintf("`%s`", appLog.Resource.Labels.VersionID)).
			AddShortField("Timestamp", fmt.Sprintf("`%s`", appLog.Timestamp)).
			AddField("Caller", fmt.Sprintf("`%s`", appLog.JsonPayload.Caller)).
			AddField("Message", fmt.Sprintf("```%s```", appLog.JsonPayload.Message))
	default:
		return fmt.Errorf("Unknown resource type: %s\n", appLog.Resource.Type)
	}

	msg := &slack.SlackMessage{
		Text:        fmt.Sprintf("%s%s速報", slack.MentionAtChannel, appLog.Severity),
		Attachments: []slack.Attachment{attachment},
	}
	if err := uc.slackClient.SetMessage(msg).Notify(); err != nil {
		return err
	}
	return nil
}
