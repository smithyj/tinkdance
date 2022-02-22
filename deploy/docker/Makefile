# 启动容器
start:
	docker-compose up -d
# 停止容器
stop:
	docker-compose stop
# 重启容器
restart:
	docker-compose restart
# 删除容器
rm:
	make stop
	docker-compose rm -f
# 编译镜像
build:
	docker-compose build
# 删除镜像
rmi:
	docker images | grep tinkdance | awk '{print $$3}' | xargs -rn 1 docker rmi || true
# 删除容器与镜像
clean:
	make stop
	make rm
	make rmi
# 进入 Golang 容器
exec-golang:
	docker exec -it tinkdance_golang /bin/bash