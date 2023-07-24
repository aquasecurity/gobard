module github.com/aquasecurity/askamanpage

go 1.19

require (
	github.com/aquasecurity/gobard v0.0.0-20230723051236-42c43ca83672
	github.com/go-resty/resty/v2 v2.7.0
	github.com/urfave/cli/v2 v2.25.7
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/net v0.10.0 // indirect
	k8s.io/apimachinery v0.27.4 // indirect
)

replace github.com/aquasecurity/gobard => ../../
