from langgraph.graph import StateGraph , START , END
from llm.graph.state import CommentState
from llm.graph.nodes import sentiment_node, positive_node , negative_node
from llm.routes.sentimentRoutes import route_sentiment

graph_builder = StateGraph(CommentState)

graph_builder.add_node("sentiment", sentiment_node)
graph_builder.add_node("positive", positive_node)
graph_builder.add_node("negative", negative_node)

graph_builder.add_edge(START, "sentiment")
graph_builder.add_conditional_edges("sentiment", route_sentiment, {
    "positive": "positive",
    "negative": "negative",
})

# Both branches finish the graph
graph_builder.add_edge("positive", END)
graph_builder.add_edge("negative", END)

graph = graph_builder.compile()