package tg

import "fmt"

func (t *Update) GetUpdateID() int { return t.UpdateID }
func (t *Update) SetUpdateID(v int) { t.UpdateID = v }
func (t *Update) WithUpdateID(v int) *Update { t.UpdateID = v; return t }
func (t *Update) GetMessage() *Message { return t.Message }
func (t *Update) SetMessage(v *Message) { t.Message = v }
func (t *Update) WithMessage(v *Message) *Update { t.Message = v; return t }
func (t *Update) GetEditedMessage() *Message { return t.EditedMessage }
func (t *Update) SetEditedMessage(v *Message) { t.EditedMessage = v }
func (t *Update) WithEditedMessage(v *Message) *Update { t.EditedMessage = v; return t }
func (t *Update) GetChannelPost() *Message { return t.ChannelPost }
func (t *Update) SetChannelPost(v *Message) { t.ChannelPost = v }
func (t *Update) WithChannelPost(v *Message) *Update { t.ChannelPost = v; return t }
func (t *Update) GetEditedChannelPost() *Message { return t.EditedChannelPost }
func (t *Update) SetEditedChannelPost(v *Message) { t.EditedChannelPost = v }
func (t *Update) WithEditedChannelPost(v *Message) *Update { t.EditedChannelPost = v; return t }
func (t *Update) GetInlineQuery() *InlineQuery { return t.InlineQuery }
func (t *Update) SetInlineQuery(v *InlineQuery) { t.InlineQuery = v }
func (t *Update) WithInlineQuery(v *InlineQuery) *Update { t.InlineQuery = v; return t }
func (t *Update) GetChosenInlineResult() *ChosenInlineResult { return t.ChosenInlineResult }
func (t *Update) SetChosenInlineResult(v *ChosenInlineResult) { t.ChosenInlineResult = v }
func (t *Update) WithChosenInlineResult(v *ChosenInlineResult) *Update { t.ChosenInlineResult = v; return t }
func (t *Update) GetCallbackQuery() *CallbackQuery { return t.CallbackQuery }
func (t *Update) SetCallbackQuery(v *CallbackQuery) { t.CallbackQuery = v }
func (t *Update) WithCallbackQuery(v *CallbackQuery) *Update { t.CallbackQuery = v; return t }

func (t *GetUpdates) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetUpdates) SetMethod(v MethodName) { t.Method = &v }
func (t *GetUpdates) WithMethod(v MethodName) *GetUpdates { t.Method = &v; return t }
func (t *GetUpdates) SetNilMethod() { t.Method = nil }
func (t *GetUpdates) WithNilMethod() *GetUpdates { t.Method = nil; return t }
func (t *GetUpdates) GetOffset() int { if t.Offset == nil { return 0 } else { return *t.Offset } }
func (t *GetUpdates) SetOffset(v int) { t.Offset = &v }
func (t *GetUpdates) WithOffset(v int) *GetUpdates { t.Offset = &v; return t }
func (t *GetUpdates) SetNilOffset() { t.Offset = nil }
func (t *GetUpdates) WithNilOffset() *GetUpdates { t.Offset = nil; return t }
func (t *GetUpdates) GetLimit() int { if t.Limit == nil { return 0 } else { return *t.Limit } }
func (t *GetUpdates) SetLimit(v int) { t.Limit = &v }
func (t *GetUpdates) WithLimit(v int) *GetUpdates { t.Limit = &v; return t }
func (t *GetUpdates) SetNilLimit() { t.Limit = nil }
func (t *GetUpdates) WithNilLimit() *GetUpdates { t.Limit = nil; return t }
func (t *GetUpdates) GetTimeout() int { if t.Timeout == nil { return 0 } else { return *t.Timeout } }
func (t *GetUpdates) SetTimeout(v int) { t.Timeout = &v }
func (t *GetUpdates) WithTimeout(v int) *GetUpdates { t.Timeout = &v; return t }
func (t *GetUpdates) SetNilTimeout() { t.Timeout = nil }
func (t *GetUpdates) WithNilTimeout() *GetUpdates { t.Timeout = nil; return t }
func (t *GetUpdates) GetAllowedUpdates() []UpdateType { if t.AllowedUpdates == nil { return nil } else { return *t.AllowedUpdates } }
func (t *GetUpdates) SetAllowedUpdates(v []UpdateType) { if v == nil { t.AllowedUpdates = nil } else { t.AllowedUpdates = &v } }
func (t *GetUpdates) WithAllowedUpdates(v []UpdateType) *GetUpdates { if v == nil { t.AllowedUpdates = nil } else { t.AllowedUpdates = &v }; return t }

func (t *SetWebhook) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SetWebhook) SetMethod(v MethodName) { t.Method = &v }
func (t *SetWebhook) WithMethod(v MethodName) *SetWebhook { t.Method = &v; return t }
func (t *SetWebhook) SetNilMethod() { t.Method = nil }
func (t *SetWebhook) WithNilMethod() *SetWebhook { t.Method = nil; return t }
func (t *SetWebhook) GetURL() string { return t.URL }
func (t *SetWebhook) SetURL(v string) { t.URL = v }
func (t *SetWebhook) WithURL(v string) *SetWebhook { t.URL = v; return t }
func (t *SetWebhook) SetFmtURL(f string, a ...interface{}) { t.SetURL(fmt.Sprintf(f, a...)) }
func (t *SetWebhook) WithFmtURL(f string, a ...interface{}) *SetWebhook { return t.WithURL(fmt.Sprintf(f, a...)) }
func (t *SetWebhook) GetCertificateFile() *InputFile { return t.CertificateFile }
func (t *SetWebhook) SetCertificateFile(v *InputFile) { t.CertificateFile = v }
func (t *SetWebhook) WithCertificateFile(v *InputFile) *SetWebhook { t.CertificateFile = v; return t }
func (t *SetWebhook) GetMaxConnections() int { if t.MaxConnections == nil { return 0 } else { return *t.MaxConnections } }
func (t *SetWebhook) SetMaxConnections(v int) { t.MaxConnections = &v }
func (t *SetWebhook) WithMaxConnections(v int) *SetWebhook { t.MaxConnections = &v; return t }
func (t *SetWebhook) SetNilMaxConnections() { t.MaxConnections = nil }
func (t *SetWebhook) WithNilMaxConnections() *SetWebhook { t.MaxConnections = nil; return t }
func (t *SetWebhook) GetAllowedUpdates() []UpdateType { if t.AllowedUpdates == nil { return nil } else { return *t.AllowedUpdates } }
func (t *SetWebhook) SetAllowedUpdates(v []UpdateType) { if v == nil { t.AllowedUpdates = nil } else { t.AllowedUpdates = &v } }
func (t *SetWebhook) WithAllowedUpdates(v []UpdateType) *SetWebhook { if v == nil { t.AllowedUpdates = nil } else { t.AllowedUpdates = &v }; return t }

func (t *DeleteWebhook) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *DeleteWebhook) SetMethod(v MethodName) { t.Method = &v }
func (t *DeleteWebhook) WithMethod(v MethodName) *DeleteWebhook { t.Method = &v; return t }
func (t *DeleteWebhook) SetNilMethod() { t.Method = nil }
func (t *DeleteWebhook) WithNilMethod() *DeleteWebhook { t.Method = nil; return t }

func (t *GetWebhookInfo) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetWebhookInfo) SetMethod(v MethodName) { t.Method = &v }
func (t *GetWebhookInfo) WithMethod(v MethodName) *GetWebhookInfo { t.Method = &v; return t }
func (t *GetWebhookInfo) SetNilMethod() { t.Method = nil }
func (t *GetWebhookInfo) WithNilMethod() *GetWebhookInfo { t.Method = nil; return t }

func (t *WebhookInfo) GetURL() string { return t.URL }
func (t *WebhookInfo) SetURL(v string) { t.URL = v }
func (t *WebhookInfo) WithURL(v string) *WebhookInfo { t.URL = v; return t }
func (t *WebhookInfo) SetFmtURL(f string, a ...interface{}) { t.SetURL(fmt.Sprintf(f, a...)) }
func (t *WebhookInfo) WithFmtURL(f string, a ...interface{}) *WebhookInfo { return t.WithURL(fmt.Sprintf(f, a...)) }
func (t *WebhookInfo) GetHasCustomCertificate() bool { return t.HasCustomCertificate }
func (t *WebhookInfo) SetHasCustomCertificate(v bool) { t.HasCustomCertificate = v }
func (t *WebhookInfo) WithHasCustomCertificate(v bool) *WebhookInfo { t.HasCustomCertificate = v; return t }
func (t *WebhookInfo) GetPendingUpdateCount() int { return t.PendingUpdateCount }
func (t *WebhookInfo) SetPendingUpdateCount(v int) { t.PendingUpdateCount = v }
func (t *WebhookInfo) WithPendingUpdateCount(v int) *WebhookInfo { t.PendingUpdateCount = v; return t }
func (t *WebhookInfo) GetLastErrorDate() int { if t.LastErrorDate == nil { return 0 } else { return *t.LastErrorDate } }
func (t *WebhookInfo) SetLastErrorDate(v int) { t.LastErrorDate = &v }
func (t *WebhookInfo) WithLastErrorDate(v int) *WebhookInfo { t.LastErrorDate = &v; return t }
func (t *WebhookInfo) SetNilLastErrorDate() { t.LastErrorDate = nil }
func (t *WebhookInfo) WithNilLastErrorDate() *WebhookInfo { t.LastErrorDate = nil; return t }
func (t *WebhookInfo) GetLastErrorMessage() string { if t.LastErrorMessage == nil { return "" } else { return *t.LastErrorMessage } }
func (t *WebhookInfo) SetLastErrorMessage(v string) { t.LastErrorMessage = &v }
func (t *WebhookInfo) WithLastErrorMessage(v string) *WebhookInfo { t.LastErrorMessage = &v; return t }
func (t *WebhookInfo) SetNilLastErrorMessage() { t.LastErrorMessage = nil }
func (t *WebhookInfo) WithNilLastErrorMessage() *WebhookInfo { t.LastErrorMessage = nil; return t }
func (t *WebhookInfo) SetFmtLastErrorMessage(f string, a ...interface{}) { t.SetLastErrorMessage(fmt.Sprintf(f, a...)) }
func (t *WebhookInfo) WithFmtLastErrorMessage(f string, a ...interface{}) *WebhookInfo { return t.WithLastErrorMessage(fmt.Sprintf(f, a...)) }
func (t *WebhookInfo) GetMaxConnections() int { if t.MaxConnections == nil { return 0 } else { return *t.MaxConnections } }
func (t *WebhookInfo) SetMaxConnections(v int) { t.MaxConnections = &v }
func (t *WebhookInfo) WithMaxConnections(v int) *WebhookInfo { t.MaxConnections = &v; return t }
func (t *WebhookInfo) SetNilMaxConnections() { t.MaxConnections = nil }
func (t *WebhookInfo) WithNilMaxConnections() *WebhookInfo { t.MaxConnections = nil; return t }
func (t *WebhookInfo) GetAllowedUpdates() []UpdateType { if t.AllowedUpdates == nil { return nil } else { return *t.AllowedUpdates } }
func (t *WebhookInfo) SetAllowedUpdates(v []UpdateType) { if v == nil { t.AllowedUpdates = nil } else { t.AllowedUpdates = &v } }
func (t *WebhookInfo) WithAllowedUpdates(v []UpdateType) *WebhookInfo { if v == nil { t.AllowedUpdates = nil } else { t.AllowedUpdates = &v }; return t }

func (t *User) GetID() int { return t.ID }
func (t *User) SetID(v int) { t.ID = v }
func (t *User) WithID(v int) *User { t.ID = v; return t }
func (t *User) GetFirstName() string { return t.FirstName }
func (t *User) SetFirstName(v string) { t.FirstName = v }
func (t *User) WithFirstName(v string) *User { t.FirstName = v; return t }
func (t *User) SetFmtFirstName(f string, a ...interface{}) { t.SetFirstName(fmt.Sprintf(f, a...)) }
func (t *User) WithFmtFirstName(f string, a ...interface{}) *User { return t.WithFirstName(fmt.Sprintf(f, a...)) }
func (t *User) GetLastName() string { if t.LastName == nil { return "" } else { return *t.LastName } }
func (t *User) SetLastName(v string) { t.LastName = &v }
func (t *User) WithLastName(v string) *User { t.LastName = &v; return t }
func (t *User) SetNilLastName() { t.LastName = nil }
func (t *User) WithNilLastName() *User { t.LastName = nil; return t }
func (t *User) SetFmtLastName(f string, a ...interface{}) { t.SetLastName(fmt.Sprintf(f, a...)) }
func (t *User) WithFmtLastName(f string, a ...interface{}) *User { return t.WithLastName(fmt.Sprintf(f, a...)) }
func (t *User) GetUserName() string { if t.UserName == nil { return "" } else { return *t.UserName } }
func (t *User) SetUserName(v string) { t.UserName = &v }
func (t *User) WithUserName(v string) *User { t.UserName = &v; return t }
func (t *User) SetNilUserName() { t.UserName = nil }
func (t *User) WithNilUserName() *User { t.UserName = nil; return t }
func (t *User) SetFmtUserName(f string, a ...interface{}) { t.SetUserName(fmt.Sprintf(f, a...)) }
func (t *User) WithFmtUserName(f string, a ...interface{}) *User { return t.WithUserName(fmt.Sprintf(f, a...)) }

func (t *Chat) GetID() int64 { return t.ID }
func (t *Chat) SetID(v int64) { t.ID = v }
func (t *Chat) WithID(v int64) *Chat { t.ID = v; return t }
func (t *Chat) GetType() ChatType { return t.Type }
func (t *Chat) SetType(v ChatType) { t.Type = v }
func (t *Chat) WithType(v ChatType) *Chat { t.Type = v; return t }
func (t *Chat) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *Chat) SetTitle(v string) { t.Title = &v }
func (t *Chat) WithTitle(v string) *Chat { t.Title = &v; return t }
func (t *Chat) SetNilTitle() { t.Title = nil }
func (t *Chat) WithNilTitle() *Chat { t.Title = nil; return t }
func (t *Chat) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *Chat) WithFmtTitle(f string, a ...interface{}) *Chat { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *Chat) GetUserName() string { if t.UserName == nil { return "" } else { return *t.UserName } }
func (t *Chat) SetUserName(v string) { t.UserName = &v }
func (t *Chat) WithUserName(v string) *Chat { t.UserName = &v; return t }
func (t *Chat) SetNilUserName() { t.UserName = nil }
func (t *Chat) WithNilUserName() *Chat { t.UserName = nil; return t }
func (t *Chat) SetFmtUserName(f string, a ...interface{}) { t.SetUserName(fmt.Sprintf(f, a...)) }
func (t *Chat) WithFmtUserName(f string, a ...interface{}) *Chat { return t.WithUserName(fmt.Sprintf(f, a...)) }
func (t *Chat) GetFirstName() string { if t.FirstName == nil { return "" } else { return *t.FirstName } }
func (t *Chat) SetFirstName(v string) { t.FirstName = &v }
func (t *Chat) WithFirstName(v string) *Chat { t.FirstName = &v; return t }
func (t *Chat) SetNilFirstName() { t.FirstName = nil }
func (t *Chat) WithNilFirstName() *Chat { t.FirstName = nil; return t }
func (t *Chat) SetFmtFirstName(f string, a ...interface{}) { t.SetFirstName(fmt.Sprintf(f, a...)) }
func (t *Chat) WithFmtFirstName(f string, a ...interface{}) *Chat { return t.WithFirstName(fmt.Sprintf(f, a...)) }
func (t *Chat) GetLastName() string { if t.LastName == nil { return "" } else { return *t.LastName } }
func (t *Chat) SetLastName(v string) { t.LastName = &v }
func (t *Chat) WithLastName(v string) *Chat { t.LastName = &v; return t }
func (t *Chat) SetNilLastName() { t.LastName = nil }
func (t *Chat) WithNilLastName() *Chat { t.LastName = nil; return t }
func (t *Chat) SetFmtLastName(f string, a ...interface{}) { t.SetLastName(fmt.Sprintf(f, a...)) }
func (t *Chat) WithFmtLastName(f string, a ...interface{}) *Chat { return t.WithLastName(fmt.Sprintf(f, a...)) }
func (t *Chat) GetAllMembersAreAdministrators() bool { if t.AllMembersAreAdministrators == nil { return false } else { return *t.AllMembersAreAdministrators } }
func (t *Chat) SetAllMembersAreAdministrators(v bool) { t.AllMembersAreAdministrators = &v }
func (t *Chat) WithAllMembersAreAdministrators(v bool) *Chat { t.AllMembersAreAdministrators = &v; return t }
func (t *Chat) SetNilAllMembersAreAdministrators() { t.AllMembersAreAdministrators = nil }
func (t *Chat) WithNilAllMembersAreAdministrators() *Chat { t.AllMembersAreAdministrators = nil; return t }

func (t *Message) GetMessageID() int { return t.MessageID }
func (t *Message) SetMessageID(v int) { t.MessageID = v }
func (t *Message) WithMessageID(v int) *Message { t.MessageID = v; return t }
func (t *Message) GetFrom() *User { return t.From }
func (t *Message) SetFrom(v *User) { t.From = v }
func (t *Message) WithFrom(v *User) *Message { t.From = v; return t }
func (t *Message) GetDate() int { return t.Date }
func (t *Message) SetDate(v int) { t.Date = v }
func (t *Message) WithDate(v int) *Message { t.Date = v; return t }
func (t *Message) GetChat() *Chat { return t.Chat }
func (t *Message) SetChat(v *Chat) { t.Chat = v }
func (t *Message) WithChat(v *Chat) *Message { t.Chat = v; return t }
func (t *Message) GetForwardFrom() *User { return t.ForwardFrom }
func (t *Message) SetForwardFrom(v *User) { t.ForwardFrom = v }
func (t *Message) WithForwardFrom(v *User) *Message { t.ForwardFrom = v; return t }
func (t *Message) GetForwardFromChat() *Chat { return t.ForwardFromChat }
func (t *Message) SetForwardFromChat(v *Chat) { t.ForwardFromChat = v }
func (t *Message) WithForwardFromChat(v *Chat) *Message { t.ForwardFromChat = v; return t }
func (t *Message) GetForwardFromMessageID() int { if t.ForwardFromMessageID == nil { return 0 } else { return *t.ForwardFromMessageID } }
func (t *Message) SetForwardFromMessageID(v int) { t.ForwardFromMessageID = &v }
func (t *Message) WithForwardFromMessageID(v int) *Message { t.ForwardFromMessageID = &v; return t }
func (t *Message) SetNilForwardFromMessageID() { t.ForwardFromMessageID = nil }
func (t *Message) WithNilForwardFromMessageID() *Message { t.ForwardFromMessageID = nil; return t }
func (t *Message) GetForwardDate() int { if t.ForwardDate == nil { return 0 } else { return *t.ForwardDate } }
func (t *Message) SetForwardDate(v int) { t.ForwardDate = &v }
func (t *Message) WithForwardDate(v int) *Message { t.ForwardDate = &v; return t }
func (t *Message) SetNilForwardDate() { t.ForwardDate = nil }
func (t *Message) WithNilForwardDate() *Message { t.ForwardDate = nil; return t }
func (t *Message) GetReplyToMessage() *Message { return t.ReplyToMessage }
func (t *Message) SetReplyToMessage(v *Message) { t.ReplyToMessage = v }
func (t *Message) WithReplyToMessage(v *Message) *Message { t.ReplyToMessage = v; return t }
func (t *Message) GetEditDate() int { if t.EditDate == nil { return 0 } else { return *t.EditDate } }
func (t *Message) SetEditDate(v int) { t.EditDate = &v }
func (t *Message) WithEditDate(v int) *Message { t.EditDate = &v; return t }
func (t *Message) SetNilEditDate() { t.EditDate = nil }
func (t *Message) WithNilEditDate() *Message { t.EditDate = nil; return t }
func (t *Message) GetText() string { if t.Text == nil { return "" } else { return *t.Text } }
func (t *Message) SetText(v string) { t.Text = &v }
func (t *Message) WithText(v string) *Message { t.Text = &v; return t }
func (t *Message) SetNilText() { t.Text = nil }
func (t *Message) WithNilText() *Message { t.Text = nil; return t }
func (t *Message) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *Message) WithFmtText(f string, a ...interface{}) *Message { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *Message) GetEntities() []*MessageEntity { if t.Entities == nil { return nil } else { return *t.Entities } }
func (t *Message) SetEntities(v []*MessageEntity) { if v == nil { t.Entities = nil } else { t.Entities = &v } }
func (t *Message) WithEntities(v []*MessageEntity) *Message { if v == nil { t.Entities = nil } else { t.Entities = &v }; return t }
func (t *Message) GetAudio() *Audio { return t.Audio }
func (t *Message) SetAudio(v *Audio) { t.Audio = v }
func (t *Message) WithAudio(v *Audio) *Message { t.Audio = v; return t }
func (t *Message) GetDocument() *Document { return t.Document }
func (t *Message) SetDocument(v *Document) { t.Document = v }
func (t *Message) WithDocument(v *Document) *Message { t.Document = v; return t }
func (t *Message) GetGame() *Game { return t.Game }
func (t *Message) SetGame(v *Game) { t.Game = v }
func (t *Message) WithGame(v *Game) *Message { t.Game = v; return t }
func (t *Message) GetPhoto() []*PhotoSize { if t.Photo == nil { return nil } else { return *t.Photo } }
func (t *Message) SetPhoto(v []*PhotoSize) { if v == nil { t.Photo = nil } else { t.Photo = &v } }
func (t *Message) WithPhoto(v []*PhotoSize) *Message { if v == nil { t.Photo = nil } else { t.Photo = &v }; return t }
func (t *Message) GetSticker() *Sticker { return t.Sticker }
func (t *Message) SetSticker(v *Sticker) { t.Sticker = v }
func (t *Message) WithSticker(v *Sticker) *Message { t.Sticker = v; return t }
func (t *Message) GetVideo() *Video { return t.Video }
func (t *Message) SetVideo(v *Video) { t.Video = v }
func (t *Message) WithVideo(v *Video) *Message { t.Video = v; return t }
func (t *Message) GetVoice() *Voice { return t.Voice }
func (t *Message) SetVoice(v *Voice) { t.Voice = v }
func (t *Message) WithVoice(v *Voice) *Message { t.Voice = v; return t }
func (t *Message) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *Message) SetCaption(v string) { t.Caption = &v }
func (t *Message) WithCaption(v string) *Message { t.Caption = &v; return t }
func (t *Message) SetNilCaption() { t.Caption = nil }
func (t *Message) WithNilCaption() *Message { t.Caption = nil; return t }
func (t *Message) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *Message) WithFmtCaption(f string, a ...interface{}) *Message { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *Message) GetContact() *Contact { return t.Contact }
func (t *Message) SetContact(v *Contact) { t.Contact = v }
func (t *Message) WithContact(v *Contact) *Message { t.Contact = v; return t }
func (t *Message) GetLocation() *Location { return t.Location }
func (t *Message) SetLocation(v *Location) { t.Location = v }
func (t *Message) WithLocation(v *Location) *Message { t.Location = v; return t }
func (t *Message) GetVenue() *Venue { return t.Venue }
func (t *Message) SetVenue(v *Venue) { t.Venue = v }
func (t *Message) WithVenue(v *Venue) *Message { t.Venue = v; return t }
func (t *Message) GetNewChatMember() *User { return t.NewChatMember }
func (t *Message) SetNewChatMember(v *User) { t.NewChatMember = v }
func (t *Message) WithNewChatMember(v *User) *Message { t.NewChatMember = v; return t }
func (t *Message) GetLeftChatMember() *User { return t.LeftChatMember }
func (t *Message) SetLeftChatMember(v *User) { t.LeftChatMember = v }
func (t *Message) WithLeftChatMember(v *User) *Message { t.LeftChatMember = v; return t }
func (t *Message) GetNewChatTitle() string { if t.NewChatTitle == nil { return "" } else { return *t.NewChatTitle } }
func (t *Message) SetNewChatTitle(v string) { t.NewChatTitle = &v }
func (t *Message) WithNewChatTitle(v string) *Message { t.NewChatTitle = &v; return t }
func (t *Message) SetNilNewChatTitle() { t.NewChatTitle = nil }
func (t *Message) WithNilNewChatTitle() *Message { t.NewChatTitle = nil; return t }
func (t *Message) SetFmtNewChatTitle(f string, a ...interface{}) { t.SetNewChatTitle(fmt.Sprintf(f, a...)) }
func (t *Message) WithFmtNewChatTitle(f string, a ...interface{}) *Message { return t.WithNewChatTitle(fmt.Sprintf(f, a...)) }
func (t *Message) GetNewChatPhoto() []*PhotoSize { if t.NewChatPhoto == nil { return nil } else { return *t.NewChatPhoto } }
func (t *Message) SetNewChatPhoto(v []*PhotoSize) { if v == nil { t.NewChatPhoto = nil } else { t.NewChatPhoto = &v } }
func (t *Message) WithNewChatPhoto(v []*PhotoSize) *Message { if v == nil { t.NewChatPhoto = nil } else { t.NewChatPhoto = &v }; return t }
func (t *Message) GetDeleteChatPhoto() bool { if t.DeleteChatPhoto == nil { return false } else { return *t.DeleteChatPhoto } }
func (t *Message) SetDeleteChatPhoto(v bool) { t.DeleteChatPhoto = &v }
func (t *Message) WithDeleteChatPhoto(v bool) *Message { t.DeleteChatPhoto = &v; return t }
func (t *Message) SetNilDeleteChatPhoto() { t.DeleteChatPhoto = nil }
func (t *Message) WithNilDeleteChatPhoto() *Message { t.DeleteChatPhoto = nil; return t }
func (t *Message) GetGroupChatCreated() bool { if t.GroupChatCreated == nil { return false } else { return *t.GroupChatCreated } }
func (t *Message) SetGroupChatCreated(v bool) { t.GroupChatCreated = &v }
func (t *Message) WithGroupChatCreated(v bool) *Message { t.GroupChatCreated = &v; return t }
func (t *Message) SetNilGroupChatCreated() { t.GroupChatCreated = nil }
func (t *Message) WithNilGroupChatCreated() *Message { t.GroupChatCreated = nil; return t }
func (t *Message) GetSuperGroupChatCreated() bool { if t.SuperGroupChatCreated == nil { return false } else { return *t.SuperGroupChatCreated } }
func (t *Message) SetSuperGroupChatCreated(v bool) { t.SuperGroupChatCreated = &v }
func (t *Message) WithSuperGroupChatCreated(v bool) *Message { t.SuperGroupChatCreated = &v; return t }
func (t *Message) SetNilSuperGroupChatCreated() { t.SuperGroupChatCreated = nil }
func (t *Message) WithNilSuperGroupChatCreated() *Message { t.SuperGroupChatCreated = nil; return t }
func (t *Message) GetChannelChatCreated() bool { if t.ChannelChatCreated == nil { return false } else { return *t.ChannelChatCreated } }
func (t *Message) SetChannelChatCreated(v bool) { t.ChannelChatCreated = &v }
func (t *Message) WithChannelChatCreated(v bool) *Message { t.ChannelChatCreated = &v; return t }
func (t *Message) SetNilChannelChatCreated() { t.ChannelChatCreated = nil }
func (t *Message) WithNilChannelChatCreated() *Message { t.ChannelChatCreated = nil; return t }
func (t *Message) GetMigrateToChatID() int64 { if t.MigrateToChatID == nil { return 0 } else { return *t.MigrateToChatID } }
func (t *Message) SetMigrateToChatID(v int64) { t.MigrateToChatID = &v }
func (t *Message) WithMigrateToChatID(v int64) *Message { t.MigrateToChatID = &v; return t }
func (t *Message) SetNilMigrateToChatID() { t.MigrateToChatID = nil }
func (t *Message) WithNilMigrateToChatID() *Message { t.MigrateToChatID = nil; return t }
func (t *Message) GetMigrateFromChatID() int64 { if t.MigrateFromChatID == nil { return 0 } else { return *t.MigrateFromChatID } }
func (t *Message) SetMigrateFromChatID(v int64) { t.MigrateFromChatID = &v }
func (t *Message) WithMigrateFromChatID(v int64) *Message { t.MigrateFromChatID = &v; return t }
func (t *Message) SetNilMigrateFromChatID() { t.MigrateFromChatID = nil }
func (t *Message) WithNilMigrateFromChatID() *Message { t.MigrateFromChatID = nil; return t }
func (t *Message) GetPinnedMessage() *Message { return t.PinnedMessage }
func (t *Message) SetPinnedMessage(v *Message) { t.PinnedMessage = v }
func (t *Message) WithPinnedMessage(v *Message) *Message { t.PinnedMessage = v; return t }

func (t *MessageEntity) GetType() MessageEntityType { return t.Type }
func (t *MessageEntity) SetType(v MessageEntityType) { t.Type = v }
func (t *MessageEntity) WithType(v MessageEntityType) *MessageEntity { t.Type = v; return t }
func (t *MessageEntity) GetOffset() int { return t.Offset }
func (t *MessageEntity) SetOffset(v int) { t.Offset = v }
func (t *MessageEntity) WithOffset(v int) *MessageEntity { t.Offset = v; return t }
func (t *MessageEntity) GetLength() int { return t.Length }
func (t *MessageEntity) SetLength(v int) { t.Length = v }
func (t *MessageEntity) WithLength(v int) *MessageEntity { t.Length = v; return t }
func (t *MessageEntity) GetURL() string { if t.URL == nil { return "" } else { return *t.URL } }
func (t *MessageEntity) SetURL(v string) { t.URL = &v }
func (t *MessageEntity) WithURL(v string) *MessageEntity { t.URL = &v; return t }
func (t *MessageEntity) SetNilURL() { t.URL = nil }
func (t *MessageEntity) WithNilURL() *MessageEntity { t.URL = nil; return t }
func (t *MessageEntity) SetFmtURL(f string, a ...interface{}) { t.SetURL(fmt.Sprintf(f, a...)) }
func (t *MessageEntity) WithFmtURL(f string, a ...interface{}) *MessageEntity { return t.WithURL(fmt.Sprintf(f, a...)) }
func (t *MessageEntity) GetUser() *User { return t.User }
func (t *MessageEntity) SetUser(v *User) { t.User = v }
func (t *MessageEntity) WithUser(v *User) *MessageEntity { t.User = v; return t }

func (t *PhotoSize) GetFileID() string { return t.FileID }
func (t *PhotoSize) SetFileID(v string) { t.FileID = v }
func (t *PhotoSize) WithFileID(v string) *PhotoSize { t.FileID = v; return t }
func (t *PhotoSize) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *PhotoSize) WithFmtFileID(f string, a ...interface{}) *PhotoSize { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *PhotoSize) GetWidth() int { return t.Width }
func (t *PhotoSize) SetWidth(v int) { t.Width = v }
func (t *PhotoSize) WithWidth(v int) *PhotoSize { t.Width = v; return t }
func (t *PhotoSize) GetHeight() int { return t.Height }
func (t *PhotoSize) SetHeight(v int) { t.Height = v }
func (t *PhotoSize) WithHeight(v int) *PhotoSize { t.Height = v; return t }
func (t *PhotoSize) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *PhotoSize) SetFileSize(v int) { t.FileSize = &v }
func (t *PhotoSize) WithFileSize(v int) *PhotoSize { t.FileSize = &v; return t }
func (t *PhotoSize) SetNilFileSize() { t.FileSize = nil }
func (t *PhotoSize) WithNilFileSize() *PhotoSize { t.FileSize = nil; return t }

func (t *Audio) GetFileID() string { return t.FileID }
func (t *Audio) SetFileID(v string) { t.FileID = v }
func (t *Audio) WithFileID(v string) *Audio { t.FileID = v; return t }
func (t *Audio) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *Audio) WithFmtFileID(f string, a ...interface{}) *Audio { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *Audio) GetDuration() int { return t.Duration }
func (t *Audio) SetDuration(v int) { t.Duration = v }
func (t *Audio) WithDuration(v int) *Audio { t.Duration = v; return t }
func (t *Audio) GetPerformer() string { if t.Performer == nil { return "" } else { return *t.Performer } }
func (t *Audio) SetPerformer(v string) { t.Performer = &v }
func (t *Audio) WithPerformer(v string) *Audio { t.Performer = &v; return t }
func (t *Audio) SetNilPerformer() { t.Performer = nil }
func (t *Audio) WithNilPerformer() *Audio { t.Performer = nil; return t }
func (t *Audio) SetFmtPerformer(f string, a ...interface{}) { t.SetPerformer(fmt.Sprintf(f, a...)) }
func (t *Audio) WithFmtPerformer(f string, a ...interface{}) *Audio { return t.WithPerformer(fmt.Sprintf(f, a...)) }
func (t *Audio) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *Audio) SetTitle(v string) { t.Title = &v }
func (t *Audio) WithTitle(v string) *Audio { t.Title = &v; return t }
func (t *Audio) SetNilTitle() { t.Title = nil }
func (t *Audio) WithNilTitle() *Audio { t.Title = nil; return t }
func (t *Audio) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *Audio) WithFmtTitle(f string, a ...interface{}) *Audio { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *Audio) GetMimeType() string { if t.MimeType == nil { return "" } else { return *t.MimeType } }
func (t *Audio) SetMimeType(v string) { t.MimeType = &v }
func (t *Audio) WithMimeType(v string) *Audio { t.MimeType = &v; return t }
func (t *Audio) SetNilMimeType() { t.MimeType = nil }
func (t *Audio) WithNilMimeType() *Audio { t.MimeType = nil; return t }
func (t *Audio) SetFmtMimeType(f string, a ...interface{}) { t.SetMimeType(fmt.Sprintf(f, a...)) }
func (t *Audio) WithFmtMimeType(f string, a ...interface{}) *Audio { return t.WithMimeType(fmt.Sprintf(f, a...)) }
func (t *Audio) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *Audio) SetFileSize(v int) { t.FileSize = &v }
func (t *Audio) WithFileSize(v int) *Audio { t.FileSize = &v; return t }
func (t *Audio) SetNilFileSize() { t.FileSize = nil }
func (t *Audio) WithNilFileSize() *Audio { t.FileSize = nil; return t }

func (t *Document) GetFileID() string { return t.FileID }
func (t *Document) SetFileID(v string) { t.FileID = v }
func (t *Document) WithFileID(v string) *Document { t.FileID = v; return t }
func (t *Document) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *Document) WithFmtFileID(f string, a ...interface{}) *Document { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *Document) GetThumb() *PhotoSize { return t.Thumb }
func (t *Document) SetThumb(v *PhotoSize) { t.Thumb = v }
func (t *Document) WithThumb(v *PhotoSize) *Document { t.Thumb = v; return t }
func (t *Document) GetFileName() string { if t.FileName == nil { return "" } else { return *t.FileName } }
func (t *Document) SetFileName(v string) { t.FileName = &v }
func (t *Document) WithFileName(v string) *Document { t.FileName = &v; return t }
func (t *Document) SetNilFileName() { t.FileName = nil }
func (t *Document) WithNilFileName() *Document { t.FileName = nil; return t }
func (t *Document) SetFmtFileName(f string, a ...interface{}) { t.SetFileName(fmt.Sprintf(f, a...)) }
func (t *Document) WithFmtFileName(f string, a ...interface{}) *Document { return t.WithFileName(fmt.Sprintf(f, a...)) }
func (t *Document) GetMIMEType() string { if t.MIMEType == nil { return "" } else { return *t.MIMEType } }
func (t *Document) SetMIMEType(v string) { t.MIMEType = &v }
func (t *Document) WithMIMEType(v string) *Document { t.MIMEType = &v; return t }
func (t *Document) SetNilMIMEType() { t.MIMEType = nil }
func (t *Document) WithNilMIMEType() *Document { t.MIMEType = nil; return t }
func (t *Document) SetFmtMIMEType(f string, a ...interface{}) { t.SetMIMEType(fmt.Sprintf(f, a...)) }
func (t *Document) WithFmtMIMEType(f string, a ...interface{}) *Document { return t.WithMIMEType(fmt.Sprintf(f, a...)) }
func (t *Document) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *Document) SetFileSize(v int) { t.FileSize = &v }
func (t *Document) WithFileSize(v int) *Document { t.FileSize = &v; return t }
func (t *Document) SetNilFileSize() { t.FileSize = nil }
func (t *Document) WithNilFileSize() *Document { t.FileSize = nil; return t }

func (t *Sticker) GetFileID() string { return t.FileID }
func (t *Sticker) SetFileID(v string) { t.FileID = v }
func (t *Sticker) WithFileID(v string) *Sticker { t.FileID = v; return t }
func (t *Sticker) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *Sticker) WithFmtFileID(f string, a ...interface{}) *Sticker { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *Sticker) GetWidth() int { return t.Width }
func (t *Sticker) SetWidth(v int) { t.Width = v }
func (t *Sticker) WithWidth(v int) *Sticker { t.Width = v; return t }
func (t *Sticker) GetHeight() int { return t.Height }
func (t *Sticker) SetHeight(v int) { t.Height = v }
func (t *Sticker) WithHeight(v int) *Sticker { t.Height = v; return t }
func (t *Sticker) GetThumb() *PhotoSize { return t.Thumb }
func (t *Sticker) SetThumb(v *PhotoSize) { t.Thumb = v }
func (t *Sticker) WithThumb(v *PhotoSize) *Sticker { t.Thumb = v; return t }
func (t *Sticker) GetEmoji() string { if t.Emoji == nil { return "" } else { return *t.Emoji } }
func (t *Sticker) SetEmoji(v string) { t.Emoji = &v }
func (t *Sticker) WithEmoji(v string) *Sticker { t.Emoji = &v; return t }
func (t *Sticker) SetNilEmoji() { t.Emoji = nil }
func (t *Sticker) WithNilEmoji() *Sticker { t.Emoji = nil; return t }
func (t *Sticker) SetFmtEmoji(f string, a ...interface{}) { t.SetEmoji(fmt.Sprintf(f, a...)) }
func (t *Sticker) WithFmtEmoji(f string, a ...interface{}) *Sticker { return t.WithEmoji(fmt.Sprintf(f, a...)) }
func (t *Sticker) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *Sticker) SetFileSize(v int) { t.FileSize = &v }
func (t *Sticker) WithFileSize(v int) *Sticker { t.FileSize = &v; return t }
func (t *Sticker) SetNilFileSize() { t.FileSize = nil }
func (t *Sticker) WithNilFileSize() *Sticker { t.FileSize = nil; return t }

func (t *Video) GetFileID() string { return t.FileID }
func (t *Video) SetFileID(v string) { t.FileID = v }
func (t *Video) WithFileID(v string) *Video { t.FileID = v; return t }
func (t *Video) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *Video) WithFmtFileID(f string, a ...interface{}) *Video { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *Video) GetWidth() int { return t.Width }
func (t *Video) SetWidth(v int) { t.Width = v }
func (t *Video) WithWidth(v int) *Video { t.Width = v; return t }
func (t *Video) GetHeight() int { return t.Height }
func (t *Video) SetHeight(v int) { t.Height = v }
func (t *Video) WithHeight(v int) *Video { t.Height = v; return t }
func (t *Video) GetDuration() int { return t.Duration }
func (t *Video) SetDuration(v int) { t.Duration = v }
func (t *Video) WithDuration(v int) *Video { t.Duration = v; return t }
func (t *Video) GetThumb() *PhotoSize { return t.Thumb }
func (t *Video) SetThumb(v *PhotoSize) { t.Thumb = v }
func (t *Video) WithThumb(v *PhotoSize) *Video { t.Thumb = v; return t }
func (t *Video) GetMimeType() string { if t.MimeType == nil { return "" } else { return *t.MimeType } }
func (t *Video) SetMimeType(v string) { t.MimeType = &v }
func (t *Video) WithMimeType(v string) *Video { t.MimeType = &v; return t }
func (t *Video) SetNilMimeType() { t.MimeType = nil }
func (t *Video) WithNilMimeType() *Video { t.MimeType = nil; return t }
func (t *Video) SetFmtMimeType(f string, a ...interface{}) { t.SetMimeType(fmt.Sprintf(f, a...)) }
func (t *Video) WithFmtMimeType(f string, a ...interface{}) *Video { return t.WithMimeType(fmt.Sprintf(f, a...)) }
func (t *Video) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *Video) SetFileSize(v int) { t.FileSize = &v }
func (t *Video) WithFileSize(v int) *Video { t.FileSize = &v; return t }
func (t *Video) SetNilFileSize() { t.FileSize = nil }
func (t *Video) WithNilFileSize() *Video { t.FileSize = nil; return t }

func (t *Voice) GetFileID() string { return t.FileID }
func (t *Voice) SetFileID(v string) { t.FileID = v }
func (t *Voice) WithFileID(v string) *Voice { t.FileID = v; return t }
func (t *Voice) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *Voice) WithFmtFileID(f string, a ...interface{}) *Voice { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *Voice) GetDuration() int { return t.Duration }
func (t *Voice) SetDuration(v int) { t.Duration = v }
func (t *Voice) WithDuration(v int) *Voice { t.Duration = v; return t }
func (t *Voice) GetMimeType() string { if t.MimeType == nil { return "" } else { return *t.MimeType } }
func (t *Voice) SetMimeType(v string) { t.MimeType = &v }
func (t *Voice) WithMimeType(v string) *Voice { t.MimeType = &v; return t }
func (t *Voice) SetNilMimeType() { t.MimeType = nil }
func (t *Voice) WithNilMimeType() *Voice { t.MimeType = nil; return t }
func (t *Voice) SetFmtMimeType(f string, a ...interface{}) { t.SetMimeType(fmt.Sprintf(f, a...)) }
func (t *Voice) WithFmtMimeType(f string, a ...interface{}) *Voice { return t.WithMimeType(fmt.Sprintf(f, a...)) }
func (t *Voice) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *Voice) SetFileSize(v int) { t.FileSize = &v }
func (t *Voice) WithFileSize(v int) *Voice { t.FileSize = &v; return t }
func (t *Voice) SetNilFileSize() { t.FileSize = nil }
func (t *Voice) WithNilFileSize() *Voice { t.FileSize = nil; return t }

func (t *Contact) GetPhoneNumber() string { return t.PhoneNumber }
func (t *Contact) SetPhoneNumber(v string) { t.PhoneNumber = v }
func (t *Contact) WithPhoneNumber(v string) *Contact { t.PhoneNumber = v; return t }
func (t *Contact) SetFmtPhoneNumber(f string, a ...interface{}) { t.SetPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *Contact) WithFmtPhoneNumber(f string, a ...interface{}) *Contact { return t.WithPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *Contact) GetFirstName() string { return t.FirstName }
func (t *Contact) SetFirstName(v string) { t.FirstName = v }
func (t *Contact) WithFirstName(v string) *Contact { t.FirstName = v; return t }
func (t *Contact) SetFmtFirstName(f string, a ...interface{}) { t.SetFirstName(fmt.Sprintf(f, a...)) }
func (t *Contact) WithFmtFirstName(f string, a ...interface{}) *Contact { return t.WithFirstName(fmt.Sprintf(f, a...)) }
func (t *Contact) GetLastName() string { if t.LastName == nil { return "" } else { return *t.LastName } }
func (t *Contact) SetLastName(v string) { t.LastName = &v }
func (t *Contact) WithLastName(v string) *Contact { t.LastName = &v; return t }
func (t *Contact) SetNilLastName() { t.LastName = nil }
func (t *Contact) WithNilLastName() *Contact { t.LastName = nil; return t }
func (t *Contact) SetFmtLastName(f string, a ...interface{}) { t.SetLastName(fmt.Sprintf(f, a...)) }
func (t *Contact) WithFmtLastName(f string, a ...interface{}) *Contact { return t.WithLastName(fmt.Sprintf(f, a...)) }
func (t *Contact) GetUserID() int { if t.UserID == nil { return 0 } else { return *t.UserID } }
func (t *Contact) SetUserID(v int) { t.UserID = &v }
func (t *Contact) WithUserID(v int) *Contact { t.UserID = &v; return t }
func (t *Contact) SetNilUserID() { t.UserID = nil }
func (t *Contact) WithNilUserID() *Contact { t.UserID = nil; return t }

func (t *Location) GetLongitude() float64 { return t.Longitude }
func (t *Location) SetLongitude(v float64) { t.Longitude = v }
func (t *Location) WithLongitude(v float64) *Location { t.Longitude = v; return t }
func (t *Location) GetLatitude() float64 { return t.Latitude }
func (t *Location) SetLatitude(v float64) { t.Latitude = v }
func (t *Location) WithLatitude(v float64) *Location { t.Latitude = v; return t }

func (t *Venue) GetLocation() *Location { return t.Location }
func (t *Venue) SetLocation(v *Location) { t.Location = v }
func (t *Venue) WithLocation(v *Location) *Venue { t.Location = v; return t }
func (t *Venue) GetTitle() string { return t.Title }
func (t *Venue) SetTitle(v string) { t.Title = v }
func (t *Venue) WithTitle(v string) *Venue { t.Title = v; return t }
func (t *Venue) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *Venue) WithFmtTitle(f string, a ...interface{}) *Venue { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *Venue) GetAddress() string { return t.Address }
func (t *Venue) SetAddress(v string) { t.Address = v }
func (t *Venue) WithAddress(v string) *Venue { t.Address = v; return t }
func (t *Venue) SetFmtAddress(f string, a ...interface{}) { t.SetAddress(fmt.Sprintf(f, a...)) }
func (t *Venue) WithFmtAddress(f string, a ...interface{}) *Venue { return t.WithAddress(fmt.Sprintf(f, a...)) }
func (t *Venue) GetFoursquareID() string { if t.FoursquareID == nil { return "" } else { return *t.FoursquareID } }
func (t *Venue) SetFoursquareID(v string) { t.FoursquareID = &v }
func (t *Venue) WithFoursquareID(v string) *Venue { t.FoursquareID = &v; return t }
func (t *Venue) SetNilFoursquareID() { t.FoursquareID = nil }
func (t *Venue) WithNilFoursquareID() *Venue { t.FoursquareID = nil; return t }
func (t *Venue) SetFmtFoursquareID(f string, a ...interface{}) { t.SetFoursquareID(fmt.Sprintf(f, a...)) }
func (t *Venue) WithFmtFoursquareID(f string, a ...interface{}) *Venue { return t.WithFoursquareID(fmt.Sprintf(f, a...)) }

func (t *UserProfilePhotos) GetTotalCount() int { return t.TotalCount }
func (t *UserProfilePhotos) SetTotalCount(v int) { t.TotalCount = v }
func (t *UserProfilePhotos) WithTotalCount(v int) *UserProfilePhotos { t.TotalCount = v; return t }
func (t *UserProfilePhotos) GetPhotos() [][]*PhotoSize { return t.Photos }
func (t *UserProfilePhotos) SetPhotos(v [][]*PhotoSize) { t.Photos = v }
func (t *UserProfilePhotos) WithPhotos(v [][]*PhotoSize) *UserProfilePhotos { t.Photos = v; return t }

func (t *File) GetFileID() string { return t.FileID }
func (t *File) SetFileID(v string) { t.FileID = v }
func (t *File) WithFileID(v string) *File { t.FileID = v; return t }
func (t *File) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *File) WithFmtFileID(f string, a ...interface{}) *File { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *File) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *File) SetFileSize(v int) { t.FileSize = &v }
func (t *File) WithFileSize(v int) *File { t.FileSize = &v; return t }
func (t *File) SetNilFileSize() { t.FileSize = nil }
func (t *File) WithNilFileSize() *File { t.FileSize = nil; return t }
func (t *File) GetFilePath() string { if t.FilePath == nil { return "" } else { return *t.FilePath } }
func (t *File) SetFilePath(v string) { t.FilePath = &v }
func (t *File) WithFilePath(v string) *File { t.FilePath = &v; return t }
func (t *File) SetNilFilePath() { t.FilePath = nil }
func (t *File) WithNilFilePath() *File { t.FilePath = nil; return t }
func (t *File) SetFmtFilePath(f string, a ...interface{}) { t.SetFilePath(fmt.Sprintf(f, a...)) }
func (t *File) WithFmtFilePath(f string, a ...interface{}) *File { return t.WithFilePath(fmt.Sprintf(f, a...)) }

func (t *ReplyKeyboardMarkup) GetKeyboard() [][]*KeyboardButton { return t.Keyboard }
func (t *ReplyKeyboardMarkup) SetKeyboard(v [][]*KeyboardButton) { t.Keyboard = v }
func (t *ReplyKeyboardMarkup) WithKeyboard(v [][]*KeyboardButton) *ReplyKeyboardMarkup { t.Keyboard = v; return t }
func (t *ReplyKeyboardMarkup) GetResizeKeyboard() bool { if t.ResizeKeyboard == nil { return false } else { return *t.ResizeKeyboard } }
func (t *ReplyKeyboardMarkup) SetResizeKeyboard(v bool) { t.ResizeKeyboard = &v }
func (t *ReplyKeyboardMarkup) WithResizeKeyboard(v bool) *ReplyKeyboardMarkup { t.ResizeKeyboard = &v; return t }
func (t *ReplyKeyboardMarkup) SetNilResizeKeyboard() { t.ResizeKeyboard = nil }
func (t *ReplyKeyboardMarkup) WithNilResizeKeyboard() *ReplyKeyboardMarkup { t.ResizeKeyboard = nil; return t }
func (t *ReplyKeyboardMarkup) GetOneTimeKeyboard() bool { if t.OneTimeKeyboard == nil { return false } else { return *t.OneTimeKeyboard } }
func (t *ReplyKeyboardMarkup) SetOneTimeKeyboard(v bool) { t.OneTimeKeyboard = &v }
func (t *ReplyKeyboardMarkup) WithOneTimeKeyboard(v bool) *ReplyKeyboardMarkup { t.OneTimeKeyboard = &v; return t }
func (t *ReplyKeyboardMarkup) SetNilOneTimeKeyboard() { t.OneTimeKeyboard = nil }
func (t *ReplyKeyboardMarkup) WithNilOneTimeKeyboard() *ReplyKeyboardMarkup { t.OneTimeKeyboard = nil; return t }
func (t *ReplyKeyboardMarkup) GetSelective() bool { if t.Selective == nil { return false } else { return *t.Selective } }
func (t *ReplyKeyboardMarkup) SetSelective(v bool) { t.Selective = &v }
func (t *ReplyKeyboardMarkup) WithSelective(v bool) *ReplyKeyboardMarkup { t.Selective = &v; return t }
func (t *ReplyKeyboardMarkup) SetNilSelective() { t.Selective = nil }
func (t *ReplyKeyboardMarkup) WithNilSelective() *ReplyKeyboardMarkup { t.Selective = nil; return t }

func (t *KeyboardButton) GetText() string { return t.Text }
func (t *KeyboardButton) SetText(v string) { t.Text = v }
func (t *KeyboardButton) WithText(v string) *KeyboardButton { t.Text = v; return t }
func (t *KeyboardButton) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *KeyboardButton) WithFmtText(f string, a ...interface{}) *KeyboardButton { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *KeyboardButton) GetRequestContact() bool { if t.RequestContact == nil { return false } else { return *t.RequestContact } }
func (t *KeyboardButton) SetRequestContact(v bool) { t.RequestContact = &v }
func (t *KeyboardButton) WithRequestContact(v bool) *KeyboardButton { t.RequestContact = &v; return t }
func (t *KeyboardButton) SetNilRequestContact() { t.RequestContact = nil }
func (t *KeyboardButton) WithNilRequestContact() *KeyboardButton { t.RequestContact = nil; return t }
func (t *KeyboardButton) GetRequestLocation() bool { if t.RequestLocation == nil { return false } else { return *t.RequestLocation } }
func (t *KeyboardButton) SetRequestLocation(v bool) { t.RequestLocation = &v }
func (t *KeyboardButton) WithRequestLocation(v bool) *KeyboardButton { t.RequestLocation = &v; return t }
func (t *KeyboardButton) SetNilRequestLocation() { t.RequestLocation = nil }
func (t *KeyboardButton) WithNilRequestLocation() *KeyboardButton { t.RequestLocation = nil; return t }

func (t *ReplyKeyboardRemove) GetRemoveKeyboard() bool { return t.RemoveKeyboard }
func (t *ReplyKeyboardRemove) SetRemoveKeyboard(v bool) { t.RemoveKeyboard = v }
func (t *ReplyKeyboardRemove) WithRemoveKeyboard(v bool) *ReplyKeyboardRemove { t.RemoveKeyboard = v; return t }
func (t *ReplyKeyboardRemove) GetSelective() bool { if t.Selective == nil { return false } else { return *t.Selective } }
func (t *ReplyKeyboardRemove) SetSelective(v bool) { t.Selective = &v }
func (t *ReplyKeyboardRemove) WithSelective(v bool) *ReplyKeyboardRemove { t.Selective = &v; return t }
func (t *ReplyKeyboardRemove) SetNilSelective() { t.Selective = nil }
func (t *ReplyKeyboardRemove) WithNilSelective() *ReplyKeyboardRemove { t.Selective = nil; return t }

func (t *InlineKeyboardMarkup) GetInlineKeyboard() [][]*InlineKeyboardButton { return t.InlineKeyboard }
func (t *InlineKeyboardMarkup) SetInlineKeyboard(v [][]*InlineKeyboardButton) { t.InlineKeyboard = v }
func (t *InlineKeyboardMarkup) WithInlineKeyboard(v [][]*InlineKeyboardButton) *InlineKeyboardMarkup { t.InlineKeyboard = v; return t }

func (t *InlineKeyboardButton) GetText() string { return t.Text }
func (t *InlineKeyboardButton) SetText(v string) { t.Text = v }
func (t *InlineKeyboardButton) WithText(v string) *InlineKeyboardButton { t.Text = v; return t }
func (t *InlineKeyboardButton) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) WithFmtText(f string, a ...interface{}) *InlineKeyboardButton { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) GetURL() string { if t.URL == nil { return "" } else { return *t.URL } }
func (t *InlineKeyboardButton) SetURL(v string) { t.URL = &v }
func (t *InlineKeyboardButton) WithURL(v string) *InlineKeyboardButton { t.URL = &v; return t }
func (t *InlineKeyboardButton) SetNilURL() { t.URL = nil }
func (t *InlineKeyboardButton) WithNilURL() *InlineKeyboardButton { t.URL = nil; return t }
func (t *InlineKeyboardButton) SetFmtURL(f string, a ...interface{}) { t.SetURL(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) WithFmtURL(f string, a ...interface{}) *InlineKeyboardButton { return t.WithURL(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) GetCallbackData() string { if t.CallbackData == nil { return "" } else { return *t.CallbackData } }
func (t *InlineKeyboardButton) SetCallbackData(v string) { t.CallbackData = &v }
func (t *InlineKeyboardButton) WithCallbackData(v string) *InlineKeyboardButton { t.CallbackData = &v; return t }
func (t *InlineKeyboardButton) SetNilCallbackData() { t.CallbackData = nil }
func (t *InlineKeyboardButton) WithNilCallbackData() *InlineKeyboardButton { t.CallbackData = nil; return t }
func (t *InlineKeyboardButton) SetFmtCallbackData(f string, a ...interface{}) { t.SetCallbackData(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) WithFmtCallbackData(f string, a ...interface{}) *InlineKeyboardButton { return t.WithCallbackData(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) GetSwitchInlineQuery() string { if t.SwitchInlineQuery == nil { return "" } else { return *t.SwitchInlineQuery } }
func (t *InlineKeyboardButton) SetSwitchInlineQuery(v string) { t.SwitchInlineQuery = &v }
func (t *InlineKeyboardButton) WithSwitchInlineQuery(v string) *InlineKeyboardButton { t.SwitchInlineQuery = &v; return t }
func (t *InlineKeyboardButton) SetNilSwitchInlineQuery() { t.SwitchInlineQuery = nil }
func (t *InlineKeyboardButton) WithNilSwitchInlineQuery() *InlineKeyboardButton { t.SwitchInlineQuery = nil; return t }
func (t *InlineKeyboardButton) SetFmtSwitchInlineQuery(f string, a ...interface{}) { t.SetSwitchInlineQuery(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) WithFmtSwitchInlineQuery(f string, a ...interface{}) *InlineKeyboardButton { return t.WithSwitchInlineQuery(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) GetSwitchInlineQueryCurrentChat() string { if t.SwitchInlineQueryCurrentChat == nil { return "" } else { return *t.SwitchInlineQueryCurrentChat } }
func (t *InlineKeyboardButton) SetSwitchInlineQueryCurrentChat(v string) { t.SwitchInlineQueryCurrentChat = &v }
func (t *InlineKeyboardButton) WithSwitchInlineQueryCurrentChat(v string) *InlineKeyboardButton { t.SwitchInlineQueryCurrentChat = &v; return t }
func (t *InlineKeyboardButton) SetNilSwitchInlineQueryCurrentChat() { t.SwitchInlineQueryCurrentChat = nil }
func (t *InlineKeyboardButton) WithNilSwitchInlineQueryCurrentChat() *InlineKeyboardButton { t.SwitchInlineQueryCurrentChat = nil; return t }
func (t *InlineKeyboardButton) SetFmtSwitchInlineQueryCurrentChat(f string, a ...interface{}) { t.SetSwitchInlineQueryCurrentChat(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) WithFmtSwitchInlineQueryCurrentChat(f string, a ...interface{}) *InlineKeyboardButton { return t.WithSwitchInlineQueryCurrentChat(fmt.Sprintf(f, a...)) }
func (t *InlineKeyboardButton) GetCallbackGame() *CallbackGame { return t.CallbackGame }
func (t *InlineKeyboardButton) SetCallbackGame(v *CallbackGame) { t.CallbackGame = v }
func (t *InlineKeyboardButton) WithCallbackGame(v *CallbackGame) *InlineKeyboardButton { t.CallbackGame = v; return t }

func (t *CallbackQuery) GetID() string { return t.ID }
func (t *CallbackQuery) SetID(v string) { t.ID = v }
func (t *CallbackQuery) WithID(v string) *CallbackQuery { t.ID = v; return t }
func (t *CallbackQuery) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) WithFmtID(f string, a ...interface{}) *CallbackQuery { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) GetFrom() *User { return t.From }
func (t *CallbackQuery) SetFrom(v *User) { t.From = v }
func (t *CallbackQuery) WithFrom(v *User) *CallbackQuery { t.From = v; return t }
func (t *CallbackQuery) GetMessage() *Message { return t.Message }
func (t *CallbackQuery) SetMessage(v *Message) { t.Message = v }
func (t *CallbackQuery) WithMessage(v *Message) *CallbackQuery { t.Message = v; return t }
func (t *CallbackQuery) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *CallbackQuery) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *CallbackQuery) WithInlineMessageID(v string) *CallbackQuery { t.InlineMessageID = &v; return t }
func (t *CallbackQuery) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *CallbackQuery) WithNilInlineMessageID() *CallbackQuery { t.InlineMessageID = nil; return t }
func (t *CallbackQuery) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) WithFmtInlineMessageID(f string, a ...interface{}) *CallbackQuery { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) GetChatInstance() string { return t.ChatInstance }
func (t *CallbackQuery) SetChatInstance(v string) { t.ChatInstance = v }
func (t *CallbackQuery) WithChatInstance(v string) *CallbackQuery { t.ChatInstance = v; return t }
func (t *CallbackQuery) SetFmtChatInstance(f string, a ...interface{}) { t.SetChatInstance(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) WithFmtChatInstance(f string, a ...interface{}) *CallbackQuery { return t.WithChatInstance(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) GetData() string { if t.Data == nil { return "" } else { return *t.Data } }
func (t *CallbackQuery) SetData(v string) { t.Data = &v }
func (t *CallbackQuery) WithData(v string) *CallbackQuery { t.Data = &v; return t }
func (t *CallbackQuery) SetNilData() { t.Data = nil }
func (t *CallbackQuery) WithNilData() *CallbackQuery { t.Data = nil; return t }
func (t *CallbackQuery) SetFmtData(f string, a ...interface{}) { t.SetData(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) WithFmtData(f string, a ...interface{}) *CallbackQuery { return t.WithData(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) GetGameShortName() string { if t.GameShortName == nil { return "" } else { return *t.GameShortName } }
func (t *CallbackQuery) SetGameShortName(v string) { t.GameShortName = &v }
func (t *CallbackQuery) WithGameShortName(v string) *CallbackQuery { t.GameShortName = &v; return t }
func (t *CallbackQuery) SetNilGameShortName() { t.GameShortName = nil }
func (t *CallbackQuery) WithNilGameShortName() *CallbackQuery { t.GameShortName = nil; return t }
func (t *CallbackQuery) SetFmtGameShortName(f string, a ...interface{}) { t.SetGameShortName(fmt.Sprintf(f, a...)) }
func (t *CallbackQuery) WithFmtGameShortName(f string, a ...interface{}) *CallbackQuery { return t.WithGameShortName(fmt.Sprintf(f, a...)) }

func (t *ForceReply) GetForceReply() bool { return t.ForceReply }
func (t *ForceReply) SetForceReply(v bool) { t.ForceReply = v }
func (t *ForceReply) WithForceReply(v bool) *ForceReply { t.ForceReply = v; return t }
func (t *ForceReply) GetSelective() bool { if t.Selective == nil { return false } else { return *t.Selective } }
func (t *ForceReply) SetSelective(v bool) { t.Selective = &v }
func (t *ForceReply) WithSelective(v bool) *ForceReply { t.Selective = &v; return t }
func (t *ForceReply) SetNilSelective() { t.Selective = nil }
func (t *ForceReply) WithNilSelective() *ForceReply { t.Selective = nil; return t }

func (t *ChatMember) GetUser() *User { return t.User }
func (t *ChatMember) SetUser(v *User) { t.User = v }
func (t *ChatMember) WithUser(v *User) *ChatMember { t.User = v; return t }
func (t *ChatMember) GetStatus() ChatMemberStatus { return t.Status }
func (t *ChatMember) SetStatus(v ChatMemberStatus) { t.Status = v }
func (t *ChatMember) WithStatus(v ChatMemberStatus) *ChatMember { t.Status = v; return t }

func (t *ResponseParameters) GetMigrateToChatID() int64 { if t.MigrateToChatID == nil { return 0 } else { return *t.MigrateToChatID } }
func (t *ResponseParameters) SetMigrateToChatID(v int64) { t.MigrateToChatID = &v }
func (t *ResponseParameters) WithMigrateToChatID(v int64) *ResponseParameters { t.MigrateToChatID = &v; return t }
func (t *ResponseParameters) SetNilMigrateToChatID() { t.MigrateToChatID = nil }
func (t *ResponseParameters) WithNilMigrateToChatID() *ResponseParameters { t.MigrateToChatID = nil; return t }
func (t *ResponseParameters) GetRetryAfter() int { if t.RetryAfter == nil { return 0 } else { return *t.RetryAfter } }
func (t *ResponseParameters) SetRetryAfter(v int) { t.RetryAfter = &v }
func (t *ResponseParameters) WithRetryAfter(v int) *ResponseParameters { t.RetryAfter = &v; return t }
func (t *ResponseParameters) SetNilRetryAfter() { t.RetryAfter = nil }
func (t *ResponseParameters) WithNilRetryAfter() *ResponseParameters { t.RetryAfter = nil; return t }

func (t *GetMe) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetMe) SetMethod(v MethodName) { t.Method = &v }
func (t *GetMe) WithMethod(v MethodName) *GetMe { t.Method = &v; return t }
func (t *GetMe) SetNilMethod() { t.Method = nil }
func (t *GetMe) WithNilMethod() *GetMe { t.Method = nil; return t }

func (t *SendMessage) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendMessage) SetMethod(v MethodName) { t.Method = &v }
func (t *SendMessage) WithMethod(v MethodName) *SendMessage { t.Method = &v; return t }
func (t *SendMessage) SetNilMethod() { t.Method = nil }
func (t *SendMessage) WithNilMethod() *SendMessage { t.Method = nil; return t }
func (t *SendMessage) GetChatID() ChatID { return t.ChatID }
func (t *SendMessage) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendMessage) WithChatID(v ChatID) *SendMessage { t.ChatID = v; return t }
func (t *SendMessage) GetText() string { return t.Text }
func (t *SendMessage) SetText(v string) { t.Text = v }
func (t *SendMessage) WithText(v string) *SendMessage { t.Text = v; return t }
func (t *SendMessage) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *SendMessage) WithFmtText(f string, a ...interface{}) *SendMessage { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *SendMessage) GetParseMode() ParseMode { if t.ParseMode == nil { return "" } else { return *t.ParseMode } }
func (t *SendMessage) SetParseMode(v ParseMode) { t.ParseMode = &v }
func (t *SendMessage) WithParseMode(v ParseMode) *SendMessage { t.ParseMode = &v; return t }
func (t *SendMessage) SetNilParseMode() { t.ParseMode = nil }
func (t *SendMessage) WithNilParseMode() *SendMessage { t.ParseMode = nil; return t }
func (t *SendMessage) GetDisableWebPagePreview() bool { if t.DisableWebPagePreview == nil { return false } else { return *t.DisableWebPagePreview } }
func (t *SendMessage) SetDisableWebPagePreview(v bool) { t.DisableWebPagePreview = &v }
func (t *SendMessage) WithDisableWebPagePreview(v bool) *SendMessage { t.DisableWebPagePreview = &v; return t }
func (t *SendMessage) SetNilDisableWebPagePreview() { t.DisableWebPagePreview = nil }
func (t *SendMessage) WithNilDisableWebPagePreview() *SendMessage { t.DisableWebPagePreview = nil; return t }
func (t *SendMessage) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendMessage) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendMessage) WithDisableNotification(v bool) *SendMessage { t.DisableNotification = &v; return t }
func (t *SendMessage) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendMessage) WithNilDisableNotification() *SendMessage { t.DisableNotification = nil; return t }
func (t *SendMessage) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendMessage) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendMessage) WithReplyToMessageID(v int) *SendMessage { t.ReplyToMessageID = &v; return t }
func (t *SendMessage) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendMessage) WithNilReplyToMessageID() *SendMessage { t.ReplyToMessageID = nil; return t }
func (t *SendMessage) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendMessage) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendMessage) WithReplyMarkup(v ReplyMarkup) *SendMessage { t.ReplyMarkup = v; return t }

func (t *ForwardMessage) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *ForwardMessage) SetMethod(v MethodName) { t.Method = &v }
func (t *ForwardMessage) WithMethod(v MethodName) *ForwardMessage { t.Method = &v; return t }
func (t *ForwardMessage) SetNilMethod() { t.Method = nil }
func (t *ForwardMessage) WithNilMethod() *ForwardMessage { t.Method = nil; return t }
func (t *ForwardMessage) GetChatID() ChatID { return t.ChatID }
func (t *ForwardMessage) SetChatID(v ChatID) { t.ChatID = v }
func (t *ForwardMessage) WithChatID(v ChatID) *ForwardMessage { t.ChatID = v; return t }
func (t *ForwardMessage) GetFromChatID() ChatID { return t.FromChatID }
func (t *ForwardMessage) SetFromChatID(v ChatID) { t.FromChatID = v }
func (t *ForwardMessage) WithFromChatID(v ChatID) *ForwardMessage { t.FromChatID = v; return t }
func (t *ForwardMessage) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *ForwardMessage) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *ForwardMessage) WithDisableNotification(v bool) *ForwardMessage { t.DisableNotification = &v; return t }
func (t *ForwardMessage) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *ForwardMessage) WithNilDisableNotification() *ForwardMessage { t.DisableNotification = nil; return t }
func (t *ForwardMessage) GetMessageID() int { return t.MessageID }
func (t *ForwardMessage) SetMessageID(v int) { t.MessageID = v }
func (t *ForwardMessage) WithMessageID(v int) *ForwardMessage { t.MessageID = v; return t }

func (t *SendPhoto) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendPhoto) SetMethod(v MethodName) { t.Method = &v }
func (t *SendPhoto) WithMethod(v MethodName) *SendPhoto { t.Method = &v; return t }
func (t *SendPhoto) SetNilMethod() { t.Method = nil }
func (t *SendPhoto) WithNilMethod() *SendPhoto { t.Method = nil; return t }
func (t *SendPhoto) GetChatID() ChatID { return t.ChatID }
func (t *SendPhoto) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendPhoto) WithChatID(v ChatID) *SendPhoto { t.ChatID = v; return t }
func (t *SendPhoto) GetPhoto() string { if t.Photo == nil { return "" } else { return *t.Photo } }
func (t *SendPhoto) SetPhoto(v string) { t.Photo = &v }
func (t *SendPhoto) WithPhoto(v string) *SendPhoto { t.Photo = &v; return t }
func (t *SendPhoto) SetNilPhoto() { t.Photo = nil }
func (t *SendPhoto) WithNilPhoto() *SendPhoto { t.Photo = nil; return t }
func (t *SendPhoto) SetFmtPhoto(f string, a ...interface{}) { t.SetPhoto(fmt.Sprintf(f, a...)) }
func (t *SendPhoto) WithFmtPhoto(f string, a ...interface{}) *SendPhoto { return t.WithPhoto(fmt.Sprintf(f, a...)) }
func (t *SendPhoto) GetPhotoFile() *InputFile { return t.PhotoFile }
func (t *SendPhoto) SetPhotoFile(v *InputFile) { t.PhotoFile = v }
func (t *SendPhoto) WithPhotoFile(v *InputFile) *SendPhoto { t.PhotoFile = v; return t }
func (t *SendPhoto) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *SendPhoto) SetCaption(v string) { t.Caption = &v }
func (t *SendPhoto) WithCaption(v string) *SendPhoto { t.Caption = &v; return t }
func (t *SendPhoto) SetNilCaption() { t.Caption = nil }
func (t *SendPhoto) WithNilCaption() *SendPhoto { t.Caption = nil; return t }
func (t *SendPhoto) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *SendPhoto) WithFmtCaption(f string, a ...interface{}) *SendPhoto { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *SendPhoto) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendPhoto) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendPhoto) WithDisableNotification(v bool) *SendPhoto { t.DisableNotification = &v; return t }
func (t *SendPhoto) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendPhoto) WithNilDisableNotification() *SendPhoto { t.DisableNotification = nil; return t }
func (t *SendPhoto) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendPhoto) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendPhoto) WithReplyToMessageID(v int) *SendPhoto { t.ReplyToMessageID = &v; return t }
func (t *SendPhoto) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendPhoto) WithNilReplyToMessageID() *SendPhoto { t.ReplyToMessageID = nil; return t }
func (t *SendPhoto) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendPhoto) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendPhoto) WithReplyMarkup(v ReplyMarkup) *SendPhoto { t.ReplyMarkup = v; return t }

func (t *SendAudio) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendAudio) SetMethod(v MethodName) { t.Method = &v }
func (t *SendAudio) WithMethod(v MethodName) *SendAudio { t.Method = &v; return t }
func (t *SendAudio) SetNilMethod() { t.Method = nil }
func (t *SendAudio) WithNilMethod() *SendAudio { t.Method = nil; return t }
func (t *SendAudio) GetChatID() ChatID { return t.ChatID }
func (t *SendAudio) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendAudio) WithChatID(v ChatID) *SendAudio { t.ChatID = v; return t }
func (t *SendAudio) GetAudio() string { if t.Audio == nil { return "" } else { return *t.Audio } }
func (t *SendAudio) SetAudio(v string) { t.Audio = &v }
func (t *SendAudio) WithAudio(v string) *SendAudio { t.Audio = &v; return t }
func (t *SendAudio) SetNilAudio() { t.Audio = nil }
func (t *SendAudio) WithNilAudio() *SendAudio { t.Audio = nil; return t }
func (t *SendAudio) SetFmtAudio(f string, a ...interface{}) { t.SetAudio(fmt.Sprintf(f, a...)) }
func (t *SendAudio) WithFmtAudio(f string, a ...interface{}) *SendAudio { return t.WithAudio(fmt.Sprintf(f, a...)) }
func (t *SendAudio) GetAudioFile() *InputFile { return t.AudioFile }
func (t *SendAudio) SetAudioFile(v *InputFile) { t.AudioFile = v }
func (t *SendAudio) WithAudioFile(v *InputFile) *SendAudio { t.AudioFile = v; return t }
func (t *SendAudio) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *SendAudio) SetCaption(v string) { t.Caption = &v }
func (t *SendAudio) WithCaption(v string) *SendAudio { t.Caption = &v; return t }
func (t *SendAudio) SetNilCaption() { t.Caption = nil }
func (t *SendAudio) WithNilCaption() *SendAudio { t.Caption = nil; return t }
func (t *SendAudio) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *SendAudio) WithFmtCaption(f string, a ...interface{}) *SendAudio { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *SendAudio) GetDuration() int { if t.Duration == nil { return 0 } else { return *t.Duration } }
func (t *SendAudio) SetDuration(v int) { t.Duration = &v }
func (t *SendAudio) WithDuration(v int) *SendAudio { t.Duration = &v; return t }
func (t *SendAudio) SetNilDuration() { t.Duration = nil }
func (t *SendAudio) WithNilDuration() *SendAudio { t.Duration = nil; return t }
func (t *SendAudio) GetPerformer() string { if t.Performer == nil { return "" } else { return *t.Performer } }
func (t *SendAudio) SetPerformer(v string) { t.Performer = &v }
func (t *SendAudio) WithPerformer(v string) *SendAudio { t.Performer = &v; return t }
func (t *SendAudio) SetNilPerformer() { t.Performer = nil }
func (t *SendAudio) WithNilPerformer() *SendAudio { t.Performer = nil; return t }
func (t *SendAudio) SetFmtPerformer(f string, a ...interface{}) { t.SetPerformer(fmt.Sprintf(f, a...)) }
func (t *SendAudio) WithFmtPerformer(f string, a ...interface{}) *SendAudio { return t.WithPerformer(fmt.Sprintf(f, a...)) }
func (t *SendAudio) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *SendAudio) SetTitle(v string) { t.Title = &v }
func (t *SendAudio) WithTitle(v string) *SendAudio { t.Title = &v; return t }
func (t *SendAudio) SetNilTitle() { t.Title = nil }
func (t *SendAudio) WithNilTitle() *SendAudio { t.Title = nil; return t }
func (t *SendAudio) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *SendAudio) WithFmtTitle(f string, a ...interface{}) *SendAudio { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *SendAudio) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendAudio) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendAudio) WithDisableNotification(v bool) *SendAudio { t.DisableNotification = &v; return t }
func (t *SendAudio) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendAudio) WithNilDisableNotification() *SendAudio { t.DisableNotification = nil; return t }
func (t *SendAudio) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendAudio) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendAudio) WithReplyToMessageID(v int) *SendAudio { t.ReplyToMessageID = &v; return t }
func (t *SendAudio) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendAudio) WithNilReplyToMessageID() *SendAudio { t.ReplyToMessageID = nil; return t }
func (t *SendAudio) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendAudio) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendAudio) WithReplyMarkup(v ReplyMarkup) *SendAudio { t.ReplyMarkup = v; return t }

func (t *SendDocument) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendDocument) SetMethod(v MethodName) { t.Method = &v }
func (t *SendDocument) WithMethod(v MethodName) *SendDocument { t.Method = &v; return t }
func (t *SendDocument) SetNilMethod() { t.Method = nil }
func (t *SendDocument) WithNilMethod() *SendDocument { t.Method = nil; return t }
func (t *SendDocument) GetChatID() ChatID { return t.ChatID }
func (t *SendDocument) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendDocument) WithChatID(v ChatID) *SendDocument { t.ChatID = v; return t }
func (t *SendDocument) GetDocument() string { if t.Document == nil { return "" } else { return *t.Document } }
func (t *SendDocument) SetDocument(v string) { t.Document = &v }
func (t *SendDocument) WithDocument(v string) *SendDocument { t.Document = &v; return t }
func (t *SendDocument) SetNilDocument() { t.Document = nil }
func (t *SendDocument) WithNilDocument() *SendDocument { t.Document = nil; return t }
func (t *SendDocument) SetFmtDocument(f string, a ...interface{}) { t.SetDocument(fmt.Sprintf(f, a...)) }
func (t *SendDocument) WithFmtDocument(f string, a ...interface{}) *SendDocument { return t.WithDocument(fmt.Sprintf(f, a...)) }
func (t *SendDocument) GetDocumentFile() *InputFile { return t.DocumentFile }
func (t *SendDocument) SetDocumentFile(v *InputFile) { t.DocumentFile = v }
func (t *SendDocument) WithDocumentFile(v *InputFile) *SendDocument { t.DocumentFile = v; return t }
func (t *SendDocument) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *SendDocument) SetCaption(v string) { t.Caption = &v }
func (t *SendDocument) WithCaption(v string) *SendDocument { t.Caption = &v; return t }
func (t *SendDocument) SetNilCaption() { t.Caption = nil }
func (t *SendDocument) WithNilCaption() *SendDocument { t.Caption = nil; return t }
func (t *SendDocument) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *SendDocument) WithFmtCaption(f string, a ...interface{}) *SendDocument { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *SendDocument) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendDocument) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendDocument) WithDisableNotification(v bool) *SendDocument { t.DisableNotification = &v; return t }
func (t *SendDocument) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendDocument) WithNilDisableNotification() *SendDocument { t.DisableNotification = nil; return t }
func (t *SendDocument) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendDocument) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendDocument) WithReplyToMessageID(v int) *SendDocument { t.ReplyToMessageID = &v; return t }
func (t *SendDocument) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendDocument) WithNilReplyToMessageID() *SendDocument { t.ReplyToMessageID = nil; return t }
func (t *SendDocument) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendDocument) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendDocument) WithReplyMarkup(v ReplyMarkup) *SendDocument { t.ReplyMarkup = v; return t }

func (t *SendSticker) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendSticker) SetMethod(v MethodName) { t.Method = &v }
func (t *SendSticker) WithMethod(v MethodName) *SendSticker { t.Method = &v; return t }
func (t *SendSticker) SetNilMethod() { t.Method = nil }
func (t *SendSticker) WithNilMethod() *SendSticker { t.Method = nil; return t }
func (t *SendSticker) GetChatID() ChatID { return t.ChatID }
func (t *SendSticker) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendSticker) WithChatID(v ChatID) *SendSticker { t.ChatID = v; return t }
func (t *SendSticker) GetSticker() string { if t.Sticker == nil { return "" } else { return *t.Sticker } }
func (t *SendSticker) SetSticker(v string) { t.Sticker = &v }
func (t *SendSticker) WithSticker(v string) *SendSticker { t.Sticker = &v; return t }
func (t *SendSticker) SetNilSticker() { t.Sticker = nil }
func (t *SendSticker) WithNilSticker() *SendSticker { t.Sticker = nil; return t }
func (t *SendSticker) SetFmtSticker(f string, a ...interface{}) { t.SetSticker(fmt.Sprintf(f, a...)) }
func (t *SendSticker) WithFmtSticker(f string, a ...interface{}) *SendSticker { return t.WithSticker(fmt.Sprintf(f, a...)) }
func (t *SendSticker) GetStickerFile() *InputFile { return t.StickerFile }
func (t *SendSticker) SetStickerFile(v *InputFile) { t.StickerFile = v }
func (t *SendSticker) WithStickerFile(v *InputFile) *SendSticker { t.StickerFile = v; return t }
func (t *SendSticker) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendSticker) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendSticker) WithDisableNotification(v bool) *SendSticker { t.DisableNotification = &v; return t }
func (t *SendSticker) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendSticker) WithNilDisableNotification() *SendSticker { t.DisableNotification = nil; return t }
func (t *SendSticker) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendSticker) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendSticker) WithReplyToMessageID(v int) *SendSticker { t.ReplyToMessageID = &v; return t }
func (t *SendSticker) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendSticker) WithNilReplyToMessageID() *SendSticker { t.ReplyToMessageID = nil; return t }
func (t *SendSticker) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendSticker) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendSticker) WithReplyMarkup(v ReplyMarkup) *SendSticker { t.ReplyMarkup = v; return t }

func (t *SendVideo) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendVideo) SetMethod(v MethodName) { t.Method = &v }
func (t *SendVideo) WithMethod(v MethodName) *SendVideo { t.Method = &v; return t }
func (t *SendVideo) SetNilMethod() { t.Method = nil }
func (t *SendVideo) WithNilMethod() *SendVideo { t.Method = nil; return t }
func (t *SendVideo) GetChatID() ChatID { return t.ChatID }
func (t *SendVideo) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendVideo) WithChatID(v ChatID) *SendVideo { t.ChatID = v; return t }
func (t *SendVideo) GetVideo() string { if t.Video == nil { return "" } else { return *t.Video } }
func (t *SendVideo) SetVideo(v string) { t.Video = &v }
func (t *SendVideo) WithVideo(v string) *SendVideo { t.Video = &v; return t }
func (t *SendVideo) SetNilVideo() { t.Video = nil }
func (t *SendVideo) WithNilVideo() *SendVideo { t.Video = nil; return t }
func (t *SendVideo) SetFmtVideo(f string, a ...interface{}) { t.SetVideo(fmt.Sprintf(f, a...)) }
func (t *SendVideo) WithFmtVideo(f string, a ...interface{}) *SendVideo { return t.WithVideo(fmt.Sprintf(f, a...)) }
func (t *SendVideo) GetVideoFile() *InputFile { return t.VideoFile }
func (t *SendVideo) SetVideoFile(v *InputFile) { t.VideoFile = v }
func (t *SendVideo) WithVideoFile(v *InputFile) *SendVideo { t.VideoFile = v; return t }
func (t *SendVideo) GetDuration() int { if t.Duration == nil { return 0 } else { return *t.Duration } }
func (t *SendVideo) SetDuration(v int) { t.Duration = &v }
func (t *SendVideo) WithDuration(v int) *SendVideo { t.Duration = &v; return t }
func (t *SendVideo) SetNilDuration() { t.Duration = nil }
func (t *SendVideo) WithNilDuration() *SendVideo { t.Duration = nil; return t }
func (t *SendVideo) GetWidth() int { if t.Width == nil { return 0 } else { return *t.Width } }
func (t *SendVideo) SetWidth(v int) { t.Width = &v }
func (t *SendVideo) WithWidth(v int) *SendVideo { t.Width = &v; return t }
func (t *SendVideo) SetNilWidth() { t.Width = nil }
func (t *SendVideo) WithNilWidth() *SendVideo { t.Width = nil; return t }
func (t *SendVideo) GetHeight() int { if t.Height == nil { return 0 } else { return *t.Height } }
func (t *SendVideo) SetHeight(v int) { t.Height = &v }
func (t *SendVideo) WithHeight(v int) *SendVideo { t.Height = &v; return t }
func (t *SendVideo) SetNilHeight() { t.Height = nil }
func (t *SendVideo) WithNilHeight() *SendVideo { t.Height = nil; return t }
func (t *SendVideo) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *SendVideo) SetCaption(v string) { t.Caption = &v }
func (t *SendVideo) WithCaption(v string) *SendVideo { t.Caption = &v; return t }
func (t *SendVideo) SetNilCaption() { t.Caption = nil }
func (t *SendVideo) WithNilCaption() *SendVideo { t.Caption = nil; return t }
func (t *SendVideo) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *SendVideo) WithFmtCaption(f string, a ...interface{}) *SendVideo { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *SendVideo) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendVideo) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendVideo) WithDisableNotification(v bool) *SendVideo { t.DisableNotification = &v; return t }
func (t *SendVideo) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendVideo) WithNilDisableNotification() *SendVideo { t.DisableNotification = nil; return t }
func (t *SendVideo) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendVideo) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendVideo) WithReplyToMessageID(v int) *SendVideo { t.ReplyToMessageID = &v; return t }
func (t *SendVideo) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendVideo) WithNilReplyToMessageID() *SendVideo { t.ReplyToMessageID = nil; return t }
func (t *SendVideo) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendVideo) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendVideo) WithReplyMarkup(v ReplyMarkup) *SendVideo { t.ReplyMarkup = v; return t }

func (t *SendVoice) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendVoice) SetMethod(v MethodName) { t.Method = &v }
func (t *SendVoice) WithMethod(v MethodName) *SendVoice { t.Method = &v; return t }
func (t *SendVoice) SetNilMethod() { t.Method = nil }
func (t *SendVoice) WithNilMethod() *SendVoice { t.Method = nil; return t }
func (t *SendVoice) GetChatID() ChatID { return t.ChatID }
func (t *SendVoice) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendVoice) WithChatID(v ChatID) *SendVoice { t.ChatID = v; return t }
func (t *SendVoice) GetVoice() string { if t.Voice == nil { return "" } else { return *t.Voice } }
func (t *SendVoice) SetVoice(v string) { t.Voice = &v }
func (t *SendVoice) WithVoice(v string) *SendVoice { t.Voice = &v; return t }
func (t *SendVoice) SetNilVoice() { t.Voice = nil }
func (t *SendVoice) WithNilVoice() *SendVoice { t.Voice = nil; return t }
func (t *SendVoice) SetFmtVoice(f string, a ...interface{}) { t.SetVoice(fmt.Sprintf(f, a...)) }
func (t *SendVoice) WithFmtVoice(f string, a ...interface{}) *SendVoice { return t.WithVoice(fmt.Sprintf(f, a...)) }
func (t *SendVoice) GetVoiceFile() *InputFile { return t.VoiceFile }
func (t *SendVoice) SetVoiceFile(v *InputFile) { t.VoiceFile = v }
func (t *SendVoice) WithVoiceFile(v *InputFile) *SendVoice { t.VoiceFile = v; return t }
func (t *SendVoice) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *SendVoice) SetCaption(v string) { t.Caption = &v }
func (t *SendVoice) WithCaption(v string) *SendVoice { t.Caption = &v; return t }
func (t *SendVoice) SetNilCaption() { t.Caption = nil }
func (t *SendVoice) WithNilCaption() *SendVoice { t.Caption = nil; return t }
func (t *SendVoice) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *SendVoice) WithFmtCaption(f string, a ...interface{}) *SendVoice { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *SendVoice) GetDuration() int { if t.Duration == nil { return 0 } else { return *t.Duration } }
func (t *SendVoice) SetDuration(v int) { t.Duration = &v }
func (t *SendVoice) WithDuration(v int) *SendVoice { t.Duration = &v; return t }
func (t *SendVoice) SetNilDuration() { t.Duration = nil }
func (t *SendVoice) WithNilDuration() *SendVoice { t.Duration = nil; return t }
func (t *SendVoice) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendVoice) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendVoice) WithDisableNotification(v bool) *SendVoice { t.DisableNotification = &v; return t }
func (t *SendVoice) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendVoice) WithNilDisableNotification() *SendVoice { t.DisableNotification = nil; return t }
func (t *SendVoice) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendVoice) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendVoice) WithReplyToMessageID(v int) *SendVoice { t.ReplyToMessageID = &v; return t }
func (t *SendVoice) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendVoice) WithNilReplyToMessageID() *SendVoice { t.ReplyToMessageID = nil; return t }
func (t *SendVoice) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendVoice) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendVoice) WithReplyMarkup(v ReplyMarkup) *SendVoice { t.ReplyMarkup = v; return t }

func (t *SendLocation) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendLocation) SetMethod(v MethodName) { t.Method = &v }
func (t *SendLocation) WithMethod(v MethodName) *SendLocation { t.Method = &v; return t }
func (t *SendLocation) SetNilMethod() { t.Method = nil }
func (t *SendLocation) WithNilMethod() *SendLocation { t.Method = nil; return t }
func (t *SendLocation) GetChatID() ChatID { return t.ChatID }
func (t *SendLocation) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendLocation) WithChatID(v ChatID) *SendLocation { t.ChatID = v; return t }
func (t *SendLocation) GetLatitude() float64 { return t.Latitude }
func (t *SendLocation) SetLatitude(v float64) { t.Latitude = v }
func (t *SendLocation) WithLatitude(v float64) *SendLocation { t.Latitude = v; return t }
func (t *SendLocation) GetLongitude() float64 { return t.Longitude }
func (t *SendLocation) SetLongitude(v float64) { t.Longitude = v }
func (t *SendLocation) WithLongitude(v float64) *SendLocation { t.Longitude = v; return t }
func (t *SendLocation) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendLocation) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendLocation) WithDisableNotification(v bool) *SendLocation { t.DisableNotification = &v; return t }
func (t *SendLocation) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendLocation) WithNilDisableNotification() *SendLocation { t.DisableNotification = nil; return t }
func (t *SendLocation) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendLocation) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendLocation) WithReplyToMessageID(v int) *SendLocation { t.ReplyToMessageID = &v; return t }
func (t *SendLocation) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendLocation) WithNilReplyToMessageID() *SendLocation { t.ReplyToMessageID = nil; return t }
func (t *SendLocation) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendLocation) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendLocation) WithReplyMarkup(v ReplyMarkup) *SendLocation { t.ReplyMarkup = v; return t }

func (t *SendVenue) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendVenue) SetMethod(v MethodName) { t.Method = &v }
func (t *SendVenue) WithMethod(v MethodName) *SendVenue { t.Method = &v; return t }
func (t *SendVenue) SetNilMethod() { t.Method = nil }
func (t *SendVenue) WithNilMethod() *SendVenue { t.Method = nil; return t }
func (t *SendVenue) GetChatID() ChatID { return t.ChatID }
func (t *SendVenue) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendVenue) WithChatID(v ChatID) *SendVenue { t.ChatID = v; return t }
func (t *SendVenue) GetLatitude() float64 { return t.Latitude }
func (t *SendVenue) SetLatitude(v float64) { t.Latitude = v }
func (t *SendVenue) WithLatitude(v float64) *SendVenue { t.Latitude = v; return t }
func (t *SendVenue) GetLongitude() float64 { return t.Longitude }
func (t *SendVenue) SetLongitude(v float64) { t.Longitude = v }
func (t *SendVenue) WithLongitude(v float64) *SendVenue { t.Longitude = v; return t }
func (t *SendVenue) GetTitle() string { return t.Title }
func (t *SendVenue) SetTitle(v string) { t.Title = v }
func (t *SendVenue) WithTitle(v string) *SendVenue { t.Title = v; return t }
func (t *SendVenue) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *SendVenue) WithFmtTitle(f string, a ...interface{}) *SendVenue { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *SendVenue) GetAddress() string { return t.Address }
func (t *SendVenue) SetAddress(v string) { t.Address = v }
func (t *SendVenue) WithAddress(v string) *SendVenue { t.Address = v; return t }
func (t *SendVenue) SetFmtAddress(f string, a ...interface{}) { t.SetAddress(fmt.Sprintf(f, a...)) }
func (t *SendVenue) WithFmtAddress(f string, a ...interface{}) *SendVenue { return t.WithAddress(fmt.Sprintf(f, a...)) }
func (t *SendVenue) GetFoursquareID() string { if t.FoursquareID == nil { return "" } else { return *t.FoursquareID } }
func (t *SendVenue) SetFoursquareID(v string) { t.FoursquareID = &v }
func (t *SendVenue) WithFoursquareID(v string) *SendVenue { t.FoursquareID = &v; return t }
func (t *SendVenue) SetNilFoursquareID() { t.FoursquareID = nil }
func (t *SendVenue) WithNilFoursquareID() *SendVenue { t.FoursquareID = nil; return t }
func (t *SendVenue) SetFmtFoursquareID(f string, a ...interface{}) { t.SetFoursquareID(fmt.Sprintf(f, a...)) }
func (t *SendVenue) WithFmtFoursquareID(f string, a ...interface{}) *SendVenue { return t.WithFoursquareID(fmt.Sprintf(f, a...)) }
func (t *SendVenue) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendVenue) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendVenue) WithDisableNotification(v bool) *SendVenue { t.DisableNotification = &v; return t }
func (t *SendVenue) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendVenue) WithNilDisableNotification() *SendVenue { t.DisableNotification = nil; return t }
func (t *SendVenue) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendVenue) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendVenue) WithReplyToMessageID(v int) *SendVenue { t.ReplyToMessageID = &v; return t }
func (t *SendVenue) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendVenue) WithNilReplyToMessageID() *SendVenue { t.ReplyToMessageID = nil; return t }
func (t *SendVenue) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendVenue) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendVenue) WithReplyMarkup(v ReplyMarkup) *SendVenue { t.ReplyMarkup = v; return t }

func (t *SendContact) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendContact) SetMethod(v MethodName) { t.Method = &v }
func (t *SendContact) WithMethod(v MethodName) *SendContact { t.Method = &v; return t }
func (t *SendContact) SetNilMethod() { t.Method = nil }
func (t *SendContact) WithNilMethod() *SendContact { t.Method = nil; return t }
func (t *SendContact) GetChatID() ChatID { return t.ChatID }
func (t *SendContact) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendContact) WithChatID(v ChatID) *SendContact { t.ChatID = v; return t }
func (t *SendContact) GetPhoneNumber() string { return t.PhoneNumber }
func (t *SendContact) SetPhoneNumber(v string) { t.PhoneNumber = v }
func (t *SendContact) WithPhoneNumber(v string) *SendContact { t.PhoneNumber = v; return t }
func (t *SendContact) SetFmtPhoneNumber(f string, a ...interface{}) { t.SetPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *SendContact) WithFmtPhoneNumber(f string, a ...interface{}) *SendContact { return t.WithPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *SendContact) GetFirstName() string { return t.FirstName }
func (t *SendContact) SetFirstName(v string) { t.FirstName = v }
func (t *SendContact) WithFirstName(v string) *SendContact { t.FirstName = v; return t }
func (t *SendContact) SetFmtFirstName(f string, a ...interface{}) { t.SetFirstName(fmt.Sprintf(f, a...)) }
func (t *SendContact) WithFmtFirstName(f string, a ...interface{}) *SendContact { return t.WithFirstName(fmt.Sprintf(f, a...)) }
func (t *SendContact) GetLastName() string { if t.LastName == nil { return "" } else { return *t.LastName } }
func (t *SendContact) SetLastName(v string) { t.LastName = &v }
func (t *SendContact) WithLastName(v string) *SendContact { t.LastName = &v; return t }
func (t *SendContact) SetNilLastName() { t.LastName = nil }
func (t *SendContact) WithNilLastName() *SendContact { t.LastName = nil; return t }
func (t *SendContact) SetFmtLastName(f string, a ...interface{}) { t.SetLastName(fmt.Sprintf(f, a...)) }
func (t *SendContact) WithFmtLastName(f string, a ...interface{}) *SendContact { return t.WithLastName(fmt.Sprintf(f, a...)) }
func (t *SendContact) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendContact) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendContact) WithDisableNotification(v bool) *SendContact { t.DisableNotification = &v; return t }
func (t *SendContact) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendContact) WithNilDisableNotification() *SendContact { t.DisableNotification = nil; return t }
func (t *SendContact) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendContact) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendContact) WithReplyToMessageID(v int) *SendContact { t.ReplyToMessageID = &v; return t }
func (t *SendContact) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendContact) WithNilReplyToMessageID() *SendContact { t.ReplyToMessageID = nil; return t }
func (t *SendContact) GetReplyMarkup() ReplyMarkup { return t.ReplyMarkup }
func (t *SendContact) SetReplyMarkup(v ReplyMarkup) { t.ReplyMarkup = v }
func (t *SendContact) WithReplyMarkup(v ReplyMarkup) *SendContact { t.ReplyMarkup = v; return t }

func (t *SendChatAction) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendChatAction) SetMethod(v MethodName) { t.Method = &v }
func (t *SendChatAction) WithMethod(v MethodName) *SendChatAction { t.Method = &v; return t }
func (t *SendChatAction) SetNilMethod() { t.Method = nil }
func (t *SendChatAction) WithNilMethod() *SendChatAction { t.Method = nil; return t }
func (t *SendChatAction) GetChatID() ChatID { return t.ChatID }
func (t *SendChatAction) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendChatAction) WithChatID(v ChatID) *SendChatAction { t.ChatID = v; return t }
func (t *SendChatAction) GetAction() ChatAction { return t.Action }
func (t *SendChatAction) SetAction(v ChatAction) { t.Action = v }
func (t *SendChatAction) WithAction(v ChatAction) *SendChatAction { t.Action = v; return t }

func (t *GetUserProfilePhotos) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetUserProfilePhotos) SetMethod(v MethodName) { t.Method = &v }
func (t *GetUserProfilePhotos) WithMethod(v MethodName) *GetUserProfilePhotos { t.Method = &v; return t }
func (t *GetUserProfilePhotos) SetNilMethod() { t.Method = nil }
func (t *GetUserProfilePhotos) WithNilMethod() *GetUserProfilePhotos { t.Method = nil; return t }
func (t *GetUserProfilePhotos) GetUserID() int { return t.UserID }
func (t *GetUserProfilePhotos) SetUserID(v int) { t.UserID = v }
func (t *GetUserProfilePhotos) WithUserID(v int) *GetUserProfilePhotos { t.UserID = v; return t }
func (t *GetUserProfilePhotos) GetOffset() int { if t.Offset == nil { return 0 } else { return *t.Offset } }
func (t *GetUserProfilePhotos) SetOffset(v int) { t.Offset = &v }
func (t *GetUserProfilePhotos) WithOffset(v int) *GetUserProfilePhotos { t.Offset = &v; return t }
func (t *GetUserProfilePhotos) SetNilOffset() { t.Offset = nil }
func (t *GetUserProfilePhotos) WithNilOffset() *GetUserProfilePhotos { t.Offset = nil; return t }
func (t *GetUserProfilePhotos) GetLimit() int { if t.Limit == nil { return 0 } else { return *t.Limit } }
func (t *GetUserProfilePhotos) SetLimit(v int) { t.Limit = &v }
func (t *GetUserProfilePhotos) WithLimit(v int) *GetUserProfilePhotos { t.Limit = &v; return t }
func (t *GetUserProfilePhotos) SetNilLimit() { t.Limit = nil }
func (t *GetUserProfilePhotos) WithNilLimit() *GetUserProfilePhotos { t.Limit = nil; return t }

func (t *GetFile) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetFile) SetMethod(v MethodName) { t.Method = &v }
func (t *GetFile) WithMethod(v MethodName) *GetFile { t.Method = &v; return t }
func (t *GetFile) SetNilMethod() { t.Method = nil }
func (t *GetFile) WithNilMethod() *GetFile { t.Method = nil; return t }
func (t *GetFile) GetFileID() string { return t.FileID }
func (t *GetFile) SetFileID(v string) { t.FileID = v }
func (t *GetFile) WithFileID(v string) *GetFile { t.FileID = v; return t }
func (t *GetFile) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *GetFile) WithFmtFileID(f string, a ...interface{}) *GetFile { return t.WithFileID(fmt.Sprintf(f, a...)) }

func (t *KickChatMember) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *KickChatMember) SetMethod(v MethodName) { t.Method = &v }
func (t *KickChatMember) WithMethod(v MethodName) *KickChatMember { t.Method = &v; return t }
func (t *KickChatMember) SetNilMethod() { t.Method = nil }
func (t *KickChatMember) WithNilMethod() *KickChatMember { t.Method = nil; return t }
func (t *KickChatMember) GetChatID() ChatID { return t.ChatID }
func (t *KickChatMember) SetChatID(v ChatID) { t.ChatID = v }
func (t *KickChatMember) WithChatID(v ChatID) *KickChatMember { t.ChatID = v; return t }
func (t *KickChatMember) GetUserID() int { return t.UserID }
func (t *KickChatMember) SetUserID(v int) { t.UserID = v }
func (t *KickChatMember) WithUserID(v int) *KickChatMember { t.UserID = v; return t }

func (t *LeaveChat) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *LeaveChat) SetMethod(v MethodName) { t.Method = &v }
func (t *LeaveChat) WithMethod(v MethodName) *LeaveChat { t.Method = &v; return t }
func (t *LeaveChat) SetNilMethod() { t.Method = nil }
func (t *LeaveChat) WithNilMethod() *LeaveChat { t.Method = nil; return t }
func (t *LeaveChat) GetChatID() ChatID { return t.ChatID }
func (t *LeaveChat) SetChatID(v ChatID) { t.ChatID = v }
func (t *LeaveChat) WithChatID(v ChatID) *LeaveChat { t.ChatID = v; return t }

func (t *UnbanChatMember) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *UnbanChatMember) SetMethod(v MethodName) { t.Method = &v }
func (t *UnbanChatMember) WithMethod(v MethodName) *UnbanChatMember { t.Method = &v; return t }
func (t *UnbanChatMember) SetNilMethod() { t.Method = nil }
func (t *UnbanChatMember) WithNilMethod() *UnbanChatMember { t.Method = nil; return t }
func (t *UnbanChatMember) GetChatID() ChatID { return t.ChatID }
func (t *UnbanChatMember) SetChatID(v ChatID) { t.ChatID = v }
func (t *UnbanChatMember) WithChatID(v ChatID) *UnbanChatMember { t.ChatID = v; return t }
func (t *UnbanChatMember) GetUserID() int { return t.UserID }
func (t *UnbanChatMember) SetUserID(v int) { t.UserID = v }
func (t *UnbanChatMember) WithUserID(v int) *UnbanChatMember { t.UserID = v; return t }

func (t *GetChat) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetChat) SetMethod(v MethodName) { t.Method = &v }
func (t *GetChat) WithMethod(v MethodName) *GetChat { t.Method = &v; return t }
func (t *GetChat) SetNilMethod() { t.Method = nil }
func (t *GetChat) WithNilMethod() *GetChat { t.Method = nil; return t }
func (t *GetChat) GetChatID() ChatID { return t.ChatID }
func (t *GetChat) SetChatID(v ChatID) { t.ChatID = v }
func (t *GetChat) WithChatID(v ChatID) *GetChat { t.ChatID = v; return t }

func (t *GetChatAdministrators) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetChatAdministrators) SetMethod(v MethodName) { t.Method = &v }
func (t *GetChatAdministrators) WithMethod(v MethodName) *GetChatAdministrators { t.Method = &v; return t }
func (t *GetChatAdministrators) SetNilMethod() { t.Method = nil }
func (t *GetChatAdministrators) WithNilMethod() *GetChatAdministrators { t.Method = nil; return t }
func (t *GetChatAdministrators) GetChatID() ChatID { return t.ChatID }
func (t *GetChatAdministrators) SetChatID(v ChatID) { t.ChatID = v }
func (t *GetChatAdministrators) WithChatID(v ChatID) *GetChatAdministrators { t.ChatID = v; return t }

func (t *GetChatMembersCount) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetChatMembersCount) SetMethod(v MethodName) { t.Method = &v }
func (t *GetChatMembersCount) WithMethod(v MethodName) *GetChatMembersCount { t.Method = &v; return t }
func (t *GetChatMembersCount) SetNilMethod() { t.Method = nil }
func (t *GetChatMembersCount) WithNilMethod() *GetChatMembersCount { t.Method = nil; return t }
func (t *GetChatMembersCount) GetChatID() ChatID { return t.ChatID }
func (t *GetChatMembersCount) SetChatID(v ChatID) { t.ChatID = v }
func (t *GetChatMembersCount) WithChatID(v ChatID) *GetChatMembersCount { t.ChatID = v; return t }

func (t *GetChatMember) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetChatMember) SetMethod(v MethodName) { t.Method = &v }
func (t *GetChatMember) WithMethod(v MethodName) *GetChatMember { t.Method = &v; return t }
func (t *GetChatMember) SetNilMethod() { t.Method = nil }
func (t *GetChatMember) WithNilMethod() *GetChatMember { t.Method = nil; return t }
func (t *GetChatMember) GetChatID() ChatID { return t.ChatID }
func (t *GetChatMember) SetChatID(v ChatID) { t.ChatID = v }
func (t *GetChatMember) WithChatID(v ChatID) *GetChatMember { t.ChatID = v; return t }
func (t *GetChatMember) GetUserID() int { return t.UserID }
func (t *GetChatMember) SetUserID(v int) { t.UserID = v }
func (t *GetChatMember) WithUserID(v int) *GetChatMember { t.UserID = v; return t }

func (t *AnswerCallbackQuery) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *AnswerCallbackQuery) SetMethod(v MethodName) { t.Method = &v }
func (t *AnswerCallbackQuery) WithMethod(v MethodName) *AnswerCallbackQuery { t.Method = &v; return t }
func (t *AnswerCallbackQuery) SetNilMethod() { t.Method = nil }
func (t *AnswerCallbackQuery) WithNilMethod() *AnswerCallbackQuery { t.Method = nil; return t }
func (t *AnswerCallbackQuery) GetCallbackQueryID() string { return t.CallbackQueryID }
func (t *AnswerCallbackQuery) SetCallbackQueryID(v string) { t.CallbackQueryID = v }
func (t *AnswerCallbackQuery) WithCallbackQueryID(v string) *AnswerCallbackQuery { t.CallbackQueryID = v; return t }
func (t *AnswerCallbackQuery) SetFmtCallbackQueryID(f string, a ...interface{}) { t.SetCallbackQueryID(fmt.Sprintf(f, a...)) }
func (t *AnswerCallbackQuery) WithFmtCallbackQueryID(f string, a ...interface{}) *AnswerCallbackQuery { return t.WithCallbackQueryID(fmt.Sprintf(f, a...)) }
func (t *AnswerCallbackQuery) GetText() string { if t.Text == nil { return "" } else { return *t.Text } }
func (t *AnswerCallbackQuery) SetText(v string) { t.Text = &v }
func (t *AnswerCallbackQuery) WithText(v string) *AnswerCallbackQuery { t.Text = &v; return t }
func (t *AnswerCallbackQuery) SetNilText() { t.Text = nil }
func (t *AnswerCallbackQuery) WithNilText() *AnswerCallbackQuery { t.Text = nil; return t }
func (t *AnswerCallbackQuery) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *AnswerCallbackQuery) WithFmtText(f string, a ...interface{}) *AnswerCallbackQuery { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *AnswerCallbackQuery) GetShowAlert() bool { if t.ShowAlert == nil { return false } else { return *t.ShowAlert } }
func (t *AnswerCallbackQuery) SetShowAlert(v bool) { t.ShowAlert = &v }
func (t *AnswerCallbackQuery) WithShowAlert(v bool) *AnswerCallbackQuery { t.ShowAlert = &v; return t }
func (t *AnswerCallbackQuery) SetNilShowAlert() { t.ShowAlert = nil }
func (t *AnswerCallbackQuery) WithNilShowAlert() *AnswerCallbackQuery { t.ShowAlert = nil; return t }
func (t *AnswerCallbackQuery) GetURL() string { if t.URL == nil { return "" } else { return *t.URL } }
func (t *AnswerCallbackQuery) SetURL(v string) { t.URL = &v }
func (t *AnswerCallbackQuery) WithURL(v string) *AnswerCallbackQuery { t.URL = &v; return t }
func (t *AnswerCallbackQuery) SetNilURL() { t.URL = nil }
func (t *AnswerCallbackQuery) WithNilURL() *AnswerCallbackQuery { t.URL = nil; return t }
func (t *AnswerCallbackQuery) SetFmtURL(f string, a ...interface{}) { t.SetURL(fmt.Sprintf(f, a...)) }
func (t *AnswerCallbackQuery) WithFmtURL(f string, a ...interface{}) *AnswerCallbackQuery { return t.WithURL(fmt.Sprintf(f, a...)) }
func (t *AnswerCallbackQuery) GetCacheTime() int { if t.CacheTime == nil { return 0 } else { return *t.CacheTime } }
func (t *AnswerCallbackQuery) SetCacheTime(v int) { t.CacheTime = &v }
func (t *AnswerCallbackQuery) WithCacheTime(v int) *AnswerCallbackQuery { t.CacheTime = &v; return t }
func (t *AnswerCallbackQuery) SetNilCacheTime() { t.CacheTime = nil }
func (t *AnswerCallbackQuery) WithNilCacheTime() *AnswerCallbackQuery { t.CacheTime = nil; return t }

func (t *EditMessageText) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *EditMessageText) SetMethod(v MethodName) { t.Method = &v }
func (t *EditMessageText) WithMethod(v MethodName) *EditMessageText { t.Method = &v; return t }
func (t *EditMessageText) SetNilMethod() { t.Method = nil }
func (t *EditMessageText) WithNilMethod() *EditMessageText { t.Method = nil; return t }
func (t *EditMessageText) GetChatID() ChatID { return t.ChatID }
func (t *EditMessageText) SetChatID(v ChatID) { t.ChatID = v }
func (t *EditMessageText) WithChatID(v ChatID) *EditMessageText { t.ChatID = v; return t }
func (t *EditMessageText) GetMessageID() int { if t.MessageID == nil { return 0 } else { return *t.MessageID } }
func (t *EditMessageText) SetMessageID(v int) { t.MessageID = &v }
func (t *EditMessageText) WithMessageID(v int) *EditMessageText { t.MessageID = &v; return t }
func (t *EditMessageText) SetNilMessageID() { t.MessageID = nil }
func (t *EditMessageText) WithNilMessageID() *EditMessageText { t.MessageID = nil; return t }
func (t *EditMessageText) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *EditMessageText) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *EditMessageText) WithInlineMessageID(v string) *EditMessageText { t.InlineMessageID = &v; return t }
func (t *EditMessageText) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *EditMessageText) WithNilInlineMessageID() *EditMessageText { t.InlineMessageID = nil; return t }
func (t *EditMessageText) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *EditMessageText) WithFmtInlineMessageID(f string, a ...interface{}) *EditMessageText { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *EditMessageText) GetText() string { return t.Text }
func (t *EditMessageText) SetText(v string) { t.Text = v }
func (t *EditMessageText) WithText(v string) *EditMessageText { t.Text = v; return t }
func (t *EditMessageText) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *EditMessageText) WithFmtText(f string, a ...interface{}) *EditMessageText { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *EditMessageText) GetParseMode() ParseMode { if t.ParseMode == nil { return "" } else { return *t.ParseMode } }
func (t *EditMessageText) SetParseMode(v ParseMode) { t.ParseMode = &v }
func (t *EditMessageText) WithParseMode(v ParseMode) *EditMessageText { t.ParseMode = &v; return t }
func (t *EditMessageText) SetNilParseMode() { t.ParseMode = nil }
func (t *EditMessageText) WithNilParseMode() *EditMessageText { t.ParseMode = nil; return t }
func (t *EditMessageText) GetDisableWebPagePreview() bool { if t.DisableWebPagePreview == nil { return false } else { return *t.DisableWebPagePreview } }
func (t *EditMessageText) SetDisableWebPagePreview(v bool) { t.DisableWebPagePreview = &v }
func (t *EditMessageText) WithDisableWebPagePreview(v bool) *EditMessageText { t.DisableWebPagePreview = &v; return t }
func (t *EditMessageText) SetNilDisableWebPagePreview() { t.DisableWebPagePreview = nil }
func (t *EditMessageText) WithNilDisableWebPagePreview() *EditMessageText { t.DisableWebPagePreview = nil; return t }
func (t *EditMessageText) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *EditMessageText) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *EditMessageText) WithReplyMarkup(v *InlineKeyboardMarkup) *EditMessageText { t.ReplyMarkup = v; return t }

func (t *EditMessageCaption) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *EditMessageCaption) SetMethod(v MethodName) { t.Method = &v }
func (t *EditMessageCaption) WithMethod(v MethodName) *EditMessageCaption { t.Method = &v; return t }
func (t *EditMessageCaption) SetNilMethod() { t.Method = nil }
func (t *EditMessageCaption) WithNilMethod() *EditMessageCaption { t.Method = nil; return t }
func (t *EditMessageCaption) GetChatID() ChatID { return t.ChatID }
func (t *EditMessageCaption) SetChatID(v ChatID) { t.ChatID = v }
func (t *EditMessageCaption) WithChatID(v ChatID) *EditMessageCaption { t.ChatID = v; return t }
func (t *EditMessageCaption) GetMessageID() int { if t.MessageID == nil { return 0 } else { return *t.MessageID } }
func (t *EditMessageCaption) SetMessageID(v int) { t.MessageID = &v }
func (t *EditMessageCaption) WithMessageID(v int) *EditMessageCaption { t.MessageID = &v; return t }
func (t *EditMessageCaption) SetNilMessageID() { t.MessageID = nil }
func (t *EditMessageCaption) WithNilMessageID() *EditMessageCaption { t.MessageID = nil; return t }
func (t *EditMessageCaption) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *EditMessageCaption) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *EditMessageCaption) WithInlineMessageID(v string) *EditMessageCaption { t.InlineMessageID = &v; return t }
func (t *EditMessageCaption) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *EditMessageCaption) WithNilInlineMessageID() *EditMessageCaption { t.InlineMessageID = nil; return t }
func (t *EditMessageCaption) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *EditMessageCaption) WithFmtInlineMessageID(f string, a ...interface{}) *EditMessageCaption { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *EditMessageCaption) GetCaption() string { return t.Caption }
func (t *EditMessageCaption) SetCaption(v string) { t.Caption = v }
func (t *EditMessageCaption) WithCaption(v string) *EditMessageCaption { t.Caption = v; return t }
func (t *EditMessageCaption) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *EditMessageCaption) WithFmtCaption(f string, a ...interface{}) *EditMessageCaption { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *EditMessageCaption) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *EditMessageCaption) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *EditMessageCaption) WithReplyMarkup(v *InlineKeyboardMarkup) *EditMessageCaption { t.ReplyMarkup = v; return t }

func (t *EditMessageReplyMarkup) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *EditMessageReplyMarkup) SetMethod(v MethodName) { t.Method = &v }
func (t *EditMessageReplyMarkup) WithMethod(v MethodName) *EditMessageReplyMarkup { t.Method = &v; return t }
func (t *EditMessageReplyMarkup) SetNilMethod() { t.Method = nil }
func (t *EditMessageReplyMarkup) WithNilMethod() *EditMessageReplyMarkup { t.Method = nil; return t }
func (t *EditMessageReplyMarkup) GetChatID() ChatID { return t.ChatID }
func (t *EditMessageReplyMarkup) SetChatID(v ChatID) { t.ChatID = v }
func (t *EditMessageReplyMarkup) WithChatID(v ChatID) *EditMessageReplyMarkup { t.ChatID = v; return t }
func (t *EditMessageReplyMarkup) GetMessageID() int { if t.MessageID == nil { return 0 } else { return *t.MessageID } }
func (t *EditMessageReplyMarkup) SetMessageID(v int) { t.MessageID = &v }
func (t *EditMessageReplyMarkup) WithMessageID(v int) *EditMessageReplyMarkup { t.MessageID = &v; return t }
func (t *EditMessageReplyMarkup) SetNilMessageID() { t.MessageID = nil }
func (t *EditMessageReplyMarkup) WithNilMessageID() *EditMessageReplyMarkup { t.MessageID = nil; return t }
func (t *EditMessageReplyMarkup) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *EditMessageReplyMarkup) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *EditMessageReplyMarkup) WithInlineMessageID(v string) *EditMessageReplyMarkup { t.InlineMessageID = &v; return t }
func (t *EditMessageReplyMarkup) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *EditMessageReplyMarkup) WithNilInlineMessageID() *EditMessageReplyMarkup { t.InlineMessageID = nil; return t }
func (t *EditMessageReplyMarkup) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *EditMessageReplyMarkup) WithFmtInlineMessageID(f string, a ...interface{}) *EditMessageReplyMarkup { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *EditMessageReplyMarkup) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *EditMessageReplyMarkup) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *EditMessageReplyMarkup) WithReplyMarkup(v *InlineKeyboardMarkup) *EditMessageReplyMarkup { t.ReplyMarkup = v; return t }

func (t *InlineQuery) GetID() string { return t.ID }
func (t *InlineQuery) SetID(v string) { t.ID = v }
func (t *InlineQuery) WithID(v string) *InlineQuery { t.ID = v; return t }
func (t *InlineQuery) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQuery) WithFmtID(f string, a ...interface{}) *InlineQuery { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQuery) GetFrom() *User { return t.From }
func (t *InlineQuery) SetFrom(v *User) { t.From = v }
func (t *InlineQuery) WithFrom(v *User) *InlineQuery { t.From = v; return t }
func (t *InlineQuery) GetLocation() *Location { return t.Location }
func (t *InlineQuery) SetLocation(v *Location) { t.Location = v }
func (t *InlineQuery) WithLocation(v *Location) *InlineQuery { t.Location = v; return t }
func (t *InlineQuery) GetQuery() string { return t.Query }
func (t *InlineQuery) SetQuery(v string) { t.Query = v }
func (t *InlineQuery) WithQuery(v string) *InlineQuery { t.Query = v; return t }
func (t *InlineQuery) SetFmtQuery(f string, a ...interface{}) { t.SetQuery(fmt.Sprintf(f, a...)) }
func (t *InlineQuery) WithFmtQuery(f string, a ...interface{}) *InlineQuery { return t.WithQuery(fmt.Sprintf(f, a...)) }
func (t *InlineQuery) GetOffset() string { return t.Offset }
func (t *InlineQuery) SetOffset(v string) { t.Offset = v }
func (t *InlineQuery) WithOffset(v string) *InlineQuery { t.Offset = v; return t }
func (t *InlineQuery) SetFmtOffset(f string, a ...interface{}) { t.SetOffset(fmt.Sprintf(f, a...)) }
func (t *InlineQuery) WithFmtOffset(f string, a ...interface{}) *InlineQuery { return t.WithOffset(fmt.Sprintf(f, a...)) }

func (t *AnswerInlineQuery) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *AnswerInlineQuery) SetMethod(v MethodName) { t.Method = &v }
func (t *AnswerInlineQuery) WithMethod(v MethodName) *AnswerInlineQuery { t.Method = &v; return t }
func (t *AnswerInlineQuery) SetNilMethod() { t.Method = nil }
func (t *AnswerInlineQuery) WithNilMethod() *AnswerInlineQuery { t.Method = nil; return t }
func (t *AnswerInlineQuery) GetInlineQueryID() string { return t.InlineQueryID }
func (t *AnswerInlineQuery) SetInlineQueryID(v string) { t.InlineQueryID = v }
func (t *AnswerInlineQuery) WithInlineQueryID(v string) *AnswerInlineQuery { t.InlineQueryID = v; return t }
func (t *AnswerInlineQuery) SetFmtInlineQueryID(f string, a ...interface{}) { t.SetInlineQueryID(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) WithFmtInlineQueryID(f string, a ...interface{}) *AnswerInlineQuery { return t.WithInlineQueryID(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) GetResults() []InlineQueryResult { return t.Results }
func (t *AnswerInlineQuery) SetResults(v []InlineQueryResult) { t.Results = v }
func (t *AnswerInlineQuery) WithResults(v []InlineQueryResult) *AnswerInlineQuery { t.Results = v; return t }
func (t *AnswerInlineQuery) GetCacheTime() int { if t.CacheTime == nil { return 0 } else { return *t.CacheTime } }
func (t *AnswerInlineQuery) SetCacheTime(v int) { t.CacheTime = &v }
func (t *AnswerInlineQuery) WithCacheTime(v int) *AnswerInlineQuery { t.CacheTime = &v; return t }
func (t *AnswerInlineQuery) SetNilCacheTime() { t.CacheTime = nil }
func (t *AnswerInlineQuery) WithNilCacheTime() *AnswerInlineQuery { t.CacheTime = nil; return t }
func (t *AnswerInlineQuery) GetIsPersonal() bool { if t.IsPersonal == nil { return false } else { return *t.IsPersonal } }
func (t *AnswerInlineQuery) SetIsPersonal(v bool) { t.IsPersonal = &v }
func (t *AnswerInlineQuery) WithIsPersonal(v bool) *AnswerInlineQuery { t.IsPersonal = &v; return t }
func (t *AnswerInlineQuery) SetNilIsPersonal() { t.IsPersonal = nil }
func (t *AnswerInlineQuery) WithNilIsPersonal() *AnswerInlineQuery { t.IsPersonal = nil; return t }
func (t *AnswerInlineQuery) GetNextOffset() string { if t.NextOffset == nil { return "" } else { return *t.NextOffset } }
func (t *AnswerInlineQuery) SetNextOffset(v string) { t.NextOffset = &v }
func (t *AnswerInlineQuery) WithNextOffset(v string) *AnswerInlineQuery { t.NextOffset = &v; return t }
func (t *AnswerInlineQuery) SetNilNextOffset() { t.NextOffset = nil }
func (t *AnswerInlineQuery) WithNilNextOffset() *AnswerInlineQuery { t.NextOffset = nil; return t }
func (t *AnswerInlineQuery) SetFmtNextOffset(f string, a ...interface{}) { t.SetNextOffset(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) WithFmtNextOffset(f string, a ...interface{}) *AnswerInlineQuery { return t.WithNextOffset(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) GetSwitchPMText() string { if t.SwitchPMText == nil { return "" } else { return *t.SwitchPMText } }
func (t *AnswerInlineQuery) SetSwitchPMText(v string) { t.SwitchPMText = &v }
func (t *AnswerInlineQuery) WithSwitchPMText(v string) *AnswerInlineQuery { t.SwitchPMText = &v; return t }
func (t *AnswerInlineQuery) SetNilSwitchPMText() { t.SwitchPMText = nil }
func (t *AnswerInlineQuery) WithNilSwitchPMText() *AnswerInlineQuery { t.SwitchPMText = nil; return t }
func (t *AnswerInlineQuery) SetFmtSwitchPMText(f string, a ...interface{}) { t.SetSwitchPMText(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) WithFmtSwitchPMText(f string, a ...interface{}) *AnswerInlineQuery { return t.WithSwitchPMText(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) GetSwitchPMParameter() string { if t.SwitchPMParameter == nil { return "" } else { return *t.SwitchPMParameter } }
func (t *AnswerInlineQuery) SetSwitchPMParameter(v string) { t.SwitchPMParameter = &v }
func (t *AnswerInlineQuery) WithSwitchPMParameter(v string) *AnswerInlineQuery { t.SwitchPMParameter = &v; return t }
func (t *AnswerInlineQuery) SetNilSwitchPMParameter() { t.SwitchPMParameter = nil }
func (t *AnswerInlineQuery) WithNilSwitchPMParameter() *AnswerInlineQuery { t.SwitchPMParameter = nil; return t }
func (t *AnswerInlineQuery) SetFmtSwitchPMParameter(f string, a ...interface{}) { t.SetSwitchPMParameter(fmt.Sprintf(f, a...)) }
func (t *AnswerInlineQuery) WithFmtSwitchPMParameter(f string, a ...interface{}) *AnswerInlineQuery { return t.WithSwitchPMParameter(fmt.Sprintf(f, a...)) }

func (t *InlineQueryResultArticle) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultArticle) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultArticle) WithType(v InlineQueryResultType) *InlineQueryResultArticle { t.Type = v; return t }
func (t *InlineQueryResultArticle) GetID() string { return t.ID }
func (t *InlineQueryResultArticle) SetID(v string) { t.ID = v }
func (t *InlineQueryResultArticle) WithID(v string) *InlineQueryResultArticle { t.ID = v; return t }
func (t *InlineQueryResultArticle) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) WithFmtID(f string, a ...interface{}) *InlineQueryResultArticle { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) GetTitle() string { return t.Title }
func (t *InlineQueryResultArticle) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultArticle) WithTitle(v string) *InlineQueryResultArticle { t.Title = v; return t }
func (t *InlineQueryResultArticle) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultArticle { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultArticle) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultArticle) WithInputMessageContent(v InputMessageContent) *InlineQueryResultArticle { t.InputMessageContent = v; return t }
func (t *InlineQueryResultArticle) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultArticle) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultArticle) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultArticle { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultArticle) GetURL() string { if t.URL == nil { return "" } else { return *t.URL } }
func (t *InlineQueryResultArticle) SetURL(v string) { t.URL = &v }
func (t *InlineQueryResultArticle) WithURL(v string) *InlineQueryResultArticle { t.URL = &v; return t }
func (t *InlineQueryResultArticle) SetNilURL() { t.URL = nil }
func (t *InlineQueryResultArticle) WithNilURL() *InlineQueryResultArticle { t.URL = nil; return t }
func (t *InlineQueryResultArticle) SetFmtURL(f string, a ...interface{}) { t.SetURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) WithFmtURL(f string, a ...interface{}) *InlineQueryResultArticle { return t.WithURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) GetHideURL() bool { if t.HideURL == nil { return false } else { return *t.HideURL } }
func (t *InlineQueryResultArticle) SetHideURL(v bool) { t.HideURL = &v }
func (t *InlineQueryResultArticle) WithHideURL(v bool) *InlineQueryResultArticle { t.HideURL = &v; return t }
func (t *InlineQueryResultArticle) SetNilHideURL() { t.HideURL = nil }
func (t *InlineQueryResultArticle) WithNilHideURL() *InlineQueryResultArticle { t.HideURL = nil; return t }
func (t *InlineQueryResultArticle) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultArticle) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultArticle) WithDescription(v string) *InlineQueryResultArticle { t.Description = &v; return t }
func (t *InlineQueryResultArticle) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultArticle) WithNilDescription() *InlineQueryResultArticle { t.Description = nil; return t }
func (t *InlineQueryResultArticle) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultArticle { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) GetThumbURL() string { if t.ThumbURL == nil { return "" } else { return *t.ThumbURL } }
func (t *InlineQueryResultArticle) SetThumbURL(v string) { t.ThumbURL = &v }
func (t *InlineQueryResultArticle) WithThumbURL(v string) *InlineQueryResultArticle { t.ThumbURL = &v; return t }
func (t *InlineQueryResultArticle) SetNilThumbURL() { t.ThumbURL = nil }
func (t *InlineQueryResultArticle) WithNilThumbURL() *InlineQueryResultArticle { t.ThumbURL = nil; return t }
func (t *InlineQueryResultArticle) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultArticle { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultArticle) GetThumbWidth() int { if t.ThumbWidth == nil { return 0 } else { return *t.ThumbWidth } }
func (t *InlineQueryResultArticle) SetThumbWidth(v int) { t.ThumbWidth = &v }
func (t *InlineQueryResultArticle) WithThumbWidth(v int) *InlineQueryResultArticle { t.ThumbWidth = &v; return t }
func (t *InlineQueryResultArticle) SetNilThumbWidth() { t.ThumbWidth = nil }
func (t *InlineQueryResultArticle) WithNilThumbWidth() *InlineQueryResultArticle { t.ThumbWidth = nil; return t }
func (t *InlineQueryResultArticle) GetThumbHeight() int { if t.ThumbHeight == nil { return 0 } else { return *t.ThumbHeight } }
func (t *InlineQueryResultArticle) SetThumbHeight(v int) { t.ThumbHeight = &v }
func (t *InlineQueryResultArticle) WithThumbHeight(v int) *InlineQueryResultArticle { t.ThumbHeight = &v; return t }
func (t *InlineQueryResultArticle) SetNilThumbHeight() { t.ThumbHeight = nil }
func (t *InlineQueryResultArticle) WithNilThumbHeight() *InlineQueryResultArticle { t.ThumbHeight = nil; return t }

func (t *InlineQueryResultPhoto) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultPhoto) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultPhoto) WithType(v InlineQueryResultType) *InlineQueryResultPhoto { t.Type = v; return t }
func (t *InlineQueryResultPhoto) GetID() string { return t.ID }
func (t *InlineQueryResultPhoto) SetID(v string) { t.ID = v }
func (t *InlineQueryResultPhoto) WithID(v string) *InlineQueryResultPhoto { t.ID = v; return t }
func (t *InlineQueryResultPhoto) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) WithFmtID(f string, a ...interface{}) *InlineQueryResultPhoto { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) GetPhotoURL() string { return t.PhotoURL }
func (t *InlineQueryResultPhoto) SetPhotoURL(v string) { t.PhotoURL = v }
func (t *InlineQueryResultPhoto) WithPhotoURL(v string) *InlineQueryResultPhoto { t.PhotoURL = v; return t }
func (t *InlineQueryResultPhoto) SetFmtPhotoURL(f string, a ...interface{}) { t.SetPhotoURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) WithFmtPhotoURL(f string, a ...interface{}) *InlineQueryResultPhoto { return t.WithPhotoURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) GetThumbURL() string { return t.ThumbURL }
func (t *InlineQueryResultPhoto) SetThumbURL(v string) { t.ThumbURL = v }
func (t *InlineQueryResultPhoto) WithThumbURL(v string) *InlineQueryResultPhoto { t.ThumbURL = v; return t }
func (t *InlineQueryResultPhoto) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultPhoto { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) GetPhotoWidth() int { if t.PhotoWidth == nil { return 0 } else { return *t.PhotoWidth } }
func (t *InlineQueryResultPhoto) SetPhotoWidth(v int) { t.PhotoWidth = &v }
func (t *InlineQueryResultPhoto) WithPhotoWidth(v int) *InlineQueryResultPhoto { t.PhotoWidth = &v; return t }
func (t *InlineQueryResultPhoto) SetNilPhotoWidth() { t.PhotoWidth = nil }
func (t *InlineQueryResultPhoto) WithNilPhotoWidth() *InlineQueryResultPhoto { t.PhotoWidth = nil; return t }
func (t *InlineQueryResultPhoto) GetPhotoHeight() int { if t.PhotoHeight == nil { return 0 } else { return *t.PhotoHeight } }
func (t *InlineQueryResultPhoto) SetPhotoHeight(v int) { t.PhotoHeight = &v }
func (t *InlineQueryResultPhoto) WithPhotoHeight(v int) *InlineQueryResultPhoto { t.PhotoHeight = &v; return t }
func (t *InlineQueryResultPhoto) SetNilPhotoHeight() { t.PhotoHeight = nil }
func (t *InlineQueryResultPhoto) WithNilPhotoHeight() *InlineQueryResultPhoto { t.PhotoHeight = nil; return t }
func (t *InlineQueryResultPhoto) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *InlineQueryResultPhoto) SetTitle(v string) { t.Title = &v }
func (t *InlineQueryResultPhoto) WithTitle(v string) *InlineQueryResultPhoto { t.Title = &v; return t }
func (t *InlineQueryResultPhoto) SetNilTitle() { t.Title = nil }
func (t *InlineQueryResultPhoto) WithNilTitle() *InlineQueryResultPhoto { t.Title = nil; return t }
func (t *InlineQueryResultPhoto) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultPhoto { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultPhoto) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultPhoto) WithDescription(v string) *InlineQueryResultPhoto { t.Description = &v; return t }
func (t *InlineQueryResultPhoto) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultPhoto) WithNilDescription() *InlineQueryResultPhoto { t.Description = nil; return t }
func (t *InlineQueryResultPhoto) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultPhoto { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultPhoto) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultPhoto) WithCaption(v string) *InlineQueryResultPhoto { t.Caption = &v; return t }
func (t *InlineQueryResultPhoto) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultPhoto) WithNilCaption() *InlineQueryResultPhoto { t.Caption = nil; return t }
func (t *InlineQueryResultPhoto) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultPhoto { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultPhoto) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultPhoto) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultPhoto) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultPhoto { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultPhoto) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultPhoto) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultPhoto) WithInputMessageContent(v InputMessageContent) *InlineQueryResultPhoto { t.InputMessageContent = v; return t }

func (t *InlineQueryResultGIF) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultGIF) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultGIF) WithType(v InlineQueryResultType) *InlineQueryResultGIF { t.Type = v; return t }
func (t *InlineQueryResultGIF) GetID() string { return t.ID }
func (t *InlineQueryResultGIF) SetID(v string) { t.ID = v }
func (t *InlineQueryResultGIF) WithID(v string) *InlineQueryResultGIF { t.ID = v; return t }
func (t *InlineQueryResultGIF) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) WithFmtID(f string, a ...interface{}) *InlineQueryResultGIF { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) GetGIFURL() string { return t.GIFURL }
func (t *InlineQueryResultGIF) SetGIFURL(v string) { t.GIFURL = v }
func (t *InlineQueryResultGIF) WithGIFURL(v string) *InlineQueryResultGIF { t.GIFURL = v; return t }
func (t *InlineQueryResultGIF) SetFmtGIFURL(f string, a ...interface{}) { t.SetGIFURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) WithFmtGIFURL(f string, a ...interface{}) *InlineQueryResultGIF { return t.WithGIFURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) GetGIFWidth() int { if t.GIFWidth == nil { return 0 } else { return *t.GIFWidth } }
func (t *InlineQueryResultGIF) SetGIFWidth(v int) { t.GIFWidth = &v }
func (t *InlineQueryResultGIF) WithGIFWidth(v int) *InlineQueryResultGIF { t.GIFWidth = &v; return t }
func (t *InlineQueryResultGIF) SetNilGIFWidth() { t.GIFWidth = nil }
func (t *InlineQueryResultGIF) WithNilGIFWidth() *InlineQueryResultGIF { t.GIFWidth = nil; return t }
func (t *InlineQueryResultGIF) GetGIFHeight() int { if t.GIFHeight == nil { return 0 } else { return *t.GIFHeight } }
func (t *InlineQueryResultGIF) SetGIFHeight(v int) { t.GIFHeight = &v }
func (t *InlineQueryResultGIF) WithGIFHeight(v int) *InlineQueryResultGIF { t.GIFHeight = &v; return t }
func (t *InlineQueryResultGIF) SetNilGIFHeight() { t.GIFHeight = nil }
func (t *InlineQueryResultGIF) WithNilGIFHeight() *InlineQueryResultGIF { t.GIFHeight = nil; return t }
func (t *InlineQueryResultGIF) GetThumbURL() string { return t.ThumbURL }
func (t *InlineQueryResultGIF) SetThumbURL(v string) { t.ThumbURL = v }
func (t *InlineQueryResultGIF) WithThumbURL(v string) *InlineQueryResultGIF { t.ThumbURL = v; return t }
func (t *InlineQueryResultGIF) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultGIF { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *InlineQueryResultGIF) SetTitle(v string) { t.Title = &v }
func (t *InlineQueryResultGIF) WithTitle(v string) *InlineQueryResultGIF { t.Title = &v; return t }
func (t *InlineQueryResultGIF) SetNilTitle() { t.Title = nil }
func (t *InlineQueryResultGIF) WithNilTitle() *InlineQueryResultGIF { t.Title = nil; return t }
func (t *InlineQueryResultGIF) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultGIF { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultGIF) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultGIF) WithCaption(v string) *InlineQueryResultGIF { t.Caption = &v; return t }
func (t *InlineQueryResultGIF) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultGIF) WithNilCaption() *InlineQueryResultGIF { t.Caption = nil; return t }
func (t *InlineQueryResultGIF) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultGIF { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGIF) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultGIF) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultGIF) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultGIF { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultGIF) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultGIF) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultGIF) WithInputMessageContent(v InputMessageContent) *InlineQueryResultGIF { t.InputMessageContent = v; return t }

func (t *InlineQueryResultMPEG4GIF) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultMPEG4GIF) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultMPEG4GIF) WithType(v InlineQueryResultType) *InlineQueryResultMPEG4GIF { t.Type = v; return t }
func (t *InlineQueryResultMPEG4GIF) GetID() string { return t.ID }
func (t *InlineQueryResultMPEG4GIF) SetID(v string) { t.ID = v }
func (t *InlineQueryResultMPEG4GIF) WithID(v string) *InlineQueryResultMPEG4GIF { t.ID = v; return t }
func (t *InlineQueryResultMPEG4GIF) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) WithFmtID(f string, a ...interface{}) *InlineQueryResultMPEG4GIF { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) GetMPEG4URL() string { return t.MPEG4URL }
func (t *InlineQueryResultMPEG4GIF) SetMPEG4URL(v string) { t.MPEG4URL = v }
func (t *InlineQueryResultMPEG4GIF) WithMPEG4URL(v string) *InlineQueryResultMPEG4GIF { t.MPEG4URL = v; return t }
func (t *InlineQueryResultMPEG4GIF) SetFmtMPEG4URL(f string, a ...interface{}) { t.SetMPEG4URL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) WithFmtMPEG4URL(f string, a ...interface{}) *InlineQueryResultMPEG4GIF { return t.WithMPEG4URL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) GetMPEG4Width() int { if t.MPEG4Width == nil { return 0 } else { return *t.MPEG4Width } }
func (t *InlineQueryResultMPEG4GIF) SetMPEG4Width(v int) { t.MPEG4Width = &v }
func (t *InlineQueryResultMPEG4GIF) WithMPEG4Width(v int) *InlineQueryResultMPEG4GIF { t.MPEG4Width = &v; return t }
func (t *InlineQueryResultMPEG4GIF) SetNilMPEG4Width() { t.MPEG4Width = nil }
func (t *InlineQueryResultMPEG4GIF) WithNilMPEG4Width() *InlineQueryResultMPEG4GIF { t.MPEG4Width = nil; return t }
func (t *InlineQueryResultMPEG4GIF) GetMPEG4Height() int { if t.MPEG4Height == nil { return 0 } else { return *t.MPEG4Height } }
func (t *InlineQueryResultMPEG4GIF) SetMPEG4Height(v int) { t.MPEG4Height = &v }
func (t *InlineQueryResultMPEG4GIF) WithMPEG4Height(v int) *InlineQueryResultMPEG4GIF { t.MPEG4Height = &v; return t }
func (t *InlineQueryResultMPEG4GIF) SetNilMPEG4Height() { t.MPEG4Height = nil }
func (t *InlineQueryResultMPEG4GIF) WithNilMPEG4Height() *InlineQueryResultMPEG4GIF { t.MPEG4Height = nil; return t }
func (t *InlineQueryResultMPEG4GIF) GetThumbURL() string { return t.ThumbURL }
func (t *InlineQueryResultMPEG4GIF) SetThumbURL(v string) { t.ThumbURL = v }
func (t *InlineQueryResultMPEG4GIF) WithThumbURL(v string) *InlineQueryResultMPEG4GIF { t.ThumbURL = v; return t }
func (t *InlineQueryResultMPEG4GIF) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultMPEG4GIF { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *InlineQueryResultMPEG4GIF) SetTitle(v string) { t.Title = &v }
func (t *InlineQueryResultMPEG4GIF) WithTitle(v string) *InlineQueryResultMPEG4GIF { t.Title = &v; return t }
func (t *InlineQueryResultMPEG4GIF) SetNilTitle() { t.Title = nil }
func (t *InlineQueryResultMPEG4GIF) WithNilTitle() *InlineQueryResultMPEG4GIF { t.Title = nil; return t }
func (t *InlineQueryResultMPEG4GIF) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultMPEG4GIF { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultMPEG4GIF) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultMPEG4GIF) WithCaption(v string) *InlineQueryResultMPEG4GIF { t.Caption = &v; return t }
func (t *InlineQueryResultMPEG4GIF) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultMPEG4GIF) WithNilCaption() *InlineQueryResultMPEG4GIF { t.Caption = nil; return t }
func (t *InlineQueryResultMPEG4GIF) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultMPEG4GIF { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultMPEG4GIF) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultMPEG4GIF) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultMPEG4GIF) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultMPEG4GIF { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultMPEG4GIF) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultMPEG4GIF) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultMPEG4GIF) WithInputMessageContent(v InputMessageContent) *InlineQueryResultMPEG4GIF { t.InputMessageContent = v; return t }

func (t *InlineQueryResultVideo) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultVideo) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultVideo) WithType(v InlineQueryResultType) *InlineQueryResultVideo { t.Type = v; return t }
func (t *InlineQueryResultVideo) GetID() string { return t.ID }
func (t *InlineQueryResultVideo) SetID(v string) { t.ID = v }
func (t *InlineQueryResultVideo) WithID(v string) *InlineQueryResultVideo { t.ID = v; return t }
func (t *InlineQueryResultVideo) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtID(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetVideoURL() string { return t.VideoURL }
func (t *InlineQueryResultVideo) SetVideoURL(v string) { t.VideoURL = v }
func (t *InlineQueryResultVideo) WithVideoURL(v string) *InlineQueryResultVideo { t.VideoURL = v; return t }
func (t *InlineQueryResultVideo) SetFmtVideoURL(f string, a ...interface{}) { t.SetVideoURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtVideoURL(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithVideoURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetMimeType() string { return t.MimeType }
func (t *InlineQueryResultVideo) SetMimeType(v string) { t.MimeType = v }
func (t *InlineQueryResultVideo) WithMimeType(v string) *InlineQueryResultVideo { t.MimeType = v; return t }
func (t *InlineQueryResultVideo) SetFmtMimeType(f string, a ...interface{}) { t.SetMimeType(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtMimeType(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithMimeType(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetThumbURL() string { return t.ThumbURL }
func (t *InlineQueryResultVideo) SetThumbURL(v string) { t.ThumbURL = v }
func (t *InlineQueryResultVideo) WithThumbURL(v string) *InlineQueryResultVideo { t.ThumbURL = v; return t }
func (t *InlineQueryResultVideo) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetTitle() string { return t.Title }
func (t *InlineQueryResultVideo) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultVideo) WithTitle(v string) *InlineQueryResultVideo { t.Title = v; return t }
func (t *InlineQueryResultVideo) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultVideo) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultVideo) WithCaption(v string) *InlineQueryResultVideo { t.Caption = &v; return t }
func (t *InlineQueryResultVideo) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultVideo) WithNilCaption() *InlineQueryResultVideo { t.Caption = nil; return t }
func (t *InlineQueryResultVideo) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetVideoWidth() int { if t.VideoWidth == nil { return 0 } else { return *t.VideoWidth } }
func (t *InlineQueryResultVideo) SetVideoWidth(v int) { t.VideoWidth = &v }
func (t *InlineQueryResultVideo) WithVideoWidth(v int) *InlineQueryResultVideo { t.VideoWidth = &v; return t }
func (t *InlineQueryResultVideo) SetNilVideoWidth() { t.VideoWidth = nil }
func (t *InlineQueryResultVideo) WithNilVideoWidth() *InlineQueryResultVideo { t.VideoWidth = nil; return t }
func (t *InlineQueryResultVideo) GetVideoHeight() int { if t.VideoHeight == nil { return 0 } else { return *t.VideoHeight } }
func (t *InlineQueryResultVideo) SetVideoHeight(v int) { t.VideoHeight = &v }
func (t *InlineQueryResultVideo) WithVideoHeight(v int) *InlineQueryResultVideo { t.VideoHeight = &v; return t }
func (t *InlineQueryResultVideo) SetNilVideoHeight() { t.VideoHeight = nil }
func (t *InlineQueryResultVideo) WithNilVideoHeight() *InlineQueryResultVideo { t.VideoHeight = nil; return t }
func (t *InlineQueryResultVideo) GetVideoDuration() int { if t.VideoDuration == nil { return 0 } else { return *t.VideoDuration } }
func (t *InlineQueryResultVideo) SetVideoDuration(v int) { t.VideoDuration = &v }
func (t *InlineQueryResultVideo) WithVideoDuration(v int) *InlineQueryResultVideo { t.VideoDuration = &v; return t }
func (t *InlineQueryResultVideo) SetNilVideoDuration() { t.VideoDuration = nil }
func (t *InlineQueryResultVideo) WithNilVideoDuration() *InlineQueryResultVideo { t.VideoDuration = nil; return t }
func (t *InlineQueryResultVideo) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultVideo) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultVideo) WithDescription(v string) *InlineQueryResultVideo { t.Description = &v; return t }
func (t *InlineQueryResultVideo) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultVideo) WithNilDescription() *InlineQueryResultVideo { t.Description = nil; return t }
func (t *InlineQueryResultVideo) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultVideo { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVideo) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultVideo) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultVideo) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultVideo { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultVideo) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultVideo) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultVideo) WithInputMessageContent(v InputMessageContent) *InlineQueryResultVideo { t.InputMessageContent = v; return t }

func (t *InlineQueryResultAudio) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultAudio) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultAudio) WithType(v InlineQueryResultType) *InlineQueryResultAudio { t.Type = v; return t }
func (t *InlineQueryResultAudio) GetID() string { return t.ID }
func (t *InlineQueryResultAudio) SetID(v string) { t.ID = v }
func (t *InlineQueryResultAudio) WithID(v string) *InlineQueryResultAudio { t.ID = v; return t }
func (t *InlineQueryResultAudio) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) WithFmtID(f string, a ...interface{}) *InlineQueryResultAudio { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) GetAudioURL() string { return t.AudioURL }
func (t *InlineQueryResultAudio) SetAudioURL(v string) { t.AudioURL = v }
func (t *InlineQueryResultAudio) WithAudioURL(v string) *InlineQueryResultAudio { t.AudioURL = v; return t }
func (t *InlineQueryResultAudio) SetFmtAudioURL(f string, a ...interface{}) { t.SetAudioURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) WithFmtAudioURL(f string, a ...interface{}) *InlineQueryResultAudio { return t.WithAudioURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) GetTitle() string { return t.Title }
func (t *InlineQueryResultAudio) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultAudio) WithTitle(v string) *InlineQueryResultAudio { t.Title = v; return t }
func (t *InlineQueryResultAudio) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultAudio { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultAudio) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultAudio) WithCaption(v string) *InlineQueryResultAudio { t.Caption = &v; return t }
func (t *InlineQueryResultAudio) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultAudio) WithNilCaption() *InlineQueryResultAudio { t.Caption = nil; return t }
func (t *InlineQueryResultAudio) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultAudio { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) GetPerformer() string { if t.Performer == nil { return "" } else { return *t.Performer } }
func (t *InlineQueryResultAudio) SetPerformer(v string) { t.Performer = &v }
func (t *InlineQueryResultAudio) WithPerformer(v string) *InlineQueryResultAudio { t.Performer = &v; return t }
func (t *InlineQueryResultAudio) SetNilPerformer() { t.Performer = nil }
func (t *InlineQueryResultAudio) WithNilPerformer() *InlineQueryResultAudio { t.Performer = nil; return t }
func (t *InlineQueryResultAudio) SetFmtPerformer(f string, a ...interface{}) { t.SetPerformer(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) WithFmtPerformer(f string, a ...interface{}) *InlineQueryResultAudio { return t.WithPerformer(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultAudio) GetAudioDuration() int { if t.AudioDuration == nil { return 0 } else { return *t.AudioDuration } }
func (t *InlineQueryResultAudio) SetAudioDuration(v int) { t.AudioDuration = &v }
func (t *InlineQueryResultAudio) WithAudioDuration(v int) *InlineQueryResultAudio { t.AudioDuration = &v; return t }
func (t *InlineQueryResultAudio) SetNilAudioDuration() { t.AudioDuration = nil }
func (t *InlineQueryResultAudio) WithNilAudioDuration() *InlineQueryResultAudio { t.AudioDuration = nil; return t }
func (t *InlineQueryResultAudio) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultAudio) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultAudio) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultAudio { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultAudio) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultAudio) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultAudio) WithInputMessageContent(v InputMessageContent) *InlineQueryResultAudio { t.InputMessageContent = v; return t }

func (t *InlineQueryResultVoice) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultVoice) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultVoice) WithType(v InlineQueryResultType) *InlineQueryResultVoice { t.Type = v; return t }
func (t *InlineQueryResultVoice) GetID() string { return t.ID }
func (t *InlineQueryResultVoice) SetID(v string) { t.ID = v }
func (t *InlineQueryResultVoice) WithID(v string) *InlineQueryResultVoice { t.ID = v; return t }
func (t *InlineQueryResultVoice) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) WithFmtID(f string, a ...interface{}) *InlineQueryResultVoice { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) GetVoiceURL() string { return t.VoiceURL }
func (t *InlineQueryResultVoice) SetVoiceURL(v string) { t.VoiceURL = v }
func (t *InlineQueryResultVoice) WithVoiceURL(v string) *InlineQueryResultVoice { t.VoiceURL = v; return t }
func (t *InlineQueryResultVoice) SetFmtVoiceURL(f string, a ...interface{}) { t.SetVoiceURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) WithFmtVoiceURL(f string, a ...interface{}) *InlineQueryResultVoice { return t.WithVoiceURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) GetTitle() string { return t.Title }
func (t *InlineQueryResultVoice) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultVoice) WithTitle(v string) *InlineQueryResultVoice { t.Title = v; return t }
func (t *InlineQueryResultVoice) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultVoice { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultVoice) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultVoice) WithCaption(v string) *InlineQueryResultVoice { t.Caption = &v; return t }
func (t *InlineQueryResultVoice) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultVoice) WithNilCaption() *InlineQueryResultVoice { t.Caption = nil; return t }
func (t *InlineQueryResultVoice) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultVoice { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVoice) GetVoiceDuration() int { if t.VoiceDuration == nil { return 0 } else { return *t.VoiceDuration } }
func (t *InlineQueryResultVoice) SetVoiceDuration(v int) { t.VoiceDuration = &v }
func (t *InlineQueryResultVoice) WithVoiceDuration(v int) *InlineQueryResultVoice { t.VoiceDuration = &v; return t }
func (t *InlineQueryResultVoice) SetNilVoiceDuration() { t.VoiceDuration = nil }
func (t *InlineQueryResultVoice) WithNilVoiceDuration() *InlineQueryResultVoice { t.VoiceDuration = nil; return t }
func (t *InlineQueryResultVoice) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultVoice) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultVoice) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultVoice { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultVoice) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultVoice) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultVoice) WithInputMessageContent(v InputMessageContent) *InlineQueryResultVoice { t.InputMessageContent = v; return t }

func (t *InlineQueryResultDocument) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultDocument) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultDocument) WithType(v InlineQueryResultType) *InlineQueryResultDocument { t.Type = v; return t }
func (t *InlineQueryResultDocument) GetID() string { return t.ID }
func (t *InlineQueryResultDocument) SetID(v string) { t.ID = v }
func (t *InlineQueryResultDocument) WithID(v string) *InlineQueryResultDocument { t.ID = v; return t }
func (t *InlineQueryResultDocument) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtID(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetTitle() string { return t.Title }
func (t *InlineQueryResultDocument) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultDocument) WithTitle(v string) *InlineQueryResultDocument { t.Title = v; return t }
func (t *InlineQueryResultDocument) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultDocument) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultDocument) WithCaption(v string) *InlineQueryResultDocument { t.Caption = &v; return t }
func (t *InlineQueryResultDocument) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultDocument) WithNilCaption() *InlineQueryResultDocument { t.Caption = nil; return t }
func (t *InlineQueryResultDocument) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetDocumentURL() string { return t.DocumentURL }
func (t *InlineQueryResultDocument) SetDocumentURL(v string) { t.DocumentURL = v }
func (t *InlineQueryResultDocument) WithDocumentURL(v string) *InlineQueryResultDocument { t.DocumentURL = v; return t }
func (t *InlineQueryResultDocument) SetFmtDocumentURL(f string, a ...interface{}) { t.SetDocumentURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtDocumentURL(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithDocumentURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetMimeType() string { return t.MimeType }
func (t *InlineQueryResultDocument) SetMimeType(v string) { t.MimeType = v }
func (t *InlineQueryResultDocument) WithMimeType(v string) *InlineQueryResultDocument { t.MimeType = v; return t }
func (t *InlineQueryResultDocument) SetFmtMimeType(f string, a ...interface{}) { t.SetMimeType(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtMimeType(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithMimeType(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultDocument) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultDocument) WithDescription(v string) *InlineQueryResultDocument { t.Description = &v; return t }
func (t *InlineQueryResultDocument) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultDocument) WithNilDescription() *InlineQueryResultDocument { t.Description = nil; return t }
func (t *InlineQueryResultDocument) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultDocument) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultDocument) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultDocument { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultDocument) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultDocument) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultDocument) WithInputMessageContent(v InputMessageContent) *InlineQueryResultDocument { t.InputMessageContent = v; return t }
func (t *InlineQueryResultDocument) GetThumbURL() string { if t.ThumbURL == nil { return "" } else { return *t.ThumbURL } }
func (t *InlineQueryResultDocument) SetThumbURL(v string) { t.ThumbURL = &v }
func (t *InlineQueryResultDocument) WithThumbURL(v string) *InlineQueryResultDocument { t.ThumbURL = &v; return t }
func (t *InlineQueryResultDocument) SetNilThumbURL() { t.ThumbURL = nil }
func (t *InlineQueryResultDocument) WithNilThumbURL() *InlineQueryResultDocument { t.ThumbURL = nil; return t }
func (t *InlineQueryResultDocument) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultDocument { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultDocument) GetThumbWidth() int { if t.ThumbWidth == nil { return 0 } else { return *t.ThumbWidth } }
func (t *InlineQueryResultDocument) SetThumbWidth(v int) { t.ThumbWidth = &v }
func (t *InlineQueryResultDocument) WithThumbWidth(v int) *InlineQueryResultDocument { t.ThumbWidth = &v; return t }
func (t *InlineQueryResultDocument) SetNilThumbWidth() { t.ThumbWidth = nil }
func (t *InlineQueryResultDocument) WithNilThumbWidth() *InlineQueryResultDocument { t.ThumbWidth = nil; return t }
func (t *InlineQueryResultDocument) GetThumbHeight() int { if t.ThumbHeight == nil { return 0 } else { return *t.ThumbHeight } }
func (t *InlineQueryResultDocument) SetThumbHeight(v int) { t.ThumbHeight = &v }
func (t *InlineQueryResultDocument) WithThumbHeight(v int) *InlineQueryResultDocument { t.ThumbHeight = &v; return t }
func (t *InlineQueryResultDocument) SetNilThumbHeight() { t.ThumbHeight = nil }
func (t *InlineQueryResultDocument) WithNilThumbHeight() *InlineQueryResultDocument { t.ThumbHeight = nil; return t }

func (t *InlineQueryResultLocation) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultLocation) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultLocation) WithType(v InlineQueryResultType) *InlineQueryResultLocation { t.Type = v; return t }
func (t *InlineQueryResultLocation) GetID() string { return t.ID }
func (t *InlineQueryResultLocation) SetID(v string) { t.ID = v }
func (t *InlineQueryResultLocation) WithID(v string) *InlineQueryResultLocation { t.ID = v; return t }
func (t *InlineQueryResultLocation) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultLocation) WithFmtID(f string, a ...interface{}) *InlineQueryResultLocation { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultLocation) GetLatitude() float64 { return t.Latitude }
func (t *InlineQueryResultLocation) SetLatitude(v float64) { t.Latitude = v }
func (t *InlineQueryResultLocation) WithLatitude(v float64) *InlineQueryResultLocation { t.Latitude = v; return t }
func (t *InlineQueryResultLocation) GetLongitude() float64 { return t.Longitude }
func (t *InlineQueryResultLocation) SetLongitude(v float64) { t.Longitude = v }
func (t *InlineQueryResultLocation) WithLongitude(v float64) *InlineQueryResultLocation { t.Longitude = v; return t }
func (t *InlineQueryResultLocation) GetTitle() string { return t.Title }
func (t *InlineQueryResultLocation) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultLocation) WithTitle(v string) *InlineQueryResultLocation { t.Title = v; return t }
func (t *InlineQueryResultLocation) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultLocation) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultLocation { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultLocation) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultLocation) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultLocation) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultLocation { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultLocation) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultLocation) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultLocation) WithInputMessageContent(v InputMessageContent) *InlineQueryResultLocation { t.InputMessageContent = v; return t }
func (t *InlineQueryResultLocation) GetThumbURL() string { if t.ThumbURL == nil { return "" } else { return *t.ThumbURL } }
func (t *InlineQueryResultLocation) SetThumbURL(v string) { t.ThumbURL = &v }
func (t *InlineQueryResultLocation) WithThumbURL(v string) *InlineQueryResultLocation { t.ThumbURL = &v; return t }
func (t *InlineQueryResultLocation) SetNilThumbURL() { t.ThumbURL = nil }
func (t *InlineQueryResultLocation) WithNilThumbURL() *InlineQueryResultLocation { t.ThumbURL = nil; return t }
func (t *InlineQueryResultLocation) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultLocation) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultLocation { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultLocation) GetThumbWidth() int { if t.ThumbWidth == nil { return 0 } else { return *t.ThumbWidth } }
func (t *InlineQueryResultLocation) SetThumbWidth(v int) { t.ThumbWidth = &v }
func (t *InlineQueryResultLocation) WithThumbWidth(v int) *InlineQueryResultLocation { t.ThumbWidth = &v; return t }
func (t *InlineQueryResultLocation) SetNilThumbWidth() { t.ThumbWidth = nil }
func (t *InlineQueryResultLocation) WithNilThumbWidth() *InlineQueryResultLocation { t.ThumbWidth = nil; return t }
func (t *InlineQueryResultLocation) GetThumbHeight() int { if t.ThumbHeight == nil { return 0 } else { return *t.ThumbHeight } }
func (t *InlineQueryResultLocation) SetThumbHeight(v int) { t.ThumbHeight = &v }
func (t *InlineQueryResultLocation) WithThumbHeight(v int) *InlineQueryResultLocation { t.ThumbHeight = &v; return t }
func (t *InlineQueryResultLocation) SetNilThumbHeight() { t.ThumbHeight = nil }
func (t *InlineQueryResultLocation) WithNilThumbHeight() *InlineQueryResultLocation { t.ThumbHeight = nil; return t }

func (t *InlineQueryResultVenue) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultVenue) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultVenue) WithType(v InlineQueryResultType) *InlineQueryResultVenue { t.Type = v; return t }
func (t *InlineQueryResultVenue) GetID() string { return t.ID }
func (t *InlineQueryResultVenue) SetID(v string) { t.ID = v }
func (t *InlineQueryResultVenue) WithID(v string) *InlineQueryResultVenue { t.ID = v; return t }
func (t *InlineQueryResultVenue) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) WithFmtID(f string, a ...interface{}) *InlineQueryResultVenue { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) GetLatitude() float64 { return t.Latitude }
func (t *InlineQueryResultVenue) SetLatitude(v float64) { t.Latitude = v }
func (t *InlineQueryResultVenue) WithLatitude(v float64) *InlineQueryResultVenue { t.Latitude = v; return t }
func (t *InlineQueryResultVenue) GetLongitude() float64 { return t.Longitude }
func (t *InlineQueryResultVenue) SetLongitude(v float64) { t.Longitude = v }
func (t *InlineQueryResultVenue) WithLongitude(v float64) *InlineQueryResultVenue { t.Longitude = v; return t }
func (t *InlineQueryResultVenue) GetTitle() string { return t.Title }
func (t *InlineQueryResultVenue) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultVenue) WithTitle(v string) *InlineQueryResultVenue { t.Title = v; return t }
func (t *InlineQueryResultVenue) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultVenue { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) GetAddress() string { return t.Address }
func (t *InlineQueryResultVenue) SetAddress(v string) { t.Address = v }
func (t *InlineQueryResultVenue) WithAddress(v string) *InlineQueryResultVenue { t.Address = v; return t }
func (t *InlineQueryResultVenue) SetFmtAddress(f string, a ...interface{}) { t.SetAddress(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) WithFmtAddress(f string, a ...interface{}) *InlineQueryResultVenue { return t.WithAddress(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) GetFoursquareID() string { if t.FoursquareID == nil { return "" } else { return *t.FoursquareID } }
func (t *InlineQueryResultVenue) SetFoursquareID(v string) { t.FoursquareID = &v }
func (t *InlineQueryResultVenue) WithFoursquareID(v string) *InlineQueryResultVenue { t.FoursquareID = &v; return t }
func (t *InlineQueryResultVenue) SetNilFoursquareID() { t.FoursquareID = nil }
func (t *InlineQueryResultVenue) WithNilFoursquareID() *InlineQueryResultVenue { t.FoursquareID = nil; return t }
func (t *InlineQueryResultVenue) SetFmtFoursquareID(f string, a ...interface{}) { t.SetFoursquareID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) WithFmtFoursquareID(f string, a ...interface{}) *InlineQueryResultVenue { return t.WithFoursquareID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultVenue) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultVenue) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultVenue { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultVenue) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultVenue) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultVenue) WithInputMessageContent(v InputMessageContent) *InlineQueryResultVenue { t.InputMessageContent = v; return t }
func (t *InlineQueryResultVenue) GetThumbURL() string { if t.ThumbURL == nil { return "" } else { return *t.ThumbURL } }
func (t *InlineQueryResultVenue) SetThumbURL(v string) { t.ThumbURL = &v }
func (t *InlineQueryResultVenue) WithThumbURL(v string) *InlineQueryResultVenue { t.ThumbURL = &v; return t }
func (t *InlineQueryResultVenue) SetNilThumbURL() { t.ThumbURL = nil }
func (t *InlineQueryResultVenue) WithNilThumbURL() *InlineQueryResultVenue { t.ThumbURL = nil; return t }
func (t *InlineQueryResultVenue) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultVenue { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultVenue) GetThumbWidth() int { if t.ThumbWidth == nil { return 0 } else { return *t.ThumbWidth } }
func (t *InlineQueryResultVenue) SetThumbWidth(v int) { t.ThumbWidth = &v }
func (t *InlineQueryResultVenue) WithThumbWidth(v int) *InlineQueryResultVenue { t.ThumbWidth = &v; return t }
func (t *InlineQueryResultVenue) SetNilThumbWidth() { t.ThumbWidth = nil }
func (t *InlineQueryResultVenue) WithNilThumbWidth() *InlineQueryResultVenue { t.ThumbWidth = nil; return t }
func (t *InlineQueryResultVenue) GetThumbHeight() int { if t.ThumbHeight == nil { return 0 } else { return *t.ThumbHeight } }
func (t *InlineQueryResultVenue) SetThumbHeight(v int) { t.ThumbHeight = &v }
func (t *InlineQueryResultVenue) WithThumbHeight(v int) *InlineQueryResultVenue { t.ThumbHeight = &v; return t }
func (t *InlineQueryResultVenue) SetNilThumbHeight() { t.ThumbHeight = nil }
func (t *InlineQueryResultVenue) WithNilThumbHeight() *InlineQueryResultVenue { t.ThumbHeight = nil; return t }

func (t *InlineQueryResultContact) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultContact) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultContact) WithType(v InlineQueryResultType) *InlineQueryResultContact { t.Type = v; return t }
func (t *InlineQueryResultContact) GetID() string { return t.ID }
func (t *InlineQueryResultContact) SetID(v string) { t.ID = v }
func (t *InlineQueryResultContact) WithID(v string) *InlineQueryResultContact { t.ID = v; return t }
func (t *InlineQueryResultContact) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) WithFmtID(f string, a ...interface{}) *InlineQueryResultContact { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) GetPhoneNumber() string { return t.PhoneNumber }
func (t *InlineQueryResultContact) SetPhoneNumber(v string) { t.PhoneNumber = v }
func (t *InlineQueryResultContact) WithPhoneNumber(v string) *InlineQueryResultContact { t.PhoneNumber = v; return t }
func (t *InlineQueryResultContact) SetFmtPhoneNumber(f string, a ...interface{}) { t.SetPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) WithFmtPhoneNumber(f string, a ...interface{}) *InlineQueryResultContact { return t.WithPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) GetFirstName() string { return t.FirstName }
func (t *InlineQueryResultContact) SetFirstName(v string) { t.FirstName = v }
func (t *InlineQueryResultContact) WithFirstName(v string) *InlineQueryResultContact { t.FirstName = v; return t }
func (t *InlineQueryResultContact) SetFmtFirstName(f string, a ...interface{}) { t.SetFirstName(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) WithFmtFirstName(f string, a ...interface{}) *InlineQueryResultContact { return t.WithFirstName(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) GetLastName() string { if t.LastName == nil { return "" } else { return *t.LastName } }
func (t *InlineQueryResultContact) SetLastName(v string) { t.LastName = &v }
func (t *InlineQueryResultContact) WithLastName(v string) *InlineQueryResultContact { t.LastName = &v; return t }
func (t *InlineQueryResultContact) SetNilLastName() { t.LastName = nil }
func (t *InlineQueryResultContact) WithNilLastName() *InlineQueryResultContact { t.LastName = nil; return t }
func (t *InlineQueryResultContact) SetFmtLastName(f string, a ...interface{}) { t.SetLastName(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) WithFmtLastName(f string, a ...interface{}) *InlineQueryResultContact { return t.WithLastName(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultContact) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultContact) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultContact { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultContact) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultContact) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultContact) WithInputMessageContent(v InputMessageContent) *InlineQueryResultContact { t.InputMessageContent = v; return t }
func (t *InlineQueryResultContact) GetThumbURL() string { if t.ThumbURL == nil { return "" } else { return *t.ThumbURL } }
func (t *InlineQueryResultContact) SetThumbURL(v string) { t.ThumbURL = &v }
func (t *InlineQueryResultContact) WithThumbURL(v string) *InlineQueryResultContact { t.ThumbURL = &v; return t }
func (t *InlineQueryResultContact) SetNilThumbURL() { t.ThumbURL = nil }
func (t *InlineQueryResultContact) WithNilThumbURL() *InlineQueryResultContact { t.ThumbURL = nil; return t }
func (t *InlineQueryResultContact) SetFmtThumbURL(f string, a ...interface{}) { t.SetThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) WithFmtThumbURL(f string, a ...interface{}) *InlineQueryResultContact { return t.WithThumbURL(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultContact) GetThumbWidth() int { if t.ThumbWidth == nil { return 0 } else { return *t.ThumbWidth } }
func (t *InlineQueryResultContact) SetThumbWidth(v int) { t.ThumbWidth = &v }
func (t *InlineQueryResultContact) WithThumbWidth(v int) *InlineQueryResultContact { t.ThumbWidth = &v; return t }
func (t *InlineQueryResultContact) SetNilThumbWidth() { t.ThumbWidth = nil }
func (t *InlineQueryResultContact) WithNilThumbWidth() *InlineQueryResultContact { t.ThumbWidth = nil; return t }
func (t *InlineQueryResultContact) GetThumbHeight() int { if t.ThumbHeight == nil { return 0 } else { return *t.ThumbHeight } }
func (t *InlineQueryResultContact) SetThumbHeight(v int) { t.ThumbHeight = &v }
func (t *InlineQueryResultContact) WithThumbHeight(v int) *InlineQueryResultContact { t.ThumbHeight = &v; return t }
func (t *InlineQueryResultContact) SetNilThumbHeight() { t.ThumbHeight = nil }
func (t *InlineQueryResultContact) WithNilThumbHeight() *InlineQueryResultContact { t.ThumbHeight = nil; return t }

func (t *InlineQueryResultGame) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultGame) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultGame) WithType(v InlineQueryResultType) *InlineQueryResultGame { t.Type = v; return t }
func (t *InlineQueryResultGame) GetID() string { return t.ID }
func (t *InlineQueryResultGame) SetID(v string) { t.ID = v }
func (t *InlineQueryResultGame) WithID(v string) *InlineQueryResultGame { t.ID = v; return t }
func (t *InlineQueryResultGame) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGame) WithFmtID(f string, a ...interface{}) *InlineQueryResultGame { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGame) GetGameShortName() string { return t.GameShortName }
func (t *InlineQueryResultGame) SetGameShortName(v string) { t.GameShortName = v }
func (t *InlineQueryResultGame) WithGameShortName(v string) *InlineQueryResultGame { t.GameShortName = v; return t }
func (t *InlineQueryResultGame) SetFmtGameShortName(f string, a ...interface{}) { t.SetGameShortName(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGame) WithFmtGameShortName(f string, a ...interface{}) *InlineQueryResultGame { return t.WithGameShortName(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultGame) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultGame) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultGame) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultGame { t.ReplyMarkup = v; return t }

func (t *InlineQueryResultCachedPhoto) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedPhoto) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedPhoto) WithType(v InlineQueryResultType) *InlineQueryResultCachedPhoto { t.Type = v; return t }
func (t *InlineQueryResultCachedPhoto) GetID() string { return t.ID }
func (t *InlineQueryResultCachedPhoto) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedPhoto) WithID(v string) *InlineQueryResultCachedPhoto { t.ID = v; return t }
func (t *InlineQueryResultCachedPhoto) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedPhoto { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) GetPhotoFileID() string { return t.PhotoFileID }
func (t *InlineQueryResultCachedPhoto) SetPhotoFileID(v string) { t.PhotoFileID = v }
func (t *InlineQueryResultCachedPhoto) WithPhotoFileID(v string) *InlineQueryResultCachedPhoto { t.PhotoFileID = v; return t }
func (t *InlineQueryResultCachedPhoto) SetFmtPhotoFileID(f string, a ...interface{}) { t.SetPhotoFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) WithFmtPhotoFileID(f string, a ...interface{}) *InlineQueryResultCachedPhoto { return t.WithPhotoFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *InlineQueryResultCachedPhoto) SetTitle(v string) { t.Title = &v }
func (t *InlineQueryResultCachedPhoto) WithTitle(v string) *InlineQueryResultCachedPhoto { t.Title = &v; return t }
func (t *InlineQueryResultCachedPhoto) SetNilTitle() { t.Title = nil }
func (t *InlineQueryResultCachedPhoto) WithNilTitle() *InlineQueryResultCachedPhoto { t.Title = nil; return t }
func (t *InlineQueryResultCachedPhoto) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultCachedPhoto { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultCachedPhoto) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultCachedPhoto) WithDescription(v string) *InlineQueryResultCachedPhoto { t.Description = &v; return t }
func (t *InlineQueryResultCachedPhoto) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultCachedPhoto) WithNilDescription() *InlineQueryResultCachedPhoto { t.Description = nil; return t }
func (t *InlineQueryResultCachedPhoto) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultCachedPhoto { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedPhoto) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedPhoto) WithCaption(v string) *InlineQueryResultCachedPhoto { t.Caption = &v; return t }
func (t *InlineQueryResultCachedPhoto) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedPhoto) WithNilCaption() *InlineQueryResultCachedPhoto { t.Caption = nil; return t }
func (t *InlineQueryResultCachedPhoto) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedPhoto { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedPhoto) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedPhoto) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedPhoto) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedPhoto { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedPhoto) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedPhoto) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedPhoto) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedPhoto { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedGIF) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedGIF) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedGIF) WithType(v InlineQueryResultType) *InlineQueryResultCachedGIF { t.Type = v; return t }
func (t *InlineQueryResultCachedGIF) GetID() string { return t.ID }
func (t *InlineQueryResultCachedGIF) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedGIF) WithID(v string) *InlineQueryResultCachedGIF { t.ID = v; return t }
func (t *InlineQueryResultCachedGIF) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedGIF { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) GetGIFFileID() string { return t.GIFFileID }
func (t *InlineQueryResultCachedGIF) SetGIFFileID(v string) { t.GIFFileID = v }
func (t *InlineQueryResultCachedGIF) WithGIFFileID(v string) *InlineQueryResultCachedGIF { t.GIFFileID = v; return t }
func (t *InlineQueryResultCachedGIF) SetFmtGIFFileID(f string, a ...interface{}) { t.SetGIFFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) WithFmtGIFFileID(f string, a ...interface{}) *InlineQueryResultCachedGIF { return t.WithGIFFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *InlineQueryResultCachedGIF) SetTitle(v string) { t.Title = &v }
func (t *InlineQueryResultCachedGIF) WithTitle(v string) *InlineQueryResultCachedGIF { t.Title = &v; return t }
func (t *InlineQueryResultCachedGIF) SetNilTitle() { t.Title = nil }
func (t *InlineQueryResultCachedGIF) WithNilTitle() *InlineQueryResultCachedGIF { t.Title = nil; return t }
func (t *InlineQueryResultCachedGIF) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultCachedGIF { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedGIF) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedGIF) WithCaption(v string) *InlineQueryResultCachedGIF { t.Caption = &v; return t }
func (t *InlineQueryResultCachedGIF) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedGIF) WithNilCaption() *InlineQueryResultCachedGIF { t.Caption = nil; return t }
func (t *InlineQueryResultCachedGIF) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedGIF { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedGIF) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedGIF) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedGIF) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedGIF { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedGIF) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedGIF) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedGIF) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedGIF { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedMPEG4GIF) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedMPEG4GIF) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedMPEG4GIF) WithType(v InlineQueryResultType) *InlineQueryResultCachedMPEG4GIF { t.Type = v; return t }
func (t *InlineQueryResultCachedMPEG4GIF) GetID() string { return t.ID }
func (t *InlineQueryResultCachedMPEG4GIF) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedMPEG4GIF) WithID(v string) *InlineQueryResultCachedMPEG4GIF { t.ID = v; return t }
func (t *InlineQueryResultCachedMPEG4GIF) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedMPEG4GIF { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) GetMPEG4FileID() string { return t.MPEG4FileID }
func (t *InlineQueryResultCachedMPEG4GIF) SetMPEG4FileID(v string) { t.MPEG4FileID = v }
func (t *InlineQueryResultCachedMPEG4GIF) WithMPEG4FileID(v string) *InlineQueryResultCachedMPEG4GIF { t.MPEG4FileID = v; return t }
func (t *InlineQueryResultCachedMPEG4GIF) SetFmtMPEG4FileID(f string, a ...interface{}) { t.SetMPEG4FileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) WithFmtMPEG4FileID(f string, a ...interface{}) *InlineQueryResultCachedMPEG4GIF { return t.WithMPEG4FileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) GetTitle() string { if t.Title == nil { return "" } else { return *t.Title } }
func (t *InlineQueryResultCachedMPEG4GIF) SetTitle(v string) { t.Title = &v }
func (t *InlineQueryResultCachedMPEG4GIF) WithTitle(v string) *InlineQueryResultCachedMPEG4GIF { t.Title = &v; return t }
func (t *InlineQueryResultCachedMPEG4GIF) SetNilTitle() { t.Title = nil }
func (t *InlineQueryResultCachedMPEG4GIF) WithNilTitle() *InlineQueryResultCachedMPEG4GIF { t.Title = nil; return t }
func (t *InlineQueryResultCachedMPEG4GIF) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultCachedMPEG4GIF { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedMPEG4GIF) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedMPEG4GIF) WithCaption(v string) *InlineQueryResultCachedMPEG4GIF { t.Caption = &v; return t }
func (t *InlineQueryResultCachedMPEG4GIF) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedMPEG4GIF) WithNilCaption() *InlineQueryResultCachedMPEG4GIF { t.Caption = nil; return t }
func (t *InlineQueryResultCachedMPEG4GIF) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedMPEG4GIF { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedMPEG4GIF) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedMPEG4GIF) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedMPEG4GIF) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedMPEG4GIF { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedMPEG4GIF) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedMPEG4GIF) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedMPEG4GIF) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedMPEG4GIF { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedSticker) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedSticker) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedSticker) WithType(v InlineQueryResultType) *InlineQueryResultCachedSticker { t.Type = v; return t }
func (t *InlineQueryResultCachedSticker) GetID() string { return t.ID }
func (t *InlineQueryResultCachedSticker) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedSticker) WithID(v string) *InlineQueryResultCachedSticker { t.ID = v; return t }
func (t *InlineQueryResultCachedSticker) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedSticker) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedSticker { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedSticker) GetStickerFileID() string { return t.StickerFileID }
func (t *InlineQueryResultCachedSticker) SetStickerFileID(v string) { t.StickerFileID = v }
func (t *InlineQueryResultCachedSticker) WithStickerFileID(v string) *InlineQueryResultCachedSticker { t.StickerFileID = v; return t }
func (t *InlineQueryResultCachedSticker) SetFmtStickerFileID(f string, a ...interface{}) { t.SetStickerFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedSticker) WithFmtStickerFileID(f string, a ...interface{}) *InlineQueryResultCachedSticker { return t.WithStickerFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedSticker) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedSticker) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedSticker) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedSticker { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedSticker) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedSticker) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedSticker) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedSticker { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedDocument) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedDocument) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedDocument) WithType(v InlineQueryResultType) *InlineQueryResultCachedDocument { t.Type = v; return t }
func (t *InlineQueryResultCachedDocument) GetID() string { return t.ID }
func (t *InlineQueryResultCachedDocument) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedDocument) WithID(v string) *InlineQueryResultCachedDocument { t.ID = v; return t }
func (t *InlineQueryResultCachedDocument) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedDocument { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) GetTitle() string { return t.Title }
func (t *InlineQueryResultCachedDocument) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultCachedDocument) WithTitle(v string) *InlineQueryResultCachedDocument { t.Title = v; return t }
func (t *InlineQueryResultCachedDocument) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultCachedDocument { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) GetDocmentFileID() string { return t.DocmentFileID }
func (t *InlineQueryResultCachedDocument) SetDocmentFileID(v string) { t.DocmentFileID = v }
func (t *InlineQueryResultCachedDocument) WithDocmentFileID(v string) *InlineQueryResultCachedDocument { t.DocmentFileID = v; return t }
func (t *InlineQueryResultCachedDocument) SetFmtDocmentFileID(f string, a ...interface{}) { t.SetDocmentFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) WithFmtDocmentFileID(f string, a ...interface{}) *InlineQueryResultCachedDocument { return t.WithDocmentFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultCachedDocument) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultCachedDocument) WithDescription(v string) *InlineQueryResultCachedDocument { t.Description = &v; return t }
func (t *InlineQueryResultCachedDocument) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultCachedDocument) WithNilDescription() *InlineQueryResultCachedDocument { t.Description = nil; return t }
func (t *InlineQueryResultCachedDocument) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultCachedDocument { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedDocument) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedDocument) WithCaption(v string) *InlineQueryResultCachedDocument { t.Caption = &v; return t }
func (t *InlineQueryResultCachedDocument) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedDocument) WithNilCaption() *InlineQueryResultCachedDocument { t.Caption = nil; return t }
func (t *InlineQueryResultCachedDocument) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedDocument { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedDocument) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedDocument) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedDocument) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedDocument { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedDocument) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedDocument) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedDocument) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedDocument { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedVideo) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedVideo) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedVideo) WithType(v InlineQueryResultType) *InlineQueryResultCachedVideo { t.Type = v; return t }
func (t *InlineQueryResultCachedVideo) GetID() string { return t.ID }
func (t *InlineQueryResultCachedVideo) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedVideo) WithID(v string) *InlineQueryResultCachedVideo { t.ID = v; return t }
func (t *InlineQueryResultCachedVideo) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedVideo { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) GetVideoFileID() string { return t.VideoFileID }
func (t *InlineQueryResultCachedVideo) SetVideoFileID(v string) { t.VideoFileID = v }
func (t *InlineQueryResultCachedVideo) WithVideoFileID(v string) *InlineQueryResultCachedVideo { t.VideoFileID = v; return t }
func (t *InlineQueryResultCachedVideo) SetFmtVideoFileID(f string, a ...interface{}) { t.SetVideoFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) WithFmtVideoFileID(f string, a ...interface{}) *InlineQueryResultCachedVideo { return t.WithVideoFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) GetTitle() string { return t.Title }
func (t *InlineQueryResultCachedVideo) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultCachedVideo) WithTitle(v string) *InlineQueryResultCachedVideo { t.Title = v; return t }
func (t *InlineQueryResultCachedVideo) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultCachedVideo { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) GetDescription() string { if t.Description == nil { return "" } else { return *t.Description } }
func (t *InlineQueryResultCachedVideo) SetDescription(v string) { t.Description = &v }
func (t *InlineQueryResultCachedVideo) WithDescription(v string) *InlineQueryResultCachedVideo { t.Description = &v; return t }
func (t *InlineQueryResultCachedVideo) SetNilDescription() { t.Description = nil }
func (t *InlineQueryResultCachedVideo) WithNilDescription() *InlineQueryResultCachedVideo { t.Description = nil; return t }
func (t *InlineQueryResultCachedVideo) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) WithFmtDescription(f string, a ...interface{}) *InlineQueryResultCachedVideo { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedVideo) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedVideo) WithCaption(v string) *InlineQueryResultCachedVideo { t.Caption = &v; return t }
func (t *InlineQueryResultCachedVideo) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedVideo) WithNilCaption() *InlineQueryResultCachedVideo { t.Caption = nil; return t }
func (t *InlineQueryResultCachedVideo) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedVideo { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVideo) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedVideo) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedVideo) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedVideo { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedVideo) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedVideo) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedVideo) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedVideo { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedVoice) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedVoice) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedVoice) WithType(v InlineQueryResultType) *InlineQueryResultCachedVoice { t.Type = v; return t }
func (t *InlineQueryResultCachedVoice) GetID() string { return t.ID }
func (t *InlineQueryResultCachedVoice) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedVoice) WithID(v string) *InlineQueryResultCachedVoice { t.ID = v; return t }
func (t *InlineQueryResultCachedVoice) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedVoice { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) GetVoiceFileID() string { return t.VoiceFileID }
func (t *InlineQueryResultCachedVoice) SetVoiceFileID(v string) { t.VoiceFileID = v }
func (t *InlineQueryResultCachedVoice) WithVoiceFileID(v string) *InlineQueryResultCachedVoice { t.VoiceFileID = v; return t }
func (t *InlineQueryResultCachedVoice) SetFmtVoiceFileID(f string, a ...interface{}) { t.SetVoiceFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) WithFmtVoiceFileID(f string, a ...interface{}) *InlineQueryResultCachedVoice { return t.WithVoiceFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) GetTitle() string { return t.Title }
func (t *InlineQueryResultCachedVoice) SetTitle(v string) { t.Title = v }
func (t *InlineQueryResultCachedVoice) WithTitle(v string) *InlineQueryResultCachedVoice { t.Title = v; return t }
func (t *InlineQueryResultCachedVoice) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) WithFmtTitle(f string, a ...interface{}) *InlineQueryResultCachedVoice { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedVoice) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedVoice) WithCaption(v string) *InlineQueryResultCachedVoice { t.Caption = &v; return t }
func (t *InlineQueryResultCachedVoice) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedVoice) WithNilCaption() *InlineQueryResultCachedVoice { t.Caption = nil; return t }
func (t *InlineQueryResultCachedVoice) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedVoice { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedVoice) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedVoice) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedVoice) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedVoice { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedVoice) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedVoice) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedVoice) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedVoice { t.InputMessageContent = v; return t }

func (t *InlineQueryResultCachedAudio) GetType() InlineQueryResultType { return t.Type }
func (t *InlineQueryResultCachedAudio) SetType(v InlineQueryResultType) { t.Type = v }
func (t *InlineQueryResultCachedAudio) WithType(v InlineQueryResultType) *InlineQueryResultCachedAudio { t.Type = v; return t }
func (t *InlineQueryResultCachedAudio) GetID() string { return t.ID }
func (t *InlineQueryResultCachedAudio) SetID(v string) { t.ID = v }
func (t *InlineQueryResultCachedAudio) WithID(v string) *InlineQueryResultCachedAudio { t.ID = v; return t }
func (t *InlineQueryResultCachedAudio) SetFmtID(f string, a ...interface{}) { t.SetID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedAudio) WithFmtID(f string, a ...interface{}) *InlineQueryResultCachedAudio { return t.WithID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedAudio) GetAudioFileID() string { return t.AudioFileID }
func (t *InlineQueryResultCachedAudio) SetAudioFileID(v string) { t.AudioFileID = v }
func (t *InlineQueryResultCachedAudio) WithAudioFileID(v string) *InlineQueryResultCachedAudio { t.AudioFileID = v; return t }
func (t *InlineQueryResultCachedAudio) SetFmtAudioFileID(f string, a ...interface{}) { t.SetAudioFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedAudio) WithFmtAudioFileID(f string, a ...interface{}) *InlineQueryResultCachedAudio { return t.WithAudioFileID(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedAudio) GetCaption() string { if t.Caption == nil { return "" } else { return *t.Caption } }
func (t *InlineQueryResultCachedAudio) SetCaption(v string) { t.Caption = &v }
func (t *InlineQueryResultCachedAudio) WithCaption(v string) *InlineQueryResultCachedAudio { t.Caption = &v; return t }
func (t *InlineQueryResultCachedAudio) SetNilCaption() { t.Caption = nil }
func (t *InlineQueryResultCachedAudio) WithNilCaption() *InlineQueryResultCachedAudio { t.Caption = nil; return t }
func (t *InlineQueryResultCachedAudio) SetFmtCaption(f string, a ...interface{}) { t.SetCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedAudio) WithFmtCaption(f string, a ...interface{}) *InlineQueryResultCachedAudio { return t.WithCaption(fmt.Sprintf(f, a...)) }
func (t *InlineQueryResultCachedAudio) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *InlineQueryResultCachedAudio) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *InlineQueryResultCachedAudio) WithReplyMarkup(v *InlineKeyboardMarkup) *InlineQueryResultCachedAudio { t.ReplyMarkup = v; return t }
func (t *InlineQueryResultCachedAudio) GetInputMessageContent() InputMessageContent { return t.InputMessageContent }
func (t *InlineQueryResultCachedAudio) SetInputMessageContent(v InputMessageContent) { t.InputMessageContent = v }
func (t *InlineQueryResultCachedAudio) WithInputMessageContent(v InputMessageContent) *InlineQueryResultCachedAudio { t.InputMessageContent = v; return t }

func (t *InputTextMessageContent) GetMessageText() string { return t.MessageText }
func (t *InputTextMessageContent) SetMessageText(v string) { t.MessageText = v }
func (t *InputTextMessageContent) WithMessageText(v string) *InputTextMessageContent { t.MessageText = v; return t }
func (t *InputTextMessageContent) SetFmtMessageText(f string, a ...interface{}) { t.SetMessageText(fmt.Sprintf(f, a...)) }
func (t *InputTextMessageContent) WithFmtMessageText(f string, a ...interface{}) *InputTextMessageContent { return t.WithMessageText(fmt.Sprintf(f, a...)) }
func (t *InputTextMessageContent) GetParseMode() string { if t.ParseMode == nil { return "" } else { return *t.ParseMode } }
func (t *InputTextMessageContent) SetParseMode(v string) { t.ParseMode = &v }
func (t *InputTextMessageContent) WithParseMode(v string) *InputTextMessageContent { t.ParseMode = &v; return t }
func (t *InputTextMessageContent) SetNilParseMode() { t.ParseMode = nil }
func (t *InputTextMessageContent) WithNilParseMode() *InputTextMessageContent { t.ParseMode = nil; return t }
func (t *InputTextMessageContent) SetFmtParseMode(f string, a ...interface{}) { t.SetParseMode(fmt.Sprintf(f, a...)) }
func (t *InputTextMessageContent) WithFmtParseMode(f string, a ...interface{}) *InputTextMessageContent { return t.WithParseMode(fmt.Sprintf(f, a...)) }
func (t *InputTextMessageContent) GetDisableWebPagePreview() bool { if t.DisableWebPagePreview == nil { return false } else { return *t.DisableWebPagePreview } }
func (t *InputTextMessageContent) SetDisableWebPagePreview(v bool) { t.DisableWebPagePreview = &v }
func (t *InputTextMessageContent) WithDisableWebPagePreview(v bool) *InputTextMessageContent { t.DisableWebPagePreview = &v; return t }
func (t *InputTextMessageContent) SetNilDisableWebPagePreview() { t.DisableWebPagePreview = nil }
func (t *InputTextMessageContent) WithNilDisableWebPagePreview() *InputTextMessageContent { t.DisableWebPagePreview = nil; return t }

func (t *InputLocationMessageContent) GetLatitude() float64 { return t.Latitude }
func (t *InputLocationMessageContent) SetLatitude(v float64) { t.Latitude = v }
func (t *InputLocationMessageContent) WithLatitude(v float64) *InputLocationMessageContent { t.Latitude = v; return t }
func (t *InputLocationMessageContent) GetLongitude() float64 { return t.Longitude }
func (t *InputLocationMessageContent) SetLongitude(v float64) { t.Longitude = v }
func (t *InputLocationMessageContent) WithLongitude(v float64) *InputLocationMessageContent { t.Longitude = v; return t }

func (t *InputVenueMessageContent) GetLatitude() float64 { return t.Latitude }
func (t *InputVenueMessageContent) SetLatitude(v float64) { t.Latitude = v }
func (t *InputVenueMessageContent) WithLatitude(v float64) *InputVenueMessageContent { t.Latitude = v; return t }
func (t *InputVenueMessageContent) GetLongitude() float64 { return t.Longitude }
func (t *InputVenueMessageContent) SetLongitude(v float64) { t.Longitude = v }
func (t *InputVenueMessageContent) WithLongitude(v float64) *InputVenueMessageContent { t.Longitude = v; return t }
func (t *InputVenueMessageContent) GetTitle() string { return t.Title }
func (t *InputVenueMessageContent) SetTitle(v string) { t.Title = v }
func (t *InputVenueMessageContent) WithTitle(v string) *InputVenueMessageContent { t.Title = v; return t }
func (t *InputVenueMessageContent) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *InputVenueMessageContent) WithFmtTitle(f string, a ...interface{}) *InputVenueMessageContent { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *InputVenueMessageContent) GetAddress() string { return t.Address }
func (t *InputVenueMessageContent) SetAddress(v string) { t.Address = v }
func (t *InputVenueMessageContent) WithAddress(v string) *InputVenueMessageContent { t.Address = v; return t }
func (t *InputVenueMessageContent) SetFmtAddress(f string, a ...interface{}) { t.SetAddress(fmt.Sprintf(f, a...)) }
func (t *InputVenueMessageContent) WithFmtAddress(f string, a ...interface{}) *InputVenueMessageContent { return t.WithAddress(fmt.Sprintf(f, a...)) }
func (t *InputVenueMessageContent) GetFoursquareID() string { if t.FoursquareID == nil { return "" } else { return *t.FoursquareID } }
func (t *InputVenueMessageContent) SetFoursquareID(v string) { t.FoursquareID = &v }
func (t *InputVenueMessageContent) WithFoursquareID(v string) *InputVenueMessageContent { t.FoursquareID = &v; return t }
func (t *InputVenueMessageContent) SetNilFoursquareID() { t.FoursquareID = nil }
func (t *InputVenueMessageContent) WithNilFoursquareID() *InputVenueMessageContent { t.FoursquareID = nil; return t }
func (t *InputVenueMessageContent) SetFmtFoursquareID(f string, a ...interface{}) { t.SetFoursquareID(fmt.Sprintf(f, a...)) }
func (t *InputVenueMessageContent) WithFmtFoursquareID(f string, a ...interface{}) *InputVenueMessageContent { return t.WithFoursquareID(fmt.Sprintf(f, a...)) }

func (t *InputContactMessageContent) GetPhoneNumber() string { return t.PhoneNumber }
func (t *InputContactMessageContent) SetPhoneNumber(v string) { t.PhoneNumber = v }
func (t *InputContactMessageContent) WithPhoneNumber(v string) *InputContactMessageContent { t.PhoneNumber = v; return t }
func (t *InputContactMessageContent) SetFmtPhoneNumber(f string, a ...interface{}) { t.SetPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *InputContactMessageContent) WithFmtPhoneNumber(f string, a ...interface{}) *InputContactMessageContent { return t.WithPhoneNumber(fmt.Sprintf(f, a...)) }
func (t *InputContactMessageContent) GetFirstName() string { return t.FirstName }
func (t *InputContactMessageContent) SetFirstName(v string) { t.FirstName = v }
func (t *InputContactMessageContent) WithFirstName(v string) *InputContactMessageContent { t.FirstName = v; return t }
func (t *InputContactMessageContent) SetFmtFirstName(f string, a ...interface{}) { t.SetFirstName(fmt.Sprintf(f, a...)) }
func (t *InputContactMessageContent) WithFmtFirstName(f string, a ...interface{}) *InputContactMessageContent { return t.WithFirstName(fmt.Sprintf(f, a...)) }
func (t *InputContactMessageContent) GetLastName() string { if t.LastName == nil { return "" } else { return *t.LastName } }
func (t *InputContactMessageContent) SetLastName(v string) { t.LastName = &v }
func (t *InputContactMessageContent) WithLastName(v string) *InputContactMessageContent { t.LastName = &v; return t }
func (t *InputContactMessageContent) SetNilLastName() { t.LastName = nil }
func (t *InputContactMessageContent) WithNilLastName() *InputContactMessageContent { t.LastName = nil; return t }
func (t *InputContactMessageContent) SetFmtLastName(f string, a ...interface{}) { t.SetLastName(fmt.Sprintf(f, a...)) }
func (t *InputContactMessageContent) WithFmtLastName(f string, a ...interface{}) *InputContactMessageContent { return t.WithLastName(fmt.Sprintf(f, a...)) }

func (t *ChosenInlineResult) GetResultID() string { return t.ResultID }
func (t *ChosenInlineResult) SetResultID(v string) { t.ResultID = v }
func (t *ChosenInlineResult) WithResultID(v string) *ChosenInlineResult { t.ResultID = v; return t }
func (t *ChosenInlineResult) SetFmtResultID(f string, a ...interface{}) { t.SetResultID(fmt.Sprintf(f, a...)) }
func (t *ChosenInlineResult) WithFmtResultID(f string, a ...interface{}) *ChosenInlineResult { return t.WithResultID(fmt.Sprintf(f, a...)) }
func (t *ChosenInlineResult) GetFrom() *User { return t.From }
func (t *ChosenInlineResult) SetFrom(v *User) { t.From = v }
func (t *ChosenInlineResult) WithFrom(v *User) *ChosenInlineResult { t.From = v; return t }
func (t *ChosenInlineResult) GetLocation() *Location { return t.Location }
func (t *ChosenInlineResult) SetLocation(v *Location) { t.Location = v }
func (t *ChosenInlineResult) WithLocation(v *Location) *ChosenInlineResult { t.Location = v; return t }
func (t *ChosenInlineResult) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *ChosenInlineResult) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *ChosenInlineResult) WithInlineMessageID(v string) *ChosenInlineResult { t.InlineMessageID = &v; return t }
func (t *ChosenInlineResult) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *ChosenInlineResult) WithNilInlineMessageID() *ChosenInlineResult { t.InlineMessageID = nil; return t }
func (t *ChosenInlineResult) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *ChosenInlineResult) WithFmtInlineMessageID(f string, a ...interface{}) *ChosenInlineResult { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *ChosenInlineResult) GetQuery() string { return t.Query }
func (t *ChosenInlineResult) SetQuery(v string) { t.Query = v }
func (t *ChosenInlineResult) WithQuery(v string) *ChosenInlineResult { t.Query = v; return t }
func (t *ChosenInlineResult) SetFmtQuery(f string, a ...interface{}) { t.SetQuery(fmt.Sprintf(f, a...)) }
func (t *ChosenInlineResult) WithFmtQuery(f string, a ...interface{}) *ChosenInlineResult { return t.WithQuery(fmt.Sprintf(f, a...)) }

func (t *SendGame) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SendGame) SetMethod(v MethodName) { t.Method = &v }
func (t *SendGame) WithMethod(v MethodName) *SendGame { t.Method = &v; return t }
func (t *SendGame) SetNilMethod() { t.Method = nil }
func (t *SendGame) WithNilMethod() *SendGame { t.Method = nil; return t }
func (t *SendGame) GetChatID() ChatID { return t.ChatID }
func (t *SendGame) SetChatID(v ChatID) { t.ChatID = v }
func (t *SendGame) WithChatID(v ChatID) *SendGame { t.ChatID = v; return t }
func (t *SendGame) GetGameShortName() string { return t.GameShortName }
func (t *SendGame) SetGameShortName(v string) { t.GameShortName = v }
func (t *SendGame) WithGameShortName(v string) *SendGame { t.GameShortName = v; return t }
func (t *SendGame) SetFmtGameShortName(f string, a ...interface{}) { t.SetGameShortName(fmt.Sprintf(f, a...)) }
func (t *SendGame) WithFmtGameShortName(f string, a ...interface{}) *SendGame { return t.WithGameShortName(fmt.Sprintf(f, a...)) }
func (t *SendGame) GetDisableNotification() bool { if t.DisableNotification == nil { return false } else { return *t.DisableNotification } }
func (t *SendGame) SetDisableNotification(v bool) { t.DisableNotification = &v }
func (t *SendGame) WithDisableNotification(v bool) *SendGame { t.DisableNotification = &v; return t }
func (t *SendGame) SetNilDisableNotification() { t.DisableNotification = nil }
func (t *SendGame) WithNilDisableNotification() *SendGame { t.DisableNotification = nil; return t }
func (t *SendGame) GetReplyToMessageID() int { if t.ReplyToMessageID == nil { return 0 } else { return *t.ReplyToMessageID } }
func (t *SendGame) SetReplyToMessageID(v int) { t.ReplyToMessageID = &v }
func (t *SendGame) WithReplyToMessageID(v int) *SendGame { t.ReplyToMessageID = &v; return t }
func (t *SendGame) SetNilReplyToMessageID() { t.ReplyToMessageID = nil }
func (t *SendGame) WithNilReplyToMessageID() *SendGame { t.ReplyToMessageID = nil; return t }
func (t *SendGame) GetReplyMarkup() *InlineKeyboardMarkup { return t.ReplyMarkup }
func (t *SendGame) SetReplyMarkup(v *InlineKeyboardMarkup) { t.ReplyMarkup = v }
func (t *SendGame) WithReplyMarkup(v *InlineKeyboardMarkup) *SendGame { t.ReplyMarkup = v; return t }

func (t *Game) GetTitle() string { return t.Title }
func (t *Game) SetTitle(v string) { t.Title = v }
func (t *Game) WithTitle(v string) *Game { t.Title = v; return t }
func (t *Game) SetFmtTitle(f string, a ...interface{}) { t.SetTitle(fmt.Sprintf(f, a...)) }
func (t *Game) WithFmtTitle(f string, a ...interface{}) *Game { return t.WithTitle(fmt.Sprintf(f, a...)) }
func (t *Game) GetDescription() string { return t.Description }
func (t *Game) SetDescription(v string) { t.Description = v }
func (t *Game) WithDescription(v string) *Game { t.Description = v; return t }
func (t *Game) SetFmtDescription(f string, a ...interface{}) { t.SetDescription(fmt.Sprintf(f, a...)) }
func (t *Game) WithFmtDescription(f string, a ...interface{}) *Game { return t.WithDescription(fmt.Sprintf(f, a...)) }
func (t *Game) GetPhoto() []*PhotoSize { return t.Photo }
func (t *Game) SetPhoto(v []*PhotoSize) { t.Photo = v }
func (t *Game) WithPhoto(v []*PhotoSize) *Game { t.Photo = v; return t }
func (t *Game) GetText() string { if t.Text == nil { return "" } else { return *t.Text } }
func (t *Game) SetText(v string) { t.Text = &v }
func (t *Game) WithText(v string) *Game { t.Text = &v; return t }
func (t *Game) SetNilText() { t.Text = nil }
func (t *Game) WithNilText() *Game { t.Text = nil; return t }
func (t *Game) SetFmtText(f string, a ...interface{}) { t.SetText(fmt.Sprintf(f, a...)) }
func (t *Game) WithFmtText(f string, a ...interface{}) *Game { return t.WithText(fmt.Sprintf(f, a...)) }
func (t *Game) GetTextEntities() []*MessageEntity { if t.TextEntities == nil { return nil } else { return *t.TextEntities } }
func (t *Game) SetTextEntities(v []*MessageEntity) { if v == nil { t.TextEntities = nil } else { t.TextEntities = &v } }
func (t *Game) WithTextEntities(v []*MessageEntity) *Game { if v == nil { t.TextEntities = nil } else { t.TextEntities = &v }; return t }
func (t *Game) GetAnimation() *Animation { return t.Animation }
func (t *Game) SetAnimation(v *Animation) { t.Animation = v }
func (t *Game) WithAnimation(v *Animation) *Game { t.Animation = v; return t }

func (t *Animation) GetFileID() string { return t.FileID }
func (t *Animation) SetFileID(v string) { t.FileID = v }
func (t *Animation) WithFileID(v string) *Animation { t.FileID = v; return t }
func (t *Animation) SetFmtFileID(f string, a ...interface{}) { t.SetFileID(fmt.Sprintf(f, a...)) }
func (t *Animation) WithFmtFileID(f string, a ...interface{}) *Animation { return t.WithFileID(fmt.Sprintf(f, a...)) }
func (t *Animation) GetThumb() *PhotoSize { return t.Thumb }
func (t *Animation) SetThumb(v *PhotoSize) { t.Thumb = v }
func (t *Animation) WithThumb(v *PhotoSize) *Animation { t.Thumb = v; return t }
func (t *Animation) GetFileName() string { if t.FileName == nil { return "" } else { return *t.FileName } }
func (t *Animation) SetFileName(v string) { t.FileName = &v }
func (t *Animation) WithFileName(v string) *Animation { t.FileName = &v; return t }
func (t *Animation) SetNilFileName() { t.FileName = nil }
func (t *Animation) WithNilFileName() *Animation { t.FileName = nil; return t }
func (t *Animation) SetFmtFileName(f string, a ...interface{}) { t.SetFileName(fmt.Sprintf(f, a...)) }
func (t *Animation) WithFmtFileName(f string, a ...interface{}) *Animation { return t.WithFileName(fmt.Sprintf(f, a...)) }
func (t *Animation) GetMimeType() string { if t.MimeType == nil { return "" } else { return *t.MimeType } }
func (t *Animation) SetMimeType(v string) { t.MimeType = &v }
func (t *Animation) WithMimeType(v string) *Animation { t.MimeType = &v; return t }
func (t *Animation) SetNilMimeType() { t.MimeType = nil }
func (t *Animation) WithNilMimeType() *Animation { t.MimeType = nil; return t }
func (t *Animation) SetFmtMimeType(f string, a ...interface{}) { t.SetMimeType(fmt.Sprintf(f, a...)) }
func (t *Animation) WithFmtMimeType(f string, a ...interface{}) *Animation { return t.WithMimeType(fmt.Sprintf(f, a...)) }
func (t *Animation) GetFileSize() int { if t.FileSize == nil { return 0 } else { return *t.FileSize } }
func (t *Animation) SetFileSize(v int) { t.FileSize = &v }
func (t *Animation) WithFileSize(v int) *Animation { t.FileSize = &v; return t }
func (t *Animation) SetNilFileSize() { t.FileSize = nil }
func (t *Animation) WithNilFileSize() *Animation { t.FileSize = nil; return t }

func (t *SetGameScore) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *SetGameScore) SetMethod(v MethodName) { t.Method = &v }
func (t *SetGameScore) WithMethod(v MethodName) *SetGameScore { t.Method = &v; return t }
func (t *SetGameScore) SetNilMethod() { t.Method = nil }
func (t *SetGameScore) WithNilMethod() *SetGameScore { t.Method = nil; return t }
func (t *SetGameScore) GetUserID() int { return t.UserID }
func (t *SetGameScore) SetUserID(v int) { t.UserID = v }
func (t *SetGameScore) WithUserID(v int) *SetGameScore { t.UserID = v; return t }
func (t *SetGameScore) GetScore() int { return t.Score }
func (t *SetGameScore) SetScore(v int) { t.Score = v }
func (t *SetGameScore) WithScore(v int) *SetGameScore { t.Score = v; return t }
func (t *SetGameScore) GetForce() bool { if t.Force == nil { return false } else { return *t.Force } }
func (t *SetGameScore) SetForce(v bool) { t.Force = &v }
func (t *SetGameScore) WithForce(v bool) *SetGameScore { t.Force = &v; return t }
func (t *SetGameScore) SetNilForce() { t.Force = nil }
func (t *SetGameScore) WithNilForce() *SetGameScore { t.Force = nil; return t }
func (t *SetGameScore) GetDisableEditMessage() bool { if t.DisableEditMessage == nil { return false } else { return *t.DisableEditMessage } }
func (t *SetGameScore) SetDisableEditMessage(v bool) { t.DisableEditMessage = &v }
func (t *SetGameScore) WithDisableEditMessage(v bool) *SetGameScore { t.DisableEditMessage = &v; return t }
func (t *SetGameScore) SetNilDisableEditMessage() { t.DisableEditMessage = nil }
func (t *SetGameScore) WithNilDisableEditMessage() *SetGameScore { t.DisableEditMessage = nil; return t }
func (t *SetGameScore) GetChatID() int64 { if t.ChatID == nil { return 0 } else { return *t.ChatID } }
func (t *SetGameScore) SetChatID(v int64) { t.ChatID = &v }
func (t *SetGameScore) WithChatID(v int64) *SetGameScore { t.ChatID = &v; return t }
func (t *SetGameScore) SetNilChatID() { t.ChatID = nil }
func (t *SetGameScore) WithNilChatID() *SetGameScore { t.ChatID = nil; return t }
func (t *SetGameScore) GetMessageID() int { if t.MessageID == nil { return 0 } else { return *t.MessageID } }
func (t *SetGameScore) SetMessageID(v int) { t.MessageID = &v }
func (t *SetGameScore) WithMessageID(v int) *SetGameScore { t.MessageID = &v; return t }
func (t *SetGameScore) SetNilMessageID() { t.MessageID = nil }
func (t *SetGameScore) WithNilMessageID() *SetGameScore { t.MessageID = nil; return t }
func (t *SetGameScore) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *SetGameScore) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *SetGameScore) WithInlineMessageID(v string) *SetGameScore { t.InlineMessageID = &v; return t }
func (t *SetGameScore) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *SetGameScore) WithNilInlineMessageID() *SetGameScore { t.InlineMessageID = nil; return t }
func (t *SetGameScore) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *SetGameScore) WithFmtInlineMessageID(f string, a ...interface{}) *SetGameScore { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }

func (t *GetGameHighScores) GetMethod() MethodName { if t.Method == nil { return "" } else { return *t.Method } }
func (t *GetGameHighScores) SetMethod(v MethodName) { t.Method = &v }
func (t *GetGameHighScores) WithMethod(v MethodName) *GetGameHighScores { t.Method = &v; return t }
func (t *GetGameHighScores) SetNilMethod() { t.Method = nil }
func (t *GetGameHighScores) WithNilMethod() *GetGameHighScores { t.Method = nil; return t }
func (t *GetGameHighScores) GetUserID() int { return t.UserID }
func (t *GetGameHighScores) SetUserID(v int) { t.UserID = v }
func (t *GetGameHighScores) WithUserID(v int) *GetGameHighScores { t.UserID = v; return t }
func (t *GetGameHighScores) GetChatID() int64 { if t.ChatID == nil { return 0 } else { return *t.ChatID } }
func (t *GetGameHighScores) SetChatID(v int64) { t.ChatID = &v }
func (t *GetGameHighScores) WithChatID(v int64) *GetGameHighScores { t.ChatID = &v; return t }
func (t *GetGameHighScores) SetNilChatID() { t.ChatID = nil }
func (t *GetGameHighScores) WithNilChatID() *GetGameHighScores { t.ChatID = nil; return t }
func (t *GetGameHighScores) GetMessageID() int { if t.MessageID == nil { return 0 } else { return *t.MessageID } }
func (t *GetGameHighScores) SetMessageID(v int) { t.MessageID = &v }
func (t *GetGameHighScores) WithMessageID(v int) *GetGameHighScores { t.MessageID = &v; return t }
func (t *GetGameHighScores) SetNilMessageID() { t.MessageID = nil }
func (t *GetGameHighScores) WithNilMessageID() *GetGameHighScores { t.MessageID = nil; return t }
func (t *GetGameHighScores) GetInlineMessageID() string { if t.InlineMessageID == nil { return "" } else { return *t.InlineMessageID } }
func (t *GetGameHighScores) SetInlineMessageID(v string) { t.InlineMessageID = &v }
func (t *GetGameHighScores) WithInlineMessageID(v string) *GetGameHighScores { t.InlineMessageID = &v; return t }
func (t *GetGameHighScores) SetNilInlineMessageID() { t.InlineMessageID = nil }
func (t *GetGameHighScores) WithNilInlineMessageID() *GetGameHighScores { t.InlineMessageID = nil; return t }
func (t *GetGameHighScores) SetFmtInlineMessageID(f string, a ...interface{}) { t.SetInlineMessageID(fmt.Sprintf(f, a...)) }
func (t *GetGameHighScores) WithFmtInlineMessageID(f string, a ...interface{}) *GetGameHighScores { return t.WithInlineMessageID(fmt.Sprintf(f, a...)) }

func (t *GameHighScore) GetPosition() int { return t.Position }
func (t *GameHighScore) SetPosition(v int) { t.Position = v }
func (t *GameHighScore) WithPosition(v int) *GameHighScore { t.Position = v; return t }
func (t *GameHighScore) GetUser() *User { return t.User }
func (t *GameHighScore) SetUser(v *User) { t.User = v }
func (t *GameHighScore) WithUser(v *User) *GameHighScore { t.User = v; return t }
func (t *GameHighScore) GetScore() int { return t.Score }
func (t *GameHighScore) SetScore(v int) { t.Score = v }
func (t *GameHighScore) WithScore(v int) *GameHighScore { t.Score = v; return t }
