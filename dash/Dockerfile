FROM python:3.8-slim-buster

RUN mkdir app
WORKDIR app

COPY . ./

RUN pip3 install -r requirements.txt

EXPOSE 8050

CMD [ "gunicorn", "--workers=1", "--threads=1", "-b 0.0.0.0:8050", "index:server"]