package medium

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

//ABCDEF
// A
// B  F
// C E
// D

func Convert(s string, numRows int) string {
	var j_const_number int = numRows + numRows - 2
	var j_1_num int
	var j_2_num int
	var j_cur int
	var j_bool bool = true
	var index int = 0
	var runs rune

	result := make([]rune, 0, len(s))

	if len(s) < 2 || numRows == 1 {
		return s
	}

	if len(s) < numRows {
		numRows = len(s)
	}

	for row := 0; row < numRows; row++ {
		index = row
		appended_arr := make([]rune, 0, len(s)/numRows)
		result = append(result, rune(s[row]))

		j_1_num = j_const_number - j_2_num

		j_bool = true

		if j_1_num == 0 {
			j_bool = false
		} else if j_2_num == 0 {
			j_bool = true
		}

		for {
			if j_bool {
				j_cur = j_1_num
			} else {
				j_cur = j_2_num
			}

			if j_cur+index >= len(s) {
				break
			}

			runs = rune(s[index+j_cur])

			appended_arr = append(appended_arr, runs)
			if !(j_1_num == 0 || j_2_num == 0) {
				j_bool = !j_bool
			}
			index += j_cur
		}

		result = append(result, appended_arr...)
		j_2_num += 2
	}

	return string(result)
}
