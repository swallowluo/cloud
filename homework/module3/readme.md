1、将Dockerfile与httpserver.go上传到同一目录下。

编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化

2、构建本地镜像

docker build -t "jzhisun/cncamp-httpserver:V1.0" ./

按照docker官方镜像库的规范，定义镜像名称。

3、将镜像推送至 docker 官方镜像仓库

docker push jzhisun/cncamp-httpserver:V1.0 #上传镜像

docker pull jzhisun/cncamp-httpserver:V1.0 #下载镜像

4、启动镜像

docker run -d -p 8080:80/tcp jzhisun/cncamp-httpserver:V1.0 #启动镜像

5、查看容器IP地址

```
export CONTAIN_ID=`docker ps |grep -i jzhisun |grep -v grep |awk '{print $1}'`
echo $CONTAIN_ID
export CONTAIN_PID=`docker inspect --format "{{.State.Pid}}" $CONTAIN_ID`
echo $CONTAIN_PID
nsenter --target $CONTAIN_PID -n ip a
```
