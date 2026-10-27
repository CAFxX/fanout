package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int64; F2 int32 }
type GenS7 struct { F0 bool; F1 uint8; F2 string }
type GenS8 struct { F0 uint64; F1 int64; F2 string }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int; F1 int32; F2 int16 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int64; F1 uint8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint8 }
func genFn0(p0 string, p1 int8, p2 uint16) uint8 {
    v5 := []int{332, 4294967295, 42, -2147483648}
    p0 = (p0 + "foo bar")
    switch 127 {
    case 1:
        if false {
            var v6 string = p0
            if true {
                v7 := GenS1{F0: -9223372036854775808, F1: "a", F2: (p1 - -1)}
                for fi8 := 0; fi8 < 15; fi8++ {
                    var v9 string = v6
                    if (true && true) {
                        p1 = 114
                        if len(v5) > 0 {
                            v5[0] = (fi8 << ((65535) & 63))
                        }
                        if (("" + ("foo bar" + "")) == p0) {
                            v9 = (v9 + "")
                            for fi10 := 0; fi10 < 8; fi10++ {
                                v11 := GenS1{F0: fi10, F1: "\n\t", F2: p1}
                                v6 = ""
                                _ = fi10
                                _ = v11
                            }
                            v5 = append(v5, 1)
                        }
                    } else {
                        if (true || ((false || false) || (!(true || false)))) {
                            _ = len(v5)
                        } else {
                            if false {
                                switch p0 {
                                case "foo bar":
                                    for fi12 := 0; fi12 < 19; fi12++ {
                                        switch 0 {
                                        case 127:
                                            v13 := GenS1{F0: fi12, F1: "hello", F2: 1}
                                            _ = v13
                                        case -9223372036854775808:
                                            v14 := GenS2{F0: uint(p2)}
                                            _ = v14
                                        default:
                                            var v15 int8 = -103
                                            _ = v15
                                        }
                                        _ = fi12
                                    }
                                case "hello":
                                    v16 := GenS3{F0: ((1e10 == 1e10) && (true && true))}
                                    _ = v16
                                case "a":
                                    v9 = v6
                                default:
                                    switch fi8 {
                                    case -9223372036854775808:
                                        v17 := GenS2{F0: 2}
                                        _ = v17
                                    case 4294967295:
                                        _ = len(v5)
                                    }
                                }
                                v5 = append(v5, fi8)
                            }
                            var v18 string = p0
                            p0 = "a"
                            _ = v18
                        }
                        fmt.Println(p1, v6, p2)
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            p1 -= p1
                            v19 := GenS3{F0: true}
                            _ = v19
                        }
                    }
                    _ = fi8
                    _ = v9
                }
                _ = len(v5)
                _ = v7
            }
            _ = v6
        } else {
            if (!(0 > 42)) {
                v20 := GenS4{F0: p1, F1: (p2 / ((2) | 1)), F2: (!(uint(p2) != 18446744073709551615))}
                if true {
                    if (false && true) {
                        var v21 uint8 = 42
                        var v22 float64 = (3.14159 * -2.25)
                        p2 = uint16(v21)
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS4{F0: (p1 - (p1 - (-(p1)))), F1: p2, F2: true}
                    _ = v23
                }
                var v24 float32 = 3.14159
                _ = v20
                _ = v24
            } else {
                if true {
                    v25 := GenS3{F0: false}
                    var v26 float64 = (((1e10 / 1.5) + 1e10) / ((float64(0.0) * (1e10 * 0.1)) + (0.1 * 3.14159)))
                    _ = v25
                    _ = v26
                }
                if (!(!(false || false))) {
                    var v27 int = 32768
                    var v28 uint64 = 1
                    v29 := GenS4{F0: 127, F1: uint16(v28), F2: ((false && false) && false)}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    var v30 float32 = 0.0
                    for fi31 := 0; fi31 < 3; fi31++ {
                        fmt.Println(p2, fi31, p2)
                        v32 := GenS5{F0: int(p2)}
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                if (true && (!false)) {
                    switch 1 {
                    case 42:
                        p2 = p2
                    case 1:
                        switch p0 {
                        case "\n\t":
                            _ = len(v5)
                        case "":
                            v33 := GenS2{F0: 255}
                            _ = v33
                        case "a":
                            v34 := GenS3{F0: (p0 > p0)}
                            _ = v34
                        }
                    }
                    switch -1 {
                    case 0:
                        switch p0 {
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = 692
                            }
                        }
                    case -737:
                        v35 := GenS6{F0: 923, F1: 1, F2: -1}
                        _ = v35
                    default:
                        if (245 != 199) {
                            v5 = append(v5, 2147483647)
                        }
                    }
                    for fi36 := 0; fi36 < 4; fi36++ {
                        v5 = append(v5, (fi36 + (fi36 - (fi36 << ((233) & 63)))))
                        _ = fi36
                    }
                } else {
                    fmt.Println(p2)
                    for fi37 := 0; fi37 < 9; fi37++ {
                        var v38 uint = 65535
                        switch 63 {
                        case -9223372036854775808:
                            if len(v5) > 0 {
                                v5[0] = fi37
                            }
                        case -672:
                            _ = len(v5)
                        default:
                            if len(v5) > 0 {
                                v5[0] = 127
                            }
                        }
                        _ = fi37
                        _ = v38
                    }
                    v39 := GenS1{F0: 0, F1: (((p0 + "\n\t") + p0) + ((("hello" + (p0 + ("\n\t" + p0))) + p0) + p0)), F2: (-(p1))}
                    _ = v39
                }
            }
            v40 := GenS2{F0: 18446744073709551615}
            _ = len(v5)
            _ = v40
        }
    case 838:
        p1 = p1
    }
    if true {
        for fi41 := 0; fi41 < 17; fi41++ {
            if len(v5) > 0 {
                v5[0] = fi41
            }
            v5 = append(v5, fi41)
            _ = fi41
        }
        _ = len(v5)
        switch 127 {
        case 443:
            p0 = p0
        case 4294967295:
            switch p0 {
            case "\n\t":
                if false {
                    _ = len(v5)
                } else {
                    if false {
                        for fi42 := 0; fi42 < 6; fi42++ {
                            if (!true) {
                                p2 += 255
                            } else {
                                if (2147483647 > -1) {
                                    fmt.Println(p1)
                                } else {
                                    p2 += 0
                                    if (!(-2147483648 != -1)) {
                                        if (75 != 0) {
                                            for fi43 := 0; fi43 < 14; fi43++ {
                                                var v44 string = p0
                                                v45 := GenS1{F0: (-(fi42)), F1: p0, F2: p1}
                                                _ = fi43
                                                _ = v44
                                                _ = v45
                                            }
                                        }
                                        for fi46 := 0; fi46 < 3; fi46++ {
                                            var v47 bool = true
                                            if ((fi46 / (((-(fi46))) | 1)) <= (fi42 << ((18446744073709551615) & 63))) {
                                                v5 = append(v5, -862)
                                                if (!true) {
                                                    switch 476 {
                                                    case -660:
                                                        p1 = p1
                                                    case 805:
                                                        var v48 int64 = 42
                                                        _ = v48
                                                    }
                                                    v47 = ((0.0 == 3.14159) && (2 == 18446744073709551615))
                                                    var v49 float32 = 0.0
                                                    p0 = ("\n\t" + (p0 + p0))
                                                    if (false && (((3.14159 + ((0.1 + 0.1) + ((1.5 * -2.25) + ((0.1 - 0.1) + 0.1)))) + -2.25) > -2.25)) {
                                                        switch p0 {
                                                        case "foo bar":
                                                            var v50 int16 = -1
                                                            _ = v50
                                                        case "":
                                                            var v51 uint64 = 0
                                                            _ = v51
                                                        default:
                                                            if v47 {
                                                                p1 += (-(p1))
                                                                var v52 int = (fi42 % (((fi42 ^ fi46)) | 1))
                                                                _ = v52
                                                            } else {
                                                                p0 = (p0 + "\n\t")
                                                                v53 := GenS6{F0: (p2 + (p2 - p2)), F1: -92, F2: 42}
                                                                if false {
                                                                    if (!v47) {
                                                                        var v54 int = fi46
                                                                        _ = v54
                                                                    } else {
                                                                        if v47 {
                                                                            p2 -= (p2 + (p2 % (((p2 * p2)) | 1)))
                                                                        }
                                                                        switch fi46 {
                                                                        case 0:
                                                                            var v55 uint64 = 42
                                                                            _ = v55
                                                                        case 430:
                                                                            if (32767 != 1000) {
                                                                                if v47 {
                                                                                    p2 = (p2 + p2)
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = -1
                                                                                    }
                                                                                    v56 := GenS4{F0: -128, F1: (p2 ^ p2), F2: true}
                                                                                    var v57 uint = 42
                                                                                    var v58 float64 = (0.0 / 1.0)
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                } else {
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    if v47 {
                                                                                        fmt.Println(fi46)
                                                                                        _ = len(v5)
                                                                                        fmt.Println(v49, fi42, p1)
                                                                                    }
                                                                                    fmt.Println(p1, v49, fi42)
                                                                                }
                                                                                if ((true || true) && v47) {
                                                                                    p1 = -1
                                                                                    var v59 uint32 = 995
                                                                                    _ = v59
                                                                                } else {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = -1
                                                                                    }
                                                                                    var v60 uint16 = p2
                                                                                    _ = v60
                                                                                }
                                                                                switch fi42 {
                                                                                case 0:
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    case "":
                                                                                        var v61 int = 36
                                                                                        _ = v61
                                                                                    case "a":
                                                                                        v62 := GenS6{F0: 16, F1: 1, F2: 1}
                                                                                        _ = v62
                                                                                    }
                                                                                    v63 := GenS2{F0: 2}
                                                                                    var v64 uint32 = 2147483647
                                                                                    _ = v63
                                                                                    _ = v64
                                                                                case 42:
                                                                                    fmt.Println(p0)
                                                                                default:
                                                                                    switch (fi42 | fi46) {
                                                                                    case -881:
                                                                                        if true {
                                                                                            v49 += v49
                                                                                            _ = len(v5)
                                                                                        } else {
                                                                                            var v65 bool = v47
                                                                                            var v66 int64 = 757
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = fi46
                                                                                            }
                                                                                            _ = v65
                                                                                            _ = v66
                                                                                        }
                                                                                    case -22:
                                                                                        p2 = p2
                                                                                    case 567:
                                                                                        if true {
                                                                                            var v67 uint16 = 1
                                                                                            _ = len(v5)
                                                                                            _ = v67
                                                                                        } else {
                                                                                            if (0 == 42) {
                                                                                                var v68 int16 = -1
                                                                                                if v47 {
                                                                                                    v69 := GenS3{F0: (v47 || v47)}
                                                                                                    if (v49 > (1e10 * v49)) {
                                                                                                        v49 += 0.1
                                                                                                        p2 += (^(p2))
                                                                                                        if (42 <= p1) {
                                                                                                            fmt.Println(v49, v68)
                                                                                                            v70 := GenS7{F0: (0 != 0), F1: 42, F2: p0}
                                                                                                            _ = v70
                                                                                                        }
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            if len(v5) > 0 {
                                                                                                                v5[0] = 173
                                                                                                            }
                                                                                                            v5 = append(v5, fi42)
                                                                                                            switch (("hello" + "hello") + "\n\t") {
                                                                                                            case "foo bar":
                                                                                                                if (true || v47) {
                                                                                                                    v47 = (true && false)
                                                                                                                } else {
                                                                                                                    v68 = v68
                                                                                                                    if false {
                                                                                                                        var v71 int64 = -9223372036854775808
                                                                                                                        fmt.Println(v47, p0, v71)
                                                                                                                        if true {
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                            if (v47 || (480 >= 4294967295)) {
                                                                                                                                if v47 {
                                                                                                                                    v72 := GenS5{F0: (fi42 - fi42)}
                                                                                                                                    var v73 float32 = (3.14159 + (float32(3.14159) + 1e10))
                                                                                                                                    _ = v72
                                                                                                                                    _ = v73
                                                                                                                                } else {
                                                                                                                                    var v74 float32 = v49
                                                                                                                                    v49 -= 3.14159
                                                                                                                                    var v75 bool = v47
                                                                                                                                    _ = v74
                                                                                                                                    _ = v75
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                p2 -= p2
                                                                                                                            }
                                                                                                                            switch p0 {
                                                                                                                            case "\n\t":
                                                                                                                                var v76 float64 = 0.0
                                                                                                                                _ = v76
                                                                                                                            case "hello":
                                                                                                                                if true {
                                                                                                                                    if v47 {
                                                                                                                                        fmt.Println(v47, v49)
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    v77 := GenS1{F0: fi42, F1: p0, F2: 127}
                                                                                                                                    _ = v77
                                                                                                                                }
                                                                                                                            case "a":
                                                                                                                                v78 := GenS7{F0: true, F1: 236, F2: (p0 + (p0 + p0))}
                                                                                                                                _ = v78
                                                                                                                            default:
                                                                                                                                if (!v47) {
                                                                                                                                    v71 += v71
                                                                                                                                    var v79 uint32 = 0
                                                                                                                                    var v80 int16 = v68
                                                                                                                                    _ = v79
                                                                                                                                    _ = v80
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v49)
                                                                                                                                    fmt.Println(p0, v68)
                                                                                                                                    p1 = 49
                                                                                                                                    v47 = (847 < 931)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if v47 {
                                                                                                                                if len(v5) > 0 {
                                                                                                                                    v5[0] = (fi42 | fi46)
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v81 uint32 = 65535
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        fmt.Println(fi46, p0)
                                                                                                                        if (!(false || v47)) {
                                                                                                                            v47 = false
                                                                                                                            if v47 {
                                                                                                                                v82 := GenS8{F0: 0, F1: (v71 * (v71 >> ((1) & 63))), F2: ""}
                                                                                                                                _ = v82
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v71
                                                                                                                    }
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                if (127 < 255) {
                                                                                                                    var v83 string = p0
                                                                                                                    v5 = append(v5, (fi42 + (-(fi42))))
                                                                                                                    fmt.Println(v83, fi42)
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v5 = append(v5, fi46)
                                                                                                            default:
                                                                                                                v49 += v49
                                                                                                            }
                                                                                                        }
                                                                                                        var v84 uint = 255
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    v5 = append(v5, (fi46 - fi46))
                                                                                                    var v85 string = p0
                                                                                                    _ = v69
                                                                                                    _ = v85
                                                                                                }
                                                                                                var v86 int8 = p1
                                                                                                _ = v68
                                                                                                _ = v86
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch (("a" + ("a" + (p0 + "a"))) + p0) {
                                                                                        case "\n\t":
                                                                                            p0 = ((p0 + p0) + ((("hello" + p0) + ((p0 + p0) + "foo bar")) + (p0 + (((((p0 + p0) + (("hello" + p0) + (p0 + (p0 + (p0 + p0))))) + (p0 + (p0 + p0))) + (p0 + ((p0 + "") + "a"))) + ""))))
                                                                                        case "a":
                                                                                            fmt.Println(p0, p0, v47)
                                                                                        case "":
                                                                                            v47 = false
                                                                                        default:
                                                                                            v87 := GenS9{F0: 42, F1: false}
                                                                                            _ = v87
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    var v88 float64 = (((0.1 / (1e10 * -2.25)) - 0.0) * 0.1)
                                                                                    if (42 != 401) {
                                                                                        if ((3.14159 / 0.1) <= v88) {
                                                                                            if (fi46 <= (fi42 & (fi46 & (fi46 | fi42)))) {
                                                                                                v89 := GenS10{F0: 770}
                                                                                                v90 := GenS11{F0: 32767, F1: 1, F2: 32767}
                                                                                                fmt.Println(v49, v47)
                                                                                                if (v47 || (false && v47)) {
                                                                                                    if v47 {
                                                                                                        fmt.Println(fi46, p1)
                                                                                                        v91 := GenS12{F0: 42}
                                                                                                        v47 = v47
                                                                                                        _ = v91
                                                                                                    }
                                                                                                } else {
                                                                                                    var v92 uint16 = 2
                                                                                                    p1 = 86
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                    _ = v92
                                                                                                }
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            } else {
                                                                                                fmt.Println(fi46, v49, p2)
                                                                                                v93 := GenS13{F0: 2, F1: 4294967295, F2: 42}
                                                                                                v49 = v49
                                                                                                _ = v93
                                                                                            }
                                                                                            var v94 uint = 255
                                                                                            _ = v94
                                                                                        } else {
                                                                                            v88 = v88
                                                                                            var v95 int32 = 964
                                                                                            _ = v95
                                                                                        }
                                                                                        v49 -= (-2.25 / (v49 - v49))
                                                                                    } else {
                                                                                        v49 -= v49
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (fi42 / (((fi46 + -2147483648)) | 1))
                                                                                    }
                                                                                    _ = v88
                                                                                } else {
                                                                                    switch (fi46 >> ((1) & 63)) {
                                                                                    case -2147483648:
                                                                                        _ = len(v5)
                                                                                    case 32768:
                                                                                        fmt.Println(p2, v49)
                                                                                    case 579:
                                                                                        var v96 int8 = -119
                                                                                        _ = v96
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 9223372036854775807:
                                                                            if false {
                                                                                if v47 {
                                                                                    v97 := GenS14{F0: p0}
                                                                                    _ = v97
                                                                                } else {
                                                                                    var v98 int64 = 1
                                                                                    _ = v98
                                                                                }
                                                                                v99 := GenS4{F0: (-(p1)), F1: p2, F2: v47}
                                                                                v100 := GenS15{F0: 100}
                                                                                _ = v100
                                                                                _ = v99
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi46, v49, v47)
                                                                        }
                                                                        v101 := GenS3{F0: v47}
                                                                        var v102 uint16 = 32767
                                                                        if v47 {
                                                                            var v103 int = (fi46 & (fi46 - fi46))
                                                                            _ = v103
                                                                        } else {
                                                                            v5 = append(v5, 128)
                                                                            fmt.Println(p0)
                                                                            var v104 int64 = 0
                                                                            _ = v104
                                                                        }
                                                                        if v47 {
                                                                            v5 = append(v5, (fi46 + (fi42 / ((451) | 1))))
                                                                            if (0 == -803) {
                                                                                switch ("foo bar" + (("" + "") + p0)) {
                                                                                case "a":
                                                                                    fmt.Println(fi42, v49)
                                                                                    switch 128 {
                                                                                    case -856:
                                                                                        var v105 uint32 = 65535
                                                                                        var v106 int32 = -796
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                    case 209:
                                                                                        if (-32768 >= -10) {
                                                                                            if (750 == 42) {
                                                                                                var v107 int64 = 42
                                                                                                v47 = v47
                                                                                                _ = len(v5)
                                                                                                _ = v107
                                                                                            }
                                                                                            var v108 uint32 = 913
                                                                                            _ = v108
                                                                                        } else {
                                                                                            var v109 uint8 = 1
                                                                                            var v110 uint16 = 0
                                                                                            _ = v109
                                                                                            _ = v110
                                                                                        }
                                                                                    case 0:
                                                                                        var v111 int32 = 0
                                                                                        _ = v111
                                                                                    }
                                                                                    v47 = true
                                                                                    v112 := GenS2{F0: 2}
                                                                                    _ = v112
                                                                                default:
                                                                                    fmt.Println(v102, p1)
                                                                                }
                                                                                v113 := GenS6{F0: (v102 + 1), F1: 42, F2: 873}
                                                                                var v114 uint16 = (v102 >> ((42) & 63))
                                                                                _ = v113
                                                                                _ = v114
                                                                            } else {
                                                                                p2 = v102
                                                                                v115 := GenS16{F0: 0, F1: 1}
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        switch 4294967295 {
                                                                        case 2147483647:
                                                                            if false {
                                                                                if (1 >= 18446744073709551615) {
                                                                                    var v116 int8 = p1
                                                                                    _ = v116
                                                                                } else {
                                                                                    var v117 int16 = -32768
                                                                                    _ = v117
                                                                                }
                                                                                _ = len(v5)
                                                                            } else {
                                                                                var v118 string = ((("\n\t" + p0) + p0) + p0)
                                                                                v102 -= (v102 ^ 1)
                                                                                _ = v118
                                                                            }
                                                                        default:
                                                                            v119 := GenS16{F0: 0, F1: 255}
                                                                            _ = v119
                                                                        }
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                    if true {
                                                                        switch ("\n\t" + p0) {
                                                                        case "foo bar":
                                                                            fmt.Println(p2, p0)
                                                                        case "a":
                                                                            v120 := GenS17{F0: (p2 << ((255) & 63))}
                                                                            _ = v120
                                                                        case "":
                                                                            if false {
                                                                                p0 = p0
                                                                            } else {
                                                                                var v121 string = "a"
                                                                                p0 = "hello"
                                                                                switch fi46 {
                                                                                case 1:
                                                                                    v122 := GenS4{F0: (p1 % ((p1) | 1)), F1: p2, F2: v47}
                                                                                    _ = v122
                                                                                case 9223372036854775807:
                                                                                    if v47 {
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        switch ("" + (("" + v121) + "a")) {
                                                                                        case "a":
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                        case "hello":
                                                                                            if (-2.25 > float64(-2.25)) {
                                                                                                p2 -= p2
                                                                                                if true {
                                                                                                    var v123 uint64 = 654
                                                                                                    switch "\n\t" {
                                                                                                    case "a":
                                                                                                        v47 = v47
                                                                                                    case "hello":
                                                                                                        var v124 uint32 = 42
                                                                                                        _ = v124
                                                                                                    case "\n\t":
                                                                                                        var v125 uint16 = (p2 | (p2 & p2))
                                                                                                        _ = v125
                                                                                                    }
                                                                                                    _ = v123
                                                                                                }
                                                                                                p1 -= p1
                                                                                            } else {
                                                                                                switch p0 {
                                                                                                case "\n\t":
                                                                                                    var v126 uint32 = 2147483647
                                                                                                    _ = v126
                                                                                                }
                                                                                                var v127 uint64 = 42
                                                                                                _ = v127
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v128 float64 = 3.14159
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            v129 := GenS11{F0: (fi46 / ((32768) | 1)), F1: int32(p1), F2: -32768}
                                                                                            _ = v128
                                                                                            _ = v129
                                                                                        default:
                                                                                            var v130 bool = (v47 && true)
                                                                                            _ = v130
                                                                                        }
                                                                                        fmt.Println(p2, v121)
                                                                                    } else {
                                                                                        v47 = v47
                                                                                        switch (p0 + p0) {
                                                                                        case "foo bar":
                                                                                            v49 = v49
                                                                                        case "":
                                                                                            v131 := GenS3{F0: false}
                                                                                            _ = v131
                                                                                        default:
                                                                                            p2 += (p2 & (p2 * p2))
                                                                                        }
                                                                                        var v132 int8 = 81
                                                                                        v132 = -123
                                                                                        p0 = ("" + v121)
                                                                                        _ = v132
                                                                                    }
                                                                                case 42:
                                                                                    v47 = (v47 || true)
                                                                                default:
                                                                                    p1 += (p1 & (^(p1)))
                                                                                }
                                                                                _ = v121
                                                                            }
                                                                        default:
                                                                            var v133 uint64 = 18446744073709551615
                                                                            _ = v133
                                                                        }
                                                                    }
                                                                    v47 = (1 == 73)
                                                                    if ((p0 + p0) > p0) {
                                                                        v49 -= v49
                                                                        var v134 uint = 42
                                                                        v135 := GenS13{F0: 255, F1: 253, F2: 128}
                                                                        _ = v134
                                                                        _ = v135
                                                                    } else {
                                                                        v47 = ((float32(1.5) * -2.25) == (((1.5 - v49) * float32((3.14159 / (((1e10 - float64(1.5)) * 0.1) + (0.1 * -2.25))))) * v49))
                                                                        fmt.Println(p0, v49)
                                                                        v47 = true
                                                                        p0 = p0
                                                                    }
                                                                    switch (fi42 & (fi46 - -2147483648)) {
                                                                    case -2147483648:
                                                                        var v136 uint64 = 18446744073709551615
                                                                        _ = v136
                                                                    case 127:
                                                                        v5 = append(v5, (fi42 >> ((0) & 63)))
                                                                    }
                                                                    p1 -= p1
                                                                } else {
                                                                    fmt.Println(fi42)
                                                                    fmt.Println(v47, fi46, p2)
                                                                }
                                                                _ = v53
                                                            }
                                                        }
                                                    }
                                                    var v137 uint8 = 8
                                                    _ = v137
                                                    _ = v49
                                                } else {
                                                    p2 -= 2
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi46 / (((fi42 >> ((18446744073709551615) & 63))) | 1))
                                                    }
                                                    _ = len(v5)
                                                }
                                                if (3.14159 < 0.0) {
                                                    switch (fi46 / (((^(fi46))) | 1)) {
                                                    case -9223372036854775808:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case 937:
                                                        var v138 uint32 = 994
                                                        _ = v138
                                                    default:
                                                        v139 := GenS7{F0: (65535 != 255), F1: 255, F2: (p0 + p0)}
                                                        _ = v139
                                                    }
                                                    var v140 bool = v47
                                                    v141 := GenS7{F0: v140, F1: 127, F2: p0}
                                                    _ = v140
                                                    _ = v141
                                                }
                                            }
                                            _ = fi46
                                            _ = v47
                                        }
                                    }
                                    fmt.Println(p0, p1, fi42)
                                }
                            }
                            var v142 bool = (42 != 585)
                            p2 -= p2
                            _ = fi42
                            _ = v142
                        }
                    }
                    if false {
                        switch 9223372036854775807 {
                        case 42:
                            if false {
                                p2 = (^(p2))
                                p0 = p0
                                for fi143 := 0; fi143 < 13; fi143++ {
                                    if (p1 >= (p1 + p1)) {
                                        v144 := GenS18{F0: 182}
                                        v145 := GenS11{F0: (fi143 >> ((0) & 63)), F1: 42, F2: 570}
                                        _ = v144
                                        _ = v145
                                    }
                                    _ = fi143
                                }
                            } else {
                                fmt.Println(p2, p2, p2)
                                var v146 float64 = ((-2.25 / -2.25) / -2.25)
                                _ = v146
                            }
                        default:
                            p1 = (-(p1))
                        }
                        v147 := GenS17{F0: 255}
                        _ = v147
                    }
                }
            }
        }
    } else {
        p0 = p0
    }
    p2 += 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 128
}
func genFn1() bool {
    var v148 float32 = ((1.5 / (float32((float64(-2.25) * 1e10)) * 3.14159)) - (3.14159 * (((3.14159 - 1e10) * 1e10) - (0.1 - -2.25))))
    v149 := []int{0, 681, 128, 911}
    _ = v148
    _ = v149
    return (!true)
}

func main() {
    var v1 int8 = -103
    var v2 float64 = (1e10 / 1.0)
    var v3 string = ("foo bar" + (("" + "a") + ""))
    v4 := make(map[uint32]uint32)
    var v150 int = 0
    v3 = "\n\t"
    _ = len(v4)
    var v151 uint16 = 0
    if true {
        fmt.Println(v3)
        var v152 uint32 = 2147483647
        v153 := []int{v150, v150, (v150 * (-(v150)))}
        _ = v152
        _ = v153
    }
    var v154 int8 = (v1 / ((-128) | 1))
    switch "foo bar" {
    case "a":
        fmt.Println(v154)
    case "hello":
        v155 := GenS12{F0: v151}
        _ = v155
    case "foo bar":
        v150 -= 32767
    default:
        v154 = v154
    }
    if (0 >= 1) {
        var v156 float32 = (-2.25 * 0.0)
        for fi157 := 0; fi157 < 19; fi157++ {
            if false {
                v2 = v2
                var v158 uint = 4294967295
                _ = v158
            } else {
                for fi159 := 0; fi159 < 6; fi159++ {
                    fmt.Println(v2)
                    v4[149] = 2147483647
                    switch (((v3 + "\n\t") + v3) + "\n\t") {
                    case "a":
                        v160 := GenS15{F0: 127}
                        _ = v160
                    }
                    var v161 uint16 = 65535
                    v3 = v3
                    _ = fi159
                    _ = v161
                }
            }
            for fi162 := 0; fi162 < 7; fi162++ {
                v154 += v1
                v163 := []int{(v150 >> ((996) & 63)), -50, v150, (fi162 << ((0) & 63))}
                v1 = (^(v1))
                fmt.Println(v150, v154, v156)
                _ = fi162
                _ = v163
            }
            _ = fi157
        }
        _ = v156
    } else {
        v3 = v3
    }
    fmt.Println(v2, v150)
    v164 := []int{(v150 - v150), v150, (v150 + 4294967295), -468}
    _ = len(v164)
    v165 := GenS4{F0: v1, F1: v151, F2: false}
    v154 = (v1 % (((v154 / (((v1 & 127)) | 1))) | 1))
    var v166 bool = false
    switch (v150 & int(v154)) {
    case 42:
        for fi167 := 0; fi167 < 8; fi167++ {
            if (0.1 == -2.25) {
                switch 478 {
                case 2147483647:
                    v4[2147483648] = uint32(v151)
                case -9223372036854775808:
                    v150 += fi167
                default:
                    v164 = append(v164, 1)
                }
                switch v150 {
                case 127:
                    v3 = (v3 + "a")
                case 2147483647:
                    if (!v166) {
                        v164 = append(v164, 128)
                    }
                case -751:
                    v2 = (1e10 + (-2.25 * -2.25))
                default:
                    var v168 int8 = -128
                    _ = v168
                }
            } else {
                switch fi167 {
                case 42:
                    v169 := v4[94]
                    _ = v169
                case -9223372036854775808:
                    fmt.Println(v166, v151)
                default:
                    v2 = v2
                }
                var v170 uint = 65535
                _ = v170
            }
            _ = fi167
        }
    }
    for fi171 := 0; fi171 < 3; fi171++ {
        for fi172 := 0; fi172 < 8; fi172++ {
            var v173 uint = 1
            v174 := GenS14{F0: v3}
            fmt.Println(v150, v150, v166)
            var v175 uint8 = 128
            var v176 bool = v166
            _ = fi172
            _ = v173
            _ = v174
            _ = v175
            _ = v176
        }
        var v177 float32 = ((1e10 - float32(1.5)) + 1.5)
        if (1 == 18446744073709551615) {
            for fi178 := 0; fi178 < 16; fi178++ {
                var v179 uint16 = v151
                _ = fi178
                _ = v179
            }
            v180 := GenS6{F0: v151, F1: 944, F2: int32(v150)}
            fmt.Println(v1, v2, v166)
            _ = v180
        }
        _ = fi171
        _ = v177
    }
    v151 = v151
    fmt.Println("v1:", v1)
    fmt.Println("v150:", v150)
    fmt.Println("v151:", v151)
    fmt.Println("v154:", v154)
    fmt.Println("slicelen v164:", len(v164), v164)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v150
    _ = v151
    _ = v154
    _ = v164
    _ = v165
    _ = v166
    _ = v2
    _ = v3
    _ = v4
}
