[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/DoctorOgg/sensu-check-http-go)
![goreleaser](https://github.com/DoctorOgg/sensu-check-http-go/workflows/goreleaser/badge.svg)

# sensu-check-http-go

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)

## Overview

This is a simple http check plugin, it does not have all the features of the ruby version.

## Files

* sensu-check-http-go

## Usage examples

```bash
$ sensu-check-http-go -u https://google.com -c "oogle"

OK: https://google.com, status: 200, String found: oogle
```

Help:

```bash
$ sensu-check-http-go -h

A simple replacement for the ruby based http check for sensu

Usage:
  sensu-check-http-go [flags]
  sensu-check-http-go [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --checkstring string   String to Match
  -h, --help                 help for sensu-check-http-go
  -t, --timeout int          Timeout value in seconds (default 10)
  -z, --tlstimeout int       TLS handshake timeout in milliseconds (default 1000)
  -u, --url string           URL to check
```
## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add DoctorOgg/sensu-check-http-go
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/DoctorOgg/sensu-check-http-go].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-check-http-go
  namespace: default
spec:
  command: sensu-check-http-go -u https://google.com -c "oogle" 
  subscriptions:
  - system
  runtime_assets:
  - DoctorOgg/sensu-check-http-go
```

