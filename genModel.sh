#!/bin/bash
# 使用方法：
# ./genModel.sh {表名} 比如： ./genModel.sh user
# 再将 ./model 目录下的文件剪切到对应服务的 model 目录里面

# 将数据库配置信息修改成你自己的
Host=127.0.0.1
Port=3306
DBName=go_zero_tutorial
UserName=root
PassWD=123456

# 需要对哪张表生成模型
Table=$1
# 表生成的模型存放的目录
ModelDir=./model

Green="\033[32m"
Red="\033[31m"
GreenBG="\033[42;37m"
RedBG="\033[41;37m"
Font="\033[0m"

OK="${GreenBG}OK!${Font}"
Error="${RedBG}Error:${Font}"

if [[ -z $1 ]];
then
  echo -e "${Error} ${Red} 请输入表名! ${Font}"
  exit 1
fi

echo -e "${OK} ${Green} 正在对数据库【 $DBName 】中的 【 $Table 】数据表创建模型 ${Font}"
# 或者通过 ddl 生成
# -c 或者 -cache=true 代表生成带缓存的模型
# goctl model mysql ddl -c -src ./sql/user.sql -dir ./model --style=goZero
goctl model mysql datasource -url="${UserName}:${PassWD}@tcp(${Host}:${Port})/${DBName}" -table="${Table}"  -dir="${ModelDir}" -cache=true --style=goZero

exit 0