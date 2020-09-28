### Load libraries ----
library(tidyverse)
library(janitor)
library(readxl)

### Source accuracy diagnostic scripts ----
source("R/scripts/00_accuracy_functions.R")

d <- GET("https://covid-19-search-diagnostics-db.bnext.org/api/diagnostics") %>% 
  .[["content"]] %>%
  rawToChar() %>% 
  fromJSON() %>%
  as_tibble()

