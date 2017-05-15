# try

Try is a command-line utility that makes it easy to test-drive other people's projects without side effects.

Usage:

    $ try [repo]
    
Try will clone the repo, set up the environment, install dependencies, and then drop you into a shell where you can play around. Once you're done, exit the shell and everything will be deleted.

Examples:

    $ try https://github.com/ggbrw/boolr
    $ try https://github.com/shivammathur/IPpy
    $ try https://github.com/maxhallinan/my-clippings-to-json
    
Currently supported environments:

- Node / JavaScript

If you'd like to add support for another environment, create a new strategy in [`runner/strategy/`](runner/strategy/). See [`runner/strategy/node.go`](runner/strategy/node.go) for an example.

Installation:

    $ go get github.com/zachlatta/try/cmd/try

## Why doesn't try use Docker?

If you're like me, you've spent a bunch of time customizing your development environment and you want access to your custom dotfiles, non-traditional shell, and your editor when playing around with new repositories.

Try opts for using temp directories over Docker so you have access to all of your tooling without having to re-install anything.

## How do you capitalize this thing?

Try should be capitalized at the beginning of sentences, but lowercased everywhere else.

Examples:

> Why doesn't try support Docker?

> Try doesn't support Docker because of the reason explained above.

## Feature Wishlist

Run `git ls-files | xargs grep "TODO"` to get an up-to-date feature wishlist for try.

## License

Try is licensed under the MIT license. See [`LICENSE`](LICENSE) for the full text.
