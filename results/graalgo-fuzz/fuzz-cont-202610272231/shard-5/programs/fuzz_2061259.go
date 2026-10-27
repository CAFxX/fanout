package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int; F2 uint32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint16) int8 {
    if false {
        p1 += 2
    }
    p0 += 127
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int16 = int16(p0)
        var v8 int = int(p0)
        var v9 int64 = int64(p0)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{551, 42, -506, -1}
    _ = len(v10)
    _ = p0
    _ = p1
    _ = v10
    return 50
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = (0.0 + 0.1)
    v3 := make(map[string]bool)
    var v4 uint8 = 2
    var v5 int16 = -32768
    for fi11 := 0; fi11 < 10; fi11++ {
        if true {
            var v12 string = ""
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint16 = 577
            var v15 string = (("" + "\n\t") + "\n\t")
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 bool = false
            _ = fi16
            _ = v17
        }
        _ = fi11
    }
    if true {
        if true {
            if true {
                var v18 float32 = ((0.0 * float32(v2)) - (0.1 * 3.14159))
                _ = v18
            }
            v19 := []int{-2147483648, 320, 0, int(v1)}
            var v20 bool = (true || true)
            _ = v19
            _ = v20
        } else {
            switch 128 {
            case 0:
                v21 := GenS1{F0: 2, F1: true}
                _ = v21
            case -2147483648:
                var v22 int64 = -375
                _ = v22
            default:
                if (4294967295 > 657) {
                    switch 0 {
                    case 182:
                        var v23 int64 = 288
                        _ = v23
                    }
                    v24 := v3["foo bar"]
                    switch -1 {
                    case 128:
                        if false {
                            v25 := []int{32767, int(v5), 42}
                            _ = v25
                        }
                    case -2147483648:
                        fmt.Println(v4, v1)
                    }
                    _ = v24
                } else {
                    var v26 float64 = -2.25
                    _ = v26
                }
            }
            v5 = v5
            var v27 uint = 42
            _ = v27
        }
    } else {
        var v28 uint8 = v4
        v5 = v5
        _ = v28
    }
    var v29 int64 = 42
    fmt.Println(v5, v2, v29)
    var v30 uint32 = 42
    v31 := []int{1}
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS1{F0: 2, F1: (true && (!false))}
        if true {
            if false {
                if len(v31) > 0 {
                    v31[0] = fi32
                }
                switch (("" + "a") + "a") {
                case "\n\t":
                    var v34 uint32 = 33
                    _ = v34
                case "foo bar":
                    var v35 int32 = int32(v4)
                    _ = v35
                }
            } else {
                for fi36 := 0; fi36 < 11; fi36++ {
                    v29 += (v29 ^ v29)
                    v37 := GenS2{F0: v1, F1: true}
                    _ = fi36
                    _ = v37
                }
                v38 := GenS1{F0: 913, F1: false}
                v39 := GenS1{F0: 42, F1: true}
                _ = v38
                _ = v39
            }
            _ = len(v3)
            v40 := GenS1{F0: 18446744073709551615, F1: (true || (!(int8(v5) <= -1)))}
            _ = v40
        }
        var v41 int8 = 6
        _ = fi32
        _ = v33
        _ = v41
    }
    fmt.Println(v30, v30)
    for fi42 := 0; fi42 < 5; fi42++ {
        if false {
            fmt.Println(v29, v29)
            switch fi42 {
            case 972:
                v43 := GenS2{F0: v1, F1: (((3.14159 * (3.14159 * 1.5)) / 3.14159) <= 3.14159)}
                _ = v43
            case -721:
                switch ("" + ("" + "hello")) {
                case "foo bar":
                    if true {
                        if true {
                            var v44 int = (fi42 >> ((18446744073709551615) & 63))
                            var v45 int16 = (v5 + (v5 & 42))
                            for fi46 := 0; fi46 < 20; fi46++ {
                                if false {
                                    v47 := GenS1{F0: 0, F1: false}
                                    _ = v47
                                } else {
                                    if ((!false) || true) {
                                        v4 = 144
                                    }
                                    var v48 float32 = -2.25
                                    v30 = (v30 >> ((65535) & 63))
                                    if true {
                                        switch "" {
                                        case "foo bar":
                                            var v49 int8 = -101
                                            _ = v49
                                        default:
                                            if true {
                                                if true {
                                                    v45 -= v45
                                                    switch (fi42 | 2147483647) {
                                                    case 32768:
                                                        var v50 uint16 = (v1 / ((v1) | 1))
                                                        _ = v50
                                                    case 898:
                                                        var v51 uint16 = v1
                                                        _ = v51
                                                    case -1:
                                                        v52 := GenS1{F0: 18446744073709551615, F1: ("a" >= ("\n\t" + ""))}
                                                        _ = v52
                                                    default:
                                                        var v53 string = ("" + ("a" + ""))
                                                        _ = v53
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch ("foo bar" + "foo bar") {
                                                    case "a":
                                                        fmt.Println(v45)
                                                    case "":
                                                        var v54 bool = (2147483647 > 0)
                                                        _ = v54
                                                    case "\n\t":
                                                        v55 := GenS2{F0: 577, F1: false}
                                                        _ = v55
                                                    }
                                                    v48 -= v48
                                                    fmt.Println(v5)
                                                    var v56 int16 = v5
                                                    var v57 int = 161
                                                    _ = v56
                                                    _ = v57
                                                case "":
                                                    var v58 int64 = (v29 - 812)
                                                    _ = v58
                                                default:
                                                    if (320 != 2) {
                                                        v1 = 65535
                                                    }
                                                }
                                                var v59 int16 = v5
                                                _ = v59
                                            } else {
                                                v60 := GenS3{F0: (v4 >> ((42) & 63))}
                                                _ = v60
                                            }
                                            if true {
                                                v61 := GenS2{F0: 199, F1: true}
                                                v3["a"] = ("foo bar" == ("" + ("" + "a")))
                                                switch v44 {
                                                case -1:
                                                    v29 -= (v29 ^ -994)
                                                case 42:
                                                    v4 = v4
                                                default:
                                                    var v62 uint16 = (v1 + (^(v1)))
                                                    _ = v62
                                                }
                                                _ = v61
                                            } else {
                                                var v63 int = (fi42 / (((v44 + (v44 << ((491) & 63)))) | 1))
                                                var v64 uint32 = 0
                                                if ((!(false && ((v1 | (v1 >> ((0) & 63))) == 42))) || false) {
                                                    var v65 float32 = 0.0
                                                    v63 += fi42
                                                    if false {
                                                        switch int(v4) {
                                                        case -2147483648:
                                                            var v66 uint8 = v4
                                                            _ = v66
                                                        case 127:
                                                            if (true && (533 <= (v30 - (v64 >> ((1) & 63))))) {
                                                                fmt.Println(v45)
                                                                v4 += (^(v4))
                                                                fmt.Println(v64, fi46)
                                                            } else {
                                                                switch fi42 {
                                                                case 32767:
                                                                    var v67 string = "a"
                                                                    _ = v67
                                                                case 1:
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v65 = 3.14159
                                                                    default:
                                                                        v68 := GenS3{F0: v4}
                                                                        _ = v68
                                                                    }
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            v69 := GenS2{F0: 1, F1: false}
                                                            _ = v69
                                                        default:
                                                            var v70 int64 = (^(v29))
                                                            _ = v70
                                                        }
                                                        var v71 uint16 = v1
                                                        v3[("hello" + ("foo bar" + "hello"))] = true
                                                        _ = v71
                                                    }
                                                    _ = v65
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                    } else {
                                        var v72 int8 = 127
                                        var v73 bool = false
                                        var v74 uint64 = 18446744073709551615
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                    if false {
                                        v75 := GenS4{F0: (fi46 + v44), F1: (v44 / ((803) | 1)), F2: v30}
                                        var v76 int = v44
                                        var v77 uint32 = v30
                                        v48 -= (((v48 + ((-2.25 - v48) / v48)) + 3.14159) - (float32(v2) * 1e10))
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    }
                                    _ = v48
                                }
                                switch "a" {
                                case "foo bar":
                                    v30 = 2
                                }
                                _ = fi46
                            }
                            _ = v44
                            _ = v45
                        } else {
                            _ = len(v31)
                        }
                        switch 42 {
                        case 0:
                            if true {
                                var v78 bool = true
                                _ = v78
                            }
                        case -591:
                            for fi79 := 0; fi79 < 20; fi79++ {
                                var v80 int = (^(fi42))
                                var v81 int16 = v5
                                _ = fi79
                                _ = v80
                                _ = v81
                            }
                        case 32768:
                            for fi82 := 0; fi82 < 16; fi82++ {
                                v2 = v2
                                _ = fi82
                            }
                        default:
                            v30 += (v30 | v30)
                        }
                    }
                case "":
                    switch (fi42 + 2147483647) {
                    case 671:
                        for fi83 := 0; fi83 < 5; fi83++ {
                            delete(v3, "a")
                            if false {
                                switch fi83 {
                                case -2147483648:
                                    v2 += float64((1.5 / (((((-2.25 - (0.0 * ((((1.5 / 1e10) / float32(1.5)) * 1.5) / (3.14159 / 3.14159)))) + 1.5) - 1e10) / 0.1) + -2.25)))
                                }
                            }
                            var v84 string = "foo bar"
                            _ = fi83
                            _ = v84
                        }
                    case 9223372036854775807:
                        if (v1 <= (v1 % ((0) | 1))) {
                            for fi85 := 0; fi85 < 18; fi85++ {
                                v5 = (v5 & (^(v5)))
                                switch "foo bar" {
                                case "hello":
                                    _ = len(v31)
                                case "\n\t":
                                    var v86 string = "hello"
                                    _ = v86
                                case "":
                                    v2 += v2
                                default:
                                    v29 += (v29 / (((v29 + 590)) | 1))
                                }
                                var v87 int8 = 0
                                _ = fi85
                                _ = v87
                            }
                            var v88 uint = 296
                            fmt.Println(v88, v1)
                            _ = v88
                        }
                    }
                case "\n\t":
                    if true {
                        v5 += 1
                        for fi89 := 0; fi89 < 10; fi89++ {
                            switch "foo bar" {
                            case "hello":
                                delete(v3, "")
                            case "\n\t":
                                var v90 float32 = (-2.25 - (-2.25 - 0.0))
                                _ = v90
                            default:
                                switch 0 {
                                case 44:
                                    if len(v31) > 0 {
                                        v31[0] = fi42
                                    }
                                case 297:
                                    v2 -= ((v2 * (((((v2 - 0.1) / float64(1e10)) / float64(0.1)) - v2) + 1.5)) + v2)
                                }
                                fmt.Println(fi89)
                                var v91 uint32 = v30
                                _ = v91
                            }
                            _ = fi89
                        }
                        for fi92 := 0; fi92 < 17; fi92++ {
                            if (196 != 225) {
                                var v93 int8 = 100
                                v94 := GenS4{F0: (fi92 ^ 32767), F1: int(v4), F2: (v30 * 1)}
                                if len(v31) > 0 {
                                    v31[0] = -2147483648
                                }
                                _ = v93
                                _ = v94
                            }
                            _ = fi92
                        }
                    } else {
                        if len(v31) > 0 {
                            v31[0] = (^(fi42))
                        }
                    }
                }
            case -75:
                v31 = append(v31, fi42)
            default:
                delete(v3, ("hello" + "foo bar"))
            }
        } else {
            var v95 int8 = 1
            _ = v95
        }
        fmt.Println(v1, v4)
        switch (fi42 * fi42) {
        case -1:
            var v96 uint = 0
            _ = v96
        case 42:
            switch (("a" + ("a" + "hello")) + ("hello" + "a")) {
            case "hello":
                for fi97 := 0; fi97 < 9; fi97++ {
                    v3[("\n\t" + "")] = ((v29 | v29) == (v29 << ((255) & 63)))
                    var v98 uint16 = v1
                    _ = fi97
                    _ = v98
                }
            case "\n\t":
                v99 := GenS5{F0: -1, F1: v29, F2: 4294967295}
                _ = v99
            default:
                var v100 float64 = (0.1 + v2)
                _ = v100
            }
        }
        _ = fi42
    }
    var v101 int = 127
    delete(v3, "foo bar")
    for fi102 := 0; fi102 < 19; fi102++ {
        switch (^(v101)) {
        case 32767:
            v101 = (fi102 & (fi102 * (v101 >> ((65535) & 63))))
        case 1:
            v103 := GenS6{F0: v1}
            _ = v103
        case 596:
            fmt.Println(v101, v2, v4)
        default:
            switch "foo bar" {
            case "":
                var v104 int8 = 0
                _ = v104
            }
        }
        var v105 int8 = 42
        var v106 string = "\n\t"
        _ = fi102
        _ = v105
        _ = v106
    }
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
