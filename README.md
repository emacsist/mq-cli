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

# 用例


```bash
[14:46:33] emacsist:mq-cli git:(master) $ ./mq_cli-macos
目标队列名为: hello
send data OK =>hello world
send data OK =>alsdkfslkfs
send data OK =>fsaf
send data OK =>sdf
send data OK =>as
send data OK =>f
send data OK =>asdf
send data OK =>asf
send data OK =>sa
send data OK =>fsa
send data OK =>f
send data OK =>sf
send data OK =>saf
send data OK =>as
send data OK =>f
send data OK =>sdf
send data OK =>asf
send data OK =>as
send data OK =>f
send data OK =>asf
Done.
```
