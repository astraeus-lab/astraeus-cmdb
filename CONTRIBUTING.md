
# Contributing

## Issues

It is a great way to contribute by reporting an issue. Well-written and complete bug reports are always welcome! Please open an issue and follow the template to fill in required information.

Before opening any issue, please look up the existing issues to avoid submitting a duplication. If you find a match, you can "subscribe" to it to get notified on updates. If you have additional helpful information about the issue, please leave a comment.

- With issues:
    - Which version you are using.
    - Steps to reproduce the issue.
    - Snapshots or log files if needed.

Note: Because the issues are open to the public, when submitting files, be sure to remove any sensitive information, e.g. user name, password, IP address, and company name. You can
replace those parts with "REDACTED" or other strings like "****".

## Pull Request

PR are always welcome, even if they only contain small fixes like typos or a few lines of code. If there will be a significant effort, please document it as an issue and get a discussion going before starting to work on it.

Please submit a PR broken down into small changes bit by bit. A PR consisting of a lot features and code changes may be hard to review. It is recommended to submit PRs in an incremental fashion.

- With pull requests:
    - Open your pull request against `master`
    - Your pull request should have no more than two commits, if not you should squash them.
    - It should pass all tests in the available continuous integration systems such as GitHub Actions.
    - If your pull request contains a new feature, please document it on the README.
    - You should add/modify tests to cover your proposed code changes.

Note: If you split your pull request to small changes, please make sure any of the changes goes to master will not break anything. Otherwise, it can not be merged until this feature complete.

### Coding style

* See [Google Style](https://google.github.io/styleguide/go/).
