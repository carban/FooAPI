import json
import redis
from redis.commands.json.path import Path
import random 
import requests 


# Opening JSON file
f = open('comments_raw.json')
final_data = []

# returns JSON object as 
# a dictionary
comments_raw = json.load(f)

n = lambda x: (51-x, x+1)
i = 1
postId = 1
obj = {}

for p_comments in comments_raw:
    comment1 = p_comments[0]
    comment2 = p_comments[1]
    first_user, second_user = n(postId)
    obj["id"] = str(i)
    obj["comment"] = comment1
    obj["reactions"] = random.randint(0, 5)
    obj["postId"] = str(postId)
    first = requests.get("http://localhost:8080/api/users/"+str(first_user)).json()["data"]
    obj["user"] = {
        "id": first["id"],
        "name": first["name"],
        "lastname": first["lastname"],
        "username": first["username"],
    }
    final_data.append(obj.copy())
    i += 1
    obj["id"] = str(i)
    obj["comment"] = comment2
    obj["reactions"] = random.randint(0, 5)
    second = requests.get("http://localhost:8080/api/users/"+str(second_user)).json()["data"]
    obj["user"] = {
        "id": second["id"],
        "name": second["name"],
        "lastname": second["lastname"],
        "username": second["username"],
    }
    final_data.append(obj.copy())
    i += 1
    postId += 1

with open("comments.json", "w") as outfile:
    json.dump(final_data, outfile)
 
# Closing file
f.close()