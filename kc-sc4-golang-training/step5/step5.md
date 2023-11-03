# Environment variables

For this tutorial, let's start in a brand new directory and initialize a new module:

```bash
mkdir -p $HOME/go/src/environment
cd $HOME/go/src/environment
go mod init training/basics/environment
```{{exec}}

If you have done the Hello World tutorial correctly, you should have the following code:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo !")
}
```{{}}

Create a new file `environment.go`{{}} with this content.  
`cp ../hello-world/hello.go environment.go`{{exec}}  

In this example, let's say you want detect the user language, and display a custom message based on this value:

- if it's an English user, show the English message
- if it's a Spanish user, show the Spanish version of the same message

The user language will be stored in an environment variable called **language**. This variable will follow the [ISO 3166 standard](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes):

```bash
# for (true) English users
language=en-GB

# for (American) English users
language=en-US

# for Spanish users
language=es
```{{}}

> [More values are available](http://www.lingoes.net/en/translator/langcode.htm) if you want.

The package `os`{{}} is part of the [Go Standard Library](https://pkg.go.dev/std) and provide some functions to work with environment variables.

The first thing to do is to import this package. Change the **import** statement accordingly:

```diff
diff --git a/src/environment/environment.go b/src/environment/environment.go
index 1f49991..61e8302 100644
@@ -1,7 +1,26 @@
 package main

-import "fmt"
+import (
+       "fmt"
+       "os"
+)

 func main() {
        fmt.Println("Hola Mundo !")
```{{}}

Now you can use the `os.Getenv`{{}} function to read the environment variable value.

```diff
diff --git a/src/environment/environment.go b/src/environment/environment.go
index 44696b3..344c792 100644
@@ -6,5 +6,7 @@ import (
 )

 func main() {
+       language := os.Getenv("language")
+
        fmt.Println("Hola Mundo !")
 }
```{{}}

Finally, you can decide which message to show based on the value:

```diff
diff --git a/src/environment/environment.go b/src/environment/environment.go
index 344c792..d4bdc18 100644
@@ -8,5 +8,9 @@ import (
 func main() {
        language := os.Getenv("language")

-       fmt.Println("Hola Mundo !")
+       if language == "es" {
+               fmt.Println("Hola Mundo !")
+       } else if language == "en-GB" || language == "en-US" {
+               fmt.Println("Hello World !")
+       }
 }
```{{}}

> tips : you could use the command `go fmt`{{exec}} to automaticaly format your Go source code. You could see [https://go.dev](https://go.dev/blog/gofmt) for more details.

You can try to run this code with various values for the `language` variable and look at the output:

`language=es go run environment.go`{{exec}}
should log "Hola Mundo !"

`language=en-GB go run environment.go`{{exec}}
should log "Hello World !"

`language=en-US go run environment.go`{{exec}}
should log "Hello World !" again

`language=invalid go run environment.go`{{exec}}
should log nothing


You can try to run the code with no environment variable, or with an empty value:

`go run environment.go`{{exec}}
should log nothing

`language= go run environment.go`{{exec}}`
should log nothing


There is another function, `os.LookupEnv`, which returns an additional boolean if the variable is present. Let's use this:

```diff
diff --git a/src/environment/environment.go b/src/environment/environment.go
index d4bdc18..e13a94f 100644
@@ -6,7 +6,15 @@ import (
 )

 func main() {
-       language := os.Getenv("language")
+       language, present := os.LookupEnv("language")
+
+       if !present {
+               // we don't know which language the user speaks...
+               // so we're showing this error message in English by default
+               err := fmt.Errorf("environment variable 'language' is missing")
+               fmt.Println(err.Error())
+               os.Exit(1)
+       }

        if language == "es" {
                fmt.Println("Hola Mundo !")
```{{}}

Try to execute your new code with differents values and see how the program behaves.

You should have something similar to this:

`go run environment.go`{{exec}}
# environment variable 'language' is missing
# exit status 1

`language=es go run environment.go`{{exec}}
# Hola Mundo !

`language= go run environment.go`{{exec}}
# should log nothing


The `os.LookupEnv`{{}} function can also distinguish if the variable is absent (unset) versus if the variable is present but has an empty value. In our case, an empty value is an unsupported value. Let's do some refactoring:

```diff
diff --git a/src/environment/environment.go b/src/environment/environment.go
index 7b89e7e..af018bc 100644
@@ -16,9 +16,23 @@ func main() {
                os.Exit(1)
        }

-       if language == "es" {
-               fmt.Println("Hola Mundo !")
-       } else if language == "en-GB" || language == "en-US" {
-               fmt.Println("Hello World !")
+       var message string
+
+       switch language {
+       case "en-GB": // English (United Kingdom)
+               message = "Hello World !"
+
+       case "en-US": // English (United States)
+               message = "Hello America !"
+
+       case "es": // Spanish
+               message = "Hola Mundo !"
+
+       default:
+               err := fmt.Errorf("language '%v' is not supported", language)
+               fmt.Println(err.Error())
+               os.Exit(1)
        }
+
+       fmt.Println(message)
 }
```{{}}

Let's try this code once again, using a loop for simplicity:

```bash
clear
for lang in es en-GB en-US fr invalid-value ; do
    echo "Trying code with language set to '${lang}'"
    language=${lang} go run environment.go
    echo ""
done

# with no value
go run environment.go

# with empty value
language= go run environment.go
```{{exec}}

This will output the following:

```bash
Trying code with language set to 'es'
Hola Mundo !

Trying code with language set to 'en-GB'
Hello World !

Trying code with language set to 'en-US'
Hello America !

Trying code with language set to 'fr'
language 'fr' is not supported
exit status 1

Trying code with language set to 'invalid-value'
language 'invalid-value' is not supported
exit status 1

environment variable 'language' is missing
exit status 1

language '' is not supported
exit status 1
```{{}}

## Final notes

Congratulations, now you know how to read environment variables and adapt your program accordingly.

In the [os package](https://pkg.go.dev/os), there are other functions related to environment variables, such as `os.Setenv("name", "value")`{{}} to define an environment variable, or the useful `os.ExpandEnv` that we can use ourself:

```diff
diff --git a/src/environment/environment.go b/src/environment/environment.go
index af018bc..8cc696c 100644
--- a/src/environment/environment.go
+++ b/src/environment/environment.go
@@ -34,5 +34,6 @@ func main() {
                os.Exit(1)
        }

+       fmt.Println(os.ExpandEnv("You have choosen '${language}' lang."))
        fmt.Println(message)
 }
```{{}}

Output:

`language=es go run environment.go`{{exec}}

```
# You have choosen 'es' lang.
# Hola Mundo !
```{{}}


OK, that's enougth for environement variables, next step will be about command-line arguments.
