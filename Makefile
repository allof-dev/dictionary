WORDNET_FILE=english-wordnet-2024.xml
DB_NAME=dict.db

run: $(DB_NAME)
	go run ./cmd/server $(DB_NAME)

batch.jsonl: $(DB_NAME)
	go run ./cmd/generate/batch-translation $(DB_NAME) $@
	@echo "Check '$@'. You must resolve follow duplications manually"
	@cat $@ | jq .custom_id | sort | uniq -d

$(DB_NAME): $(WORDNET_FILE)
	go run ./cmd/generate/db $(WORDNET_FILE) $@

$(WORDNET_FILE):
	curl https://en-word.net/static/$(WORDNET_FILE).gz -o $@.gz
	gunzip $@.gz