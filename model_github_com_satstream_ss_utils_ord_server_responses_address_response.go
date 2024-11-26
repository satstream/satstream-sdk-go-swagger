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

type GithubComSatstreamSsUtilsOrdServerResponsesAddressResponse struct {
	Inscriptions []string `json:"inscriptions,omitempty"`
	Outputs []string `json:"outputs,omitempty"`
	RunesBalances []GithubComSatstreamSsUtilsOrdServerResponsesRunesBalance `json:"runes_balances,omitempty"`
	SatBalance int32 `json:"sat_balance,omitempty"`
}
