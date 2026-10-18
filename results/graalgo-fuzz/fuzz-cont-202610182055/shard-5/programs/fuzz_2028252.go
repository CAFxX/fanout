package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int, p1 int8) uint16 {
    v5 := []int{p0, (^(p0)), p0}
    switch "" {
    case "a":
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(p0, fi6, fi6)
            if true {
                _ = len(v5)
                if (616 <= 4294967295) {
                    fmt.Println(p0, fi6)
                    for fi7 := 0; fi7 < 18; fi7++ {
                        var v8 int = int(p1)
                        _ = fi7
                        _ = v8
                    }
                } else {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v10 := GenS1{F0: 0, F1: 1000}
                        p1 = p1
                        _ = fi9
                        _ = v10
                    }
                    var v11 string = ("foo bar" + ("a" + "foo bar"))
                    _ = v11
                }
            } else {
                v12 := GenS2{F0: -69, F1: 18446744073709551615}
                var v13 float32 = 0.0
                fmt.Println(p1, fi6, p0)
                _ = v12
                _ = v13
            }
            v5 = append(v5, 42)
            _ = fi6
        }
    default:
        p1 = (p1 * (p1 | (^(p1))))
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 bool, p2 uint) int8 {
    p0 = 128
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]bool)
    var v3 uint16 = 0
    var v4 uint32 = 42
    v14 := []int{}
    var v15 float64 = (float64((0.0 - -2.25)) + 0.0)
    var v16 int16 = int16(v3)
    var v17 uint64 = 42
    v1 = (v1 ^ 2)
    var v18 bool = (1 <= 817)
    fmt.Println(v15)
    if v18 {
        for fi19 := 0; fi19 < 11; fi19++ {
            v18 = (!(true || ((v17 - v17) > (v17 | uint64(v1)))))
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 bool = (uint8(v17) == 230)
                var v22 int64 = 1
                v23 := GenS2{F0: -37, F1: v17}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            var v24 int32 = -681
            _ = fi19
            _ = v24
        }
    } else {
        if v18 {
            if v18 {
                _ = len(v14)
                var v25 int8 = 0
                v17 -= v17
                _ = v25
            } else {
                v26 := v2[128]
                fmt.Println(v15, v18, v4)
                _ = v26
            }
            _ = len(v2)
        } else {
            v4 = v4
            v27 := GenS3{F0: int16(v3), F1: -1, F2: uint64(v1)}
            switch ("hello" + "a") {
            case "foo bar":
                var v28 int8 = 42
                _ = v28
            case "hello":
                v3 = (v3 - v3)
            default:
                if v18 {
                    if v18 {
                        v29 := GenS1{F0: 42, F1: 32767}
                        var v30 uint16 = 683
                        switch 687 {
                        case 4294967295:
                            var v31 int8 = 1
                            _ = v31
                        case 1:
                            var v32 int8 = 100
                            _ = v32
                        default:
                            var v33 bool = v18
                            _ = v33
                        }
                        _ = v29
                        _ = v30
                    } else {
                        _ = len(v2)
                        v4 = (v4 / (((^(v4))) | 1))
                        var v34 uint32 = (v4 % ((v4) | 1))
                        _ = v34
                    }
                }
            }
            _ = v27
        }
    }
    switch -851 {
    case 1:
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS3{F0: (v16 * (v16 / (((v16 & (-(v16)))) | 1))), F1: (v16 / ((v16) | 1)), F2: 858}
            if (((0.1 * float64((0.0 / (3.14159 - (0.0 * 1.5))))) + (float64(1.5) * 0.0)) > v15) {
                v17 -= v17
                var v37 uint8 = 127
                v2[4294967295] = v18
                _ = v37
            }
            _ = fi35
            _ = v36
        }
    case -444:
        v3 += genFn0(-858, -1)
    case 32767:
        _ = len(v2)
    }
    var v38 int16 = 0
    for fi39 := 0; fi39 < 5; fi39++ {
        v2[(fi39 >> ((v1) & 63))] = false
        if false {
            delete(v2, int(v3))
            if (-128 > -1) {
                v40 := GenS4{F0: -722}
                switch "a" {
                case "a":
                    for fi41 := 0; fi41 < 2; fi41++ {
                        var v42 int = (fi41 & fi41)
                        _ = fi41
                        _ = v42
                    }
                case "\n\t":
                    var v43 int64 = -355
                    _ = v43
                }
                v4 -= v4
                _ = v40
            } else {
                fmt.Println(v1, v1)
                v1 -= (v1 ^ v1)
            }
        }
        _ = fi39
    }
    v18 = false
    fmt.Println(v3, v17, v17)
    v14 = append(v14, -2147483648)
    v3 = v3
    for fi44 := 0; fi44 < 8; fi44++ {
        fmt.Println(v17)
        fmt.Println(v1)
        var v45 uint64 = (v17 | (^(v17)))
        _ = fi44
        _ = v45
    }
    v17 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
