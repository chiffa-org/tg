package tg

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type typeField struct {
	index     int
	name      string
	kind      reflect.Kind
	takeElem  bool
	omitEmpty bool
}

func inspectTag(tag string) (name string, omitEmpty bool) {
	opts := strings.Split(tag, ",")
	if len(opts) == 0 {
		return "", false
	}
	for _, opt := range opts[1:] {
		if opt == "omitempty" {
			return opts[0], true
		}
	}
	return opts[0], false
}

func inspectField(f reflect.StructField) (field typeField, include bool) {
	tag := f.Tag.Get("json")
	if tag == "-" {
		return typeField{}, false
	}
	name, omitEmpty := inspectTag(tag)
	if name == "" {
		name = f.Name
	}
	kind := f.Type.Kind()
	if omitEmpty && kind != reflect.Ptr && kind != reflect.Interface {
		panic(fmt.Errorf("field %s has omitempty tag but is not nillable", f.Name))
	}
	takeElem := false
	if kind == reflect.Ptr {
		kind = f.Type.Elem().Kind()
		switch kind {
		case reflect.Bool, reflect.String, reflect.Float32, reflect.Float64,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if !omitEmpty {
				panic(fmt.Errorf("field %s is nillable but has not omitempty tag", f.Name))
			}
			takeElem = true
		}
	}
	return typeField{0, name, kind, takeElem, omitEmpty}, true
}

func inspectType(t reflect.Type) []typeField {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Errorf("type %s is not a struct", t))
	}
	n := t.NumField()
	fields := make([]typeField, 0, n)
	for i := 0; i < n; i++ {
		field, include := inspectField(t.Field(i))
		if include {
			field.index = i
			fields = append(fields, field)
		}
	}
	return fields
}

var typeFields = make(map[reflect.Type][]typeField)

func init() {
	requests := []Request{
		(*GetUpdates)(nil),
		(*SetWebhook)(nil),
		(*DeleteWebhook)(nil),
		(*GetWebhookInfo)(nil),
		(*GetMe)(nil),
		(*SendMessage)(nil),
		(*ForwardMessage)(nil),
		(*SendPhoto)(nil),
		(*SendAudio)(nil),
		(*SendDocument)(nil),
		(*SendSticker)(nil),
		(*SendVideo)(nil),
		(*SendVoice)(nil),
		(*SendLocation)(nil),
		(*SendVenue)(nil),
		(*SendContact)(nil),
		(*SendChatAction)(nil),
		(*GetUserProfilePhotos)(nil),
		(*GetFile)(nil),
		(*KickChatMember)(nil),
		(*LeaveChat)(nil),
		(*UnbanChatMember)(nil),
		(*GetChat)(nil),
		(*GetChatAdministrators)(nil),
		(*GetChatMembersCount)(nil),
		(*GetChatMember)(nil),
		(*AnswerCallbackQuery)(nil),
		(*EditMessageText)(nil),
		(*EditMessageCaption)(nil),
		(*EditMessageReplyMarkup)(nil),
		(*AnswerInlineQuery)(nil),
		(*SendGame)(nil),
		(*SetGameScore)(nil),
		(*GetGameHighScores)(nil),
	}
	for _, r := range requests {
		t := reflect.TypeOf(r)
		typeFields[t] = inspectType(t)
	}
}

func values(r Request) (url.Values, error) {
	fields, ok := typeFields[reflect.TypeOf(r)]
	if !ok {
		return nil, fmt.Errorf("type %s is not supported", reflect.TypeOf(r))
	}
	v := reflect.ValueOf(r).Elem()
	urlValues := make(url.Values, len(fields))
	for _, field := range fields {
		v := v.Field(field.index)
		if field.omitEmpty && v.IsNil() {
			continue
		}
		if field.takeElem {
			v = v.Elem()
		}
		var value string
		switch field.kind {
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		case reflect.Float32:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 64)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = strconv.FormatUint(v.Uint(), 10)
		default:
			bs, err := json.Marshal(v.Interface())
			if err != nil {
				return nil, err
			}
			value = string(bs)
		}
		urlValues[field.name] = append(urlValues[field.name], value)
	}
	return urlValues, nil
}
