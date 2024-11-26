/*
 * Satstream API
 *
 * Satstream API
 *
 * API version: 1.0
 * Contact: team@satstream.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package satstream_go_sdk

// Block statistics
type AllOfresponsesGetBlockStatsResponseData struct {
	// Average fee in the block
	Avgfee float64 `json:"avgfee,omitempty"`
	// Average feerate (in satoshis per virtual byte)
	Avgfeerate float64 `json:"avgfeerate,omitempty"`
	// Average transaction size
	Avgtxsize float64 `json:"avgtxsize,omitempty"`
	// The block hash (to check for potential reorgs)
	Blockhash string `json:"blockhash,omitempty"`
	// Feerates at the 10th, 25th, 50th, 75th, and 90th percentile
	FeeratePercentiles []float64 `json:"feerate_percentiles,omitempty"`
	// The height of the block
	Height int32 `json:"height,omitempty"`
	// The number of inputs (excluding coinbase)
	Ins int32 `json:"ins,omitempty"`
	// Maximum fee in the block
	Maxfee float64 `json:"maxfee,omitempty"`
	// Maximum feerate (in satoshis per virtual byte)
	Maxfeerate float64 `json:"maxfeerate,omitempty"`
	// Maximum transaction size
	Maxtxsize int32 `json:"maxtxsize,omitempty"`
	// Truncated median fee in the block
	Medianfee float64 `json:"medianfee,omitempty"`
	// The block median time past
	Mediantime int32 `json:"mediantime,omitempty"`
	// Truncated median transaction size
	Mediantxsize int32 `json:"mediantxsize,omitempty"`
	// Minimum fee in the block
	Minfee float64 `json:"minfee,omitempty"`
	// Minimum feerate (in satoshis per virtual byte)
	Minfeerate float64 `json:"minfeerate,omitempty"`
	// Minimum transaction size
	Mintxsize int32 `json:"mintxsize,omitempty"`
	// The number of outputs
	Outs int32 `json:"outs,omitempty"`
	// The block subsidy
	Subsidy float64 `json:"subsidy,omitempty"`
	// Total size of all segwit transactions
	SwtotalSize int32 `json:"swtotal_size,omitempty"`
	// Total weight of all segwit transactions
	SwtotalWeight int32 `json:"swtotal_weight,omitempty"`
	// The number of segwit transactions
	Swtxs int32 `json:"swtxs,omitempty"`
	// The block time
	Time int32 `json:"time,omitempty"`
	// Total amount in all outputs
	TotalOut float64 `json:"total_out,omitempty"`
	// Total size of all non-coinbase transactions
	TotalSize int32 `json:"total_size,omitempty"`
	// Total weight of all non-coinbase transactions
	TotalWeight int32 `json:"total_weight,omitempty"`
	// The fee total
	Totalfee float64 `json:"totalfee,omitempty"`
	// The number of transactions (excluding coinbase)
	Txs int32 `json:"txs,omitempty"`
	// The increase/decrease in the number of unspent outputs
	UtxoIncrease int32 `json:"utxo_increase,omitempty"`
	// The increase/decrease in size for the utxo index
	UtxoSizeInc int32 `json:"utxo_size_inc,omitempty"`
}
