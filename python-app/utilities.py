
import sys


def error_handler(text: str):
    return f'{sys._getframe().f_back.f_code.co_name}() method error: {text}'
