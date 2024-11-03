FROM python:3.12-slim

WORKDIR /tgbot

COPY . .

RUN pip install -r requirements.txt

ENV PYTHONPATH="${PYTHONPATH}:/tgbot"

ENTRYPOINT ["python3.12", "/tgbot/bot/main.py"]