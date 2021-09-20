package valuerenderer

import (
	"context"
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/dustin/go-humanize"

	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// TODO
// 1.Look at some library wqith a lot of stars that formats nicely integers and floats.
// 2.If there is no good solution for step1, implement own solution.

// ValueRenderer defines an interface to produce formated output for Int,Dec,Coin types as well as parse a string to Coin or Uint.
type ValueRenderer interface {
	Format(context.Context, interface{}) (string, error)
	Parse(context.Context, string) (interface{}, error)
}

// denomQuerierFunc takes a context and a denom as arguments and returns metadata, error
type denomQuerierFunc func(context.Context, string) (banktypes.Metadata, error)

// DefaultValueRenderer defines a struct that implements ValueRenderer interface
type DefaultValueRenderer struct {
	denomQuerier denomQuerierFunc
}

var _ ValueRenderer = &DefaultValueRenderer{}

// NewDefaultValueRenderer  initiates denomToMetadataMap field and returns DefaultValueRenderer struct
func NewDefaultValueRenderer(denomQuerier denomQuerierFunc) DefaultValueRenderer {
	return DefaultValueRenderer{
		denomQuerier: denomQuerier,
	}
}

// Format converts an empty interface into a string depending on interface type.
func (dvr DefaultValueRenderer) Format(c context.Context, x interface{}) (string, error) {
	var sb strings.Builder

	switch v := x.(type) {
	case types.Dec:
		s := v.String()
		if len(s) == 0 {
			return "", errors.New("empty string")
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return "", errors.New("empty string")
		}

		sb.WriteString(humanize.Ftoa(f))

	case types.Int:
		s := v.String()
		if len(s) == 0 {
			return "", errors.New("empty string")
		}

		sb.WriteString(humanize.Comma(v.Int64()))

	case types.Coin:
		metadata, err := dvr.denomQuerier(c, convertToBaseDenom(v.Denom))
		if err != nil {
			return "", err
		}

		expSub := computeExponentSubtraction(v.Denom, metadata)

		formatedAmount := dvr.ComputeAmount(v.Amount.Int64(), expSub)

		sb.WriteString(formatedAmount)
		sb.WriteString(metadata.Display)

	default:
		panic("type is invalid")
	}

	return sb.String(), nil
}

func convertToBaseDenom(denom string) string {
	switch {
	// e.g. uregen => uregen
	case strings.HasPrefix(denom, "u"):
		return denom
	// e.g. mregen => uregen
	case strings.HasPrefix(denom, "m"):
		return "u" + denom[1:]
	// has no prefix regen => uregen
	default:
		return "u" + denom
	}
}

func computeExponentSubtraction(denom string, metadata banktypes.Metadata) float64 {
	var coinExp, displayExp int64
	for _, denomUnit := range metadata.DenomUnits {
		if denomUnit.Denom == denom {
			coinExp = int64(denomUnit.Exponent)
		}

		if denomUnit.Denom == metadata.Display {
			displayExp = int64(denomUnit.Exponent)
		}
	}

	return float64(coinExp - displayExp)
}

// countTrailingZeroes counts the amount of trailing zeroes in a string
func countTrailingZeroes(str string) int {
	counter := 0
	for i := len(str) - 1; i > 0; i-- {
		if rune(str[i]) == rune('0') {
			counter++
		} else {
			break
		}
	}
	return counter
}

// ComputeAmount calculates an amount to produce formated output
func (dvr DefaultValueRenderer) ComputeAmount(amount int64, expSub float64) string {

	// check if amount is nil
	// check if expSub is zero
	var res string

	switch {
	// negative , convert mregen to regen less zeroes 23 => 0,023, expSub -3
	case math.Signbit(expSub):

		stringValue := strconv.FormatInt(amount, 10)
		count := countTrailingZeroes(stringValue)
		if count >= int(math.Abs(expSub)) {
			// case 1 if number of zeroes >= Abs(expSub)  23000, 3 => 23 (int64)
			x := amount / int64(math.Pow(10, math.Abs(expSub)))
			res = humanize.Comma(x)
		} else {
			// case 2 number of trailing zeroes < abs(expSub)  23, 3,=> 0.023(float64)
			x := float64(float64(amount) / math.Pow(10, math.Abs(expSub)))
			res = humanize.Ftoa(x)
		}
	// positive, e.g.convert mregen to uregen
	case !math.Signbit(expSub):
		x := amount * int64(math.Pow(10, expSub))
		res = humanize.Comma(x)
	// == 0, convert regen to regen, amount does not change
	default:
		res = humanize.Comma(amount)
	}

	return res

}

// Parse parses a string and takes a decision whether to convert it into Coin or Uint
func (dvr DefaultValueRenderer) Parse(ctx context.Context, s string) (interface{}, error) {
	if s == "" {
		return nil, errors.New("unable to parse empty string")
	}
	// remove all commas
	str := strings.ReplaceAll(s, ",", "")
	re := regexp.MustCompile(`\d+[mu]?regen`) // TODO remove regen
	// case 1: "1000000regen" => Coin
	if re.MatchString(str) {
		coin, err := coinFromString(str)
		if err != nil {
			return nil, err
		}

		return coin, nil
	}

	// case2: convert it to Uint
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil, err
	}

	return types.NewUint(i), nil
}

// should I add tests for coinFromString?
// coinFromString converts a string to coin
// TODO use regex below intead of this algo
func coinFromString(s string) (types.Coin, error) {
	index := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			continue
		}

		index = i
		break
	}

	if index == len(s)-1 {
		return types.Coin{}, errors.New("unable to find a denonination")
	}

	amount, denom := s[:index+1], s[index+1:]
	// convert to int64 to make up Coin later
	amountInt, ok := types.NewIntFromString(amount)
	if !ok {
		return types.Coin{}, errors.New("unable to convert amountStr into int64")
	}

	return types.NewCoin(denom, amountInt), nil
}
