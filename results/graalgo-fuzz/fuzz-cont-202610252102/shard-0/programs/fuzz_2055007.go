package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 uint) uint16 {
    if false {
        p0 += (p0 ^ 837)
    }
    if (42 >= 42) {
        fmt.Println(p0)
        p0 += (^(p0))
        p0 += p0
    }
    _ = p0
    return 996
}

func main() {
    var v1 float32 = ((-2.25 * 0.1) * 1.5)
    var v2 uint16 = 2
    v3 := []string{"hello", "a"}
    v4 := []uint32{}
    var v5 bool = true
    var v6 uint8 = 42
    v1 = v1
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 6; fi8++ {
        v2 = (v2 | v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v7 = (fi8 >> ((2) & 63))
        _ = fi8
    }
    v9 := GenS1{F0: 2147483647, F1: 0, F2: -32768}
    v6 += v6
    switch "hello" {
    case "hello":
        v2 = 271
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v1, v7)
            _ = fi10
        }
    case "":
        _ = len(v4)
    }
    fmt.Println(v7, v1)
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float64 = ((3.14159 + 3.14159) / -2.25)
        v13 := GenS2{F0: true, F1: 1}
        _ = fi11
        _ = v12
        _ = v13
    }
    v7 -= 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
