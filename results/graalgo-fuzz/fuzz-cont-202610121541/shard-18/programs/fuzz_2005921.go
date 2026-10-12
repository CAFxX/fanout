package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
func genFn0(p0 bool, p1 int64) uint {
    var v5 bool = (32767 != -4)
    var v6 int64 = (p1 ^ p1)
    var v7 int64 = p1
    fmt.Println(v5, v5, v7)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 653
}
func genFn1(p0 int8) int8 {
    var v8 uint16 = 255
    if (!false) {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{(fi9 | fi9), -948, 32768}
            fmt.Println(fi9, v8)
            _ = fi9
            _ = v10
        }
        var v11 int32 = 42
        v8 += v8
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            switch "a" {
            case "hello":
                v14 := []int{int(v8), (fi13 << ((68) & 63)), 265}
                _ = v14
            case "foo bar":
                var v15 uint16 = v8
                _ = v15
            case "":
                switch "\n\t" {
                case "foo bar":
                    p0 = (p0 >> ((255) & 63))
                case "hello":
                    v16 := GenS1{F0: "", F1: v8}
                    _ = v16
                case "a":
                    if false {
                        var v17 string = "\n\t"
                        var v18 uint64 = 18446744073709551615
                        var v19 int16 = 542
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
            default:
                fmt.Println(fi13, fi13, v8)
            }
            fmt.Println(fi12, v8)
            _ = fi13
        }
        var v20 uint32 = 1
        _ = fi12
        _ = v20
    }
    v8 = v8
    _ = p0
    _ = v8
    return p0
}

func main() {
    v1 := []string{"a", ((("\n\t" + "\n\t") + "a") + "\n\t"), "\n\t", "\n\t"}
    var v2 string = ""
    var v3 int32 = 2147483647
    v4 := []uint32{}
    v21 := GenS1{F0: v2, F1: 42}
    var v22 uint = 340
    var v23 int32 = v3
    var v24 uint = 158
    var v25 int64 = 9223372036854775807
    for fi26 := 0; fi26 < 4; fi26++ {
        v22 = (v22 / ((v24) | 1))
        _ = fi26
    }
    fmt.Println(v3)
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 bool = (false && false)
        v2 = (v2 + v2)
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
