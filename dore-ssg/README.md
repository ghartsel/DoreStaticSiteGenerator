# dore-ssg

The Dore static site generator.

This takes content authored in ReStructuredText and converts it to HTML embedded in a technical documentation theme.

## Prerequisites

- Golang
- Docutils

## Usage

``` bash
    go run .
```

tbd

## Notes

- tbd
- tbd
- tbd

## TODO:

- search
  - return search page w/ ?highlight=<word> param
  - preview result
    - save 3 words before and after hit word
- server
  - form handler
  - security checks
  - custom 404 page
- only regenerate changed topics
- improved function comments
- cleanup css & js
- better error detection & handling
- populate <meta> fields from .toml
- document constraints on ReST directives
- document process differences
  - no auto-build on source change (not a blog)
  - ssg dependencies (Pygments?, docutlls, ...)
- syntax highlighting: map additional syntactic elements
- doctype templates (spec, UG, cookbook, API/SDK ref)
- progress notifications
- publish to aws, netlify, github pages
- feedback (discus, or ...)
- version selection
- add Google Analytics
