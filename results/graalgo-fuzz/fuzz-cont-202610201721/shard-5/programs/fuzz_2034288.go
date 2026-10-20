package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int16, p1 uint16, p2 int8) int64 {
    v6 := GenS1{F0: p0, F1: 941}
    var v7 int8 = (-(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint) int8 {
    var v8 uint = (^(p0))
    p0 += p0
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]uint32)
    var v3 int64 = -360
    v4 := []uint32{(v1 | (^(v1))), (v1 - v1)}
    v5 := make(map[int]string)
    if (!false) {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: 507, F1: 413}
            if true {
                v11 := GenS1{F0: 918, F1: 0}
                v12 := GenS1{F0: -32768, F1: 2147483647}
                fmt.Println(fi9, v1, fi9)
                _ = v11
                _ = v12
            } else {
                v3 = (v3 % ((9223372036854775807) | 1))
                delete(v5, 851)
                v13 := GenS1{F0: 42, F1: -2147483648}
                _ = v13
            }
            if ((fi9 | (fi9 >> ((18446744073709551615) & 63))) >= fi9) {
                fmt.Println(fi9)
                for fi14 := 0; fi14 < 3; fi14++ {
                    v15 := GenS1{F0: 250, F1: -2147483648}
                    _ = fi14
                    _ = v15
                }
                v3 -= genFn0(32767, 42, genFn1(2))
            } else {
                var v16 int = fi9
                v1 -= 65535
                _ = v16
            }
            _ = fi9
            _ = v10
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(v1, fi17, fi18)
                v19 := GenS2{F0: genFn0(-32768, 2, -85)}
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    var v20 int64 = 42
    if len(v4) > 0 {
        v4[0] = v1
    }
    var v21 float32 = (float32(0.1) / (0.0 - float32(3.14159)))
    v1 = v1
    for fi22 := 0; fi22 < 14; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            if (-2147483648 >= -1) {
                v24 := GenS2{F0: v3}
                _ = v24
            } else {
                fmt.Println(fi23, v21, fi23)
                fmt.Println(fi22, v1, v1)
                var v25 uint32 = 2147483648
                _ = v25
            }
            _ = fi23
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS2{F0: 290}
            _ = fi26
            _ = v27
        }
        switch "hello" {
        case "hello":
            var v28 float64 = 1e10
            _ = v28
        case "":
            v21 += 1.5
        case "\n\t":
            v29 := GenS2{F0: v3}
            _ = v29
        }
        _ = fi22
    }
    var v30 float64 = (-2.25 * 3.14159)
    v31 := v2[127]
    for fi32 := 0; fi32 < 10; fi32++ {
        v30 = -2.25
        for fi33 := 0; fi33 < 2; fi33++ {
            v20 += -1
            _ = fi33
        }
        var v34 int16 = 0
        _ = fi32
        _ = v34
    }
    fmt.Println(v20, v31)
    v1 = (v31 | v31)
    for fi35 := 0; fi35 < 19; fi35++ {
        fmt.Println(v21, v31, v30)
        _ = fi35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 int8 = 0
        _ = fi36
        _ = v37
    }
    switch 206 {
    case 127:
        for fi38 := 0; fi38 < 20; fi38++ {
            if false {
                var v39 float32 = 3.14159
                var v40 uint32 = (v1 ^ v1)
                _ = v39
                _ = v40
            }
            v20 += (v20 + -322)
            _ = fi38
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v21, v21)
    fmt.Println(v30, v20, v1)
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
