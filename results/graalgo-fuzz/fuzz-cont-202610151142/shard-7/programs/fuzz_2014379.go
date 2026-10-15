package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16, p1 int16) int16 {
    var v7 string = "hello"
    v8 := GenS1{F0: -137}
    switch v7 {
    case "foo bar":
        p1 = (p1 >> ((65535) & 63))
    case "a":
        p0 -= p1
    default:
        if false {
            p1 += p0
            switch v7 {
            case "a":
                p1 -= (p0 ^ (p0 * p0))
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 616
}

func main() {
    var v1 uint = 1
    v2 := make(map[string]bool)
    var v3 int64 = 9223372036854775807
    var v4 int16 = 0
    var v5 uint32 = 673
    var v6 int8 = -128
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int16 = v4
        _ = fi9
        _ = v10
    }
    var v11 int8 = -128
    switch "" {
    case "":
        switch "a" {
        case "hello":
            switch (("hello" + ("" + "hello")) + ("foo bar" + ("\n\t" + "hello"))) {
            case "":
                fmt.Println(v11, v6, v1)
            default:
                var v12 int16 = v4
                _ = v12
            }
        }
    case "hello":
        v13 := GenS2{F0: (v11 << ((v1) & 63)), F1: v5, F2: false}
        _ = v13
    default:
        v5 = v5
    }
    var v14 int16 = genFn0((v4 % ((1000) | 1)), (v4 & (v4 + -32768)))
    var v15 float32 = 1.5
    if (2 > 18446744073709551615) {
        v6 = -1
    } else {
        fmt.Println(v3)
        v4 -= genFn0((v14 + (-(v4))), (v14 | 1))
        var v16 float64 = 1e10
        _ = v16
    }
    if true {
        var v17 uint64 = 677
        var v18 uint = v1
        _ = v17
        _ = v18
    } else {
        _ = len(v2)
    }
    fmt.Println(v4, v11, v5)
    v15 = v15
    delete(v2, (("\n\t" + "foo bar") + "foo bar"))
    switch 2147483647 {
    case 42:
        var v19 string = "\n\t"
        _ = v19
    case 1:
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 float64 = 3.14159
            if false {
                v22 := GenS2{F0: (v6 << ((v1) & 63)), F1: v5, F2: (!(!(false || true)))}
                for fi23 := 0; fi23 < 4; fi23++ {
                    _ = len(v2)
                    v24 := GenS3{F0: -726, F1: v5, F2: v5}
                    if true {
                        var v25 uint16 = 837
                        _ = v25
                    }
                    v26 := GenS4{F0: v4}
                    var v27 int = fi23
                    fmt.Println(v14, v3, v11)
                    if (757 < 718) {
                        if true {
                            v28 := GenS3{F0: -1, F1: v5, F2: v5}
                            _ = v28
                        }
                    } else {
                        v29 := GenS5{F0: (-(v3))}
                        _ = v29
                    }
                    fmt.Println(v3)
                    var v30 uint8 = 255
                    v27 = v27
                    fmt.Println(v4, v30)
                    _ = fi23
                    _ = v24
                    _ = v26
                    _ = v27
                    _ = v30
                }
                var v31 int64 = (-(v3))
                _ = v22
                _ = v31
            }
            _ = fi20
            _ = v21
        }
    case 9223372036854775807:
        v32 := GenS3{F0: -1, F1: (^(v5)), F2: (v5 >> (((v1 % ((uint(v5)) | 1))) & 63))}
        _ = v32
    default:
        v33 := GenS1{F0: 2147483647}
        _ = v33
    }
    v6 += (v6 >> ((v1) & 63))
    var v34 int64 = 1
    _ = len(v2)
    v5 = (v5 | v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
