# base image
FROM python:3.9

RUN apt update
    
WORKDIR /app

COPY . .

ENV PYTHONDONTWRITEBYTECODE 1

ENV PYTHONUNBUFFERED 1

RUN pip install --upgrade setuptools pip

RUN pip install -r requirements.txt --use-deprecated=legacy-resolver

ENTRYPOINT ["python3", "app.py"]

