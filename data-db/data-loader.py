# python3 data-loader.py users.json users_array

import json
import redis
from redis.commands.json.path import Path
import os
import sys

file_path = sys.argv[1]
name_obj = sys.argv[2]

# Opening JSON file
if os.path.exists(file_path):
    f = open(file_path)

# returns JSON object as 
# a dictionary
data = json.load(f)

r = redis.Redis(host='localhost', port=6379)

r.json().set(name_obj, Path.root_path(), data)