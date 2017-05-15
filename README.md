# try

Try is a command-line utility that makes it easy to test-drive other people's projects without side effects.

Installation:

    $ go get github.com/zachlatta/try/cmd/try

Usage:

    $ try [repo]
    
Try will clone the repo, set up the environment, install dependencies, and then drop you into a shell where you can play around. Once you're done, exit the shell and everything will be deleted.

Examples:

    $ try https://github.com/ggbrw/boolr
    $ try https://github.com/shivammathur/IPpy
    $ try https://github.com/maxhallinan/my-clippings-to-json
    
## Why doesn't try use Docker?

If you're like me, you've spent a bunch of time customizing your development environment and you want access to your custom dotfiles, non-traditional shell, and your editor when playing around with new repositories.

Try opts for using temp directories over Docker so you have access to all of your tooling without having to re-install anything.

## License

Try is licensed under the MIT license. See [`LICENSE`](LICENSE) for the full text.
