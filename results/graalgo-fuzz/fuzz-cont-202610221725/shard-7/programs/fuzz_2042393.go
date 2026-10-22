package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint; F2 string }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 bool }
func genFn0() int16 {
    switch ("\n\t" + ("a" + "foo bar")) {
    case "foo bar":
        v4 := GenS1{F0: 1, F1: -458, F2: 42}
        _ = v4
    case "hello":
        if ((2 > 255) || ((320 != 1) && (3.14159 == (-2.25 * float64(1.5))))) {
            var v5 uint = 255
            switch ("hello" + "hello") {
            case "a":
                if (!true) {
                    if (!(9223372036854775807 != 9223372036854775807)) {
                        var v6 int32 = 826
                        _ = v6
                    }
                } else {
                    for fi7 := 0; fi7 < 10; fi7++ {
                        var v8 uint8 = 29
                        for fi9 := 0; fi9 < 11; fi9++ {
                            v5 = (^(v5))
                            _ = fi9
                        }
                        var v10 int32 = -841
                        _ = fi7
                        _ = v10
                        _ = v8
                    }
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if true {
                            v12 := []int{fi11, -218}
                            _ = v12
                        }
                        v5 += (v5 ^ v5)
                        _ = fi11
                    }
                }
            case "\n\t":
                switch -9223372036854775808 {
                case 127:
                    fmt.Println(v5)
                case 137:
                    v5 -= 376
                case 330:
                    v13 := []int{-287, -2147483648, 9223372036854775807, 992}
                    _ = v13
                }
            case "":
                v5 = (v5 | (^(v5)))
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                v5 += (^(v5))
                if false {
                    v15 := GenS1{F0: -128, F1: -48, F2: 0}
                    if (2 >= v5) {
                        var v16 int32 = 2147483647
                        v17 := []int{1, int(v16), (fi14 % (((-(fi14))) | 1))}
                        _ = v16
                        _ = v17
                    }
                    if true {
                        v18 := []int{32767, -1}
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 string = "foo bar"
                            if (-1 == 1) {
                                fmt.Println(fi19, fi19)
                                if false {
                                    v21 := GenS2{F0: 0}
                                    var v22 int8 = -1
                                    v23 := GenS1{F0: (v22 / ((14) | 1)), F1: -9223372036854775808, F2: -2147483648}
                                    v18 = append(v18, fi14)
                                    switch v20 {
                                    case "\n\t":
                                        if (!false) {
                                            if true {
                                                var v24 uint16 = uint16(v5)
                                                v25 := GenS3{F0: false}
                                                if true {
                                                    if true {
                                                        fmt.Println(v5, v24, v20)
                                                        fmt.Println(v5, fi19, fi19)
                                                        v18 = append(v18, -1)
                                                        fmt.Println(fi19, v22, fi19)
                                                    } else {
                                                        v5 = (v5 / ((v5) | 1))
                                                        switch "" {
                                                        case "a":
                                                            v26 := GenS3{F0: true}
                                                            _ = v26
                                                        case "foo bar":
                                                            if len(v18) >= 2 {
                                                                v18 = v18[1:len(v18)]
                                                            }
                                                        default:
                                                            var v27 int32 = 0
                                                            _ = v27
                                                        }
                                                        v18 = append(v18, fi19)
                                                        if len(v18) >= 2 {
                                                            v18 = v18[1:len(v18)]
                                                        }
                                                    }
                                                    v20 = v20
                                                }
                                                switch -677 {
                                                case 672:
                                                    fmt.Println(fi14)
                                                }
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v5 -= (^(v5))
                                            }
                                        }
                                    case "":
                                        fmt.Println(v20, v20, v20)
                                    default:
                                        fmt.Println(fi14, fi14, v22)
                                        switch v20 {
                                        case "hello":
                                            v22 = (v22 ^ (v22 >> ((v5) & 63)))
                                        case "foo bar":
                                            _ = len(v18)
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                        case "\n\t":
                                            var v28 int32 = 42
                                            _ = v28
                                        default:
                                            var v29 bool = (uint16(v5) < 65535)
                                            _ = v29
                                        }
                                    }
                                    v30 := GenS4{F0: (-(fi19)), F1: 18446744073709551615}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v30
                                }
                                var v31 int32 = -843
                                v32 := GenS3{F0: ((!true) && false)}
                                _ = v31
                                _ = v32
                            } else {
                                if false {
                                    var v33 int = fi14
                                    var v34 uint8 = 2
                                    _ = v33
                                    _ = v34
                                }
                            }
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    }
                    _ = v15
                }
                fmt.Println(v5, fi14)
                _ = fi14
            }
            _ = v5
        } else {
            if false {
                v35 := []int{-9223372036854775808, 663}
                _ = v35
            } else {
                switch 807 {
                case -909:
                    v36 := []int{42, 127, -280, -430}
                    _ = v36
                case 32768:
                    var v37 uint64 = 846
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 4; fi38++ {
                        switch ("foo bar" + ("\n\t" + ("foo bar" + "a"))) {
                        case "":
                            var v39 int32 = 0
                            _ = v39
                        case "hello":
                            switch "hello" {
                            case "foo bar":
                                var v40 uint8 = 212
                                _ = v40
                            case "a":
                                for fi41 := 0; fi41 < 14; fi41++ {
                                    switch (fi38 | fi41) {
                                    case 492:
                                        var v42 uint = 255
                                        _ = v42
                                    default:
                                        v43 := []int{(-(fi41))}
                                        _ = v43
                                    }
                                    _ = fi41
                                }
                            case "\n\t":
                                var v44 int = (-(fi38))
                                _ = v44
                            default:
                                var v45 int8 = 2
                                _ = v45
                            }
                        case "a":
                            switch 32767 {
                            case 127:
                                if false {
                                    var v46 uint32 = 2147483647
                                    var v47 uint32 = v46
                                    _ = v46
                                    _ = v47
                                }
                            case -2147483648:
                                v48 := GenS5{F0: 18446744073709551615}
                                _ = v48
                            case -433:
                                if ((0.0 / 0.1) > 0.0) {
                                    var v49 string = "a"
                                    var v50 uint8 = 14
                                    var v51 float64 = 0.0
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                            }
                        default:
                            fmt.Println(fi38)
                        }
                        _ = fi38
                    }
                }
                var v52 string = "hello"
                _ = v52
            }
            var v53 int64 = 120
            v54 := []int{}
            _ = v53
            _ = v54
        }
    default:
        var v55 float64 = (0.1 + (1e10 + 1e10))
        _ = v55
    }
    var v56 uint64 = 952
    var v57 int8 = -38
    v58 := GenS1{F0: (v57 % ((-103) | 1)), F1: 0, F2: -1}
    _ = v56
    _ = v57
    _ = v58
    return 0
}

func main() {
    v1 := []string{"", ("foo bar" + "hello"), "foo bar"}
    var v2 int8 = -128
    v3 := make(map[string]string)
    v1 = append(v1, "hello")
    v59 := GenS6{F0: 2147483647}
    var v60 int32 = -1
    if false {
        for fi61 := 0; fi61 < 19; fi61++ {
            switch (fi61 << ((311) & 63)) {
            case -904:
                if false {
                    var v62 int32 = -2147483648
                    var v63 string = (("\n\t" + (("\n\t" + "a") + "hello")) + "\n\t")
                    _ = v62
                    _ = v63
                } else {
                    if (-2.25 != (-2.25 * ((float64(1.5) - (((1.5 - -2.25) / -2.25) / 3.14159)) + (float64(3.14159) * 0.1)))) {
                        if (!(0.0 <= ((1e10 * (1.5 - float64((0.1 / (1e10 * (1e10 * 0.1)))))) / 3.14159))) {
                            var v64 uint = 42
                            _ = v64
                        }
                        var v65 int16 = 1000
                        for fi66 := 0; fi66 < 8; fi66++ {
                            var v67 int16 = (-(v65))
                            fmt.Println(fi66)
                            v68 := GenS2{F0: 476}
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                        _ = v65
                    }
                    v69 := GenS7{F0: 18446744073709551615, F1: 255, F2: "hello"}
                    _ = v69
                }
            case -1:
                var v70 float32 = 3.14159
                _ = v70
            }
            v71 := GenS7{F0: 2, F1: 18446744073709551615, F2: ((((("a" + "foo bar") + "hello") + "") + "") + "hello")}
            _ = fi61
            _ = v71
        }
        fmt.Println(v2, v60)
    }
    v2 = 127
    v2 = v2
    if false {
        var v72 float64 = -2.25
        _ = v72
    }
    fmt.Println(v60)
    if false {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v73 uint64 = 2
            switch 4294967295 {
            case 0:
                v60 -= -2147483648
            case 1:
                fmt.Println(v73, v73)
            }
            _ = v73
        } else {
            for fi74 := 0; fi74 < 3; fi74++ {
                var v75 bool = true
                v76 := GenS8{F0: 127, F1: 2}
                var v77 uint16 = 1
                _ = fi74
                _ = v75
                _ = v76
                _ = v77
            }
        }
        switch 127 {
        case -9223372036854775808:
            for fi78 := 0; fi78 < 11; fi78++ {
                switch "" {
                case "\n\t":
                    v2 -= v2
                case "foo bar":
                    v79 := GenS9{F0: (v60 - (-(v60))), F1: 32767}
                    _ = v79
                case "hello":
                    for fi80 := 0; fi80 < 20; fi80++ {
                        v81 := GenS10{F0: 4294967295, F1: v60}
                        _ = fi80
                        _ = v81
                    }
                }
                var v82 uint16 = 65535
                for fi83 := 0; fi83 < 14; fi83++ {
                    switch (^(fi78)) {
                    case 9223372036854775807:
                        v84 := GenS11{F0: (((!false) && (false || true)) || false)}
                        _ = v84
                    case -509:
                        if false {
                            v82 += 65535
                            fmt.Println(v60)
                            v2 = -1
                        }
                    case 583:
                        if (-9223372036854775808 >= 324) {
                            if true {
                                v82 -= (v82 * (v82 % ((997) | 1)))
                            } else {
                                switch fi83 {
                                case 42:
                                    var v85 bool = (1000 >= 1)
                                    _ = v85
                                default:
                                    if false {
                                        _ = len(v1)
                                        if true {
                                            if (v2 == -1) {
                                                delete(v3, "a")
                                                var v86 uint8 = 166
                                                _ = v86
                                            }
                                        }
                                        var v87 int32 = v60
                                        _ = v87
                                    } else {
                                        fmt.Println(v82)
                                        var v88 uint16 = 42
                                        _ = v88
                                    }
                                    var v89 uint = 1
                                    _ = v89
                                }
                                v82 -= 0
                            }
                            switch -1 {
                            case 0:
                                v82 += (^(v82))
                            case -9223372036854775808:
                                if false {
                                    v3["\n\t"] = "hello"
                                } else {
                                    var v90 int8 = v2
                                    _ = v90
                                }
                            case -797:
                                v3[("hello" + "hello")] = "\n\t"
                            }
                            var v91 int = 128
                            _ = v91
                        }
                    }
                    var v92 int8 = (^(v2))
                    var v93 int16 = 0
                    _ = fi83
                    _ = v92
                    _ = v93
                }
                _ = fi78
                _ = v82
            }
        case -926:
            v1 = append(v1, "foo bar")
        default:
            for fi94 := 0; fi94 < 16; fi94++ {
                var v95 int64 = -645
                _ = fi94
                _ = v95
            }
        }
    } else {
        v96 := v3[("" + "a")]
        _ = v96
    }
    for fi97 := 0; fi97 < 17; fi97++ {
        if true {
            if (!false) {
                var v98 uint = 633
                var v99 uint64 = uint64(v98)
                fmt.Println(v99, v98)
                _ = v98
                _ = v99
            }
            v60 += -748
            v2 -= (v2 / (((v2 / ((v2) | 1))) | 1))
        } else {
            var v100 int32 = 20
            v101 := GenS7{F0: 2, F1: 239, F2: "\n\t"}
            fmt.Println(v2, v100)
            _ = v100
            _ = v101
        }
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v2 = 100
        default:
            fmt.Println(fi97, v60, v60)
        }
        var v102 uint16 = 255
        _ = fi97
        _ = v102
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v60
}
