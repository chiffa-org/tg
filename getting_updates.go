package tg

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
}

const GetUpdatesMethod MethodName = "getUpdates"

type GetUpdates struct {
	Method         *MethodName   `json:"method,omitempty"`
	Offset         *int          `json:"offset,omitempty"`
	Limit          *int          `json:"limit,omitempty"`
	Timeout        *int          `json:"timeout,omitempty"`
	AllowedUpdates *[]UpdateType `json:"allowed_updates,omitempty"`
}

type GetUpdatesResult []*Update

const SetWebhookMethod MethodName = "setWebhook"

type SetWebhook struct {
	Method          *MethodName   `json:"method,omitempty"`
	URL             string        `json:"url"`
	CertificateFile *InputFile    `json:"-"`
	MaxConnections  *int          `json:"max_connections,omitempty"`
	AllowedUpdates  *[]UpdateType `json:"allowed_updates,omitempty"`
}

type SetWebhookResult bool

const DeleteWebhookMethod MethodName = "deleteWebhook"

type DeleteWebhook struct {
	Method *MethodName `json:"method,omitempty"`
}

type DeleteWebhookResult bool

const GetWebhookInfoMethod MethodName = "getWebhookInfo"

type GetWebhookInfo struct {
	Method *MethodName `json:"method,omitempty"`
}

type GetWebhookInfoResult *WebhookInfo

type WebhookInfo struct {
	URL                  string        `json:"url"`
	HasCustomCertificate bool          `json:"has_custom_certificate"`
	PendingUpdateCount   int           `json:"pending_update_count"`
	LastErrorDate        *int          `json:"last_error_date,omitempty"`
	LastErrorMessage     *string       `json:"last_error_message,omitempty"`
	MaxConnections       *int          `json:"max_connections,omitempty"`
	AllowedUpdates       *[]UpdateType `json:"allowed_updates,omitempty"`
}

type UpdateType string

const (
	UpdateTypeMessage            UpdateType = "message"
	UpdateTypeEditedMessage      UpdateType = "edited_message"
	UpdateTypeChannelPost        UpdateType = "channel_post"
	UpdateTypeEditedChannelPost  UpdateType = "edited_channel_post"
	UpdateTypeInlineQuery        UpdateType = "inline_query"
	UpdateTypeChosenInlineResult UpdateType = "chosen_inline_result"
	UpdateTypeCallbackQuery      UpdateType = "callback_query"
)
