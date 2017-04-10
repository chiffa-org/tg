package tg

import (
	"net/url"
)

type Request interface {
	MethodName() MethodName
	IncludeMethodName()
	Values() (url.Values, error)
}

type FileRequest interface {
	Request
	InputFiles() map[string]*InputFile
}

func (*GetUpdates) MethodName() MethodName             { return GetUpdatesMethod }
func (*SetWebhook) MethodName() MethodName             { return SetWebhookMethod }
func (*DeleteWebhook) MethodName() MethodName          { return DeleteWebhookMethod }
func (*GetWebhookInfo) MethodName() MethodName         { return GetWebhookInfoMethod }
func (*GetMe) MethodName() MethodName                  { return GetMeMethod }
func (*SendMessage) MethodName() MethodName            { return SendMessageMethod }
func (*ForwardMessage) MethodName() MethodName         { return ForwardMessageMethod }
func (*SendPhoto) MethodName() MethodName              { return SendPhotoMethod }
func (*SendAudio) MethodName() MethodName              { return SendAudioMethod }
func (*SendDocument) MethodName() MethodName           { return SendDocumentMethod }
func (*SendSticker) MethodName() MethodName            { return SendStickerMethod }
func (*SendVideo) MethodName() MethodName              { return SendVideoMethod }
func (*SendVoice) MethodName() MethodName              { return SendVoiceMethod }
func (*SendLocation) MethodName() MethodName           { return SendLocationMethod }
func (*SendVenue) MethodName() MethodName              { return SendVenueMethod }
func (*SendContact) MethodName() MethodName            { return SendContactMethod }
func (*SendChatAction) MethodName() MethodName         { return SendChatActionMethod }
func (*GetUserProfilePhotos) MethodName() MethodName   { return GetUserProfilePhotosMethod }
func (*GetFile) MethodName() MethodName                { return GetFileMethod }
func (*KickChatMember) MethodName() MethodName         { return KickChatMemberMethod }
func (*LeaveChat) MethodName() MethodName              { return LeaveChatMethod }
func (*UnbanChatMember) MethodName() MethodName        { return UnbanChatMemberMethod }
func (*GetChat) MethodName() MethodName                { return GetChatMethod }
func (*GetChatAdministrators) MethodName() MethodName  { return GetChatAdministratorsMethod }
func (*GetChatMembersCount) MethodName() MethodName    { return GetChatMembersCountMethod }
func (*GetChatMember) MethodName() MethodName          { return GetChatMemberMethod }
func (*AnswerCallbackQuery) MethodName() MethodName    { return AnswerCallbackQueryMethod }
func (*EditMessageText) MethodName() MethodName        { return EditMessageTextMethod }
func (*EditMessageCaption) MethodName() MethodName     { return EditMessageCaptionMethod }
func (*EditMessageReplyMarkup) MethodName() MethodName { return EditMessageReplyMarkupMethod }
func (*AnswerInlineQuery) MethodName() MethodName      { return AnswerInlineQueryMethod }
func (*SendGame) MethodName() MethodName               { return SendGameMethod }
func (*SetGameScore) MethodName() MethodName           { return SetGameScoreMethod }
func (*GetGameHighScores) MethodName() MethodName      { return GetGameHighScoresMethod }

func (t *GetUpdates) IncludeMethodName()             { t.SetMethod(GetUpdatesMethod) }
func (t *SetWebhook) IncludeMethodName()             { t.SetMethod(SetWebhookMethod) }
func (t *DeleteWebhook) IncludeMethodName()          { t.SetMethod(DeleteWebhookMethod) }
func (t *GetWebhookInfo) IncludeMethodName()         { t.SetMethod(GetWebhookInfoMethod) }
func (t *GetMe) IncludeMethodName()                  { t.SetMethod(GetMeMethod) }
func (t *SendMessage) IncludeMethodName()            { t.SetMethod(SendMessageMethod) }
func (t *ForwardMessage) IncludeMethodName()         { t.SetMethod(ForwardMessageMethod) }
func (t *SendPhoto) IncludeMethodName()              { t.SetMethod(SendPhotoMethod) }
func (t *SendAudio) IncludeMethodName()              { t.SetMethod(SendAudioMethod) }
func (t *SendDocument) IncludeMethodName()           { t.SetMethod(SendDocumentMethod) }
func (t *SendSticker) IncludeMethodName()            { t.SetMethod(SendStickerMethod) }
func (t *SendVideo) IncludeMethodName()              { t.SetMethod(SendVideoMethod) }
func (t *SendVoice) IncludeMethodName()              { t.SetMethod(SendVoiceMethod) }
func (t *SendLocation) IncludeMethodName()           { t.SetMethod(SendLocationMethod) }
func (t *SendVenue) IncludeMethodName()              { t.SetMethod(SendVenueMethod) }
func (t *SendContact) IncludeMethodName()            { t.SetMethod(SendContactMethod) }
func (t *SendChatAction) IncludeMethodName()         { t.SetMethod(SendChatActionMethod) }
func (t *GetUserProfilePhotos) IncludeMethodName()   { t.SetMethod(GetUserProfilePhotosMethod) }
func (t *GetFile) IncludeMethodName()                { t.SetMethod(GetFileMethod) }
func (t *KickChatMember) IncludeMethodName()         { t.SetMethod(KickChatMemberMethod) }
func (t *LeaveChat) IncludeMethodName()              { t.SetMethod(LeaveChatMethod) }
func (t *UnbanChatMember) IncludeMethodName()        { t.SetMethod(UnbanChatMemberMethod) }
func (t *GetChat) IncludeMethodName()                { t.SetMethod(GetChatMethod) }
func (t *GetChatAdministrators) IncludeMethodName()  { t.SetMethod(GetChatAdministratorsMethod) }
func (t *GetChatMembersCount) IncludeMethodName()    { t.SetMethod(GetChatMembersCountMethod) }
func (t *GetChatMember) IncludeMethodName()          { t.SetMethod(GetChatMemberMethod) }
func (t *AnswerCallbackQuery) IncludeMethodName()    { t.SetMethod(AnswerCallbackQueryMethod) }
func (t *EditMessageText) IncludeMethodName()        { t.SetMethod(EditMessageTextMethod) }
func (t *EditMessageCaption) IncludeMethodName()     { t.SetMethod(EditMessageCaptionMethod) }
func (t *EditMessageReplyMarkup) IncludeMethodName() { t.SetMethod(EditMessageReplyMarkupMethod) }
func (t *AnswerInlineQuery) IncludeMethodName()      { t.SetMethod(AnswerInlineQueryMethod) }
func (t *SendGame) IncludeMethodName()               { t.SetMethod(SendGameMethod) }
func (t *SetGameScore) IncludeMethodName()           { t.SetMethod(SetGameScoreMethod) }
func (t *GetGameHighScores) IncludeMethodName()      { t.SetMethod(GetGameHighScoresMethod) }

func (t *GetUpdates) Values() (url.Values, error)             { return values(t) }
func (t *SetWebhook) Values() (url.Values, error)             { return values(t) }
func (t *DeleteWebhook) Values() (url.Values, error)          { return values(t) }
func (t *GetWebhookInfo) Values() (url.Values, error)         { return values(t) }
func (t *GetMe) Values() (url.Values, error)                  { return values(t) }
func (t *SendMessage) Values() (url.Values, error)            { return values(t) }
func (t *ForwardMessage) Values() (url.Values, error)         { return values(t) }
func (t *SendPhoto) Values() (url.Values, error)              { return values(t) }
func (t *SendAudio) Values() (url.Values, error)              { return values(t) }
func (t *SendDocument) Values() (url.Values, error)           { return values(t) }
func (t *SendSticker) Values() (url.Values, error)            { return values(t) }
func (t *SendVideo) Values() (url.Values, error)              { return values(t) }
func (t *SendVoice) Values() (url.Values, error)              { return values(t) }
func (t *SendLocation) Values() (url.Values, error)           { return values(t) }
func (t *SendVenue) Values() (url.Values, error)              { return values(t) }
func (t *SendContact) Values() (url.Values, error)            { return values(t) }
func (t *SendChatAction) Values() (url.Values, error)         { return values(t) }
func (t *GetUserProfilePhotos) Values() (url.Values, error)   { return values(t) }
func (t *GetFile) Values() (url.Values, error)                { return values(t) }
func (t *KickChatMember) Values() (url.Values, error)         { return values(t) }
func (t *LeaveChat) Values() (url.Values, error)              { return values(t) }
func (t *UnbanChatMember) Values() (url.Values, error)        { return values(t) }
func (t *GetChat) Values() (url.Values, error)                { return values(t) }
func (t *GetChatAdministrators) Values() (url.Values, error)  { return values(t) }
func (t *GetChatMembersCount) Values() (url.Values, error)    { return values(t) }
func (t *GetChatMember) Values() (url.Values, error)          { return values(t) }
func (t *AnswerCallbackQuery) Values() (url.Values, error)    { return values(t) }
func (t *EditMessageText) Values() (url.Values, error)        { return values(t) }
func (t *EditMessageCaption) Values() (url.Values, error)     { return values(t) }
func (t *EditMessageReplyMarkup) Values() (url.Values, error) { return values(t) }
func (t *AnswerInlineQuery) Values() (url.Values, error)      { return values(t) }
func (t *SendGame) Values() (url.Values, error)               { return values(t) }
func (t *SetGameScore) Values() (url.Values, error)           { return values(t) }
func (t *GetGameHighScores) Values() (url.Values, error)      { return values(t) }

func (t *SetWebhook) InputFiles() map[string]*InputFile {
	if t.CertificateFile == nil {
		return nil
	}
	return map[string]*InputFile{"certificate": t.CertificateFile}
}

func (t *SendPhoto) InputFiles() map[string]*InputFile {
	if t.PhotoFile == nil {
		return nil
	}
	return map[string]*InputFile{"photo": t.PhotoFile}
}

func (t *SendAudio) InputFiles() map[string]*InputFile {
	if t.AudioFile == nil {
		return nil
	}
	return map[string]*InputFile{"audio": t.AudioFile}
}

func (t *SendDocument) InputFiles() map[string]*InputFile {
	if t.DocumentFile == nil {
		return nil
	}
	return map[string]*InputFile{"document": t.DocumentFile}
}

func (t *SendSticker) InputFiles() map[string]*InputFile {
	if t.StickerFile == nil {
		return nil
	}
	return map[string]*InputFile{"sticker": t.StickerFile}
}

func (t *SendVideo) InputFiles() map[string]*InputFile {
	if t.VideoFile == nil {
		return nil
	}
	return map[string]*InputFile{"video": t.VideoFile}
}

func (t *SendVoice) InputFiles() map[string]*InputFile {
	if t.VoiceFile == nil {
		return nil
	}
	return map[string]*InputFile{"voice": t.VoiceFile}
}
