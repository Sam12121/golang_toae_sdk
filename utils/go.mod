module github.com/Sam12121/golang_toae_sdk/utils

go 1.20

replace github.com/Sam12121/golang_toae_sdk/client => ../client

require (
	github.com/Sam12121/golang_toae_sdk/client v0.0.0-00010101000000-000000000000
	github.com/hashicorp/go-retryablehttp v0.7.5
)

require github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
