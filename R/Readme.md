README for ETL scripts written in R

# Quick start:

## Notes
All file paths are relative to the root directory of this Git repo. In other words, one directory up from here. (This file is R/Readme.md)

Raw data, whether hand-downloaded downloaded from FindDx or hand filled in by IQT staff are in the R/data_raw directory. Remember these file paths.

## Steps
1. Upload data to R/data_raw as necessary. Beware of overwriting previous files unless you know you wnat to. :)
2. Update the paths to these raw files in R/01_clean_eua.R and R/02_clean_find.R
3. Run the scripts in R/scripts in order, from 01_\*.R to 04_\*.R
4. .csv files of cleaned & formatted data are in the R/data_derived directory

As of this writing the files for upload to the DB or other uses are
* combined_performance.csv - one row per company + test for seeing sensitivity/specificity
* combined_ppv_npv_mat.csv - one row per company + test + prevalence for plotting NPV/PPV accuracy
* flu_plot_mat.csv - one row per prevalence level. Used for plotting flu rapid diagnostic performance data as a comparison to COVID tests.

# Notes or other design considerations
As of this writing, data come from both manually curated EUA test data (26 diagnostic tests) and only those FindDx tests where there was a single reported clinical trial.

Open issues include:
1. Figuring out how to use/combine multiple clinical trial data from FindDx
2. Proper matching and deconfliction of tests in Find's DB to tests in the EUA DB. As of now, spelling or description changes aren't considered in the curation scripts.
3. Incorporating data from additional non-covid tests if we want them available for context

# Data dictionary as of 2020-07-08

Asterisks denote columns, which when combined, uniquely identify each row.

## combined_performance.csv

*company: manufacturer name
*test_name: test name
test_type: type of test (should be PCR for all tests right now)
clinical_lod_or_both: Do we have data from clinical trials, limit of detection, or both?
ppa: Positive percent agreement = tp / n_pos
npa: Negative percent agreement = tn / n_neg
target_specimen: List of methods for specimen collection for performance estimate
notes: Notes from collecting performance data from FDA IFU documents
tp: True positives
fp: False positives
tn: True negatives
fn: False negatives
n_pos: Total number of positive samples = TP + FN
n_neg: Total number of negative samples = TN + FP
specimen: If provided, specimen collected for performance estimate. If total, aggregated across specimens listed in target_specimen
sensitivity: Point estimate of sensitivity
specificity: Point estimate of specificity
sensitivity_95ci_low: lower bound of the 95% CI on sensitivity
sensitivity_95ci_high: upper bound of the 95% CI on sensitivity
specificity_95ci_low: lower bound of the 95% CI on specificity
specificity_95ci_high: upper bound of the 95% CI on specificity
source: One of FindDx or FDA EUA - depending on where we got the data

## combined_ppv_npv_mat.csv

*company: manufacturer name
*test_name: test name
*prevalence: prevalence of covid-19 (x axis of plots)
ppv: point estimate of positive predictive value for given prevalence level & test
ppv_95ci_low: 95% CI lower bound of ppv
ppv_95ci_high: 95% CI upper bound of ppv
npv: point estimate of negative predictive value for given prevalence level & test
npv_95ci_low: 95% CI lower bound of npv
npv_95ci_high: 95% CI upper bound of npv
p_pos: point estimate of the percent of tests that will return positive for a given prevalence level & test
p_pos_low: 95% CI lower bound of the percent of tests that will return positive
p_pos_high: 95% CI upper bound of the percent of tests that will return positive
fdr: point estimate for the false discovery rate for a given prevalence level & test
fdr_95ci_low: 95% lower bound of fdr
fdr_95ci_high: 95% upper bound for fdr
fomr: point estimate for the false omission rate for a given prevalence level & test
fomr_95ci_low: 95% lower bound of fomr
fomr_95ci_high: 95% upper bound of fomr
