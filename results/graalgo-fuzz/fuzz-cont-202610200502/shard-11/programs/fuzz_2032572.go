package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint8, p1 int8, p2 int32) uint32 {
    v5 := []int{9223372036854775807, int(p1), -552}
    p2 -= 0
    switch "foo bar" {
    case "hello":
        var v6 bool = (((true || (!(0 == 1000))) && false) || (42 <= 2))
        _ = v6
    case "\n\t":
        switch "a" {
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = -158
            }
        }
    default:
        v5 = append(v5, 32768)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}
func genFn1() uint32 {
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        v7 = v7
        _ = fi8
    }
    if (true && true) {
        v7 = v7
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint32 = 664
            v7 += 32767
            _ = fi9
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            switch (fi11 + fi11) {
            case -1:
                v7 += 65535
            case 9223372036854775807:
                for fi12 := 0; fi12 < 20; fi12++ {
                    fmt.Println(v7)
                    _ = fi12
                }
            default:
                fmt.Println(v7)
            }
            _ = fi11
        }
        v7 = (v7 ^ v7)
        var v13 int32 = 0
        _ = v13
    }
    if false {
        v14 := GenS1{F0: -9223372036854775808, F1: 42, F2: false}
        v15 := GenS1{F0: -1, F1: 1, F2: true}
        _ = v14
        _ = v15
    }
    _ = v7
    return 0
}

func main() {
    var v1 int = 8
    var v2 uint16 = 32767
    var v3 uint = 4294967295
    var v4 uint = (^(v3))
    if true {
        var v16 bool = (-688 > 505)
        switch 15 {
        case 955:
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS2{F0: 4294967295}
                v4 = (v4 % ((v3) | 1))
                if false {
                    var v19 uint8 = 127
                    fmt.Println(v2, v1)
                    v20 := []int{(fi17 << (((v3 & v3)) & 63)), -421, (fi17 - fi17), (v1 * 32767)}
                    _ = v19
                    _ = v20
                } else {
                    v3 += (^(v4))
                }
                _ = fi17
                _ = v18
            }
        case -2147483648:
            var v21 int64 = -198
            _ = v21
        case 930:
            v2 = (^(v2))
        default:
            v22 := GenS3{F0: 128, F1: 80, F2: -2147483648}
            _ = v22
        }
        _ = v16
    }
    if (57 >= 42) {
        v23 := GenS3{F0: 2, F1: 128, F2: -223}
        v1 = (v1 >> ((v3) & 63))
        _ = v23
    } else {
        switch "hello" {
        case "hello":
            fmt.Println(v3, v4, v3)
        }
        var v24 uint16 = v2
        v25 := []int{0, v1, v1}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v2 += (v2 / (((^(v2))) | 1))
        fmt.Println(v1, v3, v3)
        _ = fi26
    }
    v27 := GenS2{F0: genFn0(128, 89, 42)}
    if true {
        fmt.Println(v3)
        fmt.Println(v3, v1, v1)
        var v28 float64 = 0.1
        _ = v28
    }
    var v29 uint8 = 127
    v30 := []int{v1, v1, (v1 << ((uint(v29)) & 63)), (-(v1))}
    fmt.Println(v1, v29)
    if len(v30) > 0 {
        v30[0] = 127
    }
    _ = len(v30)
    v31 := GenS4{F0: -128}
    v32 := GenS3{F0: 128, F1: v29, F2: -902}
    if (-2.25 > (1.5 / 1e10)) {
        if (-2.25 >= ((1e10 / -2.25) - -2.25)) {
            switch int(v3) {
            case 2147483647:
                if ((1 >= 9223372036854775807) || false) {
                    var v33 bool = false
                    v4 = v4
                    _ = v33
                }
            case -1:
                v30 = append(v30, 0)
            case -9223372036854775808:
                var v34 uint8 = v29
                _ = v34
            default:
                var v35 bool = false
                _ = v35
            }
            v36 := GenS3{F0: (v29 ^ (v29 | 77)), F1: v29, F2: -2147483648}
            if len(v30) > 0 {
                v30[0] = (v1 << ((121) & 63))
            }
            _ = v36
        }
    } else {
        if false {
            v37 := GenS5{F0: -32768}
            var v38 uint8 = (^(v29))
            for fi39 := 0; fi39 < 20; fi39++ {
                var v40 bool = (false || false)
                v4 += (v4 * (v4 - v4))
                v2 -= (^(v2))
                _ = fi39
                _ = v40
            }
            _ = v37
            _ = v38
        }
        if false {
            fmt.Println(v4)
            var v41 int16 = 10
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v41
        } else {
            fmt.Println(v3, v29)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
