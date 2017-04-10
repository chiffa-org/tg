package tg

import (
	"io"
)

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name,omitempty"`
	UserName  *string `json:"username,omitempty"`
}

type Chat struct {
	ID                          int64    `json:"id"`
	Type                        ChatType `json:"type"`
	Title                       *string  `json:"title,omitempty"`
	UserName                    *string  `json:"username,omitempty"`
	FirstName                   *string  `json:"first_name,omitempty"`
	LastName                    *string  `json:"last_name,omitempty"`
	AllMembersAreAdministrators *bool    `json:"all_members_are_administrators,omitempty"`
}

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type Message struct {
	MessageID             int               `json:"message_id"`
	From                  *User             `json:"from,omitempty"`
	Date                  int               `json:"date"`
	Chat                  *Chat             `json:"chat"`
	ForwardFrom           *User             `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat             `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  *int              `json:"forward_from_message_id,omitempty"`
	ForwardDate           *int              `json:"forward_date,omitempty"`
	ReplyToMessage        *Message          `json:"reply_to_message,omitempty"`
	EditDate              *int              `json:"edit_date,omitempty"`
	Text                  *string           `json:"text,omitempty"`
	Entities              *[]*MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio            `json:"audio,omitempty"`
	Document              *Document         `json:"document,omitempty"`
	Game                  *Game             `json:"game,omitempty"`
	Photo                 *[]*PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker          `json:"sticker,omitempty"`
	Video                 *Video            `json:"video,omitempty"`
	Voice                 *Voice            `json:"voice,omitempty"`
	Caption               *string           `json:"caption,omitempty"`
	Contact               *Contact          `json:"contact,omitempty"`
	Location              *Location         `json:"location,omitempty"`
	Venue                 *Venue            `json:"venue,omitempty"`
	NewChatMember         *User             `json:"new_chat_member,omitempty"`
	LeftChatMember        *User             `json:"left_chat_member,omitempty"`
	NewChatTitle          *string           `json:"new_chat_title,omitempty"`
	NewChatPhoto          *[]*PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       *bool             `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      *bool             `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated *bool             `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    *bool             `json:"channel_chat_created,omitempty"`
	MigrateToChatID       *int64            `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     *int64            `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message          `json:"pinned_message,omitempty"`
}

type MessageEntity struct {
	Type   MessageEntityType `json:"type"`
	Offset int               `json:"offset"`
	Length int               `json:"length"`
	URL    *string           `json:"url,omitempty"`
	User   *User             `json:"user,omitempty"`
}

type MessageEntityType string

const (
	MessageEntityTypeMention     MessageEntityType = "mention"
	MessageEntityTypeHashtag     MessageEntityType = "hashtag"
	MessageEntityTypeBotCommand  MessageEntityType = "bot_command"
	MessageEntityTypeURL         MessageEntityType = "url"
	MessageEntityTypeEmail       MessageEntityType = "email"
	MessageEntityTypeBold        MessageEntityType = "bold"
	MessageEntityTypeItalic      MessageEntityType = "italic"
	MessageEntityTypeCode        MessageEntityType = "code"
	MessageEntityTypePre         MessageEntityType = "pre"
	MessageEntityTypeTextLink    MessageEntityType = "text_link"
	MessageEntityTypeTextMention MessageEntityType = "text_mention"
)

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize *int   `json:"file_size,omitempty"`
}

type Audio struct {
	FileID    string  `json:"file_id"`
	Duration  int     `json:"duration"`
	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`
	MimeType  *string `json:"mime_type,omitempty"`
	FileSize  *int    `json:"file_size,omitempty"`
}

type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MIMEType *string    `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    *string    `json:"emoji,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Voice struct {
	FileID   string  `json:"file_id"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *int    `json:"file_size,omitempty"`
}

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int    `json:"user_id,omitempty"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location     *Location `json:"location"`
	Title        string    `json:"title"`
	Address      string    `json:"address"`
	FoursquareID *string   `json:"foursquare_id,omitempty"`
}

type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

type File struct {
	FileID   string  `json:"file_id"`
	FileSize *int    `json:"file_size,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
}

type ReplyMarkup interface {
	replyMarkup()
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  *bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard *bool               `json:"one_time_keyboard,omitempty"`
	Selective       *bool               `json:"selective,omitempty"`
}

func (*ReplyKeyboardMarkup) replyMarkup() {}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  *bool  `json:"request_contact,omitempty"`
	RequestLocation *bool  `json:"request_location,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

func (*ReplyKeyboardRemove) replyMarkup() {}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

func (*InlineKeyboardMarkup) replyMarkup() {}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply bool  `json:"force_reply"`
	Selective  *bool `json:"selective,omitempty"`
}

func (*ForceReply) replyMarkup() {}

type ChatMember struct {
	User   *User            `json:"user"`
	Status ChatMemberStatus `json:"status"`
}

type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

type ResponseParameters struct {
	MigrateToChatID *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int   `json:"retry_after,omitempty"`
}

type InputFile struct {
	io.Reader
	Name string
	Size int
}
