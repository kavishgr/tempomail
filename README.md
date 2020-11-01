# tempomail

**`tempomail`** is a standalone binary that allows you to create a temporary `email address` in **1 Second** and receive emails. It uses 1secmail's [API](https://www.1secmail.com/api/). No dependencies required!

> Insert GIF 

# Installation

### From Binary

Download the pre-built binaries for different platforms from the [releases](https://github.com/kavishgr/tempomail/releases/) page. Extract them using tar, move it to your `$PATH` and you're ready to go.

```sh
▶ # download release from https://github.com/kavishgr/tempomail/releases/
▶ tar -xzvf linux-amd64-tempomail.tgz
▶ mv tempomail /usr/local/bin/
▶ tempomail -h
```


### From Github

```sh
git clone https://github.com/kavishgr/tempomail.git
cd tempomail
go build .
mv tempomail /usr/local/bin/
tempomail -h
```

# Usage

By default, all emails are saved in **/tmp/1secmails/**. It only has one flag `--path` to specify a directory to store your emails:

```
Usage of tempomail:
  -path string
    	specify directory to store emails (default "/tmp/1secmails/")
```

## Does it need improvement ?

Open an issue.
