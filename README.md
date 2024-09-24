# wechatBot

基于WindowsHooks库开发的微信机器人，目前已对接GLM模型，实现了AI聊天功能。文档地址：https://www.showdoc.com.cn/WeChatProject/8929096202824749


## 运行环境

- Golang 1.20+
- Docker 1.13+

## 快速开始

要快速启动项目，请按照以下步骤进行操作：

1. **配置环境变量**

   根据项目目录下的`.env.example`文件，创建`.env`文件，并将配置文件中的变量值替换为实际的值。

2. **构建Docker镜像**

   使用以下命令在本地构建Docker镜像：
    ```bash
    docker build -t wechatbot .
    ```

3. **运行容器**

   运行刚刚构建的镜像作为一个Docker容器：
   ```bash
   docker run -d \
   --name wechatbot \
   -p 8080:8080 \
   --env-file /path/to/.env \
   wechatbot
   ```

