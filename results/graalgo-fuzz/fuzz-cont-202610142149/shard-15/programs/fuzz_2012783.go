package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int }
func genFn0(p0 uint8) int32 {
    var v4 uint = 0
    for fi5 := 0; fi5 < 3; fi5++ {
        if true {
            for fi6 := 0; fi6 < 19; fi6++ {
                if (true && (246 >= 0)) {
                    var v7 uint16 = 0
                    fmt.Println(v4)
                    _ = v7
                }
                p0 = (p0 * (p0 * p0))
                _ = fi6
            }
            v8 := GenS1{F0: (p0 | (p0 - (p0 >> ((v4) & 63)))), F1: 255, F2: 18446744073709551615}
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint32 = 2147483647
                var v11 float32 = (-2.25 + float32(-2.25))
                var v12 uint32 = 2147483647
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v8
        } else {
            if true {
                var v13 uint32 = 2147483648
                _ = v13
            } else {
                var v14 int64 = -411
                v15 := []int{(fi5 & fi5), 4294967295, (fi5 << (((v4 ^ v4)) & 63))}
                if (v14 > v14) {
                    v14 += v14
                    _ = len(v15)
                    if ((((((-2.25 * 0.1) * 1.5) * 3.14159) - (float32((3.14159 / 3.14159)) - float32((3.14159 * (-2.25 - -2.25))))) * (0.0 - (3.14159 * ((1e10 + (3.14159 + 1.5)) * (float32(0.0) - 3.14159))))) > 0.0) {
                        switch fi5 {
                        case 719:
                            fmt.Println(fi5)
                        case -838:
                            switch 697 {
                            case 32767:
                                v14 = 42
                            case 4294967295:
                                if (127 > (^(p0))) {
                                    p0 = (p0 ^ (p0 ^ p0))
                                    v15 = append(v15, fi5)
                                    var v16 int16 = 899
                                    _ = v16
                                }
                            case -9223372036854775808:
                                if true {
                                    fmt.Println(fi5)
                                } else {
                                    p0 += (^(p0))
                                    var v17 float64 = float64((3.14159 / 1.5))
                                    _ = v17
                                }
                            default:
                                v18 := GenS1{F0: (^(p0)), F1: 32767, F2: 18446744073709551615}
                                _ = v18
                            }
                        case -531:
                            v19 := GenS2{F0: fi5, F1: 0, F2: 127}
                            _ = v19
                        }
                    }
                }
                _ = v14
                _ = v15
            }
            p0 += (p0 + 42)
            var v20 int = 2147483647
            _ = v20
        }
        if false {
            var v21 float64 = (3.14159 * (((((float64(((0.0 - -2.25) * (1.5 - ((1.5 + ((-2.25 / ((-2.25 + 0.1) / 1.0)) / (1.5 + -2.25))) + (1e10 / (0.0 + 0.1)))))) / 1.5) * 3.14159) / (((((0.1 - 0.0) * 3.14159) / 1e10) - 3.14159) / 1e10)) + -2.25) / (3.14159 / 1.5)))
            _ = v21
        } else {
            p0 = (p0 >> ((v4) & 63))
        }
        _ = fi5
    }
    var v22 float64 = 0.0
    v4 = (v4 % (((^(v4))) | 1))
    v23 := []int{641, 32768, 527, 32767}
    _ = p0
    _ = v22
    _ = v23
    _ = v4
    return 1
}
func genFn1(p0 int, p1 uint8) uint64 {
    p1 -= (p1 % (((^(p1))) | 1))
    var v24 int64 = 719
    _ = p0
    _ = p1
    _ = v24
    return 4294967295
}

func main() {
    var v1 int8 = 0
    var v2 uint = 255
    var v3 uint = (v2 + (v2 << ((558) & 63)))
    v25 := []int{int(v1)}
    v1 = (v1 | v1)
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS3{F0: 42, F1: v3, F2: 87}
        v3 -= (v3 << ((v2) & 63))
        v2 -= v3
        _ = fi26
        _ = v27
    }
    v3 -= v3
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    v3 += (v2 ^ v3)
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    switch 127 {
    case 32767:
        v28 := GenS4{F0: (!false)}
        _ = v28
    }
    v29 := GenS3{F0: 0, F1: 784, F2: 2}
    fmt.Println(v3)
    v2 -= v3
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS5{F0: 719, F1: -9223372036854775808}
        _ = fi30
        _ = v31
    }
    switch -561 {
    case 1:
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 float32 = -2.25
            for fi34 := 0; fi34 < 16; fi34++ {
                _ = len(v25)
                _ = fi34
            }
            _ = fi32
            _ = v33
        }
    case -1:
        for fi35 := 0; fi35 < 8; fi35++ {
            fmt.Println(fi35, v2)
            _ = fi35
        }
    default:
        for fi36 := 0; fi36 < 5; fi36++ {
            fmt.Println(v1, v1, v2)
            var v37 int64 = 298
            _ = fi36
            _ = v37
        }
    }
    fmt.Println(v2, v3)
    var v38 bool = (65535 < 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v38
}
