package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int64; F2 int32 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 uint; F1 int; F2 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 int; F2 string }
type GenS12 struct { F0 int8; F1 int16 }
type GenS13 struct { F0 uint; F1 int8; F2 uint16 }
type GenS14 struct { F0 bool; F1 int16 }
type GenS15 struct { F0 bool }

func main() {
    v1 := []int{1, 127, -2147483648}
    v2 := make(map[uint32]int)
    v3 := []int{-209}
    v4 := make(map[uint32]string)
    var v5 bool = true
    if (!((((3.14159 + float64(float32(-2.25))) - (1.5 / -2.25)) + 0.0) == 1.5)) {
        fmt.Println(v5)
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 int = 40
            _ = fi6
            _ = v7
        }
        var v8 int8 = -39
        _ = v8
    }
    v9 := GenS1{F0: -2147483648, F1: ("foo bar" + "")}
    if (4294967295 < 550) {
        if v5 {
            v10 := GenS2{F0: 32767}
            switch 0 {
            case 32767:
                _ = len(v1)
            case -1:
                for fi11 := 0; fi11 < 19; fi11++ {
                    if (v5 || v5) {
                        v5 = v5
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v12 bool = v5
                        _ = v12
                    }
                    _ = fi11
                }
            case 2147483647:
                var v13 uint64 = 0
                _ = v13
            default:
                var v14 uint32 = 599
                _ = v14
            }
            v5 = v5
            _ = v10
        }
    } else {
        var v15 int8 = -128
        _ = v15
    }
    if (false || v5) {
        switch "" {
        case "hello":
            v5 = (!(100 < 1))
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v5)
    switch "a" {
    case "\n\t":
        if v5 {
            if (!true) {
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch 127 {
                    case 9223372036854775807:
                        v17 := GenS3{F0: -2147483648}
                        _ = v17
                    case 0:
                        if v5 {
                            if len(v1) > 0 {
                                v1[0] = -809
                            }
                            if (926 <= 4294967295) {
                                fmt.Println(fi16, v5)
                                v5 = (-24 >= 586)
                            } else {
                                var v18 bool = (false || v5)
                                for fi19 := 0; fi19 < 10; fi19++ {
                                    var v20 string = "a"
                                    var v21 float64 = 0.0
                                    switch fi16 {
                                    case 127:
                                        v21 = v21
                                    case 4294967295:
                                        var v22 float32 = float32(-2.25)
                                        _ = v22
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            v5 = v5
                                        case "":
                                            v23 := GenS4{F0: 328, F1: 42, F2: 0}
                                            _ = v23
                                        }
                                    }
                                    var v24 string = v20
                                    var v25 string = "foo bar"
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                    _ = v24
                                    _ = v25
                                }
                                _ = v18
                            }
                            v26 := GenS5{F0: ((false && v5) && v5)}
                            _ = v26
                        } else {
                            switch 42 {
                            case 369:
                                switch fi16 {
                                case 1:
                                    var v27 string = "\n\t"
                                    _ = v27
                                case 42:
                                    var v28 uint64 = 0
                                    _ = v28
                                default:
                                    if ((float64(1.5) * 0.1) > 0.0) {
                                        v29 := v2[1]
                                        _ = v29
                                    }
                                }
                            case 32767:
                                v4[42] = ((("" + "a") + ("" + "")) + "")
                            case 1:
                                var v30 int64 = 1
                                _ = v30
                            default:
                                delete(v4, 0)
                            }
                            var v31 string = ""
                            _ = v31
                        }
                    case 128:
                        var v32 uint32 = 42
                        _ = v32
                    default:
                        fmt.Println(v5, v5)
                    }
                    var v33 float64 = float64((0.0 / float32(-2.25)))
                    for fi34 := 0; fi34 < 16; fi34++ {
                        var v35 float32 = (0.1 - 1e10)
                        _ = fi34
                        _ = v35
                    }
                    _ = fi16
                    _ = v33
                }
                _ = len(v3)
            }
        }
    case "":
        if v5 {
            if v5 {
                v5 = false
                v5 = (!(310 != 42))
                if (128 <= -2147483648) {
                    var v36 uint32 = 984
                    v37 := GenS2{F0: -946}
                    _ = v36
                    _ = v37
                } else {
                    var v38 float64 = ((((0.0 + 0.1) + float64(-2.25)) / 1.5) * (3.14159 * 1.5))
                    v38 = v38
                    if v5 {
                        for fi39 := 0; fi39 < 12; fi39++ {
                            for fi40 := 0; fi40 < 17; fi40++ {
                                switch fi39 {
                                case 4294967295:
                                    v38 = (((float64(((1.5 - 3.14159) / 1.0)) / 1.0) - v38) * (v38 / (3.14159 - (1e10 * 0.0))))
                                case 9223372036854775807:
                                    var v41 int32 = -1
                                    _ = v41
                                }
                                switch "a" {
                                case "foo bar":
                                    var v42 uint16 = 32767
                                    _ = v42
                                case "a":
                                    var v43 uint16 = 42
                                    _ = v43
                                case "\n\t":
                                    fmt.Println(v38)
                                }
                                _ = fi40
                            }
                            if true {
                                delete(v4, 1)
                                v5 = true
                                if v5 {
                                    v44 := GenS6{F0: "hello", F1: 0, F2: -2147483648}
                                    _ = v44
                                }
                            }
                            switch (-(fi39)) {
                            case 4294967295:
                                var v45 uint8 = 0
                                _ = v45
                            default:
                                for fi46 := 0; fi46 < 4; fi46++ {
                                    var v47 string = "foo bar"
                                    v5 = v5
                                    v47 = ((("hello" + "hello") + (v47 + ("hello" + v47))) + (("" + v47) + (v47 + "\n\t")))
                                    _ = fi46
                                    _ = v47
                                }
                            }
                            _ = fi39
                        }
                        if v5 {
                            if false {
                                switch 809 {
                                case -566:
                                    for fi48 := 0; fi48 < 2; fi48++ {
                                        switch ("hello" + "") {
                                        case "a":
                                            if true {
                                                delete(v4, 4294967295)
                                            }
                                        case "\n\t":
                                            if (1 < 42) {
                                                v5 = ("a" != "hello")
                                                fmt.Println(fi48, v5)
                                                for fi49 := 0; fi49 < 8; fi49++ {
                                                    v38 = ((v38 / v38) / v38)
                                                    _ = fi49
                                                }
                                            } else {
                                                var v50 uint = 18446744073709551615
                                                v51 := GenS5{F0: v5}
                                                _ = v50
                                                _ = v51
                                            }
                                        case "foo bar":
                                            v52 := GenS7{F0: 41, F1: 2}
                                            _ = v52
                                        }
                                        var v53 int32 = 515
                                        _ = fi48
                                        _ = v53
                                    }
                                case 32768:
                                    var v54 string = "hello"
                                    _ = v54
                                default:
                                    var v55 int = 0
                                    _ = v55
                                }
                            }
                        }
                        var v56 uint64 = 762
                        _ = v56
                    }
                    _ = v38
                }
            } else {
                var v57 float32 = (0.1 + (((1.5 * 1e10) / 1.5) * 1e10))
                var v58 uint = 65535
                v58 = (^(v58))
                _ = v57
                _ = v58
            }
        }
    default:
        v5 = false
    }
    var v59 int16 = -32768
    switch 837 {
    case -152:
        switch -2147483648 {
        case 548:
            for fi60 := 0; fi60 < 8; fi60++ {
                switch (fi60 >> ((255) & 63)) {
                case 2147483647:
                    if (!v5) {
                        for fi61 := 0; fi61 < 15; fi61++ {
                            switch fi60 {
                            case -9223372036854775808:
                                if false {
                                    v59 = (v59 + v59)
                                }
                            case -393:
                                v59 = v59
                            default:
                                if (false && false) {
                                    var v62 int8 = 68
                                    switch fi61 {
                                    case 957:
                                        switch ("" + ("hello" + "hello")) {
                                        case "\n\t":
                                            v63 := GenS8{F0: 2, F1: fi60, F2: ("a" + "a")}
                                            _ = v63
                                        case "a":
                                            v64 := GenS9{F0: 255}
                                            _ = v64
                                        case "hello":
                                            v5 = true
                                        default:
                                            if (v5 && true) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                var v65 string = "\n\t"
                                                _ = v65
                                            }
                                        }
                                    case 32767:
                                        v62 += v62
                                    default:
                                        v62 -= 42
                                    }
                                    fmt.Println(fi61)
                                    _ = v62
                                }
                                if (!v5) {
                                    switch "" {
                                    case "":
                                        v5 = v5
                                    default:
                                        v66 := GenS3{F0: 489}
                                        _ = v66
                                    }
                                    var v67 uint16 = 32767
                                    v68 := GenS10{F0: (v67 >> ((977) & 63))}
                                    _ = v67
                                    _ = v68
                                } else {
                                    if (!(false || v5)) {
                                        fmt.Println(fi60)
                                        var v69 uint16 = 32767
                                        var v70 uint16 = v69
                                        _ = v69
                                        _ = v70
                                    }
                                    switch fi60 {
                                    case 32767:
                                        v59 -= (v59 / (((^(v59))) | 1))
                                    case 1:
                                        var v71 int64 = -91
                                        _ = v71
                                    default:
                                        var v72 uint64 = 1
                                        var v73 uint32 = 2
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = len(v4)
                                }
                            }
                            _ = fi61
                        }
                    } else {
                        if (-219 < -300) {
                            var v74 uint32 = 4294967295
                            v3 = append(v3, (fi60 ^ int(v59)))
                            if false {
                                v75 := v2[v74]
                                var v76 int64 = -1
                                _ = v75
                                _ = v76
                            } else {
                                v77 := GenS11{F0: -91, F1: fi60, F2: "\n\t"}
                                _ = v77
                            }
                            _ = v74
                        } else {
                            var v78 float64 = 0.1
                            _ = v78
                        }
                    }
                case -9223372036854775808:
                    delete(v2, 0)
                default:
                    if (18446744073709551615 == 700) {
                        fmt.Println(fi60, fi60, v59)
                    }
                }
                delete(v2, 42)
                v5 = (66 < 127)
                _ = fi60
            }
        case 1:
            v1 = append(v1, 128)
        case 376:
            fmt.Println(v59)
        default:
            fmt.Println(v59, v59)
        }
    case 9223372036854775807:
        delete(v4, 1)
    }
    v5 = v5
    v5 = (v5 && (!true))
    v79 := GenS2{F0: v59}
    var v80 int16 = -1
    var v81 uint64 = 630
    var v82 string = ""
    if (48 != 42) {
        v81 = (v81 ^ v81)
        if v5 {
            if true {
                for fi83 := 0; fi83 < 2; fi83++ {
                    v82 = v82
                    v81 = 43
                    if (42 != 0) {
                        var v84 int64 = -188
                        var v85 int64 = v84
                        _ = v84
                        _ = v85
                    }
                    _ = fi83
                }
                var v86 int32 = -769
                var v87 string = (v82 + "\n\t")
                _ = v86
                _ = v87
            } else {
                fmt.Println(v82, v81)
            }
        }
        for fi88 := 0; fi88 < 16; fi88++ {
            v89 := GenS9{F0: 2}
            v3 = append(v3, fi88)
            for fi90 := 0; fi90 < 14; fi90++ {
                fmt.Println(v80, v82)
                if false {
                    if true {
                        v91 := GenS12{F0: 1, F1: (^(v59))}
                        _ = v91
                    }
                    if true {
                        var v92 uint32 = 900
                        v92 = v92
                        _ = v92
                    } else {
                        v5 = v5
                        fmt.Println(v82)
                        if ((482 != (^(fi88))) && (!(v5 && (v82 < "")))) {
                            v5 = ((v80 - 32767) >= v80)
                        } else {
                            fmt.Println(fi88, v5, v5)
                        }
                    }
                    v93 := GenS13{F0: uint(v81), F1: -1, F2: 264}
                    _ = v93
                }
                var v94 int32 = -144
                _ = fi90
                _ = v94
            }
            _ = fi88
            _ = v89
        }
    } else {
        v95 := GenS9{F0: 340}
        for fi96 := 0; fi96 < 12; fi96++ {
            switch fi96 {
            case -171:
                fmt.Println(v80, v59, fi96)
            case 32767:
                for fi97 := 0; fi97 < 6; fi97++ {
                    v5 = v5
                    v81 += v81
                    _ = fi97
                }
            case 1:
                for fi98 := 0; fi98 < 17; fi98++ {
                    var v99 float32 = (1.5 / float32((0.1 + 0.1)))
                    var v100 int8 = 127
                    _ = fi98
                    _ = v100
                    _ = v99
                }
            }
            if (-128 >= -63) {
                v81 -= (^(v81))
                v82 = (v82 + ("" + "a"))
            } else {
                _ = len(v3)
                if v5 {
                    for fi101 := 0; fi101 < 19; fi101++ {
                        if ((-855 >= v80) || v5) {
                            _ = len(v1)
                        } else {
                            v102 := GenS9{F0: 0}
                            _ = v102
                        }
                        if v5 {
                            v5 = v5
                            v5 = (!false)
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 43
                            }
                        }
                        _ = fi101
                    }
                    fmt.Println(v80)
                    v103 := GenS11{F0: 881, F1: fi96, F2: v82}
                    _ = v103
                } else {
                    v81 = (^(v81))
                    var v104 int8 = 0
                    _ = v104
                }
                switch -1 {
                case -84:
                    switch fi96 {
                    case -1:
                        v59 -= (v59 >> ((801) & 63))
                    case -139:
                        v105 := GenS14{F0: (!(v5 && (65535 != 348))), F1: -32768}
                        _ = v105
                    case 128:
                        for fi106 := 0; fi106 < 4; fi106++ {
                            v3 = append(v3, -9223372036854775808)
                            var v107 uint64 = v81
                            _ = fi106
                            _ = v107
                        }
                    }
                case -9223372036854775808:
                    var v108 uint8 = 1
                    _ = v108
                default:
                    v109 := GenS15{F0: v5}
                    _ = v109
                }
            }
            v82 = ((("\n\t" + "\n\t") + ("foo bar" + v82)) + "a")
            _ = fi96
        }
        _ = v95
    }
    if (4294967295 == 614) {
        v110 := GenS2{F0: v59}
        _ = v110
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v9
}
