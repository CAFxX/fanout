package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint8 = 1
    var v2 int16 = 32767
    var v3 bool = false
    var v4 float64 = ((float64(1e10) - 3.14159) + (3.14159 + float64((float32(((-2.25 - (0.1 - float64(3.14159))) + 0.0)) - 1.5))))
    v5 := []string{}
    var v6 uint = 4294967295
    v6 = (v6 ^ 18446744073709551615)
    var v7 bool = ((v3 || (!v3)) || (!(1 < 483)))
    v3 = v7
    var v8 int64 = 0
    if false {
        v9 := GenS1{F0: v1, F1: v6}
        _ = v9
    }
    if v3 {
        if len(v5) > 0 {
            v5[0] = ("foo bar" + "")
        }
    } else {
        v2 = (v2 * (v2 | (^(v2))))
        v5 = append(v5, ("\n\t" + "foo bar"))
    }
    var v10 float64 = v4
    fmt.Println(v6, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
