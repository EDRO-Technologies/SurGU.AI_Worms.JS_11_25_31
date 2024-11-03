import os


class Config(object):
    def __init__(self):
        self.BACKEND_URL = os.getenv("BACKEND_URL")
        self.TG_API_KEY = os.getenv("TG_API_KEY")
        self.FRONTEND_URL = os.getenv("FRONTEND_URL")

        for var in vars(self):
            if self.__getattribute__(var) is None:
                raise ValueError(f"tg_bot: переменная '{var}' должна быть установлена в окружении")

    def __new__(cls):
        if not hasattr(cls, 'instance'):
            cls.instance = super(Config, cls).__new__(cls)
        return cls.instance
