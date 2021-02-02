#
curl -X POST -k http://localhost:50051/wednesday/1/roman --header "Content-Type:application/json" -d '{"value":"XI"}'
#
"success":true,"message":"Success","code":200,"data":11}



go test base/utils/utils_test.go -v
=== RUN   TestRomanToInteger
2021/02/02 17:43:19 Item: 88
2021/02/02 17:43:19 Item: 73
2021/02/02 17:43:19 Test IX: 9
--- PASS: TestRomanToInteger (0.00s)
PASS
ok      command-line-arguments  (cached)




