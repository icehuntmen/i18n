# discordgo-i18n
[![GoDoc](https://godoc.org/github.com/kaysoro/discordgo-i18n?status.svg)](https://godoc.org/github.com/kaysoro/discordgo-i18n)
[![Golangci-lint](https://github.com/Kaysoro/discordgo-i18n/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Kaysoro/discordgo-i18n/actions/workflows/golangci-lint.yml)
[![Test](https://github.com/Kaysoro/discordgo-i18n/actions/workflows/test.yml/badge.svg)](https://github.com/Kaysoro/discordgo-i18n/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kaysoro/discordgo-i18n/branch/main/graph/badge.svg)](https://codecov.io/gh/kaysoro/discordgo-i18n) 

discordgo-i18n is a simple and lightweight Go package that helps you translate Go programs into [languages supported by Discord](https://discord.com/developers/docs/reference#LOCALES).

- Built to ease usage of [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)
- Less verbose than [go-i18n](https://github.com/nicksnyder/go-i18n)
- Supports multiple strings per key to make your bot "more alive"
- Supports strings with named variables using [text/template](http://golang.org/pkg/text/template/) syntax
- Supports message files of JSON format

# Getting started

## Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the main branch.*

```sh
go get github.com/icehuntmen/i18n
```

**NOTICE**: this package has been built to ease usage of [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo), it can be used for other projects but will be less practical.

## Usage

Import the package into your project.

```go
import i18n "github.com/icehuntmen/i18n"
```

Load bundles for locales to support.

```go
err := i18n.LoadBundle(discordgo.French, "path/to/your/file.json")
```

The bundle format must respect the schema below; note [text/template](http://golang.org/pkg/text/template/) syntax is used to inject variables.  
For a given key, value can be string, string array to randomize translations or even deep structures to group translations as wanted. In case any other type is provided, it is mapped to string automatically.

```json
{
    "hello_world": "Hello world!",
    "hello_anyone": "Hello {{ .anyone }}!",
    "image": "https://media2.giphy.com/media/Ju7l5y9osyymQ/giphy.gif",
    "bye": ["See you", "Bye!"],
    "command": {
        "scream": {
            "name": "scream",
            "description": "Screams something",
            "dog": "Waf waf! 🐶",
            "cat": "Miaw! 🐱"
        }
    }
}
```

By default, the locale fallback used when a key does not have any translations is `discordgo.EnglishUS`. To change it, use the following method.

```go
i18n.SetDefault(discordgo.ChineseCN)
```
If you need to connect loggger, specify this
```go
var logger  := logrus.New()
i18n.SetLogger(logger)
```

To get translations use the below thread-safe method; if any translation cannot be found or an error occurred even with the fallback, key is returned.

```go
helloWorld := i18n.Get(discordgo.EnglishUS, "hello_world")
fmt.Println(helloWorld)
// Prints "Hello world!"

helloAnyone := i18n.Get(discordgo.EnglishUS, "hello_anyone")
fmt.Println(helloAnyone)
// Prints "Hello {{ .anyone }}!"

helloAnyone = i18n.Get(discordgo.EnglishUS, "hello_anyone", i18n.Vars{"anyone": "Nick"})
fmt.Println(helloAnyone)
// Prints "Hello Nick!"

bye := i18n.Get(discordgo.EnglishUS, "bye")
fmt.Println(bye)
// Prints randomly "See you" or "Bye!"

keyDoesNotExist := i18n.Get(discordgo.EnglishUS, "key_does_not_exist")
fmt.Println(keyDoesNotExist)
// Prints "key_does_not_exist"

dog := i18n.Get(discordgo.EnglishUS, "command.scream.dog")
fmt.Println(dog)
// Prints "Waf waf! 🐶"
```

To get localizations for a command name, description, options or other fields, use the below thread-safe method. It retrieves a `*map[discordgo.Locale]string` based on the loaded bundles.

```go
screamCommand := discordgo.ApplicationCommand{
    Name:                     "scream",
    Description:              "Here a default description for my command",
    NameLocalizations:        i18n.GetLocalizations("command.scream.name"),
    DescriptionLocalizations: i18n.GetLocalizations("command.scream.description"),
}
```

Here an example of how it can work with interactions.

```go
func HelloWorld(s *discordgo.Session, i *discordgo.InteractionCreate) {

    err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
                {
                    Title:       i18n.Get(i.Locale, "hello_world"),
                    Description: i18n.Get(i.Locale, "hello_anyone", i18n.Vars{"anyone": i.Member.Nick}),
                    Image:       &discordgo.MessageEmbedImage{URL: i18n.Get(i.Locale, "image")},
                },
            },
		},
	})

    // ...
}
```

## Contributing

Contributions are very welcomed, however please follow the below guidelines.

- First open an issue describing the bug or enhancement so it can be
discussed.  
- Try to match current naming conventions as closely as possible.  
- Create a Pull Request with your changes against the main branch.

# Licence

discordgo-i18n is available under the MIT license. See the [LICENSE](LICENSE) file for more info.