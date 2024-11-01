import json
import redis
from redis.commands.json.path import Path
import random 
import requests 


# Opening JSON file
f = open('todos_raw.json')
final_data = []

# returns JSON object as 
# a dictionary
todos_raw = json.load(f)

print(len(todos_raw))

i = 1
userId = 1
obj = {}

for p_todos in todos_raw:
    todo1 = p_todos[0]
    todo2 = p_todos[1]
    obj["id"] = str(i)
    obj["todo"] = todo1
    obj["state"] = random.choice(["Pending", "Working", "Completed"])
    obj["closed"] = True if random.randint(0, 1) else False
    obj["userId"] = str(userId)
    final_data.append(obj.copy())
    i += 1
    obj["id"] = str(i)
    obj["todo"] = todo2
    obj["state"] = random.choice(["Pending", "Working", "Completed"])
    obj["closed"] = True if random.randint(0, 1) else False
    obj["userId"] = str(userId)
    final_data.append(obj.copy())
    i += 1
    userId += 1

with open("todos.json", "w") as outfile:
    json.dump(final_data, outfile)
 
# Closing file
f.close()