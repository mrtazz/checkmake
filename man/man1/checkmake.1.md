---
title: checkmake(1) checkmake User Manuals | checkmake User Manuals
author: Daniel Schauenberg <d@unwiredcouch.com>
date: REPLACE_DATE
---

# NAME
checkmake - experimental linter for Makefiles

# SYNOPSIS

```
checkmake [--debug|--config=<configPath>] <makefile>
```

# DESCRIPTION
`checkmake` is an experimental linter for Makefiles. It allows for a set of
configurable rules being run against a Makefile.

# OPTIONS

```
-h --help               Show this screen.
--version               Show version.
--debug                 Enable debug mode
--config=<configPath>   Configuration file to read
--list-rules            List registered rules
```

# CONFIGURATION
By default checkmake looks for a `checkmake.ini` file in the same folder it's
executed in. This can be overridden by passing the `--config=` argument
pointing it to a different configuration file. With the configuration file
the `[default]` section is for checkmake itself while sections named after the
rule names are passed to the rules as their configuration. All keys/values are
hereby treated as strings and passed to the rule in a string/string map.


# SEE ALSO
make(1)
