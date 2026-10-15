package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 string = "foo bar"
    var v3 uint64 = 2
    v4 := []uint32{}
    var v5 bool = (!(true && true))
    v2 = v2
    switch "\n\t" {
    case "\n\t":
        var v6 uint64 = 42
        _ = v6
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 81}
        v9 := GenS2{F0: 42, F1: -343, F2: 32767}
        var v10 int64 = -9223372036854775808
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS2{F0: 65535, F1: -562, F2: 119}
    v12 := GenS3{F0: 922}
    fmt.Println(v1, v5, v5)
    if v5 {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 bool = true
            var v15 int = (fi13 - 2147483647)
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint8 = 42
        _ = v16
    }
    fmt.Println(v5, v3, v1)
    var v17 uint32 = 114
    v18 := GenS4{F0: 179}
    var v19 uint32 = 2147483648
    fmt.Println(v5)
    switch 9223372036854775807 {
    case 4294967295:
        var v20 uint = 255
        _ = v20
    case -2147483648:
        var v21 uint8 = 1
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
