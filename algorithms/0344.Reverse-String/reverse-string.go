package problem0344

func reverseString(s []byte) {
	// 頭尾互換，並一層一層地向內。
	// 在 for 迴圈中，
	// 初始化： i 為起點位置、 j 為終點位置；
	// 條件為：起點位置 小於 終點位置 的時候；
	// Afterthought 事後動作：起點往後、終點往前。
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	// 使用 golang 的方式即可將 js 的解題方式寫成兩行，
	// 世界又更美好了，Inner Peace。
}
