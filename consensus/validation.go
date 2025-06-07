package consensus

// received votes must be greater than minm required
func CheckConsensus(votes map[string]bool, minVotes int) bool {
	yayVotes := 0

	for _, vote := range votes {
		if vote {
			yayVotes++
		}
	}
	return yayVotes >= minVotes
}