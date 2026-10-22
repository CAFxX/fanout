package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 string; F1 uint64; F2 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int16; F2 int64 }
type GenS10 struct { F0 string; F1 uint32; F2 int64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16; F1 uint16 }
type GenS14 struct { F0 int16; F1 uint16 }
type GenS15 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS16 struct { F0 int8; F1 int8; F2 int8 }
type GenS17 struct { F0 int; F1 string }
func genFn0(p0 string) bool {
    var v7 float64 = 1.5
    switch p0 {
    case "hello":
        v8 := GenS1{F0: 255}
        _ = v8
    case "foo bar":
        p0 = ("\n\t" + "\n\t")
    }
    var v9 uint = 1
    var v10 uint = 4294967295
    if false {
        switch 128 {
        case -809:
            var v11 uint = v10
            _ = v11
        default:
            v12 := []int{127, 1, 127, 127}
            _ = v12
        }
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v9
    return true
}

func main() {
    var v1 uint16 = 449
    v2 := make(map[string]int)
    v3 := []int{9223372036854775807, 135, -365, 32768}
    var v4 uint8 = 2
    var v5 string = ""
    v6 := make(map[uint32]bool)
    v13 := GenS1{F0: 255}
    v6[494] = false
    v5 = (v5 + (v5 + (("a" + "hello") + v5)))
    var v14 uint64 = 4294967295
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v4)
        for fi16 := 0; fi16 < 19; fi16++ {
            if (!(!((65535 <= 255) || true))) {
                switch (fi16 & 4294967295) {
                case 1:
                    v1 = 255
                case -510:
                    v1 += (v1 + (v1 ^ v1))
                case 808:
                    var v17 int64 = 0
                    _ = v17
                default:
                    fmt.Println(v4, fi16)
                }
                var v18 uint8 = (v4 | 108)
                _ = v18
            } else {
                switch (fi15 << ((854) & 63)) {
                case 2147483647:
                    var v19 uint32 = 52
                    _ = v19
                case 32768:
                    switch "a" {
                    case "\n\t":
                        switch v5 {
                        case "":
                            var v20 int8 = -121
                            _ = len(v3)
                            switch v5 {
                            case "\n\t":
                                if (!(float64(1.5) >= 0.0)) {
                                    if (!false) {
                                        v21 := GenS2{F0: false, F1: "hello"}
                                        v22 := GenS3{F0: (v4 << ((0) & 63))}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v2[("a" + (v5 + v5))] = fi15
                                        v20 += v20
                                    }
                                } else {
                                    v23 := GenS1{F0: 1}
                                    _ = v23
                                }
                            case "hello":
                                v4 += v4
                            case "foo bar":
                                v3 = append(v3, 756)
                            }
                            fmt.Println(v20)
                            var v24 int16 = int16(v4)
                            v1 += v1
                            if false {
                                var v25 string = v5
                                v26 := GenS1{F0: 2}
                                v5 = "foo bar"
                                _ = v25
                                _ = v26
                            } else {
                                _ = len(v3)
                                var v27 int16 = 54
                                var v28 uint32 = 42
                                _ = v27
                                _ = v28
                            }
                            _ = v20
                            _ = v24
                        }
                    default:
                        if (0 <= 1000) {
                            v4 -= v4
                        } else {
                            switch (fi16 * 42) {
                            case 4294967295:
                                if (false || true) {
                                    var v29 int32 = -480
                                    _ = v29
                                } else {
                                    v1 += 42
                                    var v30 uint8 = v4
                                    v31 := GenS4{F0: 42, F1: -1}
                                    v32 := GenS5{F0: 0, F1: v5}
                                    v33 := GenS6{F0: 18446744073709551615, F1: 32767}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            case 128:
                                v34 := GenS7{F0: v5, F1: 1, F2: -1}
                                _ = v34
                            case 32768:
                                fmt.Println(v4)
                            }
                            var v35 uint8 = (v4 | uint8(v1))
                            _ = v35
                        }
                    }
                case 32767:
                    v3 = append(v3, (fi15 & fi15))
                default:
                    var v36 uint8 = v4
                    _ = v36
                }
                if (0 < 65535) {
                    var v37 bool = ((!((("\n\t" + (v5 + "a")) + v5) != "")) && ((^(v1)) > (v1 & 0)))
                    switch "hello" {
                    case "\n\t":
                        switch 42 {
                        case 128:
                            if false {
                                if (!v37) {
                                    fmt.Println(v4, fi16)
                                }
                                v1 -= v1
                            }
                        case 15:
                            fmt.Println(v4)
                        default:
                            var v38 uint8 = (v4 * v4)
                            _ = v38
                        }
                    }
                    _ = v37
                } else {
                    var v39 int = fi15
                    switch fi16 {
                    case 2147483647:
                        var v40 bool = ((127 >= 100) || (218 == 42))
                        _ = v40
                    case 0:
                        var v41 uint8 = 0
                        _ = v41
                    default:
                        switch (fi16 - (fi16 - -9223372036854775808)) {
                        case -731:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 42:
                            if (!false) {
                                v42 := GenS5{F0: 100, F1: ""}
                                var v43 int = 127
                                _ = v42
                                _ = v43
                            }
                        default:
                            switch v39 {
                            case -9223372036854775808:
                                var v44 int8 = -128
                                _ = v44
                            case 0:
                                switch (v39 + (fi16 + -864)) {
                                case 4294967295:
                                    var v45 uint = 1
                                    _ = v45
                                }
                            case 1:
                                fmt.Println(v4)
                            default:
                                v14 += (v14 + 2)
                            }
                        }
                    }
                    if (1 >= 694) {
                        if (-694 == -519) {
                            fmt.Println(v14, v4)
                            var v46 uint = 42
                            if (fi16 == (v39 >> ((255) & 63))) {
                                v4 = v4
                                v47 := GenS6{F0: v14, F1: v39}
                                v48 := GenS8{F0: 0}
                                switch (fi15 << ((878) & 63)) {
                                case -1:
                                    fmt.Println(v1)
                                case -184:
                                    v46 = (v46 | (v46 ^ (v46 + v46)))
                                case 42:
                                    switch fi16 {
                                    case -2147483648:
                                        v14 = (v14 % ((v14) | 1))
                                    case 91:
                                        v1 += v1
                                        var v49 uint8 = v4
                                        if false {
                                            v50 := GenS9{F0: uint8(v1), F1: 32767, F2: 42}
                                            _ = v50
                                        }
                                        _ = v49
                                    case 32768:
                                        var v51 uint = uint(v4)
                                        _ = v51
                                    default:
                                        if ((!(1 >= 127)) || true) {
                                            if len(v3) > 0 {
                                                v3[0] = v39
                                            }
                                        }
                                    }
                                default:
                                    var v52 uint64 = v14
                                    _ = v52
                                }
                                _ = v47
                                _ = v48
                            }
                            _ = v46
                        } else {
                            v53 := GenS10{F0: v5, F1: 2, F2: 9223372036854775807}
                            v54 := GenS11{F0: v1}
                            _ = v53
                            _ = v54
                        }
                    }
                    if false {
                        var v55 bool = (!(v4 != uint8(v1)))
                        _ = v55
                    }
                    v5 = (v5 + v5)
                    _ = v39
                }
                if ((v1 / ((v1) | 1)) < v1) {
                    if false {
                        v14 = (v14 * v14)
                        v56 := GenS12{F0: false}
                        _ = v56
                    }
                    var v57 uint64 = (^(v14))
                    if (true && true) {
                        if (!true) {
                            _ = len(v2)
                            var v58 uint32 = 0
                            _ = v58
                        } else {
                            v59 := GenS7{F0: (v5 + "foo bar"), F1: (^(v57)), F2: -1}
                            v14 = 1
                            _ = v59
                        }
                        if false {
                            if (!false) {
                                v57 = (^(v57))
                                v6[2147483648] = (false && true)
                            } else {
                                if (18446744073709551615 < 361) {
                                    if (-2.25 != 0.1) {
                                        switch v5 {
                                        case "foo bar":
                                            var v60 uint16 = (^(v1))
                                            _ = v60
                                        case "\n\t":
                                            var v61 int32 = 912
                                            _ = v61
                                        default:
                                            var v62 uint = 255
                                            _ = v62
                                        }
                                        if false {
                                            switch (fi16 << ((647) & 63)) {
                                            case -2147483648:
                                                var v63 int16 = -1
                                                _ = v63
                                            }
                                            var v64 uint16 = v1
                                            _ = v64
                                        } else {
                                            v2[((v5 + "foo bar") + v5)] = fi15
                                            var v65 bool = (false && false)
                                            v66 := GenS13{F0: 1, F1: 42}
                                            v14 -= uint64(v1)
                                            var v67 int = fi15
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = fi16
                                        }
                                    }
                                    v68 := GenS12{F0: (0 != (v57 / ((825) | 1)))}
                                    _ = v68
                                }
                                v3 = append(v3, 32767)
                                fmt.Println(v1, v57, fi16)
                                if (0 != 2147483647) {
                                    fmt.Println(v1, fi15, v1)
                                    v3 = append(v3, 815)
                                } else {
                                    if true {
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = (fi16 & (fi16 - 128))
                                            }
                                            var v69 uint16 = v1
                                            var v70 int64 = 199
                                            _ = v69
                                            _ = v70
                                        }
                                        v6[0] = false
                                    } else {
                                        if true {
                                            var v71 uint = 18446744073709551615
                                            _ = v71
                                        } else {
                                            v1 = (^(v1))
                                        }
                                    }
                                }
                                fmt.Println(v1, v14, v14)
                                if true {
                                    v72 := GenS5{F0: 127, F1: v5}
                                    v57 -= 18446744073709551615
                                    v5 = v5
                                    _ = v72
                                } else {
                                    var v73 uint16 = 204
                                    _ = v73
                                }
                                var v74 bool = ((0 == 32767) || ((v5 >= "\n\t") || true))
                                _ = v74
                            }
                            v1 = (^(v1))
                        }
                    } else {
                        v5 = ""
                        v4 = (v4 - v4)
                    }
                    _ = v57
                }
                fmt.Println(v4)
            }
            _ = fi16
        }
        v14 += v14
        _ = fi15
    }
    v75 := GenS3{F0: 171}
    v5 = (v5 + "hello")
    for fi76 := 0; fi76 < 3; fi76++ {
        if true {
            if true {
                var v77 uint8 = (v4 + (v4 | (v4 % (((^(v4))) | 1))))
                for fi78 := 0; fi78 < 6; fi78++ {
                    var v79 uint16 = 255
                    var v80 bool = ((1e10 <= 0.0) || true)
                    _ = fi78
                    _ = v79
                    _ = v80
                }
                fmt.Println(v14, v5)
                _ = v77
            } else {
                v4 = v4
                if false {
                    var v81 uint16 = 2
                    for fi82 := 0; fi82 < 14; fi82++ {
                        if (v4 > v4) {
                            v5 = v5
                        }
                        v83 := GenS14{F0: 32767, F1: v81}
                        v4 -= (v4 / (((v4 - (^(v4)))) | 1))
                        _ = fi82
                        _ = v83
                    }
                    var v84 int16 = 0
                    _ = v81
                    _ = v84
                }
                if (!(4294967295 <= 42)) {
                    var v85 uint = 4294967295
                    var v86 float64 = -2.25
                    for fi87 := 0; fi87 < 16; fi87++ {
                        v88 := GenS15{F0: v4, F1: -2147483648, F2: 2}
                        v89 := GenS5{F0: 0, F1: v5}
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                    _ = v85
                    _ = v86
                }
            }
            for fi90 := 0; fi90 < 7; fi90++ {
                v4 += v4
                v14 += v14
                switch (("hello" + "\n\t") + ("a" + (v5 + "a"))) {
                case "hello":
                    v6[822] = ((((32767 < -981) || (!true)) || (!true)) && (-2147483648 <= 584))
                case "foo bar":
                    var v91 uint = 255
                    _ = v91
                case "":
                    v2[v5] = (fi90 - (^(fi90)))
                default:
                    fmt.Println(fi90, fi90, v5)
                }
                _ = fi90
            }
            var v92 uint64 = 18446744073709551615
            _ = v92
        } else {
            var v93 int64 = -753
            v94 := GenS16{F0: -1, F1: 100, F2: 42}
            for fi95 := 0; fi95 < 8; fi95++ {
                switch ("\n\t" + (v5 + v5)) {
                case "\n\t":
                    var v96 float32 = 0.0
                    _ = v96
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch v5 {
                    case "a":
                        var v97 string = v5
                        _ = v97
                    case "foo bar":
                        if (false || (uint16(v14) <= (v1 & (^(v1))))) {
                            var v98 int32 = 2147483647
                            v99 := GenS14{F0: -1, F1: (v1 ^ v1)}
                            fmt.Println(fi95, v93, fi95)
                            fmt.Println(fi95, v93)
                            var v100 int32 = v98
                            _ = v100
                            _ = v98
                            _ = v99
                        } else {
                            var v101 int32 = 1
                            v102 := GenS17{F0: fi76, F1: ""}
                            v93 = v93
                            _ = v101
                            _ = v102
                        }
                    default:
                        var v103 uint16 = 865
                        _ = v103
                    }
                case "foo bar":
                    var v104 bool = (true || (-27 != 127))
                    _ = v104
                default:
                    v105 := GenS13{F0: 1, F1: 181}
                    _ = v105
                }
                delete(v6, 42)
                _ = fi95
            }
            _ = v93
            _ = v94
        }
        var v106 float32 = (float32(1.5) * (0.1 - 1e10))
        v5 = ""
        _ = fi76
        _ = v106
    }
    v5 = "a"
    fmt.Println(v4, v4)
    for fi107 := 0; fi107 < 7; fi107++ {
        v108 := GenS16{F0: 114, F1: -104, F2: -1}
        _ = fi107
        _ = v108
    }
    var v109 float64 = 0.0
    var v110 bool = (1 >= 40)
    if (2 > v14) {
        for fi111 := 0; fi111 < 17; fi111++ {
            v3 = append(v3, (-(fi111)))
            v109 = v109
            _ = fi111
        }
        for fi112 := 0; fi112 < 9; fi112++ {
            v5 = "hello"
            v14 = 230
            _ = fi112
        }
        var v113 int32 = 123
        _ = v113
    } else {
        var v114 int32 = 1
        _ = v114
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v109
    _ = v110
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
}
