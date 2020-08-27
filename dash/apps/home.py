import dash_core_components as dcc
import dash_html_components as html
import dash_bootstrap_components as dbc
from config import APP_PATH

# needed only if running this as a single page app
#external_stylesheets = [dbc.themes.LUX]


markdown_text_intro = '''
Diagnostic tests are a critical tool to contain epidemics, to support medical care, and for public health measures.
Understanding when they are accurate and inaccurate is necessary for knowing which individuals have the novel coronavirus, need isolation, and need their contacts traced.

Measuring the accuracy of a diagnostic test is complicated. This is true for diagnostic tests of novel coronavirus and for more common diseases like influenza.

We have made this post to help you know what goes into understanding and interpreting the accuracy of various diagnostic tests for the novel coronavirus.
'''

markdown_text = '''

## Determinants of diagnostic test accuracy

Some components to measuring the accuracy of diagnostic tests are:

* What the test measures
* Characteristics of a diagnostic test
* What question(s) the test should answer
* Prevalence of the disease in the tested population

### What does "accuracy" mean for diagnostic tests?

There are four statistics of "accuracy" for diagnostic tests.
All four are important, measure different things, and are interconnected. 

* **Sensitivity** is the fraction of people with a disease in a tested population that receive a positive test result.
* **Specificity** is the fraction of people who do _not_ have a disease in a tested population that receive a negative test result.

Sensitivity and specificity are characteristics of the diagnostic test alone.
They are usually measured and reported during a clinical trial prior to a diagnostic test's approval for public use.

* **Positive predictive value (PPV)** is the fraction of positive tests where the person actually had the disease.
* **Negative predictive value (NPV)** is the fraction of negative tests where the person actually did not have the disease.

PPV and NPV are calculated from sensitivity, specificity, and the prevalence of the disease in the population being tested. 

### What do various diagnostic tests for novel coronavirus measure?

There are three types of diagnostic test for the novel coronavirus. 

* **Molecular PCR** tests measure the presence of novel coronavirus _directly_. These tests detect the presence of viral RNA. 
* **Antibody** tests and **antigen** tests measure the presence of novel coronavirus _indirectly_. These tests detect the body's reaction to the presence of novel coronavirus. 

Molecular PCR tests and antigen tests are used to diagnose current infection of novel coronavirus. 
Antibody tests are used to diagnose current or recent past infection of novel coronavirus. 
However, [current CDC guidance](https://www.cdc.gov/coronavirus/2019-ncov/lab/resources/antibody-tests-professional.html)
states that **antibody tests should not be used to diagnose an active infection of the novel coronavirus**.
Molecular PCR tests can generally detect infection sooner than antibody or antigen tests, as it may take days or more than a week for the body to ramp up its response.


### Determining characteristics of a test

During clinical trials, manufacturers collect samples from people who both do and do not have a disease.
They use these samples to calculate sensitivity and specificity. 
There is uncertainty about the measurements of sensitivity and specificity because the clinical trials use a sample of people. 
Having larger samples from positive people increases confidence about the measure of sensitivity. 
Having larger samples from negative people increases confidence about the measure of specificity. 
Researchers measure the uncertainty about sensitivity and specificity using 
[confidence intervals](https://www.graphpad.com/guides/prism/7/statistics/stat_more_about_confidence_interval.htm) and carry that uncertainty to downstream calculations.

### Questions a diagnostic test can answer

When a medical or public health offical administers a diagnostic test, they may seek answers to two questions.

* Does the person being tested have the disease? To answer this question, the test must have a high PPV. If PPV is low, it is likely that a diagnostic test will produce many false positives.
* Does the person being tested _not_ have the disease? To answer this question, the test must have a high NPV. If NPV is low, it is likely that a diagnostic test will produce many false negatives.

The prevalence of the disease also affects PPV and NPV.
For a fixed sensitivity and specificity, low disease prevalence lowers PPV; high disease prevalence lowers NPV.

[The CDC guidelines on influenza rapid diagnostic testing](https://www.cdc.gov/flu/professionals/diagnosis/clinician_guidance_ridt.htm)
reflect that the rapid flu diagnostic shows a similar pattern in relation to disease prevalence.
From the guidelines:

> False-positive  results are more likely to occur when influenza  prevalence in the population tested (e.g. community) is low, which is generally at the beginning and end of the influenza season and during the summer.
> False-negative  results are more likely to occur when influenza  prevalence is high in the community.
### How to think about "prevalence" of a disease in the tested population

The simplest definition of prevalence is the percent of people in a population that have a disease, like novel coronavirus.
Knowing the true prevalence in the population is difficult, especially with a shortage of tests.
But there are heuristics and strategies that can be helpful.

Consider geographic differences.
For example it's known that the rate of infection with novel coronavirus is much higher in New York City than, say, Santa Clara, CA. 
(Whatever the criticisms of the respective prevalence studies, there's no debate about the considerable difference in the rates of infection between the two regions.)

Consider seasonal differences. Researchers don't yet know about the seasonal effect of the novel coronavirus pandemic.
But other diseases are seasonal, as noted by the quote from the CDC about influenza, above. 

Consider other information.
For example there are sub-populations with higher exposure, such as front-line healthcare workers or contacts of someone known to be infected. 

Clinicians may also use the presence or absence of symptoms or results of other diagnostic tests to help interpret results.

**[The University of Minnesota's Center for Infectious Disease Research and Policy has guidance for novel coronavirus testing that discusses prevalence and testing for various subpopulations.](https://www.cidrap.umn.edu/sites/default/files/public/downloads/cidrap-covid19-viewpoint-part3.pdf)**

#### Examples of Coronavirus Prevalence


* **1.3% to 4.7%** - In April of 2020, [researchers estimated](https://www.medrxiv.org/content/10.1101/2020.04.14.20062463v2) that 2.8% (95% confidence interval of 1.3% to 4.7%) of the residents of Santa Clara County, CA had antibodies for the novel coronavirus. Note that this study [was criticized](https://www.theguardian.com/world/2020/apr/23/coronavirus-antibody-studies-california-stanford) for its [statistical methods and interpretation of test accuracy](https://www.statschat.org.nz/page/4/) and may have _over_ estimated prevalence.
* **1.8% to 3.9%** - In April of 2020, [researcher estimated](https://www.medrxiv.org/content/10.1101/2020.04.26.20079822v1.full.pdf) that between 2.7 % (95% confidence interval of 1.8% to 3.9%) of residnece of Kobe, Japan had antibodies for the novel coronavirus.
* **2.8%** - In May of 2020, [researchers estimated](https://news.iu.edu/stories/2020/05/iupui/releases/13-preliminary-findings-impact-covid-19-indiana-coronavirus.html) about 2.8% of the residents of Indiana had antibodies for the novel coronavirus.
* **19.2%** - In January of 2020, [public health officials found](https://www.cdc.gov/mmwr/volumes/69/wr/mm6912e3.htm) that 712 of 3,711 passengers on the Diamond Cruise Ship tested positive for the novel coronavirus with a molecular test.
* **21.5% to 24.0%** - In April of 2020, [researchers estimated](https://www.medrxiv.org/content/10.1101/2020.05.25.20113050v1) that 22.7% (95% confidence interval of 21.5% to 24.0%) of New York City residents had antibodies for the novel coronavirus indicating current or past infection.
* **85.6%** - In March of 2020, [researchers found](https://www.nejm.org/doi/full/10.1056/NEJMoa2005412) that 101 out of 118 residents of a long term care facility in King County, Washington tested positive for Coronavirus using a molecular diagnostic test.
'''


markdown_text_disclaimer = '''
## Disclaimer

This page and the related application were developed and are being released by IQT in collaboration with B.Next.
Please note that the application displays standard statistical calculations of 3rd party data as reported.
B.Next and IQT cannot guarantee the accuracy of the reported data.
'''

#app = dash.Dash(__name__, external_stylesheets=external_stylesheets)

# change to app.layout if running as single page app instead
layout = html.Div(children=[
    dbc.Container([
        dbc.Row([
                 dbc.Col(html.H1("Understanding Diagnostic Test Accuracy", className="text-center"), className="mb-5 mt-5")
                 #dbc.Col(html.H1("Welcome to IQT Labs' COVID Toolkit", className="text-center"), className="mb-5 mt-5")
                ]),
    
        # Intro markdown text from above
        dcc.Markdown(children=markdown_text_intro, id='markdown'),
        
        # Link to results tool
        html.A(dbc.Button(
               'Go to Results Tool',
               color='primary',
               className='three columns',
              ),
              id='results-link',
              href="{}/home".format(APP_PATH),
              target="_blank"
        ),

        html.Br(),
        html.Br(),

        # Markdown text from above
        dcc.Markdown(children=markdown_text, id='markdown'),

        # Link to results tool
        html.A(dbc.Button(
               'Go to Results Tool',
               color='primary',
               className='three columns',
              ),
              id='results-link',
              href="{}/home".format(APP_PATH),
              target="_blank"
        ),

        html.Hr(),

        # Lower page buttons to links
        dbc.Row([
                 dbc.Col(dbc.Card(children=[html.H3(children='Access data & code used to build this dashboard',
                                                    className="text-center"),
                                            html.A([
                                            html.Img(src="{}/assets/GitLab_Logo.png".format(APP_PATH),
                                                     style={
                                                            'height' : '7vw',
                                                            'min-height' : '1vw',
                                                            'padding-top' : 10,
                                                            'padding-bottom' : 10,
                                                           }
                                                    )
                                                   ], className='text-center', href='https://gitlab.iqt.org/zhampelarias/covid-dash-app'),
                                           ], body=True, color="dark", outline=True), width=2, lg=4, className="mb-4"),

                 dbc.Col(dbc.Card(children=[html.H3(children='Explore other work from B.Next',
                                                    className="text-center"),
                                            html.A([
                                            html.Img(src="{}/assets/BNext_Logo.png".format(APP_PATH),
                                                     style={
                                                            'height' : '7vw',
                                                            'min-height' : '1vw',
                                                            'padding-top' : 10,
                                                            'padding-bottom' : 10,
                                                            'className': 'text-center'
                                                           }
                                                    )
                                                   ], className='text-center', href='https://www.bnext.org/'),
                                           ], body=True, color="dark", outline=True), width=2, lg=4, className="mb-4"),

                 dbc.Col(dbc.Card(children=[html.H3(children='Explore other research areas of IQT Labs',
                                                    className="text-center"),
                                            html.A([
                                            html.Img(src="{}/assets/IQT_Labs_Logo.png".format(APP_PATH),
                                                     style={
                                                            'height' : '7vw',
                                                            'min-height' : '1vw',
                                                            'padding-top' : 10,
                                                            'padding-bottom' : 10,
                                                           }
                                                    )
                                                   ], className='text-center', href='https://www.iqt.org/labs/'),
                                           ], body=True, color="dark", outline=True), width=2, lg=4, className="mb-4")
                ], justify='around', className="mb-5"),
                #], justify='around', align='stretch', className="mb-5"),
        
        html.Hr(),
        
        # Intro markdown text from above
        dcc.Markdown(children=markdown_text_disclaimer, id='markdown-disc'),

        html.Hr(),

        html.A("Special thanks to Tommy Jones for main development of toolkit in R."),
        
        html.Hr(),

    ])

])

# needed only if running this as a single page app
# if __name__ == '__main__':
#     app.run_server(host='127.0.0.1', debug=True)
