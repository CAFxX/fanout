package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0() uint {
    var v4 string = ("a" + "\n\t")
    v5 := []int{}
    v4 = v4
    switch -1 {
    case 284:
        switch "a" {
        case "":
            if (false && ((true || (!false)) || (true && (!(!(!(true || true))))))) {
                switch "foo bar" {
                case "foo bar":
                    v6 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: 0}
                    _ = v6
                case "hello":
                    v4 = v4
                default:
                    var v7 uint16 = 1
                    _ = v7
                }
                if true {
                    var v8 uint16 = 32767
                    var v9 int = 4294967295
                    _ = v8
                    _ = v9
                }
                v5 = append(v5, 32767)
            }
        default:
            v10 := GenS2{F0: 67, F1: -94}
            _ = v10
        }
    case 0:
        v4 = v4
    case 32768:
        var v11 int64 = -153
        _ = v11
    default:
        v4 = v4
    }
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 int32, p1 uint8) uint {
    var v12 uint32 = 311
    v13 := []int{-116, -201, -2147483648}
    fmt.Println(p1, p0)
    var v14 bool = false
    v15 := GenS3{F0: 65535, F1: 0, F2: v12}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 255
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 1
    var v3 float32 = -2.25
    v16 := []int{32768, 751}
    v2 += (v2 << ((0) & 63))
    var v17 uint64 = 1
    switch "foo bar" {
    case "hello":
        var v18 uint = 913
        _ = v18
    case "foo bar":
        if true {
            v19 := GenS2{F0: 41, F1: 127}
            _ = v19
        } else {
            fmt.Println(v1, v1)
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    case "\n\t":
        var v20 uint8 = v2
        _ = v20
    default:
        v1 = v1
    }
    if true {
        switch -723 {
        case -9223372036854775808:
            v2 = (^(v2))
        default:
            fmt.Println(v17, v2, v2)
        }
        v2 = v2
    }
    _ = len(v16)
    fmt.Println(v3)
    v17 += (v17 | (v17 | v17))
    var v21 int = 2147483647
    var v22 uint8 = 24
    if (false || ((false || true) || (!((^(v1)) > v1)))) {
        var v23 int16 = 1
        _ = v23
    }
    var v24 int32 = -96
    for fi25 := 0; fi25 < 3; fi25++ {
        v2 = 42
        _ = fi25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v21 = (v21 ^ v21)
        v21 -= (v21 + v21)
        var v27 int64 = 1
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
