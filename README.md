# Gitpher
[![ReportCard][ReportCard-Image]][ReportCard-Url]

[ReportCard-Url]: http://goreportcard.com/report/kegsay/gitpher
[ReportCard-Image]: http://goreportcard.com/badge/kegsay/gitpher

*An experimental Go web server for git wrangling*

## [!] Warning [!]
This project is experimental, run at your own risk! The web server exposes raw `git` commands and a Github User Content proxy **without any security**. This is also my first Go, so excuse the mess!

# Aims of this project
- Create an HTTP API which exposes useful functionality from `git`.
- Add additional features (e.g. Comment database) to reach feature parity with Github.
- Provide support for [Git Pull Review](https://github.com/Kegsay/github-pull-review) (GPR) clients.

## Where we're at
 - [x] Proxy Github User Content requests to support GPR file expanding.
 - Reach feature parity with [gitrust](https://github.com/illicitonion/gitrust), the current backend for GPR.
