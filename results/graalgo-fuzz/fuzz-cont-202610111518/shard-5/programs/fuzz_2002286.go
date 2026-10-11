package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint8; F2 int32 }
type GenS5 struct { F0 int64; F1 int; F2 uint }
type GenS6 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 uint16, p1 int64, p2 uint32) uint32 {
    p2 = p2
    if ((true && false) || true) {
        p2 = p2
    }
    p2 += 42
    var v6 int = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 ^ uint32(p0))
}
func genFn1(p0 int8, p1 int8, p2 uint64) int8 {
    var v7 int = 128
    p0 -= 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1e10
    var v3 int16 = 32767
    var v4 uint16 = 0
    v5 := make(map[int]uint32)
    v8 := GenS1{F0: 1, F1: -2147483648}
    v9 := GenS1{F0: -775, F1: 42}
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                v12 := GenS1{F0: -9223372036854775808, F1: -461}
                var v13 int = (fi11 + 128)
                switch (fi10 - v13) {
                case -1:
                    var v14 int8 = 1
                    _ = v14
                case 985:
                    fmt.Println(v4)
                case 42:
                    v1[(fi11 | (v13 - v13))] = true
                }
                _ = v12
                _ = v13
            } else {
                var v15 float64 = ((v2 + v2) + -2.25)
                _ = v15
            }
            v2 -= float64((-2.25 * 3.14159))
            _ = fi11
        }
        if false {
            var v16 uint = 65535
            v16 -= (v16 ^ 331)
            v17 := GenS1{F0: -1, F1: 2147483647}
            _ = v16
            _ = v17
        }
        _ = fi10
    }
    v4 += v4
    var v18 float32 = 1.5
    fmt.Println(v18, v2)
    v2 += (((v2 - v2) - v2) - float64((v18 + v18)))
    v4 += (v4 ^ (v4 << ((42) & 63)))
    for fi19 := 0; fi19 < 8; fi19++ {
        if false {
            var v20 int8 = -48
            v20 -= (v20 / ((100) | 1))
            _ = v20
        }
        delete(v5, (fi19 / ((fi19) | 1)))
        _ = fi19
    }
    if (((14 <= 2) && (v18 <= v18)) && true) {
        v3 += v3
    } else {
        v21 := GenS1{F0: 42, F1: -1}
        var v22 int32 = 0
        if ((true || true) || true) {
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 uint32 = 1
                _ = fi23
                _ = v24
            }
            v25 := []int{9223372036854775807, 9223372036854775807, 836}
            if (!true) {
                if true {
                    switch "" {
                    case "foo bar":
                        v18 += ((0.1 * v18) * (1.5 + (v18 - v18)))
                    case "":
                        v2 = ((v2 + (v2 + ((v2 * v2) / v2))) * v2)
                    case "hello":
                        var v26 int8 = genFn1(-128, 0, 538)
                        _ = v26
                    default:
                        v27 := GenS1{F0: 9223372036854775807, F1: (-(v22))}
                        _ = v27
                    }
                    v28 := GenS1{F0: -1, F1: (v22 >> ((42) & 63))}
                    _ = v28
                }
                if (false && (!(-1 <= 2147483647))) {
                    v3 -= v3
                    var v29 bool = (149 < 1)
                    _ = v29
                }
                v2 = (-2.25 * (v2 * float64(1.5)))
            } else {
                if (503 == 1) {
                    v2 = (v2 * (0.0 + v2))
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v3 = (v3 ^ v3)
                        _ = fi30
                    }
                } else {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        switch fi31 {
                        case -836:
                            v32 := GenS1{F0: 887, F1: (v22 + 2147483647)}
                            _ = v32
                        case -2147483648:
                            v22 = v22
                        case -76:
                            for fi33 := 0; fi33 < 3; fi33++ {
                                v22 = (v22 & v22)
                                v2 += (v2 * (v2 + v2))
                                _ = fi33
                            }
                        }
                        switch fi31 {
                        case -9223372036854775808:
                            if false {
                                v22 += v22
                            } else {
                                for fi34 := 0; fi34 < 17; fi34++ {
                                    var v35 float64 = (float64(0.0) - (1.5 * (v2 / 3.14159)))
                                    var v36 float32 = (v18 - 1.5)
                                    fmt.Println(v36, v18)
                                    var v37 int64 = 1
                                    if (!((v36 >= v18) || (1 <= uint8(v4)))) {
                                        fmt.Println(v18, v36)
                                        if (true && false) {
                                            switch fi31 {
                                            case 32767:
                                                switch "a" {
                                                case "\n\t":
                                                    v38 := GenS1{F0: (-(v37)), F1: -949}
                                                    _ = v38
                                                }
                                            default:
                                                var v39 float64 = v35
                                                _ = v39
                                            }
                                            if true {
                                                switch ("\n\t" + "foo bar") {
                                                case "foo bar":
                                                    v40 := GenS2{F0: (!(float64(v18) != (v2 + (((v2 - v2) + -2.25) + (v2 * 0.0)))))}
                                                    if false {
                                                        v41 := GenS2{F0: (2 >= 2)}
                                                        v37 = 9223372036854775807
                                                        _ = v41
                                                    }
                                                    _ = len(v1)
                                                    _ = v40
                                                case "\n\t":
                                                    v25 = append(v25, (fi34 << ((2) & 63)))
                                                case "a":
                                                    _ = len(v25)
                                                }
                                            } else {
                                                switch (fi31 / ((-725) | 1)) {
                                                case 128:
                                                    var v42 int = 370
                                                    _ = v42
                                                case 4294967295:
                                                    if (42 < 2) {
                                                        var v43 float32 = 1.5
                                                        v18 = (3.14159 / v36)
                                                        var v44 float32 = (((((float32(((1e10 - v2) + float64(0.1))) + 0.0) - float32(float64(v43))) / v18) + 3.14159) / v43)
                                                        _ = len(v5)
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        if (true && (1 <= 127)) {
                                                            delete(v1, 4294967295)
                                                        }
                                                        if true {
                                                            v45 := GenS1{F0: v37, F1: (v22 - -451)}
                                                            fmt.Println(v36)
                                                            _ = v45
                                                        } else {
                                                            var v46 uint32 = 2
                                                            if ((v37 | 0) > -9223372036854775808) {
                                                                var v47 int16 = 42
                                                                v48 := GenS3{F0: -128}
                                                                v4 += (v4 + (v4 << ((42) & 63)))
                                                                v25 = append(v25, (fi34 & (-(fi31))))
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                fmt.Println(v37, v35)
                                                            }
                                                            var v49 int8 = 3
                                                            _ = v46
                                                            _ = v49
                                                        }
                                                        switch fi31 {
                                                        case -574:
                                                            if ((-954 < (v22 - (v22 | v22))) || false) {
                                                                v3 = 1
                                                            }
                                                        case -933:
                                                            var v50 uint32 = 2147483647
                                                            _ = v50
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                    }
                                                case 32768:
                                                    if (!((1 != v4) || false)) {
                                                        var v51 uint64 = uint64(v4)
                                                        _ = v51
                                                    } else {
                                                        v3 = v3
                                                        v52 := v1[42]
                                                        v53 := GenS3{F0: -118}
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                }
                                            }
                                            v54 := GenS4{F0: "", F1: 2, F2: v22}
                                            _ = v54
                                        }
                                        var v55 uint64 = 2
                                        _ = v55
                                    }
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                var v56 uint64 = 42
                                v57 := GenS1{F0: 42, F1: v22}
                                _ = v56
                                _ = v57
                            }
                        case -2147483648:
                            var v58 float32 = (float32(1e10) * float32((0.1 - (1e10 * ((1e10 / -2.25) + (1.5 + v2))))))
                            _ = v58
                        default:
                            var v59 int32 = (v22 << ((4294967295) & 63))
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 20; fi60++ {
                            var v61 float64 = v2
                            v18 += (v18 - v18)
                            v62 := v5[(fi31 / ((fi60) | 1))]
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        _ = fi31
                    }
                }
                v2 -= ((v2 + (v2 + v2)) + v2)
            }
            _ = v25
        } else {
            var v63 int = -283
            _ = v63
        }
        _ = v21
        _ = v22
    }
    var v64 uint = 65535
    v3 = (v3 >> (((v64 ^ v64)) & 63))
    v5[127] = 4294967295
    for fi65 := 0; fi65 < 6; fi65++ {
        if (false || true) {
            v66 := GenS5{F0: 982, F1: fi65, F2: v64}
            v3 += (v3 + (v3 - (v3 << ((v64) & 63))))
            _ = v66
        } else {
            switch "" {
            case "foo bar":
                for fi67 := 0; fi67 < 16; fi67++ {
                    v4 = (v4 - (v4 >> (((v64 / (((^(v64))) | 1))) & 63)))
                    var v68 uint32 = 4294967295
                    if true {
                        delete(v1, 127)
                        v69 := GenS6{F0: 127, F1: 819, F2: -2147483648}
                        v70 := GenS1{F0: -1, F1: -1}
                        _ = v69
                        _ = v70
                    }
                    _ = fi67
                    _ = v68
                }
            case "a":
                var v71 uint = 275
                _ = v71
            }
            if false {
                var v72 int8 = 127
                _ = v72
            } else {
                var v73 uint64 = 4294967295
                _ = v73
            }
        }
        var v74 int64 = 42
        _ = fi65
        _ = v74
    }
    fmt.Println(v4, v4, v64)
    fmt.Println(v64, v18, v18)
    v64 = 226
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v8
    _ = v9
}
