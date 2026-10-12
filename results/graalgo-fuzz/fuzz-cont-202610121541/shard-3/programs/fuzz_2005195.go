package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
func genFn0(p0 string, p1 int8) uint16 {
    for fi5 := 0; fi5 < 5; fi5++ {
        p0 = p0
        _ = fi5
    }
    switch "hello" {
    case "\n\t":
        p0 = ("foo bar" + "\n\t")
    default:
        v6 := []int{1, 32768, -778, 128}
        _ = v6
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int64 = (v1 % (((v1 ^ v1)) | 1))
    var v3 float32 = 0.0
    v4 := make(map[uint32]uint32)
    v4[0] = 2147483647
    switch "foo bar" {
    case "foo bar":
        var v7 float64 = 0.1
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v3 = 3.14159
        fmt.Println(v1)
        _ = fi8
    }
    if true {
        v2 += 188
    } else {
        var v9 uint16 = 255
        var v10 uint16 = (v9 & (v9 | 65535))
        v4[720] = 42
        _ = v10
        _ = v9
    }
    var v11 float64 = ((-2.25 * 1.5) + 3.14159)
    if true {
        var v12 uint64 = 18446744073709551615
        _ = v12
    }
    var v13 uint32 = 2
    v2 += v1
    for fi14 := 0; fi14 < 17; fi14++ {
        v2 = (v2 & -9223372036854775808)
        var v15 uint64 = 2
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 424, F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
