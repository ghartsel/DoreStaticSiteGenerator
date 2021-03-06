# Welcome to the *DoreStaticSiteGenerator*

The Dore static site generator is inspired by Hugo, for implementation efficiency, and Sphinx, for a documentation-centered authoring environment. Using structure, convention, and semantic richness, Dore encourages authoring at a higher level to the benefit of users. Dore is not intended as a replacement for existing static site generators that might fit a particular niche but is expected to meet the documentation needs of most projects with much less effort while ensuring much higher quality.

See the mock [demoyard](https://ghartsel.github.io/demoyard/) documentation for an example of the content generated by Dore. The Dore static site generator, itself, can be viewed [here](https://github.com/ghartsel/DoreStaticSiteGenerator/tree/master/dore-ssg).

## Features

- Written in Go for fast document generation and powerful templating.
- Uses .toml markup as a clean, easy-to-read, and ideally suited configuration specification.
- Server-hosted documentation for secure hosting and host-anywhere publishing.
- Complete separation of content from metadata and styling - separation of authoring from presentation.
- Supports and encourages topic-based authoring.
- Fully customizable functionality and styling.
- Integrated search with up-to-date index generation.

## Repository Content

### ReSTCheatSheet

The [ReSTCheatSheet](https://github.com/ghartsel/DoreStaticSiteGenerator/tree/master/ReSTCheatSheet) folder presents the ReStruturedText markup and directives supported by Dore.

### demo/pub

The [demo](https://github.com/ghartsel/DoreStaticSiteGenerator/tree/master/demo) folder contains the source content for the [demoyard](https://ghartsel.github.io/demoyard/) example.

### dore-ssg

The [dore-ssg](https://github.com/ghartsel/DoreStaticSiteGenerator/tree/master/dore-ssg) folder contains the source code for the Dore static site generator.

### server

The [server](https://github.com/ghartsel/DoreStaticSiteGenerator/tree/master/server) folder contains the server source code, which serves the published content.
