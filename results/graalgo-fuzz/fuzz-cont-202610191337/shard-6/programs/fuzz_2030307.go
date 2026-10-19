package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int16; F1 int16; F2 int8 }
type GenS7 struct { F0 int16; F1 uint16; F2 int32 }
type GenS8 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int64, p1 uint) uint32 {
    if false {
        p0 += (p0 & -1)
    }
    fmt.Println(p0)
    if true {
        var v5 float64 = 1.5
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int64 = p0
        switch (^(fi6)) {
        case 127:
            for fi8 := 0; fi8 < 3; fi8++ {
                v7 = (^(p0))
                _ = fi8
            }
        case 1:
            v9 := GenS1{F0: -1, F1: 255, F2: -747}
            _ = v9
        case -2147483648:
            p1 += p1
        default:
            v10 := GenS2{F0: 267}
            _ = v10
        }
        var v11 int16 = 381
        _ = fi6
        _ = v11
        _ = v7
    }
    _ = p0
    _ = p1
    return 150
}
func genFn1(p0 int16, p1 uint8) int8 {
    fmt.Println(p0, p1)
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: 9}
    var v15 bool = (("\n\t" != ("hello" + ((("\n\t" + (("\n\t" + ("\n\t" + "foo bar")) + "a")) + "") + ((("hello" + "a") + ((("foo bar" + "foo bar") + (("" + "\n\t") + (("\n\t" + "foo bar") + ("foo bar" + "\n\t")))) + "\n\t")) + "hello")))) && (127 != 2147483647))
    v16 := GenS4{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 116
}

func main() {
    var v1 uint32 = 1
    v2 := []int{}
    var v3 int32 = -2147483648
    var v4 uint32 = 4294967295
    for fi17 := 0; fi17 < 13; fi17++ {
        if len(v2) > 0 {
            v2[0] = (^(fi17))
        }
        var v18 int = fi17
        v3 = v3
        _ = fi17
        _ = v18
    }
    v3 += (v3 & v3)
    var v19 uint64 = 1
    switch int(v19) {
    case 32767:
        var v20 int16 = 32767
        _ = v20
    default:
        v3 -= v3
    }
    switch 42 {
    case -1:
        _ = len(v2)
    case 1:
        switch 192 {
        case 32768:
            v21 := GenS1{F0: -86, F1: 0, F2: -696}
            _ = v21
        case 32767:
            v22 := GenS1{F0: -36, F1: 42, F2: 1}
            _ = v22
        case -9223372036854775808:
            if true {
                v23 := GenS5{F0: 2, F1: 42}
                for fi24 := 0; fi24 < 13; fi24++ {
                    v4 -= (v4 ^ 20)
                    _ = fi24
                }
                fmt.Println(v4, v3, v3)
                _ = v23
            }
        }
    default:
        _ = len(v2)
    }
    switch 127 {
    case 2147483647:
        v25 := GenS6{F0: 1, F1: 42, F2: 0}
        _ = v25
    case 0:
        v2 = append(v2, 32767)
    }
    v2 = append(v2, 899)
    if ("" != "a") {
        v1 += (v4 / ((v1) | 1))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v4 = (v1 ^ genFn0(891, 42))
    if ((4294967295 >= 1) || false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        switch "hello" {
        case "a":
            for fi27 := 0; fi27 < 14; fi27++ {
                if (66 < 0) {
                    fmt.Println(fi27, fi26, fi26)
                    var v28 uint8 = 142
                    v4 = v4
                    _ = v28
                } else {
                    var v29 uint16 = 894
                    if (!false) {
                        fmt.Println(v19)
                        var v30 int32 = v3
                        _ = v30
                    } else {
                        v3 = 2147483647
                        v4 = genFn0(-974, 2)
                    }
                    switch (fi26 | (fi27 << ((18446744073709551615) & 63))) {
                    case 32768:
                        v31 := GenS2{F0: (^(v3))}
                        _ = v31
                    default:
                        switch "" {
                        case "":
                            var v32 float64 = (((((0.1 / 1.0) + 0.0) / 1.5) - 0.0) - (float64((((float32(((-2.25 + (((1e10 / ((float64(-2.25) + 0.0) + float64(0.0))) - 1.5) - (1.5 * (1e10 - float64(1e10))))) / float64(((1.5 - ((-2.25 - (((1e10 + float32(float64((0.1 + 0.0)))) + (0.0 - 0.0)) * 0.0)) / 1.5)) - 1.5)))) / 1.0) * float32((3.14159 / 1.0))) * 0.1)) / 1.0))
                            _ = v32
                        case "\n\t":
                            v19 -= (v19 >> ((4294967295) & 63))
                        default:
                            v33 := GenS7{F0: 0, F1: (v29 & (v29 + (v29 ^ (^(v29))))), F2: v3}
                            _ = v33
                        }
                    }
                    v29 += (v29 + v29)
                    if len(v2) > 0 {
                        v2[0] = (-(fi27))
                    }
                    _ = v29
                }
                _ = fi27
            }
        default:
            v4 = 2147483647
        }
        _ = fi26
    }
    switch 32768 {
    case -690:
        v34 := GenS8{F0: 42, F1: (v1 ^ (^(v4))), F2: uint8(v1)}
        _ = v34
    case -9223372036854775808:
        v2 = append(v2, 889)
    case -2147483648:
        v1 -= v1
    default:
        v4 = v4
    }
    v2 = append(v2, 63)
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(fi35)
        var v36 uint64 = v19
        if (!false) {
            if (false && ((v19 / (((^(v19))) | 1)) <= 2)) {
                switch (fi35 % ((1) | 1)) {
                case 4294967295:
                    v37 := GenS9{F0: 0, F1: 128, F2: 18446744073709551615}
                    _ = v37
                case 128:
                    var v38 bool = false
                    _ = v38
                default:
                    var v39 int = 495
                    _ = v39
                }
            }
            v1 = (v4 << ((2) & 63))
            if (!(0 < uint8(v1))) {
                v19 = v36
                var v40 int32 = (v3 - v3)
                v2 = append(v2, fi35)
                _ = v40
            } else {
                fmt.Println(v1, v19, v1)
                v41 := GenS10{F0: uint64(v1), F1: 2147483647}
                if (!(false && false)) {
                    var v42 int16 = 32767
                    _ = v42
                }
                _ = v41
            }
        }
        _ = fi35
        _ = v36
    }
    v1 -= genFn0(-9223372036854775808, 96)
    v43 := GenS2{F0: v3}
    v44 := GenS3{F0: 694}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
