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

type InscriptionResponse struct {
	Address string `json:"address,omitempty"`
	Charms []string `json:"charms,omitempty"`
	Children []string `json:"children,omitempty"`
	ContentLength int32 `json:"content_length,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	EffectiveContentType string `json:"effective_content_type,omitempty"`
	Fee int32 `json:"fee,omitempty"`
	Height int32 `json:"height,omitempty"`
	Id string `json:"id,omitempty"`
	Next string `json:"next,omitempty"`
	Number int32 `json:"number,omitempty"`
	Parents []string `json:"parents,omitempty"`
	Previous string `json:"previous,omitempty"`
	Rune_ string `json:"rune,omitempty"`
	Sat int32 `json:"sat,omitempty"`
	Satpoint string `json:"satpoint,omitempty"`
	Timestamp int32 `json:"timestamp,omitempty"`
	Value int32 `json:"value,omitempty"`
}
