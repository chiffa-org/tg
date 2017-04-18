package main

import (
	"bytes"
	"fmt"
	"github.com/chiffa-org/tg"
	"io"
	"os"
	"reflect"
	"text/template"
)

const (
	Get = `func (t *{{.Type}}) Get{{.FName}}() {{.FType}} { return t.{{.FName}} }`
	Set = `func (t *{{.Type}}) Set{{.FName}}(v {{.FType}}) { t.{{.FName}} = v }
func (t *{{.Type}}) With{{.FName}}(v {{.FType}}) *{{.Type}} { t.{{.FName}} = v; return t }`
	SetNil = `func (t *{{.Type}}) SetNil{{.FName}}() { t.{{.FName}} = nil }
func (t *{{.Type}}) WithNil{{.FName}}() *{{.Type}} { t.{{.FName}} = nil; return t }`
	SetFmt = `func (t *{{.Type}}) SetFmt{{.FName}}(f string, a ...interface{}) { t.Set{{.FName}}(fmt.Sprintf(f, a...)) }
func (t *{{.Type}}) WithFmt{{.FName}}(f string, a ...interface{}) *{{.Type}} { return t.With{{.FName}}(fmt.Sprintf(f, a...)) }`
	GetZeroPtr = `func (t *{{.Type}}) Get{{.FName}}() {{.FType}} { if t.{{.FName}} == nil { return {{.Zero}} } else { return *t.{{.FName}} } }`
	SetPrimPtr = `func (t *{{.Type}}) Set{{.FName}}(v {{.FType}}) { t.{{.FName}} = &v }
func (t *{{.Type}}) With{{.FName}}(v {{.FType}}) *{{.Type}} { t.{{.FName}} = &v; return t }`
	SetSlicePtr = `func (t *{{.Type}}) Set{{.FName}}(v {{.FType}}) { if v == nil { t.{{.FName}} = nil } else { t.{{.FName}} = &v } }
func (t *{{.Type}}) With{{.FName}}(v {{.FType}}) *{{.Type}} { if v == nil { t.{{.FName}} = nil } else { t.{{.FName}} = &v }; return t }`
)

var (
	GetTmpl         = template.Must(template.New("Get").Parse(Get))
	SetTmpl         = template.Must(template.New("Set").Parse(Set))
	SetNilTmpl      = template.Must(template.New("SetNil").Parse(SetNil))
	SetFmtTmpl      = template.Must(template.New("SetFmt").Parse(SetFmt))
	GetZeroPtrTmpl  = template.Must(template.New("GetZeroPtr").Parse(GetZeroPtr))
	SetPrimPtrTmpl  = template.Must(template.New("SetPrimPtr").Parse(SetPrimPtr))
	SetSlicePtrTmpl = template.Must(template.New("SetSlicePtr").Parse(SetSlicePtr))
)

type Params struct {
	Type  string
	FName string
	FType string
	Zero  string
}

func TypeString(t reflect.Type) string {
	b := new(bytes.Buffer)
l:
	for {
		switch t.Kind() {
		case reflect.Ptr:
			b.WriteString("*")
			t = t.Elem()
		case reflect.Slice:
			b.WriteString("[]")
			t = t.Elem()
		default:
			break l
		}
	}
	b.WriteString(t.Name())
	return b.String()
}

func ZeroString(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice:
		return "nil"
	case reflect.Struct:
		return t.Name() + "{}"
	default:
		return fmt.Sprintf("%#v", reflect.Zero(t).Interface())
	}
}

func Execute(w io.Writer, v interface{}) error {
	t := reflect.TypeOf(v)
	typeName := t.Name()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type
		params := Params{
			Type:  typeName,
			FName: f.Name,
		}
		switch ft.Kind() {
		case reflect.Ptr:
			switch ft.Elem().Kind() {
			case reflect.Struct:
				params.FType = TypeString(ft)
				if err := GetTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
				if err := SetTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
			case reflect.Slice:
				ft = ft.Elem()
				params.FType = TypeString(ft)
				params.Zero = ZeroString(ft)
				if err := GetZeroPtrTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
				if err := SetSlicePtrTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
			default:
				ft = ft.Elem()
				params.FType = TypeString(ft)
				params.Zero = ZeroString(ft)
				if err := GetZeroPtrTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
				if err := SetPrimPtrTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
				if err := SetNilTmpl.Execute(w, params); err != nil {
					return err
				}
				fmt.Fprintln(w)
			}
		default:
			params.FType = TypeString(ft)
			if err := GetTmpl.Execute(w, params); err != nil {
				return err
			}
			fmt.Fprintln(w)
			if err := SetTmpl.Execute(w, params); err != nil {
				return err
			}
			fmt.Fprintln(w)
		}
		if ft == reflect.TypeOf("") {
			if err := SetFmtTmpl.Execute(w, params); err != nil {
				return err
			}
			fmt.Fprintln(w)
		}
	}
	fmt.Fprintln(w)
	return nil
}

func main() {
	types := []interface{}{
		tg.Update{},
		tg.GetUpdates{},
		tg.SetWebhook{},
		tg.DeleteWebhook{},
		tg.GetWebhookInfo{},
		tg.WebhookInfo{},

		tg.User{},
		tg.Chat{},
		tg.Message{},
		tg.MessageEntity{},
		tg.PhotoSize{},
		tg.Audio{},
		tg.Document{},
		tg.Sticker{},
		tg.Video{},
		tg.Voice{},
		tg.Contact{},
		tg.Location{},
		tg.Venue{},
		tg.UserProfilePhotos{},
		tg.File{},
		tg.ReplyKeyboardMarkup{},
		tg.KeyboardButton{},
		tg.ReplyKeyboardRemove{},
		tg.InlineKeyboardMarkup{},
		tg.InlineKeyboardButton{},
		tg.CallbackQuery{},
		tg.ForceReply{},
		tg.ChatMember{},
		tg.ResponseParameters{},

		tg.GetMe{},
		tg.SendMessage{},
		tg.ForwardMessage{},
		tg.SendPhoto{},
		tg.SendAudio{},
		tg.SendDocument{},
		tg.SendSticker{},
		tg.SendVideo{},
		tg.SendVoice{},
		tg.SendLocation{},
		tg.SendVenue{},
		tg.SendContact{},
		tg.SendChatAction{},
		tg.GetUserProfilePhotos{},
		tg.GetFile{},
		tg.KickChatMember{},
		tg.LeaveChat{},
		tg.UnbanChatMember{},
		tg.GetChat{},
		tg.GetChatAdministrators{},
		tg.GetChatMembersCount{},
		tg.GetChatMember{},
		tg.AnswerCallbackQuery{},

		tg.EditMessageText{},
		tg.EditMessageCaption{},
		tg.EditMessageReplyMarkup{},

		tg.InlineQuery{},
		tg.AnswerInlineQuery{},
		tg.InlineQueryResultArticle{},
		tg.InlineQueryResultPhoto{},
		tg.InlineQueryResultGIF{},
		tg.InlineQueryResultMPEG4GIF{},
		tg.InlineQueryResultVideo{},
		tg.InlineQueryResultAudio{},
		tg.InlineQueryResultVoice{},
		tg.InlineQueryResultDocument{},
		tg.InlineQueryResultLocation{},
		tg.InlineQueryResultVenue{},
		tg.InlineQueryResultContact{},
		tg.InlineQueryResultGame{},
		tg.InlineQueryResultCachedPhoto{},
		tg.InlineQueryResultCachedGIF{},
		tg.InlineQueryResultCachedMPEG4GIF{},
		tg.InlineQueryResultCachedSticker{},
		tg.InlineQueryResultCachedDocument{},
		tg.InlineQueryResultCachedVideo{},
		tg.InlineQueryResultCachedVoice{},
		tg.InlineQueryResultCachedAudio{},
		tg.InputTextMessageContent{},
		tg.InputLocationMessageContent{},
		tg.InputVenueMessageContent{},
		tg.InputContactMessageContent{},
		tg.ChosenInlineResult{},

		tg.SendGame{},
		tg.Game{},
		tg.Animation{},
		//tg.CallbackGame{},
		tg.SetGameScore{},
		tg.GetGameHighScores{},
		tg.GameHighScore{},
	}

	for _, t := range types {
		Execute(os.Stdout, t)
	}
}
