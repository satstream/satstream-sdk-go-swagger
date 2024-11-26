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

type MempoolFees struct {
	// Ancestor transaction fees in BTC
	Ancestor float64 `json:"ancestor,omitempty"`
	// Base transaction fee in BTC
	Base float64 `json:"base,omitempty"`
	// Descendant transaction fees in BTC
	Descendant float64 `json:"descendant,omitempty"`
	// Modified transaction fee in BTC
	Modified float64 `json:"modified,omitempty"`
}
