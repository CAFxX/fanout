package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 int64, p1 uint64) string {
    for fi4 := 0; fi4 < 10; fi4++ {
        p0 = p0
        p1 = 484
        _ = fi4
    }
    var v5 int64 = 0
    _ = p0
    _ = p1
    _ = v5
    return ((("" + ("foo bar" + ("" + "hello"))) + "\n\t") + "foo bar")
}
func genFn1(p0 int8, p1 uint8) uint16 {
    if (255 <= 4294967295) {
        v6 := []int{32767}
        fmt.Println(p1, p0)
        v7 := GenS1{F0: 810, F1: -103}
        _ = v6
        _ = v7
    }
    p1 += 0
    var v8 uint16 = 42
    _ = p0
    _ = p1
    _ = v8
    return v8
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"\n\t", ("foo bar" + "\n\t"), (("" + "") + "foo bar"), ("" + "hello")}
    v3 := []uint32{}
    fmt.Println(v1)
    var v9 uint32 = 2
    var v10 uint8 = 255
    var v11 uint8 = (v10 | 190)
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint = 662
        _ = fi12
        _ = v13
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v11, v1)
    default:
        if (59 < 0) {
            fmt.Println(v11, v9)
            v9 += (v9 + (v9 | v9))
        }
    }
    v11 -= (v11 | 2)
    if false {
        v10 -= v10
    } else {
        fmt.Println(v10, v10)
        if false {
            v10 = (v11 >> ((255) & 63))
            var v14 int = -2147483648
            _ = v14
        } else {
            v15 := GenS2{F0: 1, F1: 0}
            var v16 bool = (32768 <= 4294967295)
            _ = v15
            _ = v16
        }
        fmt.Println(v10, v1)
    }
    if len(v3) > 0 {
        v3[0] = (v9 | 4294967295)
    }
    v9 += v9
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v17 := GenS1{F0: (v9 * (^(v9))), F1: -40}
    v1 += v1
    v9 = 2147483648
    v9 = 943
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int16 = (^(v1))
        var v20 int16 = (v1 - v19)
        var v21 int = (fi18 * fi18)
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
