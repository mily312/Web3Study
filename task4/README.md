# **个人博客系统——后端项目**



## 运行环境

- go version go1.25.0
- gin，gorm框架
- mysql数据库



## 启动方式

在项目根目录下，运行main.go文件

go run main.go



## 后端接口实现功能

### 用户注册

用户注册不需要鉴权，password采用密文方式存储

![image-20250926144549974](C:\Users\mily\AppData\Roaming\Typora\typora-user-images\image-20250926144549974.png)

![image-20250926144633162](C:\Users\mily\AppData\Roaming\Typora\typora-user-images\image-20250926144633162.png)

### 用户登录

用户登录，成功后返回jwt token

![image-20250926144855423](C:\Users\mily\AppData\Roaming\Typora\typora-user-images\image-20250926144855423.png)

密码验证失败

![image-20250926144943207](C:\Users\mily\AppData\Roaming\Typora\typora-user-images\image-20250926144943207.png)

### 文章创建

### 文章读取

### 文章修改

### 文章删除

