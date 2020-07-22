/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Uri struct {
	Absolute bool `json:"absolute,omitempty"`

	Authority string `json:"authority,omitempty"`

	Fragment string `json:"fragment,omitempty"`

	Host string `json:"host,omitempty"`

	Opaque bool `json:"opaque,omitempty"`

	Path string `json:"path,omitempty"`

	Port int32 `json:"port,omitempty"`

	Query string `json:"query,omitempty"`

	RawAuthority string `json:"rawAuthority,omitempty"`

	RawFragment string `json:"rawFragment,omitempty"`

	RawPath string `json:"rawPath,omitempty"`

	RawQuery string `json:"rawQuery,omitempty"`

	RawSchemeSpecificPart string `json:"rawSchemeSpecificPart,omitempty"`

	RawUserInfo string `json:"rawUserInfo,omitempty"`

	Scheme string `json:"scheme,omitempty"`

	SchemeSpecificPart string `json:"schemeSpecificPart,omitempty"`

	UserInfo string `json:"userInfo,omitempty"`
}
