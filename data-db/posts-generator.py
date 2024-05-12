import json
import redis
from redis.commands.json.path import Path
import random 

# Opening JSON file
f = open('posts_raw.json')
u = open('users.json')


# returns JSON object as 
# a dictionary
data_and_titles = json.load(f)
data = data_and_titles[0]
titles = data_and_titles[1]

users = json.load(u)
final_data = []

for index, p in enumerate(data):
    p = p.split(" ")
    tags = []
    for i, w in enumerate(p):
        if "#" in w:
            tags.append(w)
            p.pop(i)
    if "#" in p[-1]:
        print(p[-1])
        tags.append(p[-1])
        p.pop(-1)       
    msg = " ".join(p)
    final_data.append({
        "id": str(index+1),
        "title": titles[index],
        "content": msg,
        "visibility": "public" if random.randint(0, 1) else "friends only",
        "tags": tags,
        "reactions": random.randint(0, 15),
        "user": {
            "id": users[index]["id"],
            "name": users[index]["name"],
            "lastName": users[index]["lastName"],
            "userName": users[index]["username"]
        }
    })

with open("posts.json", "w") as outfile:
    json.dump(final_data, outfile)
 
# Closing file
f.close()