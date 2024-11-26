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

type GithubComSatstreamSsApiServerApiTransactionRequestsCreatePsbtInput struct {
	// The sequence number Optional, default depends on the value of the 'replaceable' and 'locktime' arguments
	Sequence int32 `json:"sequence,omitempty"`
	// The transaction id
	Txid string `json:"txid"`
	// The output number
	Vout int32 `json:"vout"`
}
