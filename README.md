# 注意

填写 vhost 参数时，如果是 */* 开头的，要进行编码（%2F）。

比如，你的 vhost 为 */beta* ，则amqp URL填写为:

```bash
"amqp": "amqp://guest:guest@127.0.0.1:5672/%2Fbeta"
```


如果是默认的 vhost （/），则URL填写为:

```bash
"amqp": "amqp://guest:guest@127.0.0.1:5672/%2F"
```

# app.json

这个是配置 rabbitmq 的文件

# data.txt

则是存放数据的文件，格式为：每一行一条数据。

程序会忽略空行，以及内容前后的空格
