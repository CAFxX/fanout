package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int32; F1 int32; F2 int64 }
type GenS8 struct { F0 uint16; F1 int32; F2 int32 }
type GenS9 struct { F0 int64; F1 uint64; F2 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int16; F1 string }
type GenS13 struct { F0 uint16; F1 uint8; F2 int }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint32; F1 int32 }
type GenS17 struct { F0 int; F1 uint }
type GenS18 struct { F0 int8; F1 uint }
type GenS19 struct { F0 int32; F1 bool }
type GenS20 struct { F0 int }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 int8; F1 int32; F2 int8 }
type GenS24 struct { F0 int32 }
type GenS25 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint32) uint {
    v6 := []int{127}
    if true {
        var v7 float64 = 1e10
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 bool) int {
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := []int{fi8}
        if len(v9) > 0 {
            v9[0] = (fi8 & 42)
        }
        fmt.Println(p0)
        _ = fi8
        _ = v9
    }
    var v10 int8 = 32
    var v11 uint16 = 0
    _ = p0
    _ = v10
    _ = v11
    return 2147483647
}

func main() {
    var v1 string = "a"
    var v2 int = 32767
    v3 := []string{}
    var v4 int = v2
    var v5 uint64 = 42
    switch v2 {
    case -419:
        v5 = (v5 | v5)
    default:
        var v12 int16 = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 bool = (true && false)
        _ = fi13
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS1{F0: 27, F1: 953}
            for fi17 := 0; fi17 < 5; fi17++ {
                _ = len(v3)
                if true {
                    v2 = genFn1((((v5 - v5) > (v5 - (^(v5)))) || false))
                } else {
                    v3 = append(v3, ("\n\t" + (v1 + v1)))
                    v18 := GenS1{F0: 45, F1: genFn0(-644, 836)}
                    v19 := GenS1{F0: 255, F1: 255}
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
            if false {
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := GenS1{F0: 0, F1: 18446744073709551615}
                    v22 := GenS1{F0: 128, F1: 4294967295}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                v23 := GenS2{F0: (v1 + "hello")}
                if false {
                    fmt.Println(v4)
                    if false {
                        v24 := GenS3{F0: 100, F1: "a"}
                        var v25 uint8 = 1
                        v26 := GenS3{F0: 76, F1: ("hello" + v1)}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v1 = (v1 + v1)
                }
                v2 -= fi15
                _ = v23
            }
            _ = fi15
            _ = v16
        }
    }
    switch 1 {
    case -444:
        for fi27 := 0; fi27 < 9; fi27++ {
            _ = len(v3)
            fmt.Println(v4, v1, v2)
            if (!("a" == ((v1 + "") + v1))) {
                v28 := GenS1{F0: 42, F1: genFn0(2147483647, 2147483648)}
                _ = v28
            }
            _ = fi27
        }
    default:
        if true {
            v1 = ("hello" + "\n\t")
        } else {
            var v29 uint = 585
            var v30 int = v2
            v31 := GenS4{F0: -1, F1: (^(v5))}
            _ = v29
            _ = v30
            _ = v31
        }
    }
    fmt.Println(v5, v1, v5)
    if ((-1 != 1) && true) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v32 := GenS5{F0: 18446744073709551615, F1: (v4 << ((0) & 63))}
    if true {
        if false {
            switch (v4 << ((255) & 63)) {
            case 42:
                v33 := GenS5{F0: 52, F1: v4}
                _ = v33
            default:
                _ = len(v3)
            }
            var v34 int = v2
            v35 := GenS6{F0: (v1 + ""), F1: false}
            _ = v34
            _ = v35
        }
        v36 := GenS7{F0: 2147483647, F1: 2147483647, F2: 469}
        _ = v36
    } else {
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "\n\t")
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            for fi38 := 0; fi38 < 19; fi38++ {
                v5 += 18446744073709551615
                v1 = v1
                fmt.Println(fi37)
                _ = fi38
            }
            if false {
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 string = v1
                    switch v40 {
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        if (false && false) {
                            if false {
                                var v41 bool = (!true)
                                _ = v41
                            } else {
                                v40 = v1
                                v42 := GenS8{F0: 822, F1: 0, F2: 360}
                                _ = v42
                            }
                            if (!false) {
                                var v43 uint8 = 128
                                _ = v43
                            }
                        } else {
                            switch 4294967295 {
                            case 128:
                                var v44 int8 = -128
                                _ = v44
                            }
                        }
                    }
                    switch (v1 + ("a" + ((("" + (("hello" + "hello") + (v1 + v40))) + v40) + v1))) {
                    case "foo bar":
                        v45 := GenS9{F0: 799, F1: (v5 % (((^(v5))) | 1)), F2: 1}
                        _ = v45
                    case "\n\t":
                        v1 = v1
                        v2 = (v4 | genFn1(true))
                    default:
                        fmt.Println(v1, v40)
                    }
                    _ = fi39
                    _ = v40
                }
            } else {
                var v46 string = ""
                v5 += 42
                _ = v46
            }
            _ = fi37
        }
        var v47 uint = 471
        _ = v47
    }
    switch (v2 & v4) {
    case -1:
        switch 113 {
        case 32767:
            switch (^(v4)) {
            case 45:
                fmt.Println(v4)
            default:
                if true {
                    v3 = append(v3, "hello")
                }
            }
        }
    case 128:
        switch v4 {
        case 0:
            switch v1 {
            case "hello":
                fmt.Println(v1, v5)
            }
        case 9223372036854775807:
            switch v1 {
            case "\n\t":
                v5 -= (v5 - 0)
            case "foo bar":
                var v48 float32 = (1e10 + 1.5)
                _ = v48
            case "":
                for fi49 := 0; fi49 < 20; fi49++ {
                    var v50 float32 = 0.0
                    _ = fi49
                    _ = v50
                }
            }
        case 480:
            if true {
                if (!true) {
                    fmt.Println(v2, v5, v1)
                } else {
                    if (!true) {
                        v51 := GenS6{F0: ((v1 + "foo bar") + "foo bar"), F1: (255 == 1)}
                        _ = v51
                    }
                    v4 += genFn1((0.1 < 1.5))
                    v4 -= (v2 ^ genFn1((0 == 938)))
                }
                v1 = (v1 + ((((v1 + "") + ("a" + "")) + v1) + v1))
            }
        }
    case 4294967295:
        switch ((v1 + "\n\t") + ("foo bar" + v1)) {
        case "\n\t":
            var v52 uint = 564
            _ = v52
        }
    default:
        for fi53 := 0; fi53 < 5; fi53++ {
            fmt.Println(v4, v5, v4)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi53
        }
    }
    fmt.Println(v1, v1)
    for fi54 := 0; fi54 < 14; fi54++ {
        v5 = 42
        _ = fi54
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        v5 += (^(v5))
        _ = fi55
    }
    switch v1 {
    case "foo bar":
        v56 := GenS7{F0: 2147483647, F1: -1, F2: 9223372036854775807}
        _ = v56
    case "":
        if true {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v57 := GenS10{F0: -1}
                _ = v57
            } else {
                for fi58 := 0; fi58 < 9; fi58++ {
                    v4 = (^(v2))
                    _ = fi58
                }
            }
            for fi59 := 0; fi59 < 6; fi59++ {
                for fi60 := 0; fi60 < 13; fi60++ {
                    if true {
                        v5 -= 1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(fi59)
                    }
                    var v61 uint32 = 1
                    _ = fi60
                    _ = v61
                }
                if ((true && true) || (false || (42 == 496))) {
                    switch (v2 + fi59) {
                    case 9223372036854775807:
                        for fi62 := 0; fi62 < 18; fi62++ {
                            var v63 int32 = 2147483647
                            switch (-(v2)) {
                            case -256:
                                _ = len(v3)
                                fmt.Println(v4)
                                var v64 int16 = 32767
                                _ = v64
                            default:
                                switch fi62 {
                                case 32767:
                                    fmt.Println(v4, v63)
                                default:
                                    if false {
                                        fmt.Println(v5)
                                        if (!false) {
                                            var v65 uint32 = 2
                                            if true {
                                                var v66 int = fi59
                                                v63 = v63
                                                var v67 uint16 = 255
                                                _ = v66
                                                _ = v67
                                            } else {
                                                fmt.Println(v65)
                                                if true {
                                                    if (true && false) {
                                                        fmt.Println(fi62, v1)
                                                        if (59 != uint16(v5)) {
                                                            v1 = v1
                                                            var v68 uint16 = 65535
                                                            if (true && true) {
                                                                if true {
                                                                    var v69 uint16 = v68
                                                                    var v70 int32 = v63
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                fmt.Println(v1, fi59, fi59)
                                                                v3 = append(v3, "\n\t")
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                v2 = v4
                                                            case "a":
                                                                v68 -= (v68 + v68)
                                                            }
                                                            _ = v68
                                                        } else {
                                                            if (127 <= 128) {
                                                                switch v2 {
                                                                case 68:
                                                                    v63 = v63
                                                                case 396:
                                                                    v2 = v4
                                                                case 127:
                                                                    v71 := GenS9{F0: 20, F1: (^(v5)), F2: -32768}
                                                                    _ = v71
                                                                }
                                                                fmt.Println(v65, v1, v5)
                                                                if ((^(v65)) == v65) {
                                                                    v3 = append(v3, "hello")
                                                                    if ((45 >= 0) && ((v65 % ((v65) | 1)) < (v65 | (v65 & 2147483648)))) {
                                                                        _ = len(v3)
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v72 := GenS11{F0: true, F1: 100}
                                                                            _ = v72
                                                                        default:
                                                                            v63 -= (-(v63))
                                                                        }
                                                                        var v73 uint16 = 65535
                                                                        _ = v73
                                                                    } else {
                                                                        _ = len(v3)
                                                                        if true {
                                                                            var v74 int = v4
                                                                            v75 := GenS8{F0: 0, F1: (^(v63)), F2: -1}
                                                                            _ = v74
                                                                            _ = v75
                                                                        } else {
                                                                            v76 := GenS12{F0: -325, F1: v1}
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                        v77 := GenS9{F0: 1, F1: (^(v5)), F2: -32768}
                                                                        _ = v77
                                                                    }
                                                                    v4 = (v2 >> ((484) & 63))
                                                                } else {
                                                                    v3 = append(v3, v1)
                                                                    v3 = append(v3, "\n\t")
                                                                    var v78 int = v2
                                                                    fmt.Println(v2, v2, v2)
                                                                    _ = v78
                                                                }
                                                            }
                                                            if (!true) {
                                                                if true {
                                                                    var v79 int16 = 1
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v1
                                                                    }
                                                                    _ = v79
                                                                }
                                                                v5 = (^(v5))
                                                                v80 := GenS13{F0: 0, F1: 42, F2: genFn1(((!(("" + "") >= (v1 + ((v1 + v1) + ((v1 + "hello") + "a"))))) || ((v63 | (v63 - (v63 | v63))) >= (v63 - v63))))}
                                                                _ = v80
                                                            } else {
                                                                fmt.Println(fi59, v65, v5)
                                                                v65 -= (^(v65))
                                                                v3 = append(v3, v1)
                                                            }
                                                        }
                                                        if ((!(!true)) || ((0 != 0) && (false && (-1 == -1)))) {
                                                            switch v1 {
                                                            case "foo bar":
                                                                v5 = v5
                                                            default:
                                                                if (0.1 < 0.1) {
                                                                    var v81 float64 = 0.0
                                                                    _ = v81
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = ""
                                                            }
                                                        } else {
                                                            v65 = v65
                                                        }
                                                    } else {
                                                        if true {
                                                            v65 += (^(v65))
                                                        }
                                                        if true {
                                                            v82 := GenS11{F0: (v63 != v63), F1: -46}
                                                            if (true || false) {
                                                                switch fi59 {
                                                                case 4294967295:
                                                                    _ = len(v3)
                                                                case 120:
                                                                    if (false || false) {
                                                                        v5 += (v5 % ((427) | 1))
                                                                        var v83 float64 = 0.0
                                                                        v84 := GenS12{F0: 32767, F1: ((v1 + ((v1 + "") + v1)) + ("\n\t" + v1))}
                                                                        _ = v83
                                                                        _ = v84
                                                                    } else {
                                                                        fmt.Println(fi59)
                                                                    }
                                                                case 32767:
                                                                    if true {
                                                                        var v85 string = ((v1 + v1) + v1)
                                                                        fmt.Println(fi62, fi62, v5)
                                                                        var v86 string = "foo bar"
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "foo bar"
                                                                    }
                                                                    switch (-(v4)) {
                                                                    case -233:
                                                                        var v87 bool = false
                                                                        _ = v87
                                                                    case 4294967295:
                                                                        if (!true) {
                                                                            v4 = (^(fi62))
                                                                            var v88 uint64 = 768
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                    var v89 float64 = (1.5 / 0.1)
                                                                    _ = v89
                                                                default:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v1
                                                                    }
                                                                }
                                                                v90 := GenS2{F0: "\n\t"}
                                                                _ = v90
                                                            } else {
                                                                if ((-2147483648 < 0) || true) {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v1
                                                                    }
                                                                } else {
                                                                    v5 = (v5 - v5)
                                                                    var v91 int8 = 51
                                                                    if false {
                                                                        var v92 int16 = -708
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v93 := GenS14{F0: 1}
                                                                        _ = v92
                                                                        _ = v93
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v94 uint64 = v5
                                                                        fmt.Println(v1)
                                                                        _ = v94
                                                                    }
                                                                    fmt.Println(v5)
                                                                    v1 = ((((((v1 + "hello") + v1) + "a") + ((v1 + v1) + v1)) + v1) + v1)
                                                                    if true {
                                                                        v91 -= (v91 + v91)
                                                                        fmt.Println(v5, v4, v65)
                                                                    } else {
                                                                        var v95 bool = false
                                                                        _ = v95
                                                                    }
                                                                    _ = v91
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = v1
                                                            }
                                                            _ = v82
                                                        }
                                                        var v96 uint = 521
                                                        _ = v96
                                                    }
                                                    if true {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v97 uint16 = 2
                                                        if true {
                                                            v97 += (v97 + (v97 | (v97 & v97)))
                                                            if (v63 > (v63 * 676)) {
                                                                v98 := GenS12{F0: 32767, F1: ((((v1 + "\n\t") + v1) + "a") + v1)}
                                                                v5 = 18446744073709551615
                                                                _ = v98
                                                            }
                                                            switch (fi59 / (((fi59 * genFn1((false && false)))) | 1)) {
                                                            case 1:
                                                                v3 = append(v3, v1)
                                                            case -972:
                                                                var v99 float64 = 1.5
                                                                _ = v99
                                                            }
                                                        } else {
                                                            if false {
                                                                v97 += v97
                                                                v65 = (v65 - 2147483647)
                                                                switch genFn1(true) {
                                                                case 32767:
                                                                    if true {
                                                                        v4 += 128
                                                                        v63 -= v63
                                                                        var v100 bool = (v1 == v1)
                                                                        _ = v100
                                                                    }
                                                                case 128:
                                                                    v3 = append(v3, (v1 + v1))
                                                                case 0:
                                                                    var v101 bool = (1000 > 1000)
                                                                    _ = v101
                                                                }
                                                            } else {
                                                                var v102 uint32 = (v65 + uint32(v5))
                                                                v102 -= v102
                                                                switch v2 {
                                                                case 32767:
                                                                    _ = len(v3)
                                                                case -1:
                                                                    if (!true) {
                                                                        v103 := GenS15{F0: 255}
                                                                        var v104 int32 = v63
                                                                        _ = v103
                                                                        _ = v104
                                                                    } else {
                                                                        var v105 uint32 = (v65 & (v102 + v102))
                                                                        _ = v105
                                                                    }
                                                                case 0:
                                                                    fmt.Println(v63)
                                                                default:
                                                                    if true {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    } else {
                                                                        v4 = (v2 ^ v4)
                                                                        var v106 float64 = 0.1
                                                                        if (126 >= 100) {
                                                                            var v107 float64 = v106
                                                                            if true {
                                                                                switch v4 {
                                                                                case 603:
                                                                                    var v108 int64 = 628
                                                                                    _ = v108
                                                                                case 641:
                                                                                    if true {
                                                                                        if true {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            fmt.Println(v1, v106, v2)
                                                                                            var v109 bool = false
                                                                                            _ = v109
                                                                                        }
                                                                                    } else {
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        v2 = v4
                                                                                        var v110 uint16 = (v97 ^ v97)
                                                                                        _ = v110
                                                                                    }
                                                                                }
                                                                            }
                                                                            v5 = v5
                                                                            _ = v107
                                                                        }
                                                                        _ = v106
                                                                    }
                                                                }
                                                                _ = v102
                                                            }
                                                            v97 -= v97
                                                            if (42 >= (v65 % ((v65) | 1))) {
                                                                var v111 float32 = 1.5
                                                                v65 += (v65 - 2)
                                                                var v112 int32 = (v63 / (((v63 / ((int32(v5)) | 1))) | 1))
                                                                if (v65 >= (v65 + 2)) {
                                                                    v113 := GenS13{F0: v97, F1: 2, F2: (v2 | fi59)}
                                                                    v112 += v112
                                                                    _ = v113
                                                                } else {
                                                                    fmt.Println(fi62)
                                                                }
                                                                v114 := GenS12{F0: 459, F1: (("\n\t" + (v1 + ("foo bar" + v1))) + ("a" + "foo bar"))}
                                                                _ = v111
                                                                _ = v112
                                                                _ = v114
                                                            } else {
                                                                if true {
                                                                    v3 = append(v3, ("" + ""))
                                                                    v115 := GenS16{F0: (v65 | v65), F1: v63}
                                                                    var v116 uint = 65535
                                                                    _ = v115
                                                                    _ = v116
                                                                }
                                                                switch genFn1(true) {
                                                                case 128:
                                                                    var v117 int64 = 0
                                                                    _ = v117
                                                                case 32768:
                                                                    if (false && false) {
                                                                        switch -858 {
                                                                        case 488:
                                                                            var v118 uint32 = 4294967295
                                                                            _ = v118
                                                                        case -1:
                                                                            if false {
                                                                                switch (v2 * int(v5)) {
                                                                                case 825:
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v63 = v63
                                                                                    }
                                                                                case 32768:
                                                                                    v119 := GenS7{F0: 2147483647, F1: v63, F2: 772}
                                                                                    var v120 float32 = 3.14159
                                                                                    switch ((v1 + "\n\t") + (v1 + "")) {
                                                                                    case "hello":
                                                                                        var v121 float64 = 0.0
                                                                                        _ = v121
                                                                                    case "\n\t":
                                                                                        if (v63 < v63) {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = len(v3)
                                                                                        } else {
                                                                                            var v122 float32 = (v120 * 3.14159)
                                                                                            _ = v122
                                                                                        }
                                                                                    }
                                                                                    v2 = genFn1((!(false && true)))
                                                                                    var v123 float64 = (1e10 / 0.1)
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                    _ = v123
                                                                                case -2147483648:
                                                                                    switch v1 {
                                                                                    case "":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    case "a":
                                                                                        if (-1 <= -9223372036854775808) {
                                                                                            v124 := GenS15{F0: 175}
                                                                                            if (false && false) {
                                                                                                fmt.Println(fi62, v63)
                                                                                                if (-9223372036854775808 > 622) {
                                                                                                    v97 += v97
                                                                                                    if ((1 >= 127) || ((100 == 19) || false)) {
                                                                                                        if true {
                                                                                                            v63 = 0
                                                                                                            fmt.Println(v63, fi59, v5)
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = ""
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v125 uint8 = 153
                                                                                                            fmt.Println(v125)
                                                                                                            _ = len(v3)
                                                                                                            if true {
                                                                                                                _ = len(v3)
                                                                                                                v3 = append(v3, v1)
                                                                                                                if (false || true) {
                                                                                                                    switch 32767 {
                                                                                                                    case 793:
                                                                                                                        v63 = -1
                                                                                                                    default:
                                                                                                                        _ = len(v3)
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        switch genFn1(("" >= (((v1 + v1) + v1) + v1))) {
                                                                                                                        case -766:
                                                                                                                            v63 = (v63 - (v63 >> ((676) & 63)))
                                                                                                                        case -2147483648:
                                                                                                                            var v126 float64 = ((((0.0 / 1.0) * (0.0 + float64(1e10))) - 1.5) + ((-2.25 * float64(-2.25)) * 0.0))
                                                                                                                            v5 += v5
                                                                                                                            var v127 int32 = (v63 % (((v63 + 1)) | 1))
                                                                                                                            v128 := GenS15{F0: v125}
                                                                                                                            v129 := GenS17{F0: (fi62 - genFn1((635 >= 592))), F1: 4294967295}
                                                                                                                            _ = v126
                                                                                                                            _ = v127
                                                                                                                            _ = v128
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v130 := GenS17{F0: fi62, F1: genFn0((v63 << ((1) & 63)), v65)}
                                                                                                                _ = v130
                                                                                                            }
                                                                                                            v131 := GenS2{F0: v1}
                                                                                                            _ = v125
                                                                                                            _ = v131
                                                                                                        }
                                                                                                        v132 := GenS8{F0: (v97 - 1), F1: v63, F2: v63}
                                                                                                        fmt.Println(fi59, v1)
                                                                                                        _ = v132
                                                                                                    } else {
                                                                                                        _ = len(v3)
                                                                                                        if ((96 == -128) || true) {
                                                                                                            var v133 int16 = 32767
                                                                                                            fmt.Println(fi59, v133, fi59)
                                                                                                            _ = v133
                                                                                                        }
                                                                                                        v3 = append(v3, "\n\t")
                                                                                                        v134 := GenS18{F0: int8(v4), F1: 65535}
                                                                                                        _ = v134
                                                                                                    }
                                                                                                } else {
                                                                                                    v4 -= int(v97)
                                                                                                    v97 += v97
                                                                                                }
                                                                                                _ = len(v3)
                                                                                            } else {
                                                                                                _ = len(v3)
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            _ = v124
                                                                                        }
                                                                                    }
                                                                                    switch v2 {
                                                                                    case 32768:
                                                                                        v2 += (fi62 / (((^(v4))) | 1))
                                                                                    case -1:
                                                                                        var v135 int32 = (v63 & 0)
                                                                                        _ = v135
                                                                                    case 2147483647:
                                                                                        v3 = append(v3, v1)
                                                                                    default:
                                                                                        var v136 int16 = 32767
                                                                                        _ = v136
                                                                                    }
                                                                                    v65 = uint32(v97)
                                                                                    switch fi62 {
                                                                                    case 1:
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    case -147:
                                                                                        v5 = 404
                                                                                    case -2147483648:
                                                                                        _ = len(v3)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v1 = v1
                                                                            }
                                                                        case 32768:
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v3 = append(v3, "hello")
                                                                            case "a":
                                                                                var v137 uint64 = 1
                                                                                _ = v137
                                                                            case "\n\t":
                                                                                fmt.Println(v1, v2)
                                                                            default:
                                                                                var v138 uint32 = v65
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                        switch (v1 + (v1 + "a")) {
                                                                        case "a":
                                                                            if false {
                                                                                switch (fi59 * fi62) {
                                                                                case 134:
                                                                                    if true {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ((("foo bar" + v1) + "\n\t") + "")
                                                                                        }
                                                                                        var v139 bool = false
                                                                                        var v140 uint32 = 42
                                                                                        _ = v139
                                                                                        _ = v140
                                                                                    } else {
                                                                                        v4 = v2
                                                                                        v4 = (fi59 & v2)
                                                                                        switch (v1 + "") {
                                                                                        case "\n\t":
                                                                                            v1 = v1
                                                                                        default:
                                                                                            v141 := GenS19{F0: (^(v63)), F1: (true || (true && true))}
                                                                                            _ = v141
                                                                                        }
                                                                                    }
                                                                                case 32767:
                                                                                    var v142 uint = uint(v5)
                                                                                    _ = v142
                                                                                case 128:
                                                                                    v63 = (v63 >> ((genFn0((v63 >> ((4294967295) & 63)), (v65 - 830))) & 63))
                                                                                default:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                }
                                                                                if ((v5 + (v5 - 18446744073709551615)) != v5) {
                                                                                    switch v1 {
                                                                                    case "hello":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ("hello" + ((v1 + v1) + (v1 + v1)))
                                                                                        }
                                                                                    case "":
                                                                                        v143 := GenS1{F0: 128, F1: 0}
                                                                                        _ = v143
                                                                                    }
                                                                                    v1 = ("" + v1)
                                                                                    var v144 float64 = 1.5
                                                                                    _ = v144
                                                                                } else {
                                                                                    if (0 < (fi62 * fi62)) {
                                                                                        if true {
                                                                                            _ = len(v3)
                                                                                        }
                                                                                    } else {
                                                                                        var v145 int = 127
                                                                                        fmt.Println(v2, v65, v65)
                                                                                        v146 := GenS12{F0: 1, F1: (v1 + ("a" + v1))}
                                                                                        v147 := GenS20{F0: fi62}
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                        _ = v147
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v148 float64 = (0.0 / 1.0)
                                                                                _ = v148
                                                                            }
                                                                        }
                                                                    }
                                                                case -9223372036854775808:
                                                                    if false {
                                                                        v2 -= (v2 << ((65535) & 63))
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v97
                                                    }
                                                }
                                                if (!(true && (!false))) {
                                                    v149 := GenS17{F0: (fi62 | genFn1(true)), F1: genFn0(728, 650)}
                                                    v5 += (v5 >> ((262) & 63))
                                                    v150 := GenS21{F0: true}
                                                    v2 = (^(fi62))
                                                    _ = v149
                                                    _ = v150
                                                } else {
                                                    fmt.Println(fi62, fi62, fi62)
                                                    var v151 int16 = 170
                                                    var v152 string = v1
                                                    _ = v151
                                                    _ = v152
                                                }
                                                v1 = ((("a" + v1) + v1) + "foo bar")
                                            }
                                            _ = v65
                                        } else {
                                            var v153 uint8 = 0
                                            v3 = append(v3, ("" + v1))
                                            if (false && (!(18446744073709551615 < 960))) {
                                                var v154 uint8 = (v153 ^ (^(v153)))
                                                var v155 int16 = 42
                                                _ = v154
                                                _ = v155
                                            } else {
                                                v156 := GenS22{F0: 184}
                                                switch v1 {
                                                case "foo bar":
                                                    var v157 string = v1
                                                    _ = v157
                                                case "hello":
                                                    switch (v2 / (((-(fi62))) | 1)) {
                                                    case 128:
                                                        v2 = (fi59 % ((9223372036854775807) | 1))
                                                    case -597:
                                                        var v158 bool = ((1.5 > ((0.0 * -2.25) + 1.5)) && (false && false))
                                                        _ = v158
                                                    }
                                                default:
                                                    switch (v1 + (v1 + ((v1 + v1) + v1))) {
                                                    case "hello":
                                                        v4 = v4
                                                    case "a":
                                                        var v159 uint16 = 32767
                                                        _ = v159
                                                    case "\n\t":
                                                        if true {
                                                            if false {
                                                                var v160 string = "\n\t"
                                                                v161 := GenS23{F0: 81, F1: v63, F2: -113}
                                                                _ = v160
                                                                _ = v161
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = "foo bar"
                                                                }
                                                                fmt.Println(v153, v2)
                                                                fmt.Println(v63, fi59, v63)
                                                                var v162 uint64 = v5
                                                                _ = v162
                                                            }
                                                        } else {
                                                            if ((-(v63)) <= 0) {
                                                                switch v2 {
                                                                case -1:
                                                                    _ = len(v3)
                                                                case 42:
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v1 + "")
                                                                        }
                                                                        if false {
                                                                            _ = len(v3)
                                                                        }
                                                                        _ = len(v3)
                                                                        v1 = v1
                                                                    case "hello":
                                                                        v3 = append(v3, v1)
                                                                    case "":
                                                                        fmt.Println(fi62)
                                                                    default:
                                                                        var v163 int = (fi62 - (fi62 << ((genFn0(v63, 2)) & 63)))
                                                                        _ = v163
                                                                    }
                                                                default:
                                                                    var v164 int16 = 32767
                                                                    _ = v164
                                                                }
                                                                var v165 string = v1
                                                                _ = v165
                                                            } else {
                                                                var v166 int8 = -1
                                                                v63 -= v63
                                                                _ = v166
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v156
                                            }
                                            _ = v153
                                        }
                                    } else {
                                        fmt.Println(v2, v4, fi59)
                                        if true {
                                            v3 = append(v3, (((v1 + v1) + (v1 + ((v1 + v1) + "hello"))) + v1))
                                            switch v2 {
                                            case -1:
                                                v5 = (v5 % ((2) | 1))
                                            case 128:
                                                var v167 int = genFn1((!false))
                                                _ = v167
                                            case 42:
                                                var v168 int = (v2 - int(v5))
                                                _ = v168
                                            default:
                                                fmt.Println(v1, fi59)
                                                v4 += (v4 * (fi59 & (v2 | v2)))
                                                var v169 uint64 = (^(v5))
                                                _ = v169
                                            }
                                            fmt.Println(fi59, fi59, v4)
                                        } else {
                                            var v170 uint16 = 32767
                                            v3 = append(v3, v1)
                                            var v171 int16 = int16(v5)
                                            v170 = (v170 * v170)
                                            _ = v170
                                            _ = v171
                                        }
                                        if false {
                                            var v172 int32 = 1
                                            var v173 uint = 1
                                            _ = v172
                                            _ = v173
                                        } else {
                                            v2 -= (^(fi59))
                                        }
                                    }
                                }
                            }
                            v1 = ((v1 + (v1 + v1)) + ("\n\t" + v1))
                            _ = fi62
                            _ = v63
                        }
                    case 16:
                        switch v2 {
                        case 32767:
                            fmt.Println(v5, v5, v4)
                        case 1:
                            var v174 int64 = -917
                            _ = v174
                        }
                    case -639:
                        switch v4 {
                        case -2147483648:
                            v175 := GenS11{F0: true, F1: int8(v2)}
                            _ = v175
                        }
                    default:
                        var v176 uint64 = (v5 ^ (v5 * 119))
                        _ = v176
                    }
                    var v177 uint8 = 127
                    for fi178 := 0; fi178 < 12; fi178++ {
                        if true {
                            v179 := GenS9{F0: -488, F1: 4294967295, F2: -1}
                            v3 = append(v3, v1)
                            v2 += v2
                            _ = v179
                        } else {
                            switch (fi178 + v4) {
                            case 32768:
                                var v180 int16 = 756
                                _ = v180
                            }
                        }
                        v5 -= (^(v5))
                        _ = fi178
                    }
                    _ = v177
                } else {
                    switch ((v1 + v1) + v1) {
                    case "foo bar":
                        switch "foo bar" {
                        case "foo bar":
                            for fi181 := 0; fi181 < 11; fi181++ {
                                var v182 int16 = -610
                                _ = fi181
                                _ = v182
                            }
                        }
                    case "a":
                        _ = len(v3)
                    case "\n\t":
                        if false {
                            var v183 uint64 = (v5 * v5)
                            _ = len(v3)
                            var v184 string = v1
                            _ = v183
                            _ = v184
                        }
                    }
                    if false {
                        v5 += (v5 >> ((240) & 63))
                        switch v1 {
                        case "\n\t":
                            if len(v3) > 0 {
                                v3[0] = (v1 + v1)
                            }
                        case "hello":
                            if true {
                                v5 = (v5 / ((v5) | 1))
                                fmt.Println(v2)
                            }
                        case "a":
                            if (65535 > 42) {
                                var v185 uint = genFn0(842, 1)
                                _ = v185
                            } else {
                                var v186 int32 = 1
                                _ = v186
                            }
                        }
                        _ = len(v3)
                    }
                }
                var v187 float64 = -2.25
                _ = fi59
                _ = v187
            }
            v1 = "a"
        }
    case "\n\t":
        v2 -= 42
    default:
        if (-1 != -328) {
            var v188 uint32 = 2147483648
            if (444 <= 18446744073709551615) {
                switch (v1 + v1) {
                case "a":
                    for fi189 := 0; fi189 < 7; fi189++ {
                        switch "foo bar" {
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                        case "":
                            switch ("\n\t" + v1) {
                            case "a":
                                v190 := GenS10{F0: 32767}
                                _ = v190
                            case "\n\t":
                                v3 = append(v3, "foo bar")
                            }
                        default:
                            var v191 float32 = 0.0
                            _ = v191
                        }
                        v192 := GenS24{F0: 1}
                        var v193 uint16 = 1
                        _ = fi189
                        _ = v192
                        _ = v193
                    }
                default:
                    if (!true) {
                        _ = len(v3)
                        if len(v3) > 0 {
                            v3[0] = v1
                        }
                    } else {
                        _ = len(v3)
                    }
                }
                v3 = append(v3, ((v1 + "\n\t") + v1))
                v194 := GenS14{F0: 255}
                _ = v194
            } else {
                fmt.Println(v4, v4)
                for fi195 := 0; fi195 < 12; fi195++ {
                    v188 = v188
                    v3 = append(v3, v1)
                    v196 := GenS17{F0: v4, F1: 99}
                    _ = fi195
                    _ = v196
                }
                var v197 uint64 = (v5 | v5)
                _ = v197
            }
            _ = v188
        } else {
            for fi198 := 0; fi198 < 17; fi198++ {
                if len(v3) > 0 {
                    v3[0] = (v1 + v1)
                }
                var v199 uint16 = 42
                _ = fi198
                _ = v199
            }
            switch "a" {
            case "\n\t":
                v2 += (v4 + v4)
            }
        }
    }
    if false {
        if ("foo bar" == "hello") {
            var v200 bool = false
            _ = v200
        }
    } else {
        switch v2 {
        case 1:
            v2 = v4
        case -327:
            v2 -= v2
        case -93:
            v201 := GenS25{F0: 65535}
            _ = v201
        default:
            var v202 float32 = 0.1
            _ = v202
        }
        switch (v4 | genFn1(false)) {
        case 1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 2147483647:
            v5 += (v5 / (((^(v5))) | 1))
        default:
            v5 += v5
        }
    }
    for fi203 := 0; fi203 < 12; fi203++ {
        if len(v3) > 0 {
            v3[0] = v1
        }
        v204 := GenS7{F0: 0, F1: 2147483647, F2: 42}
        var v205 int32 = 7
        _ = fi203
        _ = v204
        _ = v205
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
