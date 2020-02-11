import pymongo
from pymongo import MongoClient

cluster = MongoClient("mongodb+srv://dbUser:<password>@cluster0-5nzly.mongodb.net/test?retryWrites=true&w=majority")

class Class:
    classDB = cluster["classdb"]
    classCollection = classDB["feedbackCollection"]
    userID=None
    name=None
    time=None
    prof=None
    building=None
    choice=None

    post0 = {"_id": 0, "className": "sampleClass", "startTime": "12:00 PM", "professor": "sampleProfessor", "building": "Bascom Hall", "choice": 'g'}

    def __init__(self, mName, mTime, mProf, mBuilding):
        self.userID = 1
        self.name = mName
        self.time = mTime
        self.prof = mProf
        self.building = mBuilding
        self.classCollection.insert_one(self.post0)
    
    def addChoice(self, char):
        tempPost = {"_id": self.userID, "className": "sampleClass", "startTime": "12:00 PM", "professor": "sampleProfessor", "building": "Bascom Hall", "choice": char}
        self.classCollection.insert_one(tempPost)

    
