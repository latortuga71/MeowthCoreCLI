package main

import (
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/latortuga71/MeowthCoreCLI/cmd"
)

func main() {
	fmt.Printf(`
			    .',,,'.                          
			   ;ONWWWNO;                         
			  :O0KXXXXX0;                        
			.ONkdxdddONO.                             
		.:c'..;cOWM0l;;:lOWWOl;..,lc.               
		.OOloolkNMNX0dcxXKXW0kdoooko.               
		.Oo .;oxkxk0xoodk0kolxd,  c:                
		.Oo   .lOOkOkl;;lOO00c.  .dd.               
		lk;.ckKNWWWkllkNMWWN0o,.:0:                
		;KXXWMMMW0oxXMMMMMMMWMN0KO'                
		;KMXd:;;;..dWMMMMWNNWMMMMK,                
		;KMx.      lWMMMMKc:0MMMMK;                
		;XM0;    .,OWWMMMW0ONMMMMK,                
		.xWWO;:dx0NNXXXWWMMMMMMMNd.                
		'xxcxNMMMWKkkkXMMMMMMMNd.                 
		. .. .'xNWWMMWK0XKXWMMMMWKl. .  .             
		... ..;lkKWWNWMNNWWKkl;.   ..               
		.:c:,..  .lO0K000x:. ...,c:.                
		.o0XK0d,.. ..........oK0XNk'                
		.;oc,odoc'''...':ldcoko;.                 
		..;cl;;dc,,'...                      
		..;c,.,;';llo:..                      
		'cll:::,,. .   ..,,;c:,.                   
		.xKd'                .,lxxl,.              
		..                     ,dc.               
							.   
`)
	fmt.Println()
	grumble.Main(cmd.App)

}
