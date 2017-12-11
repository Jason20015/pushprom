for((i=0.001;i<10240;i+0.001))
do
curl -H "Content-type: application/json" -X POST -d '{"type": "histogram", "name":"response_time","help": "response time for request", "method": "observe","value":'$i',"labels":{"service":"video","inter":"getvideolist"}}' http://192.168.209.23:9091/
done
