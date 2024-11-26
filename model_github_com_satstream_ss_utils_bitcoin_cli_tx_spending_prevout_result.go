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

type GithubComSatstreamSsUtilsBitcoinCliTxSpendingPrevoutResult struct {
	// The transaction id of the mempool transaction spending this output
	Spendingtxid string `json:"spendingtxid,omitempty"`
	// Whether the output is spent
	Spent bool `json:"spent,omitempty"`
	// The transaction id of the checked output
	Txid string `json:"txid,omitempty"`
	// The vout value of the checked output
	Vout int32 `json:"vout,omitempty"`
}
