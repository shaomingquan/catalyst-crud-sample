govendor fetch github.com/shaomingquan/webcore/gene
govendor fetch github.com/shaomingquan/webcore/core
govendor fetch github.com/shaomingquan/webcore

echo "update runtime libs"

go get github.com/shaomingquan/webcore

echo "update exe"