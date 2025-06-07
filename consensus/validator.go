package consensus

type Validator struct {
	ID		string	// Node Identifier
	Stake	int		// Staked amount
}

type ValidatorPool struct {
	Validators []Validator
}

// functions
// addValidator
// getValidators [TODO: We'll see if they are required]
// getValidatorById

func NewValidatorPool() *ValidatorPool {
	return &ValidatorPool{}
}

func (vp *ValidatorPool) AddValidator(id string, stake int) {
	vp.Validators = append(vp.Validators, Validator{ID: id, Stake: stake})
}

// get total stake amts
func (vp *ValidatorPool) GetTotalStake() int {
	total := 0
	for _,v := range vp.Validators {
		total += v.Stake
	}
	return total
}
