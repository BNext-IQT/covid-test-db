################################################################################
# Eventually this may combined multiple non-covid tests
# This script formats a plot matrix for a rapid influenza diagnostic test.
# The sensitivity/specificity numbers are hard coded here. But the source is
# https://pubmed.ncbi.nlm.nih.gov/22371850/
################################################################################


### Source accuracy diagnostic scripts ----
source("R/scripts/00_accuracy_functions.R")

# source of influenza https://pubmed.ncbi.nlm.nih.gov/22371850/
x <- seq(0.005, 1, by = 0.005) %>% round(10)

flu_sensitivity_point <- .623

flu_sensitivity <- c(.579, .666)

flu_specificity_point <- .982

flu_specificity <- c(.975, .987)

flu_plot_mat <- tibble(
  prevalence = x,
  ppv = calc_ppv(
    x = x, 
    sensitivity = flu_sensitivity_point, 
    specificity = flu_specificity_point),
  ppv_low = calc_ppv(
    x = x, 
    sensitivity = flu_sensitivity[1], 
    specificity = flu_specificity[1]),
  ppv_high = calc_ppv(
    x = x, 
    sensitivity = flu_sensitivity[2], 
    specificity = flu_specificity[2]),
  npv = calc_npv(
    x = x, 
    sensitivity = flu_sensitivity_point, 
    specificity = flu_specificity_point),
  npv_low = calc_npv(
    x = x, 
    sensitivity = flu_sensitivity[1], 
    specificity = flu_specificity[1]),
  npv_high = calc_npv(
    x = x, 
    sensitivity = flu_sensitivity[2], 
    specificity = flu_specificity[2]),
  p_pos = flu_sensitivity_point * x + (1 - flu_specificity_point) * (1 - x),
  p_pos_low = flu_sensitivity[2] * x + (1 - flu_specificity[2]) * (1 - x),
  p_pos_high = flu_sensitivity[1] * x + (1 - flu_specificity[1]) * (1 - x)
)  %>% 
  mutate(
    fdr = 1 - ppv,
    fdr_low = 1 - ppv_high,
    fdr_high = 1 - ppv_low,
    fomr = 1 - npv,
    fomr_low = 1 - npv_high,
    fomr_high = 1 - npv_low
  )

save(
  flu_plot_mat,
  file = "R/data_derived/flu_calculated.RData"
)

write_csv(
  x = flu_plot_mat,
  path = "R/data_derived/flu_plot_mat.csv",
  na = ""
)


