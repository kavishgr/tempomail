# tempomail

**`tempomail`** is a standalone binary that allows you to create a temporary `email address` in **1 Second** and receive emails. It uses 1secmail's [API](https://www.1secmail.com/api/). No dependencies required!

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

# Does it need improvement ?

Open an issue.

## Support and Contributions

<a href="https://www.buymeacoffee.com/kavishgr" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
