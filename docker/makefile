nodes = 4
txs = 1000

up: build conf start watch

build :
	./scripts/build-images.sh

conf : 
	rm -rf conf
	./scripts/build-conf.sh $(nodes)

start:
	./scripts/run-testnet.sh $(nodes)

watch: 
	./scripts/watch.sh $(nodes)

stop: 
	./scripts/stop-testnet.sh

demo: 
	./scripts/demo.sh $(nodes)

bombard: 
	./scripts/bombard.sh $(nodes) $(txs)

.PHONY: up build conf start watch stop demo bombard
	