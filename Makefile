.PHONY: gen tidy test lint all

all: gen tidy test lint

gen:
	sh ./scripts/gen.sh

tidy:
	sh ./scripts/tidy.sh

test:
	sh ./scripts/test.sh

lint:
	sh ./scripts/lint.sh


.PHONY: run
run: ## run {svc} server. example: make run svc=product
	@scripts/run.sh ${svc}

.PHONY: gensvc
gensvc: ## generate thrift code
	@scripts/gensvc.sh ${svc}