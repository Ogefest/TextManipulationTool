# Text Manipulation Tool

CLI tool created for manipulate texts files as replacement for `grep` `sed` `tr` `cut` in one tool. Simple commands in `imagemagick` style allow to create chains of commands to change text files or extract data. Manipulate texts without opening manuals or hard to remember switches. Filter lines, replace strings or change column order.

## Usage

`tmt` read from STDIN, procesing commands and write to STDOUT as result. Example usage

`
cat /tmp/some.file.log | tmt filter somestring replace str1 str2 cut 10:20 > /tmp/output.log
`

You can create as many commands as you need creating chains of commands.

## Commands


| command | params | description |
|---------|--------|-------------|
| **replace** | STRING_FROM STRING_TO | replace one string into another |
| **lower** | - | convert string to lower case |
| **upper** | - | convert string to upper case |
| **trim**  | - | remove white spaces |
| **duplicate** | - | create duplicated string in new line |
| **filter** | STRING | remove if string not contain STRING |
| **if** | STRING | proceed only if string contains STRING|
| **ifnot** | STRING | proceed only if string not contains STRING |
| **prefix** | STRING | add STRING before | 
| **suffix** | STRING | add STRING after |
| **cut** | START_INT:END_INT | extract string from string |
| **columnseparator** | STRING | set string as column separator, could be used multiple times. Default " " (space)|
| **column** | COLUMN_NUMBER | extract column from string |
| **columnremove** | COLUMN_NUMBER | remove column |
| **columnadd** | COLUMN_NUMBER STRING | add STRING after COLUMN_NUMBER in string |
| **columnselect** | COLUMN_NUMBER | select column for processing other commands only with this column |
| **columndeselect** | - | disable column selection |
| **columnorder** | COL1,COL2,COL3... | change column order in string |


## Examples

All commands are described using example log file. 
```
431338047 app 10.36.0.14:57592     
431340107 root 17.224.118.13:36114 
431340618 app 21.138.232.192:58918 
431340618 usr 111.238.212.12:52918 
```

### Replace root user to app

`cat input.log | tmt replace root app`

```
431338047 app 10.36.0.14:57592     
431340107 app 17.224.118.13:36114 
431340618 app 21.138.232.192:58918 
431340618 usr 111.238.212.12:52918
```

### Add string when user is not app

`cat input.log | tmt ifnot app suffix MARKED`

```
431338047 app 10.36.0.14:57592     
431340107 root 17.224.118.13:36114 MARKED
431340618 app 21.138.232.192:58918 
431340618 usr 111.238.212.12:52918 MARKED
```

### Get first digt from port number

`cat input.log | tmt columnseparator : column 2 cut :1`

```
5
3
5
5
```

### Replace port with IP

`cat input.log | tmt replace ":" " " columselect 4 trim suffix : columnorder 1,2,4,3 columndeselect replace ": " ":"`

```
431338047 app 57592:10.36.0.14
431340107 root 36114:17.224.118.13
431340618 app 58918:21.138.232.192
431340618 usr 52918:111.238.212.12
```

## Install

Download [from GitHub release page](https://github.com/Ogefest/TextManipulationTool/releases) and put `tmt` file into `/usr/bin/tmt`

`sudo cp tmt /usr/bin/tmt`


If you want to build from source
```
git clone git@github.com:Ogefest/TextManipulationTool.git tmt
cd tmt
make
sudo make install
```

## Contribution

If you have idea for new commands, new features or found bugs let me know using GitHub issues link