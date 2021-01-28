## halo博客备份本地

由于本人服务器临近到期不想丢失辛苦创作的博客，且囊肿羞涩不想续费，遂开发此项目。

### 主要功能

- 按照分类将halo的博客拉取到本地
- 将blog中的图片链接导到图床

### 需要配置的地方

- `write/savecontent.go` 
  - filepath : 保存路径
  - oldImgPath : 服务器上halo配置的图片地址
  - newImgPath : 图床地址
- `fetcher/fetch.go` (需要浏览器登录后抓包获取)
  - userAgent
  - Authorization
  - cookie
- `main.go`
  - seed ： halo后台的文章列表URL



### 效果展示

![halo中的博客](https://gitee.com/BzmAi/picture-bed/raw/master/image-20210128161530406.png)



**拉取到本地后**

![](https://gitee.com/BzmAi/picture-bed/raw/master/image-20210128161559739.png)