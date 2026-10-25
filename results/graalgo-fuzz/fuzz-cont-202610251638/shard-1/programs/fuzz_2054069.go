package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 int32; F2 int32 }
type GenS7 struct { F0 int32; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS9 struct { F0 uint8; F1 int; F2 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 string; F1 string }
func genFn0(p0 uint64) uint {
    if (true || (1 > 42)) {
        var v5 uint8 = 127
        if false {
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: -2147483648}
                var v8 float64 = float64(0.0)
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 int64 = -638
            _ = v9
        } else {
            p0 += p0
            p0 -= 1
            p0 = 1
        }
        if true {
            if (((float32(3.14159) > 1.5) || ((((!true) && (true && true)) || (!((485 < 0) && true))) || (42 <= -9223372036854775808))) && (!false)) {
                v5 += (v5 + uint8(p0))
            } else {
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 float64 = 1.5
                    _ = fi10
                    _ = v11
                }
                v12 := []int{128, 32768, -2147483648}
                var v13 int16 = -1
                _ = v12
                _ = v13
            }
            switch "" {
            case "":
                switch 1 {
                case 42:
                    switch 9223372036854775807 {
                    case -1:
                        for fi14 := 0; fi14 < 8; fi14++ {
                            p0 -= 42
                            var v15 uint32 = 965
                            var v16 bool = ((int32(v15) > -219) || (p0 < (p0 / (((p0 + p0)) | 1))))
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    case 42:
                        v17 := []int{}
                        _ = v17
                    default:
                        v18 := GenS1{F0: -773}
                        _ = v18
                    }
                case -867:
                    fmt.Println(v5, v5)
                }
            case "hello":
                fmt.Println(p0)
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                v5 = v5
                var v20 string = "hello"
                if (681 <= 1) {
                    var v21 uint = 42
                    var v22 uint64 = p0
                    var v23 int = (fi19 & fi19)
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
        } else {
            var v24 uint64 = p0
            var v25 uint = 42
            for fi26 := 0; fi26 < 14; fi26++ {
                switch fi26 {
                case 1:
                    var v27 uint64 = p0
                    _ = v27
                case 32768:
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 float32 = 0.1
                        switch "\n\t" {
                        case "foo bar":
                            v5 = 64
                        default:
                            v30 := GenS1{F0: -2147483648}
                            _ = v30
                        }
                        v31 := []int{32767, (fi28 * (fi28 & fi28)), 4294967295, (fi28 >> (((v25 | (^(v25)))) & 63))}
                        _ = fi28
                        _ = v29
                        _ = v31
                    }
                case 128:
                    if false {
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            fmt.Println(v25)
                        default:
                            for fi32 := 0; fi32 < 10; fi32++ {
                                v33 := []int{32768, 32767}
                                v25 += (v25 / ((956) | 1))
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                default:
                    var v34 uint32 = uint32(v24)
                    _ = v34
                }
                _ = fi26
            }
            _ = v24
            _ = v25
        }
        _ = v5
    }
    var v35 int32 = 2147483647
    if (v35 < (v35 >> ((1) & 63))) {
        var v36 string = ""
        var v37 float64 = (3.14159 / (1e10 * 1.5))
        _ = v36
        _ = v37
    } else {
        if ((!(18446744073709551615 > 2)) && ((float32(3.14159) + 1e10) <= 0.0)) {
            v35 = (v35 % ((-462) | 1))
            v35 = (-(v35))
        } else {
            switch 127 {
            case 213:
                v38 := GenS2{F0: 156}
                _ = v38
            case 32767:
                fmt.Println(p0, p0)
            case 956:
                if (4294967295 < 4294967295) {
                    for fi39 := 0; fi39 < 10; fi39++ {
                        var v40 float32 = (3.14159 + 0.0)
                        _ = fi39
                        _ = v40
                    }
                    v35 = 1
                    for fi41 := 0; fi41 < 4; fi41++ {
                        p0 -= p0
                        v42 := []int{}
                        _ = fi41
                        _ = v42
                    }
                } else {
                    var v43 uint16 = 65535
                    _ = v43
                }
            default:
                p0 = p0
            }
            v35 = v35
            var v44 uint16 = 65535
            _ = v44
        }
        v45 := []int{42, 637}
        v46 := GenS3{F0: 8, F1: true}
        _ = v45
        _ = v46
    }
    v47 := []int{-1}
    fmt.Println(p0)
    _ = p0
    _ = v35
    _ = v47
    return 2
}
func genFn1() int64 {
    v48 := []int{9223372036854775807, 32767, -673, 870}
    v48 = append(v48, 0)
    var v49 float64 = -2.25
    var v50 bool = true
    _ = v48
    _ = v49
    _ = v50
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 string = "a"
    var v3 float32 = 0.1
    var v4 float64 = 0.0
    switch ((v2 + (((v2 + v2) + (("a" + v2) + (v2 + v2))) + (((((v2 + v2) + (v2 + "foo bar")) + "\n\t") + "hello") + (v2 + "")))) + ("foo bar" + (v2 + (v2 + "hello")))) {
    case "":
        var v51 int64 = -191
        _ = v51
    case "a":
        v4 = (v4 - (v4 + v4))
    default:
        v52 := GenS1{F0: -19}
        _ = v52
    }
    var v53 uint32 = 0
    v54 := []int{}
    var v55 int32 = 1
    if len(v54) >= 2 {
        v54 = v54[1:len(v54)]
    }
    if false {
        switch v2 {
        case "\n\t":
            v56 := GenS4{F0: v55, F1: 0}
            _ = v56
        default:
            v57 := GenS5{F0: 2}
            _ = v57
        }
        v58 := GenS3{F0: 255, F1: (!(v3 >= v3))}
        _ = v58
    }
    for fi59 := 0; fi59 < 14; fi59++ {
        v1 = (^(v1))
        var v60 float32 = float32(float64(v3))
        _ = fi59
        _ = v60
    }
    _ = len(v54)
    v61 := GenS5{F0: uint64(v1)}
    if len(v54) >= 2 {
        v54 = v54[1:len(v54)]
    }
    if true {
        v55 += (v55 * v55)
        for fi62 := 0; fi62 < 16; fi62++ {
            for fi63 := 0; fi63 < 19; fi63++ {
                v64 := GenS3{F0: genFn0(42), F1: (v1 > v1)}
                var v65 float32 = v3
                v4 = v4
                _ = fi63
                _ = v64
                _ = v65
            }
            var v66 int8 = 57
            _ = fi62
            _ = v66
        }
        var v67 string = v2
        _ = v67
    } else {
        var v68 uint32 = (v53 >> ((255) & 63))
        switch "a" {
        case "foo bar":
            for fi69 := 0; fi69 < 9; fi69++ {
                _ = len(v54)
                var v70 float32 = ((float32(v4) + (0.0 / -2.25)) + v3)
                if ((fi69 / (((fi69 & (^(fi69)))) | 1)) == fi69) {
                    v70 = (0.1 - (3.14159 + 1.5))
                    if false {
                        v71 := GenS6{F0: 2, F1: (v55 * v55), F2: (v55 * v55)}
                        _ = v71
                    }
                    var v72 int16 = 0
                    _ = v72
                } else {
                    v70 = ((0.1 - v3) / v3)
                    if (0 < 1) {
                        if true {
                            v73 := GenS3{F0: 18446744073709551615, F1: true}
                            v3 = v3
                            var v74 int16 = 1
                            _ = v73
                            _ = v74
                        }
                    } else {
                        v54 = append(v54, (fi69 % ((0) | 1)))
                        if true {
                            v2 = ""
                        } else {
                            v75 := GenS3{F0: 255, F1: false}
                            v76 := GenS5{F0: 568}
                            v54 = append(v54, fi69)
                            _ = v75
                            _ = v76
                        }
                    }
                    if true {
                        for fi77 := 0; fi77 < 5; fi77++ {
                            v78 := GenS6{F0: 0, F1: v55, F2: (^(v55))}
                            fmt.Println(fi69, fi69, fi77)
                            v79 := GenS7{F0: (v55 - v55), F1: 340, F2: -95}
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                        if (v68 == v53) {
                            var v80 uint8 = 42
                            v81 := GenS8{F0: (v1 | v1), F1: 32767, F2: 42}
                            var v82 uint64 = 2
                            _ = v80
                            _ = v81
                            _ = v82
                        }
                    } else {
                        switch 128 {
                        case -734:
                            var v83 uint32 = 0
                            _ = v83
                        default:
                            for fi84 := 0; fi84 < 17; fi84++ {
                                if ((1.5 / v3) > 0.0) {
                                    v53 = (^(v68))
                                    if (v70 >= v70) {
                                        v85 := GenS9{F0: (v1 % (((v1 + 255)) | 1)), F1: fi69, F2: (v55 ^ (-(v55)))}
                                        _ = v85
                                    }
                                    fmt.Println(fi69, fi69, v2)
                                }
                                switch (-(fi69)) {
                                case 128:
                                    switch fi69 {
                                    case 564:
                                        if len(v54) >= 2 {
                                            v54 = v54[1:len(v54)]
                                        }
                                    case 671:
                                        v2 = (v2 + ((v2 + v2) + ("\n\t" + ((("foo bar" + (v2 + v2)) + ("a" + ("\n\t" + v2))) + v2))))
                                    case 2147483647:
                                        v53 += v53
                                    default:
                                        if false {
                                            switch (-(fi84)) {
                                            case 32767:
                                                switch (fi69 / ((fi84) | 1)) {
                                                case 127:
                                                    v54 = append(v54, fi84)
                                                case 4294967295:
                                                    v68 = (^(v53))
                                                default:
                                                    fmt.Println(fi69, fi84)
                                                }
                                            case 632:
                                                v86 := GenS5{F0: 42}
                                                fmt.Println(v70, v1, fi84)
                                                _ = v86
                                            case -1:
                                                if false {
                                                    fmt.Println(v3, v68)
                                                    if true {
                                                        var v87 int8 = 127
                                                        var v88 int64 = 0
                                                        if (32767 < 1000) {
                                                            if (-1 >= -32768) {
                                                                if (false || (!true)) {
                                                                    if len(v54) > 0 {
                                                                        v54[0] = 1
                                                                    }
                                                                } else {
                                                                    if len(v54) >= 2 {
                                                                        v54 = v54[1:len(v54)]
                                                                    }
                                                                    v4 = v4
                                                                    v4 = 1e10
                                                                }
                                                            } else {
                                                                switch 32768 {
                                                                case 806:
                                                                    if len(v54) >= 2 {
                                                                        v54 = v54[1:len(v54)]
                                                                    }
                                                                case 4294967295:
                                                                    switch (-(fi69)) {
                                                                    case 32767:
                                                                        _ = len(v54)
                                                                    case 307:
                                                                        v68 = (v53 & v53)
                                                                    case -9223372036854775808:
                                                                        switch fi69 {
                                                                        case -881:
                                                                            if true {
                                                                                var v89 int64 = v88
                                                                                var v90 uint32 = (v68 | (^(v53)))
                                                                                v91 := GenS8{F0: 0, F1: 32767, F2: v55}
                                                                                _ = v89
                                                                                _ = v90
                                                                                _ = v91
                                                                            } else {
                                                                                v70 += v3
                                                                                fmt.Println(v53, v68)
                                                                                v54 = append(v54, -719)
                                                                            }
                                                                        case 184:
                                                                            var v92 uint64 = 35
                                                                            _ = v92
                                                                        case 122:
                                                                            var v93 int32 = (v55 * v55)
                                                                            _ = v93
                                                                        }
                                                                    default:
                                                                        v88 += (v88 - 0)
                                                                    }
                                                                case 1:
                                                                    v3 += v3
                                                                default:
                                                                    var v94 int8 = v87
                                                                    _ = v94
                                                                }
                                                                fmt.Println(v4)
                                                                var v95 int8 = v87
                                                                fmt.Println(fi69, fi84, v88)
                                                                _ = v95
                                                            }
                                                            v87 -= 86
                                                        } else {
                                                            var v96 bool = ((v87 - (-(v87))) > (v87 << ((2) & 63)))
                                                            _ = v96
                                                        }
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        fmt.Println(v1)
                                                        if false {
                                                            if (0 > 1) {
                                                                if true {
                                                                    v97 := GenS10{F0: 1}
                                                                    var v98 float32 = 0.1
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    var v99 uint16 = 2
                                                                    if len(v54) > 0 {
                                                                        v54[0] = fi84
                                                                    }
                                                                    if true {
                                                                        var v100 uint32 = uint32(v99)
                                                                        v4 = v4
                                                                        var v101 uint8 = v1
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                    _ = v99
                                                                }
                                                            } else {
                                                                var v102 float64 = (v4 / v4)
                                                                v102 += v4
                                                                _ = v102
                                                            }
                                                            var v103 int = fi69
                                                            _ = v103
                                                        } else {
                                                            v104 := GenS6{F0: 65535, F1: 42, F2: 181}
                                                            v1 += v1
                                                            _ = v104
                                                        }
                                                        var v105 uint32 = v68
                                                        switch (^(fi69)) {
                                                        case 50:
                                                            v105 = 838
                                                        }
                                                        _ = v105
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                v106 := GenS11{F0: v55, F1: 687}
                                _ = fi84
                                _ = v106
                            }
                        }
                        v107 := GenS1{F0: (v55 - -2147483648)}
                        _ = v107
                    }
                }
                _ = fi69
                _ = v70
            }
        case "\n\t":
            var v108 string = (v2 + (v2 + ((v2 + v2) + v2)))
            _ = v108
        }
        var v109 string = v2
        _ = v109
        _ = v68
    }
    v4 -= -2.25
    if (true || (!true)) {
        v54 = append(v54, 127)
        for fi110 := 0; fi110 < 7; fi110++ {
            var v111 string = "hello"
            v2 = ""
            if (!true) {
                fmt.Println(v53, v3)
            }
            _ = fi110
            _ = v111
        }
        for fi112 := 0; fi112 < 1; fi112++ {
            switch "a" {
            case "foo bar":
                v113 := GenS12{F0: v2, F1: "\n\t"}
                _ = v113
            case "\n\t":
                switch ((v2 + v2) + v2) {
                case "":
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                case "\n\t":
                    if len(v54) > 0 {
                        v54[0] = fi112
                    }
                case "hello":
                    v2 = "hello"
                }
            default:
                v53 = (v53 % ((v53) | 1))
            }
            var v114 int32 = v55
            _ = fi112
            _ = v114
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v61
}
