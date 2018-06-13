all:
	@echo make {dev,demo,prod} to deploy

dev:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r ap-southeast-1 --env dev deploy

demo:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r ap-southeast-1 --env demo deploy

prod:
	@echo $$AWS_ACCESS_KEY_ID
	apex -r ap-southeast-1 --env prod deploy

testdev:
	apex -r ap-southeast-1 --env dev invoke email < event.json

testdemo:
	apex -r ap-southeast-1 --env demo invoke email < event.json

testprod:
	apex -r ap-southeast-1 --env prod invoke email < event.json


deletedev:
	apex -r ap-southeast-1 --env dev delete


.PHONY: dev demo prod testdev testdemo testprod
