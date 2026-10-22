package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 int64 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int16; F1 int8; F2 bool }
type GenS7 struct { F0 uint16; F1 int32; F2 string }
type GenS8 struct { F0 int64; F1 string; F2 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 int64; F1 int8; F2 int16 }
type GenS12 struct { F0 bool; F1 int16; F2 int64 }
type GenS13 struct { F0 string; F1 uint64 }
func genFn0(p0 int16) uint16 {
    if true {
        fmt.Println(p0)
        var v5 uint32 = 42
        _ = v5
    }
    p0 += p0
    _ = p0
    return 611
}
func genFn1(p0 uint16) uint {
    v6 := GenS1{F0: -1, F1: 0, F2: 1}
    v7 := GenS2{F0: 266, F1: 32767}
    if false {
        if false {
            switch "hello" {
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    var v8 float64 = (-2.25 * float64(-2.25))
                    _ = v8
                case "hello":
                    switch ("\n\t" + ("" + ("a" + "foo bar"))) {
                    case "":
                        fmt.Println(p0)
                    default:
                        v9 := []int{}
                        _ = v9
                    }
                default:
                    if (false || ((805 > 4294967295) || false)) {
                        var v10 uint = 65535
                        _ = v10
                    }
                }
            case "":
                var v11 uint32 = 42
                _ = v11
            case "foo bar":
                if (196 != 2) {
                    v12 := GenS2{F0: 42, F1: 1}
                    var v13 uint32 = 42
                    _ = v12
                    _ = v13
                }
            default:
                p0 -= (p0 ^ p0)
            }
            v14 := GenS3{F0: 107, F1: 0}
            if false {
                if true {
                    if false {
                        var v15 bool = false
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v17 := GenS2{F0: -981, F1: (p0 >> ((2) & 63))}
                            v18 := GenS2{F0: 42, F1: 2}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS3{F0: -2147483648, F1: p0}
                        _ = v15
                        _ = v19
                    } else {
                        var v20 uint8 = 2
                        switch "" {
                        case "hello":
                            v21 := []int{9223372036854775807}
                            _ = v21
                        case "a":
                            var v22 uint8 = v20
                            _ = v22
                        case "foo bar":
                            var v23 int8 = 74
                            _ = v23
                        }
                        _ = v20
                    }
                    var v24 int64 = -9223372036854775808
                    var v25 uint8 = 127
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    fmt.Println(fi26)
                    _ = fi26
                }
                var v27 uint = 982
                _ = v27
            }
            _ = v14
        } else {
            p0 += (p0 + p0)
            v28 := GenS4{F0: 4294967295, F1: ("hello" == "a"), F2: 1}
            _ = v28
        }
        if true {
            if false {
                var v29 uint64 = 188
                var v30 int8 = 127
                _ = v29
                _ = v30
            } else {
                switch -9223372036854775808 {
                case 0:
                    var v31 int64 = 1
                    _ = v31
                case 32768:
                    p0 -= 2
                default:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        p0 = (^(p0))
                        if false {
                            for fi33 := 0; fi33 < 10; fi33++ {
                                if (!false) {
                                    var v34 string = "foo bar"
                                    _ = v34
                                }
                                v35 := []int{(fi33 & (fi32 | fi32))}
                                switch "" {
                                case "foo bar":
                                    p0 += (p0 + (p0 * (p0 ^ (^(p0)))))
                                case "\n\t":
                                    var v36 uint16 = (p0 / ((p0) | 1))
                                    _ = v36
                                }
                                _ = fi33
                                _ = v35
                            }
                        }
                        if ((!true) && true) {
                            p0 = (p0 >> ((4294967295) & 63))
                            if (-935 > -2147483648) {
                                v37 := GenS5{F0: -1, F1: true}
                                for fi38 := 0; fi38 < 8; fi38++ {
                                    fmt.Println(fi32)
                                    v39 := []int{(fi38 >> ((255) & 63))}
                                    switch (^(fi38)) {
                                    case 4294967295:
                                        p0 = 293
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                                _ = v37
                            } else {
                                p0 -= genFn0(783)
                                var v40 uint64 = 119
                                p0 = (p0 * 255)
                                _ = v40
                            }
                        }
                        _ = fi32
                    }
                }
            }
            fmt.Println(p0)
        }
    } else {
        p0 = (p0 ^ p0)
        v41 := GenS2{F0: 1000, F1: (p0 >> ((404) & 63))}
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 float32 = (3.14159 / (-2.25 * 0.1))
            v44 := GenS6{F0: 32767, F1: 100, F2: true}
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 65535
    var v2 int16 = 42
    var v3 bool = (-1 != 42)
    var v4 string = "\n\t"
    fmt.Println(v4, v1, v1)
    v1 = genFn0(v2)
    for fi45 := 0; fi45 < 17; fi45++ {
        var v46 bool = (!((-513 != -852) || true))
        _ = fi45
        _ = v46
    }
    if (42 >= -128) {
        v1 += v1
    }
    v47 := []int{-1, -817, 478, -417}
    v1 += v1
    v3 = v3
    if v3 {
        var v48 int32 = 320
        v49 := GenS7{F0: genFn0(v2), F1: (v48 + v48), F2: v4}
        switch 32767 {
        case 42:
            if false {
                switch "a" {
                case "":
                    v1 = 32767
                case "\n\t":
                    var v50 string = v4
                    _ = v50
                default:
                    v51 := GenS8{F0: int64(v1), F1: v4, F2: -691}
                    _ = v51
                }
                v52 := GenS7{F0: genFn0(v2), F1: int32(v1), F2: v4}
                _ = v52
            }
        }
        _ = v48
        _ = v49
    } else {
        v53 := GenS9{F0: 42}
        fmt.Println(v3, v1, v2)
        fmt.Println(v4, v3, v1)
        _ = v53
    }
    if false {
        for fi54 := 0; fi54 < 3; fi54++ {
            var v55 string = "hello"
            switch v55 {
            case "":
                if (711 > 2) {
                    v56 := GenS1{F0: 2147483647, F1: 5, F2: -9223372036854775808}
                    var v57 int = int(v2)
                    if (true || true) {
                        _ = len(v47)
                    } else {
                        v47 = append(v47, (v57 - fi54))
                        var v58 uint8 = 1
                        if false {
                            if len(v47) > 0 {
                                v47[0] = (fi54 * v57)
                            }
                            var v59 uint16 = v1
                            switch fi54 {
                            case 127:
                                var v60 bool = ((v58 % ((v58) | 1)) < v58)
                                _ = v60
                            }
                            _ = v59
                        } else {
                            var v61 bool = v3
                            if (v57 < v57) {
                                var v62 uint64 = 0
                                if len(v47) > 0 {
                                    v47[0] = fi54
                                }
                                _ = v62
                            }
                            v3 = (v58 > 127)
                            _ = v61
                        }
                        _ = v58
                    }
                    _ = v56
                    _ = v57
                } else {
                    v63 := GenS6{F0: v2, F1: -79, F2: ((("a" + (v4 + v4)) > v55) || (v3 && (!v3)))}
                    v64 := GenS10{F0: 1, F1: uint(v1)}
                    if (904 > v1) {
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                        v1 -= v1
                        for fi65 := 0; fi65 < 20; fi65++ {
                            v66 := GenS7{F0: genFn0(0), F1: -2147483648, F2: "\n\t"}
                            fmt.Println(v3)
                            _ = fi65
                            _ = v66
                        }
                    } else {
                        for fi67 := 0; fi67 < 11; fi67++ {
                            if false {
                                if v3 {
                                    var v68 int8 = -1
                                    v69 := GenS11{F0: int64(v1), F1: v68, F2: 1000}
                                    var v70 int = (fi54 & -9223372036854775808)
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                } else {
                                    switch 9223372036854775807 {
                                    case -2147483648:
                                        var v71 string = v4
                                        v2 -= (v2 | v2)
                                        var v72 uint8 = 222
                                        _ = v71
                                        _ = v72
                                    case 9223372036854775807:
                                        var v73 int16 = 32767
                                        _ = v73
                                    }
                                    if len(v47) >= 2 {
                                        v47 = v47[1:len(v47)]
                                    }
                                }
                                if len(v47) > 0 {
                                    v47[0] = (fi54 - (-(fi67)))
                                }
                                var v74 int64 = 9223372036854775807
                                var v75 uint8 = 127
                                var v76 float64 = (3.14159 + (1.5 / ((1e10 + (0.1 + (float64(-2.25) + (float64(0.1) + ((1e10 - (0.0 - 1.5)) + 1e10))))) * (0.0 - 1e10))))
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                            _ = fi67
                        }
                        for fi77 := 0; fi77 < 14; fi77++ {
                            v4 = "foo bar"
                            var v78 float64 = 3.14159
                            _ = fi77
                            _ = v78
                        }
                        v1 += (v1 & 190)
                    }
                    _ = v63
                    _ = v64
                }
            }
            var v79 uint = 145
            _ = fi54
            _ = v55
            _ = v79
        }
    } else {
        var v80 int64 = 981
        switch -660 {
        case -1:
            switch "foo bar" {
            case "foo bar":
                if v3 {
                    fmt.Println(v2)
                    v3 = (!(!(!false)))
                    var v81 int32 = -2147483648
                    _ = v81
                } else {
                    v3 = v3
                    var v82 int32 = -1
                    _ = v82
                }
            case "":
                fmt.Println(v3, v4)
            case "hello":
                if len(v47) >= 2 {
                    v47 = v47[1:len(v47)]
                }
            default:
                var v83 string = (((v4 + "hello") + "\n\t") + "hello")
                _ = v83
            }
        }
        switch 4294967295 {
        case 159:
            var v84 uint16 = (v1 << ((genFn1((v1 ^ 255))) & 63))
            _ = v84
        case 2147483647:
            if (!((v80 - (v80 & -389)) <= (v80 << ((607) & 63)))) {
                if (v80 <= (v80 & v80)) {
                    v80 -= v80
                } else {
                    if v3 {
                        v85 := GenS4{F0: 18446744073709551615, F1: v3, F2: 9223372036854775807}
                        var v86 int64 = v80
                        if v3 {
                            for fi87 := 0; fi87 < 6; fi87++ {
                                v80 = -9223372036854775808
                                _ = fi87
                            }
                            v80 = v86
                            _ = len(v47)
                        }
                        _ = v85
                        _ = v86
                    }
                    v47 = append(v47, 32768)
                    v4 = "a"
                }
                var v88 float32 = 1.5
                _ = v88
            }
        }
        _ = v80
    }
    switch v4 {
    case "":
        _ = len(v47)
    case "\n\t":
        var v89 uint32 = 879
        _ = v89
    default:
        for fi90 := 0; fi90 < 2; fi90++ {
            if (!(414 > 4294967295)) {
                var v91 int64 = 9223372036854775807
                if ((-1 > -1) && true) {
                    if (42 == 475) {
                        fmt.Println(v1, fi90, v1)
                        var v92 bool = v3
                        v93 := GenS5{F0: (-(v91)), F1: v3}
                        _ = v92
                        _ = v93
                    } else {
                        switch v4 {
                        case "foo bar":
                            switch (-(fi90)) {
                            case -2147483648:
                                v3 = (1 == 221)
                            case -672:
                                var v94 int64 = (-(v91))
                                _ = v94
                            }
                        default:
                            for fi95 := 0; fi95 < 5; fi95++ {
                                fmt.Println(v1, fi90)
                                v47 = append(v47, fi90)
                                _ = fi95
                            }
                        }
                        fmt.Println(v4, v2, v4)
                        if false {
                            if v3 {
                                switch (fi90 + fi90) {
                                case -792:
                                    var v96 uint = 255
                                    _ = v96
                                case -9223372036854775808:
                                    switch (fi90 / ((-1) | 1)) {
                                    case 2147483647:
                                        if len(v47) > 0 {
                                            v47[0] = 963
                                        }
                                    default:
                                        var v97 int = fi90
                                        _ = v97
                                    }
                                case 4294967295:
                                    v98 := GenS12{F0: true, F1: 0, F2: (^(v91))}
                                    _ = v98
                                }
                                v99 := GenS13{F0: v4, F1: 228}
                                v47 = append(v47, (fi90 + (-(fi90))))
                                _ = v99
                            }
                            switch 2147483647 {
                            case 42:
                                for fi100 := 0; fi100 < 8; fi100++ {
                                    switch fi100 {
                                    case -104:
                                        var v101 uint16 = (^(v1))
                                        _ = v101
                                    case -364:
                                        v102 := GenS8{F0: v91, F1: "\n\t", F2: 32767}
                                        _ = v102
                                    }
                                    var v103 int16 = v2
                                    _ = fi100
                                    _ = v103
                                }
                            default:
                                v47 = append(v47, int(v91))
                            }
                        }
                    }
                    for fi104 := 0; fi104 < 15; fi104++ {
                        if len(v47) > 0 {
                            v47[0] = (fi90 * 500)
                        }
                        _ = fi104
                    }
                } else {
                    for fi105 := 0; fi105 < 18; fi105++ {
                        var v106 uint64 = 241
                        _ = fi105
                        _ = v106
                    }
                }
                v3 = (v3 || ("hello" < "\n\t"))
                _ = v91
            }
            var v107 uint64 = 4294967295
            _ = fi90
            _ = v107
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
}
