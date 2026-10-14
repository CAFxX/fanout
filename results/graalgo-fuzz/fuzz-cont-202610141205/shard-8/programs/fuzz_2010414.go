package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    var v3 int64 = 0
    var v4 uint16 = 42
    v5 := []int{}
    if (1 < 223) {
        v5 = append(v5, -722)
    }
    if (!(4294967295 > 0)) {
        v4 += (^(v4))
    }
    if true {
        var v6 string = v1
        _ = v6
    }
    v7 := GenS1{F0: 150, F1: 657}
    v8 := v2[v1]
    var v9 uint32 = 42
    v9 -= v9
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            var v11 float32 = 0.1
            var v12 float64 = 1.5
            _ = v11
            _ = v12
        } else {
            var v13 float64 = (0.0 + -2.25)
            _ = len(v5)
            v1 = v1
            _ = v13
        }
        _ = fi10
    }
    switch v1 {
    case "hello":
        if true {
            var v14 uint16 = v4
            _ = v14
        }
    case "\n\t":
        if (((2147483647 <= -404) || false) || (false || false)) {
            v5 = append(v5, (v8 % (((v8 + (v8 >> ((344) & 63)))) | 1)))
        }
    case "foo bar":
        v15 := GenS2{F0: 512}
        _ = v15
    default:
        if len(v5) > 0 {
            v5[0] = (-(v8))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
