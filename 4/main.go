package main

func ChangeSlice(sl1 []string, sl2 []string) []string {
	var res []string
	for _, v1 := range sl1 {
		for _, v2 := range sl2 {
			if v1 == v2 {
				res = append(res, v1)
			}
		}
	}
	return res
}
