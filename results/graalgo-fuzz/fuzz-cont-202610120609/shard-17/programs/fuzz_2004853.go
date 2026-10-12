package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0(p0 int8, p1 int8) int8 {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p0 = p0
    v5 = append(v5, -1)
    _ = p0
    _ = p1
    _ = v5
    return -128
}

func main() {
    var v1 float32 = (3.14159 + float32(0.0))
    var v2 int8 = 42
    v3 := []int{969, 92, -129}
    v4 := make(map[int]int)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 19
        v1 -= ((v1 / -2.25) / v1)
        v7 -= 339
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 255
    if ((v8 << ((467) & 63)) != 176) {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int16 = -1
            _ = fi9
            _ = v10
        }
        fmt.Println(v8, v2, v2)
    }
    if ("\n\t" < "\n\t") {
        v11 := GenS1{F0: 2, F1: 0, F2: (v2 ^ 1)}
        _ = v11
    } else {
        var v12 int16 = 42
        fmt.Println(v1, v1, v12)
        fmt.Println(v1, v1, v8)
        _ = v12
    }
    if (0 < 786) {
        switch 4294967295 {
        case 1:
            if ("hello" != "a") {
                if (!false) {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        var v14 uint32 = 0
                        v2 -= (v2 / ((127) | 1))
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v1)
                } else {
                    v2 += v2
                }
            }
        case 349:
            switch "foo bar" {
            case "foo bar":
                v8 += 55
            case "":
                var v15 uint32 = 87
                _ = v15
            case "a":
                v16 := GenS2{F0: 42, F1: (0 < 603)}
                _ = v16
            }
        }
    } else {
        v2 -= 127
        if (((9223372036854775807 != -9223372036854775808) || (2 >= 4294967295)) && (false || (!false))) {
            var v17 uint16 = 828
            _ = v17
        }
        if ((302 == 39) || true) {
            v4[-998] = 1
            if (933 < 65535) {
                var v18 uint = 1
                _ = v18
            } else {
                v2 = v2
            }
            v19 := GenS3{F0: 1, F1: 1, F2: 1}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        _ = len(v3)
        _ = fi20
    }
    if false {
        fmt.Println(v2, v8)
    } else {
        var v21 uint = 18446744073709551615
        _ = v21
    }
    v22 := v4[32767]
    var v23 string = "foo bar"
    if false {
        var v24 string = v23
        var v25 int16 = -1
        for fi26 := 0; fi26 < 7; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v28 int32 = 0
                v24 = ("hello" + (v24 + "a"))
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        _ = v24
        _ = v25
    } else {
        switch v22 {
        case 128:
            v8 -= (^(v8))
        }
        _ = len(v4)
    }
    v29 := v4[v22]
    v30 := GenS3{F0: 111, F1: 42, F2: -1}
    var v31 int8 = 53
    if true {
        v4[v22] = (v22 & v29)
        switch v23 {
        case "hello":
            v32 := GenS4{F0: 443, F1: 65535, F2: v31}
            _ = v32
        case "foo bar":
            fmt.Println(v29, v1, v29)
        case "\n\t":
            v22 = (v22 * v29)
        default:
            var v33 uint32 = 2147483647
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v8
}
