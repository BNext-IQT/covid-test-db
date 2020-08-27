import os
import pandas as pd
import json

# Load COVID csv data into pandas dataframe
data_filename = 'data/combined_ppv_npv_mat.csv'
df = pd.read_csv(data_filename)

# Grab list of unique companies in combined file
unique_companies = df.company.unique()
n_unique_companies = unique_companies.shape[0]

# Initialize total number of tests
n_unique_tests = 0

# Dictionary to pair company names w/data files
co_name_dict = {}
co_name_json = 'data/company_names.json'

for selected_company in unique_companies:
    # Grab company specific data
    df_company_filtered = df.loc[df['company'] == selected_company]

    # See number of tests assoc with company
    avail_tests = df_company_filtered.test_name.unique()
    n_avail_tests = len(avail_tests)

    # Add to total tests
    n_unique_tests += n_avail_tests

    # Make a unique filename, condensing company name
    ofilename = ''.join(e for e in selected_company if e.isalnum()) + '.csv'
    outfile = os.path.join('data/bycompany', ofilename)
    print('Saving company {} data with {} test(s) to {}'.format(selected_company,
                                                              n_avail_tests,
                                                              outfile))
    # Save company test data to file
    df_company_filtered.to_csv(outfile, index=False, header=True)

    # Save company names and assoc filenames in dictionary
    co_name_dict[selected_company] = ofilename

# Save total number of tests
co_name_dict['n_unique_tests'] = n_unique_tests

# Save name -> datafile dictionary to json file
with open(co_name_json, 'w') as fo:
    json.dump(co_name_dict, fo)

with open(co_name_json) as fi:
  data = json.load(fi)
for key,value in data.items():
    print('key', key, 'value', value)
