package Q0854

import (
	"fmt"
	"testing"
)

type pair struct {
	s    string // current string
	i    int    // current index, s[j - 1] == s2[j - 1], j = 1, 2, 3...i - 1
	step int
}

func kSimilarity(s1 string, s2 string) (res int) {
	if len(s1) <= 1 {
		return 0
	}

	// bfs
	q := []pair{{s: s1, i: 0, step: 0}}
	vis := map[string]struct{}{s1: {}} // visited string
	for len(q) > 0 {
		nextQ := make([]pair, 0)

		for index := 0; index < len(q); index++ {
			s, i, step := q[index].s, q[index].i, q[index].step

			tmp := []byte(s)
			// find diff
			for i < len(s) && s[i] == s2[i] {
				i++
			}

			if i == len(s1) {
				res = step
				break
			}

			// swap s[i] and s[j], j > i
			for j := i + 1; j < len(s1); j++ {
				if s[j] == s2[i] && s[j] != s2[j] {
					tmp[i], tmp[j] = tmp[j], tmp[i]
					str := string(tmp)
					if _, ok := vis[str]; !ok {
						vis[str] = struct{}{}
						nextQ = append(nextQ, pair{s: str, i: i + 1, step: step + 1})
					}
					tmp[i], tmp[j] = tmp[j], tmp[i]
				}
			}
		}

		q = nextQ
	}

	return res
}

func Test(t *testing.T) {
	s1 := "bccaba"
	s2 := "abacbc"
	fmt.Println(kSimilarity(s1, s2))
}

func Test1(t *testing.T) {
	for i := 0; i < call(); i++ {
		fmt.Println(i)
	}
}

var val = 1

func call() int {
	res := val
	val += 1
	return res
}
