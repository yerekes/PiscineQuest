package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	resultInt := 0
	baseFromLen := len(baseFrom)
	for _, char := range nbr {
		index := 0
		for i, baseChar := range baseFrom {
			if baseChar == char {
				index = i
				break
			}
		}
		resultInt = resultInt*baseFromLen + index
	}
	baseToLen := len(baseTo)
	if resultInt == 0 {
		return string(baseTo[0])
	}
	resultStr := ""
	for resultInt > 0 {
		remainder := resultInt % baseToLen
		resultStr = string(baseTo[remainder]) + resultStr
		resultInt /= baseToLen
	}
	return resultStr
}
