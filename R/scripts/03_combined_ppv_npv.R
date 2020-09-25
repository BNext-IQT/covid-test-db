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

find_formatted$source <- "FindDx"

eua_performance$source <- "FDA EUA"


### Reconcile company names across our data sources ----
# This is a manual process using a case_when statement
find_formatted <-
  find_formatted %>%
  mutate(
    company = case_when(
      company == "Abbott Diagnostics" ~ "Abbott Diagnostics Scarborough, Inc.",
      company == "Abbott Diagnostics Inc." ~ "Abbott Diagnostics Scarborough, Inc.",
      company == "Abbott Molecular" ~ "Abbott Molecular Inc.",
      company == "altona Diagnostics" ~ "altona Diagnostics GmbH",
      company == "DiaSorin Molecular, LLC" ~ "DiaSorin Molecular LLC",
      company == "GenMark Diagnostics" ~ "GenMark Diagnostics, Inc.",
      company == "GenMark Diagnostics, Inc" ~ "GenMark Diagnostics, Inc.",
      company == "Hologic" ~ "Hologic, Inc.",
      company == "Jiangsu Bioperfectus Technologies Co. Ltd" ~ "Jiangsu Bioperfectus Technologies Co., Ltd.",
      company == "Roche Molecular Diagnostics" ~ "Roche Molecular Systems, Inc. (RMS)",
      company == "Sansure Biotech, Inc." ~ "Sansure BioTech Inc.",
      TRUE ~ company
    )
  )



### Reconcile tests with multiple trials ----
# Append FindDx {x} to the test name to differentiate multiple trials and flag
# that the source of this is FindDx vs FDA EUA

find_formatted <- by(
  find_formatted, 
  INDICES = paste(find_formatted$company, find_formatted$test_name), 
  function(x) x
)

find_formatted <- lapply(
  find_formatted,
  function(x) {
    x$test_name <- paste(x$test_name, "- Source: FindDx", seq_along(x$test_name))
    x
  }
)

find_formatted <- do.call(rbind, find_formatted)

### Combine find with eua data ----
# as of now filtering out antibody tests
combined_performance <- rbind(
  eua_performance,
  filter(find_formatted, test_type != "Antibody")
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
        test_name = x$test_name,
        ppv = round(ppv, 5),
        ppv_low = round(ppv_low, 5),
        ppv_high = round(ppv_high, 5),
        npv = round(npv, 5),
        npv_low = round(npv_low, 5),
        npv_high = round(npv_high, 5),
        p_pos = round(p_pos, 5),
        p_pos_low = round(p_pos_low, 5),
        p_pos_high = round(p_pos_high, 5),
        fdr = round(fdr, 5),
        fdr_low = round(fdr_low, 5),
        fdr_high = round(fdr_high, 5),
        fomr = round(fomr, 5),
        fomr_low = round(fomr_low, 5),
        fomr_high = round(fomr_high, 5)
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