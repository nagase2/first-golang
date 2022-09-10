module example/hello

go 1.19

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace example.com/greetings => ../greetings
