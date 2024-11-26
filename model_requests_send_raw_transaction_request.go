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

type RequestsSendRawTransactionRequest struct {
	// The hex string of the raw transaction
	HexString string `json:"hex_string,omitempty"`
	// Optional: Reject transactions whose fee rate is higher than this value (BTC/kvB)
	MaxFeeRate float64 `json:"max_fee_rate,omitempty"`
}
