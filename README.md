# fetchfileforblog
该小工具是用来方便我随时保存网上图片到我的七牛空间的
============
#使用方法
## 一、先修改:
1. accKey和secKey 为你自己的key
2. domain也是你自己bucket绑定的域名
3. 你自己bucket的名字

## 二、编译
可以直接clone下来，之后
```
go build fetch.go

也可以直接 go get 下来
再go install
```

## 三、安装
把文件复制到/usr/local/bin/即可
```
cp fetch /usr/local/bin/

```

# 用法
```
fetch -url=http://images.cnblogs.com/cnblogs_com/yjf512/201206/201206142215457646.png
```

output:

```
fetching http://images.cnblogs.com/cnblogs_com/yjf512/201206/201206142215457646.png
fetch success
the response url:

http://oa8dpdexh.bkt.clouddn.com//201206142215457646.png

```

返回的路径可以供你写入markdown博客。

