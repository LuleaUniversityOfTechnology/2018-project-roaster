## Development guidelines
### Lint and format
Before you commit, the affected files should be run through `goimports`. Bonus: Also run `go vet`.

 * goimports: `go get golang.org/x/tools/cmd/goimports`

### Tests
Write tests for everything.

### Branches
Develop inside your own branch until you're done, then send a pull request.

### Commits
Commits should be formatted with primary affected package as prefix, a short descriptive one liner and then an optional description of the context and what the change does. Also use GitHub's `Fixes #123` feature for closing issues. All commits must be signed with GPG.

Example:
```
math: improve Sin, Cos and Tan precision for very large arguments

The existing implementation has poor numerical properties for
large arguments, so use the McGillicutty algorithm to improve
accuracy above 1e10.

The algorithm is described at http://wikipedia.org/wiki/McGillicutty_Algorithm

Fixes #159

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
# On branch foo
# Changes not staged for commit:
#	modified:   editedfile.go
#
```

### Code style
Take a look at: https://github.com/golang/go/wiki/CodeReviewComments
