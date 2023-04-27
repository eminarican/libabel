package page

import (
	"crypto/md5"
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"io"
	"math/big"
	"math/rand"
	"strings"
)

func fromBabel(value string) *big.Int {
	return fromArbitraryBase(value, babelSet)
}

func toBabel(value *big.Int) string {
	return toArbitraryBase(value, babelSet)
}

func fromArbitraryBase(value string, set []byte) *big.Int {
	result := new(big.Int)
	base := big.NewInt(int64(len(set)))

	for _, bn := range value {
		val := big.NewInt(int64(strings.IndexRune(string(set), bn)))
		result = result.Mul(result, base).Add(result, val)
	}

	return result
}

func toArbitraryBase(value *big.Int, set []byte) string {
	if value.Sign() < 0 {
		value = value.Neg(value)
	}

	base := big.NewInt(int64(len(set)))

	arb := make([]byte, 0, 4096)

	zero := new(big.Int)
	for value.Cmp(zero) != 0 {
		newVal, rem := new(big.Int).DivMod(value, base, new(big.Int))
		arb = append(arb, set[rem.Uint64()])

		value = newVal
	}

	for i, j := 0, len(arb)-1; i < j; i, j = i+1, j-1 {
		arb[i], arb[j] = arb[j], arb[i]
	}

	return string(arb)
}

func locationHash(room cube.Pos, shulker, volume, page int) *big.Int {
	h := md5.New()
	_, _ = io.WriteString(h, fmt.Sprintf("%v:%v:%v:%v", room, shulker, volume, page))
	return new(big.Int).SetBytes(h.Sum(nil))
}

func randomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = babelSet[rand.Intn(len(babelSet))]
	}
	return string(result)
}
