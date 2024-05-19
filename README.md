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
3. Select the current lang using `GetLang`
4. Retrieve localized messages using `GetText` or `GetTextWithParam`.


## Usage

Here's a basic example of how to use `go-basic-i18n` in your application:
    go package main
    
    import ( 
        "fmt"
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


## Example
You found a simple example in `./example` or an example using getText in template file with Gin [https://github.com/pmperrin/go-basic-i18n-example](https://github.com/pmperrin/go-basic-i18n-example)


## Improvements

Here are some improvements that can be made to the go-basic-i18n package:

1. Support for language-country codes: The package should be able to handle language-country codes such as en-GB and en-US. This will allow for more specific translations based on the user's location.
2. Fallback to a generic language code: If a specific language-country code is not found, the package should be able to fall back to a generic language code. For example, if en-CA is not found, the package should use en-* as the fallback.
3. File reloading: The package should be able to reload the translation files in memory when they are modified. This will make it easier to update translations without having to restart the application.
7. Format handling: The package should be able to handle formatting of dates, numbers, and currencies. This is important for providing a localized user experience.
8. Non-Latin script support: The package should be able to handle translations for languages that use non-Latin scripts, such as Chinese, Japanese, and Arabic.
9. and more


## Contributing

Contributions to improve the library, add features, or fix bugs are welcome. Please submit pull requests to the main repository.


## License

`go-basic-i18n` is released under the MIT License. See the [LICENSE](LICENSE) file for details.

This .md file has been done by Phind 
