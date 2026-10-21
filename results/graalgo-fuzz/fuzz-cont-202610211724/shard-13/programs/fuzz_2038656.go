package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
func genFn0() uint8 {
    var v5 int = -1
    v5 -= v5
    v5 -= v5
    var v6 int8 = 46
    var v7 int8 = v6
    _ = v5
    _ = v6
    _ = v7
    return 202
}
func genFn1(p0 uint32, p1 int32, p2 int16) uint {
    v8 := []int{9223372036854775807, -267, -729, 14}
    if len(v8) > 0 {
        v8[0] = -634
    }
    p0 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    v1 := []int{}
    var v2 int64 = -995
    var v3 int = 755
    v4 := make(map[uint32]uint32)
    if false {
        var v9 int8 = -71
        if false {
            _ = len(v4)
            v10 := GenS1{F0: 2147483647, F1: 65535, F2: (v3 * v3)}
            _ = v10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            if (true || false) {
                v9 = v9
            }
            v2 += (v2 * (v2 | (v2 + v2)))
            switch "a" {
            case "":
                for fi12 := 0; fi12 < 17; fi12++ {
                    if (42 != 475) {
                        var v13 uint8 = 52
                        _ = v13
                    } else {
                        if true {
                            v4[822] = 4294967295
                        }
                        fmt.Println(fi12)
                        v2 = (v2 / ((-1) | 1))
                    }
                    _ = fi12
                }
            case "hello":
                v4[2] = 42
            default:
                var v14 uint32 = 2
                _ = v14
            }
            _ = fi11
        }
        _ = v9
    } else {
        var v15 int64 = v2
        if true {
            switch ("a" + "") {
            case "a":
                v15 = -893
            case "":
                if false {
                    fmt.Println(v3, v2)
                    var v16 uint32 = 360
                    for fi17 := 0; fi17 < 14; fi17++ {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 int16 = 0
                            _ = fi18
                            _ = v19
                        }
                        _ = fi17
                    }
                    _ = v16
                }
            case "foo bar":
                v20 := GenS2{F0: 470, F1: 255, F2: 1}
                _ = v20
            }
            if (!(v15 != -1)) {
                if (false && false) {
                    var v21 bool = (false || false)
                    if ((v21 || (65535 <= 42)) && (true || false)) {
                        v15 = (-(v15))
                        if v21 {
                            v21 = true
                            var v22 float32 = ((1e10 - ((-2.25 * -2.25) * 0.1)) + (1.5 - 3.14159))
                            for fi23 := 0; fi23 < 12; fi23++ {
                                if false {
                                    v22 += float32((1e10 * 0.1))
                                    v2 = (v15 | v15)
                                    v2 += v2
                                } else {
                                    switch fi23 {
                                    case 208:
                                        switch "\n\t" {
                                        case "a":
                                            var v24 int64 = -1
                                            _ = v24
                                        case "hello":
                                            var v25 int = v3
                                            _ = v25
                                        case "\n\t":
                                            for fi26 := 0; fi26 < 7; fi26++ {
                                                v15 -= v15
                                                var v27 uint8 = 59
                                                _ = fi26
                                                _ = v27
                                            }
                                        default:
                                            switch "hello" {
                                            case "\n\t":
                                                for fi28 := 0; fi28 < 14; fi28++ {
                                                    v29 := GenS2{F0: 42, F1: genFn0(), F2: 0}
                                                    v30 := GenS2{F0: 255, F1: 1, F2: 920}
                                                    _ = len(v4)
                                                    _ = fi28
                                                    _ = v29
                                                    _ = v30
                                                }
                                            case "a":
                                                if (true || false) {
                                                    for fi31 := 0; fi31 < 3; fi31++ {
                                                        v32 := GenS2{F0: 240, F1: 128, F2: 4294967295}
                                                        _ = fi31
                                                        _ = v32
                                                    }
                                                    v3 = 127
                                                }
                                            }
                                        }
                                    case 128:
                                        var v33 int32 = -574
                                        _ = v33
                                    case 2147483647:
                                        _ = len(v4)
                                    default:
                                        var v34 int32 = 620
                                        _ = v34
                                    }
                                    v3 -= v3
                                }
                                if v21 {
                                    var v35 int = (fi23 + (v3 | (fi23 | fi23)))
                                    var v36 uint64 = 42
                                    if (-5 <= 1) {
                                        v37 := GenS2{F0: uint16(v36), F1: 1, F2: v36}
                                        var v38 float32 = v22
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if (!((!((1e10 + (v22 * float32((float64(v22) - 0.1)))) == -2.25)) || (v21 && v21))) {
                                            v36 -= v36
                                            var v39 float64 = 0.0
                                            if (2 > genFn1(42, -843, -32768)) {
                                                if ((v22 / ((-2.25 / v22) / v22)) > v22) {
                                                    var v40 uint64 = v36
                                                    var v41 float64 = ((float64(((v22 / v22) * float32(1e10))) * 1.5) * v39)
                                                    var v42 uint32 = 0
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                }
                                                v43 := GenS2{F0: 0, F1: 151, F2: v36}
                                                switch (v35 / ((fi23) | 1)) {
                                                case 219:
                                                    if (v21 || (true || (!(-32768 == 233)))) {
                                                        switch (v35 * (-(v35))) {
                                                        case 0:
                                                            if false {
                                                                var v44 int64 = 1
                                                                var v45 uint8 = 42
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        case 32768:
                                                            for fi46 := 0; fi46 < 14; fi46++ {
                                                                v36 += (v36 & (v36 >> ((2) & 63)))
                                                                if (!(127 < 52)) {
                                                                    var v47 uint64 = (^(v36))
                                                                    if (!(4294967295 > 2)) {
                                                                        var v48 int8 = 73
                                                                        var v49 int64 = 0
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                    v4[4294967295] = 248
                                                                    _ = v47
                                                                } else {
                                                                    var v50 uint16 = 255
                                                                    var v51 int64 = v15
                                                                    if (0 >= 1) {
                                                                        fmt.Println(v21)
                                                                        var v52 int32 = 2147483647
                                                                        _ = v52
                                                                    }
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                                var v53 float32 = v22
                                                                _ = fi46
                                                                _ = v53
                                                            }
                                                        default:
                                                            var v54 int32 = -961
                                                            _ = v54
                                                        }
                                                        v22 = v22
                                                        switch ("\n\t" + "\n\t") {
                                                        case "a":
                                                            fmt.Println(v39)
                                                        }
                                                    }
                                                }
                                                _ = v43
                                            }
                                            _ = v39
                                        } else {
                                            v35 = fi23
                                            for fi55 := 0; fi55 < 4; fi55++ {
                                                if (304 >= 2) {
                                                    switch "" {
                                                    case "foo bar":
                                                        var v56 int16 = 992
                                                        _ = v56
                                                    case "a":
                                                        v57 := v4[65535]
                                                        _ = v57
                                                    default:
                                                        v4[68] = 0
                                                    }
                                                    v15 = v15
                                                    _ = len(v4)
                                                } else {
                                                    v22 = v22
                                                    v58 := v4[65535]
                                                    var v59 int16 = 1000
                                                    _ = v58
                                                    _ = v59
                                                }
                                                switch (fi55 * fi55) {
                                                case -409:
                                                    fmt.Println(v2)
                                                case 531:
                                                    var v60 int16 = -810
                                                    _ = v60
                                                }
                                                _ = fi55
                                            }
                                            for fi61 := 0; fi61 < 3; fi61++ {
                                                var v62 int8 = 42
                                                _ = fi61
                                                _ = v62
                                            }
                                        }
                                        for fi63 := 0; fi63 < 7; fi63++ {
                                            if v21 {
                                                var v64 uint = 333
                                                _ = v64
                                            }
                                            _ = fi63
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                } else {
                                    fmt.Println(v3, v3)
                                    var v65 int8 = 68
                                    _ = v65
                                }
                                _ = fi23
                            }
                            _ = v22
                        } else {
                            for fi66 := 0; fi66 < 6; fi66++ {
                                v3 = (v3 & (v3 + v3))
                                switch ((("a" + "") + "") + "a") {
                                case "\n\t":
                                    _ = len(v1)
                                default:
                                    v67 := GenS3{F0: 1, F1: ((^(fi66)) < fi66), F2: 1}
                                    _ = v67
                                }
                                _ = fi66
                            }
                            var v68 int32 = -1
                            _ = v68
                        }
                        v21 = (619 < 42)
                    } else {
                        v69 := GenS3{F0: -1, F1: (!v21), F2: 1}
                        _ = len(v4)
                        v70 := GenS3{F0: 42, F1: v21, F2: 65535}
                        _ = v69
                        _ = v70
                    }
                    _ = v21
                } else {
                    var v71 uint32 = 66
                    _ = v71
                }
                _ = len(v4)
            } else {
                var v72 string = ""
                if (355 == 65535) {
                    var v73 bool = false
                    v15 -= v15
                    _ = v73
                }
                if len(v1) > 0 {
                    v1[0] = v3
                }
                _ = v72
            }
            v3 -= v3
        } else {
            var v74 uint8 = 120
            for fi75 := 0; fi75 < 9; fi75++ {
                for fi76 := 0; fi76 < 3; fi76++ {
                    v15 = 132
                    _ = fi76
                }
                _ = fi75
            }
            _ = v74
        }
        _ = v15
    }
    var v77 float32 = (0.1 / (3.14159 + 1.5))
    var v78 float32 = ((v77 - v77) + float32(0.0))
    switch ("" + "a") {
    case "a":
        var v79 uint16 = 914
        _ = v79
    default:
        v1 = append(v1, v3)
    }
    v78 -= (v77 * 1e10)
    v78 = (v77 * v77)
    var v80 uint8 = 178
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v78
    _ = v80
}
