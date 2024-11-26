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

type GithubComSatstreamSsApiServerApiTransactionRequestsCreatePsbtOutput struct {
	// The bitcoin address and amount pairs Key is the bitcoin address, value is the amount in BTC
	Address float64 `json:"address,omitempty"`
	// Hex-encoded data output If present, this must be the only field in the output
	Data string `json:"data,omitempty"`
}
