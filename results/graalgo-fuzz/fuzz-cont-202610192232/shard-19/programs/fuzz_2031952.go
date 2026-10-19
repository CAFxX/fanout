package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int32; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 bool; F2 uint32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 int8, p1 int, p2 int64) bool {
    v6 := GenS1{F0: 172}
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int64 = -864
        v9 := []int{}
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int64 = int64(p0)
            p0 = (p0 - p0)
            v12 := GenS1{F0: 42}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v13 := []int{432, p1, p1}
    v13 = append(v13, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    return true
}
func genFn1(p0 int, p1 int32) uint {
    p0 = 601
    for fi14 := 0; fi14 < 15; fi14++ {
        p1 = p1
        for fi15 := 0; fi15 < 9; fi15++ {
            fmt.Println(fi14, p1, fi14)
            var v16 int16 = 1
            _ = fi15
            _ = v16
        }
        switch fi14 {
        case 504:
            switch (((("" + ("foo bar" + "a")) + "a") + "a") + "") {
            case "foo bar":
                var v17 string = ("foo bar" + "a")
                _ = v17
            case "hello":
                for fi18 := 0; fi18 < 2; fi18++ {
                    fmt.Println(p1, fi14)
                    fmt.Println(fi14)
                    v19 := []int{(fi14 / ((int(p1)) | 1)), (^(fi18)), (^(fi18)), fi14}
                    _ = fi18
                    _ = v19
                }
            case "\n\t":
                v20 := []int{p0, p0, -1}
                _ = v20
            default:
                switch (p0 / ((-2147483648) | 1)) {
                case 42:
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v22 := GenS2{F0: 1}
                        _ = fi21
                        _ = v22
                    }
                case 128:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        p0 += (^(fi23))
                        _ = fi23
                    }
                case 32768:
                    var v24 int64 = 0
                    _ = v24
                }
            }
        case -748:
            switch fi14 {
            case 1:
                if (!true) {
                    if (2 >= 155) {
                        var v25 uint8 = 0
                        var v26 uint16 = 358
                        _ = v25
                        _ = v26
                    } else {
                        v27 := []int{(-(fi14)), (fi14 >> ((480) & 63)), (fi14 / ((p0) | 1)), 32767}
                        if (!false) {
                            fmt.Println(p1, fi14, fi14)
                            if true {
                                if true {
                                    var v28 float32 = 3.14159
                                    if true {
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                        switch (fi14 | -477) {
                                        case -2147483648:
                                            if len(v27) > 0 {
                                                v27[0] = (p0 * (p0 * 32767))
                                            }
                                        default:
                                            switch fi14 {
                                            case 0:
                                                var v29 int16 = 0
                                                _ = v29
                                            default:
                                                switch (p0 | (fi14 / ((fi14) | 1))) {
                                                case 127:
                                                    p0 -= fi14
                                                case 2147483647:
                                                    var v30 uint16 = 1
                                                    _ = v30
                                                case -445:
                                                    for fi31 := 0; fi31 < 2; fi31++ {
                                                        var v32 int32 = p1
                                                        v32 = 0
                                                        _ = fi31
                                                        _ = v32
                                                    }
                                                default:
                                                    v33 := GenS3{F0: 179, F1: 714, F2: (!(127 != 100))}
                                                    _ = v33
                                                }
                                            }
                                        }
                                    } else {
                                        v34 := GenS1{F0: 42}
                                        fmt.Println(p1, fi14)
                                        _ = v34
                                    }
                                    _ = v28
                                }
                                p1 -= p1
                                var v35 int32 = p1
                                _ = v35
                            }
                            switch -2147483648 {
                            case 4294967295:
                                v27 = append(v27, -326)
                            default:
                                var v36 uint16 = 255
                                _ = v36
                            }
                        } else {
                            for fi37 := 0; fi37 < 8; fi37++ {
                                var v38 int16 = 32767
                                p0 -= p0
                                p1 += (p1 & 42)
                                _ = fi37
                                _ = v38
                            }
                            if (0.1 == ((1e10 * 0.0) * ((0.1 + (((3.14159 - 0.0) + -2.25) + 1e10)) * (0.0 - ((3.14159 + 1e10) / 1.5))))) {
                                var v39 uint8 = 154
                                if true {
                                    switch (fi14 + 42) {
                                    case -695:
                                        if (18446744073709551615 > 950) {
                                            for fi40 := 0; fi40 < 11; fi40++ {
                                                if false {
                                                    fmt.Println(fi14, p1)
                                                    var v41 uint16 = 32767
                                                    _ = v41
                                                } else {
                                                    if true {
                                                        fmt.Println(p0)
                                                    }
                                                }
                                                v42 := GenS4{F0: 9223372036854775807, F1: 65535, F2: 1}
                                                v39 = (v39 >> ((65535) & 63))
                                                _ = fi40
                                                _ = v42
                                            }
                                            var v43 uint64 = 945
                                            _ = v43
                                        } else {
                                            switch p0 {
                                            case 989:
                                                var v44 int = 280
                                                _ = v44
                                            case 488:
                                                var v45 int16 = -32768
                                                _ = v45
                                            }
                                        }
                                    default:
                                        if ((false && false) && ((-(p1)) >= p1)) {
                                            fmt.Println(p0, fi14, p1)
                                            var v46 uint32 = 84
                                            _ = v46
                                        }
                                    }
                                    for fi47 := 0; fi47 < 9; fi47++ {
                                        var v48 int8 = 42
                                        v49 := GenS5{F0: 0, F1: 0, F2: 531}
                                        v50 := GenS6{F0: 0, F1: 2}
                                        fmt.Println(p0, fi14)
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                } else {
                                    var v51 uint64 = 42
                                    fmt.Println(v39, p1)
                                    var v52 int32 = p1
                                    _ = v51
                                    _ = v52
                                }
                                _ = v39
                            } else {
                                for fi53 := 0; fi53 < 10; fi53++ {
                                    v54 := GenS7{F0: 536, F1: 0, F2: 18446744073709551615}
                                    _ = fi53
                                    _ = v54
                                }
                                var v55 uint16 = 42
                                _ = v55
                            }
                            switch "" {
                            case "":
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            case "hello":
                                for fi56 := 0; fi56 < 13; fi56++ {
                                    v57 := GenS4{F0: 1, F1: 825, F2: 2147483647}
                                    v58 := GenS8{F0: (p0 + (p0 * (fi56 >> ((255) & 63))))}
                                    var v59 bool = (3.14159 != 1e10)
                                    if len(v27) > 0 {
                                        v27[0] = (fi56 - fi14)
                                    }
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                            }
                        }
                        _ = v27
                    }
                }
            default:
                v60 := GenS3{F0: 1, F1: 901, F2: ((-2.25 - 0.1) > 3.14159)}
                _ = v60
            }
        case -641:
            for fi61 := 0; fi61 < 7; fi61++ {
                v62 := []int{p0, p0, p0}
                var v63 uint16 = 515
                p1 += p1
                _ = fi61
                _ = v62
                _ = v63
            }
        default:
            var v64 int32 = p1
            _ = v64
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 int8 = 100
    var v2 int16 = 1
    v3 := make(map[uint32]bool)
    v4 := []int{-1, -9223372036854775808}
    v5 := []string{("hello" + "foo bar"), "hello"}
    for fi65 := 0; fi65 < 6; fi65++ {
        for fi66 := 0; fi66 < 3; fi66++ {
            v67 := GenS9{F0: (fi65 ^ fi65), F1: true, F2: 1}
            _ = fi66
            _ = v67
        }
        v1 = v1
        _ = fi65
    }
    var v68 int8 = (v1 & v1)
    v2 = v2
    if true {
        var v69 int64 = 0
        _ = v69
    }
    fmt.Println(v68, v68, v1)
    if true {
        v2 -= (v2 / ((v2) | 1))
        if (-293 == -1) {
            var v70 uint16 = 923
            var v71 uint8 = 198
            _ = v70
            _ = v71
        } else {
            v68 = v1
            v2 = -32768
            v2 = v2
        }
        var v72 uint8 = 42
        _ = v72
    } else {
        if (v2 != v2) {
            var v73 bool = true
            for fi74 := 0; fi74 < 7; fi74++ {
                var v75 int8 = v1
                _ = fi74
                _ = v75
            }
            v2 = v2
            _ = v73
        }
        for fi76 := 0; fi76 < 1; fi76++ {
            v77 := v3[1]
            _ = fi76
            _ = v77
        }
    }
    if (0 == 284) {
        if (-2147483648 <= 0) {
            for fi78 := 0; fi78 < 10; fi78++ {
                _ = len(v3)
                if (0.0 < (0.0 - 3.14159)) {
                    var v79 uint64 = 2
                    var v80 bool = true
                    if true {
                        fmt.Println(v68, v80)
                    } else {
                        var v81 uint = genFn1(fi78, -2147483648)
                        v79 = 244
                        _ = v81
                    }
                    _ = v79
                    _ = v80
                } else {
                    if (((!true) && true) && (false && true)) {
                        for fi82 := 0; fi82 < 11; fi82++ {
                            var v83 int = 4294967295
                            _ = fi82
                            _ = v83
                        }
                    }
                    var v84 bool = (!true)
                    var v85 uint16 = 32767
                    _ = v84
                    _ = v85
                }
                if (true || false) {
                    v68 = v1
                    _ = len(v3)
                } else {
                    v1 = -21
                }
                _ = fi78
            }
        }
        v86 := GenS8{F0: 0}
        fmt.Println(v2, v2)
        _ = v86
    }
    if true {
        var v87 int16 = (^(v2))
        _ = v87
    }
    v1 -= v68
    var v88 uint64 = 1
    for fi89 := 0; fi89 < 16; fi89++ {
        for fi90 := 0; fi90 < 16; fi90++ {
            if (42 != 0) {
                delete(v3, 249)
                v1 += 127
                fmt.Println(v68)
            }
            _ = fi90
        }
        switch 13 {
        case -115:
            var v91 int16 = v2
            _ = v91
        case 32767:
            fmt.Println(v88)
        case 230:
            fmt.Println(v1, fi89, v68)
        default:
            for fi92 := 0; fi92 < 3; fi92++ {
                if true {
                    if ("hello" == "\n\t") {
                        var v93 uint64 = (v88 & 238)
                        _ = v93
                    } else {
                        var v94 float64 = 1.5
                        v2 -= (^(v2))
                        _ = v94
                    }
                    fmt.Println(v2)
                    v88 -= (v88 ^ v88)
                    v2 = v2
                    _ = len(v3)
                } else {
                    var v95 int32 = -1
                    fmt.Println(v1)
                    _ = v95
                }
                v1 = (v68 - 1)
                switch fi89 {
                case 86:
                    v68 -= (v68 / ((int8(fi92)) | 1))
                default:
                    v96 := GenS10{F0: (fi89 + -1)}
                    if (!(false && false)) {
                        fmt.Println(v68, fi92, v1)
                        var v97 int = (fi89 & int(v68))
                        switch "\n\t" {
                        case "":
                            v88 += (v88 << ((856) & 63))
                        case "hello":
                            if (2 == 489) {
                                v68 -= -128
                                var v98 uint64 = (v88 * 1)
                                _ = v98
                            }
                        case "\n\t":
                            if true {
                                v1 = v1
                                switch "foo bar" {
                                case "\n\t":
                                    var v99 uint = 895
                                    _ = v99
                                case "hello":
                                    if ((!true) && (65535 == 958)) {
                                        var v100 bool = false
                                        _ = v100
                                    } else {
                                        fmt.Println(v2, fi89, v68)
                                        v101 := GenS5{F0: 843, F1: (v2 & v2), F2: v88}
                                        _ = v101
                                    }
                                case "":
                                    v102 := GenS5{F0: -29, F1: int16(v88), F2: 18446744073709551615}
                                    _ = v102
                                default:
                                    v103 := GenS7{F0: 4294967295, F1: 4294967295, F2: 18446744073709551615}
                                    _ = v103
                                }
                                if (!(true && false)) {
                                    v104 := GenS4{F0: 9223372036854775807, F1: 552, F2: 2147483647}
                                    _ = v104
                                } else {
                                    if true {
                                        v5 = append(v5, "")
                                    } else {
                                        fmt.Println(v97)
                                        fmt.Println(v68)
                                        v97 += fi89
                                        var v105 string = "hello"
                                        _ = v105
                                    }
                                }
                            }
                        }
                        _ = v97
                    } else {
                        if (int64(v2) != -1) {
                            var v106 int8 = 42
                            _ = v106
                        }
                        v107 := GenS11{F0: v2, F1: "", F2: 727}
                        _ = v107
                    }
                    var v108 int8 = 127
                    _ = v108
                    _ = v96
                }
                _ = fi92
            }
        }
        _ = fi89
    }
    var v109 float32 = 1e10
    v110 := GenS10{F0: 0}
    v88 += 4294967295
    switch 1 {
    case -2147483648:
        _ = len(v3)
    default:
        var v111 int32 = -1
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v68:", v68)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v88
}
