#!/bin/sh

CUR_DIR=`pwd`
GOPATH=${CUR_DIR}
export GOPATH

gofmt -l -w -s src/

cd ${CUR_DIR} && go build -o city-info ./src/app/main.go && \
rm -rf output && mkdir -p output && \
cp -R deploy-meta city-info conf control.sh supervise.city-info ./output

ret=$?
if [ $ret -ne 0 ]; then
    echo "===== build failure ====="
    exit $ret
else
    echo "===== build successfully! ====="
fi

exit
