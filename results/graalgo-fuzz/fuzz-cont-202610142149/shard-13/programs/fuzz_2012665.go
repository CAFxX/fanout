package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 bool }
type GenS4 struct { F0 uint32; F1 int8; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 bool }

func main() {
    v1 := []string{"hello", "\n\t"}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 1
    v4 := []int{32767, -351, 0}
    if false {
        v5 := GenS1{F0: false, F1: 32767, F2: 568}
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint8 = 30
        v3 -= v3
        _ = fi6
        _ = v7
    }
    fmt.Println(v3)
    v3 -= v3
    var v8 uint64 = 1
    switch ("a" + ("foo bar" + ((("foo bar" + "foo bar") + "foo bar") + "a"))) {
    case "foo bar":
        v8 += v8
    }
    fmt.Println(v3)
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS2{F0: 322, F1: -2147483648}
            v12 := v2[292]
            switch ("a" + "") {
            case "hello":
                var v13 int32 = -1
                _ = v13
            default:
                switch (fi10 & (fi9 ^ (fi9 % ((fi10) | 1)))) {
                case 128:
                    if ((v12 / ((uint32(v8)) | 1)) < (v12 % ((v12) | 1))) {
                        switch 32767 {
                        case 9223372036854775807:
                            var v14 int16 = -521
                            _ = v14
                        }
                        if false {
                            switch fi10 {
                            case -675:
                                v8 = 18446744073709551615
                            case 32767:
                                v3 = v3
                            case 128:
                                fmt.Println(fi10, fi10)
                            default:
                                v12 -= 42
                            }
                            var v15 int64 = v3
                            var v16 uint = 76
                            _ = v15
                            _ = v16
                        } else {
                            _ = len(v2)
                        }
                        v17 := GenS1{F0: (11 >= 1), F1: 682, F2: 1}
                        v18 := GenS3{F0: -1, F1: false, F2: ((v3 % (((v3 | (-(v3)))) | 1)) == 9223372036854775807)}
                        _ = v17
                        _ = v18
                    } else {
                        if (("a" + ("foo bar" + "\n\t")) >= "foo bar") {
                            v19 := v2[(^(v12))]
                            v20 := GenS2{F0: -32768, F1: 771}
                            _ = v19
                            _ = v20
                        } else {
                            if (!false) {
                                v21 := GenS4{F0: (v12 + 601), F1: 42, F2: 42}
                                fmt.Println(v8, fi10)
                                v3 += v3
                                _ = v21
                            }
                            v22 := GenS4{F0: (v12 - v12), F1: -1, F2: 65535}
                            if (true || (42 == v8)) {
                                switch ("foo bar" + "a") {
                                case "":
                                    v2[v12] = v12
                                case "a":
                                    if ((v12 & v12) == 848) {
                                        var v23 int = fi9
                                        v8 -= 2
                                        switch fi9 {
                                        case -471:
                                            v24 := GenS5{F0: (^(v12))}
                                            _ = v24
                                        case 32767:
                                            switch "a" {
                                            case "\n\t":
                                                if (v23 != fi9) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v4 = append(v4, (v23 >> ((714) & 63)))
                                                    v25 := v2[v12]
                                                    v4 = append(v4, (^(fi10)))
                                                    switch fi9 {
                                                    case 42:
                                                        if (v8 == 1) {
                                                            v8 = v8
                                                            v1 = append(v1, "a")
                                                        } else {
                                                            if false {
                                                                v23 = (v23 / ((359) | 1))
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    case 485:
                                                        if len(v1) > 0 {
                                                            v1[0] = ""
                                                        }
                                                    }
                                                    _ = v25
                                                }
                                                if ((42 > 100) && true) {
                                                    if true {
                                                        var v26 int32 = 95
                                                        v27 := GenS4{F0: 4294967295, F1: 1, F2: 1}
                                                        v28 := v2[(v12 - 0)]
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    } else {
                                                        var v29 uint16 = 2
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if (((v29 | (v29 ^ 0)) != uint16(v12)) || ((-939 == fi10) && false)) {
                                                            _ = len(v2)
                                                            v4 = append(v4, (^(v23)))
                                                            if len(v1) > 0 {
                                                                v1[0] = "a"
                                                            }
                                                            if (((int8(v12) >= -1) || true) && (true || ((("\n\t" + (("" + "") + (("hello" + "") + ("" + "hello")))) == "foo bar") && (583 < 757)))) {
                                                                if false {
                                                                    fmt.Println(v12, v3)
                                                                    var v30 string = "a"
                                                                    v23 -= 656
                                                                    _ = v30
                                                                }
                                                                var v31 int32 = 0
                                                                _ = v31
                                                            } else {
                                                                fmt.Println(fi9)
                                                            }
                                                        } else {
                                                            delete(v2, 65535)
                                                            var v32 uint16 = v29
                                                            _ = v32
                                                        }
                                                        _ = v29
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v33 string = ("" + ("hello" + ("foo bar" + "")))
                                                        v34 := GenS6{F0: 1, F1: 32767}
                                                        var v35 uint16 = 968
                                                        v12 = uint32(v35)
                                                        var v36 uint = 255
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v37 := GenS7{F0: 127, F1: 50}
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    var v38 uint16 = 255
                                                    _ = v38
                                                }
                                                v8 = (v8 - (^(v8)))
                                            case "a":
                                                fmt.Println(v12, v3, fi10)
                                            case "hello":
                                                v39 := GenS4{F0: v12, F1: -104, F2: 193}
                                                _ = v39
                                            default:
                                                switch -254 {
                                                case -1:
                                                    _ = len(v2)
                                                case -154:
                                                    v40 := GenS3{F0: -327, F1: false, F2: ((true || (!false)) && true)}
                                                    _ = v40
                                                case -4:
                                                    if false {
                                                        fmt.Println(v23)
                                                    } else {
                                                        v41 := GenS7{F0: -111, F1: 153}
                                                        v8 -= (^(v8))
                                                        if (722 >= -647) {
                                                            fmt.Println(fi10, v12, v12)
                                                            switch ("foo bar" + ("" + "a")) {
                                                            case "\n\t":
                                                                v42 := GenS3{F0: -248, F1: ((!(142 <= 1)) || (false || false)), F2: (2147483647 <= 445)}
                                                                _ = v42
                                                            case "a":
                                                                v3 += v3
                                                            case "":
                                                                if (72 != -531) {
                                                                    var v43 float32 = 0.0
                                                                    _ = v43
                                                                }
                                                            default:
                                                                if (!false) {
                                                                    v8 -= (^(v8))
                                                                    v44 := GenS6{F0: 127, F1: -945}
                                                                    v23 = (^(fi10))
                                                                    _ = v44
                                                                }
                                                            }
                                                            v23 -= 4294967295
                                                        }
                                                        _ = v41
                                                    }
                                                default:
                                                    v12 -= 784
                                                }
                                            }
                                        case 4294967295:
                                            fmt.Println(v23)
                                        default:
                                            v45 := GenS5{F0: 2147483647}
                                            _ = v45
                                        }
                                        fmt.Println(fi9)
                                        var v46 int16 = 1000
                                        _ = v23
                                        _ = v46
                                    }
                                }
                                v8 += uint64(v12)
                                v3 += (v3 % (((v3 % ((v3) | 1))) | 1))
                            } else {
                                if ((!true) || ((("" + "foo bar") + "foo bar") < (("hello" + "a") + (("a" + ("hello" + ("hello" + "a"))) + "hello")))) {
                                    if true {
                                        fmt.Println(fi10, v3)
                                    }
                                } else {
                                    fmt.Println(v8, fi10)
                                }
                                var v47 int16 = 1000
                                var v48 uint64 = (v8 ^ (v8 / ((v8) | 1)))
                                if (127 > -128) {
                                    if (fi10 >= fi10) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        var v49 uint = 2
                                        _ = v49
                                    }
                                } else {
                                    var v50 int32 = -1
                                    v51 := GenS5{F0: (^(v12))}
                                    v2[v12] = (v12 ^ v12)
                                    _ = v50
                                    _ = v51
                                }
                                if ((fi10 - (fi9 & fi10)) <= 32768) {
                                    v3 = -9223372036854775808
                                }
                                if (!(true || true)) {
                                    if true {
                                        if true {
                                            if false {
                                                if true {
                                                    delete(v2, (^(v12)))
                                                } else {
                                                    if true {
                                                        var v52 int = 0
                                                        var v53 int8 = 100
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        var v54 int = fi10
                                                        _ = v54
                                                    }
                                                    v55 := v2[(v12 - (^(v12)))]
                                                    fmt.Println(fi9, v12, v55)
                                                    _ = v55
                                                }
                                                if (!false) {
                                                    if (((!(1e10 > 0.1)) || (0 > 0)) && true) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v56 int32 = 2147483647
                                                        switch fi10 {
                                                        case 4294967295:
                                                            v47 = (v47 ^ 1000)
                                                        case 0:
                                                            v57 := GenS1{F0: (((v47 >> ((255) & 63)) < v47) && true), F1: int16(v8), F2: v56}
                                                            _ = v57
                                                        }
                                                        _ = v56
                                                    }
                                                    v58 := v2[(v12 * (v12 | (^(v12))))]
                                                    _ = len(v2)
                                                    _ = v58
                                                } else {
                                                    v47 += (v47 ^ v47)
                                                    if (!(false && (0.1 >= 1e10))) {
                                                        v8 += (v8 ^ v48)
                                                        if false {
                                                            var v59 int = fi10
                                                            v48 = v8
                                                            _ = v59
                                                        }
                                                    } else {
                                                        v60 := GenS6{F0: 100, F1: (-(v47))}
                                                        _ = v60
                                                    }
                                                }
                                                fmt.Println(fi10)
                                            }
                                            if false {
                                                var v61 uint16 = 32767
                                                _ = v61
                                            } else {
                                                v2[(^(v12))] = 2147483648
                                                fmt.Println(v3, v47)
                                                if (!false) {
                                                    v47 = (v47 ^ 1)
                                                    if (v8 > v8) {
                                                        if ((v8 * (v8 + (v8 / (((v48 - v8)) | 1)))) <= (^(v48))) {
                                                            var v62 float64 = 1.5
                                                            var v63 int64 = int64(v8)
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            fmt.Println(v48)
                                                            v64 := GenS8{F0: 2147483647}
                                                            _ = v64
                                                        }
                                                        var v65 int64 = -678
                                                        _ = v65
                                                    }
                                                    v48 -= (^(v48))
                                                }
                                                var v66 uint16 = 2
                                                _ = v66
                                            }
                                            var v67 int = fi10
                                            switch "foo bar" {
                                            case "\n\t":
                                                v68 := GenS1{F0: false, F1: -979, F2: 42}
                                                _ = v68
                                            case "foo bar":
                                                if true {
                                                    switch fi10 {
                                                    case 9223372036854775807:
                                                        if false {
                                                            v8 = v48
                                                        } else {
                                                            v69 := GenS5{F0: (v12 * v12)}
                                                            _ = v69
                                                        }
                                                    }
                                                    var v70 uint64 = (v8 + (^(v48)))
                                                    _ = v70
                                                }
                                            case "":
                                                if (147 < v48) {
                                                    fmt.Println(v12, v3)
                                                    v71 := GenS9{F0: (-9223372036854775808 >= fi9)}
                                                    v3 -= v3
                                                    _ = v71
                                                }
                                            default:
                                                v48 = v48
                                            }
                                            var v72 int8 = -1
                                            _ = v67
                                            _ = v72
                                        } else {
                                            if (2 != 395) {
                                                if true {
                                                    v73 := GenS9{F0: (-2 == -40)}
                                                    v74 := GenS10{F0: 0, F1: (255 > 2)}
                                                    var v75 int16 = 42
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                            } else {
                                                var v76 int64 = (v3 >> ((2) & 63))
                                                var v77 uint = 0
                                                v48 -= v48
                                                _ = v76
                                                _ = v77
                                            }
                                            v78 := GenS5{F0: (v12 << ((452) & 63))}
                                            if true {
                                                v48 -= 2
                                                if false {
                                                    v79 := GenS7{F0: -1, F1: 2}
                                                    if (1.5 < 3.14159) {
                                                        var v80 uint8 = 2
                                                        var v81 int8 = 51
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        var v82 int32 = 42
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v82
                                                    }
                                                    var v83 bool = false
                                                    if (!true) {
                                                        var v84 uint = 820
                                                        _ = v84
                                                    } else {
                                                        delete(v2, v12)
                                                        if (65535 >= 4294967295) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            v83 = v83
                                                        } else {
                                                            v85 := GenS6{F0: 11, F1: v47}
                                                            _ = v85
                                                        }
                                                    }
                                                    _ = v79
                                                    _ = v83
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = 4294967295
                                                }
                                                switch (fi9 | fi10) {
                                                case -1:
                                                    v8 -= 0
                                                }
                                            }
                                            _ = v78
                                        }
                                        v86 := v2[v12]
                                        v48 += (v8 * (v48 ^ 2))
                                        _ = v86
                                    }
                                } else {
                                    var v87 int = (fi9 * (fi10 / ((fi10) | 1)))
                                    var v88 float64 = 3.14159
                                    switch "" {
                                    case "":
                                        switch (fi10 | -9223372036854775808) {
                                        case 144:
                                            var v89 int = (v87 * 0)
                                            _ = v89
                                        case 15:
                                            v48 = (v48 / (((v48 & (v48 ^ (v48 % (((^(v8))) | 1))))) | 1))
                                        }
                                    case "foo bar":
                                        v8 = v48
                                        fmt.Println(v88, v48, v87)
                                        v90 := GenS8{F0: (v12 & v12)}
                                        _ = v90
                                    case "a":
                                        if (false && true) {
                                            v88 = (v88 / v88)
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                    _ = v87
                                    _ = v88
                                }
                                if true {
                                    var v91 int16 = (v47 << ((451) & 63))
                                    if false {
                                        switch ("" + ("foo bar" + "hello")) {
                                        case "\n\t":
                                            var v92 float32 = (3.14159 / (1.5 - 3.14159))
                                            _ = v92
                                        case "hello":
                                            v47 = (v47 << ((255) & 63))
                                        }
                                        v2[(v12 / (((v12 & (v12 * 4294967295))) | 1))] = v12
                                    }
                                    _ = v91
                                } else {
                                    fmt.Println(fi9)
                                }
                                _ = v47
                                _ = v48
                            }
                            fmt.Println(fi9)
                            _ = v22
                        }
                        var v93 int8 = 26
                        var v94 uint32 = 1
                        _ = v93
                        _ = v94
                    }
                    if false {
                        if true {
                            var v95 int = (fi9 / ((fi10) | 1))
                            _ = v95
                        }
                        var v96 int64 = v3
                        _ = v96
                    } else {
                        var v97 float64 = 1.5
                        _ = v97
                    }
                case 32768:
                    var v98 uint16 = 0
                    _ = v98
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(fi10)
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    v8 = v8
    v3 += (-(v3))
    v99 := GenS2{F0: 11, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v99
}
