package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int16; F1 int }

func main() {
    var v1 bool = ((!(0 != 0)) && false)
    var v2 bool = v1
    var v3 int32 = 2147483647
    var v4 uint8 = 127
    v5 := []uint32{42, 239, 463}
    var v6 float32 = 1.5
    var v7 int16 = 1000
    if (239 >= 113) {
        v8 := GenS1{F0: 2}
        v1 = (v1 && v1)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            v3 += (v3 << ((470) & 63))
            if (-2.25 > 3.14159) {
                if ((18446744073709551615 < 442) || true) {
                    fmt.Println(v2)
                } else {
                    var v10 uint8 = (v4 >> ((606) & 63))
                    var v11 int16 = 112
                    _ = v10
                    _ = v11
                }
                if (-1 > 37) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 = v3
                } else {
                    if v1 {
                        var v12 string = "hello"
                        var v13 int8 = -128
                        _ = len(v5)
                        _ = v12
                        _ = v13
                    } else {
                        var v14 int8 = -55
                        _ = v14
                    }
                    var v15 uint64 = 42
                    var v16 float32 = v6
                    _ = v15
                    _ = v16
                }
                if len(v5) > 0 {
                    v5[0] = 37
                }
            } else {
                v17 := GenS1{F0: 2}
                _ = v17
            }
            _ = fi9
        }
        var v18 uint8 = (^(v4))
        _ = v18
    }
    var v19 uint8 = (^(v4))
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS2{F0: true, F1: 42}
        _ = fi20
        _ = v21
    }
    fmt.Println(v4, v3)
    if ((v6 + v6) >= 0.1) {
        var v22 float64 = (0.0 + 1e10)
        switch "" {
        case "a":
            v19 -= (v19 ^ (^(v4)))
        case "\n\t":
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(v6)
                _ = fi23
            }
        default:
            var v24 int64 = 9223372036854775807
            _ = v24
        }
        _ = v22
    } else {
        if (true || (!(!v1))) {
            fmt.Println(v1)
            if false {
                var v25 int16 = v7
                for fi26 := 0; fi26 < 20; fi26++ {
                    v27 := GenS3{F0: 1000, F1: fi26}
                    if (((v2 && (65535 >= 42)) || (((-2.25 - 1e10) > v6) || (!v2))) && true) {
                        var v28 int32 = -787
                        var v29 uint = 42
                        _ = v28
                        _ = v29
                    } else {
                        v4 += (v19 & v4)
                        v1 = (fi26 <= (fi26 / ((fi26) | 1)))
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
            var v30 int64 = -9223372036854775808
            _ = v30
        }
        var v31 uint16 = 210
        _ = v31
    }
    var v32 uint32 = 0
    v4 += 0
    v33 := GenS1{F0: 279}
    var v34 uint64 = 324
    if v1 {
        v2 = (!v1)
        var v35 uint32 = v32
        _ = v35
    } else {
        var v36 uint8 = 0
        v3 += 0
        switch -40 {
        case 42:
            fmt.Println(v2, v3)
        }
        _ = v36
    }
    var v37 int = 42
    var v38 uint64 = (v34 + (^(v34)))
    v37 = v37
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
