package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 int8; F2 int64 }
type GenS8 struct { F0 int8; F1 uint; F2 int64 }
type GenS9 struct { F0 int; F1 int32; F2 uint }
type GenS10 struct { F0 bool; F1 string }
func genFn0(p0 int, p1 string) int32 {
    fmt.Println(p1, p1)
    var v6 bool = true
    v7 := GenS1{F0: (p0 | p0), F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int8) uint8 {
    var v8 string = ("\n\t" + "hello")
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            p0 = 100
            if (true && true) {
                if (!(false || (1.5 >= 1.5))) {
                    v11 := GenS1{F0: (-(fi10)), F1: 991}
                    if true {
                        v8 = v8
                        v12 := GenS1{F0: fi10, F1: 919}
                        v8 = "\n\t"
                        switch fi9 {
                        case 127:
                            var v13 float64 = (3.14159 + -2.25)
                            _ = v13
                        case -335:
                            if (3.14159 == (0.0 - 1.5)) {
                                v14 := []int{(fi10 - (fi10 >> ((171) & 63)))}
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                var v15 uint32 = 65535
                                _ = v14
                                _ = v15
                            }
                        default:
                            fmt.Println(v8, p0, v8)
                        }
                        _ = v12
                    }
                    v16 := GenS2{F0: (!true)}
                    _ = v11
                    _ = v16
                } else {
                    v17 := GenS1{F0: fi9, F1: 4294967295}
                    switch (fi9 + 32768) {
                    case 4294967295:
                        if false {
                            v8 = v8
                            var v18 float64 = (3.14159 / float64(float32(1e10)))
                            _ = v18
                        }
                    case 1:
                        v8 = v8
                    case 2147483647:
                        if (!(65535 != 1)) {
                            v8 = "hello"
                            p0 = (-(p0))
                        }
                    default:
                        var v19 string = v8
                        _ = v19
                    }
                    _ = v17
                }
                fmt.Println(fi9, v8)
                var v20 bool = (true && false)
                _ = v20
            } else {
                v21 := GenS2{F0: false}
                _ = v21
            }
            if true {
                if false {
                    v22 := []int{(fi9 | -823), (fi9 << ((42) & 63)), (fi9 % (((fi10 ^ fi9)) | 1))}
                    _ = v22
                } else {
                    v23 := GenS3{F0: 42, F1: v8}
                    v8 = (v8 + ((v8 + ((((((v8 + (v8 + ((v8 + ("" + (v8 + (((v8 + (v8 + v8)) + v8) + v8)))) + v8))) + ((("" + v8) + ("a" + "")) + v8)) + ((v8 + v8) + v8)) + v8) + (v8 + v8)) + (("\n\t" + v8) + "foo bar"))) + ("" + v8)))
                    var v24 float32 = float32(-2.25)
                    _ = v23
                    _ = v24
                }
                p0 += p0
            }
            _ = fi10
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := []int{fi25, (fi25 & fi25), 32768}
            switch v8 {
            case "a":
                var v27 int32 = -119
                _ = v27
            default:
                v8 = v8
            }
            v8 = (v8 + v8)
            _ = fi25
            _ = v26
        }
        v28 := GenS1{F0: fi9, F1: 42}
        _ = fi9
        _ = v28
    }
    v29 := GenS2{F0: false}
    _ = p0
    _ = v29
    _ = v8
    return 230
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 2
    v3 := make(map[uint32]bool)
    var v4 int = 32768
    var v5 int16 = 42
    switch (v4 << ((4294967295) & 63)) {
    case -765:
        fmt.Println(v2)
    case 565:
        if false {
            v2 -= v2
        } else {
            var v30 int64 = 42
            v2 = 2147483648
            fmt.Println(v5)
            _ = v30
        }
    case 1:
        v4 += (v4 | (v4 >> ((65535) & 63)))
    }
    if false {
        var v31 string = ""
        _ = v31
    }
    v2 = (v2 * v2)
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 string = "hello"
        v2 = v2
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint16 = 0
            switch "\n\t" {
            case "a":
                v36 := GenS4{F0: 344}
                _ = v36
            default:
                v37 := GenS1{F0: 112, F1: 4294967295}
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    v2 += 2147483648
    var v38 uint32 = v2
    fmt.Println(v38, v2, v38)
    v38 = (v38 | v38)
    switch "foo bar" {
    case "":
        if (true && (721 <= 1)) {
            fmt.Println(v4, v38, v5)
        } else {
            if (((3.14159 * 0.0) * (-2.25 + 0.0)) != ((((-2.25 * (1.5 - 0.0)) * 3.14159) * ((float32((1.5 - 0.0)) * 0.1) * (((1.5 - 3.14159) * float32(-2.25)) / 1.5))) + 1.5)) {
                if (-676 < v4) {
                    delete(v3, 687)
                    var v39 int64 = 9223372036854775807
                    v2 = 42
                    _ = v39
                }
            }
            v4 -= (v4 + v4)
        }
    default:
        var v40 string = "\n\t"
        _ = v40
    }
    var v41 uint8 = genFn1(0)
    switch (^(v4)) {
    case 127:
        var v42 bool = (!(!((true && (!false)) && (1 != 1))))
        _ = v42
    case 128:
        var v43 int16 = v5
        _ = v43
    default:
        v44 := v1[""]
        _ = v44
    }
    v41 = (v41 >> ((281) & 63))
    if (false && false) {
        switch ("foo bar" + ("a" + "a")) {
        case "foo bar":
            if true {
                var v45 int = v4
                _ = v45
            }
        case "a":
            for fi46 := 0; fi46 < 14; fi46++ {
                v2 += (v2 & (^(v2)))
                var v47 uint32 = 42
                _ = fi46
                _ = v47
            }
        case "":
            v41 += (v41 & v41)
        }
    } else {
        for fi48 := 0; fi48 < 12; fi48++ {
            switch 1 {
            case -2147483648:
                v49 := []int{-231}
                _ = v49
            case 48:
                var v50 float32 = float32(1.5)
                _ = v50
            case 1:
                var v51 uint64 = 4294967295
                _ = v51
            }
            _ = fi48
        }
        for fi52 := 0; fi52 < 20; fi52++ {
            if false {
                if false {
                    v53 := GenS5{F0: v2}
                    v2 = v2
                    var v54 int64 = 9223372036854775807
                    _ = v53
                    _ = v54
                }
                v4 += -873
            } else {
                fmt.Println(v4)
            }
            if true {
                v55 := []int{-1, (^(fi52))}
                for fi56 := 0; fi56 < 19; fi56++ {
                    var v57 int16 = (v5 >> ((2) & 63))
                    var v58 int = (^(v4))
                    delete(v3, v38)
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                var v59 int8 = -1
                _ = v55
                _ = v59
            }
            _ = fi52
        }
        switch v4 {
        case -731:
            fmt.Println(v5)
        case 4294967295:
            switch v4 {
            case 1:
                for fi60 := 0; fi60 < 18; fi60++ {
                    if false {
                        v41 = 161
                        for fi61 := 0; fi61 < 1; fi61++ {
                            if (("" + ("foo bar" + ("hello" + "\n\t"))) < "hello") {
                                v62 := []int{fi61, (fi60 & v4), fi61, v4}
                                var v63 uint = 492
                                if false {
                                    var v64 int64 = -9223372036854775808
                                    if true {
                                        v5 += v5
                                    }
                                    _ = v64
                                }
                                _ = v62
                                _ = v63
                            } else {
                                if (!true) {
                                    var v65 uint32 = v38
                                    fmt.Println(v2, v5, fi61)
                                    _ = v65
                                }
                                v66 := GenS1{F0: v4, F1: 0}
                                v4 += fi60
                                _ = v66
                            }
                            v67 := GenS2{F0: (1.5 < 0.0)}
                            _ = fi61
                            _ = v67
                        }
                        var v68 int8 = 0
                        _ = v68
                    } else {
                        for fi69 := 0; fi69 < 20; fi69++ {
                            v38 += (v38 << ((255) & 63))
                            var v70 uint8 = 255
                            if false {
                                v38 -= (^(v2))
                                var v71 uint16 = 42
                                fmt.Println(v5, fi69)
                                _ = v71
                            } else {
                                switch -9223372036854775808 {
                                case 320:
                                    var v72 int = v4
                                    v70 = 42
                                    v73 := GenS6{F0: ("\n\t" + "")}
                                    _ = v72
                                    _ = v73
                                case 32767:
                                    v74 := []int{}
                                    _ = v74
                                case 127:
                                    v5 = int16(v2)
                                }
                                v75 := v3[(^(v2))]
                                switch ("\n\t" + ((("a" + "hello") + ("foo bar" + "\n\t")) + "")) {
                                case "a":
                                    if v75 {
                                        v76 := []int{v4, (v4 ^ fi69), 0, v4}
                                        v77 := GenS6{F0: (("" + "hello") + "\n\t")}
                                        _ = v76
                                        _ = v77
                                    }
                                case "hello":
                                    v78 := []int{32767, fi60}
                                    _ = v78
                                case "\n\t":
                                    fmt.Println(v38, fi60, fi60)
                                }
                                _ = v75
                            }
                            var v79 uint64 = 0
                            var v80 int32 = genFn0(0, "")
                            _ = fi69
                            _ = v70
                            _ = v79
                            _ = v80
                        }
                    }
                    _ = fi60
                }
            case 4294967295:
                var v81 string = "foo bar"
                _ = v81
            }
        case -979:
            var v82 int8 = 0
            _ = v82
        }
    }
    var v83 uint = 150
    switch v4 {
    case -2147483648:
        v83 = (v83 / ((v83) | 1))
    case 2147483647:
        var v84 uint32 = (v2 - v2)
        _ = v84
    case 693:
        v2 = (^(v38))
    default:
        if ((v38 - v2) < v2) {
            for fi85 := 0; fi85 < 20; fi85++ {
                if true {
                    for fi86 := 0; fi86 < 7; fi86++ {
                        switch "hello" {
                        case "foo bar":
                            v87 := GenS3{F0: genFn1(-106), F1: "a"}
                            _ = v87
                        case "a":
                            v88 := []int{}
                            _ = v88
                        case "\n\t":
                            if (!false) {
                                if (42 <= 0) {
                                    v4 = -1
                                    if (true && false) {
                                        _ = len(v3)
                                        if true {
                                            var v89 float32 = 1.5
                                            _ = v89
                                        }
                                    }
                                    switch v4 {
                                    case 403:
                                        v4 -= (^(v4))
                                    case 32768:
                                        if (false || (87 < 9223372036854775807)) {
                                            v2 += (^(v38))
                                            v90 := []int{}
                                            _ = v90
                                        } else {
                                            if (1 > 1) {
                                                v41 = 54
                                                v5 = v5
                                            } else {
                                                var v91 uint8 = (v41 ^ v41)
                                                var v92 uint32 = (v38 * v2)
                                                if ((v5 % ((-1) | 1)) < v5) {
                                                    fmt.Println(v91, fi85)
                                                    var v93 int32 = -852
                                                    v83 = (v83 ^ v83)
                                                    _ = v93
                                                }
                                                _ = v91
                                                _ = v92
                                            }
                                        }
                                    case -9223372036854775808:
                                        switch (v4 ^ (fi86 >> ((71) & 63))) {
                                        case 42:
                                            v94 := GenS2{F0: ((!true) && true)}
                                            _ = v94
                                        case 467:
                                            var v95 float64 = (0.0 * -2.25)
                                            _ = v95
                                        case -29:
                                            var v96 int64 = -9223372036854775808
                                            v97 := GenS7{F0: 33, F1: 84, F2: (^(v96))}
                                            _ = v96
                                            _ = v97
                                        }
                                    default:
                                        v38 = v38
                                        var v98 uint8 = (v41 >> ((4294967295) & 63))
                                        v38 -= 1
                                        _ = v98
                                    }
                                    v99 := GenS4{F0: 1}
                                    if false {
                                        v4 += (fi85 >> ((uint(v41)) & 63))
                                        var v100 uint32 = uint32(v41)
                                        var v101 int64 = 1
                                        _ = v100
                                        _ = v101
                                    } else {
                                        v102 := GenS8{F0: 42, F1: 284, F2: int64(v83)}
                                        if true {
                                            v83 = (v83 >> (((v83 ^ 136)) & 63))
                                            var v103 int = 32767
                                            _ = v103
                                        } else {
                                            if (!false) {
                                                switch (fi86 ^ fi85) {
                                                case 692:
                                                    var v104 uint64 = 6
                                                    _ = v104
                                                }
                                            } else {
                                                var v105 int = fi85
                                                _ = v105
                                            }
                                        }
                                        _ = v102
                                    }
                                    _ = v99
                                } else {
                                    if (!true) {
                                        v106 := GenS3{F0: v41, F1: ("a" + (("a" + "") + "hello"))}
                                        _ = v106
                                    }
                                    v4 = (fi85 ^ (v4 | v4))
                                }
                            }
                        }
                        fmt.Println(v83)
                        _ = fi86
                    }
                } else {
                    for fi107 := 0; fi107 < 3; fi107++ {
                        v108 := GenS6{F0: ("\n\t" + "a")}
                        _ = fi107
                        _ = v108
                    }
                    var v109 float64 = ((-2.25 / ((-2.25 + 0.0) * float64(-2.25))) - (1e10 * 0.1))
                    _ = v109
                }
                var v110 uint32 = 477
                _ = len(v3)
                _ = fi85
                _ = v110
            }
            var v111 float64 = (1.5 * 0.1)
            if (true || (true && (317 > 554))) {
                v112 := GenS9{F0: (v4 * 1), F1: 172, F2: v83}
                _ = v112
            } else {
                if ((^(v38)) <= (v2 / (((^(v2))) | 1))) {
                    switch v4 {
                    case 0:
                        for fi113 := 0; fi113 < 6; fi113++ {
                            if true {
                                v114 := GenS9{F0: -963, F1: -43, F2: v83}
                                v5 -= v5
                                _ = v114
                            } else {
                                v115 := GenS9{F0: (fi113 << (((v83 + v83)) & 63)), F1: 0, F2: v83}
                                _ = v115
                            }
                            switch "hello" {
                            case "":
                                for fi116 := 0; fi116 < 11; fi116++ {
                                    v117 := GenS10{F0: true, F1: "\n\t"}
                                    var v118 float32 = 0.1
                                    _ = fi116
                                    _ = v117
                                    _ = v118
                                }
                            case "a":
                                for fi119 := 0; fi119 < 5; fi119++ {
                                    var v120 int = 0
                                    fmt.Println(v4, v83, v38)
                                    _ = fi119
                                    _ = v120
                                }
                            default:
                                if ((v4 | (fi113 + v4)) <= -267) {
                                    v121 := []int{(v4 & v4), (fi113 + 127)}
                                    v41 = 255
                                    _ = v121
                                }
                            }
                            v122 := GenS5{F0: v2}
                            _ = fi113
                            _ = v122
                        }
                    default:
                        v111 -= 0.0
                    }
                } else {
                    var v123 float32 = ((((3.14159 * (-2.25 - 1.5)) + (3.14159 + 0.1)) + (((1e10 * -2.25) + (0.1 / ((1.5 / -2.25) + 3.14159))) - (3.14159 / -2.25))) * ((3.14159 - ((1.5 / 1e10) / (0.1 / 1.5))) - 0.1))
                    fmt.Println(v4, v111, v123)
                    _ = v123
                }
            }
            _ = v111
        } else {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v38, v2)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v83
}
