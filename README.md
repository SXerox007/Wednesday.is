# Wednesday.is
Roman to Integer 


## For server start
make app
## Sever will be running in 
localhost:50051

## CURL to get the roman to interger
curl -X POST -k http://localhost:50051/wednesday/1/roman --header "Content-Type:application/json" -d '{"value":"XI"}'
#
"success":true,"message":"Success","code":200,"data":11}


### To test the test util 
go test base/utils/utils_test.go -v
