package actors

import (
	"github.com/filecoin-project/lotus/chain/address"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

var AccountActorCodeCid cid.Cid
var StorageMarketActorCodeCid cid.Cid
var StorageMinerCodeCid cid.Cid
var MultisigActorCodeCid cid.Cid
var InitActorCodeCid cid.Cid
var PaymentChannelActorCodeCid cid.Cid

var InitActorAddress = mustIDAddress(0)
var NetworkAddress = mustIDAddress(1)
var StorageMarketAddress = mustIDAddress(2)
var BurntFundsAddress = mustIDAddress(99)

func mustIDAddress(i uint64) address.Address {
	a, err := address.NewIDAddress(i)
	if err != nil {
		panic(err)
	}
	return a
}

func init() {
	pref := cid.NewPrefixV1(cid.Raw, mh.ID)
	mustSum := func(s string) cid.Cid {
		c, err := pref.Sum([]byte(s))
		if err != nil {
			panic(err)
		}
		return c
	}

	AccountActorCodeCid = mustSum("account")
	StorageMarketActorCodeCid = mustSum("smarket")
	StorageMinerCodeCid = mustSum("sminer")
	MultisigActorCodeCid = mustSum("multisig")
	InitActorCodeCid = mustSum("init")
	PaymentChannelActorCodeCid = mustSum("paych")
}
