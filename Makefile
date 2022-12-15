build:
	@go build -o ./currency-api .

run:
	@./currency-api

clean:
	@rm -fv ./currency-api 