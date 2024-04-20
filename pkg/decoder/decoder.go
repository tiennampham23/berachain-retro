package decoder

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	ErrorNoMethodWithID = errors.New("no method with id")
)

func Decode(abiStr string, data []byte) (interface{}, error) {
	input := hexutil.Encode(data)
	// ignore if the input is smaller than 4bytes
	if abiStr == "" {
		return nil, nil
	}

	ABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}
	inputBytes := common.FromHex(input)
	method, err := ABI.MethodById(inputBytes)
	if err != nil {
		if errors.Is(err, ErrorNoMethodWithID) {
			return nil, nil
		}
		return nil, err
	}

	bytes, err := hex.DecodeString(input[10:])
	if err != nil {
		return nil, err
	}

	inputs, err := method.Inputs.Unpack(bytes)
	if err != nil {
		return nil, err
	}

	nonIndexedArgs := method.Inputs.NonIndexed()

	contractCall := map[string]interface{}{
		"name": method.Name,
	}

	params := make([]interface{}, 0)

	for i, input := range inputs {
		arg := nonIndexedArgs[i]
		param := map[string]interface{}{
			"name":  arg.Name,
			"value": input,
			"type":  arg.Type.String(),
		}
		params = append(params, param)
	}

	contractCall["params"] = params

	return contractCall, nil
}
