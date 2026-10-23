package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 int64 }
type GenS4 struct { F0 int32; F1 int16; F2 string }
func genFn0(p0 uint8) uint16 {
    switch "" {
    case "a":
        if ((!false) && true) {
            v4 := GenS1{F0: (0 == -9223372036854775808)}
            v5 := GenS1{F0: (!(((331 != 42) && false) || (804 < 32767)))}
            if (-75 != 16) {
                p0 -= p0
                for fi6 := 0; fi6 < 2; fi6++ {
                    if false {
                        var v7 int32 = 42
                        p0 = p0
                        _ = v7
                    } else {
                        fmt.Println(fi6, p0)
                        v8 := GenS1{F0: false}
                        var v9 int = fi6
                        _ = v8
                        _ = v9
                    }
                    if ((-(fi6)) > 781) {
                        var v10 int = (fi6 ^ fi6)
                        var v11 int = fi6
                        _ = v10
                        _ = v11
                    }
                    _ = fi6
                }
                if (42 < 42) {
                    p0 -= (p0 ^ p0)
                    p0 = p0
                } else {
                    var v12 uint32 = 93
                    var v13 int16 = 59
                    _ = v12
                    _ = v13
                }
            }
            _ = v4
            _ = v5
        } else {
            v14 := []int{-463, -9223372036854775808, 4294967295}
            _ = v14
        }
    }
    var v15 uint64 = 0
    v16 := []int{-291, 203, 32768}
    var v17 int16 = 42
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 284
}
func genFn1(p0 int, p1 int8, p2 int16) uint32 {
    fmt.Println(p0, p2, p1)
    v18 := []int{p0}
    var v19 int = -9223372036854775808
    var v20 bool = false
    fmt.Println(p0, v20, v20)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 uint32 = 2
    fmt.Println(v3)
    v3 -= (v3 - (v3 << ((18446744073709551615) & 63)))
    var v21 uint = 444
    for fi22 := 0; fi22 < 12; fi22++ {
        switch "" {
        case "hello":
            v1["hello"] = (v3 ^ v3)
        case "foo bar":
            if (true && (!(true || ((v21 < (v21 + (^(v21)))) || true)))) {
                if (150 < 42) {
                    v23 := GenS1{F0: (v3 < v3)}
                    fmt.Println(fi22, v3, v21)
                    v3 -= (v3 & (v3 ^ 4294967295))
                    _ = v23
                }
                if false {
                    v24 := v2["hello"]
                    v25 := GenS1{F0: (65535 != genFn0(128))}
                    _ = v24
                    _ = v25
                }
            }
        default:
            v21 = (^(v21))
        }
        var v26 uint = 255
        _ = fi22
        _ = v26
    }
    v27 := GenS2{F0: 729}
    v28 := GenS3{F0: true, F1: 0, F2: -9223372036854775808}
    v29 := GenS4{F0: -2147483648, F1: -32768, F2: "a"}
    v3 = 42
    v30 := GenS2{F0: 561}
    fmt.Println(v21, v21)
    var v31 int64 = -507
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 int = (fi32 | (fi32 ^ 1))
        _ = fi32
        _ = v33
    }
    var v34 uint16 = 475
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
