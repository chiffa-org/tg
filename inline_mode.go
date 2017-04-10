package tg

type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

const AnswerInlineQueryMethod MethodName = "answerInlineQuery"

type AnswerInlineQuery struct {
	Method            *MethodName         `json:"method,omitempty"`
	InlineQueryID     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"results"`
	CacheTime         *int                `json:"cache_time,omitempty"`
	IsPersonal        *bool               `json:"is_personal,omitempty"`
	NextOffset        *string             `json:"next_offset,omitempty"`
	SwitchPMText      *string             `json:"switch_pm_text,omitempty"`
	SwitchPMParameter *string             `json:"switch_pm_parameter,omitempty"`
}

type AnswerInlineQueryResult bool

type InlineQueryResult interface {
	inlineQueryResult()
}

type InlineQueryResultArticle struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	HideURL             *bool                 `json:"hide_url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

func (*InlineQueryResultArticle) inlineQueryResult() {}

type InlineQueryResultPhoto struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	PhotoURL            string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	PhotoWidth          *int                  `json:"photo_width,omitempty"`
	PhotoHeight         *int                  `json:"photo_height,omitempty"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultPhoto) inlineQueryResult() {}

type InlineQueryResultGIF struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	GIFURL              string                `json:"gif_url"`
	GIFWidth            *int                  `json:"gif_width,omitempty"`
	GIFHeight           *int                  `json:"gif_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultGIF) inlineQueryResult() {}

type InlineQueryResultMPEG4GIF struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	MPEG4URL            string                `json:"mpeg4_url"`
	MPEG4Width          *int                  `json:"mpeg4_width,omitempty"`
	MPEG4Height         *int                  `json:"mpeg4_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultMPEG4GIF) inlineQueryResult() {}

type InlineQueryResultVideo struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	VideoURL            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VideoWidth          *int                  `json:"video_width,omitempty"`
	VideoHeight         *int                  `json:"video_height,omitempty"`
	VideoDuration       *int                  `json:"video_duration,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultVideo) inlineQueryResult() {}

type InlineQueryResultAudio struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int                  `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultAudio) inlineQueryResult() {}

type InlineQueryResultVoice struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VoiceDuration       *int                  `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultVoice) inlineQueryResult() {}

type InlineQueryResultDocument struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

func (*InlineQueryResultDocument) inlineQueryResult() {}

type InlineQueryResultLocation struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

func (*InlineQueryResultLocation) inlineQueryResult() {}

type InlineQueryResultVenue struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        *string               `json:"foursquare_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

func (*InlineQueryResultVenue) inlineQueryResult() {}

type InlineQueryResultContact struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          *int                  `json:"thumb_width,omitempty"`
	ThumbHeight         *int                  `json:"thumb_height,omitempty"`
}

func (*InlineQueryResultContact) inlineQueryResult() {}

type InlineQueryResultGame struct {
	Type          InlineQueryResultType `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (*InlineQueryResultGame) inlineQueryResult() {}

type InlineQueryResultCachedPhoto struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	PhotoFileID         string                `json:"photo_file_id"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedPhoto) inlineQueryResult() {}

type InlineQueryResultCachedGIF struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	GIFFileID           string                `json:"gif_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedGIF) inlineQueryResult() {}

type InlineQueryResultCachedMPEG4GIF struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	MPEG4FileID         string                `json:"mpeg4_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedMPEG4GIF) inlineQueryResult() {}

type InlineQueryResultCachedSticker struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedSticker) inlineQueryResult() {}

type InlineQueryResultCachedDocument struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocmentFileID       string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedDocument) inlineQueryResult() {}

type InlineQueryResultCachedVideo struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedVideo) inlineQueryResult() {}

type InlineQueryResultCachedVoice struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedVoice) inlineQueryResult() {}

type InlineQueryResultCachedAudio struct {
	Type                InlineQueryResultType `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedAudio) inlineQueryResult() {}

type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeGIF      InlineQueryResultType = "gif"
	InlineQueryResultTypeMPEG4GIF InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
)

type InputMessageContent interface {
	inputMessageContent()
}

type InputTextMessageContent struct {
	MessageText           string  `json:"message_text"`
	ParseMode             *string `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool   `json:"disable_web_page_preview,omitempty"`
}

func (*InputTextMessageContent) inputMessageContent() {}

type InputLocationMessageContent struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (*InputLocationMessageContent) inputMessageContent() {}

type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID *string `json:"foursquare_id,omitempty"`
}

func (*InputVenueMessageContent) inputMessageContent() {}

type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
}

func (*InputContactMessageContent) inputMessageContent() {}

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}
