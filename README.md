# discover-animal-crossing-villagers
This project was created solely for learning purposes. There are various optimizations and changes I would make in order to make data retrieval faster. For instance, this data could live inside a database rather than each time making the Lambda call which gathers the data from the Nookipedia API.

This project API built on top of the Nookipedia API (which makes use of MediaWiki's API) that scraps villager information form the text on the page and allows the user to filter based on species in the front end.

Serverless Golang API uses AWS Lambda function connected to an API Gateway. The front end is hosted from an S3 bucket.
https://github.com/PuerkitoBio/goquery
![alt-text](./preview_gif.gif)
