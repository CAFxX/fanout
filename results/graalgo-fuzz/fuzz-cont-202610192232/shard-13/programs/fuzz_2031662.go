package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS8 struct { F0 uint; F1 int16; F2 int64 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0() int {
    var v4 uint16 = 1
    v5 := GenS1{F0: v4, F1: 128}
    v4 = v4
    var v6 uint8 = 1
    var v7 int = -653
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 int32 = -222
    var v2 uint8 = 128
    v3 := make(map[int]uint32)
    var v8 int = 185
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int = (fi9 + (fi9 * genFn0()))
        fmt.Println(v8)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v3[127] = 621
        _ = fi11
    }
    if (true || false) {
        v2 = (v2 >> ((42) & 63))
    }
    fmt.Println(v2, v8, v1)
    if (427 > 2) {
        switch ("hello" + ("a" + "foo bar")) {
        case "hello":
            v12 := GenS2{F0: 0}
            _ = v12
        case "foo bar":
            var v13 uint8 = (v2 & (v2 + 201))
            _ = v13
        default:
            for fi14 := 0; fi14 < 3; fi14++ {
                if false {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        if true {
                            var v16 float64 = 3.14159
                            if true {
                                if (true && true) {
                                    if (("" + "\n\t") >= (((("foo bar" + "a") + "a") + "foo bar") + ("hello" + ("a" + "\n\t")))) {
                                        fmt.Println(v8, fi14)
                                        v1 = 1
                                        var v17 int64 = -867
                                        var v18 bool = (892 != 0)
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 uint64 = 4294967295
                                        if (false && true) {
                                            switch ("hello" + "a") {
                                            case "a":
                                                var v20 int8 = 0
                                                _ = v20
                                            default:
                                                v21 := GenS3{F0: v2, F1: v2}
                                                _ = v21
                                            }
                                            if true {
                                                switch 2147483647 {
                                                case 331:
                                                    _ = len(v3)
                                                    v16 = 1.5
                                                }
                                                v22 := GenS4{F0: 0}
                                                _ = v22
                                            }
                                            var v23 uint8 = v2
                                            _ = v23
                                        }
                                        _ = v19
                                    }
                                    var v24 uint16 = 255
                                    switch (v8 ^ fi14) {
                                    case 934:
                                        switch "a" {
                                        case "hello":
                                            var v25 int16 = -32768
                                            _ = v25
                                        case "foo bar":
                                            v3[(fi15 - fi14)] = 1
                                        case "\n\t":
                                            v26 := GenS5{F0: (("a" + "a") + "foo bar")}
                                            _ = v26
                                        }
                                    }
                                    _ = v24
                                }
                            }
                            _ = v16
                        } else {
                            var v27 uint32 = 895
                            _ = v27
                        }
                        v1 += int32(fi15)
                        _ = fi15
                    }
                    var v28 uint64 = 4294967295
                    var v29 uint = 48
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS6{F0: 1, F1: (v1 % (((v1 / ((v1) | 1))) | 1)), F2: 1}
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        v1 -= v1
                    case "foo bar":
                        v8 += 2147483647
                    case "hello":
                        fmt.Println(v8, v8)
                    }
                    var v31 uint8 = 1
                    _ = v30
                    _ = v31
                }
                var v32 uint64 = 1
                _ = fi14
                _ = v32
            }
        }
    } else {
        switch (v8 ^ (v8 ^ (v8 ^ (v8 / ((-1) | 1))))) {
        case 42:
            var v33 int32 = -723
            _ = v33
        case 128:
            if true {
                switch ("foo bar" + "") {
                case "a":
                    var v34 uint16 = 255
                    _ = v34
                case "\n\t":
                    if false {
                        fmt.Println(v1, v2, v8)
                        v1 -= v1
                    } else {
                        switch v8 {
                        case -877:
                            if false {
                                v1 += v1
                            }
                        default:
                            for fi35 := 0; fi35 < 12; fi35++ {
                                for fi36 := 0; fi36 < 1; fi36++ {
                                    v2 = (v2 ^ 1)
                                    var v37 int32 = -1
                                    switch (("" + "") + ("" + "hello")) {
                                    case "":
                                        var v38 int8 = int8(fi35)
                                        _ = v38
                                    case "\n\t":
                                        if true {
                                            v1 -= v37
                                            v3[v8] = 4294967295
                                            var v39 int = int(v1)
                                            switch (("hello" + ("" + (("a" + ((("" + ("hello" + (("a" + "a") + "hello"))) + "") + ("foo bar" + ((("\n\t" + "\n\t") + "a") + "\n\t")))) + "\n\t"))) + "\n\t") {
                                            case "\n\t":
                                                if (2 >= 2147483647) {
                                                    var v40 int32 = v1
                                                    v41 := GenS1{F0: 65535, F1: fi35}
                                                    var v42 int32 = (v40 % ((v1) | 1))
                                                    var v43 uint = 0
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    if true {
                                                        switch 1 {
                                                        case -2147483648:
                                                            var v44 int16 = 42
                                                            _ = v44
                                                        }
                                                    } else {
                                                        v1 = (v1 ^ (^(v37)))
                                                        var v45 int16 = -32768
                                                        _ = v45
                                                    }
                                                    if true {
                                                        if true {
                                                            v46 := []int{(fi36 % (((^(fi35))) | 1)), fi35, (fi36 / (((fi36 * (-(v39)))) | 1))}
                                                            var v47 float32 = 3.14159
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            if ((-2147483648 < -744) || (false || true)) {
                                                                if (!(!false)) {
                                                                    _ = len(v3)
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        v48 := v3[int(v2)]
                                                                        _ = v48
                                                                    case "":
                                                                        fmt.Println(v8, v39)
                                                                    case "a":
                                                                        v49 := GenS7{F0: 295, F1: 873, F2: 42}
                                                                        _ = v49
                                                                    }
                                                                }
                                                            }
                                                            v50 := GenS6{F0: 48, F1: int32(v39), F2: 4294967295}
                                                            v51 := v3[-298]
                                                            v52 := GenS7{F0: 2, F1: (v51 | v51), F2: 255}
                                                            if (!false) {
                                                                if ((^(v51)) == (v51 << ((1) & 63))) {
                                                                    v2 += v2
                                                                    var v53 int = fi35
                                                                    fmt.Println(v53, fi35)
                                                                    _ = v53
                                                                }
                                                            } else {
                                                                var v54 uint32 = 641
                                                                _ = v54
                                                            }
                                                            fmt.Println(v39)
                                                            v2 = (v2 ^ (v2 / ((v2) | 1)))
                                                            _ = v50
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    } else {
                                                        var v55 int64 = -9223372036854775808
                                                        fmt.Println(v37, v37, fi35)
                                                        _ = v55
                                                    }
                                                    v56 := GenS5{F0: ("hello" + (("hello" + "") + "hello"))}
                                                    _ = v56
                                                }
                                                v57 := []int{}
                                                v58 := GenS8{F0: 1, F1: -457, F2: 1}
                                                _ = v57
                                                _ = v58
                                            case "hello":
                                                v3[32767] = 531
                                            }
                                            if ((v39 % (((v39 >> ((0) & 63))) | 1)) < genFn0()) {
                                                if (false || (0.0 > (-2.25 / 0.1))) {
                                                    if false {
                                                        v2 -= (v2 >> ((18446744073709551615) & 63))
                                                        v8 -= 42
                                                        v37 -= (v1 / ((0) | 1))
                                                    }
                                                    if (!true) {
                                                        v8 = 42
                                                        var v59 string = ""
                                                        switch ((v59 + "") + ("foo bar" + v59)) {
                                                        case "foo bar":
                                                            var v60 int64 = 9223372036854775807
                                                            if true {
                                                                if true {
                                                                    v61 := GenS8{F0: 42, F1: 54, F2: 0}
                                                                    var v62 bool = (!true)
                                                                    var v63 float64 = 0.1
                                                                    _ = v61
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                v64 := GenS1{F0: 32767, F1: (v8 ^ (fi36 * fi36))}
                                                                switch 2147483647 {
                                                                case 127:
                                                                    v2 = (v2 << ((548) & 63))
                                                                default:
                                                                    var v65 uint32 = 4294967295
                                                                    var v66 bool = (4294967295 < 18446744073709551615)
                                                                    var v67 int = fi35
                                                                    _ = v65
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                _ = v64
                                                            }
                                                            _ = v60
                                                        case "\n\t":
                                                            v68 := GenS9{F0: (v37 % (((v37 - (^(v1)))) | 1)), F1: (v37 / ((42) | 1))}
                                                            _ = v68
                                                        }
                                                        _ = v59
                                                    }
                                                }
                                                v69 := GenS10{F0: v2}
                                                _ = v69
                                            }
                                            _ = v39
                                        } else {
                                            v2 -= v2
                                            var v70 int8 = 127
                                            _ = v70
                                        }
                                    }
                                    v71 := GenS10{F0: v2}
                                    var v72 bool = (!(!("foo bar" >= ("hello" + "a"))))
                                    v72 = v72
                                    _ = fi36
                                    _ = v37
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi35
                            }
                        }
                        fmt.Println(v2)
                        v73 := GenS4{F0: 986}
                        _ = v73
                    }
                case "":
                    if false {
                        var v74 string = "hello"
                        v1 = 2147483647
                        _ = v74
                    }
                default:
                    for fi75 := 0; fi75 < 9; fi75++ {
                        var v76 int = (v8 * fi75)
                        if (255 == 42) {
                            var v77 int = v76
                            _ = v77
                        }
                        if false {
                            switch v8 {
                            case 198:
                                fmt.Println(v8)
                            case -546:
                                if true {
                                    var v78 int = (v8 - v8)
                                    var v79 int32 = (-(v1))
                                    _ = v78
                                    _ = v79
                                }
                            default:
                                for fi80 := 0; fi80 < 1; fi80++ {
                                    v8 = int(v1)
                                    v81 := GenS11{F0: 42}
                                    if false {
                                        var v82 int16 = 0
                                        _ = v82
                                    } else {
                                        if (-38 < 100) {
                                            v8 += (^(v76))
                                        } else {
                                            _ = len(v3)
                                            v83 := GenS12{F0: -81, F1: (18446744073709551615 == 119)}
                                            var v84 int8 = -128
                                            _ = v83
                                            _ = v84
                                        }
                                    }
                                    _ = fi80
                                    _ = v81
                                }
                            }
                        }
                        _ = fi75
                        _ = v76
                    }
                }
                v2 = 1
                var v85 float32 = 0.0
                _ = v85
            } else {
                delete(v3, v8)
            }
        default:
            delete(v3, genFn0())
        }
        var v86 int16 = int16(v8)
        _ = v86
    }
    v3[-279] = 65535
    var v87 bool = (0.1 <= -2.25)
    v88 := GenS13{F0: 619}
    for fi89 := 0; fi89 < 4; fi89++ {
        v90 := GenS14{F0: 594, F1: 42, F2: 1000}
        v91 := []int{fi89, 32767}
        for fi92 := 0; fi92 < 3; fi92++ {
            var v93 uint32 = 2147483647
            _ = len(v3)
            switch ((("" + "\n\t") + ("foo bar" + ("foo bar" + "hello"))) + "\n\t") {
            case "a":
                if (v93 >= (v93 - v93)) {
                    var v94 int = genFn0()
                    _ = v94
                }
            case "\n\t":
                var v95 uint64 = 756
                _ = v95
            }
            v93 = (^(v93))
            _ = fi92
            _ = v93
        }
        _ = fi89
        _ = v90
        _ = v91
    }
    var v96 uint64 = 18446744073709551615
    var v97 string = ""
    fmt.Println(v8, v87, v8)
    v98 := GenS9{F0: (v1 - 42), F1: v1}
    v99 := v3[v8]
    v100 := []int{(v8 / ((v8) | 1)), v8, genFn0()}
    var v101 uint = 3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v87:", v87)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v8
    _ = v87
    _ = v88
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
