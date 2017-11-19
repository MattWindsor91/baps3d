package list

// File controller.go defines the specific Controller logic for lists.

import (
	"fmt"

	"github.com/UniversityRadioYork/baps3d/comm"
)

// NewController constructs a new Controller for a given List.
func NewController(l *List) (*comm.Controller, *comm.Client) {
	return comm.NewController(l)
}

//
// Dump logic
//

// automodeResponse returns c's list's automode as a response.
func (l *List) autoModeResponse() AutoModeResponse {
	return AutoModeResponse{AutoMode: l.AutoMode()}
}

// freezeResponse returns c's list's frozen representation as a response.
func (l *List) freezeResponse() FreezeResponse {
	return l.Freeze()
}

// Dump handles a dump request.
func (l *List) Dump(dumpCb comm.ResponseCb) {
	// SPEC: see https://universityradioyork.github.io/baps3-spec/protocol/roles/lis
	dumpCb(l.autoModeResponse())
	dumpCb(l.freezeResponse())
	// TODO(@MattWindsor91): other items in dump
}

//
// Request handling
//

// HandleRequest handles a request for List l.
func (l *List) HandleRequest(replyCb comm.ResponseCb, bcastCb comm.ResponseCb, rbody interface{}) error {
	var err error

	switch b := rbody.(type) {
	case SetAutoModeRequest:
		if l.SetAutoMode(b.AutoMode) {
			bcastCb(l.autoModeResponse())
		}
	default:
		err = fmt.Errorf("list can't handle this request")
	}

	return err
}
