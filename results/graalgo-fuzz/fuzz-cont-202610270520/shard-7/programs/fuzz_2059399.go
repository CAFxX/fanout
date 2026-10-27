package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
func genFn0() string {
    if (!(0 != 65535)) {
        var v5 float32 = float32(0.1)
        v5 -= -2.25
        _ = v5
    }
    var v6 uint16 = 1
    var v7 int = 341
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 uint64, p1 uint16, p2 uint8) int16 {
    switch 9223372036854775807 {
    case -9223372036854775808:
        p1 = (p1 << ((42) & 63))
    case 2147483647:
        p0 = p0
    case 1:
        var v8 uint8 = 14
        _ = v8
    default:
        v9 := []int{int(p2), 1}
        _ = v9
    }
    switch "hello" {
    case "":
        var v10 int64 = -393
        _ = v10
    case "\n\t":
        fmt.Println(p1, p1)
    case "hello":
        var v11 uint16 = (^(p1))
        _ = v11
    }
    p2 -= 169
    for fi12 := 0; fi12 < 16; fi12++ {
        switch "a" {
        case "":
            var v13 int8 = 76
            _ = v13
        case "foo bar":
            var v14 string = "\n\t"
            _ = v14
        case "\n\t":
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 uint32 = 765
                _ = fi15
                _ = v16
            }
        default:
            fmt.Println(fi12, fi12, fi12)
        }
        _ = fi12
    }
    var v17 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 0
}

func main() {
    var v1 uint32 = 288
    v2 := make(map[int]int)
    v3 := []uint32{(v1 * v1), 1, 0, 823}
    var v4 uint32 = 4294967295
    if len(v3) > 0 {
        v3[0] = v1
    }
    var v18 float64 = (1e10 * 1e10)
    v4 += v4
    var v19 int64 = 42
    switch 2147483647 {
    case -9223372036854775808:
        v4 = v1
    default:
        var v20 uint64 = uint64(v4)
        _ = v20
    }
    delete(v2, -9223372036854775808)
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int = (fi21 * (fi21 + (fi21 - 455)))
        _ = fi21
        _ = v22
    }
    if (63 != 108) {
        var v23 int = 4294967295
        var v24 uint32 = (^(v4))
        _ = v23
        _ = v24
    }
    var v25 float32 = 3.14159
    v18 -= ((float64(v25) + float64(v25)) * v18)
    v25 += ((v25 - v25) * v25)
    var v26 string = ""
    if (!("hello" >= (((v26 + v26) + ("hello" + ("\n\t" + v26))) + v26))) {
        fmt.Println(v19)
        v27 := GenS1{F0: 0, F1: 203, F2: (42 < 42)}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
