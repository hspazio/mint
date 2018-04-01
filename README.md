# Mint - A simple note manager written in Go

## Usage

After installation:

```bash
# create a configuration file ~/.mint.json and a default 
# location where to keep the notes: ~/.mint/
mint init

# edit a note or create it if it does not exist
# Mint will open the note in the default editor
mint note todo

# alternatively you can let Mint listing all available notes
# and you can choose which one to edit
mint note # >> choose a note from the prompt

# list all available notes
mint ls

# remove note
mint rm todo
```

## Configure

Mint stores configurations in a `$HOME/.mint.json` file.

```javascript
{
  "dir": "/Users/fabio/.mint",
  "editor": "vim"
}
```

Configurations can be edited using the `mint config <key> <value>` command:

```bash
mint config dir "/User/fabio/Dropbox/mint"

mint config editor code
```

And `mint config` to inspect the configurations

## Contributing

bug reports and pull requests are very welcome. Please be aware that you are expected to follow the [code of conduct](https://github.com/hspazio/mint/blob/master/CODE_OF_CONDUCT.md).

## License

Copyright (c) 2018 Fabio Pitino, released under the [MIT license](http://www.opensource.org/licenses/MIT).

