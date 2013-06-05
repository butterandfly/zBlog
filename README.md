zBlog
=====
ZBlog is a blog system that can be run on google app engine. It use the zCore framework. ZBlog use a bsd-style license, please feel free to use it.
ZBlog是一个用golang编写，基于google app engine及zCore的博客程序。ZCore是我编写的一个轻量级web框架，你可以通过zBlog来学习zCore的使用。此程序使用类bsd许可，请自由、放心地使用。
本人的博客(<a href="http://www.zeno-l.com">www.zeno-l.com</a>)，使用了zBlog程序。

How to use
-----
1）编辑app.yaml，将application属性中的enteryourappname改为你的google app engine中的appid。
2）打开你的app，进入http://你的域名/addadmin，页面会要求你输入管理员的google账户（若填错或想更改，可在你的google app engine管理页面中把AdminEntity删除，然后再次进入此页面）。
3）点击首页右上方的Management按钮，会转跳到google的登陆页面，登陆你的2)中填入的用户，即可开始发表文章。

About me
-----
博客：www.zeno-l.com
