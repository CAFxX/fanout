package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 bool; F1 int64; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint; F2 int16 }
type GenS5 struct { F0 int32; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint32; F1 string; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int16; F2 bool }
type GenS9 struct { F0 uint8; F1 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint32 }
func genFn0() int8 {
    v4 := GenS1{F0: 824, F1: -2147483648}
    if true {
        var v5 uint8 = 128
        v6 := []int{-745}
        if false {
            v7 := GenS1{F0: 65535, F1: -2147483648}
            _ = v7
        }
        _ = v5
        _ = v6
    } else {
        var v8 int8 = -1
        var v9 float64 = (3.14159 * (float64((-2.25 / ((1.5 / ((3.14159 - (1.5 * 3.14159)) / 1e10)) + float32((0.0 * (0.0 * 3.14159)))))) * 3.14159))
        v10 := GenS1{F0: 2, F1: 0}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint = 18446744073709551615
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: 0, F1: (18446744073709551615 != 0)}
            if ((9223372036854775807 >= 1) && true) {
                v12 -= v12
                var v15 int8 = 84
                v12 = (v12 / (((v12 ^ (v12 & (^(v12))))) | 1))
                _ = v15
            } else {
                switch -845 {
                case 2147483647:
                    fmt.Println(fi11, fi11)
                case -2147483648:
                    v12 = (v12 << ((v12) & 63))
                default:
                    if true {
                        v16 := []int{(-(fi11)), (-(fi11)), (fi13 / (((fi13 + (fi11 * -689))) | 1)), (^(fi13))}
                        var v17 uint = (^(v12))
                        _ = v16
                        _ = v17
                    } else {
                        var v18 float64 = 3.14159
                        _ = v18
                    }
                }
            }
            _ = fi13
            _ = v14
        }
        var v19 string = ""
        _ = fi11
        _ = v12
        _ = v19
    }
    _ = v4
    return -128
}

func main() {
    var v1 int8 = 100
    var v2 uint = 65535
    v3 := make(map[uint32]int)
    switch 85 {
    case -526:
        v20 := GenS3{F0: (true || false), F1: 763, F2: 347}
        _ = v20
    case 1:
        if true {
            v21 := []int{-2147483648, -670, 574, -2147483648}
            var v22 uint8 = 2
            var v23 bool = (255 == (v22 & v22))
            _ = v21
            _ = v22
            _ = v23
        } else {
            switch (((("a" + "\n\t") + ("" + ("a" + "hello"))) + "") + "hello") {
            case "hello":
                v24 := v3[4294967295]
                _ = v24
            case "":
                fmt.Println(v2, v2)
            default:
                fmt.Println(v2)
            }
        }
    case -2147483648:
        var v25 uint16 = 32767
        _ = v25
    default:
        for fi26 := 0; fi26 < 13; fi26++ {
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 uint8 = 127
                if false {
                    _ = len(v3)
                    v29 := []int{fi26, (fi26 & -2147483648), -144}
                    switch "a" {
                    case "foo bar":
                        var v30 uint32 = 363
                        _ = v30
                    case "a":
                        v2 = 255
                        v1 = genFn0()
                        switch "" {
                        case "hello":
                            v2 = (^(v2))
                        case "a":
                            var v31 int32 = -980
                            _ = v31
                        case "":
                            var v32 uint8 = (v28 | 2)
                            _ = v32
                        }
                    default:
                        var v33 int16 = 1000
                        _ = v33
                    }
                    _ = len(v29)
                    _ = v29
                } else {
                    switch (fi27 - fi26) {
                    case -2147483648:
                        var v34 float32 = 1.5
                        _ = v34
                    case 0:
                        if (genFn0() != -128) {
                            fmt.Println(v28, v1, fi26)
                            fmt.Println(v1, v28, v2)
                            v35 := v3[150]
                            _ = v35
                        }
                    }
                    if ((^(v2)) >= v2) {
                        v36 := GenS2{F0: v28, F1: (v2 > (v2 * v2))}
                        _ = v36
                    }
                    var v37 float32 = 0.0
                    _ = v37
                }
                _ = fi27
                _ = v28
            }
            for fi38 := 0; fi38 < 15; fi38++ {
                if (2147483647 >= 2147483648) {
                    switch "" {
                    case "":
                        fmt.Println(v1, v2)
                    case "a":
                        v2 = (v2 << (((v2 & v2)) & 63))
                    case "\n\t":
                        v39 := GenS1{F0: 0, F1: -41}
                        _ = v39
                    }
                    v40 := GenS4{F0: 141, F1: v2, F2: 32767}
                    var v41 int8 = int8(fi38)
                    if false {
                        var v42 int = (fi26 << ((v2) & 63))
                        var v43 float64 = 1.5
                        v44 := []int{}
                        _ = v42
                        _ = v43
                        _ = v44
                    } else {
                        if (true || ((((3.14159 * 1.5) / 1e10) - 3.14159) > 0.0)) {
                            if (0 < 42) {
                                v1 = int8(fi26)
                                switch "\n\t" {
                                case "a":
                                    switch int(v2) {
                                    case -92:
                                        if ("a" <= "hello") {
                                            delete(v3, 4294967295)
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                v41 = v1
                                            case "hello":
                                                var v45 float32 = -2.25
                                                _ = v45
                                            case "\n\t":
                                                var v46 int64 = 718
                                                _ = v46
                                            default:
                                                v2 += 1
                                            }
                                            var v47 int = 42
                                            switch v47 {
                                            case 32768:
                                                v47 = 32767
                                            case 127:
                                                v47 = v47
                                            case 128:
                                                switch "foo bar" {
                                                case "":
                                                    v48 := GenS5{F0: 0, F1: 42, F2: 42}
                                                    var v49 uint64 = 4294967295
                                                    _ = v48
                                                    _ = v49
                                                case "\n\t":
                                                    v41 -= (v1 / (((v41 % ((-128) | 1))) | 1))
                                                case "hello":
                                                    v50 := GenS4{F0: 691, F1: 18446744073709551615, F2: 1}
                                                    _ = v50
                                                }
                                            default:
                                                var v51 int8 = (v1 << ((v2) & 63))
                                                _ = v51
                                            }
                                            _ = v47
                                        }
                                    case 128:
                                        v52 := GenS2{F0: 255, F1: ((true || true) || false)}
                                        v41 = v41
                                        _ = v52
                                    case -893:
                                        v2 = v2
                                    default:
                                        fmt.Println(v2)
                                    }
                                case "foo bar":
                                    var v53 float32 = 0.1
                                    _ = v53
                                case "\n\t":
                                    if true {
                                        switch "hello" {
                                        case "hello":
                                            v54 := []int{(fi26 % ((fi38) | 1)), fi26}
                                            _ = v54
                                        case "":
                                            if false {
                                                if (("hello" + (("a" + "a") + "")) == "\n\t") {
                                                    v55 := []int{int(v41), fi38, (-(fi26))}
                                                    var v56 int16 = -775
                                                    _ = v55
                                                    _ = v56
                                                }
                                                var v57 bool = true
                                                v58 := GenS4{F0: 0, F1: 472, F2: 42}
                                                fmt.Println(v41, fi38, fi38)
                                                if (!(0 < 0)) {
                                                    var v59 float32 = (((3.14159 / 1.5) - (3.14159 * float32(1.5))) / 1.5)
                                                    v2 = v2
                                                    _ = v59
                                                } else {
                                                    v41 = 1
                                                    v60 := GenS6{F0: 2147483648, F1: "", F2: 0}
                                                    fmt.Println(v1, v57)
                                                    _ = v60
                                                }
                                                v41 -= (v1 ^ v1)
                                                _ = v57
                                                _ = v58
                                            } else {
                                                var v61 float64 = 1e10
                                                fmt.Println(v41, v41)
                                                fmt.Println(v2)
                                                v62 := GenS7{F0: true, F1: 2147483647}
                                                var v63 uint32 = 42
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                    }
                                    var v64 uint = (v2 | v2)
                                    _ = v64
                                }
                            } else {
                                var v65 int16 = int16(v41)
                                if false {
                                    fmt.Println(v41, v65)
                                } else {
                                    switch fi38 {
                                    case 2147483647:
                                        var v66 float64 = 3.14159
                                        _ = v66
                                    case 127:
                                        var v67 uint32 = 2147483647
                                        _ = v67
                                    default:
                                        if false {
                                            if (83 >= 42) {
                                                var v68 uint64 = 42
                                                v41 -= v41
                                                var v69 bool = true
                                                fmt.Println(fi26, v41, v65)
                                                if v69 {
                                                    if ((v2 | v2) == v2) {
                                                        fmt.Println(fi26, v65, fi26)
                                                        var v70 uint32 = 2147483648
                                                        v71 := GenS2{F0: uint8(v68), F1: (!v69)}
                                                        v72 := GenS8{F0: 127, F1: (v65 % (((v65 >> ((405) & 63))) | 1)), F2: true}
                                                        v73 := v3[(v70 % (((v70 / (((v70 << (((v2 & v2)) & 63))) | 1))) | 1))]
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    delete(v3, 2147483648)
                                                    if ((v68 ^ (v68 + v68)) > v68) {
                                                        if (2 <= 815) {
                                                            switch fi26 {
                                                            case 42:
                                                                fmt.Println(v68)
                                                            }
                                                        }
                                                        switch fi26 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v69, fi38)
                                                            if v69 {
                                                                var v74 uint32 = 224
                                                                v75 := []int{}
                                                                if true {
                                                                    v2 = 2
                                                                }
                                                                var v76 bool = (255 >= 255)
                                                                if v69 {
                                                                    _ = len(v75)
                                                                    fmt.Println(v41, fi26)
                                                                }
                                                                _ = v74
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            v77 := []int{9223372036854775807}
                                                            _ = v77
                                                        }
                                                        v2 = 0
                                                    } else {
                                                        v2 = (v2 - 65535)
                                                    }
                                                } else {
                                                    if v69 {
                                                        v78 := GenS7{F0: ((v1 + (-(v41))) == -1), F1: 2}
                                                        fmt.Println(v65)
                                                        _ = v78
                                                    } else {
                                                        fmt.Println(v65, v2, v65)
                                                        if (!(0.0 == (3.14159 * -2.25))) {
                                                            switch ("" + "\n\t") {
                                                            case "\n\t":
                                                                v69 = true
                                                            default:
                                                                v79 := GenS9{F0: 2, F1: v69}
                                                                _ = v79
                                                            }
                                                            v80 := GenS3{F0: (!v69), F1: 202, F2: v68}
                                                            v41 += v1
                                                            _ = v80
                                                        } else {
                                                            fmt.Println(v69, v65)
                                                            v69 = (!(float64(((1.5 - ((0.1 * (-2.25 * -2.25)) * 1e10)) - 1.5)) <= float64((((3.14159 * -2.25) - 0.1) + (1e10 / 1.0)))))
                                                            v81 := GenS8{F0: 97, F1: v65, F2: ((0 < 42) && v69)}
                                                            _ = v81
                                                        }
                                                        v41 -= (v41 ^ genFn0())
                                                    }
                                                    v41 = (v1 >> ((v2) & 63))
                                                    var v82 uint32 = 284
                                                    _ = len(v3)
                                                    switch (^(fi38)) {
                                                    case 380:
                                                        if true {
                                                            if true {
                                                                var v83 uint16 = 1
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v84 int8 = int8(fi38)
                                                                    _ = v84
                                                                case "foo bar":
                                                                    v3[(v82 / ((v82) | 1))] = (fi38 + (fi38 + fi38))
                                                                default:
                                                                    var v85 int64 = -9223372036854775808
                                                                    _ = v85
                                                                }
                                                                _ = v83
                                                            } else {
                                                                v69 = (((v69 || (v65 > 32767)) || v69) || (v65 == (v65 - (v65 * v65))))
                                                            }
                                                            v69 = ((0.0 < ((1.5 * ((3.14159 / 1e10) / (1.5 * (0.1 / -2.25)))) + (-2.25 / (1e10 * -2.25)))) && true)
                                                        }
                                                    case 892:
                                                        var v86 int32 = -2147483648
                                                        _ = v86
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            fmt.Println(v2, v82)
                                                        }
                                                    }
                                                    _ = v82
                                                }
                                                v87 := GenS10{F0: fi26}
                                                v88 := v3[65535]
                                                _ = v68
                                                _ = v69
                                                _ = v87
                                                _ = v88
                                            } else {
                                                if false {
                                                    var v89 int = fi26
                                                    if false {
                                                        v2 = (v2 % ((v2) | 1))
                                                    } else {
                                                        v41 = v41
                                                    }
                                                    _ = v89
                                                } else {
                                                    v65 = (v65 >> ((v2) & 63))
                                                    var v90 int = int(v65)
                                                    var v91 int8 = (v41 / (((v1 & (-(v41)))) | 1))
                                                    var v92 bool = (2147483647 < 825)
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                }
                                                fmt.Println(v1)
                                                var v93 int32 = 2147483647
                                                _ = v93
                                            }
                                            v94 := GenS9{F0: 0, F1: (!false)}
                                            _ = v94
                                        } else {
                                            if true {
                                                var v95 uint16 = 0
                                                _ = v95
                                            }
                                            v96 := GenS11{F0: 42, F1: 240}
                                            var v97 uint = v2
                                            _ = v96
                                            _ = v97
                                        }
                                    }
                                }
                                _ = v65
                            }
                        } else {
                            var v98 int = 4294967295
                            v99 := GenS12{F0: 315}
                            var v100 float32 = float32(1e10)
                            _ = v100
                            _ = v98
                            _ = v99
                        }
                        if (fi38 < fi38) {
                            v41 -= (^(v41))
                            v41 -= (v1 % ((v1) | 1))
                        } else {
                            var v101 int64 = -9223372036854775808
                            v41 = -1
                            v102 := GenS3{F0: ((!true) || (true || (725 <= (-(v101))))), F1: (v101 / (((v101 - v101)) | 1)), F2: 18446744073709551615}
                            if (v2 <= (v2 & (^(v2)))) {
                                var v103 uint32 = 42
                                var v104 int = fi38
                                var v105 uint = (v2 ^ (v2 | 0))
                                _ = v103
                                _ = v104
                                _ = v105
                            } else {
                                var v106 uint32 = 4294967295
                                v107 := GenS13{F0: 215}
                                fmt.Println(v1, fi26)
                                var v108 float32 = -2.25
                                _ = v106
                                _ = v107
                                _ = v108
                            }
                            _ = v101
                            _ = v102
                        }
                    }
                    _ = v40
                    _ = v41
                }
                v109 := GenS14{F0: 1000}
                _ = fi38
                _ = v109
            }
            _ = fi26
        }
    }
    v1 += (v1 - v1)
    v1 = v1
    for fi110 := 0; fi110 < 14; fi110++ {
        var v111 float64 = 0.0
        if true {
            for fi112 := 0; fi112 < 5; fi112++ {
                v113 := GenS15{F0: 233}
                v114 := []int{fi112, fi112, (fi110 + (fi112 >> (((v2 % ((v2) | 1))) & 63)))}
                _ = fi112
                _ = v113
                _ = v114
            }
        }
        fmt.Println(v2, v111)
        _ = fi110
        _ = v111
    }
    var v115 int16 = -1
    for fi116 := 0; fi116 < 12; fi116++ {
        var v117 uint16 = 423
        _ = fi116
        _ = v117
    }
    v115 = (v115 & 423)
    if (1 != 1) {
        v118 := []int{-1, 128, 128, 32767}
        v119 := GenS16{F0: 1}
        _ = v118
        _ = v119
    }
    v1 += v1
    v1 -= (v1 & v1)
    if false {
        var v120 int64 = -1
        for fi121 := 0; fi121 < 4; fi121++ {
            v122 := GenS8{F0: 155, F1: v115, F2: (!(true && (!((fi121 <= 127) || true))))}
            var v123 string = "foo bar"
            _ = fi121
            _ = v122
            _ = v123
        }
        _ = v120
    }
    v2 += (^(v2))
    v124 := v3[669]
    if (v2 <= 165) {
        v115 = v115
    }
    v115 -= (v115 - (^(v115)))
    for fi125 := 0; fi125 < 14; fi125++ {
        v126 := GenS7{F0: true, F1: 4294967295}
        _ = fi125
        _ = v126
    }
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v115
    _ = v124
    _ = v2
    _ = v3
}
