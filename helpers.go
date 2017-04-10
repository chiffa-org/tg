package tg

import (
	"strings"
	"unicode"
)

func (u *Update) GetFrom() *User {
	switch {
	case u.Message != nil:
		return u.Message.From
	case u.EditedMessage != nil:
		return u.EditedMessage.From
	case u.ChannelPost != nil:
		return u.ChannelPost.From
	case u.EditedChannelPost != nil:
		return u.EditedChannelPost.From
	case u.InlineQuery != nil:
		return u.InlineQuery.From
	case u.ChosenInlineResult != nil:
		return u.ChosenInlineResult.From
	case u.CallbackQuery != nil:
		return u.CallbackQuery.From
	default:
		return nil
	}
}

func (u *Update) GetChat() *Chat {
	switch {
	case u.Message != nil:
		return u.Message.Chat
	case u.EditedMessage != nil:
		return u.EditedMessage.Chat
	case u.ChannelPost != nil:
		return u.ChannelPost.Chat
	case u.EditedChannelPost != nil:
		return u.EditedChannelPost.Chat
	case u.CallbackQuery != nil && u.CallbackQuery.Message != nil:
		return u.CallbackQuery.Message.Chat
	default:
		return nil
	}
}

func (m *Message) IsCommand() bool {
	return m.Text != nil && len(*m.Text) > 0 && (*m.Text)[0] == '/'
}

func (m *Message) Command() string {
	text := m.GetText()
	if idx := strings.IndexRune(text, '/'); idx != 0 {
		return ""
	}
	r := len(text)
	if idx := strings.IndexRune(text, '@'); idx > 0 && idx < r {
		r = idx
	}
	if idx := strings.IndexFunc(text, unicode.IsSpace); idx > 0 && idx < r {
		r = idx
	}
	return text[1:r]
}

func (m *Message) CommandArgs() string {
	text := m.GetText()
	if idx := strings.IndexRune(text, '/'); idx != 0 {
		return ""
	}
	if idx := strings.IndexFunc(text, unicode.IsSpace); idx > 0 {
		return text[idx:]
	}
	return ""
}

func NewKeyboardRow(row ...*KeyboardButton) []*KeyboardButton                         { return row }
func NewKeyboardRows(rows ...[]*KeyboardButton) [][]*KeyboardButton                   { return rows }
func NewInlineKeyboardRow(row ...*InlineKeyboardButton) []*InlineKeyboardButton       { return row }
func NewInlineKeyboardRows(rows ...[]*InlineKeyboardButton) [][]*InlineKeyboardButton { return rows }
