package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int16; F1 int8; F2 int32 }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int8; F1 string; F2 string }
type GenS11 struct { F0 int8; F1 uint }
func genFn0(p0 int) int16 {
    var v5 bool = (!((false || false) || true))
    for fi6 := 0; fi6 < 15; fi6++ {
        if (-2147483648 > 0) {
            if false {
                var v7 float32 = 3.14159
                if v5 {
                    switch fi6 {
                    case 128:
                        v5 = (1 >= -81)
                    case 1:
                        v8 := GenS1{F0: -217}
                        _ = v8
                    }
                } else {
                    var v9 int64 = 22
                    for fi10 := 0; fi10 < 17; fi10++ {
                        if (true && true) {
                            v11 := GenS2{F0: true}
                            _ = v11
                        } else {
                            var v12 string = "hello"
                            fmt.Println(v7)
                            _ = v12
                        }
                        _ = fi10
                    }
                    _ = v9
                }
                p0 += fi6
                _ = v7
            } else {
                if (-1 > 42) {
                    switch "" {
                    case "foo bar":
                        var v13 bool = v5
                        _ = v13
                    case "\n\t":
                        v5 = v5
                    }
                    var v14 int8 = 0
                    _ = v14
                }
            }
            fmt.Println(p0, v5)
            v15 := GenS3{F0: 0}
            _ = v15
        } else {
            var v16 uint8 = 1
            _ = v16
        }
        fmt.Println(fi6, v5, p0)
        _ = fi6
    }
    var v17 uint = 18446744073709551615
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS4{F0: -9223372036854775808}
        _ = fi18
        _ = v19
    }
    v20 := []int{1, p0}
    _ = p0
    _ = v17
    _ = v20
    _ = v5
    return -32768
}
func genFn1(p0 int) int64 {
    var v21 string = "hello"
    v21 = "hello"
    switch "hello" {
    case "":
        fmt.Println(v21)
    case "a":
        v22 := []int{(p0 ^ p0), 4294967295, (^(p0)), p0}
        _ = v22
    case "foo bar":
        v21 = "a"
    default:
        fmt.Println(p0, v21)
    }
    if (!(9223372036854775807 == p0)) {
        var v23 uint16 = 255
        v24 := []int{(p0 * (p0 ^ -9223372036854775808)), (p0 * (-(p0)))}
        for fi25 := 0; fi25 < 3; fi25++ {
            fmt.Println(v23, v21)
            if (0.1 != -2.25) {
                if (true || (!(!true))) {
                    switch p0 {
                    case -2147483648:
                        v26 := GenS5{F0: true, F1: -1}
                        _ = v26
                    case -1:
                        fmt.Println(p0, v21)
                    default:
                        if (((!false) || (!(255 <= 0))) && (true || (0.1 != 1.5))) {
                            var v27 float32 = 1e10
                            _ = v27
                        }
                    }
                    var v28 int32 = -637
                    if true {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            if false {
                                v30 := GenS3{F0: -499}
                                v21 = ((v21 + ("\n\t" + v21)) + v21)
                                _ = v30
                            }
                            var v31 int64 = -1
                            v32 := GenS6{F0: v31, F1: 1}
                            _ = fi29
                            _ = v31
                            _ = v32
                        }
                        var v33 float64 = (0.0 / 3.14159)
                        switch fi25 {
                        case 2147483647:
                            var v34 int = (fi25 + 9223372036854775807)
                            _ = v34
                        case -671:
                            var v35 int32 = -819
                            _ = v35
                        default:
                            v36 := GenS5{F0: false, F1: 1}
                            _ = v36
                        }
                        _ = v33
                    } else {
                        v21 = (((v21 + "a") + v21) + "foo bar")
                        for fi37 := 0; fi37 < 8; fi37++ {
                            switch (v21 + ((v21 + v21) + (v21 + v21))) {
                            case "":
                                var v38 uint = 18446744073709551615
                                _ = v38
                            case "hello":
                                if true {
                                    if false {
                                        var v39 int64 = -632
                                        _ = v39
                                    }
                                } else {
                                    v23 = 2
                                    _ = len(v24)
                                    v24 = append(v24, (fi25 ^ 289))
                                }
                            }
                            switch "" {
                            case "\n\t":
                                switch "" {
                                case "":
                                    fmt.Println(v23, v28, v28)
                                default:
                                    var v40 int8 = -1
                                    _ = v40
                                }
                            default:
                                v23 -= v23
                            }
                            _ = fi37
                        }
                    }
                    _ = v28
                } else {
                    p0 = fi25
                    var v41 int32 = 42
                    _ = v41
                }
                v21 = v21
            } else {
                fmt.Println(v21, p0)
                if true {
                    switch v21 {
                    case "a":
                        if false {
                            for fi42 := 0; fi42 < 18; fi42++ {
                                v21 = ((((("" + ((v21 + "foo bar") + v21)) + "") + ((v21 + v21) + "hello")) + v21) + ("a" + "foo bar"))
                                p0 = 1
                                _ = fi42
                            }
                        }
                    case "":
                        var v43 int16 = 1000
                        _ = v43
                    case "\n\t":
                        var v44 int64 = 250
                        _ = v44
                    }
                    v21 = v21
                }
            }
            _ = len(v24)
            _ = fi25
        }
        _ = v23
        _ = v24
    }
    var v45 uint = 375
    _ = p0
    _ = v21
    _ = v45
    return -887
}

func main() {
    var v1 int16 = 284
    v2 := []int{897, 127, -408, 9223372036854775807}
    var v3 uint64 = 571
    var v4 uint8 = 2
    var v46 int8 = 1
    v47 := GenS7{F0: (v1 % ((genFn0(32767)) | 1)), F1: v46, F2: 0}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        var v49 float32 = 1.5
        var v50 uint16 = 1
        _ = fi48
        _ = v49
        _ = v50
    }
    if false {
        var v51 uint32 = 779
        _ = v51
    } else {
        _ = len(v2)
        var v52 bool = ((v3 / ((v3) | 1)) < v3)
        for fi53 := 0; fi53 < 5; fi53++ {
            for fi54 := 0; fi54 < 18; fi54++ {
                var v55 uint8 = (v4 - v4)
                v3 = v3
                if (true && true) {
                    v56 := GenS8{F0: fi53, F1: 4294967295}
                    _ = v56
                } else {
                    if true {
                        if v52 {
                            v3 = (v3 + v3)
                            v1 = (v1 ^ v1)
                        } else {
                            fmt.Println(v52, v3)
                            var v57 uint64 = v3
                            _ = v57
                        }
                        v58 := GenS7{F0: int16(v3), F1: 50, F2: 42}
                        _ = v58
                    }
                    var v59 int16 = 263
                    v1 -= (v1 << ((2) & 63))
                    _ = v59
                }
                _ = fi54
                _ = v55
            }
            _ = fi53
        }
        _ = v52
    }
    switch "" {
    case "hello":
        if true {
            switch "foo bar" {
            case "":
                if (!false) {
                    _ = len(v2)
                    v1 -= v1
                    var v60 int16 = (-(v1))
                    _ = v60
                } else {
                    v3 -= (^(v3))
                    if true {
                        for fi61 := 0; fi61 < 8; fi61++ {
                            if true {
                                var v62 uint32 = uint32(v3)
                                v2 = append(v2, int(v62))
                                _ = v62
                            } else {
                                var v63 uint8 = 2
                                _ = v63
                            }
                            _ = fi61
                        }
                        v4 = 127
                    } else {
                        v46 = 47
                        switch -2147483648 {
                        case -675:
                            v3 = 18446744073709551615
                        case 128:
                            fmt.Println(v4, v4)
                        case 887:
                            v64 := GenS2{F0: true}
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 6; fi65++ {
                            if true {
                                v4 = v4
                            } else {
                                v66 := GenS9{F0: 32767, F1: 0}
                                _ = v66
                            }
                            v46 = (v46 + v46)
                            v2 = append(v2, fi65)
                            _ = fi65
                        }
                    }
                }
            case "hello":
                if (!false) {
                    v46 = -95
                    for fi67 := 0; fi67 < 11; fi67++ {
                        v68 := GenS10{F0: v46, F1: "", F2: ("\n\t" + ("hello" + ("a" + (((((("foo bar" + "a") + "\n\t") + "foo bar") + "") + "\n\t") + "\n\t"))))}
                        _ = fi67
                        _ = v68
                    }
                }
            }
            v4 = 42
        }
    case "\n\t":
        for fi69 := 0; fi69 < 15; fi69++ {
            var v70 uint32 = 0
            var v71 int8 = v46
            _ = fi69
            _ = v70
            _ = v71
        }
    }
    var v72 float64 = 1.5
    _ = len(v2)
    if false {
        _ = len(v2)
        v73 := GenS7{F0: (v1 << ((65535) & 63)), F1: 0, F2: 619}
        v74 := GenS8{F0: 1, F1: v3}
        _ = v73
        _ = v74
    } else {
        v4 += v4
        var v75 uint32 = 2147483648
        v76 := GenS5{F0: true, F1: 9223372036854775807}
        _ = v75
        _ = v76
    }
    var v77 uint64 = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += (v1 - v1)
    var v78 float64 = v72
    var v79 uint64 = v77
    if (32768 > 1) {
        if (2147483647 == -2147483648) {
            v1 = v1
        }
        if (254 < 18446744073709551615) {
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                v80 := GenS4{F0: 1}
                _ = v80
            case "":
                var v81 float32 = (1.5 + (1e10 * 1e10))
                _ = v81
            case "hello":
                v78 = 0.1
            default:
                var v82 float32 = 1e10
                _ = v82
            }
        } else {
            var v83 int32 = 0
            v84 := GenS11{F0: (v46 >> ((451) & 63)), F1: 2}
            if true {
                if false {
                    v4 -= v4
                    switch "a" {
                    case "a":
                        if true {
                            if (true && true) {
                                for fi85 := 0; fi85 < 19; fi85++ {
                                    if ("\n\t" < "\n\t") {
                                        v3 = (v3 ^ (v3 ^ 872))
                                        var v86 int8 = (v46 << ((834) & 63))
                                        if (9223372036854775807 != -799) {
                                            if true {
                                                var v87 int64 = -619
                                                if ((!(!false)) || false) {
                                                    var v88 int64 = (v87 >> ((255) & 63))
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        v2 = append(v2, fi85)
                                                    case "foo bar":
                                                        v1 += (v1 & (v1 + v1))
                                                    case "\n\t":
                                                        v3 -= v3
                                                    }
                                                    _ = v88
                                                } else {
                                                    var v89 uint8 = v4
                                                    _ = v89
                                                }
                                                switch (("foo bar" + "\n\t") + (("hello" + "\n\t") + "a")) {
                                                case "hello":
                                                    var v90 int = (fi85 & (fi85 * (fi85 + (^(fi85)))))
                                                    _ = v90
                                                case "\n\t":
                                                    for fi91 := 0; fi91 < 19; fi91++ {
                                                        var v92 uint64 = v3
                                                        v2 = append(v2, -2147483648)
                                                        var v93 int = 32768
                                                        _ = fi91
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                }
                                                _ = v87
                                            }
                                        }
                                        _ = v86
                                    }
                                    if (1 != 543) {
                                        for fi94 := 0; fi94 < 17; fi94++ {
                                            var v95 uint64 = v79
                                            _ = fi94
                                            _ = v95
                                        }
                                        fmt.Println(v79)
                                    } else {
                                        var v96 bool = (fi85 <= (fi85 / ((259) | 1)))
                                        _ = v96
                                    }
                                    _ = fi85
                                }
                                var v97 uint64 = v3
                                _ = v97
                            }
                            for fi98 := 0; fi98 < 2; fi98++ {
                                fmt.Println(v72, v4, v46)
                                for fi99 := 0; fi99 < 9; fi99++ {
                                    v4 -= v4
                                    if (1 > 131) {
                                        v100 := GenS2{F0: (525 > (v83 % ((42) | 1)))}
                                        v79 = v3
                                        _ = v100
                                    }
                                    v1 -= genFn0((fi98 % (((-(fi99))) | 1)))
                                    _ = fi99
                                }
                                _ = fi98
                            }
                            if len(v2) > 0 {
                                v2[0] = 640
                            }
                        }
                    case "hello":
                        v4 = 0
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v4 -= v4
                        }
                    }
                } else {
                    for fi101 := 0; fi101 < 17; fi101++ {
                        var v102 int16 = -991
                        var v103 int64 = -1
                        for fi104 := 0; fi104 < 15; fi104++ {
                            v1 += int16(v3)
                            var v105 string = ("foo bar" + "\n\t")
                            _ = len(v2)
                            _ = fi104
                            _ = v105
                        }
                        _ = fi101
                        _ = v102
                        _ = v103
                    }
                }
                for fi106 := 0; fi106 < 12; fi106++ {
                    if (-9223372036854775808 >= 300) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v3 += v77
                        v107 := GenS2{F0: false}
                        _ = v107
                    }
                    _ = fi106
                }
                _ = len(v2)
            } else {
                var v108 uint = 4294967295
                _ = v108
            }
            _ = v83
            _ = v84
        }
        for fi109 := 0; fi109 < 18; fi109++ {
            v110 := GenS10{F0: v46, F1: "a", F2: ""}
            for fi111 := 0; fi111 < 7; fi111++ {
                var v112 string = ""
                _ = len(v2)
                v46 += (v46 | (v46 | (-(v46))))
                _ = fi111
                _ = v112
            }
            if (!(!false)) {
                for fi113 := 0; fi113 < 16; fi113++ {
                    v46 -= v46
                    fmt.Println(v79, v3)
                    v72 = float64(1.5)
                    _ = fi113
                }
            }
            _ = fi109
            _ = v110
        }
    }
    var v114 uint8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v72:", v72)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v72
    _ = v77
    _ = v78
    _ = v79
}
