module github.com/inexio/eve-ng-restapi-go-client

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
