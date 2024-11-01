import json
import redis
from redis.commands.json.path import Path
import random 

# Opening JSON file
f = open('products_raw.json')
 
# returns JSON object as 
# a dictionary
data = json.load(f)
final_data = []
id = 1

for p in data:
    obj = {}
    obj["id"] = str(id)
    obj["title"] = p["title"]
    obj["description"] = p["description"]
    obj["brand"] = p["brand"]
    obj["category"] = p["category"]
    obj["price"] = round(random.uniform(10.0, 30.0), 1)
    obj["rating"] = round(random.uniform(2.0, 5.0), 1)
    obj["stock"] = random.randint(3, 15)
    final_data.append(obj)
    id += 1

with open("products.json", "w") as outfile:
    json.dump(final_data, outfile)
 
# Closing file
f.close()