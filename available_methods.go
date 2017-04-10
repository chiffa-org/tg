package tg

const GetMeMethod MethodName = "getMe"

type GetMe struct {
	Method *MethodName `json:"method,omitempty"`
}

type GetMeResult *User

const SendMessageMethod MethodName = "sendMessage"

type SendMessage struct {
	Method                *MethodName `json:"method,omitempty"`
	ChatID                ChatID      `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             *ParseMode  `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool       `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendMessageResult *Message

type ParseMode string

const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML     ParseMode = "HTML"
)

const ForwardMessageMethod MethodName = "forwardMessage"

type ForwardMessage struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	FromChatID          ChatID      `json:"from_chat_id"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	MessageID           int         `json:"message_id"`
}

type ForwardMessageResult *Message

const SendPhotoMethod MethodName = "sendPhoto"

type SendPhoto struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Photo               *string     `json:"photo,omitempty"`
	PhotoFile           *InputFile  `json:"-"`
	Caption             *string     `json:"caption,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendPhotoResult *Message

const SendAudioMethod MethodName = "sendAudio"

type SendAudio struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Audio               *string     `json:"audio,omitempty"`
	AudioFile           *InputFile  `json:"-"`
	Caption             *string     `json:"caption,omitempty"`
	Duration            *int        `json:"duration,omitempty"`
	Performer           *string     `json:"performer,omitempty"`
	Title               *string     `json:"title,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendAudioResult *Message

const SendDocumentMethod MethodName = "sendDocument"

type SendDocument struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Document            *string     `json:"document,omitempty"`
	DocumentFile        *InputFile  `json:"-"`
	Caption             *string     `json:"caption,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendDocumentResult *Message

const SendStickerMethod MethodName = "sendSticker"

type SendSticker struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Sticker             *string     `json:"sticker,omitempty"`
	StickerFile         *InputFile  `json:"-"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendStickerResult *Message

const SendVideoMethod MethodName = "sendVideo"

type SendVideo struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Video               *string     `json:"video,omitempty"`
	VideoFile           *InputFile  `json:"-"`
	Duration            *int        `json:"duration,omitempty"`
	Width               *int        `json:"width,omitempty"`
	Height              *int        `json:"height,omitempty"`
	Caption             *string     `json:"caption,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVideoResult *Message

const SendVoiceMethod MethodName = "sendVoice"

type SendVoice struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Voice               *string     `json:"voice,omitempty"`
	VoiceFile           *InputFile  `json:"-"`
	Caption             *string     `json:"caption,omitempty"`
	Duration            *int        `json:"duration,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVoiceResult *Message

const SendLocationMethod MethodName = "sendLocation"

type SendLocation struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendLocationResult *Message

const SendVenueMethod MethodName = "sendVenue"

type SendVenue struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	Title               string      `json:"title"`
	Address             string      `json:"address"`
	FoursquareID        *string     `json:"foursquare_id,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVenueResult *Message

const SendContactMethod MethodName = "sendContact"

type SendContact struct {
	Method              *MethodName `json:"method,omitempty"`
	ChatID              ChatID      `json:"chat_id"`
	PhoneNumber         string      `json:"phone_number"`
	FirstName           string      `json:"first_name"`
	LastName            *string     `json:"last_name,omitempty"`
	DisableNotification *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendContactResult *Message

const SendChatActionMethod MethodName = "sendChatAction"

type SendChatAction struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
	Action ChatAction  `json:"action"`
}

type ChatAction string

const (
	ChatActionTyping         ChatAction = "typing"
	ChatActionUploadPhoto    ChatAction = "upload_photo"
	ChatActionRecordVideo    ChatAction = "record_video"
	ChatActionUploadVideo    ChatAction = "upload_video"
	ChatActionRecordAudio    ChatAction = "record_audio"
	ChatActionUploadAudio    ChatAction = "upload_audio"
	ChatActionUploadDocument ChatAction = "upload_document"
	ChatActionFindLocation   ChatAction = "find_location"
)

type SendChatActionResult bool

const GetUserProfilePhotosMethod MethodName = "getUserProfilePhotos"

type GetUserProfilePhotos struct {
	Method *MethodName `json:"method,omitempty"`
	UserID int         `json:"user_id"`
	Offset *int        `json:"offset,omitempty"`
	Limit  *int        `json:"limit,omitempty"`
}

type GetUserProfilePhotosResult *UserProfilePhotos

const GetFileMethod MethodName = "getFile"

type GetFile struct {
	Method *MethodName `json:"method,omitempty"`
	FileID string      `json:"file_id"`
}

type GetFileResult *File

const KickChatMemberMethod MethodName = "kickChatMember"

type KickChatMember struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
	UserID int         `json:"user_id"`
}

type KickChatMemberResult bool

const LeaveChatMethod MethodName = "leaveChat"

type LeaveChat struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
}

type LeaveChatResult bool

const UnbanChatMemberMethod MethodName = "unbanChatMember"

type UnbanChatMember struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
	UserID int         `json:"user_id"`
}

type UnbanChatMemberResult bool

const GetChatMethod MethodName = "getChat"

type GetChat struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
}

type GetChatResult *Chat

const GetChatAdministratorsMethod MethodName = "getChatAdministrators"

type GetChatAdministrators struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
}

type GetChatAdministratorsResult []*ChatMember

const GetChatMembersCountMethod MethodName = "getChatMembersCount"

type GetChatMembersCount struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
}

type GetChatMembersCountResult int

const GetChatMemberMethod MethodName = "getChatMember"

type GetChatMember struct {
	Method *MethodName `json:"method,omitempty"`
	ChatID ChatID      `json:"chat_id"`
	UserID int         `json:"user_id"`
}

type GetChatMemberResult *ChatMember

const AnswerCallbackQueryMethod MethodName = "answerCallbackQuery"

type AnswerCallbackQuery struct {
	Method          *MethodName `json:"method,omitempty"`
	CallbackQueryID string      `json:"callback_query_id"`
	Text            *string     `json:"text,omitempty"`
	ShowAlert       *bool       `json:"show_alert,omitempty"`
	URL             *string     `json:"url,omitempty"`
	CacheTime       *int        `json:"cache_time,omitempty"`
}

type AnswerCallbackQueryResult bool
