from flask import Flask 

fruits = {"Apple", "Orange", "Pear", "Orange"}

app = Flask(__name__)

# Pass the required route to the decorator. 
@app.route("/hello") 
def hello(): 
    return (f"Here, have some {fruits}")

    
@app.route("/") 
def index(): 
    return "Booyah"
  
if __name__ == "__main__": 
    app.run(port=4000, debug=True) 