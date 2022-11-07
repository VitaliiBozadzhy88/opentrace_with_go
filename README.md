# **README**

## 📀 Preparing to launch a project

1. Clone project (need to push Code button and choose in what way you want to get project)
2. Open with IntelliJ IDEA (if you do not have this program - download it [here](https://www.jetbrains.com/idea/download/#section=mac))
3. You will also need Docker to this project. Download it from [here](https://www.docker.com)
4. You have to SET UP ROOT. How to do it you find [here](https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html)
5. Find in [handler](handlers/handler.go) fields - USERNAME, PASSWORD, HOST, PORT, DB.
   You have put here your UserName, Password, host, port and DB name that you must create before start work with this project. For quick start 
you may open MySql and write Query -> `create database usersdb` then another one Query -> `create table Users ( id int auto_increment primary key, email varchar(30), activation_code varchar(30))`


## 📌 How the project works

1. open terminal/cmd and put `docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest`
2. Find [service_create](services/service_create.go) and press Run button or ^R in Intellij IDEA to start server
3. Find [service_get](services/service_get.go) and press Run button or ^R in Intellij IDEA to start server
4. Open 3 new tabs in browser
5. Type in first tab: http://localhost:8080/create
6. Type in second tab: http://localhost:8081/get
7. Tab in third tab: http://localhost:16686/. Renew page 2 times
8. In Field SERVICE find service_get and service_create
9. Push button FIND TRACES. For familiar with JAEGER visit [site](https://www.jaegertracing.io) and read documentation
# opentrace_with_go
