# Servicemesh Webinar #2 示例代码

[SerivceMesh Webinar #2](https://mp.weixin.qq.com/s/rAmxmGrnRnGfG8rFNPCYtQ) 的示例代码，分享会的回放请看[这里](https://www.bilibili.com/video/BV1AZ4y1g7nF/)。
示例可以搭建以 [MOSN](https://mosn.io) 作为 [Istio](https://istio.io/) 数据面的 servicemesh 运行环境（替换 [Envoy](https://www.envoyproxy.io/)）

详细使用方法和介绍请看文章，或者回放，如果有任何疑问也欢迎提 Issue.

如果你也对 MOSN 有兴趣，请加入我们：）


## 目录结构

```
➜  webinar2_democode git:(master) ✗ ll
drwxrwxr-x 9 trainyao trainyao  4096 7月  23 23:05 .
drwxr-xr-x 8 trainyao trainyao  4096 7月  22 02:53 ..
-rw-r--r-- 1 trainyao trainyao 19363 7月  22 15:40 envoy_bootstrap_tmpl.json # 从 docker 镜像: mosnio/proxyv2:1.5.2 取出来的 mosn 启动配置文件的模版
-rw-r--r-- 1 trainyao trainyao 19542 7月  22 15:36 envoy_bootstrap_tmpl.with_metrics.json # 上述模版，增加了 prometheus 指标暴露配置
drwxrwxr-x 2 trainyao trainyao  4096 7月  22 20:38 grpc_traffic_manage # GRPC 流量治理 demo 用到的文件
drwxrwxr-x 5 trainyao trainyao  4096 7月  20 13:16 kv # 子仓库 https://github.com/trainyao/gloo_in_none_kubernetes_env/tree/master/kv, 一个简单的用来演示的 GRPC 服务
-rw-rw-r-- 1 trainyao trainyao  1065 7月  20 12:42 LICENSE
drwxrwxr-x 2 trainyao trainyao  4096 7月  22 20:45 limit # 限流 demo 用到的文件
-rw-rw-r-- 1 trainyao trainyao   427 7月  23 23:05 README.md
-rw-rw-r-- 1 trainyao trainyao   427 7月  23 23:05 README-cn.md
drwxr-xr-x 2 trainyao trainyao  4096 7月  22 21:08 tracing # tracing demo 用到的文件

```

## 演示了 MOSN 什么功能？

- 流量治理
    - GRPC 服务的流量治理
    - GRPC 服务故障注入
    - 限流
- 服务可见性
    - prometheus 格式的指标暴露和收集
    - 分布式跟踪, 结合[apache skywalking](https://skywalking.apache.org/)
        - 还有另外一个[katacoda 的例子](https://www.katacoda.com/mosn/courses/istio/mosn-with-skywalking)
   
# 运行环境

- Istio v1.5.2
- MOSN v0.14.0 (支持了 Istio v1.5.2)

