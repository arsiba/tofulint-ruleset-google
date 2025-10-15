plugin "terraform" {
  enabled = false
}

plugin "google" {
  enabled = true
  version = "0.0.1"
  source = "github.com/arsiba/tofulint-ruleset-google"
}
