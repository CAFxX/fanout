package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string; F2 int }
type GenS3 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint; F2 int64 }
type GenS6 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS7 struct { F0 string; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16; F1 string; F2 uint }
type GenS11 struct { F0 string }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS14 struct { F0 int; F1 int32; F2 uint16 }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 int16; F1 string; F2 uint32 }
type GenS17 struct { F0 uint8; F1 int32 }
type GenS18 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS19 struct { F0 int64 }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    var v3 int16 = (v1 ^ (^(v1)))
    v4 := []uint32{546, 904}
    if true {
        v3 = v1
    }
    _ = len(v2)
    v3 += 800
    var v5 uint = 1
    if true {
        var v6 int16 = v1
        _ = v6
    }
    if false {
        var v7 uint32 = 0
        var v8 uint16 = 255
        _ = v7
        _ = v8
    }
    var v9 int16 = v3
    v10 := GenS1{F0: 1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint = 2
        fmt.Println(v12, v9)
        _ = fi11
        _ = v12
    }
    switch -2147483648 {
    case 532:
        if (true || (244 > 2147483647)) {
            fmt.Println(v3, v9)
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            v3 = v3
            switch ("a" + (("hello" + ("hello" + ("foo bar" + ("foo bar" + ("a" + "foo bar"))))) + "foo bar")) {
            case "a":
                v9 = (v9 ^ v9)
            case "\n\t":
                var v15 uint8 = 127
                _ = v15
            case "hello":
                var v16 float64 = 3.14159
                _ = v16
            default:
                if (true || false) {
                    var v17 int8 = 42
                    var v18 uint16 = 0
                    v19 := v2[((("\n\t" + ((("\n\t" + "\n\t") + ("a" + "\n\t")) + ("a" + "\n\t"))) + "\n\t") + "\n\t")]
                    switch 42 {
                    case 32768:
                        v20 := GenS1{F0: 9223372036854775807}
                        _ = v20
                    case -791:
                        if (0.0 != 3.14159) {
                            var v21 int64 = 1
                            _ = v21
                        } else {
                            delete(v2, ("" + "hello"))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case -9223372036854775808:
                        v22 := v2[""]
                        _ = v22
                    }
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            var v23 int64 = 9223372036854775807
            _ = fi14
            _ = v23
        }
        switch "foo bar" {
        case "hello":
            if (100 == -128) {
                switch 32768 {
                case 2147483647:
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if (382 < 350) {
                            v5 = v5
                        } else {
                            v5 = 671
                        }
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        if false {
                            if (!false) {
                                var v25 int64 = 0
                                var v26 float64 = ((0.1 / ((-2.25 / (-2.25 * 1.5)) + 0.0)) + -2.25)
                                v25 -= v25
                                var v27 int = -9223372036854775808
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                fmt.Println(v5, v5)
                                _ = len(v2)
                                if true {
                                    delete(v2, ((("" + ("a" + "hello")) + "hello") + "foo bar"))
                                    v5 = v5
                                } else {
                                    v28 := GenS1{F0: 42}
                                    fmt.Println(v5)
                                    switch fi24 {
                                    case 127:
                                        v29 := GenS1{F0: 1}
                                        _ = v29
                                    case -980:
                                        if true {
                                            v9 = v3
                                        }
                                    case 2147483647:
                                        v30 := GenS1{F0: 339}
                                        _ = v30
                                    default:
                                        if false {
                                            if true {
                                                v31 := GenS2{F0: (^(fi24)), F1: "hello", F2: fi24}
                                                v1 -= 939
                                                _ = v31
                                            }
                                        }
                                    }
                                    _ = v28
                                }
                            }
                            v9 -= (v1 * (v3 % ((v9) | 1)))
                            switch 127 {
                            case 1:
                                v4 = append(v4, 2147483648)
                            case 32767:
                                if (127 <= 213) {
                                    v5 -= (v5 & v5)
                                }
                            case 4294967295:
                                switch 127 {
                                case -1:
                                    switch ("" + "hello") {
                                    case "foo bar":
                                        fmt.Println(fi13)
                                    case "hello":
                                        v32 := GenS1{F0: 42}
                                        _ = v32
                                    case "":
                                        v9 -= v3
                                    default:
                                        v1 = v3
                                    }
                                case -2147483648:
                                    if (false && (32767 != 42)) {
                                        var v33 string = "\n\t"
                                        _ = v33
                                    } else {
                                        if (-2147483648 <= -273) {
                                            v2[("hello" + ("foo bar" + "foo bar"))] = (-(fi24))
                                        }
                                        fmt.Println(v1, v1, fi24)
                                        v3 += -32768
                                    }
                                }
                            default:
                                v3 = v1
                                _ = len(v4)
                            }
                            var v34 uint16 = 1
                            _ = v34
                        } else {
                            var v35 uint64 = 978
                            if true {
                                if true {
                                    v36 := GenS1{F0: 733}
                                    if true {
                                        var v37 uint32 = 4294967295
                                        _ = v37
                                    }
                                    _ = v36
                                } else {
                                    var v38 float64 = 3.14159
                                    if ((127 > -1) && true) {
                                        var v39 int64 = 1
                                        var v40 int32 = -1
                                        _ = v39
                                        _ = v40
                                    } else {
                                        if false {
                                            v9 -= v3
                                        } else {
                                            delete(v2, "a")
                                            var v41 float32 = 3.14159
                                            _ = v41
                                        }
                                        v3 = v1
                                        var v42 bool = false
                                        _ = v42
                                    }
                                    switch -9223372036854775808 {
                                    case 9223372036854775807:
                                        if ((^(v5)) <= (^(v5))) {
                                            if true {
                                                var v43 uint64 = v35
                                                _ = v43
                                            }
                                            v5 += 197
                                            var v44 uint = 454
                                            if ((-(v3)) <= -1) {
                                                var v45 int16 = -63
                                                _ = v45
                                            }
                                            switch 4294967295 {
                                            case 2147483647:
                                                if false {
                                                    v46 := GenS3{F0: (^(v35)), F1: 0, F2: 783}
                                                    if true {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if (-1 >= 42) {
                                                            fmt.Println(fi24, fi24)
                                                            if (0 == 38) {
                                                                v35 = (v35 & v35)
                                                                v4 = append(v4, 954)
                                                                v47 := GenS2{F0: (fi13 + -9223372036854775808), F1: ("" + "\n\t"), F2: fi13}
                                                                _ = len(v4)
                                                                v48 := GenS4{F0: 0, F1: 2}
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            if false {
                                                                v1 = -736
                                                                var v49 uint64 = 42
                                                                switch 32768 {
                                                                case 525:
                                                                    var v50 float64 = 0.0
                                                                    _ = v50
                                                                case 128:
                                                                    if (!true) {
                                                                        v51 := GenS2{F0: (fi13 % ((fi13) | 1)), F1: "a", F2: (fi13 | -9223372036854775808)}
                                                                        switch (fi24 % (((fi13 >> ((65535) & 63))) | 1)) {
                                                                        case -2147483648:
                                                                            fmt.Println(v49, v35, fi24)
                                                                        case 127:
                                                                            fmt.Println(fi24)
                                                                            var v52 float32 = 3.14159
                                                                            var v53 int32 = 2147483647
                                                                            _ = v52
                                                                            _ = v53
                                                                        default:
                                                                            if (2 <= 80) {
                                                                                if true {
                                                                                    switch 128 {
                                                                                    case -774:
                                                                                        if false {
                                                                                            if false {
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    var v54 int = (fi13 << ((v5) & 63))
                                                                                                    _ = v54
                                                                                                case "":
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = 2147483648
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    _ = len(v2)
                                                                                                    if true {
                                                                                                        var v55 int16 = (v3 * v1)
                                                                                                        var v56 int64 = int64(v35)
                                                                                                        _ = v55
                                                                                                        _ = v56
                                                                                                    }
                                                                                                    if (-1 == 0) {
                                                                                                        if false {
                                                                                                            var v57 int64 = 1
                                                                                                            v35 = (^(v35))
                                                                                                            v38 -= (float64(-2.25) - (0.0 / ((-2.25 + (float64((0.1 + 3.14159)) + 0.0)) + (v38 - (0.1 + v38)))))
                                                                                                            _ = v57
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v35 -= v49
                                                                                                                if false {
                                                                                                                    if false {
                                                                                                                        v35 = v49
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                        var v58 int64 = -9223372036854775808
                                                                                                                        _ = v58
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        v59 := GenS3{F0: 18446744073709551615, F1: 1, F2: 2147483647}
                                                                                                                        _ = v59
                                                                                                                    }
                                                                                                                }
                                                                                                                var v60 int8 = 127
                                                                                                                _ = v60
                                                                                                            }
                                                                                                            var v61 int = (fi24 + (-(fi24)))
                                                                                                            _ = v61
                                                                                                        }
                                                                                                        v35 += v35
                                                                                                        var v62 uint64 = (v35 >> ((v44) & 63))
                                                                                                        _ = v62
                                                                                                    } else {
                                                                                                        fmt.Println(v44)
                                                                                                        v63 := GenS5{F0: (v35 & 42), F1: v5, F2: 9223372036854775807}
                                                                                                        _ = v63
                                                                                                    }
                                                                                                default:
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = len(v2)
                                                                                            v35 -= v35
                                                                                        } else {
                                                                                            if true {
                                                                                                var v64 uint8 = 1
                                                                                                if (1 > 2147483648) {
                                                                                                    _ = len(v4)
                                                                                                    var v65 uint32 = 65535
                                                                                                    v5 = (v44 | 18446744073709551615)
                                                                                                    var v66 string = "\n\t"
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                } else {
                                                                                                    v67 := GenS6{F0: 470, F1: 42, F2: 604}
                                                                                                    var v68 int16 = -286
                                                                                                    _ = v67
                                                                                                    _ = v68
                                                                                                }
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = 4294967295
                                                                                                }
                                                                                                if true {
                                                                                                    var v69 uint16 = 32767
                                                                                                    if ((v5 * v44) >= 42) {
                                                                                                        switch ("foo bar" + "") {
                                                                                                        case "hello":
                                                                                                            var v70 float32 = 0.1
                                                                                                            _ = v70
                                                                                                        case "a":
                                                                                                            v71 := GenS4{F0: 1, F1: (v69 - v69)}
                                                                                                            _ = v71
                                                                                                        }
                                                                                                        v1 = -32768
                                                                                                        _ = len(v4)
                                                                                                        v72 := GenS7{F0: ("a" + "\n\t"), F1: 1, F2: (v64 & v64)}
                                                                                                        v3 += (^(v1))
                                                                                                        var v73 int64 = -1
                                                                                                        _ = v72
                                                                                                        _ = v73
                                                                                                    } else {
                                                                                                        v74 := v2[("" + ("foo bar" + ("a" + (("foo bar" + (((("" + "a") + "hello") + "") + ("a" + ("hello" + "hello")))) + ("" + "a")))))]
                                                                                                        fmt.Println(fi13, v49, v5)
                                                                                                        _ = v74
                                                                                                    }
                                                                                                    _ = v69
                                                                                                } else {
                                                                                                    if true {
                                                                                                        var v75 float32 = 0.1
                                                                                                        if true {
                                                                                                            var v76 int16 = v3
                                                                                                            fmt.Println(v64, v5)
                                                                                                            _ = v76
                                                                                                        } else {
                                                                                                            var v77 bool = false
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            _ = len(v4)
                                                                                                            v2["a"] = 0
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        fmt.Println(v44, v5)
                                                                                                        _ = v75
                                                                                                    }
                                                                                                }
                                                                                                _ = v64
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch -169 {
                                                                                        case -331:
                                                                                            var v78 int = fi24
                                                                                            _ = v78
                                                                                        case 128:
                                                                                            if false {
                                                                                                var v79 int16 = v3
                                                                                                _ = v79
                                                                                            } else {
                                                                                                v80 := GenS8{F0: v49}
                                                                                                _ = v80
                                                                                            }
                                                                                        case 123:
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        v35 = (v49 + v35)
                                                                                        v81 := GenS9{F0: 128}
                                                                                        _ = v81
                                                                                    } else {
                                                                                        v1 = -1
                                                                                        var v82 int8 = 100
                                                                                        _ = v82
                                                                                    }
                                                                                    v83 := v2[(((("foo bar" + ("foo bar" + "foo bar")) + "hello") + "\n\t") + ("" + ((((("foo bar" + "hello") + "foo bar") + "a") + ("hello" + "foo bar")) + "hello")))]
                                                                                    v49 -= v35
                                                                                    v84 := GenS10{F0: v1, F1: "foo bar", F2: v5}
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                }
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 0
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v1, v44)
                                                                        _ = v51
                                                                    }
                                                                }
                                                                _ = v49
                                                            }
                                                            var v85 int16 = v9
                                                            _ = v85
                                                        }
                                                        var v86 float32 = 0.1
                                                        var v87 bool = (true && true)
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    switch (("a" + "foo bar") + "foo bar") {
                                                    case "":
                                                        if ((v38 != 1e10) && true) {
                                                            var v88 int32 = -174
                                                            v44 = v5
                                                            _ = v88
                                                        } else {
                                                            v44 = v44
                                                        }
                                                    case "foo bar":
                                                        switch ((("\n\t" + ("foo bar" + (("a" + "hello") + ("" + "")))) + (("\n\t" + "hello") + "a")) + "foo bar") {
                                                        case "a":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v89 int = (fi24 + (fi13 + (fi24 / ((9223372036854775807) | 1))))
                                                                _ = v89
                                                            case "\n\t":
                                                                v44 = (^(v44))
                                                            case "":
                                                                var v90 uint32 = 728
                                                                _ = v90
                                                            default:
                                                                var v91 int32 = int32(fi24)
                                                                _ = v91
                                                            }
                                                        case "\n\t":
                                                            var v92 uint16 = 1
                                                            _ = v92
                                                        case "hello":
                                                            v4 = append(v4, 468)
                                                        default:
                                                            switch "hello" {
                                                            case "a":
                                                                var v93 bool = true
                                                                _ = v93
                                                            case "":
                                                                if (false && ((float64(1e10) + v38) <= 1.5)) {
                                                                    v3 -= 205
                                                                }
                                                            case "foo bar":
                                                                if ((2 <= 2) && ((true || false) && (true && (!(0 <= 32767))))) {
                                                                    if true {
                                                                        var v94 int32 = 616
                                                                        _ = v94
                                                                    }
                                                                    v35 += (^(v35))
                                                                    fmt.Println(v9, v5, v1)
                                                                    v95 := v2["foo bar"]
                                                                    v3 = -1
                                                                    _ = v95
                                                                } else {
                                                                    var v96 uint8 = 0
                                                                    _ = v96
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v46
                                                }
                                            case 0:
                                                var v97 uint64 = 2
                                                _ = v97
                                            case 32767:
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483647
                                                }
                                            }
                                            _ = v44
                                        } else {
                                            v98 := GenS11{F0: ""}
                                            if false {
                                                if (!true) {
                                                    v5 = v5
                                                }
                                                v99 := GenS9{F0: 255}
                                                _ = v99
                                            } else {
                                                v100 := GenS4{F0: -128, F1: 424}
                                                _ = v100
                                            }
                                            if (true || false) {
                                                switch ("" + ("\n\t" + ("a" + ""))) {
                                                case "hello":
                                                    v101 := GenS12{F0: ""}
                                                    _ = v101
                                                case "":
                                                    v35 -= v35
                                                case "a":
                                                    if len(v4) > 0 {
                                                        v4[0] = 4294967295
                                                    }
                                                    switch 4294967295 {
                                                    case 0:
                                                        var v102 uint = (v5 << ((4294967295) & 63))
                                                        _ = v102
                                                    case 4294967295:
                                                        var v103 int32 = -2147483648
                                                        _ = v103
                                                    default:
                                                        v38 -= (((v38 * 3.14159) - v38) / float64(3.14159))
                                                    }
                                                default:
                                                    v35 -= 2
                                                }
                                                var v104 int64 = 42
                                                v3 += (v9 * v3)
                                                v104 = (v104 ^ v104)
                                                _ = v104
                                            } else {
                                                var v105 int8 = -128
                                                var v106 float64 = v38
                                                _ = v105
                                                _ = v106
                                            }
                                            v107 := GenS12{F0: ""}
                                            if (255 >= 255) {
                                                fmt.Println(v9, v1, fi24)
                                                delete(v2, "")
                                            }
                                            _ = v107
                                            _ = v98
                                        }
                                    case -1:
                                        v108 := v2[""]
                                        _ = v108
                                    default:
                                        switch ("" + "") {
                                        case "hello":
                                            var v109 uint8 = 69
                                            _ = v109
                                        case "a":
                                            v35 = (v35 % (((v35 | v35)) | 1))
                                        case "":
                                            fmt.Println(v9, v9)
                                        default:
                                            v110 := GenS13{F0: (^(v35)), F1: 0, F2: 42}
                                            _ = v110
                                        }
                                    }
                                    _ = v38
                                }
                            } else {
                                v35 -= (v35 & (v35 | v35))
                                if (1e10 >= 0.1) {
                                    var v111 uint8 = 255
                                    fmt.Println(fi13, v1)
                                    var v112 string = ""
                                    _ = v111
                                    _ = v112
                                } else {
                                    var v113 uint32 = 65535
                                    v3 -= (v1 - v1)
                                    switch (("foo bar" + "\n\t") + ("\n\t" + "\n\t")) {
                                    case "\n\t":
                                        v9 += (v3 / ((-32768) | 1))
                                    }
                                    _ = v113
                                }
                            }
                            var v114 uint16 = 0
                            _ = v114
                            _ = v35
                        }
                        _ = fi24
                    }
                case -922:
                    fmt.Println(fi13)
                default:
                    v115 := GenS10{F0: v1, F1: (("\n\t" + "\n\t") + (("\n\t" + "a") + ("" + "hello"))), F2: (v5 / ((v5) | 1))}
                    _ = v115
                }
                for fi116 := 0; fi116 < 17; fi116++ {
                    var v117 uint8 = 1
                    v5 = v5
                    _ = fi116
                    _ = v117
                }
            } else {
                fmt.Println(v1, v5, v1)
            }
        default:
            for fi118 := 0; fi118 < 13; fi118++ {
                switch (-(fi118)) {
                case 4294967295:
                    v1 = (v9 - v1)
                }
                v2[""] = 32767
                _ = fi118
            }
        }
        _ = fi13
    }
    if false {
        var v119 int32 = 673
        _ = v119
    } else {
        for fi120 := 0; fi120 < 20; fi120++ {
            for fi121 := 0; fi121 < 6; fi121++ {
                if false {
                    switch ("a" + ("\n\t" + "a")) {
                    case "":
                        switch "foo bar" {
                        case "hello":
                            switch int(v5) {
                            case 32767:
                                var v122 int = fi121
                                _ = v122
                            case 127:
                                v1 += (-(v9))
                            default:
                                v3 = 42
                            }
                            fmt.Println(fi121, fi121)
                            v5 -= 18446744073709551615
                            if (!(603 != v1)) {
                                switch (^(fi121)) {
                                case 127:
                                    if false {
                                        fmt.Println(v5, fi120, v9)
                                        var v123 float32 = (3.14159 - 1e10)
                                        _ = v123
                                    } else {
                                        var v124 uint = v5
                                        _ = v124
                                    }
                                default:
                                    fmt.Println(v1, fi120, v3)
                                }
                                var v125 uint8 = 42
                                _ = v125
                            } else {
                                switch "" {
                                case "foo bar":
                                    if (0 == 409) {
                                        v4 = append(v4, 719)
                                    } else {
                                        v9 -= v9
                                        if (32767 < 704) {
                                            fmt.Println(fi121, v5)
                                        } else {
                                            var v126 bool = (fi121 >= 4294967295)
                                            v127 := GenS14{F0: fi120, F1: -2147483648, F2: 1}
                                            v9 -= (v9 ^ v3)
                                            _ = v126
                                            _ = v127
                                        }
                                        var v128 int = 584
                                        _ = v128
                                    }
                                case "hello":
                                    fmt.Println(v9, fi120)
                                case "a":
                                    if false {
                                        v9 = (v1 << (((v5 / (((v5 | v5)) | 1))) & 63))
                                        var v129 string = "hello"
                                        v130 := GenS2{F0: (fi120 / ((fi120) | 1)), F1: ("foo bar" + "hello"), F2: (fi120 ^ (fi120 | fi120))}
                                        _ = v129
                                        _ = v130
                                    } else {
                                        v9 -= (^(v9))
                                    }
                                default:
                                    v1 = v9
                                }
                                if len(v4) > 0 {
                                    v4[0] = 311
                                }
                            }
                        case "foo bar":
                            v3 -= v3
                        case "\n\t":
                            v131 := GenS15{F0: v5, F1: 2147483647}
                            _ = v131
                        default:
                            v132 := GenS14{F0: fi120, F1: 42, F2: 0}
                            _ = v132
                        }
                    case "foo bar":
                        if (!false) {
                            var v133 string = "foo bar"
                            v134 := v2["\n\t"]
                            v4 = append(v4, 1)
                            _ = v133
                            _ = v134
                        } else {
                            v5 = (v5 ^ v5)
                            switch ("a" + ("foo bar" + "")) {
                            case "\n\t":
                                switch ("hello" + "a") {
                                case "a":
                                    if false {
                                        if (true || (((float32(((1e10 + 0.1) - 1.5)) - ((-2.25 + ((1.5 - (0.1 - -2.25)) - (-2.25 * 1.5))) * 0.1)) * (1e10 / 1.5)) <= 0.1)) {
                                            var v135 int64 = 412
                                            var v136 int16 = v1
                                            _ = v135
                                            _ = v136
                                        } else {
                                            v5 = (^(v5))
                                            switch "hello" {
                                            case "foo bar":
                                                var v137 float32 = 0.0
                                                var v138 uint32 = 249
                                                _ = v137
                                                _ = v138
                                            case "a":
                                                v3 += v9
                                            case "":
                                                fmt.Println(v5, fi121)
                                                v139 := GenS16{F0: v9, F1: (("a" + "foo bar") + ("\n\t" + "hello")), F2: 42}
                                                if false {
                                                    if false {
                                                        var v140 string = "a"
                                                        var v141 int64 = -253
                                                        v141 -= int64(v3)
                                                        _ = v140
                                                        _ = v141
                                                    }
                                                }
                                                _ = v139
                                            default:
                                                v142 := GenS7{F0: ("" + ("" + ("a" + "foo bar"))), F1: 127, F2: 42}
                                                _ = v142
                                            }
                                        }
                                        v5 = (v5 | v5)
                                        v143 := GenS9{F0: 31}
                                        _ = v143
                                    }
                                case "hello":
                                    v9 += int16(v5)
                                default:
                                    switch "foo bar" {
                                    case "\n\t":
                                        v144 := GenS15{F0: 973, F1: 630}
                                        _ = v144
                                    case "hello":
                                        if (true && true) {
                                            var v145 uint16 = 65535
                                            v145 += 255
                                            v146 := v2[("" + "")]
                                            _ = v145
                                            _ = v146
                                        }
                                    case "foo bar":
                                        if (!((-(v9)) <= (^(v9)))) {
                                            _ = len(v4)
                                            v147 := GenS17{F0: 128, F1: 1}
                                            _ = v147
                                        }
                                    }
                                }
                            case "a":
                                var v148 uint = (v5 << ((v5) & 63))
                                _ = v148
                            default:
                                switch fi121 {
                                case 0:
                                    fmt.Println(v1, v1, v3)
                                case 127:
                                    v5 = 746
                                default:
                                    var v149 uint = (^(v5))
                                    var v150 int64 = 0
                                    if true {
                                        fmt.Println(v5, fi120)
                                    } else {
                                        v151 := GenS18{F0: 733, F1: 349, F2: -1}
                                        v152 := GenS13{F0: 2, F1: v1, F2: 0}
                                        _ = v151
                                        _ = v152
                                    }
                                    v153 := GenS6{F0: uint16(v5), F1: 42, F2: 273}
                                    _ = v149
                                    _ = v150
                                    _ = v153
                                }
                            }
                        }
                    }
                    v1 = int16(v5)
                    switch ("hello" + "hello") {
                    case "hello":
                        v5 = (^(v5))
                    case "foo bar":
                        var v154 uint16 = 32767
                        _ = v154
                    }
                }
                v4 = append(v4, 71)
                v9 = (v9 + (v3 % ((v9) | 1)))
                _ = fi121
            }
            _ = fi120
        }
        fmt.Println(v9, v9)
    }
    var v155 float32 = 0.1
    v156 := GenS12{F0: (("\n\t" + ("hello" + "foo bar")) + "")}
    v157 := v2["a"]
    v9 -= (^(v3))
    if true {
        if true {
            v158 := GenS18{F0: 1, F1: 4294967295, F2: -128}
            if true {
                for fi159 := 0; fi159 < 16; fi159++ {
                    var v160 float64 = ((1e10 / 1.0) * 1e10)
                    _ = fi159
                    _ = v160
                }
            } else {
                var v161 uint = (v5 / (((v5 | v5)) | 1))
                v5 = (v161 - v5)
                _ = v161
            }
            for fi162 := 0; fi162 < 5; fi162++ {
                switch (v157 ^ fi162) {
                case 32768:
                    var v163 int16 = (v9 / ((1) | 1))
                    _ = v163
                case 2147483647:
                    v164 := GenS13{F0: 42, F1: (^(v1)), F2: 953}
                    _ = v164
                case 630:
                    if false {
                        if (!(0 > 63)) {
                            for fi165 := 0; fi165 < 14; fi165++ {
                                var v166 uint64 = 42
                                var v167 int8 = 1
                                v168 := GenS3{F0: v166, F1: 560, F2: -433}
                                fmt.Println(v155, v5)
                                v5 = v5
                                var v169 int16 = (v3 / ((0) | 1))
                                delete(v2, ("foo bar" + "foo bar"))
                                var v170 uint8 = 1
                                _ = fi165
                                _ = v166
                                _ = v167
                                _ = v168
                                _ = v169
                                _ = v170
                            }
                            _ = len(v2)
                        }
                        fmt.Println(v5, v155)
                    } else {
                        switch (("a" + "a") + ("hello" + ("hello" + ""))) {
                        case "\n\t":
                            var v171 uint64 = 751
                            _ = v171
                        default:
                            var v172 uint32 = 963
                            _ = v172
                        }
                    }
                }
                if false {
                    v173 := GenS14{F0: fi162, F1: -1, F2: 2}
                    v174 := GenS19{F0: -902}
                    v9 -= (v3 / ((v3) | 1))
                    _ = v173
                    _ = v174
                } else {
                    var v175 int32 = 0
                    fmt.Println(v175, fi162)
                    for fi176 := 0; fi176 < 16; fi176++ {
                        v3 -= -1
                        fmt.Println(v3, v155)
                        switch (fi162 + (-(v157))) {
                        case -489:
                            _ = len(v2)
                        case -2147483648:
                            fmt.Println(fi176, v9, v9)
                        }
                        _ = fi176
                    }
                    _ = v175
                }
                _ = fi162
            }
            _ = v158
        } else {
            if len(v4) > 0 {
                v4[0] = 210
            }
        }
    } else {
        fmt.Println(v157, v5, v5)
        v9 += v9
        v5 += v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("v157:", v157)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v155
    _ = v156
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
