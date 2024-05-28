package state

import (
	"github.com/ledgerwatch/erigon-lib/common"
)

type PreimageWriter struct {
	preImageMap   map[string][]byte
	savePreimages bool
}

func (pw *PreimageWriter) SetSavePreimages(save bool) {
	pw.savePreimages = save
}

func (pw *PreimageWriter) HashAddress(address common.Address, save bool) (common.Hash, error) {
	addrHash, err := common.HashData(address[:])
	if err != nil {
		return common.Hash{}, err
	}
	return addrHash, pw.savePreimage(save, addrHash[:], address[:])
}

func (pw *PreimageWriter) HashKey(key *common.Hash, save bool) (common.Hash, error) {
	keyHash, err := common.HashData(key[:])
	if err != nil {
		return common.Hash{}, err
	}
	return keyHash, pw.savePreimage(save, keyHash[:], key[:])
}

func (pw *PreimageWriter) savePreimage(save bool, hash []byte, preimage []byte) error {
	if !pw.savePreimages {
		return nil
	}

	if pw.preImageMap == nil {
		pw.preImageMap = make(map[string][]byte)
	}

	if _, ok := pw.preImageMap[string(hash)]; !ok {
		pw.preImageMap[string(hash)] = preimage
	}

	return nil
}

func (pw *PreimageWriter) GetPreimage(hash common.Hash) []byte {
	if pw.preImageMap == nil {
		return nil
	}
	return pw.preImageMap[string(hash[:])]
}
