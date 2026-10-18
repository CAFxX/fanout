package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint64; F1 string; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32; F1 uint; F2 int16 }
type GenS11 struct { F0 uint; F1 int64; F2 uint }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int64; F1 int64; F2 int8 }
type GenS15 struct { F0 int16; F1 string }
type GenS16 struct { F0 uint8; F1 string; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 uint64 = 532
    var v4 uint16 = 42
    var v5 string = "\n\t"
    var v6 bool = false
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: (fi9 + fi9), F1: 127, F2: 45}
            switch "\n\t" {
            case "foo bar":
                if v6 {
                    switch ((v5 + v5) + ("foo bar" + (v5 + v5))) {
                    case "hello":
                        v11 := GenS1{F0: fi9, F1: 42, F2: 139}
                        _ = v11
                    case "":
                        if true {
                            v12 := GenS2{F0: 65535, F1: fi9}
                            if v6 {
                                switch v5 {
                                case "a":
                                    fmt.Println(v6, fi9, v6)
                                case "\n\t":
                                    v13 := GenS1{F0: 32767, F1: 1, F2: 128}
                                    v4 = v7
                                    v6 = (!v6)
                                    if (!v6) {
                                        var v14 float64 = 0.1
                                        v15 := []int{}
                                        _ = v14
                                        _ = v15
                                    } else {
                                        fmt.Println(v5, v7)
                                    }
                                    switch ((v5 + (v5 + "a")) + v5) {
                                    case "\n\t":
                                        v16 := GenS1{F0: 738, F1: 127, F2: 0}
                                        _ = v16
                                    case "":
                                        fmt.Println(v4)
                                    }
                                    _ = v13
                                case "foo bar":
                                    v4 -= v4
                                }
                            } else {
                                if (!v6) {
                                    if (!v6) {
                                        fmt.Println(fi8, v7, v5)
                                    }
                                } else {
                                    if v6 {
                                        switch fi9 {
                                        case -721:
                                            var v17 int8 = -1
                                            _ = v17
                                        case 1:
                                            var v18 bool = v6
                                            v19 := GenS3{F0: 100}
                                            _ = v18
                                            _ = v19
                                        case -433:
                                            var v20 float64 = 1.5
                                            _ = v20
                                        }
                                    }
                                    var v21 bool = (!(v6 || false))
                                    _ = v21
                                }
                                var v22 float32 = float32(0.0)
                                v23 := GenS3{F0: -3}
                                _ = v22
                                _ = v23
                            }
                            _ = v12
                        } else {
                            v24 := []int{fi8, fi9, (fi9 & -126)}
                            _ = v24
                        }
                    case "\n\t":
                        var v25 uint16 = 42
                        _ = v25
                    }
                } else {
                    v26 := []int{(fi9 * 0), 890, fi9, (-(fi9))}
                    v7 -= v7
                    if false {
                        v3 -= (^(v3))
                        v27 := GenS4{F0: v3, F1: v3}
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        v28 := GenS2{F0: 4294967295, F1: fi9}
                        v29 := GenS5{F0: 621, F1: "", F2: (fi9 * (fi8 - (^(fi9))))}
                        v4 -= (^(v7))
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS2{F0: 42, F1: 9223372036854775807}
                        _ = v30
                    }
                    switch (fi8 % (((^(fi9))) | 1)) {
                    case -917:
                        _ = len(v1)
                    case 2147483647:
                        var v31 int8 = 1
                        v5 = (((v5 + (("\n\t" + v5) + (v5 + ("foo bar" + (((v5 + ("" + v5)) + v5) + "\n\t"))))) + "foo bar") + "")
                        v7 = 0
                        _ = len(v1)
                        v6 = v6
                        v6 = v6
                        v5 = v5
                        v6 = false
                        var v32 bool = (0 >= 1)
                        _ = v31
                        _ = v32
                    case 0:
                        if v6 {
                            v33 := GenS6{F0: 301}
                            _ = v33
                        }
                    default:
                        v34 := GenS7{F0: -920, F1: -9223372036854775808}
                        _ = v34
                    }
                    _ = v26
                }
            }
            var v35 int = 127
            _ = fi9
            _ = v10
            _ = v35
        }
        v3 -= 579
        _ = fi8
    }
    var v36 int16 = 1000
    var v37 int16 = (v36 - 1)
    v7 -= v7
    switch 42 {
    case 574:
        for fi38 := 0; fi38 < 3; fi38++ {
            v39 := GenS6{F0: 2147483647}
            if v6 {
                v6 = (-2.25 == float32(3.14159))
                var v40 int64 = 438
                _ = v40
            } else {
                switch (fi38 - 1) {
                case 32768:
                    v41 := []int{(fi38 % ((fi38) | 1))}
                    _ = v41
                case -233:
                    if (1 >= 0) {
                        v5 = v5
                    } else {
                        var v42 float32 = (1e10 * 1e10)
                        for fi43 := 0; fi43 < 20; fi43++ {
                            fmt.Println(v6, v37)
                            v44 := GenS8{F0: (v3 >> ((424) & 63))}
                            var v45 uint8 = 42
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        if (-2147483648 >= -2147483648) {
                            v5 = v5
                            v4 = 0
                        } else {
                            fmt.Println(v4, fi38)
                        }
                        _ = v42
                    }
                default:
                    v4 = (v4 ^ 828)
                }
                for fi46 := 0; fi46 < 6; fi46++ {
                    if (1 > 18446744073709551615) {
                        v3 += (v3 * (^(v3)))
                        var v47 string = "hello"
                        if (("a" + v5) >= ((v5 + v47) + v5)) {
                            v7 = (v4 % (((v7 / (((^(v4))) | 1))) | 1))
                        }
                        v48 := []int{fi38, -9223372036854775808, fi38, fi46}
                        _ = v47
                        _ = v48
                    }
                    var v49 uint64 = uint64(v4)
                    _ = fi46
                    _ = v49
                }
                fmt.Println(v36, v36, v37)
            }
            _ = fi38
            _ = v39
        }
    case -2147483648:
        if true {
            for fi50 := 0; fi50 < 1; fi50++ {
                for fi51 := 0; fi51 < 8; fi51++ {
                    switch fi50 {
                    case 930:
                        if v6 {
                            v5 = "hello"
                            var v52 float64 = (0.1 - float64(1e10))
                            var v53 uint64 = 18446744073709551615
                            var v54 uint64 = v53
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    case 42:
                        var v55 uint32 = 1
                        _ = v55
                    case 127:
                        var v56 float32 = (1.5 / 3.14159)
                        _ = v56
                    }
                    v7 = (v7 - (v7 | v4))
                    switch fi51 {
                    case 32768:
                        if (-265 < 0) {
                            v57 := []int{fi50, (fi50 & (fi51 | (fi51 / (((fi50 ^ int(v4))) | 1)))), fi50, fi50}
                            v6 = false
                            _ = v57
                        }
                    case 32767:
                        v36 = (v36 & (v36 | v37))
                    case -1:
                        v58 := []int{fi51, (fi50 / ((fi51) | 1)), (fi50 << ((65535) & 63)), (fi50 ^ (fi51 + fi51))}
                        _ = v58
                    }
                    _ = fi51
                }
                _ = fi50
            }
            fmt.Println(v6)
            if (978 <= 32767) {
                var v59 uint32 = 328
                if true {
                    fmt.Println(v59)
                    if (962 != -1) {
                        for fi60 := 0; fi60 < 11; fi60++ {
                            var v61 float64 = 1.5
                            v62 := GenS9{F0: fi60}
                            _ = len(v2)
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        v63 := GenS3{F0: -128}
                        fmt.Println(v4)
                        _ = v63
                    } else {
                        v3 += v3
                        switch 691 {
                        case 936:
                            switch "\n\t" {
                            case "a":
                                v7 -= 255
                            case "hello":
                                var v64 uint = 1
                                _ = v64
                            case "foo bar":
                                var v65 int8 = -128
                                _ = v65
                            default:
                                v59 -= 2147483647
                            }
                        case -9223372036854775808:
                            v6 = (false && false)
                        case -645:
                            var v66 int16 = (v37 + (v37 * v37))
                            _ = v66
                        default:
                            var v67 int64 = 42
                            _ = v67
                        }
                    }
                } else {
                    var v68 string = (v5 + v5)
                    var v69 float64 = (1.5 * 0.1)
                    var v70 uint32 = 242
                    _ = v68
                    _ = v69
                    _ = v70
                }
                _ = v59
            } else {
                var v71 int32 = 0
                _ = v71
            }
        } else {
            v72 := GenS10{F0: 0, F1: 370, F2: 1}
            _ = v72
        }
    case 39:
        if ("a" > v5) {
            v73 := []int{4294967295, -2147483648, -94}
            for fi74 := 0; fi74 < 6; fi74++ {
                fmt.Println(v7, fi74, v36)
                var v75 int16 = (v36 & -982)
                var v76 int = (fi74 << ((65535) & 63))
                _ = fi74
                _ = v75
                _ = v76
            }
            var v77 uint32 = 4294967295
            _ = v73
            _ = v77
        } else {
            v78 := []int{}
            switch 32768 {
            case 2147483647:
                v7 = v7
            }
            if v6 {
                var v79 int32 = -535
                v5 = (((v5 + "a") + v5) + "")
                if len(v78) >= 2 {
                    v78 = v78[1:len(v78)]
                }
                _ = v79
            }
            _ = v78
        }
    }
    switch 42 {
    case -1:
        var v80 uint8 = 1
        _ = v80
    }
    for fi81 := 0; fi81 < 20; fi81++ {
        if false {
            if (v6 && (0.0 > (0.1 + 1.5))) {
                fmt.Println(v5)
                v82 := GenS11{F0: 0, F1: 0, F2: 723}
                fmt.Println(v4, v5, v37)
                _ = v82
            } else {
                v7 = (v7 % (((v4 >> ((18446744073709551615) & 63))) | 1))
                for fi83 := 0; fi83 < 15; fi83++ {
                    if false {
                        v5 = v5
                        v84 := []int{(fi83 << ((4294967295) & 63)), (fi81 | (fi83 ^ (fi81 << ((1) & 63)))), fi83, (fi81 % ((42) | 1))}
                        fmt.Println(v5, v7)
                        _ = v84
                    }
                    var v85 int8 = 42
                    v7 -= (v7 | uint16(v3))
                    _ = fi83
                    _ = v85
                }
            }
            var v86 uint64 = v3
            _ = v86
        }
        var v87 uint16 = (v7 / ((v4) | 1))
        _ = fi81
        _ = v87
    }
    if (!false) {
        for fi88 := 0; fi88 < 7; fi88++ {
            var v89 int64 = 42
            if true {
                if v6 {
                    v90 := []int{(fi88 << ((748) & 63)), fi88}
                    v37 -= int16(v7)
                    _ = v90
                }
                var v91 float32 = -2.25
                var v92 bool = v6
                _ = v91
                _ = v92
            } else {
                var v93 uint16 = (v4 % ((245) | 1))
                fmt.Println(v93, v37, v4)
                switch fi88 {
                case 2147483647:
                    for fi94 := 0; fi94 < 20; fi94++ {
                        var v95 uint8 = 87
                        switch (fi88 * 42) {
                        case 4294967295:
                            if v6 {
                                v96 := []int{}
                                if (((v95 ^ (v95 << ((987) & 63))) > 255) || true) {
                                    v6 = (false || ((1e10 * (((float64(3.14159) + -2.25) + 0.0) + (-2.25 * -2.25))) >= 1e10))
                                    if (4294967295 == 400) {
                                        var v97 uint = 0
                                        if (!(v6 || v6)) {
                                            switch int(v97) {
                                            case 42:
                                                v7 += 901
                                            case 332:
                                                if (!true) {
                                                    v3 = (v3 | (v3 ^ 587))
                                                    v98 := GenS2{F0: (^(v97)), F1: fi94}
                                                    v89 = 1
                                                    _ = v98
                                                }
                                            }
                                            var v99 uint16 = (v93 | v93)
                                            fmt.Println(v95, v37)
                                            _ = v99
                                        } else {
                                            if true {
                                                if ((((^(v37)) != 42) && v6) || v6) {
                                                    if false {
                                                        v95 = v95
                                                    } else {
                                                        v100 := GenS6{F0: 308}
                                                        var v101 float32 = (-2.25 - -2.25)
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                }
                                                var v102 bool = (!(v89 != (v89 | v89)))
                                                _ = v102
                                            }
                                            v5 = (v5 + v5)
                                            var v103 uint = v97
                                            var v104 float64 = (1.5 - float64(1e10))
                                            _ = v103
                                            _ = v104
                                        }
                                        _ = v97
                                    }
                                    var v105 float64 = (-2.25 / 1.5)
                                    _ = v105
                                } else {
                                    v106 := GenS1{F0: int(v89), F1: 1, F2: (v95 / (((v95 % (((^(v95))) | 1))) | 1))}
                                    _ = v106
                                }
                                v3 -= 697
                                _ = v96
                            } else {
                                var v107 int16 = (v37 | (v36 * v37))
                                v3 += (v3 / ((v3) | 1))
                                _ = v107
                            }
                        case -554:
                            var v108 uint16 = uint16(v95)
                            _ = v108
                        case -15:
                            v89 = (v89 << ((18446744073709551615) & 63))
                        default:
                            fmt.Println(v37)
                        }
                        _ = fi94
                        _ = v95
                    }
                case 1:
                    _ = len(v2)
                case 353:
                    v109 := GenS9{F0: 42}
                    _ = v109
                }
                _ = v93
            }
            for fi110 := 0; fi110 < 17; fi110++ {
                v5 = v5
                switch "foo bar" {
                case "\n\t":
                    v37 = (v36 | (v36 >> ((0) & 63)))
                case "":
                    var v111 float64 = 0.1
                    _ = v111
                }
                _ = fi110
            }
            _ = fi88
            _ = v89
        }
        v37 = (v36 & -32768)
        fmt.Println(v37, v5, v3)
    }
    var v112 float64 = float64((3.14159 / 1e10))
    var v113 uint32 = 1
    var v114 string = ((v5 + ((v5 + v5) + v5)) + v5)
    switch (v114 + v114) {
    case "a":
        for fi115 := 0; fi115 < 16; fi115++ {
            for fi116 := 0; fi116 < 4; fi116++ {
                v117 := GenS12{F0: v3}
                v114 = "foo bar"
                var v118 uint = 6
                switch "\n\t" {
                case "":
                    v6 = (100 == 127)
                default:
                    v119 := v2[(v5 + v5)]
                    _ = v119
                }
                _ = fi116
                _ = v117
                _ = v118
            }
            v120 := []int{}
            _ = fi115
            _ = v120
        }
    default:
        fmt.Println(v36, v7, v112)
    }
    switch 369 {
    case 126:
        switch 32768 {
        case 2147483647:
            v2[(v5 + v5)] = v5
        case -2147483648:
            var v121 uint32 = (v113 * (v113 & 0))
            _ = v121
        case -767:
            if ((v37 & (v37 - v37)) < v36) {
                v122 := []int{}
                v123 := v2[v114]
                var v124 int64 = 0
                _ = v122
                _ = v123
                _ = v124
            } else {
                var v125 uint32 = (^(v113))
                var v126 int16 = (v37 & (^(v37)))
                if v6 {
                    var v127 int8 = -22
                    v128 := GenS9{F0: 42}
                    _ = v127
                    _ = v128
                } else {
                    for fi129 := 0; fi129 < 17; fi129++ {
                        if (-128 == int8(v113)) {
                            v5 = v114
                        } else {
                            v130 := GenS13{F0: (v3 - (v3 ^ v3))}
                            if ((v5 + v5) > ((("\n\t" + v114) + "") + "\n\t")) {
                                var v131 int8 = 58
                                _ = v131
                            }
                            _ = v130
                        }
                        for fi132 := 0; fi132 < 16; fi132++ {
                            if false {
                                v133 := GenS3{F0: -128}
                                var v134 int16 = v37
                                _ = v133
                                _ = v134
                            }
                            var v135 int64 = 9223372036854775807
                            var v136 int16 = int16(fi129)
                            v137 := GenS14{F0: 0, F1: v135, F2: 127}
                            _ = fi132
                            _ = v135
                            _ = v136
                            _ = v137
                        }
                        v2["hello"] = "hello"
                        _ = fi129
                    }
                }
                _ = v125
                _ = v126
            }
        }
    case 2147483647:
        var v138 string = (("\n\t" + "") + "\n\t")
        _ = v138
    case 127:
        if (!(255 == 1)) {
            v139 := GenS15{F0: (-(v37)), F1: v5}
            var v140 string = "foo bar"
            for fi141 := 0; fi141 < 19; fi141++ {
                fmt.Println(v7, v140, v5)
                v142 := v2[""]
                v143 := []int{(fi141 << ((1) & 63)), 127}
                _ = fi141
                _ = v142
                _ = v143
            }
            _ = v139
            _ = v140
        } else {
            switch -344 {
            case 913:
                if ((65535 >= 2147483648) && (32768 != -1)) {
                    v144 := GenS9{F0: -1}
                    v145 := []int{42, 42}
                    if false {
                        var v146 uint = 0
                        _ = v146
                    } else {
                        var v147 uint32 = (v113 - 2147483648)
                        v148 := GenS16{F0: 255, F1: (v5 + v5), F2: 255}
                        switch v5 {
                        case "hello":
                            v113 = 1
                        }
                        _ = v147
                        _ = v148
                    }
                    _ = v144
                    _ = v145
                }
            default:
                var v149 int = 128
                _ = v149
            }
        }
    default:
        fmt.Println(v113, v114)
    }
    v150 := []int{2147483647, 0, 9223372036854775807}
    v1[(v113 - (v113 << ((4294967295) & 63)))] = 640
    v150 = append(v150, -2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v150:", len(v150), v150)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v150
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
