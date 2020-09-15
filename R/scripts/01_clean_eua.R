################################################################################
# This script loads and formats diagnostic test performance data that has been
# hand extracted from IFU documents for covid diagnostic tests that have FDA
# EUA approval. The source data for this script are in a hand-created .xlsx
# file. As a result, automation may be elusive. The data set is not large,
# so be sure to go check your results and tweak the script as needed.
################################################################################

### Parameters: Change by hand when needed ----
download.file("https://github.com/BNext-IQT/covid-test-db/raw/master/scraper/Database_Master.xlsx",
              destfile = "R/data_raw/Database_Master.xlsx",
              quiet = TRUE)

path_to_data <- "R/data_raw/Database_Master.xlsx"

### Load libraries ----
library(tidyverse)
library(janitor)
library(readxl)

### Source accuracy diagnostic scripts ----
source("R/scripts/00_accuracy_functions.R")


### Read in data table ----
eua <- read_xlsx(path_to_data, sheet = 2) %>%
  clean_names()


#### Extract performance data ----
# This part may be especially brittle
eua_performance <- 
  eua %>%
  # mutate(
  #   test_type = "Molecular" # as of 2020-06-29 this was all that was in the DB, may need to change later
  # ) %>%
  select(
    company,
    test_name,
    test_type,
    clinical_lod_or_both,
    ppa,
    npa,
    performance_target_specimen,
    performance_notes
  ) 

# parse performance data, where it exists
# for now, going to aggregate across target specimenes
# create a flag column and add the note 
# "clinical performance aggregated across target specimenes by IQT"
eua_performance <- 
  eua_performance %>%
  pmap_dfr(
    .f = function(...) {
      x <- tibble(...) # preserves x as a tibble
      
      if (is.na(x$ppa)) { # no clinical trial data
        x <- 
          x %>%
          mutate(
            tp = NA,
            fp = NA,
            tn = NA,
            fn = NA,
            n_pos = NA,
            n_neg = NA
          )
      } else { # we have clinical trial data
        if (str_detect(x$ppa, "=")) { # multiple sample types
          samples_ppa <- 
            str_split(x$ppa, "\n")[[1]] %>%
            str_trim() %>%
            str_split("=") %>%
            lapply(function(row) {
              nums <- 
                str_split(row[2], "/")[[1]] %>%
                as.numeric()
              tibble(
                specimen = str_trim(row[1]), 
                tp = nums[1],
                n_pos = nums[2],
                fn = nums[2] - nums[1]
              )
            })
          
          samples_ppa <- 
            do.call(rbind, samples_ppa)  %>%
            mutate(
              company = x$company,
              test_name = x$test_name
            )
          
          samples_npa <- 
            str_split(x$npa, "\n")[[1]] %>%
            str_trim() %>%
            str_split("=") %>%
            lapply(function(row) {
              nums <- 
                str_split(row[2], "/")[[1]] %>%
                as.numeric()
              tibble(
                specimen = str_trim(row[1]), 
                tn = nums[1],
                n_neg = nums[2],
                fp = nums[2] - nums[1]
              )
            })
          
          samples_npa <-
            do.call(rbind, samples_npa) %>%
            mutate(
              company = x$company,
              test_name = x$test_name
            )
          
          x <- 
            full_join(x = x, y = samples_ppa) %>%
            full_join(samples_npa)
          
          # keep only total, aggregating across samples
          if ("total" %in% tolower(x$specimen)) {
            x <- 
              x %>%
              filter(
                tolower(specimen) == "total"
              )
          } else {
            x <- 
              x %>%
              mutate(
                tp = sum(tp),
                fn = sum(fn),
                tn = sum(tn),
                fp = sum(fp),
                n_pos = sum(n_pos),
                n_neg = sum(n_neg)
              ) %>%
              .[1, ]
          }
          x
        } else {
          samples_ppa <- 
            str_split(x$ppa, "/")[[1]] %>%
            as.numeric()
          
          samples_ppa <- 
            tibble(
              company = x$company,
              test_name = x$test_name,
              tp = samples_ppa[1],
              n_pos = samples_ppa[2],
              fn = samples_ppa[2] - samples_ppa[1]
            )
          
          samples_npa <- 
            str_split(x$npa, "/")[[1]] %>%
            as.numeric()
          
          samples_npa <- 
            tibble(
              company = x$company,
              test_name = x$test_name,
              tn = samples_npa[1],
              n_neg = samples_npa[2],
              fp = samples_npa[2] - samples_npa[1]
            )
          
          x <- 
            full_join(x, samples_ppa) %>%
            full_join(samples_npa)
          
          x
        }
        x
      }
      x
    }
  )

# remove rows without clinical performance data
eua_performance <- 
  eua_performance %>%
  filter(!(is.na(tp) | is.na(tn) | is.na(fp) | is.na(fn))) %>%
  mutate(
    sensitivity = tp / n_pos,
    specificity = tn / n_neg,
  )

### Calculate performance data with uncertainty ----

# get confidence intervals for sensitivity and specificity
conf_ints <-
  eua_performance %>%
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

eua_performance <- 
  cbind(eua_performance, conf_ints)

eua_performance$specimen <- str_to_title(eua_performance$specimen)

save(eua_performance, file = "R/data_derived/eua_performance.RData")

