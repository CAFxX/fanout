package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 1
    var v3 float64 = 1e10
    var v4 int16 = (v2 ^ (v2 + int16(v1)))
    v5 := []uint32{2147483648, 1}
    var v6 uint8 = 61
    var v7 uint8 = 0
    if false {
        var v8 int32 = 0
        switch int(v1) {
        case 42:
            fmt.Println(v3, v2)
        case 4294967295:
            fmt.Println(v6, v1, v8)
        case 32767:
            fmt.Println(v4)
        }
        _ = v8
    } else {
        var v9 int8 = 1
        _ = v9
    }
    var v10 bool = false
    switch "a" {
    case "":
        var v11 bool = ((false || (!((0.0 / (0.0 - ((3.14159 - -2.25) + 1.5))) == ((1e10 * 1e10) - 0.0)))) && v10)
        _ = v11
    }
    switch 42 {
    case 0:
        if len(v5) > 0 {
            v5[0] = 65535
        }
    }
    var v12 float32 = 1.5
    switch int(v7) {
    case -618:
        v13 := GenS1{F0: 2, F1: 1, F2: -1}
        _ = v13
    default:
        v14 := GenS1{F0: 42, F1: 2147483647, F2: 10}
        _ = v14
    }
    v15 := GenS2{F0: 102}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
