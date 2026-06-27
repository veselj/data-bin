output "endpoint_url" {
  description = "Custom-domain HTTPS endpoint"
  value       = "https://${var.domain_name}"
}

output "cloudfront_domain" {
  description = "CloudFront distribution domain (useful before DNS propagates)"
  value       = aws_cloudfront_distribution.api.domain_name
}

output "lambda_function_url" {
  description = "Raw Lambda Function URL (bypass CloudFront for debugging)"
  value       = aws_lambda_function_url.api.function_url
}

output "s3_bucket" {
  description = "Name of the S3 bucket storing payloads"
  value       = aws_s3_bucket.data.id
}

output "lambda_function_name" {
  description = "Lambda function name"
  value       = aws_lambda_function.api.function_name
}
