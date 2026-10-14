package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }

func main() {
    v1 := []string{(("foo bar" + "\n\t") + "foo bar"), "\n\t"}
    var v2 int8 = 40
    var v3 uint = 350
    var v4 int16 = -626
    var v5 int64 = 0
    var v6 uint8 = 208
    v7 := GenS1{F0: 0, F1: v6, F2: 832}
    v4 = v4
    switch "a" {
    case "foo bar":
        v6 -= (^(v6))
    default:
        fmt.Println(v4)
    }
    var v8 uint = (v3 * (v3 | (v3 ^ (^(v3)))))
    v2 = -44
    v5 += (v5 | v5)
    var v9 int64 = (v5 | int64(v3))
    var v10 bool = (!false)
    var v11 string = "a"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((false || ((v4 <= (^(v4))) && (false && (true || v10)))) || true) {
        v5 = (v9 % ((9223372036854775807) | 1))
        v12 := GenS2{F0: (^(v2))}
        for fi13 := 0; fi13 < 13; fi13++ {
            v4 = -32768
            _ = fi13
        }
        _ = v12
    } else {
        v3 += (v8 >> ((uint(v6)) & 63))
        var v14 int32 = 2147483647
        v15 := GenS3{F0: uint16(v8), F1: 2147483647, F2: 2}
        _ = v14
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
