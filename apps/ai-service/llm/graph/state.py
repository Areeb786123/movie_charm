from typing import Literal, TypedDict


class CommentState(TypedDict):
    comment: str
    sentiment: Literal["POSITIVE", "NEGATIVE"]
    message: str