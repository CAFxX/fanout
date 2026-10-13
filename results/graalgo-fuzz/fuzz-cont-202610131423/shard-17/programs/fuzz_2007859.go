package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 uint = 65535
    v3 = (v3 * (v3 & v3))
    v3 -= v3
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint32 = 42
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: 2147483647}
    var v7 int8 = 100
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 bool = (42 <= 1)
        for fi10 := 0; fi10 < 6; fi10++ {
            v7 = -128
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v11 uint = (v3 ^ v3)
    v7 -= v7
    fmt.Println(v7, v3, v7)
    var v12 int32 = 0
    var v13 int16 = 32767
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v11, v11, fi14)
        _ = fi14
    }
    v15 := GenS1{F0: 18446744073709551615, F1: 0}
    _ = len(v1)
    v16 := v2[""]
    v2["foo bar"] = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
