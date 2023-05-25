local:
	(docker-compose down --rmi all) \
	&& (docker-compose up -d)
test-back-end:
	(docker-compose down) \
	 && (docker-compose -f docker-compose.be.test.yml run --rm app-test go test -v ./...) \
    && (docker image rm cus-app-test-img)
test-front-end:
	(docker-compose down) \
	&& (docker-compose -f docker-compose.fe.test.yml run --rm frontend-test yarn run coverage) \
	&& (docker image rm cus-frontend-test-img)
	