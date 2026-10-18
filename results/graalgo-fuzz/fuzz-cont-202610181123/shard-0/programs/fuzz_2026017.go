package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0() uint8 {
    var v4 uint8 = 0
    fmt.Println(v4)
    _ = v4
    return (v4 * 42)
}
func genFn1(p0 uint8, p1 int16, p2 int32) int64 {
    var v5 uint64 = 386
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    v1 := []uint32{42}
    var v2 int32 = 1
    v3 := []uint32{2}
    var v6 uint16 = 378
    if true {
        var v7 bool = true
        switch 0 {
        case -9223372036854775808:
            v2 += v2
        }
        var v8 bool = v7
        _ = v7
        _ = v8
    } else {
        fmt.Println(v6, v2)
    }
    if len(v1) > 0 {
        v1[0] = 494
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v9 int64 = -1
    v9 = (v9 << ((65535) & 63))
    v10 := GenS1{F0: -128, F1: 2}
    var v11 uint64 = 42
    v9 += (v9 << ((74) & 63))
    var v12 string = "\n\t"
    if len(v3) > 0 {
        v3[0] = 65535
    }
    v6 = (v6 / ((v6) | 1))
    switch -463 {
    case 42:
        v6 -= 729
    case -9:
        if true {
            v13 := GenS1{F0: 1, F1: 694}
            if len(v3) > 0 {
                v3[0] = 42
            }
            _ = v13
        }
    case -535:
        switch v12 {
        case "a":
            if true {
                v12 = ("hello" + "a")
                v12 = "\n\t"
            } else {
                if true {
                    v2 -= 1
                    fmt.Println(v11, v9)
                }
                v11 -= 18446744073709551615
            }
        case "":
            v14 := GenS1{F0: 0, F1: 1}
            _ = v14
        case "\n\t":
            var v15 int32 = 2147483647
            _ = v15
        default:
            var v16 int16 = 32767
            _ = v16
        }
    default:
        _ = len(v1)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
