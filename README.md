# k2-crash-application
Application to monitor [K2](https://github.com/samsung-cnct/k2) crashes. 

## Installation

This app runs inside [K2-Tools](https://github.com/samsung-cnct/k2-tools) and in the event of a failure, sends data to an elasticsearch cluster for internal metrics. 

## Usage

This app automatically run if [K2](https://github.com/samsung-cnct/k2) ansible task fails. 

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

K2 does a lot of heavy lifting creating, updating and destroying Kubernetes clusters in the cloud. Failures can occur for a variety of reasons and this tool helps us understand when and where tasks are commonly failing. This tool does not collect any personal information and the crash data is used internally for development purposes. 

## Cutting a release

* Install dependencies:
	-	go get github.com/c4milo/github-release
	-	go get github.com/mitchellh/gox
* Create a [github personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) with repo read/write permissions and export it as GITHUB_TOKEN
* Adjust VERSION and TYPE variables in the [Makefile](Makefile) as needed
* Run ```make release```

## License

[Apache 2.0](https://github.com/samsung-cnct/k2-crash-application/blob/master/LICENSE)