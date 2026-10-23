package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int8; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS11 struct { F0 uint8; F1 int; F2 uint }
type GenS12 struct { F0 int32; F1 int32 }
func genFn0() uint8 {
    v7 := []int{4294967295, 394, 32768}
    var v8 float64 = 1.5
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint8 = 42
        v8 += v8
        v11 := GenS1{F0: 127}
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        v8 -= ((v8 / v8) - v8)
        v12 := GenS2{F0: "", F1: 247}
        _ = v12
    }
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int32) int32 {
    fmt.Println(p0)
    var v13 uint32 = 2147483647
    v13 += (v13 / ((v13) | 1))
    fmt.Println(p0)
    var v14 uint = 65535
    _ = p0
    _ = v13
    _ = v14
    return 907
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 2147483648
    var v3 uint8 = 235
    v4 := make(map[int]string)
    var v5 uint = 65535
    var v6 int = 1
    v5 -= (v5 + v5)
    if true {
        fmt.Println(v6, v5, v6)
        fmt.Println(v3, v2)
        var v15 bool = true
        _ = v15
    } else {
        var v16 uint8 = 187
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := []int{4294967295, 127, int(v2)}
            var v19 int64 = 42
            _ = fi17
            _ = v18
            _ = v19
        }
        switch "hello" {
        case "a":
            v20 := GenS1{F0: -32}
            _ = v20
        default:
            if false {
                v1[""] = (v2 ^ (v2 / ((2) | 1)))
                if ((^(v2)) == v2) {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := GenS3{F0: (v2 + 116)}
                        var v23 bool = false
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                } else {
                    if (128 != (v16 * (v3 / ((0) | 1)))) {
                        switch (^(v6)) {
                        case 0:
                            var v24 int32 = int32(v16)
                            _ = v24
                        case 128:
                            _ = len(v4)
                        }
                        switch -7 {
                        case 496:
                            var v25 bool = (!true)
                            _ = v25
                        }
                    }
                }
                v26 := GenS3{F0: (v2 & 0)}
                _ = v26
            }
        }
        _ = v16
    }
    var v27 int16 = -1
    for fi28 := 0; fi28 < 3; fi28++ {
        v27 += (v27 & (v27 - (v27 * (v27 * (v27 ^ 42)))))
        v5 += (v5 << (((v5 + v5)) & 63))
        switch (^(v6)) {
        case 9223372036854775807:
            delete(v4, -110)
        case 32768:
            var v29 int16 = -1
            _ = v29
        default:
            var v30 int16 = (v27 % (((v27 / ((v27) | 1))) | 1))
            _ = v30
        }
        _ = fi28
    }
    v5 = (v5 ^ v5)
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint64 = 0
        if false {
            for fi33 := 0; fi33 < 11; fi33++ {
                v3 = 161
                var v34 float64 = 1.5
                if true {
                    if (!true) {
                        v35 := GenS4{F0: 123, F1: 87}
                        if (!(42 <= 0)) {
                            v36 := []int{}
                            _ = v36
                        } else {
                            if true {
                                v6 = 42
                            } else {
                                var v37 int32 = -2147483648
                                if false {
                                    v38 := GenS1{F0: 1}
                                    switch -9223372036854775808 {
                                    case 132:
                                        if (v27 == -436) {
                                            var v39 uint32 = (v2 % ((65535) | 1))
                                            fmt.Println(fi33, v3)
                                            if true {
                                                if (2147483647 < genFn1(genFn1(genFn1(-113)))) {
                                                    var v40 float32 = (float32(v34) * 3.14159)
                                                    _ = v40
                                                }
                                                v2 += (v39 & uint32(v32))
                                            } else {
                                                if false {
                                                    var v41 uint16 = 448
                                                    v41 -= v41
                                                    v42 := []int{fi33, 128}
                                                    _ = v41
                                                    _ = v42
                                                }
                                                var v43 uint64 = v32
                                                v44 := GenS5{F0: (v43 & v43), F1: "hello"}
                                                v45 := []int{fi33, fi31, fi33, v6}
                                                var v46 float64 = ((v34 / ((v34 * float64(1.5)) - float64(1e10))) - (float64((0.1 + (float32(v34) * 0.0))) * 0.0))
                                                v47 := GenS3{F0: 917}
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                            v48 := GenS5{F0: (^(v32)), F1: "hello"}
                                            _ = v39
                                            _ = v48
                                        }
                                    case 4294967295:
                                        if true {
                                            if ((!false) && (0.0 < v34)) {
                                                v6 = (-(v6))
                                                if false {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v32 -= v32
                                                    case "":
                                                        var v49 bool = (v3 < (v3 % ((127) | 1)))
                                                        _ = v49
                                                    case "hello":
                                                        v50 := GenS3{F0: 42}
                                                        _ = v50
                                                    }
                                                    var v51 int = fi31
                                                    v52 := []int{-2147483648, 42, 704, (v51 & int(v37))}
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    fmt.Println(v3, fi33)
                                                    v34 = v34
                                                    var v53 float32 = ((1.5 - -2.25) * 1.5)
                                                    var v54 uint = 876
                                                    v55 := GenS2{F0: "\n\t", F1: v2}
                                                    var v56 int64 = -1
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                }
                                                v6 = (fi33 * 350)
                                                var v57 float64 = v34
                                                _ = v57
                                            } else {
                                                if (genFn0() <= (v3 * uint8(v5))) {
                                                    if (!((v5 >> ((v5) & 63)) < (v5 ^ uint(v32)))) {
                                                        v2 += (v2 >> ((255) & 63))
                                                    } else {
                                                        v58 := []int{}
                                                        switch fi33 {
                                                        case 0:
                                                            var v59 int = (-(v6))
                                                            _ = v59
                                                        case 42:
                                                            v34 = 3.14159
                                                        case -2147483648:
                                                            var v60 uint8 = (v3 | (^(v3)))
                                                            _ = len(v58)
                                                            v5 = (v5 ^ v5)
                                                            _ = v60
                                                        }
                                                        _ = v58
                                                    }
                                                    v61 := GenS5{F0: 108, F1: "a"}
                                                    _ = v61
                                                } else {
                                                    v62 := GenS6{F0: v32, F1: 613}
                                                    _ = v62
                                                }
                                                if (false || (!true)) {
                                                    if false {
                                                        v63 := GenS1{F0: -128}
                                                        var v64 int64 = -9223372036854775808
                                                        fmt.Println(v6, v34)
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v6 += fi31
                                                        switch fi33 {
                                                        case 9223372036854775807:
                                                            var v65 int64 = 1
                                                            _ = v65
                                                        case 1:
                                                            v66 := GenS3{F0: v2}
                                                            _ = v66
                                                        case 96:
                                                            fmt.Println(v32, fi33)
                                                            v67 := GenS7{F0: true, F1: (v27 + (v27 / ((v27) | 1)))}
                                                            _ = v67
                                                        default:
                                                            var v68 bool = (1 >= -9223372036854775808)
                                                            _ = v68
                                                        }
                                                        v69 := v4[fi31]
                                                        var v70 bool = ((v2 % (((v2 % ((783) | 1))) | 1)) >= v2)
                                                        v71 := GenS8{F0: 127, F1: 2, F2: 9223372036854775807}
                                                        _ = v69
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                    if false {
                                                        var v72 int8 = -101
                                                        var v73 int = v6
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        switch (fi33 | fi33) {
                                                        case 4294967295:
                                                            v37 -= v37
                                                        case 953:
                                                            var v74 int32 = v37
                                                            if false {
                                                                if (uint16(v5) == 517) {
                                                                    if true {
                                                                        var v75 int32 = v37
                                                                        _ = v75
                                                                    }
                                                                    v76 := GenS8{F0: -1, F1: 255, F2: -9223372036854775808}
                                                                    v77 := GenS7{F0: true, F1: 32767}
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                            } else {
                                                                v4[(fi33 & (fi31 >> ((18446744073709551615) & 63)))] = "a"
                                                                v74 = int32(v32)
                                                                v78 := GenS9{F0: 65535, F1: -181}
                                                                var v79 float64 = 0.0
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                            _ = v74
                                                        }
                                                        v37 += v37
                                                        v34 = float64(((-2.25 - 0.1) - 0.0))
                                                    }
                                                } else {
                                                    v80 := GenS7{F0: (9223372036854775807 == 42), F1: v27}
                                                    if true {
                                                        var v81 float64 = (float64(((1e10 * float32((1e10 - ((v34 + (v34 / 3.14159)) - -2.25)))) + 3.14159)) / v34)
                                                        _ = v81
                                                    } else {
                                                        v82 := []int{fi31, 42, fi33}
                                                        if (-1 < -87) {
                                                            v83 := GenS10{F0: (v32 ^ (v32 & 0)), F1: (v2 + v2), F2: (v3 << ((65535) & 63))}
                                                            v84 := GenS1{F0: 42}
                                                            var v85 string = "a"
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            v86 := GenS11{F0: v3, F1: (fi31 % (((fi33 / (((v6 | (v6 / ((fi31) | 1)))) | 1))) | 1)), F2: v5}
                                                            _ = v86
                                                        }
                                                        if (float64(-2.25) >= (float64(1.5) + (((v34 * v34) / 1.0) / 0.1))) {
                                                            fmt.Println(v27)
                                                        } else {
                                                            fmt.Println(v3, v6, v5)
                                                            fmt.Println(v3, v32)
                                                        }
                                                        v87 := GenS12{F0: (v37 | v37), F1: (v37 >> ((v5) & 63))}
                                                        _ = v82
                                                        _ = v87
                                                    }
                                                    v32 += (^(v32))
                                                    _ = v80
                                                }
                                                v4[(v6 * v6)] = ("\n\t" + ("a" + "a"))
                                                v37 = (v37 + (v37 | (v37 % (((v37 * 331)) | 1))))
                                                var v88 int64 = -373
                                                _ = v88
                                            }
                                            v89 := []int{fi31, (-(v6)), (fi33 % ((2147483647) | 1)), v6}
                                            var v90 int8 = 1
                                            _ = v89
                                            _ = v90
                                        }
                                    case -1:
                                        var v91 uint8 = (v3 << (((v5 + 65535)) & 63))
                                        _ = v91
                                    }
                                    delete(v4, (fi33 / ((int(v5)) | 1)))
                                    _ = v38
                                } else {
                                    var v92 bool = false
                                    v93 := []int{(fi31 - (fi33 ^ -2147483648)), fi31}
                                    v34 += (v34 + 1.5)
                                    v6 += fi33
                                    if v92 {
                                        v5 = 255
                                    } else {
                                        var v94 uint32 = 633
                                        _ = v94
                                    }
                                    _ = v92
                                    _ = v93
                                }
                                if true {
                                    var v95 float64 = v34
                                    v95 = v95
                                    _ = v95
                                }
                                _ = v37
                            }
                            var v96 float64 = v34
                            v97 := GenS2{F0: ("" + ((("hello" + "") + "a") + "foo bar")), F1: v2}
                            _ = v96
                            _ = v97
                        }
                        _ = v35
                    } else {
                        if (v6 <= fi33) {
                            var v98 int8 = 42
                            var v99 uint8 = (v3 + (v3 * (v3 - 1)))
                            _ = v98
                            _ = v99
                        }
                        fmt.Println(v6, v32)
                    }
                } else {
                    fmt.Println(v27)
                }
                _ = fi33
                _ = v34
            }
        }
        _ = fi31
        _ = v32
    }
    var v100 int64 = 42
    v2 = (^(v2))
    var v101 float32 = (1e10 * 1.5)
    fmt.Println(v2, v100, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
