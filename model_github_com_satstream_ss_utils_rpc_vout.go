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

type GithubComSatstreamSsUtilsRpcVout struct {
	N int32 `json:"n,omitempty"`
	RuneHoldings []GithubComSatstreamSsUtilsRpcUtxoRune `json:"rune_holdings,omitempty"`
	ScriptPubKey *GithubComSatstreamSsUtilsRpcScriptPubKey `json:"scriptPubKey,omitempty"`
	Value float32 `json:"value,omitempty"`
}
