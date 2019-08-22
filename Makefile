sonar:
	sonar-scanner

#Usage: make all -log="k8s client test"
all:
	git add .
	git commit -m "$(log)"
	git push origin master