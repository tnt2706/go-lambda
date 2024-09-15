.PHONY: build zip deploy update delete

ROLE := arn:aws:iam::009790545675:role/tinhtn-lambda-hello-role
FUNCTION_MAME := hello

build:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

zip:
	zip hello.zip bootstrap

deploy:
	aws lambda create-function --function-name $(FUNCTION_MAME) \
	--runtime provided.al2023 --handler bootstrap \
	--architectures arm64 \
  	--role $(ROLE) \
	--zip-file fileb://$(FUNCTION_MAME).zip

update:
	aws lambda update-function-code --function-name $(FUNCTION_MAME) \
	--zip-file fileb://$(FUNCTION_MAME).zip

delete:
	aws lambda delete-function --function-name $(FUNCTION_MAME)