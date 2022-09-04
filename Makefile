codepoint_names.go : unicode.txt convertuni.py
	./convertuni.py $< $@
