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

type RequestsJoinPsbtsRequest struct {
	// Array of base64-encoded PSBTs to join
	Psbts []string `json:"psbts,omitempty"`
}
