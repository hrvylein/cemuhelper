# cemuhelper

A little helper that lets you start your favourite games with different settings.bin for CEMU from Steam Big Picture

## Getting Started

Clone the repo to your hard drive or grab the compiled binary from releases

### Prerequisites

To compile it for yourself you would need to install go

### Installing

If you downloaded the binary just place it in the Cemu base folder and create the directory "settings".

If you cloned the repo, run "go build", copy the cemuhelper.exe to your Cemu base directory and create the directory "settings".

Start CEMU and change settings according to your game. Quit CEMU and copy settings.bin to settings/ with a specific name.

### Running

When running from CMD you would use it like f.e.
``cemuhelper.exe -file botw -game "Y:\Game\Game.rpx"``

Remember to not include .bin in the name of your settings file. Cemuhelper will add it.

You can now add cemuhelper.exe as a non-steam game to Steam and set the launch options accordingly.

No warranty given, possible data loss is on your responsibility.

## Authors

* **hrvylein** - *Initial work* - [PurpleBooth](https://github.com/hrvylein)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
