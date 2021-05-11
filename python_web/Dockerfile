FROM python:3.8.10-alpine3.13

WORKDIR /usr/src/app

COPY ./web.py ./web.py

ENV LISTEN_PORT=8000
ENV SYMBOL=MSFT
ENV APIKEY=DEMO
ENV NDAYS=5

EXPOSE 8000 

CMD [ "python", "./web.py" ]
