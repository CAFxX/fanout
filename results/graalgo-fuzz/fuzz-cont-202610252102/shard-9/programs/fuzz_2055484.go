package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16) uint32 {
    var v7 int32 = 1
    var v8 bool = (42 < 518)
    var v9 int = 127
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 574
}

func main() {
    var v1 float64 = float64(3.14159)
    var v2 float64 = (1.5 * v1)
    var v3 uint = 2
    var v4 int8 = 42
    var v5 uint16 = 1
    v6 := []string{}
    fmt.Println(v3, v3, v2)
    var v10 int64 = 42
    fmt.Println(v4)
    switch "\n\t" {
    case "":
        v1 -= v2
    case "a":
        v5 = 2
    default:
        v4 -= v4
    }
    v11 := GenS1{F0: v4, F1: 18446744073709551615, F2: 433}
    for fi12 := 0; fi12 < 2; fi12++ {
        v2 = float64((1.5 - 1e10))
        _ = fi12
    }
    v13 := GenS2{F0: 2}
    v10 += v10
    v1 += ((float64((((-2.25 * -2.25) + (-2.25 * 0.1)) / 3.14159)) / v1) + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
