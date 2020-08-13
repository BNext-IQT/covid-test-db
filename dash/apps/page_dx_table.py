import os
import dash
import dash_core_components as dcc
import dash_html_components as html
import dash_bootstrap_components as dbc

import urllib
import httpx
import dash_table as dt
import pandas as pd
from app import app

pd.set_option("display.precision", 2)

# Load csv into pandas dataframe
REPO_DIR = os.path.dirname(os.path.abspath(__file__))
DATA_DIR = os.path.join(os.path.split(REPO_DIR)[0],'data/')
data_filename = 'combined_ppv_npv_mat.csv'
data_file_path = os.path.join(DATA_DIR, data_filename)

resp = httpx.get('http://localhost/api/diagnostics')
df = pd.json_normalize(resp.json(), sep="_")
print(f'{df.to_dict("records")}')
# normalized = pd.json_normalize(df.to_json())
# print(f'{normalized}')
# Setup formatting for display
display_columns = [
           {'id':'company_name', 'name':'Company'},
           {'id':'name', 'name':'Test Name'},
           {'id':'diagnosticType_name', 'name':'Type'},
           {'id':'pcrPlatforms', 'name':'Instrument/Platform'},
           {'id':'sourceOfPerfData', 'name':'IFU/EUA'},
           {'id':'sampleTypes', 'name':'Sample Types'},
           {'id':'pointOfCare', 'name':'Point-of-Care'},
           {'id':'prepIntegrated', 'name':'Integrated Sample Prep.'},
           ]

# for col, f in formats.items():
#     df[col] = df[col].map(lambda x: f.format(x))


# Get unique company names
#unique_companies = df.company.unique()

markdown_text_disclaimer = '''
## Disclaimer

This page and the related application were developed and are being released by IQT in collaboration with its Bnext initiative.
Please note that the application displays standard statistical calculations of 3rd party data as reported.
Bnext and IQT cannot guarantee the accuracy of the reported data.
'''

#create table row
def generate_row(datarow):
    link_id = datarow['id']+'_EUA'
    return html.Tr([
        html.Td(datarow['company_name']),
        html.Td(datarow['name']),
        html.Td(datarow['diagnosticType_name']),
        html.Td([html.Div(pcr['name']) for pcr in datarow['pcrPlatforms']]),
        html.Td([html.A('IFU/EUA',
              id=link_id,
              href=datarow['sourceOfPerfData'] if datarow['sourceOfPerfData'] else '',
              target="_blank"
        )]),
        html.Td([html.Div(st['name']) for st in datarow['sampleTypes']]),
        html.Td('Y' if datarow['pointOfCare'] else 'N'),
        html.Td('Y' if datarow['prepIntegrated'] else 'N'),
    ])


# Generate the table from the pandas dataframe
def generate_table(dataframe):
    return html.Table(
        # Header
        [html.Tr([html.Th(col['name']) for col in display_columns])] +

        # Body
        [generate_row(dataframe.iloc[i]) for i in range(len(dataframe))]
    )

initial_table = generate_table(df)
#app = dash.Dash()
#
#app.layout = html.Div(children=[
layout = html.Div(children=[
      dbc.Container([
        html.Br(),
        
        # Header
        # html.H4('Select a Company and a Test'),

        # # Dropdown menu of company names
        # dcc.Dropdown(id='dropdown-company', 
        #              options=[{'label': i, 'value': i} for i in unique_companies],
        #              multi=False,
        #              placeholder='Filter by company...'),

        # html.Br(),

        # # Dropdown menu of test names
        # dcc.Dropdown(id='dropdown-test-name',
        #              value='',
        #              multi=False,
        #              placeholder='Filter by test...'),
        
        # html.Br(),

        # html.Hr(),

        # Link to download data to csv file
        # html.A(dbc.Button(
        #        'Download Selected Data',
        #        color='primary',
        #        className='three columns',
        #       ),
        #       id='download-link',
        #       download="rawdata.csv",
        #       href="",
        #       target="_blank"
        # ),

        html.Hr(),

        # Data dump of table on screen

        html.Div(id='table-cont', children=initial_table),

        html.Hr(),

        # Intro markdown text from above
        dcc.Markdown(children=markdown_text_disclaimer, id='markdown'),

        html.Hr(),

        ]
      )
  ]
)

# @app.callback(
#         dash.dependencies.Output('table-cont', 'children'),
#         [dash.dependencies.Input('dropdown-company', 'value'),
#          dash.dependencies.Input('dropdown-test-name', 'value')]
#     )
# def display_table(dropdown_value, dropdown_test_value):
#     '''Display table via options to generate it'''

#     if dropdown_value is None or dropdown_test_value is None:
#         return generate_table(df)

#     dff = df[df.company.str.contains(dropdown_value) &
#              df.test_name.str.contains(dropdown_test_value)]
#     return generate_table(dff)

# Include open source css file
app.css.append_css({"external_url": "https://codepen.io/chriddyp/pen/bWLwgP.css"})

if __name__ == '__main__':
    app.run_server(debug=True)

