package utils

func removeDuplicates(users []NameChange) []NameChange {
	uniqueIDs := make(map[int]struct{})
	result := make([]NameChange, 0)

	for _, user := range users {
		// Check if the ID is already in the map
		if _, exists := uniqueIDs[user.PlayerId]; !exists {
			// Add the object to the result slice and mark the ID as seen
			result = append(result, user)
			uniqueIDs[user.PlayerId] = struct{}{}
		}
	}

	return result
}
