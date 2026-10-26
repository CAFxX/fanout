package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 int16 }
func genFn0() string {
    var v5 uint64 = 0
    var v6 uint64 = (v5 + 4294967295)
    for fi7 := 0; fi7 < 18; fi7++ {
        v6 += 156
        fmt.Println(v6, v5, v5)
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := []int{2147483647}
            v10 := GenS1{F0: "", F1: -9223372036854775808}
            _ = len(v9)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v5 = (v6 | v5)
        if true {
            fmt.Println(fi11, v5)
            for fi12 := 0; fi12 < 20; fi12++ {
                if false {
                    var v13 int = (fi12 * fi12)
                    switch "" {
                    case "hello":
                        var v14 int = (fi11 >> ((65535) & 63))
                        _ = v14
                    }
                    var v15 int = 32767
                    v5 -= (v6 >> ((255) & 63))
                    _ = v13
                    _ = v15
                } else {
                    fmt.Println(fi12, v5)
                    v6 = v6
                }
                var v16 uint64 = 0
                _ = fi12
                _ = v16
            }
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case 16:
        switch 737 {
        case -1:
            var v17 uint = 65535
            _ = v17
        case -801:
            v18 := []int{}
            _ = v18
        }
    case 289:
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 int32 = 0
            var v21 int = (fi19 >> ((18446744073709551615) & 63))
            _ = fi19
            _ = v20
            _ = v21
        }
    default:
        v5 = 1
    }
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1() bool {
    var v22 uint32 = 858
    switch "hello" {
    case "hello":
        for fi23 := 0; fi23 < 5; fi23++ {
            if ((1e10 / 3.14159) < -2.25) {
                v24 := []int{2147483647}
                v25 := GenS1{F0: "\n\t", F1: (fi23 | fi23)}
                v24 = append(v24, -316)
                _ = v24
                _ = v25
            }
            _ = fi23
        }
    case "a":
        fmt.Println(v22)
    case "foo bar":
        v26 := []int{-9223372036854775808, 127}
        _ = v26
    }
    _ = v22
    return (71 != 128)
}

func main() {
    var v1 float64 = (float64(1.5) + float64(1e10))
    var v2 uint16 = 1
    v3 := []int{1, 32768, 127, 32768}
    v4 := make(map[int]int)
    var v27 int8 = -89
    v1 -= v1
    if true {
        var v28 uint = 952
        _ = v28
    } else {
        for fi29 := 0; fi29 < 5; fi29++ {
            if (!(!(false || true))) {
                fmt.Println(v1, v1)
                for fi30 := 0; fi30 < 9; fi30++ {
                    v4[fi29] = (-(fi30))
                    _ = fi30
                }
                v31 := GenS2{F0: 42, F1: 183, F2: 65535}
                _ = v31
            } else {
                if (!false) {
                    v2 -= (v2 * v2)
                    fmt.Println(v1)
                } else {
                    v27 = (^(v27))
                    delete(v4, (fi29 & fi29))
                }
                var v32 float64 = 1e10
                var v33 int8 = -39
                _ = v32
                _ = v33
            }
            _ = len(v4)
            _ = fi29
        }
        if false {
            _ = len(v3)
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 int = fi34
                v27 = v27
                v35 = (fi34 % ((128) | 1))
                _ = fi34
                _ = v35
            }
            var v36 int64 = 5
            _ = v36
        } else {
            var v37 int64 = -882
            if (!(!(-107 <= v27))) {
                var v38 uint = 255
                var v39 int32 = -2147483648
                if true {
                    if (161 == int64(v38)) {
                        var v40 int8 = v27
                        v41 := GenS1{F0: "\n\t", F1: -985}
                        _ = v40
                        _ = v41
                    } else {
                        var v42 uint = (^(v38))
                        _ = v42
                    }
                    fmt.Println(v38, v38)
                    delete(v4, 4294967295)
                } else {
                    switch 1 {
                    case -212:
                        v1 = v1
                    case 4294967295:
                        fmt.Println(v1)
                    }
                }
                _ = v38
                _ = v39
            }
            _ = v37
        }
    }
    var v43 float32 = 0.1
    v3 = append(v3, 2147483647)
    for fi44 := 0; fi44 < 16; fi44++ {
        if (true && true) {
            var v45 uint32 = 0
            _ = v45
        } else {
            v4[(fi44 * fi44)] = (^(fi44))
        }
        v46 := GenS3{F0: 0}
        v47 := GenS2{F0: 462, F1: 1, F2: v2}
        _ = fi44
        _ = v46
        _ = v47
    }
    v27 = v27
    v27 = (v27 & v27)
    switch "" {
    case "\n\t":
        v27 = -128
    case "foo bar":
        var v48 uint64 = 4294967295
        _ = v48
    default:
        var v49 int32 = 42
        _ = v49
    }
    if ((v2 % ((2) | 1)) >= (v2 * v2)) {
        for fi50 := 0; fi50 < 3; fi50++ {
            v1 = (v1 - v1)
            v51 := GenS3{F0: uint8(v2)}
            _ = fi50
            _ = v51
        }
    }
    v52 := GenS4{F0: 0, F1: -367}
    for fi53 := 0; fi53 < 7; fi53++ {
        for fi54 := 0; fi54 < 2; fi54++ {
            v1 += ((v1 + ((float64(v43) * v1) / v1)) * v1)
            _ = fi54
        }
        switch ("hello" + "\n\t") {
        case "hello":
            v27 = v27
        case "":
            for fi55 := 0; fi55 < 13; fi55++ {
                if ((fi55 >> ((4294967295) & 63)) <= fi55) {
                    switch "a" {
                    case "":
                        if false {
                            var v56 int32 = -655
                            v2 += v2
                            var v57 int16 = -693
                            v58 := v4[fi53]
                            v2 = (^(v2))
                            if (1e10 >= 1e10) {
                                v56 = v56
                                _ = len(v3)
                            }
                            var v59 uint16 = (v2 - v2)
                            _ = v56
                            _ = v57
                            _ = v58
                            _ = v59
                        } else {
                            var v60 int64 = -1
                            if (42 != 79) {
                                if true {
                                    if (!(274 <= 4294967295)) {
                                        var v61 uint16 = 32767
                                        if len(v3) > 0 {
                                            v3[0] = (fi55 >> ((255) & 63))
                                        }
                                        var v62 int64 = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        fmt.Println(v2, v43, fi53)
                                        var v63 uint8 = 115
                                        fmt.Println(v1, v2)
                                        _ = v63
                                    }
                                    v43 = (float32(1.5) + float32(0.1))
                                    switch (fi55 % ((int(v2)) | 1)) {
                                    case 2147483647:
                                        v64 := GenS5{F0: 2}
                                        _ = v64
                                    case 937:
                                        var v65 int16 = 0
                                        _ = v65
                                    case 164:
                                        v2 += 346
                                    }
                                } else {
                                    var v66 int64 = v60
                                    _ = v66
                                }
                                if false {
                                    v67 := GenS1{F0: ("a" + "a"), F1: (^(fi53))}
                                    _ = v67
                                }
                            }
                            _ = v60
                        }
                    case "\n\t":
                        v4[(fi55 / (((fi55 - fi55)) | 1))] = fi55
                    case "hello":
                        if true {
                            if (!false) {
                                fmt.Println(fi55)
                            }
                        }
                    }
                    if false {
                        v68 := GenS4{F0: 1, F1: -1}
                        _ = v68
                    } else {
                        _ = len(v4)
                        var v69 uint64 = 1
                        _ = v69
                    }
                    fmt.Println(fi53, v27)
                } else {
                    if (65 >= 42) {
                        if (!(424 != 0)) {
                            v27 = (v27 * v27)
                            if false {
                                if (-9223372036854775808 > 9223372036854775807) {
                                    fmt.Println(fi53, v2)
                                    switch "a" {
                                    case "foo bar":
                                        v70 := GenS2{F0: 4294967295, F1: 1, F2: 606}
                                        v71 := GenS6{F0: fi53}
                                        if len(v3) > 0 {
                                            v3[0] = 127
                                        }
                                        fmt.Println(v2, v43)
                                        var v72 int8 = v27
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                } else {
                                    var v73 uint = 1
                                    fmt.Println(v27)
                                    _ = v73
                                }
                                if false {
                                    if ((!false) || (v27 >= (v27 ^ -59))) {
                                        fmt.Println(v1)
                                    } else {
                                        var v74 uint32 = 560
                                        v75 := v4[(fi55 | fi55)]
                                        v75 -= 734
                                        v4[fi53] = -1
                                        _ = v74
                                        _ = v75
                                    }
                                } else {
                                    switch fi53 {
                                    case 127:
                                        fmt.Println(v1)
                                    case 9223372036854775807:
                                        switch (fi53 % ((750) | 1)) {
                                        case 9223372036854775807:
                                            fmt.Println(fi55, fi55)
                                        }
                                    case 32768:
                                        v1 -= (((float64(v43) / 1e10) + (v1 + ((3.14159 / 0.1) / float64(float32((v1 / float64((v43 / 0.1)))))))) + float64((((((((float32(1.5) - v43) + v43) + (v43 + ((-2.25 + (v43 / 1.0)) * v43))) - v43) / (v43 - float32(((v1 + (v1 + v1)) + v1)))) * 0.0) / 1e10)))
                                    default:
                                        if ((1.5 - v43) < (v43 + v43)) {
                                            v76 := GenS5{F0: 42}
                                            v1 -= 1e10
                                            if ((v27 >> ((255) & 63)) != v27) {
                                                var v77 uint = 18446744073709551615
                                                if (!true) {
                                                    v78 := GenS7{F0: 728, F1: -32768}
                                                    v79 := GenS7{F0: 577, F1: 681}
                                                    _ = v78
                                                    _ = v79
                                                }
                                                fmt.Println(v43, fi53, v27)
                                                _ = v77
                                            } else {
                                                var v80 int16 = 367
                                                var v81 string = "a"
                                                var v82 uint32 = 855
                                                fmt.Println(v1, fi55, fi53)
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            }
                                            _ = v76
                                        } else {
                                            _ = len(v4)
                                            var v83 string = "\n\t"
                                            _ = v83
                                        }
                                    }
                                }
                                v43 = v43
                            } else {
                                v43 -= 0.0
                                v2 = (v2 >> ((2) & 63))
                                v3 = append(v3, fi55)
                            }
                            switch "foo bar" {
                            case "":
                                fmt.Println(v43)
                            case "a":
                                v1 -= 3.14159
                            default:
                                var v84 uint8 = 42
                                _ = v84
                            }
                        }
                        v1 += v1
                    }
                }
                var v85 bool = ((0.0 + float64(float32((v1 - 1e10)))) < v1)
                _ = fi55
                _ = v85
            }
        }
        _ = fi53
    }
    v86 := GenS4{F0: 42, F1: -132}
    var v87 int32 = 1
    v2 = (v2 + (v2 << ((255) & 63)))
    v43 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v52
    _ = v86
    _ = v87
}
