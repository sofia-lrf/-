
# 将镜像推送至 docker 官方镜像仓库
docker push sofia123/cncfclassestest:testv1

# 通过 docker 命令本地启动 httpserver
root@cvm:~# docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS     NAMES
83711f80d4f4   2e93acdadb6d   "/bin/sh -c /httpser…"   22 minutes ago   Up 22 minutes   80/tcp    nostalgic_ramanujan

# 通过 nsenter 进入容器查看 IP 配置
root@cvm:~# nsenter -t 32768 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
4: eth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever