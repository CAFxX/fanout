package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 int8, p1 uint32, p2 string) uint {
    var v5 int64 = -166
    if (!false) {
        v6 := GenS1{F0: 775, F1: 2147483648, F2: true}
        fmt.Println(p0, p1, v5)
        _ = v6
    } else {
        p0 = (p0 << ((18446744073709551615) & 63))
        v7 := GenS2{F0: (^(p1)), F1: -1}
        _ = v7
    }
    fmt.Println(v5, v5, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 121
}
func genFn1(p0 bool, p1 int64, p2 string) bool {
    v8 := GenS2{F0: 2147483647, F1: 42}
    for fi9 := 0; fi9 < 6; fi9++ {
        p0 = (!p0)
        v10 := []int{(fi9 * fi9), fi9, (fi9 | fi9), fi9}
        p2 = ("hello" + "\n\t")
        _ = fi9
        _ = v10
    }
    v11 := GenS3{F0: 2}
    v12 := GenS1{F0: 427, F1: 65535, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v8
    return p0
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{2147483647, 2147483647, 2}
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint32 = 1
        switch "foo bar" {
        case "a":
            switch "foo bar" {
            case "":
                if len(v2) > 0 {
                    v2[0] = (^(v14))
                }
            case "a":
                var v15 int = (fi13 << ((0) & 63))
                _ = v15
            case "hello":
                switch "a" {
                case "":
                    v16 := GenS4{F0: 149, F1: -627, F2: 253}
                    _ = v16
                case "\n\t":
                    fmt.Println(fi13, v14)
                case "foo bar":
                    var v17 int16 = 198
                    _ = v17
                }
            }
        }
        _ = fi13
        _ = v14
    }
    v3["hello"] = 32768
    var v18 int64 = -9223372036854775808
    if false {
        v18 -= -809
    } else {
        var v19 uint32 = 1
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint8 = 192
            v19 += (v19 | v19)
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    var v22 int16 = 32767
    v1 -= (v1 >> ((65535) & 63))
    fmt.Println(v1, v1, v4)
    v3["\n\t"] = 270
    v23 := v3["foo bar"]
    var v24 string = ("a" + "")
    switch 9223372036854775807 {
    case 9223372036854775807:
        delete(v3, v24)
    default:
        v4 = v4
    }
    var v25 int64 = (v18 - v18)
    for fi26 := 0; fi26 < 7; fi26++ {
        v1 -= (v1 >> ((0) & 63))
        var v27 float32 = 3.14159
        fmt.Println(v18)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
