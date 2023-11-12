# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
Telegram, or any other method with the owners of this repository before making a change.

Please note we have a [code of conduct](CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.

## Pull Request Process

1. Ensure any install or build dependencies are removed before the end of the layer when doing a
   build.
2. Update the README.md with details of changes to the interface, this includes new environment
   variables, exposed ports, useful file locations and container parameters.
3. Increase the version numbers in any examples files and the README.md to the new version that this
   Pull Request would represent. The versioning scheme we use is [SemVer](http://semver.org/).
4. You may merge the Pull Request in once you have the sign-off of two other developers, or if you
   do not have permission to do that, you may request the second reviewer to merge it for you.

## Code hygiene

1. Please install [all dependencies](../README.md#requirements) before starting to work on the project.
2. We're using [Taskfile](https://taskfile.dev/#/) to perform common tasks.
3. Check your code with `task lint` before creating a pull request.
4. Write tests for your code and test it with `task test` before creating a pull request.
5. Please follow the [Uber Go Style Guide](https://github.com/vorobeyme/uber-go-style-guide-uk/blob/master/style.md#%D0%B2%D1%81%D1%82%D1%83%D0%BF) when writing code.
6. Or you can follow the [Effective Go](https://golang.org/doc/effective_go).
7. We're following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.
8. For project structure, please read [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ua.md#%D0%BE%D0%B3%D0%BB%D1%8F%D0%B4) recommendations.
