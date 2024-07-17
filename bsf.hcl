
packages {
  development = ["delve@1.22.0", "go-task@~3.37.2", "go@1.21.6", "golangci-lint@~1.59.1", "gotools@0.16.1"]
  runtime     = ["cacert@3.95"]
}

gomodule {
  name    = "bsf"
  src     = "./."
  ldFlags = null
  tags    = null
  doCheck = false
}

githubRelease "bsf" {
  owner = "buildsafedev"
  repo  = "bsf"
  dir   = ""
}
