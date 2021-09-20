module es

go 1.16

require (
	github.com/olivere/elastic v6.2.37+incompatible
	github.com/olivere/elastic/v7 v7.0.29
	microframe.com/logger v0.0.0-00010101000000-000000000000
)

replace microframe.com/logger => ../logger
