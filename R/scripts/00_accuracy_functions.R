# Functions below rely on the tidyverse package
library(tidyverse)

#' Calculate a confidence interval of a proportion based on Wilson's score
#' @description
#'   Calculate a confidence interval of a proportion based on Wilson's score.
#'   The traditional Wald interval fails when you have close to 100% or 0%
#'   successes. The Wilson interval is robust to these extremes and is also
#'   appropriate to use in other cases where you could use the Wald interval.
#' @param ns integer number of observed successes
#' @param n integer total number of trials. Note \code{ns <= n}.
#' @param z the z-score for the appropriate confidence level. Defaults to \code{1.96}
#'   for a 95% confidence interval.
#' @return Returns an integer vector with two entries. The first is the lower
#'   bound of the CI. The second is the upper bound.
#' @examples
#' # calculate a 95% CI 
#' calc_wilson_ci(ns = 30, n = 50, z = 1.96)
calc_wilson_ci <- function(
  ns, 
  n, 
  z = 1.96
) {
  # check inputs
  if (! ns <= n)
    stop("ns cannot be larger than n.")
  
  if (z < 0) {
    z <- abs(z)
    
    message("converting z to positive so that returned CI is in the right order")
  }
  
  # nf is the number of failures
  nf <- n - ns
  
  # CI has two terms. Lower bound of CI is a - b. Upper bound of CI is a + b
  a <- (ns + ((z ^ 2) / 2)) / (n + z ^ 2)
  
  b <- (z / (n + z ^ 2)) * sqrt((ns * nf) / n + (z ^ 2) / 4)
  
  # return result. A vector with two elements [lower bound, upper bound]
  c(a - b, a + b)
  
}

#' Calculate positive predictive value
#' @description Calculate positive predictive value (PPV) over a range of disease
#'   prevalence.
#' @param x numeric vector whose entries range beween 0 and 1; disease prevalence.
#' @param sensitivity between 0 and 1; sensitivity of the test.
#' @param specificity between 0 and 1; specificity of the test.
#' @return vector of \code{length(x)} whose values represent the PPV over the 
#'   range of \code{x}
#' @examples
#' x <- seq(0.005, 1, by = 0.005)
#' 
#' ppv <- calc_ppv(x = x, sensitivity = 0.95, specificity = 0.99)
calc_ppv <- function(
  x,
  sensitivity,
  specificity
) {
  # note that this function is vectorized over x, but not sensitivity and specificity
  
  # check inputs
  if (min(x) < 0 | max(x) > 1)
    stop("all entries of x must be between 0 and 1")
  
  if (length(specificity) != 1)
    stop("specificity must be length 1")
  
  if(length(sensitivity) != 1)
    stop("sensitivity must be length 1")
  
  if (sensitivity < 0 | sensitivity > 1)
    stop("sensitivity must be between 0 and 1")
  
  if (specificity < 0 | specificity > 1)
    stop("specificity must be between 0 and 1")
  
  
  # Calculate numerator (a) and denomenator (b) separately to preserve 
  # vectorization over x
  a <- sensitivity * x
  
  b <- sensitivity * x + (1 - specificity) * (1 - x)
  
  # return the result, a vector as long as x
  a / b
}

#' Calculate negative predictive value
#' @description Calculate negative predictive value (NPV) over a range of disease
#'   prevalence.
#' @param x numeric vector whose entries range beween 0 and 1; disease prevalence.
#' @param sensitivity between 0 and 1; sensitivity of the test.
#' @param specificity between 0 and 1; specificity of the test.
#' @return vector of \code{length(x)} whose values represent the NPV over the 
#'   range of \code{x}
#' @examples
#' x <- seq(0.005, 1, by = 0.005)
#' 
#' npv <- calc_npv(x = x, sensitivity = 0.95, specificity = 0.99)
calc_npv <- function(
  x,
  sensitivity,
  specificity
) {
  # note that this function is vectorized over x, but not sensitivity and specificity
  
  # check inputs
  if (min(x) < 0 | max(x) > 1)
    stop("all entries of x must be between 0 and 1")
  
  if (length(specificity) != 1)
    stop("specificity must be length 1")
  
  if(length(sensitivity) != 1)
    stop("sensitivity must be length 1")
  
  if (sensitivity < 0 | sensitivity > 1)
    stop("sensitivity must be between 0 and 1")
  
  if (specificity < 0 | specificity > 1)
    stop("specificity must be between 0 and 1")
  
  # Calculate numerator (a) and denomenator (b) separately to preserve 
  # vectorization over x
  a <- specificity * (1 - x)
  
  b <- specificity * (1 - x) + (1 - sensitivity) * x
  
  # return the result, a vector as long as x
  a / b
}


#' Calculate a table of posterior probabilities 
#' @description
#'   Calculate a table of posterior probabilities (and 95% confidence intervals)
#'   for diagnostic test trial results.
#' @param tp integer true positives
#' @param fp integer false positives
#' @param tn integer true negatives
#' @param fn integer false negatives
#' @return a \code{\link[tibble]{tibble}} with the following columns
#' 
#'   prevalence sequence from 0.005 to 1 representing disease prevalence or
#'   prior probability of infection
#'   
#'   ppv_low, ppv_high, and ppv are the lower and upper bounds of a 95% confidence
#'   interval and point value of positive predictive value.
#'   
#'   npv_low, npv_high, and npv are the lower and upper bounds of a 95% confidence
#'   interval and point value of negative predictive value.
#'   
#'   p_pos_low, p_pos_high, and p_pos are the lower and upper bounds of a 95% confidence
#'   interval and point value of the percent of tests coming back positive.
#'   
#'   fdr_low, fdr_high, and fdr are the lower and upper bounds of a 95% confidence
#'   interval and point value of the false discovery rate or 1 - ppv.
#'   
#'   fomr_low, fomr_high, and fomr are the lower and upper bounds of a 95% confidence
#'   interval and point value of the false ommission rate or 1 - npv.
#' @examples
#' mat <- calc_plot_matrix(tp = 30, fn = 1, tn = 802, fp = 3)
calc_plot_matrix <- function(
  tp,
  fn,
  tn,
  fp
) {
  
  # check inputs
  if (length(tp) != 1 | length(fn) != 1 | length(tn) != 1 | length(fp) != 1)
    stop("tp, fn, tn, and fp must all be integers of length 1")
  
  # point estimate for sensitivity is tp / (tp + fn)
  sensitivity <- calc_wilson_ci(ns = tp, n = tp + fn, z = 1.96)
  
  # point estimate for specificity is tn / (tn + fp)
  specificity <- calc_wilson_ci(ns = tn, n = tn + fp, z = 1.96)
  
  # calculate disease prevalence (or prior probability of infection) from 0.005 to 1
  x <- seq(0.005, 1, by = 0.005)
  
  
  # Get prevalence, ppv, npv, and CIs for each into a single table
  sens_point <- tp / (tp + fn)
  
  spec_point <- tn / (tn + fp)
  
  plot_matrix <- 
    tibble(
      prevalence = round(x, 10), # weird rounding errors at many decimal places
      ppv_low = calc_ppv(x, sensitivity[1], specificity[1]),
      ppv_high = calc_ppv(x, sensitivity[2], specificity[2]),
      ppv = calc_ppv(x, sens_point, spec_point),
      npv_low = calc_npv(x, sensitivity[1], specificity[1]),
      npv_high = calc_npv(x, sensitivity[2], specificity[2]),
      npv = calc_npv(x, sens_point, spec_point),
      p_pos_low = sensitivity[2] * x + (1 - specificity[2]) * (1 - x),
      p_pos_high = sensitivity[1] * x + (1 - specificity[1]) * (1 - x),
      p_pos = sens_point * x + (1 - spec_point) * (1 - x)
    ) %>% 
    mutate(
      fdr_low = 1 - ppv_high,
      fdr_high = 1 - ppv_low,
      fdr = 1 - ppv,
      fomr_low = 1 - npv_high,
      fomr_high = 1 - npv_low,
      fomr = 1 - npv
    )
  
  # return the result
  plot_matrix
  
}


#' Get additional test metadata for display alongside a plot
#' @description
#'   Get additional info to display about a test alongside plots of its accuracy.
#'   Designed to extract the right info once a test is selected for display in
#'   a side pannel.
#' @param test_row a single row of a tibble with appropriate variables. See
#'   details, below.
#' @param manufacturer_col character of the name of the variable in \code{test_row}
#'   that contains the manufacturer name
#' @param test_name_col character of the name of the variable in \code{test_row}
#'   that contains the test name
#' @param other_keep_cols character vector containing the names of other variables
#'   in \code{test_row} to keep for display
#' @details 
#'   \code{test_row} needs at a minimum variables for \code{manufacturer_col},
#'   \code{test_name_col}, and the following columns: \code{true_positive},
#'   \code{true_negative}, \code{false_positive}, \code{fales_negative}
get_static_test_info <- function(
  test_row,
  manufacturer_col,
  test_name_col,
  other_keep_cols
) {
  
  # check inputs
  # must be single row tibble
  # must have tp, fp, etc. cols
  # manufacturer_col and test_name_col must exist
  
  
  out <- 
    test_row %>%
    mutate(
      sensitivity = true_positive / (true_positive + false_negative),
      specificity = true_negative / (true_negative + false_positive),
      n = true_positive + true_negative + false_positive + false_negative,
      n_pos = true_positive + false_negative,
      n_neg = true_negative + false_positive
    ) %>% 
    select(
      all_of(manufacturer_col),
      all_of(test_name_col),
      sensitivity,
      specificity,
      n, 
      n_pos,
      n_neg,
      all_of(other_keep_cols)
    )
  
  out
  
}