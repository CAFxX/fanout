package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
func genFn0() int {
    var v5 int = 128
    v5 = (v5 >> ((18446744073709551615) & 63))
    v6 := []int{(v5 | v5), 608, 32768}
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 bool, p1 string, p2 uint8) string {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch fi7 {
        case 32768:
            p1 = p1
        default:
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{(fi7 / (((fi8 % ((fi8) | 1))) | 1)), fi8, fi8}
                _ = fi8
                _ = v9
            }
        }
        if p0 {
            p0 = (!((float64(0.1) - (1.5 + (1.5 + 1e10))) == 1.5))
        }
        var v10 bool = p0
        _ = fi7
        _ = v10
    }
    var v11 uint32 = 967
    v12 := []int{127, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p1
}

func main() {
    var v1 uint8 = 42
    var v2 int8 = -128
    v3 := []uint32{4, uint32(v1), 0}
    var v4 int8 = v2
    if false {
        var v13 float32 = -2.25
        switch -1 {
        case -2147483648:
            if false {
                fmt.Println(v4, v13, v4)
                switch 2147483647 {
                case 32767:
                    var v14 uint16 = 42
                    _ = v14
                }
                v2 -= 0
            }
        case 819:
            _ = len(v3)
        case -689:
            v3 = append(v3, 4294967295)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v13
    } else {
        fmt.Println(v2, v2, v2)
        for fi15 := 0; fi15 < 17; fi15++ {
            if (4294967295 > 1) {
                var v16 float64 = 1e10
                _ = v16
            } else {
                v17 := GenS1{F0: 732, F1: 0}
                var v18 float64 = 0.1
                _ = v17
                _ = v18
            }
            var v19 int64 = -9223372036854775808
            switch 2147483647 {
            case 4294967295:
                v19 += int64(v1)
            default:
                for fi20 := 0; fi20 < 7; fi20++ {
                    v2 += (v2 ^ 47)
                    var v21 int32 = 42
                    v3 = append(v3, 1)
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi15
            _ = v19
        }
    }
    switch 127 {
    case 48:
        var v22 float64 = 0.1
        _ = v22
    default:
        fmt.Println(v1, v4)
    }
    v23 := GenS1{F0: 1, F1: -32768}
    var v24 float32 = 3.14159
    v25 := GenS2{F0: ("a" > "foo bar"), F1: true}
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS3{F0: 746}
        v28 := GenS2{F0: (65535 <= 449), F1: true}
        _ = fi26
        _ = v27
        _ = v28
    }
    v4 += (-(v2))
    var v29 float64 = -2.25
    if (1 > 2147483647) {
        for fi30 := 0; fi30 < 8; fi30++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 uint = 1
                var v33 bool = true
                switch fi30 {
                case -462:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -1:
                    v34 := GenS3{F0: (v32 & (v32 + (v32 - (v32 - (v32 >> ((v32) & 63))))))}
                    var v35 float32 = 1e10
                    _ = v34
                    _ = v35
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            v2 += (v2 >> ((2) & 63))
            _ = fi30
        }
        v36 := GenS4{F0: 675, F1: v1, F2: "foo bar"}
        var v37 int64 = 9223372036854775807
        _ = v36
        _ = v37
    }
    switch 127 {
    case 210:
        var v38 int64 = 9223372036854775807
        _ = v38
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v39 uint16 = 1
        _ = v39
    default:
        v40 := GenS4{F0: 4294967295, F1: (^(v1)), F2: ("\n\t" + "a")}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
