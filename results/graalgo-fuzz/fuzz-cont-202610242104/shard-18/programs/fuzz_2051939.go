package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 string; F1 int16; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16; F1 int8; F2 string }
type GenS15 struct { F0 int16; F1 uint16; F2 int8 }
type GenS16 struct { F0 int8 }
func genFn0() int64 {
    var v4 uint16 = 705
    v4 = (v4 | v4)
    v4 -= 2
    if false {
        v5 := GenS1{F0: 856, F1: 0, F2: -124}
        var v6 uint8 = 0
        _ = v5
        _ = v6
    } else {
        fmt.Println(v4)
        var v7 bool = (false && (!(false || (!true))))
        _ = v7
    }
    _ = v4
    return -1
}
func genFn1(p0 uint, p1 uint16, p2 string) uint16 {
    var v8 uint = 458
    p1 = 415
    v8 = v8
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}

func main() {
    v1 := []string{"", "a", "foo bar", "hello"}
    var v2 int8 = -1
    var v3 string = "hello"
    var v9 int64 = 4
    var v10 int8 = 0
    v11 := GenS1{F0: 18446744073709551615, F1: -1, F2: 32767}
    v1 = append(v1, v3)
    if false {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v12 uint = 2
            _ = v12
        }
        fmt.Println(v10, v2, v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if len(v1) > 0 {
            v1[0] = (v3 + v3)
        }
        _ = len(v1)
        v10 -= (v2 - v10)
        _ = fi13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v3 = ""
        var v15 uint32 = 191
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS2{F0: "a", F1: 98, F2: 42}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v10 += v2
            if false {
                v10 -= (^(v10))
                if false {
                    v1 = append(v1, v3)
                    switch fi19 {
                    case 927:
                        if (-503 == 1) {
                            if true {
                                fmt.Println(v2, v10)
                                switch 2147483647 {
                                case 716:
                                    v20 := GenS3{F0: 42}
                                    _ = v20
                                case 575:
                                    v9 += genFn0()
                                }
                                if true {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v21 uint8 = 2
                                        switch (fi16 / (((fi19 + (-(fi19)))) | 1)) {
                                        case 42:
                                            v2 = 11
                                        case 0:
                                            fmt.Println(v2, v21)
                                        case 1:
                                            v21 += v21
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v3 = ((((v3 + ((v3 + (((v3 + "hello") + v3) + ((((v3 + (v3 + v3)) + v3) + v3) + ""))) + (v3 + "\n\t"))) + ("foo bar" + v3)) + "a") + (v3 + "hello"))
                                        _ = v21
                                    case "\n\t":
                                        if (true || true) {
                                            v22 := GenS4{F0: 804, F1: v3}
                                            switch v3 {
                                            case "a":
                                                switch -366 {
                                                case 4294967295:
                                                    if false {
                                                        var v23 uint64 = 0
                                                        _ = v23
                                                    }
                                                case -662:
                                                    v10 = v10
                                                case 32768:
                                                    fmt.Println(fi16, v9, v2)
                                                }
                                            default:
                                                if false {
                                                    fmt.Println(v2, v3)
                                                    var v24 uint8 = 231
                                                    if false {
                                                        if (v24 == (v24 | 0)) {
                                                            fmt.Println(v10)
                                                        } else {
                                                            var v25 uint16 = 65535
                                                            _ = v25
                                                        }
                                                        fmt.Println(v24, v2, fi16)
                                                    }
                                                    _ = v24
                                                } else {
                                                    var v26 int8 = (v10 - v10)
                                                    _ = v26
                                                }
                                            }
                                            var v27 int8 = (v2 >> ((255) & 63))
                                            _ = v22
                                            _ = v27
                                        } else {
                                            v10 += v2
                                            _ = len(v1)
                                            v2 += v10
                                        }
                                    case "":
                                        v28 := GenS4{F0: 4294967295, F1: (v3 + "foo bar")}
                                        _ = v28
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v29 uint = 42
                                        var v30 uint16 = 42
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 uint32 = 1
                                    var v32 uint8 = 159
                                    _ = v31
                                    _ = v32
                                } else {
                                    if true {
                                        switch (v3 + v3) {
                                        case "foo bar":
                                            v9 += 0
                                        case "a":
                                            if false {
                                                v33 := GenS5{F0: 2}
                                                _ = v33
                                            } else {
                                                var v34 uint64 = 786
                                                var v35 float32 = float32(((0.1 * 3.14159) - (-2.25 - 0.1)))
                                                _ = v34
                                                _ = v35
                                            }
                                        }
                                    } else {
                                        v9 += (^(v9))
                                    }
                                    v9 -= (^(v9))
                                    v36 := GenS6{F0: (-(v9)), F1: 0, F2: 1}
                                    var v37 float32 = 0.1
                                    _ = v36
                                    _ = v37
                                }
                                v38 := GenS3{F0: 1}
                                _ = v38
                            }
                            v39 := GenS7{F0: 4294967295}
                            switch v3 {
                            case "":
                                if (!(661 == 860)) {
                                    fmt.Println(v9)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v2)
                                    fmt.Println(v3)
                                    fmt.Println(v2)
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                    switch 9223372036854775807 {
                                    case 1:
                                        fmt.Println(v2, v3)
                                    case -146:
                                        if true {
                                            v40 := GenS8{F0: (fi19 + (-(fi19)))}
                                            v10 -= (v10 + (v2 | (v2 + (v2 | v2))))
                                            _ = v40
                                        } else {
                                            var v41 int8 = (-(v10))
                                            v2 -= v2
                                            _ = v41
                                        }
                                    case 970:
                                        v2 = int8(fi19)
                                    default:
                                        switch -636 {
                                        case 0:
                                            v3 = ""
                                        case 32767:
                                            switch (fi16 | 32767) {
                                            case 32767:
                                                v2 = v2
                                            }
                                        }
                                        switch fi16 {
                                        case 2147483647:
                                            switch (v3 + ((("hello" + v3) + ("" + "")) + ("" + ""))) {
                                            case "hello":
                                                fmt.Println(v9, fi16, v10)
                                            }
                                        default:
                                            var v42 uint = 252
                                            _ = v42
                                        }
                                    }
                                } else {
                                    v10 += (v2 << ((4294967295) & 63))
                                    v1 = append(v1, ((v3 + "hello") + (((v3 + ((v3 + (("" + v3) + "\n\t")) + "\n\t")) + v3) + "a")))
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        if (!false) {
                                            _ = len(v1)
                                            v43 := GenS2{F0: (v3 + ""), F1: 0, F2: -823}
                                            _ = v43
                                        } else {
                                            switch v3 {
                                            case "":
                                                var v44 uint64 = 899
                                                _ = v44
                                            case "hello":
                                                var v45 string = (v3 + (((v3 + v3) + ((((((((("" + "") + v3) + v3) + (("a" + (v3 + v3)) + "")) + "\n\t") + "") + v3) + (v3 + ("foo bar" + v3))) + ((v3 + v3) + "hello"))) + v3))
                                                _ = v45
                                            case "\n\t":
                                                fmt.Println(v2, v2, v10)
                                            }
                                        }
                                        if (!(255 != 555)) {
                                            fmt.Println(v9, v3, fi16)
                                        }
                                    } else {
                                        v3 = v3
                                    }
                                }
                            case "foo bar":
                                var v46 bool = (fi19 < fi19)
                                _ = v46
                            case "hello":
                                var v47 string = v3
                                _ = v47
                            }
                            _ = v39
                        } else {
                            if ((459 <= 2) || true) {
                                v3 = v3
                                if false {
                                    var v48 int8 = v10
                                    switch fi19 {
                                    case -988:
                                        v49 := GenS3{F0: 1000}
                                        var v50 uint64 = 18446744073709551615
                                        _ = v49
                                        _ = v50
                                    case 128:
                                        if true {
                                            v1 = append(v1, "")
                                            switch (fi16 | (fi16 - fi19)) {
                                            case 9223372036854775807:
                                                if true {
                                                    v1 = append(v1, (v3 + ("foo bar" + "foo bar")))
                                                    fmt.Println(fi19)
                                                }
                                                v2 -= (-(v10))
                                            case 127:
                                                var v51 int = 99
                                                _ = v51
                                            }
                                        } else {
                                            v48 = v2
                                        }
                                        v52 := GenS2{F0: v3, F1: -32768, F2: -1}
                                        var v53 uint32 = 0
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = v48
                                }
                                if true {
                                    switch "foo bar" {
                                    case "":
                                        v54 := GenS3{F0: 1}
                                        var v55 bool = true
                                        if (2 >= 0) {
                                            var v56 uint32 = 65535
                                            var v57 int16 = 1
                                            switch ((v3 + "a") + v3) {
                                            case "hello":
                                                if len(v1) > 0 {
                                                    v1[0] = ("" + v3)
                                                }
                                            case "\n\t":
                                                switch fi16 {
                                                case 42:
                                                    v58 := GenS9{F0: (v57 & v57)}
                                                    _ = v58
                                                case 1:
                                                    var v59 bool = ((fi19 != (fi16 | (fi16 << ((18446744073709551615) & 63)))) || true)
                                                    _ = v59
                                                case 32767:
                                                    if (!v55) {
                                                        v57 = (v57 & (v57 << ((628) & 63)))
                                                    } else {
                                                        v1 = append(v1, ("" + (v3 + "")))
                                                        if false {
                                                            if (244 >= 177) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v3 + v3)
                                                                }
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = "a"
                                                            }
                                                            v55 = v55
                                                            var v60 int64 = 0
                                                            v57 = (v57 >> ((2) & 63))
                                                            _ = v60
                                                        } else {
                                                            v56 += 2147483648
                                                        }
                                                    }
                                                default:
                                                    v10 -= -24
                                                }
                                            case "":
                                                var v61 int32 = 919
                                                _ = v61
                                            }
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v62 uint = 255
                                            _ = v62
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            } else {
                                var v63 string = v3
                                _ = v63
                            }
                            v64 := GenS4{F0: 255, F1: v3}
                            _ = v64
                        }
                    }
                    v10 = v10
                }
            } else {
                fmt.Println(v9)
            }
            v3 = ("a" + ((v3 + "hello") + ""))
            _ = fi19
        }
        if true {
            v10 = -128
        }
        _ = fi16
    }
    v2 = (v2 >> ((1) & 63))
    if true {
        fmt.Println(v3)
        if ((-1 != 42) && (1 == 4294967295)) {
            for fi65 := 0; fi65 < 2; fi65++ {
                if (genFn1(867, 255, v3) <= 704) {
                    if (616 < 618) {
                        fmt.Println(fi65)
                        v66 := GenS10{F0: 2147483648}
                        var v67 uint16 = 65535
                        _ = v66
                        _ = v67
                    }
                }
                _ = fi65
            }
            v2 = v10
            if (false && true) {
                var v68 float32 = 1e10
                _ = v68
            } else {
                v69 := GenS11{F0: v3}
                _ = v69
            }
        } else {
            var v70 string = (v3 + v3)
            for fi71 := 0; fi71 < 11; fi71++ {
                for fi72 := 0; fi72 < 1; fi72++ {
                    if (42 >= (v9 + (v9 ^ (v9 >> ((0) & 63))))) {
                        var v73 int8 = 81
                        if false {
                            v74 := GenS12{F0: true}
                            _ = v74
                        } else {
                            switch (-(fi72)) {
                            case 42:
                                fmt.Println(v2, v2)
                            default:
                                if true {
                                    var v75 uint64 = 0
                                    v73 += v2
                                    _ = v75
                                }
                            }
                        }
                        v76 := GenS9{F0: -58}
                        _ = v73
                        _ = v76
                    } else {
                        v77 := GenS13{F0: (fi71 + 127)}
                        var v78 uint8 = 110
                        var v79 string = v3
                        var v80 int8 = 100
                        v2 = (v2 * -59)
                        _ = v77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    if (!(314 == -1)) {
                        switch v70 {
                        case "":
                            var v81 uint8 = 21
                            _ = v81
                        default:
                            v82 := GenS14{F0: 1, F1: -67, F2: "hello"}
                            _ = v82
                        }
                    }
                    _ = fi72
                }
                for fi83 := 0; fi83 < 5; fi83++ {
                    fmt.Println(fi71)
                    switch -855 {
                    case 127:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v3 = v70
                    _ = fi83
                }
                _ = fi71
            }
            _ = len(v1)
            _ = v70
        }
        if true {
            v2 += v10
            switch 4294967295 {
            case 128:
                var v84 uint = 18446744073709551615
                _ = v84
            case -2147483648:
                var v85 int = 42
                _ = v85
            case 9223372036854775807:
                for fi86 := 0; fi86 < 20; fi86++ {
                    if (!true) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v87 bool = false
                        if true {
                            var v88 int64 = v9
                            switch v3 {
                            case "foo bar":
                                if (((((1000 > 32767) && (!true)) || ((42 <= 42) && (!false))) || v87) || v87) {
                                    v88 -= 42
                                    var v89 string = v3
                                    v2 -= v10
                                    _ = v89
                                }
                            case "":
                                v90 := GenS10{F0: 1}
                                _ = v90
                            }
                            _ = v88
                        } else {
                            switch v3 {
                            case "a":
                                switch "foo bar" {
                                case "a":
                                    var v91 uint32 = 2
                                    _ = v91
                                case "":
                                    var v92 uint32 = 2
                                    _ = v92
                                case "\n\t":
                                    for fi93 := 0; fi93 < 1; fi93++ {
                                        if true {
                                            if false {
                                                fmt.Println(v87, fi86)
                                                var v94 int16 = 898
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v94
                                            }
                                        }
                                        if v87 {
                                            var v95 uint64 = 1
                                            v96 := GenS12{F0: v87}
                                            _ = v95
                                            _ = v96
                                        } else {
                                            var v97 uint = 42
                                            v2 = v10
                                            _ = v97
                                        }
                                        _ = fi93
                                    }
                                }
                            default:
                                v2 += -107
                            }
                        }
                        _ = v87
                    }
                    var v98 string = ("\n\t" + "foo bar")
                    _ = fi86
                    _ = v98
                }
            }
        } else {
            var v99 float64 = -2.25
            _ = v99
        }
    }
    if false {
        var v100 int16 = 32767
        _ = v100
    }
    var v101 int8 = -1
    v102 := GenS15{F0: 32767, F1: 1, F2: (-(v10))}
    switch 253 {
    case 2147483647:
        for fi103 := 0; fi103 < 20; fi103++ {
            v104 := GenS16{F0: (v2 << ((4294967295) & 63))}
            _ = fi103
            _ = v104
        }
    case -601:
        var v105 bool = true
        _ = v105
    default:
        for fi106 := 0; fi106 < 15; fi106++ {
            switch (fi106 / ((fi106) | 1)) {
            case -1:
                v107 := GenS10{F0: 229}
                _ = v107
            case 32767:
                var v108 string = v3
                _ = v108
            default:
                v109 := GenS8{F0: (fi106 & fi106)}
                _ = v109
            }
            switch v3 {
            case "hello":
                switch (fi106 / ((fi106) | 1)) {
                case 42:
                    var v110 int16 = -32768
                    _ = v110
                case 2147483647:
                    v3 = v3
                case 9223372036854775807:
                    var v111 int = fi106
                    _ = v111
                }
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = v3
                }
            }
            _ = fi106
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v101
    _ = v102
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
