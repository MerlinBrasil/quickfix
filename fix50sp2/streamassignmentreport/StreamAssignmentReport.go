//Package streamassignmentreport msg type = CD.
package streamassignmentreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a StreamAssignmentReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//StreamAsgnRptID is a required field for StreamAssignmentReport.
func (m Message) StreamAsgnRptID() (*field.StreamAsgnRptIDField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReport.
func (m Message) GetStreamAsgnRptID(f *field.StreamAsgnRptIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqType is a non-required field for StreamAssignmentReport.
func (m Message) StreamAsgnReqType() (*field.StreamAsgnReqTypeField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqType reads a StreamAsgnReqType from StreamAssignmentReport.
func (m Message) GetStreamAsgnReqType(f *field.StreamAsgnReqTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqID is a non-required field for StreamAssignmentReport.
func (m Message) StreamAsgnReqID() (*field.StreamAsgnReqIDField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqID reads a StreamAsgnReqID from StreamAssignmentReport.
func (m Message) GetStreamAsgnReqID(f *field.StreamAsgnReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoAsgnReqs is a non-required field for StreamAssignmentReport.
func (m Message) NoAsgnReqs() (*field.NoAsgnReqsField, quickfix.MessageRejectError) {
	f := &field.NoAsgnReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAsgnReqs reads a NoAsgnReqs from StreamAssignmentReport.
func (m Message) GetNoAsgnReqs(f *field.NoAsgnReqsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds StreamAssignmentReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for StreamAssignmentReport.
func Builder(
	streamasgnrptid *field.StreamAsgnRptIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("CD"))
	builder.Body().Set(streamasgnrptid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CD", r
}
