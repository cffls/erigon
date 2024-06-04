package l1_data

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/crypto"
	"github.com/ledgerwatch/erigon/zk/contracts"
	"github.com/ledgerwatch/erigon/zk/da"
)

type RollupBaseEtrogBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

type ValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

func DecodeL1BatchData(txData []byte, daUrl string) ([][]byte, common.Address, error) {
	// we need to know which version of the ABI to use here so lets find it
	idAsString := fmt.Sprintf("%x", txData[:4])
	abiMapped, found := contracts.SequenceBatchesMapping[idAsString]
	if !found {
		return nil, common.Address{}, fmt.Errorf("unknown l1 call data")
	}

	smcAbi, err := abi.JSON(strings.NewReader(abiMapped))
	if err != nil {
		return nil, common.Address{}, err
	}

	method, err := smcAbi.MethodById(txData[:4])
	if err != nil {
		return nil, common.Address{}, err
	}

	// Unpack method inputs
	data, err := method.Inputs.Unpack(txData[4:])
	if err != nil {
		return nil, common.Address{}, err
	}

	var coinbase common.Address

	isValidium := false

	switch idAsString {
	case contracts.SequenceBatchesIdv5_0:
		cb, ok := data[1].(common.Address)
		if !ok {
			return nil, common.Address{}, fmt.Errorf("expected position 1 in the l1 call data to be address")
		}
		coinbase = cb
	case contracts.SequenceBatchesIdv6_6:
		cb, ok := data[3].(common.Address)
		if !ok {
			return nil, common.Address{}, fmt.Errorf("expected position 3 in the l1 call data to be address")
		}
		coinbase = cb
	case contracts.SequenceBatchesValidiumElderBerry:
		if daUrl == "" {
			return nil, common.Address{}, fmt.Errorf("data availability url is required for validium")
		}
		isValidium = true
		cb, ok := data[3].(common.Address)
		if !ok {
			return nil, common.Address{}, fmt.Errorf("expected position 3 in the l1 call data to be address")
		}
		coinbase = cb
	default:
		return nil, common.Address{}, fmt.Errorf("unknown l1 call data")
	}

	var sequences []RollupBaseEtrogBatchData

	bytedata, err := json.Marshal(data[0])
	if err != nil {
		return nil, coinbase, err
	}

	if !isValidium {
		err = json.Unmarshal(bytedata, &sequences)
		if err != nil {
			return nil, coinbase, err
		}
	} else {
		var validiumSequences []ValidiumBatchData
		err = json.Unmarshal(bytedata, &validiumSequences)

		if err != nil {
			return nil, coinbase, err
		}

		for _, validiumSequence := range validiumSequences {
			hash := common.BytesToHash(validiumSequence.TransactionsHash[:])
			data, err := da.GetOffChainData(context.Background(), daUrl, hash)
			if err != nil {
				return nil, coinbase, err
			}

			actualTransactionsHash := crypto.Keccak256Hash(data)
			if actualTransactionsHash != hash {
				return nil, coinbase, fmt.Errorf("unable to fetch off chain data for hash %s", hash.String())
			}

			sequences = append(sequences, RollupBaseEtrogBatchData{
				Transactions:         data,
				ForcedGlobalExitRoot: validiumSequence.ForcedGlobalExitRoot,
				ForcedTimestamp:      validiumSequence.ForcedTimestamp,
				ForcedBlockHashL1:    validiumSequence.ForcedBlockHashL1,
			})
		}
	}

	batchL2Datas := make([][]byte, len(sequences))
	for idx, sequence := range sequences {
		batchL2Datas[idx] = sequence.Transactions
	}

	return batchL2Datas, coinbase, err
}
