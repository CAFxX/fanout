package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 bool; F1 uint; F2 int }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS13 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS14 struct { F0 uint16; F1 uint32; F2 bool }

func main() {
    v1 := []uint32{4294967295, 2147483647, 65535, 109}
    var v2 bool = (47 <= -66)
    v3 := []int{-9223372036854775808, 744, 1}
    switch -1 {
    case -777:
        var v4 int32 = 2147483647
        _ = v4
    case 32768:
        v2 = (("a" + ("" + ("\n\t" + (("a" + "a") + "a")))) <= "a")
    case 563:
        for fi5 := 0; fi5 < 17; fi5++ {
            if (1000 > -1) {
                v6 := GenS1{F0: -9223372036854775808}
                _ = v6
            }
            if (852 < 1) {
                fmt.Println(fi5)
                v7 := GenS1{F0: 115}
                _ = v7
            } else {
                if (!v2) {
                    var v8 int64 = int64(fi5)
                    _ = v8
                }
            }
            v2 = v2
            _ = fi5
        }
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint32 = 4294967295
        v11 := GenS2{F0: (v10 ^ v10), F1: (fi9 - 32768)}
        switch fi9 {
        case -1:
            switch 277 {
            case -9223372036854775808:
                if v2 {
                    var v12 int16 = -32768
                    if ((-22 == 1) || (!v2)) {
                        v13 := GenS1{F0: 0}
                        fmt.Println(v12, fi9)
                        var v14 uint64 = 18446744073709551615
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(v2, v12, v12)
                        if v2 {
                            v2 = (!v2)
                            fmt.Println(v2, v12)
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v10 -= (v10 >> ((65535) & 63))
                    _ = v12
                }
            case 1:
                for fi15 := 0; fi15 < 10; fi15++ {
                    fmt.Println(fi15)
                    if (!(v2 || (455 < 18446744073709551615))) {
                        fmt.Println(v10)
                        v16 := GenS3{F0: false, F1: 255, F2: 0}
                        v10 = v10
                        var v17 float64 = (0.0 / 3.14159)
                        if (!((("foo bar" + "hello") + "hello") == "")) {
                            if (!false) {
                                if (1000 <= 1000) {
                                    var v18 uint32 = v10
                                    fmt.Println(v18)
                                    _ = v18
                                } else {
                                    v19 := GenS1{F0: -754}
                                    v10 -= 407
                                    switch ("\n\t" + ("foo bar" + ("hello" + "\n\t"))) {
                                    case "foo bar":
                                        fmt.Println(v17, fi9, v17)
                                    case "\n\t":
                                        if (v2 && false) {
                                            if v2 {
                                                v10 = v10
                                                v20 := GenS4{F0: 255, F1: 674}
                                                v10 -= (v10 - (v10 >> ((18446744073709551615) & 63)))
                                                _ = v20
                                            }
                                            var v21 uint16 = 2
                                            switch fi9 {
                                            case -2147483648:
                                                var v22 int32 = 42
                                                _ = v22
                                            case 567:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v23 float64 = (v17 - 0.1)
                                                    _ = v23
                                                case "hello":
                                                    var v24 float32 = ((0.1 - 0.0) * 0.0)
                                                    _ = v24
                                                case "\n\t":
                                                    var v25 bool = (fi15 < fi9)
                                                    _ = v25
                                                default:
                                                    if (v2 || v2) {
                                                        v26 := GenS5{F0: 1, F1: (v2 && (!(!(2147483647 <= fi15))))}
                                                        switch ("a" + "hello") {
                                                        case "a":
                                                            if (!true) {
                                                                v21 = (v21 << ((588) & 63))
                                                                fmt.Println(fi15, fi15, v2)
                                                                if (((0 > -1) && v2) || (false || false)) {
                                                                    if (-9223372036854775808 > -629) {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    } else {
                                                                        v10 += uint32(v21)
                                                                        v2 = false
                                                                        switch (fi15 % ((-267) | 1)) {
                                                                        case -433:
                                                                            if v2 {
                                                                                v27 := GenS4{F0: 0, F1: 32767}
                                                                                _ = v27
                                                                            }
                                                                        case -1:
                                                                            v2 = (42 > 4294967295)
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(v2, fi15, v21)
                                                                    var v28 uint64 = 4294967295
                                                                    _ = v28
                                                                }
                                                            } else {
                                                                var v29 uint16 = v21
                                                                var v30 string = ("\n\t" + "")
                                                                _ = v29
                                                                _ = v30
                                                            }
                                                        case "foo bar":
                                                            v31 := GenS1{F0: -9223372036854775808}
                                                            _ = v31
                                                        default:
                                                            fmt.Println(fi9, v10)
                                                        }
                                                        v32 := GenS6{F0: 42}
                                                        _ = v26
                                                        _ = v32
                                                    } else {
                                                        v17 += v17
                                                        v2 = v2
                                                        switch 128 {
                                                        case 4294967295:
                                                            if (false || v2) {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                _ = len(v3)
                                                            }
                                                        case -2147483648:
                                                            v17 -= v17
                                                        }
                                                    }
                                                }
                                            case 9223372036854775807:
                                                var v33 string = (("\n\t" + ("hello" + (("a" + "foo bar") + "a"))) + ("a" + "\n\t"))
                                                _ = v33
                                            default:
                                                v21 = v21
                                            }
                                            _ = v21
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = v10
                                            }
                                            v2 = (42 > 1)
                                            if v2 {
                                                if (!false) {
                                                    v34 := GenS4{F0: 2, F1: 589}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v34
                                                } else {
                                                    v2 = true
                                                    v2 = true
                                                }
                                                v35 := GenS1{F0: 0}
                                                fmt.Println(fi15)
                                                _ = v35
                                            } else {
                                                var v36 string = ""
                                                if (0 != v10) {
                                                    if len(v1) > 0 {
                                                        v1[0] = (^(v10))
                                                    }
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = v10
                                                    }
                                                    v17 = (v17 - v17)
                                                }
                                                var v37 int64 = 9223372036854775807
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                    }
                                    var v38 int32 = 0
                                    _ = v19
                                    _ = v38
                                }
                                switch (("a" + ((("" + ("\n\t" + "")) + ("hello" + (("\n\t" + ("" + "\n\t")) + "\n\t"))) + "hello")) + "\n\t") {
                                case "":
                                    var v39 int8 = 100
                                    _ = v39
                                case "a":
                                    fmt.Println(v10, fi15, v2)
                                case "\n\t":
                                    v40 := GenS2{F0: v10, F1: 128}
                                    _ = v40
                                }
                                v2 = (false && ((2 != 0) && v2))
                            }
                        } else {
                            var v41 uint64 = 42
                            _ = v41
                        }
                        _ = v16
                        _ = v17
                    }
                    _ = len(v1)
                    _ = fi15
                }
            default:
                var v42 uint8 = 42
                _ = v42
            }
        default:
            var v43 float32 = (((0.1 * 1.5) * ((-2.25 * 1.5) + ((((1e10 - float32(1e10)) * 0.0) + 1e10) / ((((3.14159 * 0.1) - 1e10) / (-2.25 / 1e10)) / 1.0)))) + float32(3.14159))
            _ = v43
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v2 = (!false)
    switch "" {
    case "a":
        if ((!v2) && (v2 && (!v2))) {
            var v44 int = 87
            switch ("" + "") {
            case "a":
                var v45 int16 = -1
                _ = v45
            case "foo bar":
                fmt.Println(v44, v2)
            default:
                fmt.Println(v44, v44)
            }
            _ = v44
        }
    default:
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
    }
    v3 = append(v3, 32768)
    var v46 uint = 65535
    v46 = (v46 + 429)
    v47 := GenS1{F0: -671}
    if (v2 || (!(v46 != (v46 * (v46 * 1))))) {
        for fi48 := 0; fi48 < 7; fi48++ {
            for fi49 := 0; fi49 < 10; fi49++ {
                v46 = (v46 & v46)
                v46 -= v46
                var v50 int16 = 432
                _ = fi49
                _ = v50
            }
            _ = fi48
        }
        v46 = v46
        if false {
            switch -585 {
            case -549:
                var v51 int = 9223372036854775807
                _ = v51
            default:
                if (true || (586 <= v46)) {
                    var v52 float32 = 0.0
                    for fi53 := 0; fi53 < 17; fi53++ {
                        if true {
                            v54 := GenS1{F0: 42}
                            _ = v54
                        } else {
                            for fi55 := 0; fi55 < 7; fi55++ {
                                if (v52 != (v52 / 3.14159)) {
                                    switch ("" + "") {
                                    case "a":
                                        v2 = v2
                                    default:
                                        var v56 uint = (v46 & v46)
                                        _ = v56
                                    }
                                }
                                fmt.Println(v2, fi53, v46)
                                v57 := GenS7{F0: 127, F1: 0}
                                _ = fi55
                                _ = v57
                            }
                            fmt.Println(fi53)
                            if (65535 >= 84) {
                                if false {
                                    fmt.Println(v52, v52)
                                    v58 := GenS3{F0: (!(1 >= -2147483648)), F1: 4294967295, F2: 844}
                                    v52 = (float32((1e10 - (0.1 * -2.25))) / (1e10 / float32(float64((float32(0.1) * v52)))))
                                    _ = v58
                                } else {
                                    fmt.Println(v52, v2)
                                    for fi59 := 0; fi59 < 6; fi59++ {
                                        fmt.Println(v52, fi53, fi53)
                                        var v60 int = (fi59 >> (((v46 ^ 4294967295)) & 63))
                                        var v61 int32 = -805
                                        v1 = append(v1, 65535)
                                        if true {
                                            var v62 string = (("" + "a") + "\n\t")
                                            _ = v62
                                        }
                                        _ = fi59
                                        _ = v60
                                        _ = v61
                                    }
                                }
                            }
                        }
                        v46 -= (v46 / ((12) | 1))
                        _ = fi53
                    }
                    if (!(!true)) {
                        if v2 {
                            var v63 uint16 = 255
                            var v64 float64 = 0.1
                            _ = v63
                            _ = v64
                        }
                    } else {
                        if (uint32(v46) != 868) {
                            v1 = append(v1, 42)
                        }
                        if v2 {
                            v2 = v2
                            if true {
                                switch ("foo bar" + "a") {
                                case "hello":
                                    if (((!(("foo bar" < "foo bar") && (!((!v2) || ((true && true) || (v2 || v2)))))) && v2) && ((231 < 390) && true)) {
                                        fmt.Println(v2, v46, v46)
                                        for fi65 := 0; fi65 < 19; fi65++ {
                                            var v66 bool = (!(true || (v2 && ((42 != 199) || v2))))
                                            v67 := GenS7{F0: 48, F1: 65535}
                                            _ = fi65
                                            _ = v66
                                            _ = v67
                                        }
                                    }
                                case "a":
                                    v68 := GenS8{F0: 535}
                                    _ = v68
                                default:
                                    v2 = v2
                                }
                            } else {
                                for fi69 := 0; fi69 < 15; fi69++ {
                                    v46 -= (v46 ^ v46)
                                    switch "\n\t" {
                                    case "hello":
                                        var v70 string = "hello"
                                        _ = v70
                                    case "foo bar":
                                        v2 = v2
                                    }
                                    _ = fi69
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        } else {
                            if (-1 < 1) {
                                var v71 string = ("a" + ("\n\t" + "a"))
                                v1 = append(v1, uint32(v46))
                                v52 = -2.25
                                _ = v71
                            } else {
                                v46 -= v46
                                v2 = (-719 >= 1000)
                                switch -9223372036854775808 {
                                case -454:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -9223372036854775808:
                                    var v72 int64 = 9223372036854775807
                                    _ = v72
                                default:
                                    for fi73 := 0; fi73 < 1; fi73++ {
                                        if v2 {
                                            for fi74 := 0; fi74 < 13; fi74++ {
                                                v75 := GenS9{F0: uint16(v46), F1: -9223372036854775808}
                                                v76 := GenS10{F0: uint16(v46), F1: -32768}
                                                _ = fi74
                                                _ = v75
                                                _ = v76
                                            }
                                            var v77 int16 = 1
                                            fmt.Println(v52)
                                            _ = v77
                                        }
                                        fmt.Println(v46)
                                        _ = fi73
                                    }
                                }
                            }
                            var v78 int16 = 0
                            fmt.Println(v2, v52)
                            _ = v78
                        }
                        if false {
                            v1 = append(v1, 85)
                            for fi79 := 0; fi79 < 13; fi79++ {
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                                var v80 uint64 = 655
                                var v81 int16 = 32767
                                _ = fi79
                                _ = v80
                                _ = v81
                            }
                        } else {
                            fmt.Println(v46, v52)
                            var v82 uint8 = 42
                            _ = v82
                        }
                    }
                    _ = v52
                }
            }
            for fi83 := 0; fi83 < 20; fi83++ {
                if v2 {
                    if v2 {
                        if true {
                            var v84 int16 = 998
                            _ = v84
                        } else {
                            v2 = (!v2)
                        }
                    }
                    var v85 uint8 = 165
                    _ = v85
                }
                v86 := GenS9{F0: 255, F1: -1}
                if ((true && v2) && false) {
                    var v87 bool = false
                    var v88 uint16 = 42
                    _ = len(v3)
                    _ = v87
                    _ = v88
                } else {
                    if len(v3) > 0 {
                        v3[0] = 663
                    }
                    var v89 uint32 = 65535
                    for fi90 := 0; fi90 < 7; fi90++ {
                        v2 = v2
                        v89 += v89
                        if v2 {
                            if (-128 != -73) {
                                if (-857 > -32768) {
                                    var v91 uint32 = 942
                                    _ = v91
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = v89
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    switch (fi83 * 0) {
                                    case -351:
                                        v92 := GenS2{F0: v89, F1: 1}
                                        _ = v92
                                    case 127:
                                        if (18446744073709551615 >= uint(v89)) {
                                            var v93 int8 = 127
                                            switch -1 {
                                            case 32767:
                                                switch "" {
                                                case "a":
                                                    var v94 string = "hello"
                                                    _ = v94
                                                case "\n\t":
                                                    var v95 uint32 = 930
                                                    _ = v95
                                                case "hello":
                                                    v3 = append(v3, (fi83 + int(v93)))
                                                default:
                                                    v1 = append(v1, v89)
                                                }
                                                v1 = append(v1, 802)
                                                if v2 {
                                                    var v96 uint32 = v89
                                                    v97 := GenS11{F0: "foo bar", F1: (!(!v2))}
                                                    _ = v96
                                                    _ = v97
                                                }
                                            case 100:
                                                var v98 bool = false
                                                _ = v98
                                            case 9223372036854775807:
                                                if v2 {
                                                    v99 := GenS12{F0: -9223372036854775808, F1: 844, F2: 0}
                                                    v1 = append(v1, 4294967295)
                                                    _ = v99
                                                }
                                            default:
                                                v100 := GenS13{F0: -82, F1: 551, F2: 153}
                                                _ = v100
                                            }
                                            _ = v93
                                        } else {
                                            var v101 uint64 = 870
                                            _ = v101
                                        }
                                    default:
                                        v89 = v89
                                    }
                                }
                                switch "" {
                                case "hello":
                                    v102 := GenS14{F0: 65535, F1: (v89 - v89), F2: (100 != 0)}
                                    _ = v102
                                case "a":
                                    var v103 int64 = 0
                                    _ = v103
                                case "":
                                    var v104 int32 = -1
                                    var v105 uint = 65535
                                    var v106 int16 = int16(fi90)
                                    _ = v104
                                    _ = v105
                                    _ = v106
                                }
                                var v107 int16 = -338
                                if false {
                                    v89 = 98
                                    var v108 int16 = 0
                                    var v109 int8 = -111
                                    var v110 uint8 = 249
                                    v3 = append(v3, fi83)
                                    _ = v108
                                    _ = v109
                                    _ = v110
                                }
                                _ = v107
                            }
                        } else {
                            var v111 string = "hello"
                            v46 = 42
                            switch "" {
                            case "hello":
                                var v112 uint64 = 1
                                _ = v112
                            case "\n\t":
                                v3 = append(v3, fi90)
                            case "a":
                                var v113 uint8 = 2
                                _ = v113
                            default:
                                if ((933 >= 2) || ((fi90 * (fi90 << ((0) & 63))) >= 32768)) {
                                    fmt.Println(v111, v2)
                                } else {
                                    var v114 int32 = 0
                                    v114 = -2147483648
                                    _ = v114
                                }
                            }
                            _ = v111
                        }
                        _ = fi90
                    }
                    _ = v89
                }
                _ = fi83
                _ = v86
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
}
