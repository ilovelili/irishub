package simulation

// DONTCOVER

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/irisnet/irishub/modules/asset/internal/keeper"
	"github.com/irisnet/irishub/modules/asset/internal/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding gov type
func DecodeStore(cdc *codec.Codec, kvA, kvB cmn.KVPair) string {
	switch {
	case strings.HasPrefix(string(kvA.Key), string(keeper.PrefixToken)):
		var tokenA, tokenB types.Tokens
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &tokenA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &tokenB)
		return fmt.Sprintf("%v\n%v", tokenA, tokenB)

	default:
		panic(fmt.Sprintf("invalid asset key %X", kvA.Key))
	}
}
