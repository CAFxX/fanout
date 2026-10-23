package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }

func main() {
    var v1 bool = (((-2.25 <= (1.5 - 0.1)) || true) || ("hello" == "hello"))
    v2 := []uint32{2147483648, 475, 2147483648, 65535}
    var v3 float64 = (-2.25 + 1.5)
    var v4 float64 = 1e10
    var v5 bool = false
    v6 := make(map[string]string)
    fmt.Println(v4)
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v7 int8 = 127
    v1 = v1
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v5)
        _ = fi8
    }
    var v9 uint64 = 4294967295
    var v10 int16 = -614
    if len(v2) > 0 {
        v2[0] = 65535
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch "\n\t" {
        case "":
            v10 = (v10 << ((18446744073709551615) & 63))
        case "a":
            var v12 uint16 = 65535
            _ = v12
        default:
            v13 := GenS1{F0: (!false), F1: 2, F2: 450}
            _ = v13
        }
        v6["hello"] = "hello"
        _ = fi11
    }
    switch "hello" {
    case "a":
        v5 = false
    case "":
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(v10)
            var v15 int32 = 42
            fmt.Println(v1)
            _ = fi14
            _ = v15
        }
    case "\n\t":
        var v16 bool = (846 != 0)
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = 428
    }
    v1 = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
