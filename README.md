## 教程链接
https://codingwithpatrik.dev/content/how-to-build-a-cli-todo-app-in-go
https://codingwithpatrik.dev/thank-you-epic


## 编译所有 .go 文件，生成 todo.exe（关键：-o 指定 .exe 后缀 Windows下）
go build -buildvcs=false -o todo.exe

## 列出所有待办
.\todo.exe -list

## 添加待办
.\todo.exe -add "Windows 测试待办"

## 编辑待办（索引0，新标题）
.\todo.exe -edit "0:修改后的测试标题"

## 切换待办状态
.\todo.exe -toggle 0

## 删除待办
.\todo.exe -del 0
