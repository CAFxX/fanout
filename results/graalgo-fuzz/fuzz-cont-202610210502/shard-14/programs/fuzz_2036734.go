package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 string) string {
    switch (p0 + p0) {
    case "":
        fmt.Println(p0)
    case "foo bar":
        var v4 int8 = -104
        _ = v4
    case "hello":
        var v5 int8 = 127
        _ = v5
    }
    var v6 int = -9223372036854775808
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 uint = 2
    var v2 uint32 = 42
    v3 := make(map[uint32]uint32)
    v1 += v1
    v2 += (v2 - (v2 & 4294967295))
    var v7 uint8 = uint8(v1)
    var v8 int8 = -47
    var v9 uint8 = 128
    if (!false) {
        fmt.Println(v7, v9, v2)
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{fi10}
            v2 = (v2 & (v2 >> (((v1 + v1)) & 63)))
            _ = fi10
            _ = v11
        }
        var v12 float64 = (1.5 * 1e10)
        _ = v12
    }
    v13 := GenS1{F0: 557, F1: 639}
    v3[(v2 ^ v2)] = 65535
    if (2147483648 > v2) {
        var v14 bool = true
        _ = v14
    }
    v15 := GenS1{F0: 2, F1: 42}
    var v16 uint = (^(v1))
    v16 = (v16 * uint(v7))
    fmt.Println(v9, v16)
    for fi17 := 0; fi17 < 1; fi17++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v9, v1, fi17)
        default:
            var v18 uint = 602
            _ = v18
        }
        v7 += 19
        fmt.Println(v8, v8)
        _ = fi17
    }
    v8 -= 100
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
