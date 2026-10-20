package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 bool; F2 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int; F1 string; F2 uint32 }
type GenS10 struct { F0 uint32; F1 string; F2 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8 }

func main() {
    v1 := []string{"\n\t", ("hello" + "a"), "\n\t", "a"}
    v2 := []string{("hello" + "a"), "foo bar", ""}
    var v3 bool = (true || (4 >= 127))
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: 845, F1: "a", F2: 252}
        switch (((("foo bar" + "\n\t") + "a") + "foo bar") + "") {
        case "\n\t":
            for fi6 := 0; fi6 < 16; fi6++ {
                if len(v1) > 0 {
                    v1[0] = (("foo bar" + "hello") + "\n\t")
                }
                _ = fi6
            }
        case "a":
            var v7 bool = v3
            _ = v7
        default:
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int = (fi8 >> ((18446744073709551615) & 63))
                _ = fi8
                _ = v9
            }
        }
        var v10 float64 = -2.25
        _ = fi4
        _ = v10
        _ = v5
    }
    _ = len(v2)
    if v3 {
        v11 := GenS1{F0: 540, F1: ("hello" + "\n\t"), F2: 42}
        _ = v11
    }
    var v12 string = "foo bar"
    if v3 {
        if (!v3) {
            if v3 {
                v1 = append(v1, v12)
                for fi13 := 0; fi13 < 11; fi13++ {
                    if v3 {
                        var v14 uint32 = 1
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            v16 := GenS2{F0: 107, F1: (("\n\t" + v12) + (v12 + (((("foo bar" + v12) + v12) + "foo bar") + "foo bar"))), F2: v3}
                            if false {
                                fmt.Println(fi13, v12)
                                var v17 float64 = 0.1
                                _ = v17
                            } else {
                                v12 = "a"
                                switch -9223372036854775808 {
                                case 128:
                                    if (42 < 100) {
                                        if v3 {
                                            v18 := GenS3{F0: 1}
                                            _ = v18
                                        } else {
                                            var v19 float32 = float32(-2.25)
                                            _ = v19
                                        }
                                        v12 = v12
                                        var v20 bool = true
                                        if v3 {
                                            v2 = append(v2, v12)
                                            if (107 <= 208) {
                                                var v21 uint8 = 127
                                                var v22 uint32 = 65535
                                                v22 = (^(v22))
                                                if len(v1) > 0 {
                                                    v1[0] = ((v12 + (v12 + v12)) + "a")
                                                }
                                                switch 128 {
                                                case 127:
                                                    v21 = (v21 >> ((2) & 63))
                                                case -1:
                                                    if v3 {
                                                        var v23 uint16 = 65535
                                                        var v24 int = (-(fi15))
                                                        v25 := GenS3{F0: 55}
                                                        switch fi13 {
                                                        case 0:
                                                            v26 := GenS3{F0: -984}
                                                            _ = v26
                                                        }
                                                        _ = v23
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                _ = v21
                                                _ = v22
                                            } else {
                                                if (!(65535 != 332)) {
                                                    switch ("hello" + (("" + ("foo bar" + (("foo bar" + "a") + ((("hello" + v12) + v12) + "\n\t")))) + v12)) {
                                                    case "a":
                                                        v27 := GenS2{F0: 1, F1: (v12 + v12), F2: (1 >= 593)}
                                                        _ = v27
                                                    case "\n\t":
                                                        if true {
                                                            if v20 {
                                                                v3 = (1 >= 42)
                                                                switch (("a" + v12) + (v12 + "")) {
                                                                case "\n\t":
                                                                    if ((!v3) && (2 <= 2147483647)) {
                                                                        fmt.Println(v20, fi13, v3)
                                                                        if v20 {
                                                                            v3 = (127 >= 127)
                                                                        } else {
                                                                            if (true || (248 < 42)) {
                                                                                v28 := GenS4{F0: true}
                                                                                v2 = append(v2, v12)
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                _ = v28
                                                                            } else {
                                                                                var v29 uint = 635
                                                                                var v30 int = 864
                                                                                _ = v29
                                                                                _ = v30
                                                                            }
                                                                            v20 = v3
                                                                            v31 := GenS4{F0: (1e10 >= float32(0.0))}
                                                                            _ = v31
                                                                        }
                                                                        if false {
                                                                            if true {
                                                                                if v3 {
                                                                                    v20 = false
                                                                                    var v32 int8 = 42
                                                                                    v3 = ((0 <= 2147483647) && ((42 == 42) || (v32 < v32)))
                                                                                    _ = v32
                                                                                }
                                                                                if false {
                                                                                    var v33 uint32 = 0
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v12
                                                                                    }
                                                                                    _ = v33
                                                                                }
                                                                                var v34 uint8 = 42
                                                                                var v35 uint8 = 42
                                                                                _ = v34
                                                                                _ = v35
                                                                            } else {
                                                                                var v36 int8 = -1
                                                                                v3 = v3
                                                                                fmt.Println(v20, v12, v12)
                                                                                _ = v36
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v37 := GenS1{F0: 2147483648, F1: v12, F2: 215}
                                                                _ = v37
                                                            } else {
                                                                if (!v3) {
                                                                    fmt.Println(v3)
                                                                } else {
                                                                    _ = len(v1)
                                                                    if ((v12 + v12) <= (((v12 + ("hello" + ((v12 + v12) + "foo bar"))) + "\n\t") + v12)) {
                                                                        if (!v3) {
                                                                            fmt.Println(fi15)
                                                                            v20 = false
                                                                            v38 := GenS2{F0: 127, F1: "hello", F2: v3}
                                                                            _ = v38
                                                                        } else {
                                                                            var v39 bool = true
                                                                            switch fi13 {
                                                                            case -9223372036854775808:
                                                                                var v40 uint16 = 606
                                                                                if v3 {
                                                                                    fmt.Println(v40)
                                                                                    if v3 {
                                                                                        v12 = v12
                                                                                    } else {
                                                                                        if (v3 || (true && false)) {
                                                                                            if (18446744073709551615 >= 569) {
                                                                                                var v41 string = "hello"
                                                                                                _ = v41
                                                                                            }
                                                                                        } else {
                                                                                            v1 = append(v1, "hello")
                                                                                            fmt.Println(fi13, v20, v12)
                                                                                            var v42 int32 = -2147483648
                                                                                            _ = v42
                                                                                        }
                                                                                        v3 = false
                                                                                    }
                                                                                }
                                                                                _ = v40
                                                                            }
                                                                            _ = v39
                                                                        }
                                                                        var v43 float32 = 1.5
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 int32 = -1
                                                                        switch ((("a" + (v12 + ("hello" + "hello"))) + ("foo bar" + "")) + (v12 + "")) {
                                                                        case "a":
                                                                            if true {
                                                                                var v45 int8 = -97
                                                                                var v46 uint8 = 255
                                                                                _ = v45
                                                                                _ = v46
                                                                            }
                                                                        case "foo bar":
                                                                            if (!(32767 <= 162)) {
                                                                                if v3 {
                                                                                    var v47 int32 = v44
                                                                                    _ = v47
                                                                                }
                                                                            }
                                                                        default:
                                                                            v20 = ((v20 && v20) || v20)
                                                                        }
                                                                        _ = v44
                                                                    }
                                                                    var v48 int32 = 0
                                                                    v3 = false
                                                                    _ = v48
                                                                }
                                                                v20 = (0.0 >= ((3.14159 + ((1.5 / -2.25) * 3.14159)) * 0.0))
                                                            }
                                                            if (0.0 == ((0.0 - float32(1e10)) + 1.5)) {
                                                                v49 := GenS5{F0: 0}
                                                                v50 := GenS4{F0: v3}
                                                                _ = len(v1)
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                var v51 string = ("a" + v12)
                                                                fmt.Println(v20, v3, v3)
                                                                _ = v51
                                                            }
                                                            v20 = v20
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v20, v3)
                                                    default:
                                                        if (!(100 > 127)) {
                                                            v52 := GenS4{F0: (2 >= 1)}
                                                            var v53 int8 = 1
                                                            v54 := GenS1{F0: 2147483648, F1: "foo bar", F2: 137}
                                                            if ((int64(fi15) > -1) && v3) {
                                                                var v55 float64 = (1.5 * -2.25)
                                                                _ = v55
                                                            }
                                                            if (449 < 0) {
                                                                v2 = append(v2, v12)
                                                                var v56 int = (fi13 * 2147483647)
                                                                var v57 float64 = (0.1 * (0.0 / ((-2.25 + float64(1.5)) / 3.14159)))
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                        } else {
                                                            fmt.Println(v20, fi15, v12)
                                                            if (2147483648 >= 2147483647) {
                                                                v3 = (18446744073709551615 == 42)
                                                                v58 := GenS6{F0: -935}
                                                                v59 := GenS7{F0: 32767, F1: (42 > -1), F2: 0}
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                var v60 string = "a"
                                                                switch fi15 {
                                                                case -2147483648:
                                                                    if (false || (-266 == -601)) {
                                                                        if (3.14159 != -2.25) {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (("foo bar" + v60) + (("hello" + (v60 + v12)) + "foo bar"))
                                                                            }
                                                                            _ = len(v2)
                                                                            switch v60 {
                                                                            case "a":
                                                                                var v61 uint8 = 2
                                                                                _ = v61
                                                                            case "\n\t":
                                                                                v20 = false
                                                                            case "":
                                                                                v3 = v3
                                                                            default:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = "hello"
                                                                                }
                                                                                v62 := GenS6{F0: 9223372036854775807}
                                                                                _ = v62
                                                                            }
                                                                            var v63 uint8 = 255
                                                                            fmt.Println(v63, v20, fi13)
                                                                            _ = v63
                                                                        }
                                                                    } else {
                                                                        var v64 int = (fi15 / ((32767) | 1))
                                                                        switch v60 {
                                                                        case "foo bar":
                                                                            v65 := GenS4{F0: false}
                                                                            _ = v65
                                                                        case "a":
                                                                            if v3 {
                                                                                if (1e10 != (3.14159 / 0.1)) {
                                                                                    var v66 int64 = 500
                                                                                    switch fi15 {
                                                                                    case 2147483647:
                                                                                        var v67 float32 = float32((1.5 / ((float64((((((1e10 / ((1e10 * -2.25) * 1.5)) - (-2.25 - float32(1.5))) * float32((float64(-2.25) * float64(1e10)))) + 1e10) - (0.1 - 0.1))) / 0.1) + -2.25)))
                                                                                        _ = v67
                                                                                    }
                                                                                    v60 = "foo bar"
                                                                                    _ = v66
                                                                                }
                                                                            } else {
                                                                                var v68 uint64 = 94
                                                                                _ = v68
                                                                            }
                                                                        default:
                                                                            v20 = (42 == 1)
                                                                        }
                                                                        _ = v64
                                                                    }
                                                                case 1:
                                                                    v12 = v12
                                                                }
                                                                if true {
                                                                    var v69 uint64 = 456
                                                                    v70 := GenS4{F0: v20}
                                                                    fmt.Println(v12, v3)
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                _ = v60
                                                            }
                                                            v71 := GenS8{F0: 18446744073709551615}
                                                            _ = v71
                                                        }
                                                    }
                                                    v12 = v12
                                                    switch (((("\n\t" + v12) + v12) + v12) + v12) {
                                                    case "\n\t":
                                                        v3 = v3
                                                        switch (fi13 & (fi15 % ((fi13) | 1))) {
                                                        case 1:
                                                            if ((v12 < v12) && v3) {
                                                                fmt.Println(fi13)
                                                            }
                                                        default:
                                                            switch (v12 + v12) {
                                                            case "a":
                                                                fmt.Println(fi15, fi13, v3)
                                                            case "foo bar":
                                                                if true {
                                                                    var v72 float32 = (((1.5 / 1.5) * (0.1 * (float32((0.1 * float64(-2.25))) + (3.14159 / 1.5)))) - 0.0)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v73 := GenS2{F0: 127, F1: ((v12 + (v12 + "a")) + "\n\t"), F2: v20}
                                                                    if v3 {
                                                                        if (!v3) {
                                                                            if v3 {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = "hello"
                                                                                }
                                                                                v74 := GenS2{F0: 109, F1: v12, F2: v20}
                                                                                v75 := GenS9{F0: fi15, F1: v12, F2: 393}
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                        } else {
                                                                            v76 := GenS8{F0: 0}
                                                                            var v77 int = (fi13 >> ((2) & 63))
                                                                            v78 := GenS8{F0: 711}
                                                                            _ = v76
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                    v79 := GenS3{F0: -9223372036854775808}
                                                                    v72 = (v72 / (float32(-2.25) + (1e10 - 1.5)))
                                                                    v2 = append(v2, ((v12 + "hello") + (v12 + ("a" + (v12 + v12)))))
                                                                    v20 = v3
                                                                    var v80 int32 = -1
                                                                    _ = v72
                                                                    _ = v73
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                            case "\n\t":
                                                                v1 = append(v1, v12)
                                                            }
                                                        }
                                                        v81 := GenS10{F0: 2147483648, F1: (v12 + v12), F2: 170}
                                                        _ = v81
                                                    case "a":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                var v82 uint16 = 635
                                                _ = v82
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                if (!(!((!(!v3)) && (!false)))) {
                                                    fmt.Println(fi15, v3, fi13)
                                                    var v83 int64 = -273
                                                    v2 = append(v2, "\n\t")
                                                    _ = v83
                                                }
                                            case "a":
                                                fmt.Println(fi15, fi13, v3)
                                                if len(v2) > 0 {
                                                    v2[0] = ("hello" + v12)
                                                }
                                            case "\n\t":
                                                if (false && (((("" + (v12 + v12)) + "") == (v12 + "\n\t")) || v20)) {
                                                    var v84 uint = 520
                                                    _ = v84
                                                } else {
                                                    var v85 uint = 0
                                                    var v86 uint16 = 255
                                                    _ = v85
                                                    _ = v86
                                                }
                                            }
                                        }
                                        var v87 int = fi15
                                        if v3 {
                                            var v88 int8 = 42
                                            _ = v88
                                        }
                                        _ = v20
                                        _ = v87
                                    } else {
                                        v12 = (v12 + v12)
                                        v89 := GenS3{F0: 0}
                                        _ = v89
                                    }
                                }
                            }
                            _ = fi15
                            _ = v16
                        }
                        fmt.Println(fi13, v12, v3)
                        if false {
                            v12 = v12
                            switch (-(fi13)) {
                            case 127:
                                for fi90 := 0; fi90 < 11; fi90++ {
                                    if (4294967295 > 42) {
                                        var v91 uint = 214
                                        _ = v91
                                    } else {
                                        v3 = v3
                                    }
                                    var v92 uint32 = 438
                                    switch fi13 {
                                    case 1:
                                        if v3 {
                                            switch (^(fi13)) {
                                            case 32768:
                                                var v93 uint8 = 0
                                                _ = v93
                                            case 128:
                                                _ = len(v2)
                                            case 1:
                                                if (v92 != v92) {
                                                    v94 := GenS11{F0: 526}
                                                    if v3 {
                                                        v3 = (96 < 127)
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v12 = v12
                                                    case "\n\t":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = v94
                                                } else {
                                                    var v95 int8 = 41
                                                    switch v12 {
                                                    case "hello":
                                                        v3 = false
                                                    case "":
                                                        var v96 float32 = 1e10
                                                        _ = v96
                                                    case "\n\t":
                                                        if (false && ((v95 >= (v95 * v95)) && ("a" <= ("" + "a")))) {
                                                            v97 := GenS10{F0: 2, F1: v12, F2: v92}
                                                            _ = v97
                                                        } else {
                                                            var v98 bool = (false && (42 >= 136))
                                                            _ = v98
                                                        }
                                                    }
                                                    _ = v95
                                                }
                                            }
                                        }
                                    case 127:
                                        if ((!v3) || (1 <= -1)) {
                                            fmt.Println(fi13, v92, v12)
                                            fmt.Println(fi13, v3)
                                            v2 = append(v2, ((("foo bar" + (v12 + v12)) + (v12 + v12)) + ("\n\t" + v12)))
                                        } else {
                                            v92 += (^(v92))
                                            if v3 {
                                                v92 = (v92 & (v92 & (v92 * v92)))
                                            }
                                            v92 -= v92
                                        }
                                    case -2147483648:
                                        switch 2147483647 {
                                        case 1:
                                            if (!true) {
                                                if len(v1) > 0 {
                                                    v1[0] = v12
                                                }
                                                var v99 float64 = 0.0
                                                if (!(v92 == (v92 * v92))) {
                                                    var v100 uint16 = 42
                                                    var v101 float32 = 1e10
                                                    var v102 string = (v12 + (v12 + (("" + ("hello" + "")) + (v12 + "hello"))))
                                                    _ = v100
                                                    _ = v101
                                                    _ = v102
                                                }
                                                _ = v99
                                            } else {
                                                v103 := GenS7{F0: 982, F1: v3, F2: -444}
                                                v1 = append(v1, v12)
                                                _ = v103
                                            }
                                            v3 = ((-2.25 != (-2.25 * 0.1)) || false)
                                        }
                                    default:
                                        if true {
                                            var v104 uint = 0
                                            var v105 int32 = 42
                                            _ = v104
                                            _ = v105
                                        }
                                    }
                                    _ = fi90
                                    _ = v92
                                }
                            case 1:
                                if ((!((v3 && v3) && ((fi13 == 1) || v3))) || (!v3)) {
                                    fmt.Println(v12, v3)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    var v106 float32 = 0.1
                                    v3 = (!(1e10 == 1.5))
                                    var v107 int = fi13
                                    _ = v106
                                    _ = v107
                                }
                            case -954:
                                for fi108 := 0; fi108 < 13; fi108++ {
                                    switch 42 {
                                    case 128:
                                        var v109 uint8 = 42
                                        _ = v109
                                    case 127:
                                        var v110 int64 = 590
                                        _ = v110
                                    }
                                    _ = fi108
                                }
                            }
                        }
                    }
                    fmt.Println(v3, v12)
                    v3 = v3
                    _ = fi13
                }
            }
            if ((v3 && v3) && (((!true) && (0.0 > 1.5)) && (!v3))) {
                v111 := GenS2{F0: -128, F1: "\n\t", F2: false}
                var v112 int = 0
                v112 += v112
                _ = v111
                _ = v112
            } else {
                fmt.Println(v3, v3)
            }
            fmt.Println(v12, v12)
        }
    } else {
        v3 = (414 != 959)
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v113 int = 42
    if (65535 <= 18446744073709551615) {
        v114 := GenS1{F0: 255, F1: (v12 + "\n\t"), F2: 255}
        _ = v114
    }
    v115 := GenS12{F0: 127}
    v3 = v3
    switch "\n\t" {
    case "hello":
        v3 = v3
    }
    v3 = (152 >= 483)
    v113 = v113
    v113 += v113
    fmt.Println(v12, v113, v12)
    var v116 int = v113
    var v117 float32 = ((3.14159 - ((float32(-2.25) * 0.0) + -2.25)) - (-2.25 * 0.1))
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v113:", v113)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v113
    _ = v115
    _ = v116
    _ = v117
    _ = v12
    _ = v2
    _ = v3
}
