# reputation-checker
@author: winstark212  
@date: 2019-10-23

Reputation checker is a threat intelligence tool. User can check IPs, domains, URLs are apparent in compromised list or not from different sources.

# Features  



# Feeds
## Common feeds for IP

* SuricataCompromised
* SuricataBotCC
* SurricataTor
* SnortIPFilter

## Common feeds for domain
* MalwareDomainHosts
* MandiantAPT

# Usage  

```  
# windows
go build -o reputation.exe cmd\reputation\main.go  
# linux  
go build -o reputation cmd/reputation/main.go
```  