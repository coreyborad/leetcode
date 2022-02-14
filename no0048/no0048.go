package no0048

func rotate(matrix [][]int) {
	length := len(matrix)
	// 對角線轉換
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 每一行做反轉
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			// 起點=j, 終點=length-j-1
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}

/*
	123
	456
	789

	147
	256
	389

	147
	258
	369


*/
