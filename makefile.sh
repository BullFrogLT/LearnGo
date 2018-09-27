#!/bin/bash

# 读取的文件，如果想生成1M以内文件，需要修改配置
# /etc/services 文件大小在600K左
FILE=/etc/services

# 生成的文件数量
FILE_NUM=1000000
#FILE_NUM=500
# 生成的文件最大字节数
MAX_BYTE=100
#MAX_BYTE=512000

# 每100个文件 sleep 时间 
SLEEP_TIME=5

# 结果目录
RESULT_PATH=./result

tmp=`cat $FILE`
mkdir -p $RESULT_PATH


TIMES=0

function sleeptime(){
    sleep $1
}


# 循环创建文件
for file in $(seq $FILE_NUM)
do
{
   # 文件名 test+数字.bcp
   name=$(printf test%02d.bcp $file) 
   #touch "$name"
   
   # 生成随机数
   _NUM=$((RANDOM %$MAX_BYTE))
   
   # 取_NUM个字符输出到文件中
   echo ${tmp:0:$_NUM} > $RESULT_PATH/$name &
   #let TIMES=$TIMES+1

   #if [ $(($TIMES%100)) = '0' ]
   #then
   #    echo $TIMES
   #    # sleeptime $SLEEP_TIME
   #fi
   
} &
# echo $file >> ./mk.log
done


