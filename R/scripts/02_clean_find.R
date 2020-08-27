################################################################################
# This script loads and formats diagnostic test performance data from FindDx.
# Original data are manually downloaded from https://finddx.shinyapps.io/COVID19DxData/
# There are three files, one for each type of test: molecular (PCR), antibody, antigen
# Output is two tables, one for tests with a single trial and one for tests with
# multiple trials 
################################################################################

### Parameters: Change by hand when needed ----
# path_to_data <- list.files(
#   "R/data_raw/2020-07-01",
#   full.names = TRUE
# )

path_to_data <- "R/data_raw/2020-08-26/COVIDDxData.xlsx"

### Load libraries ----
library(tidyverse)
library(janitor)
library(readxl)

### Source accuracy diagnostic scripts ----
source("R/scripts/00_accuracy_functions.R")

#### Read in the tables into a list ----
find <- 
  path_to_data %>%
  map(.f = function(path){
    read_xlsx(path) %>%
      clean_names() %>% 
      select(
        company = manufacturer,
        test_name,
        test_type,
        performance_target_specimen = target,
        n_pos = total_positive,
        n_neg = total_negative,
        tp = true_positive,
        tn = true_negative,
        performance_notes = comments
      ) %>%
      map(.f = function(x) {
        x[x == "Not available"] <- NA
        x
      }) %>%
      as_tibble() %>%
      mutate(
        ppa = case_when(is.na(tp) ~ "", ! is.na(tp) ~ paste0(tp, "/", n_pos)),
        npa = case_when(is.na(tn) ~ "", ! is.na(tn) ~ paste0(tn, "/", n_neg))
      ) %>% 
      mutate(
        tp = as.numeric(tp),
        tn = as.numeric(tn),
        n_pos = as.numeric(n_pos),
        n_neg = as.numeric(n_neg),
        fp = as.numeric(n_neg) - as.numeric(tn),
        fn = as.numeric(n_pos) - as.numeric(tp),
        specimen = NA,
        clinical_lod_or_both = NA
      ) %>%
      select(
        company,
        test_name,
        test_type,
        clinical_lod_or_both,
        ppa,
        npa,
        performance_target_specimen,
        performance_notes,
        tp,
        fp,
        tn,
        fn,
        n_pos,
        n_neg,
        specimen
      ) %>% 
      filter(! duplicated(.)) # remove any duplicated rows
  }) 

find <- do.call(rbind, find)

# remove any records for which we do not have performance data
# remove rows without clinical performance data
find <- 
  find %>%
  filter(! (is.na(tp) | is.na(fp) | is.na(tn) | is.na(fn))) %>%
  mutate(
    sensitivity = tp / n_pos,
    specificity = tn / n_neg,
  )

### Calculate performance data with uncertainty ----

# get confidence intervals for sensitivity and specificity
conf_ints <-
  find %>%
  pmap_dfr(function(...){
    x <- tibble(...)
    
    sensitivity_ci <- 
      calc_wilson_ci(x$tp, x$n_pos)
    
    specificity_ci <-
      calc_wilson_ci(x$tn, x$n_neg)
    
    tibble(
      sensitivity_95ci_low = sensitivity_ci[1],
      sensitivity_95ci_high = sensitivity_ci[2],
      specificity_95ci_low = specificity_ci[1],
      specificity_95ci_high = specificity_ci[2]
    )
  })

find <- 
  cbind(find, conf_ints)


### get one trial per test ----
# For tests with multiple trials, we take the one that has the largest sample size

find <- by(find, INDICES = paste(find$company, find$test_name), function(x) x)

find_formatted <- lapply(
  find,
  function(x) {
    row <- which.max(x$n_pos + x$n_neg)
    row <- row[1]
    x[row, ]
  }
)

find_formatted <- do.call(rbind, find_formatted)

save(
  find_formatted,
  file = "R/data_derived/find_performance.RData"
)

# ### split data into tests with only one trial and tests with many trials ----
# 
# find <- by(find, INDICES = paste(find$company, find$test_name), function(x) x)
# 
# num_trials <- sapply(find, nrow)
# 
# find_single_trial <- do.call(rbind, find[num_trials == 1]) %>% as_tibble()
# 
# find_multiple_trial <- do.call(rbind, find[num_trials > 1]) %>% as_tibble()
# 
# ### save result ----
# save(
#   find_single_trial,
#   find_multiple_trial,
#   file = "R/data_derived/find_performance.RData"
# )
