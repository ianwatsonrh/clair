// Code generated by "stringer -type itemKind lex.go"; DO NOT EDIT.

package dockerfile

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[itemError-0]
	_ = x[itemComment-1]
	_ = x[itemInstruction-2]
	_ = x[itemLabel-3]
	_ = x[itemArg-4]
	_ = x[itemEnv-5]
	_ = x[itemEOF-6]
}

const _itemKind_name = "itemErroritemCommentitemInstructionitemLabelitemArgitemEnvitemEOF"

var _itemKind_index = [...]uint8{0, 9, 20, 35, 44, 51, 58, 65}

func (i itemKind) String() string {
	if i < 0 || i >= itemKind(len(_itemKind_index)-1) {
		return "itemKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _itemKind_name[_itemKind_index[i]:_itemKind_index[i+1]]
}
