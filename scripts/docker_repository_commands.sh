#Go to the working directory
cd ..

#Login into docker 
docker login

#Tag the image because the namespace of the docker hub needs to be included
docker tag kube-metrics-monitor jagostinho900/kube-metrics-monitor

#push the image to the hub
docker push jagostinho900/kube-metrics-monitor