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

type GithubComSatstreamSsUtilsOrdServerResponsesBlocksResponse struct {
	Blocks []string `json:"blocks,omitempty"`
	FeaturedBlocks map[string][]string `json:"featured_blocks,omitempty"`
	Last int32 `json:"last,omitempty"`
}
