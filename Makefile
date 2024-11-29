WORDNET_FILE=english-wordnet-2024.xml


dict.db: $(WORDNET_FILE)
	go run ./cmd/generate/db $(WORDNET_FILE) $@

$(WORDNET_FILE):
	curl https://en-word.net/static/$(WORDNET_FILE).gz -o $@.gz
	gunzip $@.gz