# vc
* Validate some Certificates
* Horribly simple

## How it Works
* Initiates a TLS client session
* golang will by default fail a connection that has a terrible certificate

## Get Setup
* `brew install go`
* Setup your go environment, utilize golang docs for this
* in your go `src` path: `git clone git@jtslear/vc`

## Execution
### Build it
* `go build`

### Run it
* `vc [hostname:port]`
