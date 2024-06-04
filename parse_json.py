import json

with open("words_dictionary.json", "r") as file:
    words = json.load(file)

filtered_words = [word for word, count in words.items() if 4 == len(word)]

with open("4_letter_words.json", "w") as file:
    json.dump(filtered_words, file)