terraform {
    required_version = ">= 1.6.0"
    required_providers {
        google = {
            version = "~> 7.7.0"
            source = "hashicorp/google"
        }
    }
}