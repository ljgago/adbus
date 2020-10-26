# AdBus - Advertisement Bus

It is an implementation idea of digital signage that uses a Rasperry Pi like media player.

The system is divided into 3 parts:

- The device: is the code that runs on the device. (Not yet implemented).
- The server: it is a server gateway, handle request from the Control Panel and sends the action to corresponding device. (In progess)
- The Control Panel: is the user interface for device administration. (Not yet implemented)

### For testing

No cache for testing
`go test internal/db/* -v -count=1`

