## plz-go

A shell command to execute standard/repeatable commands in a git repo

This is a rewrite of a the [plz-cmd](https://github.com/m3brown/plz) Python project.

### Installation

TODO

### Example

plz looks for a `.plz.yaml` file either in the current directory or in the root
of the git repo you're currently in. This file can (and should) be checked into
version control.

For a .plz.yaml file located in the git root directory, commands run will be
executed relative to that directory, not the current directory.

Suppose we have the following `.plz.yaml` file:

```yaml
- id: run
  cmd: ./manage.py runserver
- id: test
  cmds:
  - ./manage.py test
  - yarn test
- id: setup
  cmds:
  - poetry install
  - poetry run ./manage.py migrate
  - yarn install
- id: ls
  cmd: ls
```

The following commands would be available:

```bash
plz run
plz test
plz setup
```

### Globbing

TODO globbing is not supported yet.

For example, the cmd `ls *.py` will not work as expected if you run the command
from a nested directory inside your git repo.

### Runtime arguments

plz supports passing custom arguments when running the plz command. For example:

```
# bind to port 8001 instead of the default 8000
plz run 127.0.0.1:8001
```

Note: arguments are only passed through for singular `cmd` definitions. When
using the plural `cmds` array, extra command line arguments are ignored.

### Development

TODO
