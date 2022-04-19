from utilities import error_handler
import requests
from decouple import config
from json import loads


class Quotes():

    def __init__(self):
        self.all_quotes: list = []
        self.quotes_list = None
        self.url = config('API_URL')

    def call_api(self):
        try:
            url = f'{self.url}'
            r = requests.get(url)
            data = r.text
            self.all_quotes.append(loads(data))
        except Exception as e:
            raise Exception(error_handler(e))

    def assemble_values(self) -> list:
        try:
            for quotes in self.all_quotes:
                date_ = quotes['date']
                quotes = quotes['rates']
                values = {'Uuid': 'valor fixo',
                          'euro': quotes['EUR'],
                          'brl': quotes['BRL'],
                          'jpy': quotes['JPY'],
                          'of_date': date_}

                self.quotes_list = values
        except Exception as e:
            raise Exception(error_handler(e))

    def return_quotes(self) -> dict:
        try:
            return self.quotes_list
            # print(self.quotes_list)
        except Exception as e:
            raise Exception(error_handler(e))

    def run(self) -> dict:
        self.call_api()
        self.assemble_values()
        return self.return_quotes()
        
