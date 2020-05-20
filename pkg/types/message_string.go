// Code generated by "stringer -type=Message"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WM_LBUTTONDOWN-513]
	_ = x[WM_LBUTTONUP-514]
	_ = x[WM_MOUSEMOVE-512]
	_ = x[WM_MOUSEWHEEL-522]
	_ = x[WM_MOUSEHWHEEL-526]
	_ = x[WM_RBUTTONDOWN-516]
	_ = x[WM_RBUTTONUP-517]
	_ = x[WM_KEYDOWN-256]
	_ = x[WM_KEYUP-257]
	_ = x[WM_SYSKEYDOWN-260]
	_ = x[WM_SYSKEYUP-261]
}

const (
	_Message_name_0 = "WM_KEYDOWNWM_KEYUP"
	_Message_name_1 = "WM_SYSKEYDOWNWM_SYSKEYUP"
	_Message_name_2 = "WM_MOUSEMOVEWM_LBUTTONDOWNWM_LBUTTONUP"
	_Message_name_3 = "WM_RBUTTONDOWNWM_RBUTTONUP"
	_Message_name_4 = "WM_MOUSEWHEEL"
	_Message_name_5 = "WM_MOUSEHWHEEL"
)

var (
	_Message_index_0 = [...]uint8{0, 10, 18}
	_Message_index_1 = [...]uint8{0, 13, 24}
	_Message_index_2 = [...]uint8{0, 12, 26, 38}
	_Message_index_3 = [...]uint8{0, 14, 26}
)

func (i Message) String() string {
	switch {
	case 256 <= i && i <= 257:
		i -= 256
		return _Message_name_0[_Message_index_0[i]:_Message_index_0[i+1]]
	case 260 <= i && i <= 261:
		i -= 260
		return _Message_name_1[_Message_index_1[i]:_Message_index_1[i+1]]
	case 512 <= i && i <= 514:
		i -= 512
		return _Message_name_2[_Message_index_2[i]:_Message_index_2[i+1]]
	case 516 <= i && i <= 517:
		i -= 516
		return _Message_name_3[_Message_index_3[i]:_Message_index_3[i+1]]
	case i == 522:
		return _Message_name_4
	case i == 526:
		return _Message_name_5
	default:
		return "Message(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}