#!/bin/bash
SET=$(seq 0 300)
for i in $SET
do
	curl -X POST https://container-service-1.29t1v2cn53l8a.ap-northeast-2.cs.amazonlightsail.com/stage/get/all
done
