package main

import (
	"fmt"
)

func main() {
	mds := make([][]string, 0, 20)
	ss := make([]string, 0, 20)
	nn := 10

	for i := 1; i <= 21; i += 2 {
		nn--
		ss := append(ss[:0])
		for z := 0; z <= nn; z++ {
			ss = append(ss, " ")
		}
		for j := 0; j < i; j++ {
			ss = append(ss, "X")
		}
		for y := 0; y <= nn; y++ {
			ss = append(ss, " ")
		}
		mds = append(mds, ss)

	}

	for _, v := range mds {
		fmt.Println(v)
	}
}

/*
#PROBLEM 1
if in line 14 instead of
ss := append(ss[:0])
I use
ss = append(ss[:0])
OUTPUT:
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]


#PROBLEM 2
If I comment from line 21 to 23
for y := 0; y <= nn; y++ {
		 	ss = append(ss, " ")
		 }
I'm not sure why it does that because:
for "z" starts the slice and I append new values in for "j"
so since append does add values at the end of the slice the output should be:
[                    X] not [X X X X X X X X X X X]

OUTPUT
[X X X X X X X X X X X]
[X X X X X X X X X X X X]
[X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X]
[X X X X X X X X X X X X X X X X X X X X X]


#SOLUTION

if I run the code in the way it has been written the output is what I was expecting to get:

OUTPUT
[                    X                    ]
[                  X X X                  ]
[                X X X X X                ]
[              X X X X X X X              ]
[            X X X X X X X X X            ]
[          X X X X X X X X X X X          ]
[        X X X X X X X X X X X X X        ]
[      X X X X X X X X X X X X X X X      ]
[    X X X X X X X X X X X X X X X X X    ]
[  X X X X X X X X X X X X X X X X X X X  ]
[X X X X X X X X X X X X X X X X X X X X X]
*/
