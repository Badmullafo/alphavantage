FROM python:3.8.10-alpine3.13

WORKDIR /usr/src/app

COPY ./alphavantage.py ./alphavantage.py

ENV LISTEN_PORT=8000
ENV SYMBOL=MSFT
ENV APIKEY=DEMO
ENV NDAYS=5

RUN python -m pip install requests

EXPOSE 8000 

ENTRYPOINT [ "python", "./alphavantage.py" ]
