# git flow 工作流
在<git核心原理>中提到git最大的一个特性： *每个提交都包括了所有文件* 。

这样就导至一个问很尴尬的问题，只要我提交到git上，代码就会被其它人提交到线上去，而这个代码并没有开发完成。

面对这种问题一般就两种解决方案：

1. 写不完坚决不提交到git
2. 写不全坚决不往主线合并代码

git flow就是第2种思路的一个解决方案

![git集中式仓库](./dots/centralized.dot)

git flow是一种集在式仓库的管理方，大家都通过一个公共git仓库进行工作。

## 分支方案
git flow分支示意图:

![gitflow分支](./dots/git-workflow.png)

这边需要特别注意的是,只有master和develop是实际上分支名，其它的都是一组分支的统称。

git flow相关[资料](http://blog.jobbole.com/76867/)

相关操作
```sh
//初始化
git flow init
//开发一个功能
git flow feature start <your feature>
//合并到开发主线
git flow feature finish
//合并到开发主线并推送到远程
git flow feature publish
//推送到远程
git flow feature pull
//发布测试版本
git flow release start
//测试版通过合并到线上主线
git flow release finish
//修补线上bug
git flow hotfix start
git flow hotfix finish
//修补开发分支bug
git flow bugfix start
git flow bugfix finish
```


