/*
 * Satstream API
 *
 * Satstream API
 *
 * API version: 1.0
 * Contact: team@satstream.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GithubComSatstreamSsUtilsRpcVin struct {
	Coinbase string `json:"coinbase,omitempty"`
	Prevout *GithubComSatstreamSsUtilsRpcPrevOut `json:"prevout,omitempty"`
	ScriptSig *GithubComSatstreamSsUtilsRpcScriptSig `json:"scriptSig,omitempty"`
	Sequence int32 `json:"sequence,omitempty"`
	Txid string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness,omitempty"`
	Vout int32 `json:"vout,omitempty"`
}
