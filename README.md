# suru

Suru (`する`) is a simple tool to help manage and measure tasks.  It
lets you take ownership of your workflows and optionally publish logs of
what you're working on in various formats (see below for an example!)

[![Status](.suru/status.png)](.suru)

You can also use it with other people in real time to focus chat and
collab online around "Topics."

## Installing Suru

For now, Suru should always be built from the latest source.

First, install [Go](https://golang.org/doc/install).  Then, open your
terminal and run:

```sh
go get -u -v gopkg.in/suru.v0/...
```

Assuming your `PATH` is set up to include the `bin` directory under your
`GOPATH`, the `suru` command will now be available in your terminal.

## Using Suru

Suru has an interactive CLI you can use to manage your Topics, or you
can run the built-in webservice.

```sh
> suru help
# ...

> suru help config
# ...

> suru help topic
# ...

> suru help live
# ...

> suru help serve
# ...

> suru help pub
# ...

> suru help sub
# ...
```

### suru config

`suru config` configures Suru.  Use `suru config gen` to generate a new
config.

### suru topic

`suru topic` manages Topics.  Topics can have discussions, tasks, and
more.  Use `suru topic <name>` to switch to a Topic.  The default Topic
is `latest`, which selects the Topic having the most recent updates or
messages.

### suru live

`suru live` runs an interactive shell session where you can see live
updates.  Use the `?` key to see a help menu and `q` to quit.

### suru serve

`suru serve` hosts a webservice which responds to gRPC messages.  The
[Protocol Buffers](https://developers.google.com/protocol-buffers/)
schema file is located in the [protocol](protocol) subdirectory.

### suru pub

`suru pub` publishes a summary of your projects/topics, and can
optionally export a URL where others can view or participate in your
projects/topics.  See `suru help pub` for more details.

### suru sub

`suru sub` subscribes to updates from another user's Suru pubfeed.  When
this is active, it can receive messages from other users.  Make sure you
are comfortable receiving messages from users on the server you're
connecting to.

## Why "Suru"

https://en.wiktionary.org/wiki/%E3%81%99%E3%82%8B#Japanese

(Plus it sounds like `sudo`.)
