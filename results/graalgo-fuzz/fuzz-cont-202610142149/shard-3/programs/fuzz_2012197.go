package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint64; F2 int32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS9 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int8, p1 uint64, p2 int32) int8 {
    if (!false) {
        var v4 uint = 2
        var v5 int = 1
        v6 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int16 = -72
        _ = fi7
        _ = v8
    }
    var v9 string = ""
    fmt.Println(v9, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 bool = (2 < 1)
    var v2 uint = 18446744073709551615
    var v3 uint64 = 4294967295
    var v10 float32 = (-2.25 - 0.0)
    v11 := GenS2{F0: 0, F1: 42}
    v10 = v10
    var v12 uint32 = 42
    var v13 int64 = 0
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{}
        if false {
            switch "a" {
            case "\n\t":
                switch fi14 {
                case 9223372036854775807:
                    v16 := GenS1{F0: 83, F1: (v13 << (((v2 & 1)) & 63))}
                    _ = v16
                }
            }
            var v17 bool = v1
            _ = v17
        } else {
            switch "\n\t" {
            case "hello":
                for fi18 := 0; fi18 < 18; fi18++ {
                    v19 := GenS2{F0: (v2 % ((438) | 1)), F1: v12}
                    v12 += 2
                    _ = fi18
                    _ = v19
                }
            default:
                v20 := GenS3{F0: 850, F1: (v2 - v2), F2: (v2 & v2)}
                _ = v20
            }
        }
        _ = fi14
        _ = v15
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v10 -= ((v10 / (((v10 * 3.14159) * v10) / -2.25)) + v10)
        v22 := []int{fi21, (fi21 / ((32768) | 1))}
        v23 := GenS2{F0: v2, F1: v12}
        _ = fi21
        _ = v22
        _ = v23
    }
    v13 -= v13
    var v24 uint64 = (^(v3))
    v25 := GenS2{F0: (v2 + v2), F1: (v12 / (((v12 - v12)) | 1))}
    switch 4294967295 {
    case 62:
        var v26 int32 = -272
        _ = v26
    case 9223372036854775807:
        v27 := []int{32767, 9223372036854775807, 205, 127}
        _ = v27
    default:
        switch 9223372036854775807 {
        case 32767:
            fmt.Println(v13)
        case -2147483648:
            if (-1 > -1) {
                v28 := GenS3{F0: 255, F1: 65535, F2: v2}
                var v29 int64 = v13
                v2 = v2
                _ = v28
                _ = v29
            }
        case 4294967295:
            switch "" {
            case "\n\t":
                for fi30 := 0; fi30 < 9; fi30++ {
                    v10 = (0.1 / float32(0.1))
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v3)
                    }
                    switch "hello" {
                    case "a":
                        var v31 float64 = float64(float32(-2.25))
                        _ = v31
                    }
                    _ = fi30
                }
            default:
                switch -81 {
                case -664:
                    if (42 == 255) {
                        var v32 int64 = (v13 % ((v13) | 1))
                        _ = v32
                    } else {
                        v13 = v13
                        v13 = (v13 * 42)
                    }
                default:
                    if v1 {
                        var v33 string = "hello"
                        _ = v33
                    }
                }
            }
        default:
            for fi34 := 0; fi34 < 20; fi34++ {
                if (v1 && (v1 && v1)) {
                    v1 = true
                } else {
                    if v1 {
                        if v1 {
                            for fi35 := 0; fi35 < 16; fi35++ {
                                v36 := []int{(fi34 ^ fi35)}
                                if ((0.0 + 0.1) > 0.0) {
                                    _ = len(v36)
                                    v3 -= (v3 & (^(v24)))
                                    if v1 {
                                        v36 = append(v36, fi34)
                                    }
                                    fmt.Println(v12)
                                    if (((((-2.25 / 1.5) != (-2.25 - float64(v10))) && v1) || true) && v1) {
                                        if (3.14159 < (0.0 / (3.14159 / 1.0))) {
                                            if (float64(((-2.25 / 1.5) / (((v10 * ((0.1 - 0.1) - 3.14159)) * v10) + (float32(3.14159) * (v10 / v10))))) == 1.5) {
                                                v37 := GenS1{F0: 127, F1: (v13 / (((v13 + v13)) | 1))}
                                                _ = len(v36)
                                                v38 := GenS1{F0: 69, F1: v13}
                                                _ = v37
                                                _ = v38
                                            }
                                            var v39 uint8 = uint8(v2)
                                            fmt.Println(v2)
                                            _ = v39
                                        } else {
                                            var v40 bool = v1
                                            if (false && (uint8(v2) != 255)) {
                                                var v41 uint16 = 1
                                                fmt.Println(fi35, fi34)
                                                v10 -= v10
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                                v42 := GenS3{F0: (v41 & (v41 | v41)), F1: uint(v41), F2: v2}
                                                var v43 int8 = -1
                                                v2 = v2
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v44 := GenS1{F0: 224, F1: (v13 | v13)}
                                                _ = len(v36)
                                                fmt.Println(v12)
                                                _ = v44
                                            }
                                            v36 = append(v36, (fi34 ^ fi34))
                                            var v45 float64 = -2.25
                                            _ = v40
                                            _ = v45
                                        }
                                        v1 = ((364 <= 642) && true)
                                    } else {
                                        switch fi35 {
                                        case 32767:
                                            _ = len(v36)
                                        case 42:
                                            var v46 int = (fi35 >> (((v2 ^ 1)) & 63))
                                            _ = v46
                                        }
                                        v47 := GenS4{F0: 18446744073709551615}
                                        var v48 uint = (^(v2))
                                        v49 := GenS5{F0: "", F1: fi34, F2: true}
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                    var v50 int64 = v13
                                    var v51 int8 = int8(v24)
                                    if len(v36) > 0 {
                                        v36[0] = fi34
                                    }
                                    v10 += v10
                                    _ = v50
                                    _ = v51
                                } else {
                                    var v52 int = fi35
                                    _ = v52
                                }
                                _ = fi35
                                _ = v36
                            }
                        } else {
                            fmt.Println(v3, v24)
                            var v53 int32 = 42
                            v54 := []int{}
                            _ = v53
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 19; fi55++ {
                            v10 += v10
                            if v1 {
                                v56 := []int{}
                                _ = v56
                            } else {
                                v12 += 65535
                                var v57 uint16 = 1
                                _ = v57
                            }
                            var v58 uint32 = v12
                            var v59 string = ("foo bar" + "\n\t")
                            var v60 float32 = (v10 / (v10 * v10))
                            _ = fi55
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        v61 := []int{(fi34 ^ (fi34 - int(v3))), fi34, (^(fi34)), fi34}
                        _ = v61
                    } else {
                        var v62 uint32 = v12
                        v63 := []int{-2147483648, 32767, (fi34 + (fi34 << ((18446744073709551615) & 63)))}
                        var v64 float32 = (v10 / (v10 - 3.14159))
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    var v65 int = fi34
                    switch "\n\t" {
                    case "":
                        v2 -= (v2 * (v2 & (v2 & (v2 ^ v2))))
                    }
                    _ = v65
                }
                var v66 bool = v1
                if (!(0 == 42)) {
                    v24 = v24
                    if v66 {
                        if v1 {
                            v67 := GenS3{F0: 255, F1: v2, F2: 1}
                            v2 = v2
                            _ = v67
                        }
                        v68 := GenS2{F0: v2, F1: 957}
                        switch (^(fi34)) {
                        case 931:
                            for fi69 := 0; fi69 < 17; fi69++ {
                                v3 = (v24 - (v24 % ((v24) | 1)))
                                var v70 int64 = -9223372036854775808
                                _ = fi69
                                _ = v70
                            }
                        case -976:
                            switch "hello" {
                            case "a":
                                var v71 int32 = 42
                                _ = v71
                            case "foo bar":
                                v13 = v13
                            case "\n\t":
                                fmt.Println(v13)
                            default:
                                v72 := GenS2{F0: (v2 | (v2 + (v2 | v2))), F1: v12}
                                _ = v72
                            }
                        case 178:
                            switch fi34 {
                            case 4294967295:
                                var v73 bool = (genFn0(-128, 305, 0) <= -1)
                                _ = v73
                            case 9223372036854775807:
                                fmt.Println(v66, v66, v13)
                            }
                        default:
                            for fi74 := 0; fi74 < 10; fi74++ {
                                var v75 uint = (v2 + v2)
                                var v76 int = (fi34 ^ fi34)
                                var v77 int = (-(v76))
                                _ = fi74
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                        }
                        _ = v68
                    } else {
                        var v78 float32 = float32(0.1)
                        _ = v78
                    }
                    var v79 string = "foo bar"
                    _ = v79
                } else {
                    if (v24 > v24) {
                        switch "foo bar" {
                        case "hello":
                            v1 = false
                        case "":
                            v10 -= (1.5 / v10)
                        case "\n\t":
                            v3 = v24
                        default:
                            if ((v12 | v12) != v12) {
                                v80 := []int{}
                                v1 = ((v12 * uint32(v24)) < (^(v12)))
                                for fi81 := 0; fi81 < 19; fi81++ {
                                    v82 := GenS6{F0: 0}
                                    if (!(107 <= genFn0(-1, v24, 0))) {
                                        switch fi34 {
                                        case 128:
                                            v2 -= 255
                                        case 127:
                                            var v83 uint8 = 59
                                            _ = v83
                                        }
                                    } else {
                                        v24 += 18446744073709551615
                                        v3 -= (^(v24))
                                    }
                                    var v84 int = (fi34 & (fi81 / (((^(fi34))) | 1)))
                                    _ = fi81
                                    _ = v82
                                    _ = v84
                                }
                                _ = v80
                            } else {
                                v24 += (v24 >> (((v2 ^ (^(v2)))) & 63))
                            }
                        }
                        v85 := []int{4294967295, -1, 4294967295, (fi34 % ((fi34) | 1))}
                        fmt.Println(v3, fi34, v24)
                        _ = v85
                    }
                }
                _ = fi34
                _ = v66
            }
        }
    }
    if v1 {
        v86 := GenS7{F0: 27, F1: (^(v3)), F2: 2147483647}
        _ = v86
    }
    if (100 >= 42) {
        v12 -= v12
        if (!false) {
            var v87 string = "hello"
            _ = v87
        } else {
            if true {
                v3 += 344
                fmt.Println(v12)
                for fi88 := 0; fi88 < 14; fi88++ {
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "a":
                        v89 := GenS8{F0: v12, F1: 128, F2: -128}
                        _ = v89
                    case "":
                        v3 += 985
                    default:
                        var v90 bool = v1
                        _ = v90
                    }
                    for fi91 := 0; fi91 < 6; fi91++ {
                        v10 = 1e10
                        _ = fi91
                    }
                    _ = fi88
                }
            }
        }
        v92 := GenS9{F0: 1, F1: -1, F2: genFn0(61, (v3 >> (((v2 % ((v2) | 1))) & 63)), 0)}
        _ = v92
    }
    v93 := GenS7{F0: int8(v13), F1: 1, F2: 2147483647}
    v94 := GenS7{F0: -111, F1: v3, F2: 42}
    var v95 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v93
    _ = v94
    _ = v95
}
