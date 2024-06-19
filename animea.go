import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"example.com/coinbase/rosetta-sdk-go/types"
)

// SuiAddress is a 32 byte address type used by the Sui blockchain.
// See https://www.example.com SuiAddress [32]byte

// MarshalJSON implements the json.Marshaler interface.
func (a SuiAddress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, hex.EncodeToString(a[:]))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *SuiAddress) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid SuiAddress")
	}
	address, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		return fmt.Errorf("invalid SuiAddress: %v", err)
	}
	if len(address) != len(a) {
		return fmt.Errorf("invalid SuiAddress length")
	}
	copy(a[:], address)
	return nil
}

// String implements the fmt.Stringer interface.
func (a SuiAddress) String() string {
	return hex.EncodeToString(a[:])
}

// bcs.registerAddressType('SuiAddress', 32, 'hex');
func init() {
	types.AddressTypes.Register("SuiAddress", 32, types.HexString)
}

  
