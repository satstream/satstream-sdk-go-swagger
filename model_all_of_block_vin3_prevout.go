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

// Detailed previous output
type AllOfBlockVin3Prevout struct {
	Generated bool `json:"generated,omitempty"`
	Height int32 `json:"height,omitempty"`
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey,omitempty"`
	Value float64 `json:"value,omitempty"`
}
