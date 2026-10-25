package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 bool }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 bool; F1 uint8; F2 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 uint }
func genFn0() uint8 {
    if false {
        var v6 uint = 396
        _ = v6
    }
    var v7 int64 = 565
    if (!(true && true)) {
        var v8 int32 = -264
        fmt.Println(v8)
        _ = v8
    } else {
        switch 42 {
        case 4294967295:
            v9 := GenS1{F0: (^(v7)), F1: true, F2: ((!(false || false)) || (!(65535 <= 65535)))}
            _ = v9
        case 1:
            v7 += v7
        case -423:
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS1{F0: -9223372036854775808, F1: (-32768 != 42), F2: (false && false)}
                _ = fi10
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 9223372036854775807, F1: false, F2: false}
                v14 := []int{(fi12 >> ((0) & 63))}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        var v15 float64 = 0.1
        var v16 string = "hello"
        _ = v15
        _ = v16
    }
    switch "foo bar" {
    case "foo bar":
        var v17 int8 = -57
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := []int{fi18, fi18}
            var v20 uint = 18446744073709551615
            for fi21 := 0; fi21 < 8; fi21++ {
                v7 = (v7 ^ (^(v7)))
                switch (fi21 + 9223372036854775807) {
                case 684:
                    var v22 int8 = -115
                    _ = v22
                }
                if false {
                    var v23 int32 = 2147483647
                    v23 -= int32(v20)
                    _ = v23
                }
                _ = fi21
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v24 := GenS2{F0: 263, F1: 42}
    _ = v24
    _ = v7
    return 0
}
func genFn1() int {
    var v25 bool = true
    var v26 float32 = 0.1
    _ = v25
    _ = v26
    return -932
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"\n\t"}
    var v3 uint16 = 0
    v4 := make(map[uint32]string)
    var v5 int16 = 0
    switch "hello" {
    case "":
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 int32 = 2147483647
            if (v1 >= 4294967295) {
                var v29 uint16 = v3
                _ = v29
            } else {
                var v30 float32 = 0.0
                v5 = v5
                _ = v30
            }
            switch -1 {
            case 207:
                for fi31 := 0; fi31 < 20; fi31++ {
                    if (3.14159 == (3.14159 - 0.1)) {
                        v4[2147483647] = ("foo bar" + ("hello" + ("hello" + "hello")))
                        v28 += (v28 % (((v28 + v28)) | 1))
                    } else {
                        v3 += v3
                    }
                    var v32 string = (("hello" + "hello") + "hello")
                    var v33 uint8 = 127
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            default:
                v4[42] = "hello"
            }
            _ = fi27
            _ = v28
        }
    case "\n\t":
        var v34 uint64 = (v1 | v1)
        _ = v34
    case "foo bar":
        if (!((v5 / ((1) | 1)) > (^(v5)))) {
            if false {
                v35 := GenS1{F0: -974, F1: (!false), F2: (((0.1 + 1e10) * (((3.14159 / 1.5) * (-2.25 + (float32((float64(float32(1.5)) - ((float64(float32((1.5 * float64(0.0)))) + 1.5) + 1e10))) + 0.0))) * float32(1.5))) != (float32((3.14159 - float64(-2.25))) / 3.14159))}
                fmt.Println(v1)
                _ = v35
            }
        }
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        if (v5 != (v5 >> ((65535) & 63))) {
            for fi37 := 0; fi37 < 6; fi37++ {
                fmt.Println(v5)
                v38 := GenS3{F0: v5, F1: 4294967295, F2: (v3 - (v3 ^ 2))}
                _ = fi37
                _ = v38
            }
            switch genFn1() {
            case 9223372036854775807:
                if true {
                    _ = len(v2)
                    switch "" {
                    case "a":
                        for fi39 := 0; fi39 < 14; fi39++ {
                            v40 := GenS1{F0: -447, F1: ((("a" + "hello") + "foo bar") >= ""), F2: (("a" + ("a" + ("hello" + ((("\n\t" + "a") + "a") + "")))) != ((("a" + "a") + ("hello" + "hello")) + "a"))}
                            fmt.Println(v5, v1, v3)
                            fmt.Println(fi39)
                            _ = fi39
                            _ = v40
                        }
                    case "\n\t":
                        var v41 float32 = 1.5
                        _ = v41
                    }
                    v42 := GenS3{F0: v5, F1: 9223372036854775807, F2: (^(v3))}
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 2; fi43++ {
                        switch ("foo bar" + "") {
                        case "a":
                            if true {
                                var v44 float32 = 1e10
                                v44 += float32(1.5)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v44
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case "\n\t":
                            v3 = v3
                        case "":
                            fmt.Println(v1)
                        default:
                            fmt.Println(fi36, v5, v5)
                            fmt.Println(fi43, fi43)
                            switch "hello" {
                            case "hello":
                                if false {
                                    v45 := GenS1{F0: 0, F1: true, F2: (255 > 128)}
                                    var v46 int8 = 100
                                    v47 := GenS2{F0: 850, F1: 0}
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                } else {
                                    var v48 string = "hello"
                                    switch (fi43 % ((fi43) | 1)) {
                                    case -1:
                                        if ((^(v1)) == (v1 >> ((414) & 63))) {
                                            var v49 uint = 255
                                            _ = v49
                                        }
                                    case 128:
                                        v5 += (^(v5))
                                    }
                                    switch v48 {
                                    case "hello":
                                        _ = len(v2)
                                    default:
                                        if (255 < 0) {
                                            var v50 int32 = 42
                                            fmt.Println(v48, fi36)
                                            _ = v50
                                        }
                                    }
                                    _ = v48
                                }
                            case "":
                                switch (("foo bar" + "") + ("\n\t" + "foo bar")) {
                                case "":
                                    v51 := GenS1{F0: 548, F1: false, F2: true}
                                    _ = v51
                                case "foo bar":
                                    if (42 > 32) {
                                        switch (fi36 - fi43) {
                                        case -2147483648:
                                            var v52 uint8 = 255
                                            _ = v52
                                        case 0:
                                            v3 = (v3 >> ((18446744073709551615) & 63))
                                        }
                                    }
                                default:
                                    var v53 uint64 = v1
                                    _ = v53
                                }
                            }
                        }
                        var v54 bool = true
                        v55 := GenS4{F0: 1, F1: 1, F2: 828}
                        _ = fi43
                        _ = v54
                        _ = v55
                    }
                    switch fi36 {
                    case 42:
                        v56 := GenS4{F0: 605, F1: (^(v1)), F2: 1}
                        _ = v56
                    case 498:
                        var v57 uint64 = (v1 >> ((42) & 63))
                        _ = v57
                    case -597:
                        if (1 == 729) {
                            v5 = int16(v3)
                            if false {
                                for fi58 := 0; fi58 < 19; fi58++ {
                                    v59 := GenS5{F0: -128}
                                    if (false && (false && true)) {
                                        v60 := GenS4{F0: 2147483647, F1: (v1 % ((v1) | 1)), F2: 410}
                                        _ = v60
                                    }
                                    fmt.Println(fi36)
                                    _ = fi58
                                    _ = v59
                                }
                                switch 366 {
                                case -9223372036854775808:
                                    switch ("\n\t" + (("hello" + "\n\t") + "foo bar")) {
                                    case "\n\t":
                                        var v61 uint = 4294967295
                                        _ = v61
                                    case "a":
                                        for fi62 := 0; fi62 < 8; fi62++ {
                                            fmt.Println(v3)
                                            switch ((("a" + ("a" + ("a" + "hello"))) + ("" + "foo bar")) + (("\n\t" + ("foo bar" + "\n\t")) + ("hello" + ("foo bar" + ("" + "a"))))) {
                                            case "foo bar":
                                                fmt.Println(v3)
                                            case "\n\t":
                                                _ = len(v2)
                                            default:
                                                var v63 int8 = 0
                                                _ = v63
                                            }
                                            var v64 uint32 = 0
                                            _ = fi62
                                            _ = v64
                                        }
                                    case "hello":
                                        v5 -= v5
                                    default:
                                        fmt.Println(v3, fi36, v5)
                                    }
                                case 42:
                                    for fi65 := 0; fi65 < 7; fi65++ {
                                        switch "\n\t" {
                                        case "hello":
                                            v5 = (v5 << ((42) & 63))
                                        case "a":
                                            var v66 int64 = -985
                                            _ = v66
                                        default:
                                            var v67 uint64 = v1
                                            _ = v67
                                        }
                                        _ = fi65
                                    }
                                }
                            } else {
                                v68 := GenS3{F0: v5, F1: int(v3), F2: 0}
                                var v69 uint8 = 72
                                _ = v68
                                _ = v69
                            }
                        } else {
                            v70 := GenS6{F0: (-(fi36)), F1: (v3 - v3)}
                            var v71 float64 = -2.25
                            switch ("foo bar" + ("hello" + "a")) {
                            case "foo bar":
                                switch "foo bar" {
                                case "hello":
                                    v72 := GenS7{F0: genFn1()}
                                    _ = v72
                                case "foo bar":
                                    v1 -= (v1 & (^(v1)))
                                case "":
                                    delete(v4, 2147483648)
                                }
                            case "":
                                v71 = v71
                            case "\n\t":
                                if ((12 != -128) || ((v3 - 65535) < v3)) {
                                    v73 := GenS8{F0: true, F1: 18446744073709551615}
                                    _ = v73
                                } else {
                                    var v74 string = "\n\t"
                                    if true {
                                        var v75 uint32 = 4294967295
                                        fmt.Println(v5, v1, v75)
                                        _ = v75
                                    } else {
                                        v1 += 4294967295
                                        v1 = 42
                                    }
                                    for fi76 := 0; fi76 < 2; fi76++ {
                                        fmt.Println(fi36, v5)
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v77 int32 = 1
                                            _ = v77
                                        default:
                                            v5 = 1
                                        }
                                        _ = fi76
                                    }
                                    _ = v74
                                }
                            default:
                                if false {
                                    v78 := GenS1{F0: 1, F1: false, F2: ((v1 % (((v1 >> ((65535) & 63))) | 1)) > v1)}
                                    var v79 int16 = v5
                                    var v80 int64 = -1
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                            }
                            _ = v70
                            _ = v71
                        }
                    }
                    for fi81 := 0; fi81 < 10; fi81++ {
                        var v82 bool = false
                        _ = fi81
                        _ = v82
                    }
                }
            case 217:
                if (0.0 != 1.5) {
                    fmt.Println(fi36, v1, fi36)
                    fmt.Println(v1, v1)
                } else {
                    v3 += 65535
                    var v83 string = ""
                    var v84 bool = ((v3 != v3) || (!false))
                    _ = v83
                    _ = v84
                }
            case -100:
                var v85 uint32 = 2147483648
                _ = v85
            default:
                var v86 uint32 = 65535
                _ = v86
            }
            var v87 float32 = 3.14159
            _ = v87
        } else {
            for fi88 := 0; fi88 < 10; fi88++ {
                var v89 int = (fi36 % ((fi88) | 1))
                _ = fi88
                _ = v89
            }
            var v90 uint64 = (v1 >> ((42) & 63))
            if true {
                var v91 bool = false
                _ = v91
            } else {
                if ("" == ("a" + "foo bar")) {
                    var v92 int16 = (-(v5))
                    _ = v92
                } else {
                    for fi93 := 0; fi93 < 5; fi93++ {
                        var v94 float32 = 0.1
                        var v95 int64 = 9223372036854775807
                        var v96 int32 = 2147483647
                        _ = fi93
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                    v4[42] = (("" + "hello") + ("foo bar" + "hello"))
                }
                v97 := GenS9{F0: true, F1: 128, F2: true}
                _ = v97
            }
            _ = v90
        }
        if (0.1 < (3.14159 + 0.1)) {
            if (genFn0() != 127) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v98 := GenS2{F0: 624, F1: 2}
                _ = v98
            } else {
                var v99 int = (fi36 << ((255) & 63))
                v1 += (^(v1))
                _ = v99
            }
            v100 := GenS9{F0: (!true), F1: 127, F2: (1 == 100)}
            _ = v100
        } else {
            if (((true && ((3.14159 / 1.0) < 0.0)) && false) && true) {
                var v101 int = (^(fi36))
                v3 -= (v3 % ((v3) | 1))
                _ = v101
            } else {
                for fi102 := 0; fi102 < 4; fi102++ {
                    var v103 uint16 = (v3 >> ((649) & 63))
                    var v104 int64 = 610
                    switch 9223372036854775807 {
                    case -230:
                        var v105 int8 = -1
                        _ = v105
                    case 1:
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    }
                    var v106 uint64 = (v1 & (^(v1)))
                    _ = fi102
                    _ = v103
                    _ = v104
                    _ = v106
                }
                _ = len(v4)
                var v107 bool = (v3 >= (v3 / ((42) | 1)))
                _ = v107
            }
        }
        v108 := GenS4{F0: -790, F1: 42, F2: 1}
        _ = fi36
        _ = v108
    }
    v109 := v4[2147483648]
    var v110 uint = 18446744073709551615
    for fi111 := 0; fi111 < 8; fi111++ {
        for fi112 := 0; fi112 < 18; fi112++ {
            switch fi112 {
            case 880:
                v113 := GenS9{F0: ((1.5 - (3.14159 * (3.14159 / (-2.25 / -2.25)))) != float32((1e10 / 1.5))), F1: 255, F2: (v109 > (("" + "\n\t") + v109))}
                _ = v113
            case 32768:
                var v114 float64 = (1.5 * -2.25)
                _ = v114
            case -2147483648:
                v115 := GenS10{F0: (v109 + "\n\t")}
                _ = v115
            default:
                switch genFn1() {
                case 685:
                    var v116 int = -1
                    _ = v116
                case 4294967295:
                    v117 := GenS11{F0: (^(fi111)), F1: (v110 - v110)}
                    _ = v117
                default:
                    var v118 float64 = 1.5
                    _ = v118
                }
            }
            var v119 uint8 = genFn0()
            v3 -= v3
            _ = fi112
            _ = v119
        }
        _ = fi111
    }
    v2 = append(v2, "foo bar")
    for fi120 := 0; fi120 < 7; fi120++ {
        v5 = 0
        for fi121 := 0; fi121 < 10; fi121++ {
            var v122 int16 = v5
            _ = fi121
            _ = v122
        }
        if false {
            var v123 int16 = v5
            if (!(false || (!false))) {
                var v124 uint64 = (v1 ^ v1)
                fmt.Println(v3, v110, fi120)
                _ = v124
            } else {
                v125 := GenS7{F0: (fi120 + 2147483647)}
                v3 -= 65535
                _ = v125
            }
            var v126 uint16 = 1
            _ = v123
            _ = v126
        }
        _ = fi120
    }
    var v127 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v127:", v127)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v109
    _ = v110
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
