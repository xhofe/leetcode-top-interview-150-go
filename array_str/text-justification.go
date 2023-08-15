package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	_words := make([]string, 0)
	cnt := 0
	for _, v := range words {
		if cnt+len(_words)+len(v) > maxWidth {
			if len(_words) == 1 {
				tmp := _words[0] + strings.Repeat(" ", maxWidth-cnt)
				ans = append(ans, tmp)
			} else {
				tmp := strings.Builder{}
				spaceCnt := maxWidth - cnt
				spaceOne := spaceCnt / (len(_words) - 1)
				spaceExtra := spaceCnt % (len(_words) - 1)
				for i, w := range _words {
					tmp.WriteString(w)
					space := spaceOne
					if spaceExtra > 0 {
						spaceExtra--
						space++
					}
					if i < len(_words)-1 {
						tmp.WriteString(strings.Repeat(" ", space))
					}
				}
				ans = append(ans, tmp.String())
			}
			cnt = 0
			_words = make([]string, 0)
		}
		_words = append(_words, v)
		cnt += len(v)
	}
	ans = append(ans, strings.Join(_words, " ")+strings.Repeat(" ", maxWidth-cnt-len(_words)+1))
	return ans
}
