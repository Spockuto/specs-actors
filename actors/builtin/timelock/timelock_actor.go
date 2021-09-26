package timelock

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	cron0 "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	"github.com/ipfs/go-cid"
	"github.com/filecoin-project/specs-actors/v6/actors/runtime"
	"github.com/filecoin-project/specs-actors/v6/actors/builtin"
)

type Actor struct{}

func (a Actor) Exports() []interface{} {
	return []interface{}{
		builtin.MethodConstructor: a.Constructor,
		2:                         a.LockTransaction,
		3:						   a.CronTick
	}
}

func (a Actor) Constructor(rt runtime.Runtime, _ abi.EmptyValue) {
	// sets up the state to have an empty map of times => transaction sets
	// Look up market actor and market state as it needs to run every epoch as well
	// DealsByEpoch
}

func (a Actor) LockTransaction(unlockEpoch abi.ChainEpoch, txEnc []byte) {
	// This method will receive the transaction and add it to our actor state

}

// This will be executed at the end of every epoch
/*	use the map of timelock transactions and decrypt them possibly here? 
	base on the drand logic
	==> ( even null blocks )
*/
func (a Actor) CronTick(){
	// loop over all tx registered for the current epoch

	// use the modified runtime method to get the beacon randomness 
	// decrypt them and add to either runtime or state

}

// TODO state

type State struct {
	// map[abi.ChainEpoch]EncTxSet
	// Drand public key 

	// Public key needs to set up at lotus genesis somehow
}