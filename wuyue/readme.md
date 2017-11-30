

### glide 安装 goquery 

```
1.先安装:
  glide get golang.org/x/net/html
2. glide get github.com/PuerkitoBio/goquery
```

### 项目运行

```
1.环境安装
  配置好GOROOT和GOPATH
  cd src
  mkdir work-codes && cd work-codes
  git clone https://code.aliyun.com/wukc/wuspider.git
  cd wuspider/wuyue
  glide install

2. 运行项目
  启动mongodb:  mongod
  启动项目： go build 

```

### 项目部署

```
pm2 

```
