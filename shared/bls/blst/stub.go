// +build darwin windows !blst_enabled

package blst

import (
	"github.com/prysmaticlabs/prysm/shared/bls/common"
)

// This stub file exists until build issues can be resolved for darwin and windows.
const err = "blst is only supported on linux with blst_enabled gotag"

// SecretKey -- stub
type SecretKey struct{}

// PublicKey -- stub
func (s SecretKey) PublicKey() common.PublicKey {
	panic(err)
}

// Sign -- stub
func (s SecretKey) Sign(_ []byte) common.Signature {
	panic(err)
}

// Marshal -- stub
func (s SecretKey) Marshal() []byte {
	panic(err)
}

// IsZero -- stub
func (s SecretKey) IsZero() bool {
	panic(err)
}

// PublicKey -- stub
type PublicKey struct{}

// Marshal -- stub
func (p PublicKey) Marshal() []byte {
	panic(err)
}

// Copy -- stub
func (p PublicKey) Copy() common.PublicKey {
	panic(err)
}

// Aggregate -- stub
func (p PublicKey) Aggregate(_ common.PublicKey) common.PublicKey {
	panic(err)
}

// IsInfinite -- stub
func (s PublicKey) IsInfinite() bool {
	panic(err)
}

// Signature -- stub
type Signature struct{}

// Verify -- stub
func (s Signature) Verify(_ common.PublicKey, _ []byte) bool {
	panic(err)
}

// AggregateVerify -- stub
func (s Signature) AggregateVerify(_ []common.PublicKey, _ [][32]byte) bool {
	panic(err)
}

// FastAggregateVerify -- stub
func (s Signature) FastAggregateVerify(_ []common.PublicKey, _ [32]byte) bool {
	panic(err)
}

// Marshal -- stub
func (s Signature) Marshal() []byte {
	panic(err)
}

// Copy -- stub
func (s Signature) Copy() common.Signature {
	panic(err)
}

// SecretKeyFromBytes -- stub
func SecretKeyFromBytes(_ []byte) (SecretKey, error) {
	panic(err)
}

// PublicKeyFromBytes -- stub
func PublicKeyFromBytes(_ []byte) (PublicKey, error) {
	panic(err)
}

// SignatureFromBytes -- stub
func SignatureFromBytes(_ []byte) (Signature, error) {
	panic(err)
}

// AggregatePublicKeys -- stub
func AggregatePublicKeys(_ [][]byte) (PublicKey, error) {
	panic(err)
}

// AggregateSignatures -- stub
func AggregateSignatures(_ []common.Signature) common.Signature {
	panic(err)
}

// VerifyMultipleSignatures -- stub
func VerifyMultipleSignatures(_ [][]byte, _ [][32]byte, _ []common.PublicKey) (bool, error) {
	panic(err)
}

// NewAggregateSignature -- stub
func NewAggregateSignature() common.Signature {
	panic(err)
}

// RandKey -- stub
func RandKey() (common.SecretKey, error) {
	panic(err)
}

// VerifyCompressed -- stub
func VerifyCompressed(_, _, _ []byte) bool {
	panic(err)
}
