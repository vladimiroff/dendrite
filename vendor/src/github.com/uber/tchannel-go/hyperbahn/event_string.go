// generated by stringer -type=Event; DO NOT EDIT

package hyperbahn

import "fmt"

const _Event_name = "UnknownEventRegistrationAttemptRegisteredRegistrationRefreshed"

var _Event_index = [...]uint8{0, 12, 31, 41, 62}

func (i Event) String() string {
	if i < 0 || i+1 >= Event(len(_Event_index)) {
		return fmt.Sprintf("Event(%d)", i)
	}
	return _Event_name[_Event_index[i]:_Event_index[i+1]]
}