# Install Guide (buliding from source, I'm too lazy to just pre-build it :P)
(requirements):
Golang

Just run these commands:

git clone https://github.com/LiterallyKirby/Shell

cd Shell

go build

(bonus) move the shell exeucuteable to /usr/bin then name it what you want for ease of use (linux only :P)

Then just setup ur terminal of choice to run this the shell on start up

if you moved it to /usr/bin just run whatever you named it

# Shell
A bad shell written in golang
It has the features from your normal shell just with some custom commands I'll add on to over time

# Encrypt
Syntax: encrypt filename file-output-name

Just encrypts a file with a given password. Nothing special I just didn't feel like installing encryption software :P

# Decrypt
Syntax: decrypt filename file-output-name

Doe's the reverse of above

# Find
Syntax: find what you want to find lol

Added simply because I don't like using grep

# Cd

Syntax: cd directory-name-here

Added because whenever I do use windows I tend to use cd by mistake

# History

Syntax:
history (to list past commands)

!x (to run that command, x represents the command number)

Added because I simply dont like the way history works in other shells.
