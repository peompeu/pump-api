package domain

type OauthProviderType string

const (
	KaKao OauthProviderType = "kakao"
	Apple OauthProviderType = "apple"
)

type RoleType string

const (
	Admin   RoleType = "admin"
	General RoleType = "general"
)

type User struct {
	Email    *string
	Password *string

	OauthProvider *OauthProviderType
	Role          RoleType
}
