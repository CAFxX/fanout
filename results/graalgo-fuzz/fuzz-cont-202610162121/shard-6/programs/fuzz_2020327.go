package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 int16) int16 {
    var v6 string = (("a" + "") + "\n\t")
    if (3.14159 < (1.5 + 1.5)) {
        fmt.Println(p1, p1, p1)
        fmt.Println(p1)
    }
    v7 := []int{1, int(p2)}
    switch "" {
    case "":
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 uint16 = 255
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case "foo bar":
        if (27 > 147) {
            switch 128 {
            case 937:
                var v11 float32 = 0.1
                _ = v11
            }
            p1 = (p1 << ((573) & 63))
            p2 = (p2 + p2)
        }
    case "hello":
        v12 := GenS1{F0: 0, F1: p1}
        _ = v12
    default:
        for fi13 := 0; fi13 < 13; fi13++ {
            p0 = (p0 / (((p0 * (p0 << ((1) & 63)))) | 1))
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        if true {
            v15 := GenS2{F0: ("a" + ((v6 + "hello") + "\n\t")), F1: 2147483647, F2: 100}
            fmt.Println(p2, p0, p2)
            _ = v15
        }
        v6 = ""
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (^(p2))
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{845}
    var v3 uint64 = 99
    v4 := []string{"hello", "\n\t"}
    v5 := []int{32767}
    if (!(!false)) {
        v16 := GenS3{F0: -525, F1: (!false)}
        _ = v16
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v3)
    v3 += (^(v3))
    var v17 int32 = 1
    v18 := GenS1{F0: 140, F1: 597}
    switch ("foo bar" + "\n\t") {
    case "":
        v17 -= (v17 << ((2) & 63))
    case "a":
        if (18446744073709551615 > 1) {
            var v19 int8 = 42
            fmt.Println(v17, v17, v3)
            _ = v19
        } else {
            var v20 float64 = 3.14159
            v2 = append(v2, 4294967295)
            _ = v20
        }
    case "\n\t":
        v21 := GenS4{F0: 9223372036854775807, F1: 939}
        _ = v21
    }
    var v22 float32 = (0.0 - 1.5)
    if true {
        if (false && false) {
            v23 := GenS5{F0: 521}
            _ = v23
        }
        var v24 string = "\n\t"
        _ = v24
    }
    fmt.Println(v3)
    fmt.Println(v3, v17)
    v25 := GenS5{F0: 42}
    var v26 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
