package main

func FindIntersections(a, b []int) (bool, []int) {
	set := make(map[int]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}

	var intersections []int
	for _, v := range b {
		if _, exists := set[v]; exists {
			intersections = append(intersections, v)
		}
	}

	return len(intersections) > 0, intersections
}
