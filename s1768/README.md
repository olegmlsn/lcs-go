## 1768. Merge Strings Alternately

`#Easy` `#LeetCode75`

[on LeetCode](https://leetcode.com/problems/merge-strings-alternately)

<details>
<summary>in English</summary>

You are given two strings `word1` and `word2`. Merge the strings by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

</details>

<details>
<summary>in Russian</summary>

Вам даны две строки `word1` и `word2`. Объедините строки, добавляя поочередно буквы, начиная с `word1`. Если одна из строк длиннее другой, добавьте оставшиеся буквы в конец объединенной строки.

Верните объединенную строку.

</details>

<details>
<summary>Examples</summary>

**Example 1:**
```
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
```

**Example 2:**

```
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b 
word2:    p   q   r   s
merged: a p b q   r   s
```

**Example 3:**

```
Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q 
merged: a p b q c   d
```

**Constraints:**

- `1 <= word1.length, word2.length <= 100`
- `word1` and `word2` consist of lowercase English letters.

</details>

## Solutions

<details>
<summary>two pointers</summary>

Проходим по строкам двумя указателями, пока оба указателя меньше длинны соответствующего слова.
После проходим по оставшейся части каждого слова, если она есть, по отдельности.

```go
// classic two pointer
func mergeAlternately(word1 string, word2 string) string {
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

```

</details>

<details>
<summary>one pointer</summary>

Итерируемся от 0 до длинны самого длинного слова. На каждой итерации складываем
в результат букву от каждого слова, проверяя что указатель меньше длинны слова.

```go
func mergeAlternately(word1 string, word2 string) string {
	result := ""
	len1 := len(word1)
	len2 := len(word2)

	cnt := max(len1, len2)
	for i := 0; i < cnt; i++ {
		if i < len1 {
			result += string(word1[i])
		}
		if i < len2 {
			result += string(word2[i])
		}
	}
	return result
}

```

</details>

</details>

<details>
<summary>modifed one pointer</summary>

```go
func mergeAlternately(word1 string, word2 string) string {
	result := ""
	len1 := len(word1)
	len2 := len(word2)

	cnt := max(len1, len2)
	for i := 0; i < cnt; i++ {
		if i == len1 && i == len2 {
			return result
		} else if i == len1 {
			result += word2[i:]
			return result
		} else if i == len2 {
			result += word1[i:]
			return result
		}
		result += string(word1[i])
		result += string(word2[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

```

</details>