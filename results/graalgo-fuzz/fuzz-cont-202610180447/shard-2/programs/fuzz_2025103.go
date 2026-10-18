package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8; F1 bool; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0() int8 {
    var v4 float32 = 0.1
    var v5 bool = true
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1000
    var v3 string = "a"
    v3 = (((("a" + v3) + (("a" + (v3 + "a")) + v3)) + "a") + v3)
    delete(v1, 128)
    fmt.Println(v2, v2)
    if (667 == 651) {
        var v6 int8 = -128
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            if (!(("\n\t" + "hello") >= v3)) {
                fmt.Println(fi8, v3)
            }
            var v9 uint16 = 32767
            v2 -= (v2 & v2)
            _ = fi8
            _ = v9
        }
        v3 = "hello"
        v10 := GenS1{F0: 17, F1: (v3 + "")}
        _ = fi7
        _ = v10
    }
    v11 := GenS2{F0: 0, F1: ((0 == 4294967295) || (true && (3.14159 < 0.1))), F2: -9223372036854775808}
    var v12 int64 = 1
    fmt.Println(v12)
    v13 := []int{}
    for fi14 := 0; fi14 < 19; fi14++ {
        v2 += (^(v2))
        _ = fi14
    }
    var v15 uint32 = 4294967295
    switch (((v3 + v3) + v3) + ((v3 + ("foo bar" + (v3 + v3))) + v3)) {
    case "a":
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint64 = 2
            for fi18 := 0; fi18 < 18; fi18++ {
                v3 = ((v3 + (("" + v3) + "\n\t")) + (v3 + ("" + (v3 + (v3 + v3)))))
                if (184 < 1) {
                    var v19 uint32 = (v15 >> ((901) & 63))
                    var v20 uint64 = (v17 ^ v17)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = 1
                var v22 uint = 2
                if ((0.1 + -2.25) >= 1.5) {
                    var v23 uint32 = 406
                    switch v3 {
                    case "a":
                        v24 := v1[fi18]
                        _ = v24
                    case "":
                        _ = len(v13)
                    }
                    if (2 == (^(v21))) {
                        switch -529 {
                        case 42:
                            if ((v12 << ((442) & 63)) < v12) {
                                fmt.Println(v12, v2, v22)
                                if false {
                                    v25 := v1[(fi18 * fi18)]
                                    v26 := GenS1{F0: 127, F1: "\n\t"}
                                    _ = v25
                                    _ = v26
                                }
                                if (!(false || true)) {
                                    var v27 int64 = (v12 + v12)
                                    v3 = "\n\t"
                                    if false {
                                        v3 = (((v3 + "a") + (v3 + v3)) + (v3 + v3))
                                        if len(v13) > 0 {
                                            v13[0] = fi16
                                        }
                                        if false {
                                            v28 := GenS1{F0: 127, F1: v3}
                                            if (!false) {
                                                fmt.Println(fi16, v21)
                                            }
                                            _ = v28
                                        } else {
                                            delete(v1, (fi18 - (fi16 ^ (fi18 & fi16))))
                                        }
                                        v13 = append(v13, (fi16 + (fi16 >> (((v22 << (((v22 % (((v22 & 862)) | 1))) & 63))) & 63))))
                                    } else {
                                        var v29 int64 = -1
                                        var v30 int64 = v29
                                        fmt.Println(fi16, v2, fi16)
                                        switch v3 {
                                        case "hello":
                                            switch (fi18 % ((9223372036854775807) | 1)) {
                                            case -859:
                                                _ = len(v13)
                                            default:
                                                var v31 float32 = (3.14159 * (float32(1e10) - (1e10 * 1.5)))
                                                if false {
                                                    v15 -= v23
                                                    _ = len(v13)
                                                }
                                                _ = v31
                                            }
                                            if (!true) {
                                                fmt.Println(v23)
                                                if true {
                                                    v15 = 584
                                                } else {
                                                    v27 = (-(v12))
                                                }
                                                v32 := v1[128]
                                                _ = v32
                                            }
                                            var v33 int32 = 42
                                            _ = v33
                                        }
                                        if false {
                                            if (false && false) {
                                                if len(v13) > 0 {
                                                    v13[0] = 32767
                                                }
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                                if false {
                                                    v3 = "a"
                                                    fmt.Println(v15)
                                                } else {
                                                    switch 42 {
                                                    case 42:
                                                        if false {
                                                            v30 -= 0
                                                        } else {
                                                            var v34 float32 = (0.0 + 3.14159)
                                                            if ("hello" >= v3) {
                                                                if len(v13) > 0 {
                                                                    v13[0] = (fi18 >> ((v22) & 63))
                                                                }
                                                                if true {
                                                                    v12 += (-(v27))
                                                                    fmt.Println(v15, v22, v3)
                                                                    switch 2147483647 {
                                                                    case 9223372036854775807:
                                                                        v35 := GenS3{F0: -128}
                                                                        _ = v35
                                                                    default:
                                                                        if len(v13) > 0 {
                                                                            v13[0] = fi18
                                                                        }
                                                                    }
                                                                }
                                                                v34 -= v34
                                                            }
                                                            _ = v34
                                                        }
                                                    }
                                                    if ((uint(v17) == v22) || false) {
                                                        if false {
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                            if (!(-40 <= 42)) {
                                                                var v36 float64 = 0.1
                                                                _ = v36
                                                            } else {
                                                                if ((286 == 0) || true) {
                                                                    v21 -= v17
                                                                    if len(v13) > 0 {
                                                                        v13[0] = fi16
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v15, v22)
                                                                        var v37 int32 = -210
                                                                        _ = v37
                                                                    }
                                                                }
                                                            }
                                                            v21 -= v17
                                                        }
                                                    }
                                                    switch fi18 {
                                                    case 128:
                                                        v2 = v2
                                                    case 1:
                                                        var v38 string = ""
                                                        _ = v38
                                                    case 9223372036854775807:
                                                        var v39 float64 = 1.5
                                                        _ = v39
                                                    default:
                                                        v30 -= 42
                                                    }
                                                }
                                            }
                                            var v40 uint64 = uint64(v22)
                                            v30 = v29
                                            _ = v40
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                    v41 := GenS2{F0: 127, F1: (1 > 255), F2: 0}
                                    _ = v27
                                    _ = v41
                                }
                            } else {
                                var v42 uint8 = 255
                                v22 = v22
                                _ = v42
                            }
                        case 362:
                            var v43 int64 = 1
                            _ = v43
                        default:
                            v21 += (v17 & 18446744073709551615)
                            v44 := GenS3{F0: 100}
                            v15 = (v15 << ((v22) & 63))
                            switch (fi16 - (fi18 + fi18)) {
                            case 1:
                                fmt.Println(v22)
                            default:
                                if (42 != v21) {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v45 := GenS3{F0: 42}
                                    var v46 string = v3
                                    v17 = 4294967295
                                    v47 := GenS3{F0: 42}
                                    var v48 int = -1
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                } else {
                                    v15 = 65535
                                }
                            }
                            if len(v13) > 0 {
                                v13[0] = (fi16 - -9223372036854775808)
                            }
                            switch (v3 + ("" + v3)) {
                            case "hello":
                                if false {
                                    var v49 uint8 = 42
                                    fmt.Println(v22, v21)
                                    v50 := GenS4{F0: v2, F1: 100}
                                    _ = v49
                                    _ = v50
                                } else {
                                    delete(v1, fi18)
                                }
                            case "\n\t":
                                var v51 int8 = 111
                                v22 -= (v22 >> ((uint(v21)) & 63))
                                var v52 int16 = 42
                                _ = v51
                                _ = v52
                            }
                            v3 = (v3 + (v3 + v3))
                            _ = v44
                        }
                        v3 = "\n\t"
                    } else {
                        if (true && false) {
                            v17 -= v21
                        } else {
                            var v53 int = -2147483648
                            _ = v53
                        }
                        var v54 string = v3
                        v55 := GenS3{F0: 127}
                        _ = v54
                        _ = v55
                    }
                    _ = v23
                } else {
                    if true {
                        if true {
                            v15 = v15
                            var v56 uint32 = (v15 - (v15 * v15))
                            if true {
                                fmt.Println(v12)
                            } else {
                                if false {
                                    v57 := GenS1{F0: uint8(v21), F1: v3}
                                    v12 = (v12 + v12)
                                    fmt.Println(v17, fi16)
                                    fmt.Println(v2, v22, v22)
                                    v58 := GenS4{F0: 1, F1: 127}
                                    _ = v57
                                    _ = v58
                                }
                                if ("hello" < "") {
                                    fmt.Println(v17, fi18)
                                }
                                if true {
                                    if true {
                                        v56 = v56
                                        var v59 uint32 = 42
                                        _ = v59
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v60 uint = uint(v15)
                                        _ = v60
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "foo bar":
                                            v12 += (v12 + v12)
                                        case "":
                                            v17 = (v17 << ((v22) & 63))
                                        default:
                                            v61 := GenS1{F0: 114, F1: v3}
                                            _ = v61
                                        }
                                    default:
                                        v56 += 471
                                    }
                                    _ = len(v13)
                                    v12 += v12
                                } else {
                                    switch v3 {
                                    case "\n\t":
                                        var v62 uint16 = 571
                                        _ = v62
                                    case "":
                                        var v63 uint16 = 1
                                        _ = v63
                                    case "a":
                                        v64 := GenS5{F0: (v12 | (v12 ^ int64(v21)))}
                                        _ = v64
                                    default:
                                        if false {
                                            var v65 float64 = 0.0
                                            v12 = v12
                                            _ = v65
                                        } else {
                                            var v66 bool = false
                                            _ = v66
                                        }
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        v17 += v17
                                    case "foo bar":
                                        v67 := v1[fi16]
                                        fmt.Println(v15)
                                        _ = v67
                                    }
                                    var v68 int16 = (v2 + v2)
                                    v69 := v1[(fi18 * fi18)]
                                    switch fi16 {
                                    case -434:
                                        switch (^(fi18)) {
                                        case 32768:
                                            var v70 uint = v22
                                            _ = v70
                                        default:
                                            v69 = ("a" + v69)
                                        }
                                    case 2147483647:
                                        v15 = (v15 | (v56 << ((0) & 63)))
                                    case -497:
                                        var v71 int16 = v68
                                        _ = v71
                                    default:
                                        v13 = append(v13, fi18)
                                    }
                                    v72 := GenS1{F0: 223, F1: v3}
                                    _ = v68
                                    _ = v69
                                    _ = v72
                                }
                                var v73 bool = (!true)
                                _ = v73
                            }
                            _ = v56
                        } else {
                            switch "foo bar" {
                            case "a":
                                v74 := GenS4{F0: int16(v15), F1: 1}
                                _ = v74
                            case "\n\t":
                                switch (fi18 | (fi16 * -631)) {
                                case -1:
                                    fmt.Println(v17, v15)
                                case 42:
                                    if ((v22 >> (((v22 % ((v22) | 1))) & 63)) <= 255) {
                                        v21 -= (v21 >> (((^(v22))) & 63))
                                    } else {
                                        if true {
                                            v75 := GenS1{F0: 60, F1: v3}
                                            v76 := GenS1{F0: 0, F1: v3}
                                            _ = v75
                                            _ = v76
                                        } else {
                                            if true {
                                                v15 += v15
                                            } else {
                                                v22 = 0
                                                switch fi16 {
                                                case 506:
                                                    if false {
                                                        v77 := GenS2{F0: -34, F1: ((v2 >> (((v22 / ((v22) | 1))) & 63)) < (v2 >> ((v22) & 63))), F2: (v12 << ((v22) & 63))}
                                                        _ = v77
                                                    } else {
                                                        _ = len(v13)
                                                        v12 = (v12 - (-(v12)))
                                                    }
                                                case 2147483647:
                                                    if true {
                                                        v2 -= v2
                                                        var v78 float32 = 3.14159
                                                        _ = v78
                                                    } else {
                                                        fmt.Println(fi16, fi18)
                                                        fmt.Println(v12)
                                                    }
                                                case 32767:
                                                    v79 := GenS6{F0: -2147483648, F1: 2}
                                                    _ = v79
                                                default:
                                                    var v80 float64 = 3.14159
                                                    _ = v80
                                                }
                                                if (0.0 <= float64(-2.25)) {
                                                    v12 = v12
                                                    var v81 uint = (v22 ^ v22)
                                                    _ = v81
                                                } else {
                                                    v3 = "a"
                                                    v22 = (v22 >> ((v22) & 63))
                                                    if (!false) {
                                                        var v82 int8 = genFn0()
                                                        var v83 string = v3
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                    fmt.Println(fi18)
                                                }
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                var v84 int = (fi16 | 9223372036854775807)
                                _ = v84
                            }
                            v2 = v2
                            v85 := GenS6{F0: -2147483648, F1: 255}
                            v2 += int16(fi18)
                            _ = v85
                        }
                    }
                }
                _ = fi18
                _ = v21
                _ = v22
            }
            fmt.Println(v2)
            _ = fi16
            _ = v17
        }
    }
    var v86 int8 = genFn0()
    switch v3 {
    case "foo bar":
        var v87 uint16 = 280
        _ = v87
    case "a":
        if (true || false) {
            if true {
                v88 := GenS6{F0: -2147483648, F1: 255}
                var v89 uint16 = 255
                fmt.Println(v15, v12)
                _ = v88
                _ = v89
            } else {
                var v90 int64 = (v12 << ((587) & 63))
                _ = v90
            }
            for fi91 := 0; fi91 < 7; fi91++ {
                if len(v13) > 0 {
                    v13[0] = fi91
                }
                var v92 float32 = -2.25
                if true {
                    _ = len(v13)
                    v12 -= v12
                }
                _ = fi91
                _ = v92
            }
        } else {
            v15 += (^(v15))
            var v93 uint8 = 6
            _ = v93
        }
    default:
        v15 -= (v15 % ((v15) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v86
}
