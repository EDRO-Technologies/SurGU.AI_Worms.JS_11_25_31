import httpx

from bot.chat_history import ChatHistory


class ApiClient:
    def __init__(self, address: str) -> None:
        self.address = address

    async def send_prompt(self, history: ChatHistory) -> dict:
        request_data = {
                'act_id': 0,
                'msg': history.dict()['chat']
        }
        print(request_data, type(request_data['msg']))
        async with httpx.AsyncClient() as client:
            r = await client.post(self.address + "/api/prompt", json=request_data)
            print(r.json())
        return r.json()
