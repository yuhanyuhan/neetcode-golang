package algorithm

func Permutation(s string) []string {
    if len(s) <= 1 {
        return []string{s}
    }

    result := []string{}
	seen := make(map[string]bool)
    for i := range s {
        c := string(s[i])
        // recursively swap between two substrings 
        for _, p := range Permutation(s[:i] + s[i+1:]) {
			newCombination := c + p
			if !seen[newCombination] {
				result = append(result, newCombination)
				seen[newCombination] = true
			}
        }
    }
    return result
}