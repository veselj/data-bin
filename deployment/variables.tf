variable "aws_region" {
  description = "AWS region to deploy into"
  type        = string
  default     = "eu-west-1"
}

variable "environment" {
  description = "Deployment environment (e.g. dev, prod)"
  type        = string
  default     = "dev"
}

variable "domain_name" {
  description = "Full custom domain for the API"
  type        = string
  default     = "data-bin.laetus.uk"
}

variable "root_domain" {
  description = "Root domain used to look up the Route 53 hosted zone"
  type        = string
  default     = "laetus.uk"
}

variable "s3_bucket_name" {
  description = "Name of the S3 bucket used to store payloads"
  type        = string
  default     = "data-bin.laetus.uk"
}

variable "lambda_zip_path" {
  description = "Path to the compiled Lambda zip archive"
  type        = string
  default     = "../build/data-bin.zip"
}
