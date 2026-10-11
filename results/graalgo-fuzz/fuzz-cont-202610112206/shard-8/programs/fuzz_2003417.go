package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint16, p2 int16) uint {
    p2 += (p2 + p2)
    var v5 string = ("\n\t" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 int, p1 int16) uint64 {
    p0 -= (p0 >> ((65) & 63))
    switch "" {
    case "":
        v6 := GenS1{F0: (p1 + (p1 % ((p1) | 1))), F1: (p0 * p0), F2: 0}
        _ = v6
    case "\n\t":
        fmt.Println(p0, p0)
    case "hello":
        v7 := GenS2{F0: false, F1: 1}
        _ = v7
    }
    switch p0 {
    case 283:
        v8 := GenS1{F0: 1, F1: (p0 - p0), F2: 127}
        _ = v8
    case -186:
        if (-1 > 331) {
            var v9 uint64 = 42
            fmt.Println(p1, p1)
            _ = v9
        } else {
            p1 += (p1 + p1)
            v10 := GenS1{F0: (p1 / ((1000) | 1)), F1: -494, F2: -128}
            _ = v10
        }
    default:
        switch ("a" + "\n\t") {
        case "foo bar":
            var v11 uint16 = 255
            _ = v11
        case "":
            var v12 int8 = 1
            _ = v12
        case "hello":
            if (56 < 127) {
                if (int32(p1) > 0) {
                    p1 = (p1 + p1)
                }
                switch ("hello" + ("foo bar" + ("hello" + "\n\t"))) {
                case "foo bar":
                    var v13 string = ("" + ("hello" + "hello"))
                    _ = v13
                default:
                    var v14 float32 = -2.25
                    _ = v14
                }
            }
        default:
            var v15 int32 = 0
            _ = v15
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 174
    v3 := make(map[string]bool)
    var v4 uint = 63
    var v16 int = 42
    v16 = 870
    v17 := GenS2{F0: true, F1: v2}
    switch ("foo bar" + ("foo bar" + "\n\t")) {
    case "\n\t":
        v18 := GenS1{F0: 0, F1: v16, F2: 42}
        _ = v18
    }
    v16 -= 128
    if (!(42 <= 489)) {
        var v19 int64 = 0
        var v20 uint8 = (v2 ^ 2)
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS3{F0: 2147483648, F1: 1}
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    } else {
        v16 -= v16
        v23 := GenS4{F0: 0}
        _ = v23
    }
    switch "\n\t" {
    case "\n\t":
        v4 = (v4 + (^(v4)))
    case "foo bar":
        switch (-(v16)) {
        case 4294967295:
            var v24 bool = false
            _ = v24
        default:
            v16 = (v16 << (((^(v4))) & 63))
        }
    case "a":
        v25 := []int{}
        _ = v25
    default:
        v2 -= 0
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        fmt.Println(v16, v4)
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
