# VSCode 远程容器编程和调试

这个教程会告诉你如何在非 Linux 环境下，用 vscode 来编译和在 Linux 容器的环境上调试你的代码。

这个教程用的是 golang 语言，可其他语言的步骤也很相似。

### 下载代码
`git clone https://github.com/ehan1990/vs-code-container.git`

### 下载 remote containers 插件
<img width="926" alt="Screen Shot 2021-05-27 at 11 56 59 AM" src="https://user-images.githubusercontent.com/1936983/119881901-a5e3fb00-bee2-11eb-8302-64ca49dc7751.png">

### 用 vscode 打开代码
<img width="727" alt="Screen Shot 2021-05-27 at 11 35 00 AM" src="https://user-images.githubusercontent.com/1936983/119878948-931bf700-bedf-11eb-9f24-09ce4c592fd1.png">

可以看出来这是个非常简单的程序。如果你在 MacOS 运行，`curl localhost:8080` 会返回 `darwin`. Linux 会返回 `linux`.

<img width="491" alt="Screen Shot 2021-05-27 at 11 47 41 AM" src="https://user-images.githubusercontent.com/1936983/119880687-59e48680-bee1-11eb-979a-4cbf8c45335d.png">

我本机是个 MacOS, 所以会显示 `darwin`.

### 在 Linux 容器运行
`F1` 会打开 command pallete. 然后选择 `Remote-Containers: Open Folder in Container`

<img width="714" alt="Screen Shot 2021-05-27 at 11 54 39 AM" src="https://user-images.githubusercontent.com/1936983/119881620-543b7080-bee2-11eb-8c8d-9de483efda0b.png">

点击 `open`

<img width="789" alt="Screen Shot 2021-05-27 at 11 55 15 AM" src="https://user-images.githubusercontent.com/1936983/119881683-69b09a80-bee2-11eb-93de-2dd76ea6f873.png">

vscode 会重新启动在容器里。 如果左下角看到这个绿色的标志的话，恭喜你。你的 vscode 启动了个 Linux 容器。现在任何代码改动都会在容器里同步。也可以在容器里编译，设断点，和调试代码。

<img width="276" alt="Screen Shot 2021-05-27 at 12 01 21 PM" src="https://user-images.githubusercontent.com/1936983/119882418-42a69880-bee3-11eb-94d4-f053400e642f.png">

可以看到如果在命令行里运行 `uname -a` 会返回 `Linux`
 ·
<img width="810" alt="Screen Shot 2021-05-27 at 12 02 05 PM" src="https://user-images.githubusercontent.com/1936983/119882522-5ce07680-bee3-11eb-9307-3c9e3c338f3b.png">

下载 `Go` 插件. 这个能让你在代码里设断点。

<img width="958" alt="Screen Shot 2021-05-27 at 12 04 38 PM" src="https://user-images.githubusercontent.com/1936983/119882841-b8126900-bee3-11eb-95ce-c96264f66c21.png">

在第十行设一个断点。然后按 `F5` 来在 vscode 运行代码。然后在命令行运行 `curl localhost:8080`. 可以看到你会在断点上停住。

<img width="642" alt="Screen Shot 2021-05-27 at 12 07 13 PM" src="https://user-images.githubusercontent.com/1936983/119883165-14758880-bee4-11eb-9c1f-0aba913d8cb7.png">

如果继续让程序跑下去的话会返回 `server is running on linux`

<img width="647" alt="Screen Shot 2021-05-27 at 12 10 24 PM" src="https://user-images.githubusercontent.com/1936983/119883568-877eff00-bee4-11eb-98a1-2b5ac59dff7c.png">
