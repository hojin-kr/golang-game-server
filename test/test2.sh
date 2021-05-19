#!/bin/bash
SET=$(seq 1 1000000)
for i in $SET
do
	curl -X  POST \
	https://container-service-1.29t1v2cn53l8a.ap-northeast-2.cs.amazonlightsail.com/stage/start \
	-H 'content-type: application/json' \
	-d '{"id":'$i'}'
done


#for i in $SET
#do
#	for j in $SETT
#	do
#	curl -X  POST \
#	https://container-service-1.29t1v2cn53l8a.ap-northeast-2.cs.amazonlightsail.com/stage/clear \
#	-H 'content-type: application/json' \
#	-d '{"id":'$j'}'
#	done
#done
