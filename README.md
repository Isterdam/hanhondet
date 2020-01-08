# HanHonDet

In older Swedish, nouns could either have a masculine, feminine or neutral gender. You would, for example, refer to a _clock_ as _she_ and the _moon_ as _he_. In contemporary Swedish, the masculine and feminine gender have been merged into a so called _common gender_, while the neutral gender is still in use. However, the three genders for nouns are still used in some dialects: in the table below there is an example of how nouns would be conjugated in traditional Scanian (the ortography is from M. Lucazins book, "Utkast till ortografi över skånska språket").

|              | Én herð**e** (A shepherd) | Éna gribbe (A girl) | Étt hjarta (A heart) |
|--------------|-------------|-----------|------------|
|    Gender     | Hanð (He)  | Honð (Her)   | Dæd (It)    |
| Indefinite article plural | Herð**a** | Gribb**er** | Hjart**en** |
| Definite article singular | Herð**en** | Gribb**ena** | Hjart**að** |
| Definite article plural | Herð**ana** | Gribb**erna** | Hjart**enen** |

In this project (which was developed during my first year of Computer Science at KTH), I developed a simple website that displays the former three grammatical genders of Swedish nouns. Firstly, I wrote a web scraper that scrapes the words with the corresponding genders from an older dictionary (Dalins ordbok) and writes them to JSON. Secondly, these were added to a PostgreSQL database. Lastly, I wrote a website primarily in Golang, some JS to make it dynamic and a simple HTML+CSS structure for front-end. The final project looks something like this in practice:
