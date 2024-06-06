// Code generated by "stringer -type ClockSyncProfile -trimprefix ClockSyncProfile -output clock_sync_profile_string.go"; DO NOT EDIT.

package roc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ClockSyncProfileDefault-0]
	_ = x[ClockSyncProfileResponsive-1]
	_ = x[ClockSyncProfileGradual-2]
}

const _ClockSyncProfile_name = "DefaultResponsiveGradual"

var _ClockSyncProfile_index = [...]uint8{0, 7, 17, 24}

func (i ClockSyncProfile) String() string {
	if i < 0 || i >= ClockSyncProfile(len(_ClockSyncProfile_index)-1) {
		return "ClockSyncProfile(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ClockSyncProfile_name[_ClockSyncProfile_index[i]:_ClockSyncProfile_index[i+1]]
}