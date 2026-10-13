package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }

func main() {
    var v1 int8 = 1
    var v2 uint = 708
    var v3 uint64 = 18446744073709551615
    var v4 int64 = -9223372036854775808
    var v5 int16 = 1
    var v6 uint32 = 369
    var v7 uint64 = (v3 - v3)
    v4 = v4
    var v8 uint64 = 4294967295
    var v9 bool = false
    fmt.Println(v7, v8)
    if (!true) {
        v3 = 2
        v9 = v9
        fmt.Println(v9, v6, v7)
    } else {
        var v10 int32 = 301
        _ = v10
    }
    fmt.Println(v9)
    var v11 int = 13
    v12 := []int{(-(v11)), (v11 + (v11 | v11)), 4294967295}
    switch "a" {
    case "hello":
        v13 := GenS1{F0: (-(v1)), F1: (v8 / (((^(v3))) | 1))}
        _ = v13
    case "a":
        var v14 int64 = (v4 * -1)
        _ = v14
    default:
        var v15 uint = 18446744073709551615
        _ = v15
    }
    v12 = append(v12, -9223372036854775808)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v16 int16 = (^(v5))
    if false {
        for fi17 := 0; fi17 < 1; fi17++ {
            v6 -= 2
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := GenS1{F0: (v1 - 0), F1: 2}
                _ = fi18
                _ = v19
            }
            fmt.Println(v8, v8)
            _ = fi17
        }
    } else {
        if (v9 || ((!v9) || false)) {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        if len(v12) > 0 {
            v12[0] = -2147483648
        }
    }
    switch "foo bar" {
    case "hello":
        var v20 int32 = -2147483648
        _ = v20
    }
    v21 := GenS2{F0: 2, F1: 85}
    for fi22 := 0; fi22 < 8; fi22++ {
        if v9 {
            v23 := GenS1{F0: v1, F1: (v3 / ((2) | 1))}
            if true {
                v24 := GenS1{F0: v1, F1: v3}
                fmt.Println(v5)
                _ = v24
            }
            switch ("foo bar" + (("a" + "") + "\n\t")) {
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    if v9 {
                        v2 += v2
                        v4 += 42
                    }
                }
            case "hello":
                if len(v12) > 0 {
                    v12[0] = (fi22 - v11)
                }
            case "\n\t":
                var v25 int = (v11 & 32767)
                _ = v25
            }
            _ = v23
        } else {
            if false {
                fmt.Println(v3, fi22, v5)
                if (v4 != v4) {
                    v16 = 0
                } else {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        var v27 uint64 = v7
                        v12 = append(v12, -448)
                        if (-549 < 457) {
                            v28 := GenS3{F0: (^(v5)), F1: v9, F2: (^(v6))}
                            _ = v28
                        } else {
                            v4 += (v4 * (v4 | v4))
                            var v29 uint = 443
                            _ = v29
                        }
                        _ = fi26
                        _ = v27
                    }
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    v4 -= (v4 / ((-9223372036854775808) | 1))
                }
                v7 = v7
            } else {
                v6 = 50
                v3 = 287
                var v30 int = v11
                _ = v30
            }
            var v31 int32 = 2147483647
            _ = v31
        }
        switch ("" + "\n\t") {
        case "hello":
            for fi32 := 0; fi32 < 20; fi32++ {
                v33 := GenS1{F0: v1, F1: v7}
                _ = fi32
                _ = v33
            }
        default:
            v5 = v5
        }
        if len(v12) > 0 {
            v12[0] = v11
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
