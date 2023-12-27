# 1. 版本说明
- golang版本 1.17
- gin版本 1.7.6
- gorm v1.25.2-0
- grpc 1.52.0
- consul/api 1.18.0
- consul服务器端 1.17.1


# 2. 部署说明


# 3.特别说明
仅作为golang相关框架学习使用，使用gin、gorm和grpc框架目前实现了rbac认证相关功能

email：wanstack@163.com

# 4. 缺陷 todo
1. 定义gorm中的model时没有指定字段大小，导致数据库中的字段太大
2. 未实现前段页面
3. 创建表未优化
4. 目前日志仅打印到终端，未使用logrus等工具进行集成