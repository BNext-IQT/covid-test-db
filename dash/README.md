# COVID-Dash-App

Dash application for displaying COVID infection probabilities provided specified test, test results and COVID prevalence table.


# Installation

First setup a standalone virtual environment:
```
virtualenv -p /usr/local/bin/python3 venv
source venv/bin/activate
```

Next, clone the repo and install the required packages:
```
git clone https://gitlab.iqt.org/zhampelarias/covid-dash-app.git
cd covid-dash-app
pip install -r requirements.txt
```

Finally, launch the app process via `python index.py` 
and then point a browser to [http://127.0.0.1:8050/](http://127.0.0.1:8050/).
