# export-env

## build

```
go get -v code.cloudfoundry.org/cli
go build export-env.go
```

## install

```
cf install-plugin export_env
```


## uninstall

```
cf uninstall-plugin ExportEnvPlugin
```


## usage

```
cf export-env
cf ee
```


## env

| env | value |
|:-----|:-----|
| CF_USER_NAME | user name |
| CF_USER_GUID | user guid |
| CF_USER_EMAIL | user email |
| CF_API_ENDPOINT | api endpoint |
| CF_API_VERSION| api version |
| CF_LOGGREGATOR_ENDPOINT | loggregator endpoint |
| CF_DOPPLER_ENDPOINT | doppler endpoint |
| CF_ACCESS_TOKEN | access token |


## notice

Mac only confirmation
