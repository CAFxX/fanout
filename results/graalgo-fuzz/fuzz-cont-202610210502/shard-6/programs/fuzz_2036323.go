package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
type GenS4 struct { F0 int; F1 uint32; F2 int }
type GenS5 struct { F0 int32; F1 bool; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0(p0 int8) int16 {
    p0 -= (p0 + 42)
    fmt.Println(p0)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 float64 = ((-2.25 * ((((-2.25 / ((1e10 / ((float64((1.5 - ((-2.25 - 0.0) + 0.1))) * 1e10) + 0.0)) / 0.1)) / 1.5) / 1e10) / 1.0)) + (((1e10 + 3.14159) - 1e10) / 1e10))
        var v8 int32 = -29
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "hello" {
    case "":
        p0 -= -117
    case "hello":
        var v9 uint16 = 1
        _ = v9
    case "\n\t":
        switch -50 {
        case 32767:
            fmt.Println(p0)
        case -263:
            var v10 string = "foo bar"
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 6; fi11++ {
                p0 = 6
                p0 = (p0 | 1)
                v12 := []int{32768}
                _ = fi11
                _ = v12
            }
        default:
            var v13 bool = false
            _ = v13
        }
    }
    v14 := []int{306, -522, 42, 32767}
    _ = p0
    _ = v14
    return 32767
}
func genFn1(p0 uint32) bool {
    p0 = 2
    fmt.Println(p0)
    var v15 int8 = 3
    _ = p0
    _ = v15
    return false
}

func main() {
    v1 := []uint32{392, 42, 2147483647}
    v2 := make(map[uint32]int)
    v3 := []uint32{683}
    var v4 uint16 = 0
    v5 := make(map[int]int)
    for fi16 := 0; fi16 < 14; fi16++ {
        v4 += v4
        _ = fi16
    }
    var v17 uint32 = 42
    v5[32768] = -880
    if true {
        v18 := GenS1{F0: (v17 % ((2147483648) | 1)), F1: ""}
        v19 := GenS2{F0: -128}
        for fi20 := 0; fi20 < 16; fi20++ {
            v17 = (v17 - v17)
            _ = fi20
        }
        _ = v18
        _ = v19
    } else {
        if true {
            v21 := GenS2{F0: 49}
            v22 := GenS3{F0: v4, F1: (false || true), F2: true}
            _ = v21
            _ = v22
        } else {
            var v23 float32 = 3.14159
            if (-9223372036854775808 == 42) {
                v23 = (v23 / 0.1)
                v4 = (v4 << ((42) & 63))
            }
            _ = v23
        }
        if true {
            if (false || false) {
                var v24 uint32 = 586
                v17 = 1
                if (true && (false || false)) {
                    v4 = v4
                }
                _ = v24
            }
            var v25 uint32 = v17
            _ = v25
        }
    }
    v2[(v17 >> ((725) & 63))] = 1
    switch "hello" {
    case "a":
        v4 += v4
    default:
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := GenS4{F0: (^(fi26)), F1: v17, F2: 0}
            v28 := GenS5{F0: 547, F1: (223 <= 42), F2: -510}
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    var v29 string = ("\n\t" + ("hello" + ("foo bar" + "a")))
    var v30 int64 = 1
    var v31 uint64 = 2
    var v32 int16 = -761
    for fi33 := 0; fi33 < 9; fi33++ {
        if true {
            fmt.Println(v30)
            v4 = v4
            v34 := GenS6{F0: v17}
            _ = v34
        }
        var v35 int64 = v30
        switch v29 {
        case "hello":
            var v36 bool = (false && false)
            _ = v36
        default:
            for fi37 := 0; fi37 < 19; fi37++ {
                if true {
                    fmt.Println(v29, v35)
                } else {
                    var v38 uint64 = v31
                    v4 = (v4 | v4)
                    var v39 uint16 = v4
                    _ = v38
                    _ = v39
                }
                v40 := GenS7{F0: v29, F1: v29}
                v41 := GenS8{F0: v17, F1: false, F2: 54}
                _ = fi37
                _ = v40
                _ = v41
            }
        }
        _ = fi33
        _ = v35
    }
    fmt.Println(v17)
    for fi42 := 0; fi42 < 10; fi42++ {
        fmt.Println(v29)
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v43 := v5[(fi42 >> ((193) & 63))]
            delete(v2, uint32(v31))
            _ = v43
        }
        v29 = "\n\t"
        _ = fi42
    }
    if (false && (true || false)) {
        for fi44 := 0; fi44 < 12; fi44++ {
            v45 := GenS5{F0: 0, F1: false, F2: (v30 - v30)}
            v46 := GenS8{F0: v17, F1: false, F2: 255}
            fmt.Println(fi44)
            _ = fi44
            _ = v45
            _ = v46
        }
        v32 += (-(v32))
    }
    v31 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
