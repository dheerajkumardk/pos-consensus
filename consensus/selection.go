package consensus

import "math/rand"

func SelectValidator(vp *ValidatorPool) string {
	totalStake := vp.GetTotalStake()
	target := rand.Intn(totalStake)
	cumulative := 0

	for _,v := range vp.Validators {
		cumulative += v.Stake
		if cumulative >= target {
			return v.ID
		}
	}
	return ""
}

// Block proposal and validation flow
// - leader: creates a new block
// 	- broadcast it to other validators

// - validators: verify the block
// 	- send votes