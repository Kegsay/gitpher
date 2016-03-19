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
 - [ ] Design a workspace for checking out and maintaining arbitrary git projects.
 - [ ] Reach feature parity with [gitrust](https://github.com/illicitonion/gitrust), the current backend for GPR:
    - [ ] Implement OAuth2 for Github
    - [ ] `/squashmerge` API
    - [ ] `/rewritehistory` API
 - [x] Proxy Github User Content requests to support GPR file expanding.
 - [ ] Implement and use a better diffing algorithm than `git diff`.

### Stretch goals
These goals work to replace GPRs dependence on Github. Its aim is to make Gitpher expose the same APIs as Github so clients can be pointed to Gitpher instead.
 - [ ] Add a database layer which can store comments, pull requests, etc
 - [ ] Support comments
 - [ ] Support creation of pull requests
