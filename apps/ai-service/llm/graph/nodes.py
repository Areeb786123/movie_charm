from llm.graph.state import CommentState
from llm.model import llm 

def sentiment_node(state: CommentState) -> dict:

        comment = state["comment"]

        prompt = f"""
        Analyze the sentiment of this movie comment.

        Comment:
        {comment}

        Return ONLY one word:
        POSITIVE
        or
        NEGATIVE
        """

        response = llm.invoke(prompt)

        sentiment = response.content.strip().upper()

        if "POSITIVE" in sentiment:
            sentiment = "POSITIVE"
        elif "NEGATIVE" in sentiment:
            sentiment = "NEGATIVE"
        else:
            raise ValueError(f"Invalid sentiment: {response}")

        return {
            "sentiment": sentiment
        }
        


def positive_node(state:CommentState)-> dict:
    comment = state["comment"]
    prompt=     prompt = f"""
    A user left this positive movie comment:

    {comment}

    Generate a short, friendly response to the user.
    """  
    
    data  = llm.invoke(prompt)
    response = data.content
    return {
        "message":response.strip(),
        "sentiment":state["sentiment"]
    }     
    
def negative_node(state:CommentState)-> dict:
    comment= state["comment"]
    prompt = f"""
    A user left this negative movie comment:

    {comment}

    Generate a short, polite and empathetic response to the user.
    """
    
    data = llm.invoke(prompt)
    response = data.content
    
   
    return {
        "sentiment": "NEGATIVE",
        "message": response.strip(),
    }
    

        