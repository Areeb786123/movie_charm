from llm.graph.graph import graph

def main():
    print("MAIN FILE LOADED")
    result = graph.invoke({"comment": "The movie is terible"})
    print(result)
    
    
if __name__ == "__main__":
    main()    