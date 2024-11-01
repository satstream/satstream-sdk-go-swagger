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

type RpcVout struct {
	N int32 `json:"n,omitempty"`
	RuneHoldings []RpcUtxoRune `json:"rune_holdings,omitempty"`
	ScriptPubKey *RpcScriptPubKey `json:"scriptPubKey,omitempty"`
	Value float32 `json:"value,omitempty"`
}
