package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
func genFn0(p0 int16, p1 uint32) uint8 {
    var v6 uint8 = 42
    fmt.Println(p0, v6)
    p0 = 0
    var v7 int16 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 string, p2 int) uint {
    p0 = 1
    v8 := GenS1{F0: -1, F1: (p0 - 44), F2: 231}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 bool = true
    v2 := []uint32{65535}
    var v3 uint16 = 48
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v3 += 255
    if v1 {
        v5[678] = (("\n\t" + ("\n\t" + "a")) + ("a" + "hello"))
        var v9 uint16 = v3
        _ = v9
    } else {
        v10 := GenS1{F0: 9, F1: -1, F2: uint8(v3)}
        _ = v10
    }
    var v11 int8 = -128
    switch "foo bar" {
    case "":
        var v12 int = 127
        _ = v12
    default:
        var v13 float32 = -2.25
        _ = v13
    }
    var v14 uint = 536
    if v1 {
        var v15 int32 = 168
        v16 := GenS2{F0: "\n\t", F1: (v3 / (((v3 | (v3 >> ((v14) & 63)))) | 1)), F2: false}
        _ = v15
        _ = v16
    }
    _ = len(v2)
    var v17 float32 = 0.1
    v3 = v3
    v14 -= 4294967295
    if (42 != 9223372036854775807) {
        v18 := GenS2{F0: ("foo bar" + ("foo bar" + "hello")), F1: (v3 >> ((v14) & 63)), F2: (uint8(v14) >= 255)}
        _ = v18
    } else {
        if ((((v1 || ((v14 - v14) >= (v14 >> ((v14) & 63)))) || (!(255 > v14))) && true) && v1) {
            if (!false) {
                var v19 int64 = -1
                _ = v19
            }
            fmt.Println(v3, v11, v11)
        }
        switch "hello" {
        case "a":
            fmt.Println(v3)
        case "foo bar":
            var v20 uint = (v14 + v14)
            _ = v20
        default:
            fmt.Println(v3, v17, v17)
        }
        switch 1 {
        case -1:
            fmt.Println(v3)
        case -336:
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 int64 = 9223372036854775807
                _ = fi21
                _ = v22
            }
        case 32767:
            v14 = (v14 >> ((v14) & 63))
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v23 int16 = 440
    v14 -= 4294967295
    for fi24 := 0; fi24 < 6; fi24++ {
        switch (^(fi24)) {
        case 9223372036854775807:
            v2 = append(v2, 42)
        default:
            var v25 int8 = v11
            _ = v25
        }
        _ = fi24
    }
    fmt.Println(v23)
    v1 = false
    switch "" {
    case "\n\t":
        _ = len(v2)
    }
    var v26 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
