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

type DecodedPsbt struct {
	// The transaction fee paid if all UTXOs slots are filled
	Fee float64 `json:"fee,omitempty"`
	// Array of inputs
	Inputs []DecodedPsbtInput `json:"inputs,omitempty"`
	// Array of outputs
	Outputs []DecodedPsbtOutput `json:"outputs,omitempty"`
	// The decoded network-serialized unsigned transaction
	Tx *AllOfDecodedPsbttx `json:"tx,omitempty"`
	// The unknown global fields
	Unknown *AllOfDecodedPsbtUnknown `json:"unknown,omitempty"`
}
