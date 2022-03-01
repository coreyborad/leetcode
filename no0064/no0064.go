package no0064

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 左邊為邊界
	// 條件為 dp[i][j]=dp[i][j−1]+grid[i][j]
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]

	}
	// 右邊為邊界
	// 條件為 dp[i][j]=dp[i−1][j]+grid[i][j]
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	// 都不為邊界
	// dp[i][j]=min(dp[i−1][j],dp[i][j−1])+grid[i][j]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]

}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
1 2 1
3 5 20
4 2 1

11

*/
