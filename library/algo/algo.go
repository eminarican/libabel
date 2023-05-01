package algo

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"math/big"
	"math/rand"
)

const (
	Length = 200
)

var (
	babelSet = []byte{
		' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', ',', '.',
	}

	base64Set = []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_',
	}
)

type Address struct {
	Hex     string
	Room    cube.Pos
	Shulker int
	Volume  int
	Page    int
}

func (a Address) String() string {
	return fmt.Sprintf(
		"%v:[%v]:[%v]:[%v]:%v",
		a.Room, a.Shulker, a.Volume, a.Page, a.Hex,
	)
}

func (a Address) Format() string {
	return text.Colourf(
		"<green>Room:</green> %v <purple>Shulker:</purple> %v\n"+
			"<aqua>Book:</aqua> %v <red>Page:</red> %v\n"+
			"<yellow>Hex:</yellow> %v",
		a.Room[:], a.Shulker, a.Volume, a.Page, a.Hex,
	)
}

func Search(value string, fill bool) Address {
	if len(value) < Length && fill {
		mis := Length - len(value)
		preLen := rand.Intn(mis + 1)
		sufLen := mis - preLen
		value = randomString(preLen) + value + randomString(sufLen)
	}

	room := cube.Pos{
		rand.Intn(1000),
		rand.Intn(14) + 1,
		rand.Intn(1000),
	}
	shulker := rand.Intn(10) + 1
	volume := rand.Intn(10) + 1
	page := rand.Intn(50) + 1

	loc := locationHash(room, shulker, volume, page)

	mul := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(5*Length)), nil)

	hexAddr := toArbitraryBase(new(big.Int).Add(fromBabel(value), new(big.Int).Mul(loc, mul)), base64Set)

	return Address{
		Hex:     hexAddr,
		Room:    room,
		Shulker: shulker,
		Volume:  volume,
		Page:    page,
	}
}

func Read(addr Address) string {
	loc := locationHash(addr.Room, addr.Shulker, addr.Volume, addr.Page)

	mul := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(5*Length)), nil)

	txt := toBabel(new(big.Int).Sub(fromArbitraryBase(addr.Hex, base64Set), new(big.Int).Mul(loc, mul)))

	if len(txt) > Length {
		return txt[:Length]
	}
	return txt
}
