//Package massquote msg type = i.
package massquote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a MassQuote wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a non-required field for MassQuote.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, quickfix.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from MassQuote.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for MassQuote.
func (m Message) QuoteID() (*field.QuoteIDField, quickfix.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MassQuote.
func (m Message) GetQuoteID(f *field.QuoteIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for MassQuote.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, quickfix.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from MassQuote.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DefBidSize is a non-required field for MassQuote.
func (m Message) DefBidSize() (*field.DefBidSizeField, quickfix.MessageRejectError) {
	f := &field.DefBidSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefBidSize reads a DefBidSize from MassQuote.
func (m Message) GetDefBidSize(f *field.DefBidSizeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DefOfferSize is a non-required field for MassQuote.
func (m Message) DefOfferSize() (*field.DefOfferSizeField, quickfix.MessageRejectError) {
	f := &field.DefOfferSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDefOfferSize reads a DefOfferSize from MassQuote.
func (m Message) GetDefOfferSize(f *field.DefOfferSizeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a required field for MassQuote.
func (m Message) NoQuoteSets() (*field.NoQuoteSetsField, quickfix.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from MassQuote.
func (m Message) GetNoQuoteSets(f *field.NoQuoteSetsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MassQuote messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MassQuote.
func Builder(
	quoteid *field.QuoteIDField,
	noquotesets *field.NoQuoteSetsField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("i"))
	builder.Body().Set(quoteid)
	builder.Body().Set(noquotesets)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "i", r
}
