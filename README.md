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

and the Image classification example:

```
root@jetson5:/home/ananos/develop/go-vaccel# export VACCEL_IMAGENET_NETWORKS=/home/ananos/develop/jetson-inference/data/networks/
(failed reverse-i-search)`epo': cat /root/.bash_history |gr^C sysctl
root@jetson5:/home/ananos/develop/go-vaccel# export VACCEL_DEBUG_LEVEL=4
root@jetson5:/home/ananos/develop/go-vaccel# export VACCEL_BACKENDS=/usr/local/lib/libvaccel-jetson.so 
root@jetson5:/home/ananos/develop/go-vaccel# ./bin/cgo 
2023.08.09-08:29:29.64 - <debug> Initializing vAccel
2023.08.09-08:29:29.64 - <info> vAccel v0.5.0-13-g4001c6e-dirty
2023.08.09-08:29:29.64 - <debug> Created top-level rundir: /run/user/0/vaccel.xf0DdM
2023.08.09-08:29:29.71 - <info> Registered plugin jetson-inference 0.1
2023.08.09-08:29:29.71 - <debug> Registered function image classification from plugin jetson-inference
2023.08.09-08:29:29.71 - <debug> Registered function image detection from plugin jetson-inference
2023.08.09-08:29:29.71 - <debug> Registered function image segmentation from plugin jetson-inference
2023.08.09-08:29:29.71 - <debug> Registered function image pose estimation from plugin jetson-inference
2023.08.09-08:29:29.71 - <debug> Registered function image depth estimation from plugin jetson-inference
2023.08.09-08:29:29.71 - <debug> Loaded plugin jetson-inference from /usr/local/lib/libvaccel-jetson.so
-------------------------------
2023.08.09-08:29:29.72 - <debug> session:1 New session
2023.08.09-08:29:29.72 - <debug> session:1 Looking for plugin implementing image classification
2023.08.09-08:29:29.72 - <debug> Found implementation in jetson-inference plugin

imageNet -- loading classification network model from:
         -- prototxt     /home/ananos/develop/jetson-inference/data/networks//googlenet.prototxt
         -- model        /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel
         -- class_labels /home/ananos/develop/jetson-inference/data/networks//ilsvrc12_synset_words.txt
         -- input_blob   'data'
         -- output_blob  'prob'
         -- batch_size   1

[TRT]    TensorRT version 8.4.1
[TRT]    loading NVIDIA plugins...
[TRT]    Registered plugin creator - ::GridAnchor_TRT version 1
[TRT]    Registered plugin creator - ::GridAnchorRect_TRT version 1
[TRT]    Registered plugin creator - ::NMS_TRT version 1
[TRT]    Registered plugin creator - ::Reorg_TRT version 1
[TRT]    Registered plugin creator - ::Region_TRT version 1
[TRT]    Registered plugin creator - ::Clip_TRT version 1
[TRT]    Registered plugin creator - ::LReLU_TRT version 1
[TRT]    Registered plugin creator - ::PriorBox_TRT version 1
[TRT]    Registered plugin creator - ::Normalize_TRT version 1
[TRT]    Registered plugin creator - ::ScatterND version 1
[TRT]    Registered plugin creator - ::RPROI_TRT version 1
[TRT]    Registered plugin creator - ::BatchedNMS_TRT version 1
[TRT]    Registered plugin creator - ::BatchedNMSDynamic_TRT version 1
[TRT]    Registered plugin creator - ::BatchTilePlugin_TRT version 1
[TRT]    Could not register plugin creator -  ::FlattenConcat_TRT version 1
[TRT]    Registered plugin creator - ::CropAndResize version 1
[TRT]    Registered plugin creator - ::CropAndResizeDynamic version 1
[TRT]    Registered plugin creator - ::DetectionLayer_TRT version 1
[TRT]    Registered plugin creator - ::EfficientNMS_TRT version 1
[TRT]    Registered plugin creator - ::EfficientNMS_ONNX_TRT version 1
[TRT]    Registered plugin creator - ::EfficientNMS_Explicit_TF_TRT version 1
[TRT]    Registered plugin creator - ::EfficientNMS_Implicit_TF_TRT version 1
[TRT]    Registered plugin creator - ::ProposalDynamic version 1
[TRT]    Registered plugin creator - ::Proposal version 1
[TRT]    Registered plugin creator - ::ProposalLayer_TRT version 1
[TRT]    Registered plugin creator - ::PyramidROIAlign_TRT version 1
[TRT]    Registered plugin creator - ::ResizeNearest_TRT version 1
[TRT]    Registered plugin creator - ::Split version 1
[TRT]    Registered plugin creator - ::SpecialSlice_TRT version 1
[TRT]    Registered plugin creator - ::InstanceNormalization_TRT version 1
[TRT]    Registered plugin creator - ::InstanceNormalization_TRT version 2
[TRT]    Registered plugin creator - ::CoordConvAC version 1
[TRT]    Registered plugin creator - ::DecodeBbox3DPlugin version 1
[TRT]    Registered plugin creator - ::GenerateDetection_TRT version 1
[TRT]    Registered plugin creator - ::MultilevelCropAndResize_TRT version 1
[TRT]    Registered plugin creator - ::MultilevelProposeROI_TRT version 1
[TRT]    Registered plugin creator - ::NMSDynamic_TRT version 1
[TRT]    Registered plugin creator - ::PillarScatterPlugin version 1
[TRT]    Registered plugin creator - ::VoxelGeneratorPlugin version 1
[TRT]    Registered plugin creator - ::MultiscaleDeformableAttnPlugin_TRT version 1
[TRT]    detected model format - caffe  (extension '.caffemodel')
[TRT]    desired precision specified for GPU: FASTEST
[TRT]    requested fasted precision for device GPU without providing valid calibrator, disabling INT8
[TRT]    [MemUsageChange] Init CUDA: CPU +219, GPU +0, now: CPU 242, GPU 7633 (MiB)
[TRT]    [MemUsageChange] Init builder kernel library: CPU +352, GPU +332, now: CPU 613, GPU 7981 (MiB)
[TRT]    native precisions detected for GPU:  FP32, FP16, INT8
[TRT]    selecting fastest native precision for GPU:  FP16
[TRT]    found engine cache file /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel.1.1.8401.GPU.FP16.engine
[TRT]    found model checksum /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel.sha256sum
[TRT]    echo "$(cat /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel.sha256sum) /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel" | sha256sum --check --status
[TRT]    model matched checksum /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel.sha256sum
[TRT]    loading network plan from engine cache... /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel.1.1.8401.GPU.FP16.engine
[TRT]    device GPU, loaded /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel
[TRT]    [MemUsageChange] Init CUDA: CPU +0, GPU +0, now: CPU 277, GPU 7996 (MiB)
[TRT]    Loaded engine size: 14 MiB
[TRT]    Using an engine plan file across different models of devices is not recommended and is likely to affect performance or even cause errors.
[TRT]    Deserialization required 21899 microseconds.
[TRT]    [MemUsageChange] TensorRT-managed allocation in engine deserialization: CPU +0, GPU +13, now: CPU 0, GPU 13 (MiB)
[TRT]    Total per-runner device persistent memory is 5632
[TRT]    Total per-runner host persistent memory is 111264
[TRT]    Allocated activation device memory of size 3613184
[TRT]    [MemUsageChange] TensorRT-managed allocation in IExecutionContext creation: CPU +0, GPU +3, now: CPU 0, GPU 16 (MiB)
[TRT]    
[TRT]    CUDA engine context initialized on device GPU:
[TRT]       -- layers       89
[TRT]       -- maxBatchSize 1
[TRT]       -- deviceMemory 3613184
[TRT]       -- bindings     2
[TRT]       binding 0
                -- index   0
                -- name    'data'
                -- type    FP32
                -- in/out  INPUT
                -- # dims  3
                -- dim #0  3
                -- dim #1  224
                -- dim #2  224
[TRT]       binding 1
                -- index   1
                -- name    'prob'
                -- type    FP32
                -- in/out  OUTPUT
                -- # dims  3
                -- dim #0  1000
                -- dim #1  1
                -- dim #2  1
[TRT]    
[TRT]    binding to input 0 data  binding index:  0
[TRT]    binding to input 0 data  dims (b=1 c=3 h=224 w=224) size=602112
[TRT]    binding to output 0 prob  binding index:  1
[TRT]    binding to output 0 prob  dims (b=1 c=1000 h=1 w=1) size=4000
[TRT]    
[TRT]    device GPU, /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel initialized.
[TRT]    loaded 1000 class labels
[TRT]    imageNet -- /home/ananos/develop/jetson-inference/data/networks//bvlc_googlenet.caffemodel initialized.
class 0954 - 0.999023  (banana)
imagenet: 99.90234% class #954 (banana)
imagenet: attempting to save output image
imagenet: completed saving
imagenet: shutting down...
2023.08.09-08:29:34.82 - <debug> session:1 Free session
```
