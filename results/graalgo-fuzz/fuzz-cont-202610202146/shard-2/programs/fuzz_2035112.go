package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 string, p1 int16, p2 bool) int64 {
    p1 = -566
    switch p0 {
    case "\n\t":
        v5 := []int{-2147483648, 127, -9223372036854775808}
        _ = v5
    default:
        v6 := GenS1{F0: 1, F1: 42}
        _ = v6
    }
    v7 := []int{}
    v8 := GenS1{F0: 36, F1: 1}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 int32) uint8 {
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            p1 -= -639
            p0 = 11
            _ = fi9
        }
    }
    var v10 int64 = 1
    p1 = (p1 & 42)
    _ = p0
    _ = p1
    _ = v10
    return 255
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int8 = 24
    var v3 float64 = 0.0
    var v4 int16 = -32768
    fmt.Println(v3, v3, v4)
    v1 += (v1 + (v1 << ((v1) & 63)))
    var v11 uint = v1
    if true {
        switch -144 {
        case -1:
            v12 := GenS1{F0: (^(v2)), F1: 112}
            _ = v12
        default:
            switch -418 {
            case 127:
                if ((v2 / (((-(v2))) | 1)) < v2) {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v1 = v1
                        v14 := []int{(fi13 >> (((v11 / ((v1) | 1))) & 63))}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v1, v1, v2)
                    v15 := GenS1{F0: (v2 - v2), F1: 40}
                    _ = v15
                }
            case 9223372036854775807:
                var v16 uint64 = 0
                _ = v16
            }
        }
        var v17 float32 = 1.5
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 int32 = int32(v4)
            if false {
                for fi20 := 0; fi20 < 13; fi20++ {
                    v3 = v3
                    _ = fi20
                }
                var v21 float32 = v17
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        var v22 string = "\n\t"
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint = v11
            v25 := []int{fi23}
            _ = fi23
            _ = v24
            _ = v25
        }
        v1 += (v1 * v11)
        _ = v22
    }
    v11 = v1
    var v26 int8 = v2
    var v27 int16 = (v4 ^ (v4 - 42))
    v28 := GenS2{F0: -9223372036854775808, F1: 1}
    if false {
        v29 := []int{64, -1, 152, 2147483647}
        switch 4294967295 {
        case 32767:
            switch 127 {
            case 1:
                for fi30 := 0; fi30 < 19; fi30++ {
                    switch (("hello" + "") + ("hello" + "a")) {
                    case "a":
                        v31 := GenS2{F0: 0, F1: 1000}
                        _ = v31
                    case "foo bar":
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 uint64 = 0
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            v11 = uint(v33)
                            _ = fi32
                            _ = v33
                        }
                    case "":
                        fmt.Println(v2, v26)
                    }
                    v2 -= -128
                    if ((fi30 != fi30) && (65535 != 65535)) {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v11)
                        case "hello":
                            for fi34 := 0; fi34 < 7; fi34++ {
                                var v35 uint64 = 42
                                if (float64((float32(float64(0.1)) * float32((1.5 * v3)))) >= 1e10) {
                                    var v36 int = (fi30 >> ((65535) & 63))
                                    v37 := GenS3{F0: v11}
                                    _ = v36
                                    _ = v37
                                }
                                _ = fi34
                                _ = v35
                            }
                        default:
                            if (((fi30 - (fi30 * fi30)) == 32767) && (v1 == 4294967295)) {
                                if (((fi30 ^ (fi30 | 127)) < (fi30 & fi30)) || (!true)) {
                                    v27 -= -1
                                }
                                v38 := GenS4{F0: fi30, F1: 867}
                                v26 = (-(v2))
                                _ = v38
                            }
                        }
                    } else {
                        for fi39 := 0; fi39 < 12; fi39++ {
                            var v40 int = (^(fi30))
                            v26 += (v26 << ((v1) & 63))
                            _ = fi39
                            _ = v40
                        }
                        v3 = (v3 / (float64(((float32(float64((3.14159 * (1.5 - 0.0)))) * (((0.0 + (float32((v3 * float64(((3.14159 - ((1.5 / 1e10) - 0.0)) * 1.5)))) - (-2.25 / 1e10))) / float32((((3.14159 - v3) * (v3 / (v3 - v3))) / float64(float32((3.14159 - v3)))))) / float32(float64(0.1)))) * 1.5)) - ((v3 / 1.0) / v3)))
                        if ((v1 >= (v1 & 42)) && false) {
                            for fi41 := 0; fi41 < 12; fi41++ {
                                v42 := GenS2{F0: -1, F1: 0}
                                _ = len(v29)
                                if len(v29) > 0 {
                                    v29[0] = (fi41 + (fi41 / (((fi41 + (fi41 % (((fi30 >> ((1) & 63))) | 1)))) | 1)))
                                }
                                _ = fi41
                                _ = v42
                            }
                            var v43 uint8 = 255
                            _ = v43
                        } else {
                            for fi44 := 0; fi44 < 20; fi44++ {
                                v27 += (v4 >> (((v11 ^ v11)) & 63))
                                _ = fi44
                            }
                            var v45 float64 = (v3 * 3.14159)
                            _ = v45
                        }
                    }
                    _ = fi30
                }
            case 0:
                v27 += (v4 | v4)
            }
        case -121:
            v27 += v27
        default:
            if (!(false || true)) {
                v46 := GenS5{F0: 32767}
                if false {
                    v11 = v1
                    for fi47 := 0; fi47 < 6; fi47++ {
                        for fi48 := 0; fi48 < 14; fi48++ {
                            if ((247 < 127) || true) {
                                v3 = v3
                            } else {
                                v4 = int16(fi48)
                            }
                            switch (("\n\t" + ("a" + "")) + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                switch fi48 {
                                case 127:
                                    v49 := GenS1{F0: 0, F1: 23}
                                    var v50 uint64 = 0
                                    _ = v49
                                    _ = v50
                                case 48:
                                    var v51 int32 = 93
                                    _ = v51
                                case 128:
                                    fmt.Println(v2, fi47)
                                default:
                                    v3 = v3
                                }
                            default:
                                switch (-(fi48)) {
                                case 127:
                                    v4 += v4
                                default:
                                    fmt.Println(v2, fi48, v11)
                                }
                            }
                            if len(v29) > 0 {
                                v29[0] = 103
                            }
                            _ = fi48
                        }
                        _ = fi47
                    }
                }
                _ = v46
            } else {
                var v52 uint32 = 104
                for fi53 := 0; fi53 < 11; fi53++ {
                    for fi54 := 0; fi54 < 8; fi54++ {
                        v27 += v27
                        if (v2 == v2) {
                            var v55 int64 = 0
                            _ = v55
                        }
                        _ = fi54
                    }
                    if len(v29) > 0 {
                        v29[0] = 128
                    }
                    v26 = v2
                    _ = fi53
                }
                var v56 bool = ((v4 <= v27) && true)
                _ = v52
                _ = v56
            }
        }
        _ = v29
    }
    fmt.Println(v26, v27, v27)
    if (true || true) {
        fmt.Println(v11, v3, v26)
        var v57 bool = false
        _ = v57
    } else {
        v58 := []int{-549, 839, -9223372036854775808}
        v59 := GenS1{F0: 100, F1: 2}
        _ = v58
        _ = v59
    }
    v60 := []int{-550, -2147483648, 32767, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v60
}
