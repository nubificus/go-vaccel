### Initial go-vaccel bindings

Apparently, it is as easy as directly calling functions from Go.

```
# export VACCEL_DEBUG_LEVEL=4
# export VACCEL_BACKENDS=/usr/local/lib/libvaccel-noop.so
# make
go build -o bin/cgo cgo.go
# ./bin/cgo 
2023.08.09-07:00:50.80 - <debug> Initializing vAccel
2023.08.09-07:00:50.80 - <info> vAccel v0.5.0-20-g914649b
2023.08.09-07:00:50.80 - <debug> Created top-level rundir: /run/user/0/vaccel.vyhPlD
2023.08.09-07:00:50.80 - <info> Registered plugin noop v0.5.0-20-g914649b
2023.08.09-07:00:50.80 - <debug> Registered function noop from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function sgemm from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function image classification from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function image detection from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function image segmentation from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function image pose estimation from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function image depth estimation from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function exec from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function TensorFlow session load from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function TensorFlow session run from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function TensorFlow session delete from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function MinMax from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Array copy from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Vector Add from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Parallel acceleration from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Matrix multiplication from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Exec with resource from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Torch jitload_forward function from plugin noop
2023.08.09-07:00:50.80 - <debug> Registered function Torch SGEMM from plugin noop
2023.08.09-07:00:50.80 - <debug> Loaded plugin noop from /usr/local/lib/libvaccel-noop.so
-------------------------------
2023.08.09-07:00:50.80 - <debug> session:1 New session
2023.08.09-07:00:50.80 - <debug> session:1 Looking for plugin implementing noop
2023.08.09-07:00:50.80 - <debug> Found implementation in noop plugin
[noop] Calling no-op for session 1
2023.08.09-07:00:50.80 - <debug> session:1 Free session
```
