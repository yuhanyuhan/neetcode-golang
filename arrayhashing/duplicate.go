package arrayhashing

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, v := range nums {
		// ok comma checks if key-value item exist in set, making it a set
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}

	return false
}
