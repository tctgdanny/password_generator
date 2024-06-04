# Password Generator
You can tell that it's a password generator by the way that it is.  

## Installation
If you work at The Core and have a Mac, this is already on your computer 😊  

To install manually, download `install.sh`, make it executable, and run it. Or clone the repo and run `go build -o pwgen main.go`.  

## Description
Generates a list of passwords in Word+Number format.  

By default, it will output 5 randomly generated passwords with 5 letter words and 4 numbers. For example, simply entering `pwgen` in terminal will output something like:  
```
Slash+1550
Dryer+9129
Bleak+6903
Viral+0504
Relic+0281
```

## Options
Quick syntax can be found by entering: `pwgen --help` or `pwgen -h`  

You can add the flag `--length` or `-l` followed by a number from 4-10, to choose how many characters are in randomly generated passwords.  

Additionally, add `--amount` or `-a` to choose how many passwords you want generated.  

For example, `pwgen --length 7 --amount 15`, or `pwgen -l 7 -a 15`, will output something like:  
```
Pushing+8143
Unitize+3331
Archive+9091
Lechery+0918
Rustler+6350
Decease+4652
Riddled+0056
Subplot+9829
Erratic+7902
Filling+2853
Coronal+2636
Gymnast+0692
Groupie+7659
Braille+2352
Brewery+1425
```
