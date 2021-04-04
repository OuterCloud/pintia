method=$1
# 组装测试命令
fileList=`ls`
cmd="go test"
for file in $fileList
do
    if echo "$file" | grep -q -E '\.go$';then
        cmd="${cmd} $file"
    fi
done
if [ ! -n "$1" ];then
echo "未指定测试方法"
else
cmd="$cmd -test.run $method"
fi
cmd="$cmd -v -gcflags=-l -coverprofile=test.out"
# 执行测试命令
$cmd
# 生成覆盖率报告
go tool cover -html=test.out -o test.html