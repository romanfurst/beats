PLATFORMS=linux/amd64 make snapshot
echo "build"
docker tag docker.elastic.co/beats/filebeat-oss:8.0.0-SNAPSHOT nexus3.kb.cz:18444/speed/filebeat:1.0.0-SNAPSHOT
docker push nexus3.kb.cz:18444/speed/filebeat:1.0.0-SNAPSHOT
echo "pushed"
