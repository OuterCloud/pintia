method=$1
fileList=`ls`
cmd="go test -v"
for file in $fileList
do
    if echo "$file" | grep -q -E '\.go$';then
        cmd="${cmd} $file"
    fi
done
cmd="$cmd -test.run $method -gcflags=-l"
$cmd