package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint16; F1 string }
func genFn0(p0 int, p1 uint16, p2 int32) uint32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = -2147483648
        var v8 int64 = 9223372036854775807
        fmt.Println(p1)
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -619 {
    case -956:
        fmt.Println(p1, p0, p0)
    }
    switch (-(p0)) {
    case 127:
        v9 := []int{(^(p0))}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 uint) string {
    fmt.Println(p0)
    var v10 bool = true
    _ = p0
    _ = v10
    return ""
}

func main() {
    var v1 float64 = -2.25
    var v2 uint = 255
    var v3 int64 = -975
    v4 := []int{-715, 32768, -198}
    v5 := make(map[int]bool)
    v11 := GenS1{F0: 119, F1: true, F2: 2147483647}
    var v12 uint64 = uint64(v2)
    var v13 uint32 = 89
    v3 = v3
    fmt.Println(v2)
    fmt.Println(v1, v2)
    v14 := GenS2{F0: -9223372036854775808, F1: v2, F2: "a"}
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS3{F0: v3}
            v18 := GenS4{F0: -772, F1: (^(fi15))}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 int32 = -1
        var v20 uint64 = (v12 / (((v12 - (v12 + 2))) | 1))
        _ = fi15
        _ = v19
        _ = v20
    }
    v3 += v3
    if (!(!false)) {
        v21 := GenS1{F0: -128, F1: true, F2: (v13 * v13)}
        var v22 int8 = 75
        _ = v21
        _ = v22
    }
    v12 -= v12
    v2 = v2
    v12 = v12
    var v23 string = ("hello" + "")
    switch "hello" {
    case "":
        var v24 uint = v2
        _ = v24
    default:
        var v25 bool = true
        _ = v25
    }
    v2 += (v2 + 0)
    switch (v23 + v23) {
    case "a":
        v5[-984] = false
    case "foo bar":
        switch "" {
        case "\n\t":
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 int16 = int16(v3)
                v1 -= 1e10
                switch -962 {
                case 9223372036854775807:
                    fmt.Println(v2)
                default:
                    switch 9223372036854775807 {
                    case 4294967295:
                        if (42 > 0) {
                            if (-42 == 0) {
                                if (false || (!true)) {
                                    var v28 int16 = int16(v3)
                                    _ = v28
                                }
                                v29 := GenS1{F0: -48, F1: (1 < 1), F2: (v13 % (((v13 >> (((v2 | (v2 ^ v2))) & 63))) | 1))}
                                _ = v29
                            }
                        } else {
                            v30 := GenS1{F0: 1, F1: true, F2: uint32(v2)}
                            switch (fi26 / (((fi26 * (fi26 & fi26))) | 1)) {
                            case -644:
                                v31 := GenS2{F0: (fi26 % (((^(fi26))) | 1)), F1: (v2 << ((v2) & 63)), F2: v23}
                                _ = v31
                            case 127:
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    if (v3 != 1) {
                                        if false {
                                            fmt.Println(fi32, v13, v3)
                                            _ = len(v4)
                                        }
                                        v27 -= -656
                                        v33 := GenS3{F0: (^(v3))}
                                        _ = v33
                                    }
                                    var v34 bool = ((((v23 + "\n\t") + v23) != ("" + "a")) || (127 != -1))
                                    var v35 bool = ((fi26 * 1) != -9223372036854775808)
                                    _ = fi32
                                    _ = v34
                                    _ = v35
                                }
                            case 42:
                                fmt.Println(v3)
                            default:
                                v36 := GenS1{F0: 127, F1: (!(!false)), F2: 65535}
                                _ = v36
                            }
                            if len(v4) > 0 {
                                v4[0] = fi26
                            }
                            _ = v30
                        }
                    case -96:
                        v37 := GenS1{F0: 42, F1: (true && true), F2: (v13 & (^(v13)))}
                        _ = v37
                    case 128:
                        if true {
                            v3 = v3
                            delete(v5, fi26)
                        }
                    }
                }
                _ = fi26
                _ = v27
            }
        }
    case "hello":
        if true {
            v38 := v5[42]
            v39 := GenS2{F0: 1, F1: (v2 | v2), F2: v23}
            _ = v38
            _ = v39
        } else {
            var v40 uint64 = (v12 + v12)
            v41 := GenS5{F0: 841, F1: v23}
            fmt.Println(v1, v13, v12)
            _ = v40
            _ = v41
        }
    default:
        var v42 float32 = (1.5 / 1.0)
        _ = v42
    }
    if (((v1 * v1) / 1.0) <= (1.5 - 3.14159)) {
        v43 := v5[127]
        _ = v43
    } else {
        var v44 int = -9223372036854775808
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
