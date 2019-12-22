# TextManipulationTool

CLI tool created for manipulate texts files as replacement for `grep` `sed` `tr` `cut` in one tool. Simple commands in `imagemagick` style allow to create chains of commands to change text files or extract data. Manipulate texts without opening manuals or hard to remember switches, in tmt commands are much cleaner.

## Usage

tmt read from STDIN, procesing commands and write to STDOUT as result. Example usage

`
cat /tmp/some.file.log | tmt filter somestring replace str1 str2 cut 10:20 > /tmp/output.log
`

You can create as many commands as you need creating chains of commands.

## Commands
- replace
- lower
- upper
- title
- trim
- duplicate
- filter
- prefix
- suffix
- cut

