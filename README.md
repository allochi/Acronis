# Acronis
Acronis technical test

Download zip
 
= Request =
 
Endpoint:
POST /downloadzip
 
Body:
[
               "FullFileName1",
               "FullFileName2",
               ...
               "FullFileName3"
]
 
For example
[
               "/home/user/log.txt",
               "/home/user/messages.txt"
]
 
= Response =
 
Binary stream with zip archive contained all files enumearted in request
