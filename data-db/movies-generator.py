import json
import redis
from redis.commands.json.path import Path
import random 
import requests 


f = open('movies_raw.json')
data = json.load(f)



# Opening JSON file
final_data = []
id = 1


for i in data:
    print(i)
    movieData = requests.get("https://www.omdbapi.com/?t="+i+"&apikey=81b3d12e").json()
    try:
        movieData.pop("Ratings")
        movieData.pop("Metascore")
        movieData.pop("DVD")
        movieData.pop("Production")
        movieData.pop("Website")
        movieData.pop("Response")
    except:
        print("Error")
    obj = {}
    obj["id"] = str(id)
    for k in list(movieData.keys()):
        obj[k] = movieData[k]
    print(obj)
    id += 1
    final_data.append(obj)

with open("movies.json", "w") as outfile:
    json.dump(final_data, outfile)

f.close()
