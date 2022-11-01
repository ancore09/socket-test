module socket-test

go 1.19

require (
	github.com/ambelovsky/gosf v0.0.0-20201109201340-237aea4d6109
	test/kafka v0.0.0-00010101000000-000000000000
)

require (
	github.com/ambelovsky/go-structs v1.1.0 // indirect
	github.com/ambelovsky/gosf-socketio v0.0.0-20220810204405-0f97832ec7af // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.36 // indirect

)

replace test/kafka => ../service
