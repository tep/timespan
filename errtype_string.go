// Code generated by "stringer -type=errType"; DO NOT EDIT

package timespan

import "fmt"

const _errType_name = "noErrmissplacedDashErrmissingCoefErrunparseableCoefErrunrecognizedMagErrmagnOrderErrmagnRestatedErrmagnOrderError"

var _errType_index = [...]uint8{0, 5, 22, 36, 54, 72, 84, 99, 113}

func (i errType) String() string {
	if i < 0 || i >= errType(len(_errType_index)-1) {
		return fmt.Sprintf("errType(%d)", i)
	}
	return _errType_name[_errType_index[i]:_errType_index[i+1]]
}