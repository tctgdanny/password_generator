import random
import json

with open("word_lists/5_letter_words.json", "r") as file:
    words_list = json.load(file)

if words_list:
    for _ in range(5):
        word = random.choice(words_list)
        number = random.randint(0, 9999)
        number_str = str(number).zfill(4)
        word = f"{word.capitalize()}+{number_str}"
        print(word)
else:
    print("Error: Could not get words list.")