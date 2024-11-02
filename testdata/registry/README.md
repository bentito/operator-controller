# Test Registry

This tool is a bare-bones image registry using the `go-containerregistry` library; it is intended to be used in a test environment only. 

Usage:
```
Usage of registry:
      --registry-address string   The address the registry binds to. (default ":12345")
```

<!--TODO makes these arguments-->
The server key and cert should be placed in:
```
	certPath = "/certs/tls.crt"
	keyPath  = "/certs/tls.key"
```
