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

type GithubComSatstreamSsUtilsOrdServerResponsesBlockResponse struct {
	BestHeight int32 `json:"best_height,omitempty"`
	Hash string `json:"hash,omitempty"`
	Height int32 `json:"height,omitempty"`
	Inscriptions []string `json:"inscriptions,omitempty"`
	Runes []string `json:"runes,omitempty"`
	Target string `json:"target,omitempty"`
	Transactions []GithubComSatstreamSsUtilsOrdServerResponsesTransaction `json:"transactions,omitempty"`
}
