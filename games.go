package tg

const SendGameMethod MethodName = "sendGame"

type SendGame struct {
	Method              *MethodName           `json:"method,omitempty"`
	ChatID              ChatID                `json:"chat_id"`
	GameShortName       string                `json:"game_short_name"`
	DisableNotification *bool                 `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int                  `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendGameResult *Message

type Game struct {
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Photo        []*PhotoSize      `json:"photo"`
	Text         *string           `json:"text,omitempty"`
	TextEntities *[]*MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation        `json:"animation,omitempty"`
}

type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type CallbackGame struct{}

const SetGameScoreMethod MethodName = "setGameScore"

type SetGameScore struct {
	Method             *MethodName `json:"method,omitempty"`
	UserID             int         `json:"user_id"`
	Score              int         `json:"score"`
	Force              *bool       `json:"force,omitempty"`
	DisableEditMessage *bool       `json:"disable_edit_message,omitempty"`
	ChatID             *int64      `json:"chat_id,omitempty"`
	MessageID          *int        `json:"message_id,omitempty"`
	InlineMessageID    *string     `json:"inline_message_id,omitempty"`
}

type SetGameScoreResult *Message

const GetGameHighScoresMethod MethodName = "getGameHighScores"

type GetGameHighScores struct {
	Method          *MethodName `json:"method,omitempty"`
	UserID          int         `json:"user_id"`
	ChatID          *int64      `json:"chat_id,omitempty"`
	MessageID       *int        `json:"message_id,omitempty"`
	InlineMessageID *string     `json:"inline_message_id,omitempty"`
}

type GetGameHighScoresResult []*GameHighScore

type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}
