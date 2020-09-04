# DoreStaticSiteGenerator

The Dore static site generator is written in Go and inspired by Hugo, for implementation efficiency, and Sphinx, for documentation-centric technical content authoring and deployment. Through structure and convention, Dore attempts to promote good documentation practices.

## Features

- tbd
- tbd
- tbd

## Motivation

- tbd
- tbd
- tbd

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
- nav accordion
  - tempermental
  - highlight/expand on focus
- only xform changed topics
- improved function comments
- cleanup css & js
- better error detection & handling
- populate <meta> fields from .toml
- document the constraints on ReST directives
- document process differences
  - no auto-build on source change (not a blog)
  - ssg dependencies (Pigments?, docutlls, ...)
- syntax highlighting: map additional syntactic elements
- doctype templates extension
- progress notifications
- publish to aws or netlify
- feedback (discus)
- version selection support
