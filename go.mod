module github.com/dominikbraun/graph

go 1.18

replace github.com/dominikbraun/graph/third_party/tencent.com/mmkv => ./third_party/tencent.com/mmkv

require (
	github.com/dominikbraun/graph/third_party/tencent.com/mmkv v0.0.0-00010101000000-000000000000
	github.com/smartystreets/goconvey v1.8.1
)

require (
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smarty/assertions v1.15.0 // indirect
)
