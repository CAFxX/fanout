package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 uint32, p1 int32, p2 int64) int32 {
    var v7 uint = 65535
    p0 += p0
    p1 = -409
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1(p0 int32) int64 {
    if false {
        var v8 uint16 = 1
        _ = v8
    } else {
        switch "\n\t" {
        case "\n\t":
            var v9 int = 0
            _ = v9
        case "":
            var v10 uint8 = 42
            _ = v10
        }
    }
    if (true || true) {
        fmt.Println(p0)
        var v11 uint16 = 20
        _ = v11
    }
    _ = p0
    return 1
}

func main() {
    v1 := []int{299, 607}
    var v2 int32 = 489
    v3 := make(map[uint32]bool)
    var v4 int16 = 0
    var v5 uint8 = 207
    var v6 uint = 632
    switch -565 {
    case 32768:
        if (-128 != 0) {
            v4 += 1
            v5 -= v5
            var v12 int32 = v2
            _ = v12
        }
    case 0:
        v4 += 42
    case -599:
        v13 := v3[329]
        _ = v13
    default:
        v4 = v4
    }
    fmt.Println(v5)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint = 2
        var v16 float32 = 1.5
        for fi17 := 0; fi17 < 15; fi17++ {
            if (!true) {
                if (42 > (v15 >> (((^(v6))) & 63))) {
                    v1 = append(v1, -675)
                    fmt.Println(v5, v4, fi17)
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v18 := GenS1{F0: 391, F1: v5, F2: (v4 ^ 32767)}
                        _ = v18
                    case "":
                        var v19 int8 = 47
                        _ = v19
                    default:
                        v20 := GenS1{F0: 194, F1: v5, F2: v4}
                        _ = v20
                    }
                } else {
                    v2 = (v2 - (v2 << (((v6 + (v15 >> ((v6) & 63)))) & 63)))
                }
            }
            var v21 uint8 = (v5 / ((uint8(v15)) | 1))
            v6 = 559
            _ = fi17
            _ = v21
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v22 int16 = (^(v4))
    for fi23 := 0; fi23 < 17; fi23++ {
        if false {
            var v24 bool = (4294967295 == 76)
            v3[2147483648] = v24
            var v25 int32 = v2
            _ = v24
            _ = v25
        }
        _ = fi23
    }
    fmt.Println(v5)
    switch ("" + "foo bar") {
    case "hello":
        fmt.Println(v6, v22)
    case "a":
        v26 := GenS1{F0: 131, F1: v5, F2: -1}
        _ = v26
    case "foo bar":
        fmt.Println(v5)
    default:
        switch "hello" {
        case "foo bar":
            fmt.Println(v6, v22, v6)
        }
    }
    if (18446744073709551615 > 787) {
        for fi27 := 0; fi27 < 20; fi27++ {
            fmt.Println(v22)
            _ = fi27
        }
        v1 = append(v1, 0)
        fmt.Println(v22, v2, v6)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
