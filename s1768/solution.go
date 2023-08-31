package s1768

// two pointers
func MergeAlternately(word1 string, word2 string) string {
	result := ""

	lenWord1 := len(word1)
	lenWord2 := len(word2)

	i := 0
	j := 0

	for i < lenWord1 && j < lenWord2 {
		result += string(word1[i])
		result += string(word2[j])
		i++
		j++
	}

	for i < lenWord1 {
		result += string(word1[i])
		i++
	}

	for j < lenWord2 {
		result += string(word2[j])
		j++
	}

	return result
}

// one pointer
// func MergeAlternately(word1 string, word2 string) string {
// 	result := ""
// 	len1 := len(word1)
// 	len2 := len(word2)

// 	cnt := max(len1, len2)
// 	for i := 0; i < cnt; i++ {
// 		if i < len1 {
// 			result += string(word1[i])
// 		}
// 		if i < len2 {
// 			result += string(word2[i])
// 		}
// 	}
// 	return result
// }

// modifed one pointer
// func MergeAlternately(word1 string, word2 string) string {
// 	result := ""
// 	len1 := len(word1)
// 	len2 := len(word2)

// 	cnt := max(len1, len2)
// 	for i := 0; i < cnt; i++ {
// 		if i == len1 && i == len2 {
// 			return result
// 		} else if i == len1 {
// 			result += word2[i:]
// 			return result
// 		} else if i == len2 {
// 			result += word1[i:]
// 			return result
// 		}
// 		result += string(word1[i])
// 		result += string(word2[i])
// 	}
// 	return result
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
