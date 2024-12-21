package main

import . "fmt"

func main() {
	k := secure_rand()
	p := rand_plain(32)
	c := enc(k,p)
	d := dec(k,c)
	Printf("k = \n%064b \n",k)
	Printf("p = \n")
	for _ , b := range p {
		Printf("%1b",b)
	}
	Printf("\nc = \n")
	for _ , w := range c {
		Printf("%064b \n", w)
	}
	Printf("d = \n")
	for _ , b := range d {
		Printf("%1b",b)
	}
	

}