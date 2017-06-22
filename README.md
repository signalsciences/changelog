# changelog
tools to help manage markdown changelogs

The parses and emits a simple markdown-like changelog.

Very little structure is imposed.  Versions are listed using the `##` h2 headline.

A sample might look like the following:

```
# name of product

any text or none

## Unreleased

This section is optional

## 1.2.3 2017-01-01

Headings are `## version-number space date`

No structure is imposed on the version number or date at least for now
```

## CLI

```
$ changelog -help
Usage of changelog:
  -last-entry
    	Show last entry only
  -last-version
    	Show last version only
  -no-comments
    	Error if HTML comments are found
  -no-unreleased
    	Error if an unreleased section is present
```

## Ideas

Pull requests very welcome!

* Validate version using https://github.com/Masterminds/semver
* Validate date using https://golang.org/pkg/time/
    * allow parenthesis around date e.g. `v1.2.3 (2107-06-21)`
    * or more generally, allow a template to parse/emit headers
* Sort by version
* Sort by date
* JSON output
* Handle markdown in markers.  Often version is a markdown link
    ```
    ## [v1.1.16](https://github.com/chef/chef-dk/tree/v1.1.16) (2016-12-14)
    ```

If you want to see other examples of changelogs try this search:
[site:github.com changelog.md]( https://www.google.com/search?&q=site:github.com+changelog.md&ie=UTF-8&oe=UTF-8#q=site:github.com+changelog.md)

