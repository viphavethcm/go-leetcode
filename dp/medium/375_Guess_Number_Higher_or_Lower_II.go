package dp_medium

/*
https://leetcode.com/problems/guess-number-higher-or-lower-ii/editorial/
*/

func getMoneyAmount(n int) int {
	// Tạo bảng DP kích thước (n+2) x (n+2).
	// Lý do: Khi ta đoán k, ta sẽ gọi dp[k+1][j]. Nếu k = n, k+1 sẽ thành n+1.
	// Việc tạo dư ra giúp Go không bị lỗi "index out of range" và các ô này mặc định bằng 0.
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// VÒNG LẶP 1: Chạy theo chiều dài của khoảng (từ 2 đến n)
	// Vì khoảng có độ dài = 1 (vd: đoán số 3 trong khoảng [3,3]) thì không tốn đồng nào (dp = 0)
	for length := 2; length <= n; length++ {

		// VÒNG LẶP 2: Xác định điểm bắt đầu 'i' của khoảng
		for i := 1; i <= n-length+1; i++ {

			// Từ i và length, ta tính được điểm kết thúc 'j'
			j := i + length - 1

			// Khởi tạo chi phí cho khoảng [i, j] bằng một số cực lớn (vô cực)
			dp[i][j] = 999999

			// VÒNG LẶP 3: Thử mọi nhát chém 'k' bên trong khoảng [i, j]
			// Mẹo tối ưu: Ta chỉ chạy k tới j-1. Vì đoán số lớn nhất (j) luôn tốn tiền
			// hơn đoán số kế cuối (j-1) mà hiệu quả chia nhỏ khoảng lại kém hơn.
			for k := i; k < j; k++ {

				// HỆ THỐNG CHƠI DƠ: Nó sẽ chọn nhánh [i, k-1] hoặc [k+1, j] tốn nhiều tiền hơn
				// Ta phải lấy max() của 2 nhánh này, cộng với số tiền 'k' vừa mất để đoán
				costIfGuessK := k + max(dp[i][k-1], dp[k+1][j])

				// NGƯỜI CHƠI TỐI ƯU: Ta thử mọi số k, và giữ lại cách đoán có chi phí nhỏ nhất
				dp[i][j] = min(dp[i][j], costIfGuessK)
			}
		}
	}

	// Kết quả cuối cùng là ô chứa chi phí chiến thắng cho toàn bộ khoảng từ 1 đến n
	return dp[1][n]
}
