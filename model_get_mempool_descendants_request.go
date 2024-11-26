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

type GetMempoolDescendantsRequest struct {
	// Required: The transaction id (must be in mempool)
	Txid string `json:"txid,omitempty"`
	// Optional: True for detailed information, false for just txids
	Verbose bool `json:"verbose,omitempty"`
}