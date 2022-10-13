package geartoothc

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func find_in_fnmap(fnname string) [2]string {
	for i := 0; i < len(fnmap); i++ {
		curr := fnmap[i]
		if curr[0] == fnname {
			return curr
		}
	}
	return [2]string{"ERROR", "ERROR"}
}
