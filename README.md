# fake-http-server
Server to receive http requests with no responses

[![CodeQL](https://github.com/guionardo/fake-http-server/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/guionardo/fake-http-server/actions/workflows/codeql-analysis.yml)

[![Publish Docker image](https://github.com/guionardo/fake-http-server/actions/workflows/publish_docker.yml/badge.svg)](https://github.com/guionardo/fake-http-server/actions/workflows/publish_docker.yml)

## Usage

Pull image from the command line:
```bash
$ docker pull docker.pkg.github.com/guionardo/fake-http-server/fake-http-server:0.3.2
```

Use as base image in DockerFile:
```
FROM docker.pkg.github.com/guionardo/fake-http-server/fake-http-server:0.3.2
```
