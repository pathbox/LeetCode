// 滑动窗口
// https://www.bilibili.com/video/BV15D4y1X7Tt?p=2

package Offer0102

func containExactly(str, aim string) int {
	count := make(map[byte]int)
	M := len(aim)
	R := 0
	for i := 0; i < M; i++ {
		count[aim[i]]++
	}
	inValidTimes := 0
	// 先让窗口拥有M个字符
	for ; R < M; R++ {
		count[str[R]]--
		if count[str[R]] <= 0 {
			inValidTimes++
		}
	}

	for ; R < len(str); R++ {
		// 判断前一个窗口符合否
		if inValidTimes == 0 {
			return R - M // 返回left索引
		}
		count[aim[R]]--
		if count[str[R]] <= 0 {
			inValidTimes++
		}
		count[str[R-M]]++
		if count[str[R-M]] < 0 {
			inValidTimes--
		}
	}
	if inValidTimes == 0 {
		return R - M
	}
	return -1
}
