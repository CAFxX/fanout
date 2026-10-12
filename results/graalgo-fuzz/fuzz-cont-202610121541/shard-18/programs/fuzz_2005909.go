package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint8, p1 int8) uint {
    p0 += (p0 % ((p0) | 1))
    var v4 int8 = (p1 & (p1 & 0))
    if (true || (1 > uint(p0))) {
        fmt.Println(p0, p1)
    }
    var v5 int32 = 275
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 471
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    var v3 uint64 = 421
    v2 = (v2 ^ (v2 | v2))
    _ = len(v1)
    var v6 uint16 = 201
    switch ("foo bar" + "hello") {
    case "hello":
        var v7 int8 = 17
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(v2, v2)
            _ = fi8
        }
    }
    var v9 uint32 = 2147483648
    v2 -= genFn0(127, -91)
    fmt.Println(v9)
    v2 = 65535
    v10 := GenS1{F0: 2147483647, F1: (^(v2))}
    v2 = v2
    v3 = v3
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 -= 18446744073709551615
        v9 = 0
        v12 := GenS2{F0: 9223372036854775807, F1: 127}
        _ = fi11
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
