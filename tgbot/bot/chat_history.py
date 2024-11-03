from dataclasses import dataclass, asdict
from typing import List


@dataclass
class ChatMessage:
    is_ai: bool
    content: str


@dataclass
class ChatHistory:
    chat: List[ChatMessage]

    def clear(self) -> None:
        self.chat.clear()

    def append(self, message: ChatMessage) -> None:
        self.chat.append(message)

    def dict(self) -> dict:
        return asdict(self)
