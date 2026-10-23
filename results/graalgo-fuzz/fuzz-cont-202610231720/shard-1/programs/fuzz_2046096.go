package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint32 = 42
    var v2 int32 = 2147483647
    var v3 int16 = -1
    v4 := []string{("foo bar" + "a"), "a"}
    fmt.Println(v3)
    var v5 uint8 = 149
    v6 := GenS1{F0: v1, F1: false, F2: v1}
    fmt.Println(v1, v3)
    v1 += (^(v1))
    switch "a" {
    case "hello":
        v7 := GenS1{F0: v1, F1: (-1 >= 127), F2: (v1 >> ((1) & 63))}
        _ = v7
    case "\n\t":
        v5 -= (^(v5))
    case "":
        v5 = 42
    default:
        v3 += v3
    }
    v8 := GenS1{F0: (v1 >> ((495) & 63)), F1: true, F2: (v1 & v1)}
    v1 = v1
    var v9 bool = true
    var v10 uint64 = 42
    v11 := GenS2{F0: 9223372036854775807}
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = "a"
        var v14 bool = (v13 < ("hello" + v13))
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
