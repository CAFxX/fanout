package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64) int {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{-786}
        _ = fi6
        _ = v7
    }
    var v8 bool = (!(false || true))
    var v9 uint8 = 127
    _ = p0
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int16, p1 uint16, p2 string) int16 {
    var v10 int64 = -9223372036854775808
    var v11 uint = 65535
    if (128 != 42) {
        var v12 int32 = 2147483647
        v13 := GenS1{F0: 119, F1: (p1 ^ uint16(v11)), F2: v10}
        var v14 float64 = -2.25
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (p0 ^ int16(p1))
}

func main() {
    var v1 float64 = float64((0.0 / 1.0))
    var v2 int8 = 1
    var v3 int16 = 0
    var v4 uint16 = 42
    var v5 int16 = 1
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := []int{}
        if (true || (genFn0(411) == genFn0(1))) {
            switch ("hello" + ("foo bar" + ("foo bar" + "a"))) {
            case "foo bar":
                _ = len(v16)
            case "hello":
                var v17 int32 = 42
                _ = v17
            }
            fmt.Println(v2, v3)
        }
        _ = fi15
        _ = v16
    }
    var v18 uint64 = 643
    v19 := []int{}
    v20 := GenS2{F0: 4294967295}
    switch 32767 {
    case -15:
        fmt.Println(v18)
    case -2147483648:
        v21 := GenS1{F0: (v2 ^ (v2 - v2)), F1: v4, F2: 160}
        _ = v21
    default:
        v22 := GenS2{F0: 32768}
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        if true {
            var v24 int8 = -1
            var v25 uint8 = 128
            v4 = 32767
            _ = v24
            _ = v25
        } else {
            switch "\n\t" {
            case "\n\t":
                if (true || ((v2 <= v2) || ((1e10 > v1) && (0 >= 644)))) {
                    v26 := GenS1{F0: (v2 | v2), F1: (v4 - v4), F2: 1}
                    v19 = append(v19, (^(fi23)))
                    fmt.Println(v2)
                    _ = v26
                }
            case "a":
                v18 = (v18 >> ((0) & 63))
            case "":
                for fi27 := 0; fi27 < 14; fi27++ {
                    if false {
                        _ = len(v19)
                        var v28 uint16 = v4
                        if true {
                            var v29 float32 = (float32(-2.25) + ((float32(v1) * -2.25) - 1e10))
                            v30 := GenS2{F0: genFn0(v18)}
                            v1 -= ((v1 + (v1 * ((3.14159 / v1) * (v1 + v1)))) * 1e10)
                            _ = v29
                            _ = v30
                        } else {
                            v28 -= 609
                        }
                        _ = v28
                    }
                    _ = fi27
                }
            }
        }
        v31 := GenS3{F0: 18446744073709551615}
        for fi32 := 0; fi32 < 19; fi32++ {
            if false {
                if (genFn1((v5 * -32768), (^(v4)), "\n\t") > genFn1(v5, v4, ("" + ("foo bar" + "hello")))) {
                    v33 := GenS1{F0: -124, F1: (v4 * 65535), F2: -1}
                    switch ("" + "a") {
                    case "\n\t":
                        fmt.Println(v5)
                    case "a":
                        var v34 uint64 = v18
                        _ = v34
                    case "":
                        _ = len(v19)
                    default:
                        var v35 bool = false
                        _ = v35
                    }
                    _ = v33
                }
            } else {
                fmt.Println(fi32, fi32, fi32)
            }
            switch ("" + "a") {
            case "":
                v18 -= v18
            case "hello":
                var v36 uint16 = 255
                _ = v36
            default:
                v4 = v4
            }
            _ = fi32
        }
        _ = fi23
        _ = v31
    }
    var v37 int16 = genFn1((v3 - (v3 & genFn1((v3 << ((18446744073709551615) & 63)), v4, "hello"))), (v4 >> ((42) & 63)), "\n\t")
    var v38 uint32 = 2147483647
    _ = len(v19)
    v4 = 0
    if false {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    }
    v39 := GenS1{F0: v2, F1: (v4 + 42), F2: -1}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
