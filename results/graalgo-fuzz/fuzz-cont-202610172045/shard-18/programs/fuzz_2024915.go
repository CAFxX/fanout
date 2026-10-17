package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int = 127
    var v2 int32 = 2147483647
    v3 := []uint32{2147483648, 2}
    var v4 uint16 = 2
    var v5 int32 = (v2 << ((4294967295) & 63))
    switch v1 {
    case -2147483648:
        var v6 int = (^(v1))
        _ = v6
    case 1:
        v2 = v2
    case 32768:
        switch "hello" {
        case "foo bar":
            v2 += 42
        }
    default:
        var v7 int64 = 0
        _ = v7
    }
    v8 := GenS1{F0: 889}
    var v9 int64 = 421
    var v10 uint = 2
    fmt.Println(v10, v5)
    switch v1 {
    case 522:
        if true {
            var v11 bool = true
            var v12 uint64 = 0
            var v13 uint = (v10 + 2)
            _ = v11
            _ = v12
            _ = v13
        }
    case -730:
        v1 = v1
    case -159:
        var v14 uint8 = 58
        _ = v14
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v15 int64 = -338
    switch v1 {
    case 312:
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS2{F0: -9223372036854775808}
            _ = fi16
            _ = v17
        }
    default:
        var v18 int = 443
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    fmt.Println(v15, v4, v15)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
