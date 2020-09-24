# dore-ssg

The Dore static site generator.

This takes content authored in ReStructuredText and converts it to HTML embedded in a technical documentation theme.

## Prerequisites

- Golang
- Docutils

## Directory Structure Dependencies

<p align="left">
    <img src="static/dirStruct.png" alt="Dore Directory Structure"/>
</p>

- The .toml configuration file must be in the /main directory.
- The /src directory contains reStructureText source files and image source files.
- The /pub directory contains published content with presentation files in the /pub/static directory.

## Usage

``` bash
    go run .
```
Or, build binary executable and install in $GOPATH.

## TODO:

- global refactoring
- code comments
- CSS and JavaScript clean up
- improved error detection and handling
- populate additional <meta> fields from .toml configuration file
- document constraints on ReST directives
- document process differences from typical SSGs
- document usability support
- syntax highlighting: map syntactic elements for additional languages
- provide templates for document types
- annunciate build progress
- provide default feedback mechanism
- provide version selection
- add analytics
