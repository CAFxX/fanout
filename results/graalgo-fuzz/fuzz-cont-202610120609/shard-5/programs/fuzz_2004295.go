package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int16; F2 int }
func genFn0(p0 string) int16 {
    p0 = ("a" + "\n\t")
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 string = (("hello" + "hello") + "\n\t")
    var v2 uint = 645
    var v3 int = -1
    v4 := GenS1{F0: 146}
    if (v3 >= v3) {
        var v5 float64 = (0.0 / 1.0)
        _ = v5
    }
    var v6 string = "hello"
    var v7 uint8 = 127
    v6 = v6
    if ((-2.25 + ((3.14159 + 0.0) * 1e10)) > (((0.0 + 0.1) + (((0.1 - 1.5) * float64((-2.25 + 3.14159))) * (3.14159 + 0.1))) / (3.14159 + ((1.5 + 1.5) + -2.25)))) {
        v8 := GenS1{F0: v2}
        _ = v8
    }
    v9 := []int{(v3 - (v3 << (((v2 / (((v2 ^ (v2 % ((65535) | 1)))) | 1))) & 63))), 2147483647}
    var v10 int = int(v2)
    if (false || (401 <= -86)) {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint16 = 2
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: v2}
                if (754 <= -1) {
                    v2 = 1
                    var v15 int8 = 110
                    v16 := GenS2{F0: (false && false), F1: 1000, F2: (v10 & -326)}
                    v3 = (v3 ^ int(v12))
                    v17 := GenS2{F0: ("\n\t" <= "a"), F1: -572, F2: (v3 | (-(v3)))}
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS1{F0: (v2 << ((65535) & 63))}
                    fmt.Println(fi13)
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v7)
    var v19 string = "hello"
    v1 = (v6 + ((v19 + (v1 + v6)) + v19))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
