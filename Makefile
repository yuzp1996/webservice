sonar:
	sonar-scanner


all: sonar
	git add .
	git commit -m "push to master"
	git push origin master