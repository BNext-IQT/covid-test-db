################################################################################
# This script loads clean EUA and FindDx data, combines them into a single table,
# calculates NPV and PPV for each test over a range of disease prevalence, and
# spits out a "performance" table where one row is one test and a "plot" table
# where one row is one test + disease prevalence level
################################################################################


### Load libraries ----
library(tidyverse)
library(janitor)
library(readxl)

### Source accuracy diagnostic scripts ----
source("R/scripts/00_accuracy_functions.R")


### Read in source data ----
load("R/data_derived/eua_performance.RData")
load("R/data_derived/find_performance.RData")


### Combine single trial tests ----
# ignoring multiple trial tests from find b/c I don't know how to handle them yet
# adding a column for data source (EUA/FindDx)

find_single_trial$source <- "FindDx"

eua_performance$source <- "FDA EUA"

combined_performance <- rbind(
  eua_performance,
  find_single_trial
)

# remove any where sensitivity/specificity is NaN (which happens)
combined_performance <- 
  combined_performance %>%
  filter(! is.nan(sensitivity) & ! is.nan(specificity)) %>%
  filter( !
            (sensitivity_95ci_low < 0 | 
               specificity_95ci_low < 0 | 
               sensitivity_95ci_low > 1 | 
               specificity_95ci_high > 1)
  )


# calculate performance data
combined_ppv_npv_mat <- 
  combined_performance %>%
  pmap_dfr(function(...) {
    x <- tibble(...)
    
    plot_mat <- 
      calc_plot_matrix(
        tp = x$tp,
        fn = x$fn,
        tn = x$tn,
        fp = x$fp
      ) %>%
      mutate(
        company = x$company,
        test_name = x$test_name
      ) %>%
      select(
        company,
        test_name,
        prevalence,
        ppv,
        ppv_95ci_low = ppv_low,
        ppv_95ci_high = ppv_high,
        npv,
        npv_95ci_low = npv_low,
        npv_95ci_high = npv_high,
        p_pos,
        p_pos_low,
        p_pos_high,
        fdr,
        fdr_95ci_low = fdr_low,
        fdr_95ci_high = fdr_high,
        fomr,
        fomr_95ci_low = fomr_low,
        fomr_95ci_high = fomr_high
      )
  })

save(
  combined_performance,
  combined_ppv_npv_mat,
  file = "R/data_derived/combined_calculated.RData"
)

write_csv(
  x = combined_performance,
  path = "R/data_derived/combined_performance.csv",
  na = ""
)

write_csv(
  x = combined_ppv_npv_mat,
  path = "R/data_derived/combined_ppv_npv_mat.csv",
  na = ""
)