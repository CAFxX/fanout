package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 bool; F1 int16; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 int; F2 int32 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint32; F1 string; F2 int16 }
type GenS10 struct { F0 string; F1 int64; F2 int }
type GenS11 struct { F0 int32; F1 uint64; F2 int8 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 uint32; F1 string; F2 uint64 }
type GenS14 struct { F0 bool; F1 uint8 }
func genFn0(p0 int16) int64 {
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 874
}
func genFn1() int {
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint = 4294967295
        v12 := GenS1{F0: -1}
        switch "hello" {
        case "a":
            if false {
                v11 = v11
                v13 := []int{fi10, fi10, fi10}
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS1{F0: 733}
                _ = fi14
                _ = v15
            }
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if false {
        if (1e10 > 1.5) {
            switch -1 {
            case 197:
                v16 := GenS1{F0: 2147483647}
                _ = v16
            }
        }
    } else {
        v17 := []int{4294967295}
        _ = v17
    }
    v18 := []int{239, 9223372036854775807}
    var v19 uint16 = 65535
    _ = v18
    _ = v19
    return 128
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 100
    v3 := []string{"hello"}
    var v4 bool = (3.14159 != 3.14159)
    var v5 string = "\n\t"
    v6 := []int{2147483647}
    v4 = (v4 || (1 >= 2))
    switch -523 {
    case 607:
        var v20 int8 = 69
        _ = v20
    case 2147483647:
        v5 = v5
    case 4294967295:
        if v4 {
            if (42 > 208) {
                var v21 uint = 18446744073709551615
                var v22 float32 = 1.5
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi23
            }
            if len(v6) > 0 {
                v6[0] = 685
            }
        }
    default:
        v24 := GenS1{F0: -824}
        _ = v24
    }
    switch 32768 {
    case 42:
        v2 = v2
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        if (-128 > -128) {
            if (490 == 1) {
                fmt.Println(v4)
                var v26 uint16 = 255
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v4)
                case "hello":
                    v2 += (v2 >> ((1) & 63))
                case "\n\t":
                    for fi27 := 0; fi27 < 12; fi27++ {
                        if true {
                            v5 = v5
                            if (3.14159 <= float32(-2.25)) {
                                v28 := GenS2{F0: 255, F1: 9223372036854775807}
                                v2 += (v2 << ((65535) & 63))
                                v29 := GenS2{F0: 574, F1: genFn0(1000)}
                                _ = v28
                                _ = v29
                            }
                        }
                        _ = fi27
                    }
                }
                _ = v26
            } else {
                var v30 float32 = 0.0
                _ = v30
            }
            switch v5 {
            case "foo bar":
                v31 := GenS3{F0: 1000, F1: genFn1()}
                _ = v31
            case "":
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 uint8 = v2
                    v2 = (^(v33))
                    switch 0 {
                    case -778:
                        v34 := GenS4{F0: v4, F1: 0, F2: -9223372036854775808}
                        _ = v34
                    case 0:
                        v35 := GenS3{F0: 32767, F1: -215}
                        if v4 {
                            _ = len(v3)
                        }
                        if (!(242 >= 1)) {
                            var v36 uint8 = (v33 >> ((0) & 63))
                            v3 = append(v3, v5)
                            v37 := GenS5{F0: -1}
                            _ = v36
                            _ = v37
                        } else {
                            var v38 uint64 = 42
                            v39 := GenS4{F0: ((0.0 >= (0.1 / ((0.1 / 1.5) + float64(-2.25)))) && false), F1: 32767, F2: -325}
                            var v40 uint32 = 225
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        if (fi25 < 42) {
                            v2 = (^(v33))
                        }
                        var v41 bool = true
                        _ = v35
                        _ = v41
                    case -9223372036854775808:
                        var v42 bool = ((18446744073709551615 > 0) || false)
                        _ = v42
                    }
                    _ = fi32
                    _ = v33
                }
            case "hello":
                var v43 int8 = 100
                _ = v43
            default:
                if len(v6) > 0 {
                    v6[0] = fi25
                }
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v4 = (!(-128 >= -1))
        _ = fi25
    }
    var v44 int64 = -1
    fmt.Println(v44, v5)
    v5 = ("\n\t" + v5)
    if len(v3) > 0 {
        v3[0] = v5
    }
    v6 = append(v6, 4294967295)
    if true {
        var v45 uint = 255
        if v4 {
            if v4 {
                if v4 {
                    switch 32768 {
                    case -17:
                        var v46 int16 = 589
                        _ = v46
                    case -9223372036854775808:
                        for fi47 := 0; fi47 < 14; fi47++ {
                            switch (fi47 - (fi47 & (^(fi47)))) {
                            case 42:
                                var v48 uint16 = 65535
                                _ = v48
                            case -284:
                                var v49 float64 = 0.1
                                _ = v49
                            case 127:
                                v5 = (v5 + "")
                            }
                            fmt.Println(v45, v4, v4)
                            v5 = v5
                            _ = fi47
                        }
                    case -2147483648:
                        v5 = ""
                    default:
                        v50 := GenS5{F0: v44}
                        _ = v50
                    }
                } else {
                    switch 9223372036854775807 {
                    case 128:
                        for fi51 := 0; fi51 < 8; fi51++ {
                            v6 = append(v6, (fi51 >> ((v45) & 63)))
                            _ = fi51
                        }
                    case -237:
                        _ = len(v3)
                    case 32768:
                        for fi52 := 0; fi52 < 14; fi52++ {
                            var v53 string = "\n\t"
                            _ = fi52
                            _ = v53
                        }
                    }
                }
                for fi54 := 0; fi54 < 13; fi54++ {
                    v55 := GenS6{F0: 1, F1: fi54, F2: 42}
                    _ = fi54
                    _ = v55
                }
                switch ((v5 + "foo bar") + v5) {
                case "a":
                    switch ("" + v5) {
                    case "foo bar":
                        var v56 int32 = -1
                        _ = v56
                    case "\n\t":
                        for fi57 := 0; fi57 < 6; fi57++ {
                            for fi58 := 0; fi58 < 4; fi58++ {
                                var v59 float64 = float64(((1.5 + 1.5) * 3.14159))
                                switch "a" {
                                case "foo bar":
                                    if false {
                                        switch (fi57 ^ genFn1()) {
                                        case -2147483648:
                                            switch v5 {
                                            case "":
                                                if false {
                                                    fmt.Println(v4, v4)
                                                    fmt.Println(v59, v5)
                                                }
                                            default:
                                                v44 += genFn0(1000)
                                            }
                                        default:
                                            v60 := GenS7{F0: 1, F1: v2}
                                            _ = v60
                                        }
                                        var v61 int16 = 32767
                                        switch "" {
                                        case "hello":
                                            var v62 uint32 = uint32(v45)
                                            _ = v62
                                        case "foo bar":
                                            v63 := GenS8{F0: uint64(v45), F1: 1, F2: 103}
                                            _ = v63
                                        default:
                                            v5 = ""
                                            switch 128 {
                                            case -2147483648:
                                                switch v5 {
                                                case "foo bar":
                                                    var v64 int16 = (v61 % ((-32768) | 1))
                                                    _ = v64
                                                case "hello":
                                                    var v65 string = "foo bar"
                                                    _ = v65
                                                }
                                            default:
                                                v66 := GenS9{F0: 2147483648, F1: v5, F2: (v61 + v61)}
                                                _ = v66
                                            }
                                        }
                                        _ = v61
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                }
                                _ = fi58
                                _ = v59
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch "" {
                            case "a":
                                if ((v45 - (v45 / ((v45) | 1))) < (v45 - 65535)) {
                                    v3 = append(v3, ((v5 + v5) + v5))
                                    for fi67 := 0; fi67 < 13; fi67++ {
                                        if v4 {
                                            v4 = (((!(0 == uint32(v2))) && v4) && true)
                                        } else {
                                            if true {
                                                v68 := GenS10{F0: "a", F1: (v44 >> (((^(v45))) & 63)), F2: (fi57 >> ((v45) & 63))}
                                                var v69 int64 = int64(v45)
                                                _ = v68
                                                _ = v69
                                            } else {
                                                var v70 int = -832
                                                switch v5 {
                                                case "a":
                                                    v6 = append(v6, (-(v70)))
                                                case "foo bar":
                                                    v71 := GenS4{F0: true, F1: 32767, F2: v44}
                                                    _ = v71
                                                case "\n\t":
                                                    if false {
                                                        var v72 uint32 = 1
                                                        _ = v72
                                                    }
                                                default:
                                                    v5 = "foo bar"
                                                }
                                                v44 -= v44
                                                _ = v70
                                            }
                                            var v73 float64 = 0.1
                                            _ = v73
                                        }
                                        switch genFn1() {
                                        case -9223372036854775808:
                                            if (fi67 != int(v2)) {
                                                var v74 uint32 = 295
                                                v2 = v2
                                                v75 := GenS9{F0: (^(v74)), F1: v5, F2: -1}
                                                _ = v74
                                                _ = v75
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    v45 += v45
                                                    v3 = append(v3, (v5 + v5))
                                                case "a":
                                                    var v76 uint32 = 384
                                                    _ = v76
                                                case "":
                                                    var v77 uint16 = 1
                                                    _ = v77
                                                default:
                                                    switch fi67 {
                                                    case 2147483647:
                                                        switch v5 {
                                                        case "\n\t":
                                                            v78 := GenS11{F0: 2147483647, F1: 42, F2: -128}
                                                            _ = v78
                                                        }
                                                    }
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = (fi67 * (-(fi57)))
                                                }
                                            }
                                        case 0:
                                            _ = len(v3)
                                        case 2147483647:
                                            var v79 int64 = (v44 - (v44 * (v44 + -9223372036854775808)))
                                            _ = v79
                                        }
                                        _ = fi67
                                    }
                                    switch 128 {
                                    case -9223372036854775808:
                                        if true {
                                            v80 := GenS12{F0: ((1.5 - 1e10) != 3.14159), F1: (true && v4)}
                                            v2 = (v2 ^ (v2 / ((v2) | 1)))
                                            v44 = (v44 % ((597) | 1))
                                            _ = v80
                                        } else {
                                            switch fi57 {
                                            case -877:
                                                for fi81 := 0; fi81 < 18; fi81++ {
                                                    v44 += -868
                                                    if len(v3) > 0 {
                                                        v3[0] = ("" + "hello")
                                                    }
                                                    _ = fi81
                                                }
                                            case 127:
                                                var v82 int64 = v44
                                                _ = v82
                                            case -390:
                                                for fi83 := 0; fi83 < 6; fi83++ {
                                                    fmt.Println(v2, v44, v44)
                                                    _ = fi83
                                                }
                                            default:
                                                var v84 uint32 = 2
                                                _ = v84
                                            }
                                        }
                                    case 32768:
                                        v5 = ""
                                    }
                                } else {
                                    if (v4 && (!(!(!v4)))) {
                                        v85 := GenS6{F0: 32767, F1: (fi57 << ((v45) & 63)), F2: 266}
                                        v2 = 127
                                        _ = v85
                                    }
                                    var v86 uint16 = 536
                                    _ = v86
                                }
                            case "hello":
                                for fi87 := 0; fi87 < 12; fi87++ {
                                    var v88 float64 = 0.1
                                    _ = len(v3)
                                    _ = fi87
                                    _ = v88
                                }
                            default:
                                if (42 >= int16(v45)) {
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v89 := GenS4{F0: ((42 > -346) || (388 <= 437)), F1: 0, F2: (v44 % ((genFn0(1)) | 1))}
                                    for fi90 := 0; fi90 < 15; fi90++ {
                                        v1 = append(v1, 2)
                                        var v91 uint32 = 226
                                        switch fi90 {
                                        case 0:
                                            v45 -= v45
                                        case 2147483647:
                                            v44 = 0
                                            v44 += (v44 ^ int64(v91))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case 32767:
                                            v91 = 955
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        _ = fi90
                                        _ = v91
                                    }
                                    _ = v89
                                }
                            }
                            _ = fi57
                        }
                    case "":
                        v45 = 929
                    default:
                        switch v5 {
                        case "hello":
                            var v92 int8 = 0
                            _ = v92
                        case "":
                            v5 = v5
                        case "\n\t":
                            v45 -= (v45 % (((v45 ^ 515)) | 1))
                        default:
                            v93 := GenS5{F0: v44}
                            _ = v93
                        }
                    }
                default:
                    v94 := GenS3{F0: 1000, F1: 79}
                    _ = v94
                }
            }
        }
        _ = v45
    }
    v95 := GenS13{F0: 2147483647, F1: "", F2: 4294967295}
    v96 := GenS7{F0: 32767, F1: v2}
    v5 = v5
    v97 := GenS14{F0: v4, F1: v2}
    for fi98 := 0; fi98 < 3; fi98++ {
        if (!v4) {
            v99 := GenS12{F0: (((0 == 0) && (127 >= 59)) || true), F1: (0.0 != (3.14159 - 0.0))}
            var v100 uint32 = 1
            _ = v100
            _ = v99
        }
        _ = fi98
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v95
    _ = v96
    _ = v97
}
