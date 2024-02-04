#!/bin/sh

# 判断当前目录下是否存在名为 dockerCopilotZspace-new 的二进制文件
if [ -f "./dockerCopilotZspace-new" ]; then
    # 如果存在，则用它覆盖 dockerCopilotZspace
    mv ./dockerCopilotZspace-new ./dockerCopilotZspace
    # 赋予 dockerCopilotZspace 执行权限
    chmod +x ./dockerCopilotZspace
fi

# 运行 dockerCopilotZspace
./dockerCopilotZspace