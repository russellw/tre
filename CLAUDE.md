The purpose of this project is to write a command line utility called 'tre' that is similar to the Windows 'tree' except produces simpler output
In particular, the output is indented but contains no other characters except for the file and directory names
This makes it more suitable for input to other programs
The 'f' flag should also be provided, albeit with switch char -
The program is to be written in Go
It should also respect .gitignore if present in target directory or any parent thereof
The target directory should be a positional command line argument, default to current directory

