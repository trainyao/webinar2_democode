# Servicemesh Webinar #2 demo code

[Chinese Introduction](./README-cn.md)

Demo code for [SerivceMesh Webinar #2](https://mp.weixin.qq.com/s/rAmxmGrnRnGfG8rFNPCYtQ), playback video is [here](https://www.bilibili.com/video/BV1AZ4y1g7nF/).
You can use this repo to test [MOSN](https://mosn.io) feature in [Istio](https://istio.io/), to create a MOSN data plane service mesh, instead of [Envoy](https://www.envoyproxy.io/).

Please follow the steps in the page or playback video, for more information, and create issue when having trouble using it. 

Join us if you are insterested in MOSN too :)


## Files

```
➜  webinar2_democode git:(master) ✗ ll
drwxrwxr-x 9 trainyao trainyao  4096 7月  23 23:05 .
drwxr-xr-x 8 trainyao trainyao  4096 7月  22 02:53 ..
-rw-r--r-- 1 trainyao trainyao 19363 7月  22 15:40 envoy_bootstrap_tmpl.json # original mosn startup config template in docker image: mosnio/proxyv2:1.5.2
-rw-r--r-- 1 trainyao trainyao 19542 7月  22 15:36 envoy_bootstrap_tmpl.with_metrics.json # mosn startup config template with metric config, for prometheus metrics demo
drwxrwxr-x 2 trainyao trainyao  4096 7月  22 20:38 grpc_traffic_manage # for GRPC traffic management demo
drwxrwxr-x 5 trainyao trainyao  4096 7月  20 13:16 kv # git submodule for https://github.com/trainyao/gloo_in_none_kubernetes_env/tree/master/kv, a simple GRPC demo server and client
-rw-rw-r-- 1 trainyao trainyao  1065 7月  20 12:42 LICENSE
drwxrwxr-x 2 trainyao trainyao  4096 7月  22 20:45 limit # for rate limit demo
-rw-rw-r-- 1 trainyao trainyao   427 7月  23 23:05 README.md # this doc
-rw-rw-r-- 1 trainyao trainyao   427 7月  23 23:05 README-cn.md
drwxr-xr-x 2 trainyao trainyao  4096 7月  22 21:08 tracing # tracing demo

```

## Feature of MOSN to demostrate

- traffic management
    - GRPC service traffic management
    - GRPC service fault injection
    - rate limit
- service observation
    - metrics exporting and graping, in prometheus format 
    - tracing, with [apache skywalking](https://skywalking.apache.org/)
        - see also in [this katacoda demo](https://www.katacoda.com/mosn/courses/istio/mosn-with-skywalking)
   
# Works on

- Istio v1.5.2
- MOSN v0.14.0 (which support Istio v1.5.2)

