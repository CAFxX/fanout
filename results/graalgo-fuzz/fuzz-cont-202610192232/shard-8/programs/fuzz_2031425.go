package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS5 struct { F0 int64; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int8; F2 int }
func genFn0() uint8 {
    v6 := []int{-765, 128, 4294967295}
    if false {
        var v7 int16 = -32768
        v6 = append(v6, 861)
        fmt.Println(v7)
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int16 = 665
            v10 += (v10 % ((42) | 1))
            v11 := GenS1{F0: -2147483648, F1: (v10 != -32768), F2: 61}
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 1, F1: false, F2: 128}
        _ = fi8
        _ = v12
    }
    _ = v6
    return 2
}
func genFn1(p0 string, p1 int8, p2 bool) uint8 {
    switch 304 {
    case -1:
        if (p2 || p2) {
            p2 = ((false || p2) && p2)
            var v13 uint16 = 1
            _ = v13
        } else {
            var v14 int64 = -9223372036854775808
            v15 := GenS2{F0: -2147483648, F1: 42}
            fmt.Println(v14, p2, p2)
            _ = v14
            _ = v15
        }
    case 127:
        v16 := GenS1{F0: 0, F1: p2, F2: 0}
        _ = v16
    case 32768:
        v17 := GenS1{F0: 2147483647, F1: true, F2: 33}
        _ = v17
    }
    if (!p2) {
        v18 := GenS2{F0: 0, F1: 609}
        var v19 uint32 = 251
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 961
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    var v5 uint32 = 2147483648
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v5, v2)
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint32 = v5
            v22 = (v5 % (((v22 << ((72) & 63))) | 1))
            v5 = 65535
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 int8 = 122
    for fi24 := 0; fi24 < 9; fi24++ {
        fmt.Println(v5, fi24, v5)
        var v25 uint = 255
        _ = fi24
        _ = v25
    }
    v2 = v2
    switch "" {
    case "\n\t":
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS3{F0: 2147483647}
            for fi28 := 0; fi28 < 13; fi28++ {
                fmt.Println(v5)
                v29 := GenS2{F0: -677, F1: 2}
                v2 += 42
                _ = fi28
                _ = v29
            }
            v30 := GenS1{F0: -2147483648, F1: false, F2: 26}
            _ = fi26
            _ = v27
            _ = v30
        }
    default:
        v23 -= v23
    }
    var v31 uint16 = 255
    var v32 uint8 = uint8(v31)
    switch "hello" {
    case "hello":
        v33 := []int{32768, 1, 32768, 127}
        _ = v33
    default:
        var v34 string = (("hello" + "\n\t") + "a")
        _ = v34
    }
    v4[32767] = 53
    for fi35 := 0; fi35 < 9; fi35++ {
        fmt.Println(fi35, v5, v5)
        switch "foo bar" {
        case "hello":
            for fi36 := 0; fi36 < 4; fi36++ {
                fmt.Println(v31)
                v1[(v5 ^ v5)] = (("hello" + "\n\t") + "\n\t")
                var v37 int = 128
                if ((((fi35 + (v37 - 42)) == fi35) || true) || (475 <= 1)) {
                    v23 += -128
                } else {
                    v38 := GenS4{F0: 42, F1: v32, F2: v2}
                    v39 := []int{fi36, -1, (fi36 + fi35)}
                    fmt.Println(v2, fi35, fi36)
                    _ = v38
                    _ = v39
                }
                switch "foo bar" {
                case "a":
                    fmt.Println(fi35)
                }
                _ = fi36
                _ = v37
            }
        case "foo bar":
            v40 := GenS1{F0: -734, F1: true, F2: (v32 >> ((1) & 63))}
            _ = v40
        case "":
            for fi41 := 0; fi41 < 10; fi41++ {
                var v42 float32 = -2.25
                v31 += 1
                v42 += ((v42 + float32((0.1 * (1.5 / 1.0)))) * 0.0)
                _ = fi41
                _ = v42
            }
        }
        var v43 uint = 723
        _ = fi35
        _ = v43
    }
    var v44 float32 = 0.1
    if (!((true || ((false && (v5 <= v5)) || false)) && (-516 < -567))) {
        var v45 string = "\n\t"
        v46 := v1[uint32(v32)]
        _ = v45
        _ = v46
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi47 := 0; fi47 < 11; fi47++ {
            fmt.Println(v2)
            if (42 < uint(v31)) {
                var v48 int = fi47
                _ = v48
            } else {
                if (1 >= -1) {
                    fmt.Println(v31, v5)
                    v49 := GenS3{F0: -2147483648}
                    var v50 int64 = 42
                    _ = v49
                    _ = v50
                } else {
                    fmt.Println(v5, v2, v31)
                    v51 := GenS5{F0: 9223372036854775807, F1: ((float32(((((-2.25 - float64((v44 - (float32(-2.25) / v44)))) * (1.5 * 1.5)) - 1.5) / 1e10)) * v44) >= ((v44 - 0.0) - ((3.14159 * v44) + float32(float64(v44))))), F2: 188}
                    _ = v51
                }
                switch ("a" + "foo bar") {
                case "\n\t":
                    v52 := []int{(fi47 - fi47)}
                    _ = v52
                case "":
                    delete(v3, "hello")
                }
                v44 = (v44 + ((1.5 + (((-2.25 - v44) - 3.14159) / ((v44 + v44) + 0.1))) * (((-2.25 / v44) / 1.5) / float32(-2.25))))
            }
            _ = fi47
        }
    case -284:
        v44 = (-2.25 / (0.1 / v44))
    default:
        if (int16(v2) == -500) {
            var v53 bool = false
            _ = v53
        } else {
            v2 = v2
            switch 32767 {
            case 4294967295:
                v23 = v23
            case 133:
                switch 1 {
                case 4294967295:
                    fmt.Println(v44)
                case -9223372036854775808:
                    if false {
                        v31 = (v31 | (^(v31)))
                    } else {
                        v54 := GenS6{F0: 0, F1: v23, F2: 32767}
                        v2 = -462
                        var v55 uint64 = 2
                        _ = v54
                        _ = v55
                    }
                default:
                    v56 := []int{2147483647, -743}
                    _ = v56
                }
            case -2147483648:
                v23 = v23
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v5
}
