# go-basic-i18n

`go-basic-i18n` is a lightweight Go library designed to provide basic internationalization (i18n) support for Go applications. It simplifies the process of managing localized messages across different languages, making it easier for developers to create multi-language applications.

## Features

- Supports loading localized messages from `.properties` files.
- Allows specifying a default language and dynamically switching between languages.
- Provides methods to retrieve localized text based on keys and optionally replace placeholders with parameters.

## Getting Started

To get started with `go-basic-i18n`, ensure you have Go installed on your system. Then, follow these steps:

1. Import the library into your Go project.
2. Initialize the internationalization system with `InitI18n`.
3. Retrieve localized messages using `GetLang` and `GetText` or `GetTextWithParam`.

## Usage

Here's a basic example of how to use `go-basic-i18n` in your application:
    go package main
    
    import ( "fmt"
    basicI18n "github.com/pmperrin/go-basic-i18n/i18n"
    )

    func main() { 
        i18nInstance, err := basicI18n.InitI18n("./path/to/language/files", "fileName", "en") 
        if err!= nil { fmt.Println(err.Error()) return }
    
        lang, err := i18nInstance.GetLang("es")
        if err!= nil {fmt.Println(err.Error())return}
    
        fmt.Println(lang.GetText("message.key"))
    }


Replace `"./path/to/language/files"` with the actual path to your language files, and `"fileName"` with the base name of your `.properties` files.

## Contributing

Contributions to improve the library, add features, or fix bugs are welcome. Please submit pull requests to the main repository.

## License

`go-basic-i18n` is released under the MIT License. See the [LICENSE](LICENSE) file for details.



This .md file has been done by Phind 
