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

type ResponsesGetAddressTimeframeBalanceItem struct {
	Balance string `json:"balance,omitempty"`
	// Time       time.Time  `json:\"time\"`
	BlockRange *AllOfresponsesGetAddressTimeframeBalanceItemBlockRange `json:"blockRange,omitempty"`
}
