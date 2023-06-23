# ASCII Printer

A command-line tool for printing ASCII art with customizable fonts and colors.

## Usage

go run . [OPTIONS] [STRING...]

The tool accepts the following options:

- `--color=<COLOR>`: Specify the color for the ASCII art. Available colors: black, red, green, yellow, blue, magenta, cyan, white. Default: default color.
- `--font=<FONT>`: Specify the font for the ASCII art. Available fonts: shadow, thinkertoy, standard. Default: standard font.

You can pass one or more strings as arguments to print them as ASCII art.

## Examples

Print a single string using the default settings:

go run . "Hello World"


Output:
 _    _          _   _                __          __                 _       _  
| |  | |        | | | |               \ \        / /                | |     | | 
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | 
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | 
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | 
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| 
                                                                                


Print a string with a specific color:

go run . --color=green "Hello World"

Output:

[32m _    _  [0m[32m       [0m[32m _  [0m[32m _  [0m[32m        [0m[32m      [0m[32m__          __ [0m[32m        [0m[32m       [0m[32m _  [0m[32m     _  [0m[32m      [0m
[32m| |  | | [0m[32m       [0m[32m| | [0m[32m| | [0m[32m        [0m[32m      [0m[32m\ \        / / [0m[32m        [0m[32m       [0m[32m| | [0m[32m    | | [0m[32m      [0m
[32m| |__| | [0m[32m  ___  [0m[32m| | [0m[32m| | [0m[32m  ___   [0m[32m      [0m[32m \ \  /\  / /  [0m[32m  ___   [0m[32m _ __  [0m[32m| | [0m[32m  __| | [0m[32m      [0m
[32m|  __  | [0m[32m / _ \ [0m[32m| | [0m[32m| | [0m[32m / _ \  [0m[32m      [0m[32m  \ \/  \/ /   [0m[32m / _ \  [0m[32m| '__| [0m[32m| | [0m[32m / _` | [0m[32m      [0m
[32m| |  | | [0m[32m|  __/ [0m[32m| | [0m[32m| | [0m[32m| (_) | [0m[32m      [0m[32m   \  /\  /    [0m[32m| (_) | [0m[32m| |    [0m[32m| | [0m[32m| (_| | [0m[32m      [0m
[32m|_|  |_| [0m[32m \___| [0m[32m|_| [0m[32m|_| [0m[32m \___/  [0m[32m      [0m[32m    \/  \/     [0m[32m \___/  [0m[32m|_|    [0m[32m|_| [0m[32m \__,_| [0m[32m      [0m
[32m         [0m[32m       [0m[32m    [0m[32m    [0m[32m        [0m[32m      [0m[32m               [0m[32m        [0m[32m       [0m[32m    [0m[32m        [0m[32m      [0m
[32m         [0m[32m       [0m[32m    [0m[32m    [0m[32m        [0m[32m      [0m[32m               [0m[32m        [0m[32m       [0m[32m    [0m[32m        [0m[32m      [0m
   
## Fonts

The tool supports the following fonts:

- `shadow`: Shadow font style.
- `thinkertoy`: Thinkertoy font style.
- `standard`: Standard font style (default).

