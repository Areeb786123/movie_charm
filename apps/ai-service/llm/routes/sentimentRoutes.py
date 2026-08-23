from llm.graph.state import CommentState

def route_sentiment(state: CommentState):
    if state["sentiment"] == "POSITIVE":
        return "positive"
    return "negative"