FROM python:3.9-slim-buster

RUN apt-get update && \
	apt-get -y upgrade && \
 	apt-get -y install make automake gcc g++ subversion python3-dev libopenblas-dev liblapack-dev gfortran

RUN mkdir app
WORKDIR app

COPY . ./
COPY app.prod ./app.py

RUN pip3 install -r requirements.txt

EXPOSE 8050

CMD [ "gunicorn", "--workers=3", "--threads=1", "-b 0.0.0.0:8050", "index:server"]