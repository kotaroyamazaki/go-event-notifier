package slack

const (
	MentionAtChannel = "<!channel>"
	MentionAtHere    = "<!here>"
)

type SlackMessage struct {
	Text        string
	Attachments []Attachment
}

type color string

const (
	AttachmentColorGood    color = "good"    // green
	AttachmentColorWarning color = "warning" // yellow
	AttachmentColorError   color = "danger"  // red
)

type Attachment struct {
	Color     color
	Title     string
	TitleLink string
	Text      string
	Fields    []AttachmentField
}

type AttachmentField struct {
	Title string
	Value string
	Short bool
}

func (a *Attachment) AddField(title, val string) *Attachment {
	a.Fields = append(a.Fields, AttachmentField{
		Title: title,
		Value: val,
		Short: false,
	})
	return a
}

func (a *Attachment) AddShortField(title, val string) *Attachment {
	a.Fields = append(a.Fields, AttachmentField{
		Title: title,
		Value: val,
		Short: true,
	})
	return a
}
