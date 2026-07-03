# escpaper
a small escaping library in golang



i needed an escaping function. and nobody i found could do it the way i want;\
so i made this;\
usage:
```
package main;
import "fmt"
func main () {
  // returns \x1b"hello\x1b"
  escapedstring ,err := escpaper.Escape(`\"hello\"`)
  //returns (theoretically) i said \x1b"hello there\x1b"
  substring,err  := escpaper.Delim(`i said \"hello there\"",other_stuff`);

}
```
