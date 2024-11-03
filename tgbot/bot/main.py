import asyncio
import logging
from aiogram import Bot, Dispatcher, types
from aiogram.enums import ParseMode
from aiogram.filters.command import Command
from chat_history import ChatHistory, ChatMessage
from config import Config
from api.client import ApiClient

logging.basicConfig(level=logging.INFO)

config = Config()

bot = Bot(token=config.TG_API_KEY)

dp = Dispatcher()

history = {}

api_client = ApiClient(config.BACKEND_URL)


@dp.message(Command("start"))
async def cmd_start(message: types.Message):
    global history
    history[message.chat.id] = ChatHistory(chat=[])
    await message.answer("Здравствуйте!\nОтправьте мне сообщение, чтобы я мог с вами поговорить")


@dp.message(Command("clear"))
async def clear(message: types.Message):
    curr_history = history.get(message.chat.id)
    if curr_history is None:
        await message.answer("Для начала диалога со мной отправьте команду /start")
        return
    curr_history.clear()
    print(curr_history)
    await message.answer("История очищена! Теперь я не помню, что вы мне говорили до этого...")


@dp.message()
async def echo(message: types.Message):
    await bot.send_chat_action(chat_id=message.chat.id, action="typing")
    curr_history = history.get(message.chat.id)
    if curr_history is None:
        await message.answer("Для начала диалога со мной отправьте команду /start")
        return
    curr_history.append(ChatMessage(is_ai=False, content=message.text))
    response = await api_client.send_prompt(history=curr_history)
    answer = response.get("federalChapter").get("answer")
    page_number = response.get("federalChapter").get("page_number")
    curr_history.append(ChatMessage(is_ai=True, content=f"{answer}"))
    print(curr_history)
    link = f"\n\n{config.FRONTEND_URL}/{page_number}" if page_number > 0 else ""
    await message.answer(answer + link, parse_mode=ParseMode.MARKDOWN)


async def main():
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
