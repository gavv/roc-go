// Code generated by "stringer -type Format -trimprefix Format -output format_string.go"; DO NOT EDIT.

package roc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FormatPcmFloat32-1]
}

const _Format_name = "PcmFloat32"

var _Format_index = [...]uint8{0, 10}

func (i Format) String() string {
	i -= 1
	if i < 0 || i >= Format(len(_Format_index)-1) {
		return "Format(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Format_name[_Format_index[i]:_Format_index[i+1]]
}