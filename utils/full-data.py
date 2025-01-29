import os

# List of items to iterate over
files = ["songs.json", "comments.json", "posts.json", "cities.geojson", "movies.json", "users.json", "todos.json", "products.json"]
names = ["songs_array", "comments_array", "posts_array", "cities_array", "movies_array", "users_array", "todos_array", "products_array"]

# Path to the script to execute
script_to_run = "data-loader.py"

for i, item in enumerate(files):
    os.system(f"python3 {script_to_run} {item} {names[i]}")
    print(f"python3 {script_to_run} {item} {names[i]} OK")
