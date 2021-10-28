sonar:
	sonar-scanner
        dd

# Usage: make all log="k8s client test"
all:sonar
	git add .
	git commit -m "$(log)"
	git push origin master
