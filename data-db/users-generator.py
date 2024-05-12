import json
import redis
from redis.commands.json.path import Path
import random 

# Opening JSON file
f = open('users.json')
 
# returns JSON object as 
# a dictionary
data = json.load(f)

first = data[0] 

male_list = ["John", "Martin", "Carlos", "Jamal", "James", "Paulo", "Bryan", "Mohamed", "Akira", "Leonardo", "Sadio", "Bill", "Alexander", "Caesar", "Ronaldo"]
female_list = ["Jane", "Rachel", "Nicol", "Sofia", "Alexa", "Rebecca", "Elizabeth", "Marie", "Lea", "Margareth", "Latika", "Alexa", "Zooey", "Lucy"]
lastnames = ["Doe", "Rodriguez", "Smith", "Muller", "Da Silva", "Rossi", "Garcia", "Nakamoto", "Shelley", "Von Trier", "Tarantino", "Mendy", "Asimov"]
countries = ["USA", "Brazil", "Italy", "Japan", "France", "India", "Colombia", "United Kingdom", "Senegal", "Greece", "Egypt", "Morocco", "Denmark", "Netherlands", "Germany", "Australia", "Mexico"]

phone_indicator = {
    "USA": "+1",
    "Brazil": "+55",
    "Italy": "+39",
    "Japan": "+81",
    "France": "+33",
    "India": "+91",
    "Colombia": "+57",
    "United Kingdom": "+44",
    "Senegal": "+221", 
    "Greece": "+30", 
    "Egypt": "+20", 
    "Morocco": "+212", 
    "Denmark": "+45", 
    "Netherlands": "+31", 
    "Germany": "+49", 
    "Australia": "+61", 
    "Mexico": "+52"
}

saved_names = {}

for i in range(1, 50):
    first = dict(first)
    first["id"] = str(i+1)
    first["gender"] = random.choice(["Male", "Female"])
    if (first["gender"] == "Male"):
        first["name"] = random.choice(male_list)
    elif (first["gender"] == "Female"):
        first["name"] = random.choice(female_list)
    first["lastName"] = random.choice(lastnames)
    if first["name"]+first["lastName"] not in saved_names:
        saved_names[first["name"]+first["lastName"]] = 1
    else:
        print("repeated")
        while(first["name"]+first["lastName"] in saved_names):
            if (first["gender"] == "Male"):
                first["name"] = random.choice(male_list)
            elif (first["gender"] == "Female"):
                first["name"] = random.choice(female_list)
            first["lastName"] = random.choice(lastnames)         
    first["username"] = random.choice(["alpha", "beta", "mega", ""])+first["lastName"][:3]+first["name"]+random.choice(["x", "xD", "gg"])+str(random.randint(10, 999))    
    first["height"] = random.randint(150, 200) 
    first["weight"] = round(random.uniform(50.0, 90.0), 2)
    first["email"] = "foo"+first["name"].lower()+first["lastName"].replace(" ", "").lower()+str(random.randint(10, 999))+"@fakeemail.com"
    first["country"] = random.choice(countries)
    first["phone"] = phone_indicator[first["country"]]+" "+str(random.randint(100, 999))+" "+str(random.randint(100, 999))+" "+str(random.randint(1000, 9999))
    year = random.choice([str(i) for i in range(1944,2006)])
    first["birthDate"] = random.choice(["0"+str(i) if i < 10 else str(i) for i in range(1,28)])+"-"+random.choice(["0"+str(i) if i < 10 else str(i) for i in range(1,12)])+"-"+year 
    first["age"] = 2023 - int(year)
    data.append(first)

with open("users.json", "w") as outfile:
    json.dump(data, outfile)
 
# Closing file
f.close()