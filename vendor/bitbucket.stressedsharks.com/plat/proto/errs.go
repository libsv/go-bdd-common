package proto

import (
	"github.com/pkg/errors"
)

var (
	// ErrNegativeBalance balance cannot be negative
	ErrNegativeBalance = errors.New("balance cannot be negative")
	// ErrEmptyTxid no txid
	ErrEmptyTxid = errors.New("Txid field cannot be empty")
	// ErrEmptySpendingTxid no spendingTxid
	ErrEmptySpendingTxid = errors.New("SpendingTxid field cannot be empty")
	// ErrEmptyData no data
	ErrEmptyData = errors.New("data cannot be empty")
	// ErrEmptyTimestamp no timestamp
	ErrEmptyTimestamp = errors.New("timestamp cannot be empty")
	// ErrEmptyLastTimestamp no lastTimestamp
	ErrEmptyLastTimestamp = errors.New("lastTimestamp cannot be empty")
	// ErrEmptyTransactionData no transaction data
	ErrEmptyTransactionData = errors.New("transaction data field cannot be empty")
	// ErrEmptyExtendedPrivateKey no extendedPrivateKey
	ErrEmptyExtendedPrivateKey = errors.New("ExtendedPrivateKey field cannot be empty")
	// ErrTxidInvalidLength txid incorrect length
	ErrTxidInvalidLength = errors.New("Txid must have a length of 64 characters (32 bytes)")
	// ErrSpendingTxidInvalidLength spendingTxid incorrect length
	ErrSpendingTxidInvalidLength = errors.New("SpendingTxid must have a length of 64 characters (32 bytes)")
	// ErrExtendedPrivateKeyInvalidLength ExtendedPrivateKey incorrect length
	ErrExtendedPrivateKeyInvalidLength = errors.New("ExtendedPrivateKey must have a length of 111 characters")
	// ErrTxidInvalidHex invalid txid hex
	ErrTxidInvalidHex = errors.New("Txid field must be valid hex")
	// ErrHashInvalidLength hash incorrect length
	ErrHashInvalidLength = errors.New("hash message must have a length of 32 bytes")
	// ErrIndexLessThanZero invalid index
	ErrIndexLessThanZero = errors.New("Index must be greater than or equal to 0")
	// ErrBlockHeightLessThanZero invalid blockHeight
	ErrBlockHeightLessThanZero = errors.New("BlockHeight must be greater than or equal to 0")
	// ErrKeyContextMissing keyContext is a required parameter
	ErrKeyContextMissing = errors.New("keyContext param missing")
	// ErrEncryptNoData no data
	ErrEncryptNoData = errors.New("no data to encrypt")
	// ErrDecryptNoData no data
	ErrDecryptNoData = errors.New("no data to decrypt")
	// ErrEmptyPublicKey empty public key
	ErrEmptyPublicKey = errors.New("empty public key")
	// ErrEmptySignature empty signature
	ErrEmptySignature = errors.New("empty signature")
	// ErrEmptyLockingSript empty lockingScript
	ErrEmptyLockingSript = errors.New("empty lockingScript")
	// ErrEmptySatoshis invalid satoshi amount
	ErrEmptySatoshis = errors.New("satoshis must be greater than 0")
	// ErrEmptyServiceName empty serviceName
	ErrEmptyServiceName = errors.New("empty ServiceName")
	// ErrServiceNameLength invalid serviceName length
	ErrServiceNameLength = errors.New("ServiceName field must be less than or equal to 255 characters")
	// ErrEmptySigner empty signer
	ErrEmptySigner = errors.New("empty signer")
	// ErrAddressFormat invalid address
	ErrAddressFormat = errors.New("invalid Address format (see signing.nchain.com:9000)")
	// ErrEmailFormat invalid email
	ErrEmailFormat = errors.New("invalid email format")
	// ErrPhoneNumberFormat invalid phone number
	ErrPhoneNumberFormat = errors.New("invalid phoneNumber format, does not follow E.164 standard")
	// ErrEmptyAlias empty alias
	ErrEmptyAlias = errors.New("empty signer alias")
	// ErrAliasLength invalid alias length
	ErrAliasLength = errors.New("alias field must be less than or equal to 255 characters")
	// ErrAliasFormat invalid alias format
	ErrAliasFormat = errors.New("alias must only contain 0-9, a-z, A-Z, _ and -")
	// ErrEmptyFunds empty funds
	ErrEmptyFunds = errors.New("empty funds")
	// ErrInvalidStartingPath invalid start path
	ErrInvalidStartingPath = errors.New("paths must not start with /")
	// ErrInvalidEndingPath invalid end path
	ErrInvalidEndingPath = errors.New("paths must not end with /")
	// ErrInvalidPathFormat invalid path format
	ErrInvalidPathFormat = errors.New("path format must only contain 0-9, a-z, and A-Z separated by /")
	// ErrInvalidDataWriterData invalid data format
	ErrInvalidDataWriterData = errors.New(`invalid data format, must be follow this format: { "data": [ [65, 66], [67, 68] ] }`)
)
