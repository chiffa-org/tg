package tg

const EditMessageTextMethod MethodName = "editMessageText"

type EditMessageText struct {
	Method                *MethodName           `json:"method,omitempty"`
	ChatID                ChatID                `json:"chat_id,omitempty"`
	MessageID             *int                  `json:"message_id,omitempty"`
	InlineMessageID       *string               `json:"inline_message_id,omitempty"`
	Text                  string                `json:"text"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool                 `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageTextResult *Message

const EditMessageCaptionMethod MethodName = "editMessageCaption"

type EditMessageCaption struct {
	Method          *MethodName           `json:"method,omitempty"`
	ChatID          ChatID                `json:"chat_id,omitempty"`
	MessageID       *int                  `json:"message_id,omitempty"`
	InlineMessageID *string               `json:"inline_message_id,omitempty"`
	Caption         string                `json:"caption"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCaptionResult *Message

const EditMessageReplyMarkupMethod MethodName = "editMessageReplyMarkup"

type EditMessageReplyMarkup struct {
	Method          *MethodName           `json:"method,omitempty"`
	ChatID          ChatID                `json:"chat_id,omitempty"`
	MessageID       *int                  `json:"message_id,omitempty"`
	InlineMessageID *string               `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkupResult *Message
