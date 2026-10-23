package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64, p1 uint64, p2 uint32) int16 {
    v4 := []int{128, -2147483648}
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint = 42
        for fi7 := 0; fi7 < 18; fi7++ {
            p2 += p2
            v8 := GenS1{F0: (p0 / ((p0) | 1)), F1: 1, F2: 368}
            p2 += (^(p2))
            var v9 bool = (!false)
            var v10 uint8 = 128
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    p0 += (p0 - p0)
    fmt.Println(p2)
    var v11 uint8 = uint8(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return 32767
}

func main() {
    var v1 uint64 = 307
    v2 := make(map[string]string)
    var v3 int64 = -287
    v12 := []int{}
    var v13 int8 = -128
    v13 = (v13 + (v13 * (v13 ^ 45)))
    v14 := GenS2{F0: 0, F1: 0}
    fmt.Println(v1, v13)
    if (!(v1 != v1)) {
        if (true || true) {
            v13 += v13
        } else {
            v1 += (^(v1))
        }
        switch 0 {
        case -9223372036854775808:
            var v15 float64 = 0.0
            _ = v15
        default:
            v3 = 0
        }
        fmt.Println(v13, v13, v13)
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint16 = 42
            var v18 float64 = -2.25
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    switch "\n\t" {
    case "hello":
        var v19 uint64 = v1
        _ = v19
    case "a":
        v13 += -1
    case "foo bar":
        v20 := GenS3{F0: -1, F1: 255, F2: v13}
        _ = v20
    default:
        switch 127 {
        case 4294967295:
            var v21 uint = 0
            _ = v21
        case -1:
            var v22 int = -842
            _ = v22
        case -574:
            _ = len(v12)
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint = 0
        _ = len(v12)
        v25 := GenS4{F0: 32767, F1: 2147483647}
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 float32 = (-2.25 / 1.5)
        var v28 bool = ((false || false) && true)
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int8 = v13
    v30 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: -279}
    fmt.Println(v1, v1)
    var v31 bool = true
    _ = len(v12)
    var v32 int16 = 42
    if ((!v31) && (v31 && true)) {
        if v31 {
            v3 -= v3
            var v33 int = 32767
            _ = v33
        }
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            if ((v3 * (v3 / ((42) | 1))) != (-(v3))) {
                v32 = genFn0(v3, v1, 962)
                if len(v12) > 0 {
                    v12[0] = fi34
                }
                if ((((false && v31) && (0 != 0)) && false) && (-192 == 1)) {
                    for fi35 := 0; fi35 < 1; fi35++ {
                        v2[("\n\t" + "\n\t")] = "\n\t"
                        v31 = v31
                        v13 += v29
                        _ = fi35
                    }
                    if len(v12) > 0 {
                        v12[0] = fi34
                    }
                    v29 = (v29 << ((667) & 63))
                }
            }
            _ = fi34
        }
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 int8 = v29
            v13 -= v37
            _ = fi36
            _ = v37
        }
    }
    if (false || (42 != 18446744073709551615)) {
        v32 -= (v32 & v32)
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS1{F0: 42, F1: -2147483648, F2: 1}
            fmt.Println(v29)
            _ = len(v12)
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
