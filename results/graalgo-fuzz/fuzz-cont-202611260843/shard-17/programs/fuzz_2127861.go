package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 uint }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64; F1 int32; F2 int16 }
func genFn0() uint32 {
    v4 := []int{960, 737, 68, 0}
    v5 := GenS1{F0: 989}
    var v6 int8 = -87
    switch ("a" + "foo bar") {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "a":
        v7 := GenS1{F0: 1}
        _ = v7
    case "":
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS1{F0: 117}
            _ = fi8
            _ = v9
        }
    }
    v6 = v6
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint8, p1 int, p2 bool) uint32 {
    p1 += (p1 + (p1 | (^(p1))))
    fmt.Println(p0, p0, p0)
    v10 := GenS1{F0: 1}
    if (0 < 42) {
        p1 = p1
        for fi11 := 0; fi11 < 7; fi11++ {
            switch p1 {
            case 2147483647:
                if p2 {
                    v12 := GenS1{F0: 561}
                    var v13 int16 = -42
                    v14 := GenS1{F0: 42}
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS2{F0: 42, F1: 4294967295, F2: ((9223372036854775807 > -9223372036854775808) && (p2 || false))}
                    _ = v15
                }
            default:
                v16 := GenS3{F0: 143, F1: 547, F2: (-9 != -104)}
                _ = v16
            }
            _ = fi11
        }
        switch (p1 | p1) {
        case 42:
            if (((!p2) || ((p1 ^ -409) > p1)) || p2) {
                if true {
                    p2 = (p2 && (!false))
                    fmt.Println(p2, p0)
                    v17 := []int{}
                    _ = v17
                }
                v18 := []int{}
                p1 = p1
                _ = v18
            }
        case -9223372036854775808:
            fmt.Println(p1, p2, p0)
        case 9223372036854775807:
            for fi19 := 0; fi19 < 18; fi19++ {
                p1 -= (fi19 << ((2) & 63))
                var v20 uint32 = 838
                _ = fi19
                _ = v20
            }
        default:
            var v21 float32 = 1.5
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 658
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    v3 := []string{"", "a"}
    if (255 != 127) {
        for fi22 := 0; fi22 < 13; fi22++ {
            fmt.Println(fi22, fi22)
            if (204 < -2147483648) {
                v2 += v2
                for fi23 := 0; fi23 < 6; fi23++ {
                    v2 -= v2
                    if false {
                        fmt.Println(v2, v2, v2)
                        fmt.Println(fi22)
                        v2 -= 2
                        if (1 >= 100) {
                            v24 := v1[42]
                            v25 := GenS1{F0: 1}
                            if true {
                                var v26 uint64 = 29
                                var v27 float64 = 0.0
                                _ = v26
                                _ = v27
                            } else {
                                if false {
                                    if ((v2 + 0) != 1) {
                                        v2 -= v2
                                    }
                                    v28 := GenS2{F0: (v2 + 32767), F1: 0, F2: (!false)}
                                    _ = v28
                                } else {
                                    v2 += 255
                                    if (-65 == -628) {
                                        var v29 uint = 255
                                        v29 += (v29 << (((v29 ^ 1)) & 63))
                                        fmt.Println(fi23)
                                        _ = v29
                                    }
                                }
                                fmt.Println(v24, v2, v2)
                                v30 := GenS3{F0: uint64(v2), F1: 42, F2: true}
                                _ = v30
                            }
                            var v31 int16 = 144
                            _ = v24
                            _ = v25
                            _ = v31
                        }
                    }
                    _ = fi23
                }
            }
            v2 = v2
            _ = fi22
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            if true {
                v33 := GenS4{F0: 255, F1: 461, F2: 42}
                switch 42 {
                case 977:
                    var v34 int64 = 1
                    _ = v34
                case -827:
                    v2 += v2
                }
                _ = v33
            } else {
                v3 = append(v3, "foo bar")
                if false {
                    v2 = 267
                    v35 := GenS1{F0: 1}
                    _ = v35
                } else {
                    v2 -= (v2 | v2)
                }
            }
            fmt.Println(fi32)
            _ = fi32
        }
        switch "\n\t" {
        case "a":
            fmt.Println(v2)
        default:
            var v36 float32 = (0.1 * float32(((3.14159 + float64((3.14159 - float32((1e10 / (1e10 / float64((0.1 / ((float32(-2.25) + 1.5) * (1e10 * (0.0 + (((1.5 * (1.5 + 0.0)) + 0.0) + 0.0)))))))))))) * 1.5)))
            _ = v36
        }
    } else {
        v2 += v2
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS5{F0: "foo bar", F1: 4294967295}
            _ = fi37
            _ = v38
        }
    }
    var v39 int64 = 42
    var v40 int = -870
    delete(v1, 798)
    switch ("\n\t" + (((((("foo bar" + "hello") + (("hello" + "hello") + "")) + "\n\t") + "foo bar") + ("a" + "hello")) + ((("hello" + "a") + "\n\t") + ""))) {
    case "hello":
        fmt.Println(v2)
    case "foo bar":
        v2 = (v2 / ((42) | 1))
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        v42 := GenS6{F0: 4294967295}
        var v43 uint64 = 42
        var v44 uint = 65535
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    if true {
        for fi45 := 0; fi45 < 15; fi45++ {
            v39 = 0
            _ = fi45
        }
        if false {
            var v46 uint16 = v2
            var v47 int32 = 813
            var v48 int16 = 1
            _ = v46
            _ = v47
            _ = v48
        } else {
            var v49 uint32 = 1
            for fi50 := 0; fi50 < 9; fi50++ {
                for fi51 := 0; fi51 < 8; fi51++ {
                    fmt.Println(fi51)
                    v52 := GenS7{F0: -511, F1: 1}
                    _ = len(v3)
                    _ = fi51
                    _ = v52
                }
                _ = fi50
            }
            var v53 int64 = (v39 + (v39 / ((v39) | 1)))
            _ = v49
            _ = v53
        }
        if (-2.25 != 0.0) {
            var v54 uint64 = 42
            _ = v54
        }
    } else {
        if (((0 >= 343) && ((v39 - v39) > (v39 % (((v39 % ((-604) | 1))) | 1)))) && (-2147483648 != 646)) {
            if true {
                if (false && ((2147483647 == -929) && (((3.14159 / 3.14159) == 1e10) || (false || (v2 > (v2 * 1)))))) {
                    var v55 uint32 = genFn1(2, v40, (2147483647 >= -2147483648))
                    _ = v55
                } else {
                    v56 := GenS1{F0: 0}
                    var v57 int = (v40 - (v40 + (v40 & v40)))
                    _ = v56
                    _ = v57
                }
            }
        } else {
            var v58 uint32 = 2147483648
            v3 = append(v3, (("" + "foo bar") + ("a" + ("hello" + ("\n\t" + (("hello" + "hello") + "a"))))))
            v59 := GenS8{F0: (true && (42 <= v2))}
            _ = v58
            _ = v59
        }
        fmt.Println(v39, v40)
    }
    v40 = v40
    if false {
        v60 := GenS3{F0: 18446744073709551615, F1: (v39 % ((v39) | 1)), F2: (!(0 < 18446744073709551615))}
        _ = v60
    }
    var v61 int32 = -118
    v40 -= v40
    v39 = v39
    switch ("a" + (("foo bar" + "hello") + "hello")) {
    case "\n\t":
        if (float32((0.1 / 1e10)) > float32(0.1)) {
            var v62 int8 = 42
            _ = v62
        }
    case "foo bar":
        fmt.Println(v2, v39, v2)
    case "":
        v2 -= (v2 ^ 42)
    }
    switch 42 {
    case 32767:
        if true {
            if true {
                if (2 > 65535) {
                    var v63 uint8 = 255
                    var v64 bool = true
                    _ = v63
                    _ = v64
                } else {
                    fmt.Println(v39, v61, v39)
                    var v65 int64 = v39
                    _ = v65
                }
                if (!("a" > ("\n\t" + (((((("\n\t" + "foo bar") + "") + "\n\t") + "foo bar") + "") + ("\n\t" + "a"))))) {
                    var v66 int32 = v61
                    v3 = append(v3, ("a" + ("a" + "a")))
                    _ = v66
                }
                for fi67 := 0; fi67 < 3; fi67++ {
                    fmt.Println(v39, v39)
                    var v68 uint = 1
                    for fi69 := 0; fi69 < 7; fi69++ {
                        var v70 int = v40
                        v61 = int32(v70)
                        var v71 uint32 = uint32(v68)
                        if false {
                            v72 := GenS7{F0: -1, F1: 47}
                            _ = v72
                        }
                        var v73 int16 = 42
                        var v74 uint = v68
                        _ = fi69
                        _ = v70
                        _ = v71
                        _ = v73
                        _ = v74
                    }
                    _ = fi67
                    _ = v68
                }
            } else {
                var v75 uint64 = 42
                if true {
                    v61 += v61
                    var v76 uint8 = 2
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    _ = v76
                } else {
                    var v77 uint16 = 32767
                    if ((v40 != (v40 ^ v40)) || (false && true)) {
                        v78 := GenS3{F0: (^(v75)), F1: v39, F2: (true || true)}
                        var v79 int32 = 0
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v78
                        _ = v79
                    }
                    v39 = v39
                    _ = v77
                }
                delete(v1, 0)
                _ = v75
            }
        }
    default:
        if (!true) {
            fmt.Println(v2, v61)
            fmt.Println(v40, v39)
            switch (v40 * int(v2)) {
            case 128:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            v61 -= (v61 - (v61 | v61))
            v80 := GenS9{F0: 1, F1: 255, F2: 600}
            _ = v80
        }
    }
    var v81 float64 = 1.5
    switch v40 {
    case -1:
        for fi82 := 0; fi82 < 16; fi82++ {
            for fi83 := 0; fi83 < 15; fi83++ {
                var v84 uint8 = 42
                _ = fi83
                _ = v84
            }
            _ = fi82
        }
    case -928:
        if ("\n\t" != "\n\t") {
            v85 := GenS10{F0: -32768}
            for fi86 := 0; fi86 < 5; fi86++ {
                for fi87 := 0; fi87 < 6; fi87++ {
                    var v88 float32 = 0.0
                    var v89 uint32 = 0
                    _ = fi87
                    _ = v88
                    _ = v89
                }
                _ = fi86
            }
            var v90 int = v40
            _ = v85
            _ = v90
        } else {
            var v91 float32 = 1.5
            if (!(!false)) {
                switch ((("a" + ("\n\t" + ("" + ("hello" + ("foo bar" + ("hello" + "hello")))))) + ("a" + "")) + "a") {
                case "hello":
                    v40 -= int(v2)
                case "foo bar":
                    var v92 int = 32767
                    _ = v92
                default:
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        if true {
                            var v93 uint32 = 2
                            v94 := GenS11{F0: (v39 << ((0) & 63)), F1: (v61 + v61), F2: 0}
                            _ = v93
                            _ = v94
                        } else {
                            v91 -= float32(0.1)
                            var v95 uint32 = 472
                            _ = v95
                        }
                    case "\n\t":
                        for fi96 := 0; fi96 < 10; fi96++ {
                            v97 := GenS10{F0: -1}
                            _ = fi96
                            _ = v97
                        }
                    }
                }
            } else {
                switch (-(v40)) {
                case 2147483647:
                    if false {
                        v98 := GenS3{F0: 18446744073709551615, F1: (v39 * 848), F2: false}
                        var v99 int16 = 410
                        for fi100 := 0; fi100 < 15; fi100++ {
                            fmt.Println(v2, v61, fi100)
                            v2 = (v2 / (((^(v2))) | 1))
                            _ = fi100
                        }
                        _ = v98
                        _ = v99
                    }
                case -459:
                    for fi101 := 0; fi101 < 20; fi101++ {
                        if ((v2 + v2) == v2) {
                            fmt.Println(v40)
                            var v102 uint8 = 12
                            switch "\n\t" {
                            case "\n\t":
                                for fi103 := 0; fi103 < 4; fi103++ {
                                    switch "\n\t" {
                                    case "hello":
                                        switch "\n\t" {
                                        case "foo bar":
                                            v104 := GenS1{F0: 18446744073709551615}
                                            _ = v104
                                        case "hello":
                                            var v105 float64 = v81
                                            _ = v105
                                        default:
                                            v106 := GenS4{F0: (v102 / (((^(v102))) | 1)), F1: 9223372036854775807, F2: 2}
                                            _ = v106
                                        }
                                    case "foo bar":
                                        var v107 uint64 = 2
                                        _ = v107
                                    default:
                                        var v108 uint64 = 701
                                        _ = v108
                                    }
                                    var v109 uint16 = v2
                                    v1[601] = (("a" + (("hello" + ("hello" + ("foo bar" + "foo bar"))) + "\n\t")) + "\n\t")
                                    _ = fi103
                                    _ = v109
                                }
                            case "a":
                                v110 := GenS9{F0: v102, F1: (^(v102)), F2: v39}
                                _ = v110
                            default:
                                v111 := GenS2{F0: 32767, F1: 804, F2: true}
                                _ = v111
                            }
                            _ = v102
                        } else {
                            v112 := GenS2{F0: v2, F1: uint64(v2), F2: (((v2 - v2) >= (v2 % ((2) | 1))) && (!(255 > 819)))}
                            var v113 string = ""
                            _ = v112
                            _ = v113
                        }
                        _ = len(v1)
                        switch "\n\t" {
                        case "foo bar":
                            var v114 string = ("" + "\n\t")
                            _ = v114
                        case "a":
                            v115 := GenS1{F0: 0}
                            _ = v115
                        default:
                            if (((444 >= 2) && false) || (503 == 2)) {
                                var v116 uint32 = 189
                                _ = v116
                            } else {
                                var v117 uint = uint(v2)
                                if ((42 <= 1) || (2147483647 >= (v61 | int32(fi101)))) {
                                    var v118 float32 = (v91 - (float32(((1.5 / v81) - v81)) / ((v91 * (3.14159 * v91)) + (v91 * (float32(1.5) - 0.0)))))
                                    _ = v118
                                } else {
                                    v40 -= 127
                                    if (fi101 <= fi101) {
                                        var v119 uint64 = 519
                                        var v120 int32 = (^(v61))
                                        v39 = v39
                                        _ = v119
                                        _ = v120
                                    } else {
                                        var v121 string = "hello"
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = v121
                                        }
                                        _ = v121
                                    }
                                    var v122 string = "\n\t"
                                    _ = v122
                                }
                                _ = v117
                            }
                        }
                        _ = fi101
                    }
                }
            }
            v1[65535] = "hello"
            _ = v91
        }
    case 9223372036854775807:
        var v123 int16 = -442
        _ = v123
    default:
        var v124 float32 = 0.1
        _ = v124
    }
    for fi125 := 0; fi125 < 10; fi125++ {
        var v126 uint32 = 2147483647
        _ = fi125
        _ = v126
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v61:", v61)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v61
    _ = v81
}
