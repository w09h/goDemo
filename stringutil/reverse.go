//package stsringutil contains util funcs for workings with strings
package stringutil

//reverse returns its arg in rev. rune-wise left to right
//ABC -> CBA
func Reverse(s string) string {

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
