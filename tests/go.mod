module github.com/rl-lasvsim/openapi-sdk-go/tests

go 1.19

require (
	github.com/rl-lasvsim/openapi-sdk-go/lasvsim v1.9999.9999
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/rl-lasvsim/openapi-sdk-go/lasvsim v1.9999.9999 => ../lasvsim
