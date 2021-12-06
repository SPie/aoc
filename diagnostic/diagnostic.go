package diagnostic

import (
	"strconv"
)

func CalculateRateFromBinaries(value1String string, value2String string) (int64, error) {
	value1, err := strconv.ParseInt(value1String, 2, 64)
	if err != nil {
		return 0, err
	}
	value2, err := strconv.ParseInt(value2String, 2, 64)
	if err != nil {
		return 0, err
	}

	return value1 * value2, nil
}