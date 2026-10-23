package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 bool; F2 uint }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool; F1 string; F2 uint }
type GenS13 struct { F0 uint8; F1 uint32; F2 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint16; F1 int32; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 int8) bool {
    var v6 int8 = -128
    var v7 int8 = p2
    if true {
        p2 = (^(p2))
        if true {
            v6 -= (^(v7))
            for fi8 := 0; fi8 < 9; fi8++ {
                fmt.Println(fi8)
                fmt.Println(v7, fi8, fi8)
                switch "\n\t" {
                case "":
                    v6 = v6
                default:
                    v9 := []int{(fi8 ^ (fi8 & fi8)), (fi8 * 1), 0}
                    _ = v9
                }
                _ = fi8
            }
            if (!false) {
                v10 := GenS1{F0: (p0 * (p0 * (p0 << ((65535) & 63))))}
                v11 := GenS1{F0: (p0 >> ((1) & 63))}
                _ = v10
                _ = v11
            } else {
                var v12 int16 = -260
                _ = v12
            }
        }
        fmt.Println(p1, v6)
    }
    if ("hello" < "a") {
        for fi13 := 0; fi13 < 15; fi13++ {
            if (v6 == (p2 % ((100) | 1))) {
                v14 := GenS2{F0: "foo bar"}
                var v15 int16 = 1
                _ = v14
                _ = v15
            }
            v16 := []int{fi13, (fi13 / ((-9223372036854775808) | 1)), (fi13 % ((4294967295) | 1)), fi13}
            v7 -= 100
            _ = fi13
            _ = v16
        }
        fmt.Println(p2)
        for fi17 := 0; fi17 < 2; fi17++ {
            v6 += (v7 | v6)
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!(166 <= 2))
}

func main() {
    v1 := make(map[string]int)
    var v2 string = (("foo bar" + "\n\t") + ("foo bar" + "\n\t"))
    var v3 int8 = 0
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    var v18 int8 = v3
    var v19 uint8 = 1
    v20 := GenS3{F0: 2, F1: 1000}
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS4{F0: "foo bar", F1: 1}
        for fi23 := 0; fi23 < 13; fi23++ {
            delete(v5, -881)
            switch (v2 + v2) {
            case "foo bar":
                var v24 uint = 42
                _ = v24
            case "hello":
                v25 := GenS4{F0: v2, F1: uint32(v19)}
                _ = v25
            case "\n\t":
                if (65535 >= 1) {
                    v18 = (v18 * v3)
                }
            default:
                if ((1 <= 1) && (!((fi23 % ((9223372036854775807) | 1)) <= fi23))) {
                    switch (fi21 | fi23) {
                    case 0:
                        var v26 int16 = 0
                        fmt.Println(v18)
                        v27 := GenS5{F0: -1}
                        _ = v26
                        _ = v27
                    default:
                        v28 := v5[fi21]
                        _ = v28
                    }
                } else {
                    v3 += v3
                    switch (fi21 >> ((18446744073709551615) & 63)) {
                    case 129:
                        v29 := GenS4{F0: ("" + v2), F1: 978}
                        _ = v29
                    case 32768:
                        var v30 uint32 = 2147483647
                        _ = v30
                    case 128:
                        var v31 float64 = ((float64(-2.25) + (3.14159 - 1.5)) * (-2.25 + (float64((float32((1.5 - (((1.5 - (1e10 / 3.14159)) - 3.14159) + 0.0))) / -2.25)) + ((0.1 - float64(1.5)) - 1.5))))
                        _ = v31
                    }
                    if false {
                        var v32 int16 = -1
                        _ = v32
                    }
                }
            }
            _ = fi23
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            v2 = ("" + v2)
            v19 = v19
            v19 = (v19 >> ((190) & 63))
            delete(v1, v2)
            _ = fi33
        }
        _ = fi21
        _ = v22
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if (v2 <= v2) {
            switch (fi34 + 2147483647) {
            case -9223372036854775808:
                var v35 float64 = 3.14159
                _ = v35
            case 32768:
                for fi36 := 0; fi36 < 17; fi36++ {
                    var v37 int32 = -1
                    var v38 int = 4294967295
                    switch ("hello" + "\n\t") {
                    case "":
                        switch v38 {
                        case -682:
                            switch "" {
                            case "foo bar":
                                fmt.Println(v38)
                            case "":
                                if true {
                                    switch 128 {
                                    case 573:
                                        v39 := GenS5{F0: 0}
                                        _ = v39
                                    case 127:
                                        v40 := GenS6{F0: (!(true && false))}
                                        _ = v40
                                    }
                                    var v41 int = -9223372036854775808
                                    var v42 int64 = 1
                                    v43 := []int{fi34, (v38 % ((int(v18)) | 1)), -705}
                                    if len(v43) > 0 {
                                        v43[0] = v38
                                    }
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                } else {
                                    if false {
                                        v44 := GenS6{F0: false}
                                        _ = v44
                                    } else {
                                        if true {
                                            v45 := GenS5{F0: 1}
                                            _ = v45
                                        } else {
                                            switch (v2 + "foo bar") {
                                            case "foo bar":
                                                v19 = (v19 % ((v19) | 1))
                                            case "a":
                                                var v46 float64 = 0.0
                                                _ = v46
                                            case "":
                                                if true {
                                                    var v47 float64 = 1.5
                                                    v2 = ""
                                                    if (!true) {
                                                        v38 -= fi36
                                                        v48 := GenS5{F0: -9223372036854775808}
                                                        v49 := GenS7{F0: v2}
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    _ = v47
                                                }
                                            default:
                                                v2 = v2
                                            }
                                            v18 = 80
                                            v38 += (fi34 * -349)
                                            v50 := GenS5{F0: -9223372036854775808}
                                            _ = v50
                                        }
                                        _ = len(v4)
                                    }
                                    fmt.Println(v38)
                                    delete(v1, v2)
                                }
                                fmt.Println(fi36, fi36, fi36)
                                if true {
                                    v2 = v2
                                }
                                switch (("\n\t" + v2) + v2) {
                                case "hello":
                                    switch "\n\t" {
                                    case "foo bar":
                                        v51 := GenS8{F0: "", F1: true, F2: 0}
                                        _ = v51
                                    case "hello":
                                        if false {
                                            v38 = 32768
                                            v38 = v38
                                            var v52 float64 = 3.14159
                                            _ = v52
                                        }
                                    case "":
                                        v37 = (v37 << ((65535) & 63))
                                    default:
                                        var v53 float32 = float32(1e10)
                                        _ = v53
                                    }
                                default:
                                    delete(v4, "hello")
                                }
                                v18 = v18
                            default:
                                v19 += (v19 * 128)
                            }
                        case 32767:
                            v18 = -126
                        default:
                            if (741 >= 178) {
                                v37 = (^(v37))
                                if (v19 != v19) {
                                    switch ((v2 + (((((v2 + v2) + ((("" + v2) + "a") + v2)) + "foo bar") + v2) + v2)) + (v2 + v2)) {
                                    case "foo bar":
                                        if false {
                                            v54 := GenS9{F0: (v37 << ((0) & 63)), F1: -9}
                                            var v55 float64 = 0.0
                                            v56 := []int{}
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                    case "hello":
                                        if true {
                                            v57 := []int{v38, (v38 % ((v38) | 1))}
                                            _ = v57
                                        } else {
                                            if (!true) {
                                                var v58 uint64 = 392
                                                var v59 uint8 = (v19 | (v19 >> ((42) & 63)))
                                                _ = v58
                                                _ = v59
                                            } else {
                                                v3 = (v3 ^ v3)
                                            }
                                        }
                                    }
                                    v19 = v19
                                }
                            }
                            v38 -= fi34
                            v60 := []int{}
                            _ = v60
                        }
                    case "\n\t":
                        v61 := GenS7{F0: (((("foo bar" + v2) + v2) + "hello") + "\n\t")}
                        var v62 bool = true
                        if (-2.25 <= (((float32((((3.14159 + (((-2.25 / (0.1 / 1.0)) / -2.25) * (float64(1.5) - (1.5 / 0.1)))) - 1.5) * 0.1)) / float32((((((1.5 * 1.5) / (0.1 / ((1e10 - (((1.5 * -2.25) - float64(float32(float64(float32(1e10))))) - (1e10 + (0.0 + 1e10)))) / 1.0))) - -2.25) * ((0.0 * (0.1 + 3.14159)) + 3.14159)) / 1e10))) * ((-2.25 / 1.5) * (float32(float64(-2.25)) + 1e10))) - (((1e10 / (1.5 + ((0.1 * 1.5) + 1e10))) - -2.25) + -2.25))) {
                            v63 := GenS7{F0: v2}
                            var v64 int32 = (-(v37))
                            var v65 float32 = 1.5
                            _ = v63
                            _ = v64
                            _ = v65
                        } else {
                            if v62 {
                                v18 += v3
                                v66 := GenS1{F0: 87}
                                _ = v66
                            } else {
                                v67 := GenS10{F0: 559, F1: 2}
                                var v68 string = (v2 + "hello")
                                v62 = (((0.0 + 0.1) * (1.5 / -2.25)) < (-2.25 - -2.25))
                                _ = v67
                                _ = v68
                            }
                        }
                        _ = v61
                        _ = v62
                    case "hello":
                        if (4294967295 != 1) {
                            fmt.Println(fi36, v3, v3)
                            var v69 int8 = (v3 ^ v18)
                            var v70 int16 = -32768
                            v71 := []int{128, (-(fi34)), (v38 / ((fi36) | 1)), (fi36 - (fi36 / ((2147483647) | 1)))}
                            fmt.Println(fi34, v69)
                            var v72 int = (fi36 ^ fi34)
                            _ = v69
                            _ = v70
                            _ = v71
                            _ = v72
                        }
                    }
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            default:
                v73 := []int{(fi34 + (fi34 & fi34)), fi34, (fi34 << ((1) & 63))}
                _ = v73
            }
        }
        _ = fi34
    }
    if true {
        for fi74 := 0; fi74 < 16; fi74++ {
            if false {
                var v75 string = "a"
                _ = v75
            } else {
                var v76 string = "foo bar"
                v77 := GenS4{F0: v76, F1: 653}
                v2 = "a"
                _ = v76
                _ = v77
            }
            var v78 string = v2
            _ = fi74
            _ = v78
        }
        for fi79 := 0; fi79 < 13; fi79++ {
            v80 := GenS11{F0: 1}
            fmt.Println(v2, v19, v19)
            _ = fi79
            _ = v80
        }
        if (true && (!(65535 != 380))) {
            for fi81 := 0; fi81 < 3; fi81++ {
                _ = len(v1)
                if (146 < (^(v19))) {
                    v18 += v3
                } else {
                    var v82 int32 = 792
                    var v83 uint32 = 0
                    v18 = v3
                    _ = v82
                    _ = v83
                }
                if false {
                    v84 := GenS12{F0: ((v19 <= (v19 & v19)) && true), F1: v2, F2: 926}
                    fmt.Println(v18, v18)
                    _ = v84
                } else {
                    fmt.Println(v2, v19, v18)
                    for fi85 := 0; fi85 < 12; fi85++ {
                        v18 = (-(v18))
                        v18 = (v3 + 127)
                        fmt.Println(v19, v3, v2)
                        _ = fi85
                    }
                    if false {
                        v86 := GenS5{F0: 42}
                        _ = v86
                    } else {
                        for fi87 := 0; fi87 < 3; fi87++ {
                            if (!false) {
                                fmt.Println(fi81, fi81)
                                if false {
                                    v18 += (v3 + 0)
                                } else {
                                    switch (fi87 | (-(fi81))) {
                                    case 0:
                                        v3 -= v18
                                    default:
                                        switch (fi87 % (((fi81 >> ((0) & 63))) | 1)) {
                                        case 835:
                                            if (-270 >= 0) {
                                                v3 = (^(v18))
                                                v3 -= v3
                                            } else {
                                                var v88 uint64 = 18446744073709551615
                                                _ = v88
                                            }
                                        default:
                                            v2 = v2
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(fi81)
                            }
                            _ = fi87
                        }
                        fmt.Println(fi81)
                    }
                }
                _ = fi81
            }
            var v89 float32 = (1e10 * 3.14159)
            _ = v89
        }
    } else {
        var v90 bool = false
        fmt.Println(v18, v3, v90)
        v91 := GenS12{F0: (!((v2 + v2) < "")), F1: ("" + ("hello" + v2)), F2: 923}
        _ = v90
        _ = v91
    }
    fmt.Println(v18, v18, v3)
    for fi92 := 0; fi92 < 8; fi92++ {
        if false {
            for fi93 := 0; fi93 < 7; fi93++ {
                var v94 uint = 1
                v2 = (v2 + (((v2 + v2) + v2) + v2))
                v2 = v2
                _ = fi93
                _ = v94
            }
            for fi95 := 0; fi95 < 1; fi95++ {
                v19 = (v19 + v19)
                v96 := GenS1{F0: 0}
                var v97 int32 = int32(fi92)
                switch (v2 + v2) {
                case "":
                    if true {
                        v2 = v2
                        v98 := GenS13{F0: (v19 * 51), F1: 42, F2: (v2 + v2)}
                        if true {
                            _ = len(v4)
                            var v99 int16 = 1000
                            var v100 uint = 18446744073709551615
                            _ = v100
                            _ = v99
                        }
                        if (true || true) {
                            if true {
                                v97 = (v97 << ((4294967295) & 63))
                                if (!false) {
                                    v101 := GenS14{F0: false}
                                    v102 := GenS15{F0: 856}
                                    _ = v101
                                    _ = v102
                                } else {
                                    var v103 uint32 = 2
                                    _ = v103
                                }
                            } else {
                                var v104 int16 = 32767
                                var v105 string = ((v2 + "hello") + ((v2 + "hello") + v2))
                                _ = v104
                                _ = v105
                            }
                            var v106 int8 = v3
                            v107 := []int{(fi95 << ((0) & 63))}
                            _ = v106
                            _ = v107
                        }
                        _ = v98
                    }
                default:
                    v108 := GenS7{F0: v2}
                    if false {
                        v109 := GenS16{F0: 177, F1: v97, F2: 1000}
                        _ = v109
                    }
                    _ = v108
                }
                v19 = 127
                _ = fi95
                _ = v96
                _ = v97
            }
            var v110 uint16 = 0
            _ = v110
        }
        for fi111 := 0; fi111 < 2; fi111++ {
            v112 := []int{}
            v1[v2] = (fi111 - (fi111 - 9223372036854775807))
            _ = fi111
            _ = v112
        }
        _ = fi92
    }
    if true {
        v113 := GenS6{F0: true}
        _ = v113
    } else {
        var v114 uint32 = 2147483647
        var v115 uint = 708
        _ = v114
        _ = v115
    }
    fmt.Println(v19, v19, v18)
    var v116 string = v2
    if (((1.5 + ((((float64(0.0) / -2.25) * 0.0) - float64(1e10)) / ((-2.25 / 1.0) / 1.5))) + -2.25) >= -2.25) {
        for fi117 := 0; fi117 < 11; fi117++ {
            for fi118 := 0; fi118 < 16; fi118++ {
                v119 := []int{-3}
                _ = fi118
                _ = v119
            }
            v19 = 250
            _ = fi117
        }
        var v120 float32 = 1e10
        _ = len(v1)
        _ = v120
    }
    v121 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v121:", len(v121), v121)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v116
    _ = v121
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
