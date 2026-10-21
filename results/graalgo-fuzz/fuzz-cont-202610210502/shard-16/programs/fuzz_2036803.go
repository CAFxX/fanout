package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32; F1 bool }
func genFn0(p0 uint64, p1 string) uint32 {
    switch p1 {
    case "foo bar":
        var v7 float32 = (1.5 / 0.1)
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int16 = -32768
            fmt.Println(v9, p1)
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            var v11 uint16 = 1
            v12 := GenS1{F0: "", F1: ""}
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 string = p1
                if false {
                    v15 := GenS2{F0: 593}
                    v16 := []int{fi10, fi13, (fi13 / ((4294967295) | 1))}
                    var v17 int64 = 9223372036854775807
                    _ = v15
                    _ = v16
                    _ = v17
                }
                if true {
                    v18 := []int{2147483647, -193, fi10, (fi13 / ((-9223372036854775808) | 1))}
                    if len(v18) > 0 {
                        v18[0] = (fi10 % ((fi13) | 1))
                    }
                    v19 := GenS2{F0: -60}
                    _ = v18
                    _ = v19
                }
                _ = fi13
                _ = v14
            }
            v20 := []int{127}
            var v21 uint8 = 128
            _ = v20
            _ = v21
        }
        _ = fi10
    }
    if (410 != 42) {
        fmt.Println(p1)
        var v22 uint64 = (p0 << ((65535) & 63))
        var v23 int = 0
        _ = v22
        _ = v23
    } else {
        if (-40 < 42) {
            if (!false) {
                var v24 float64 = (0.1 + 0.1)
                if true {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        if true {
                            switch p1 {
                            case "foo bar":
                                p0 -= p0
                            case "a":
                                v26 := []int{}
                                _ = v26
                            case "":
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    p1 = p1
                                    p0 += p0
                                    v24 = (1e10 / (((-2.25 / v24) - ((3.14159 + 0.0) + 0.0)) * ((v24 / 0.1) / 1.0)))
                                    p0 = p0
                                    _ = fi27
                                }
                            }
                        }
                        _ = fi25
                    }
                } else {
                    v28 := GenS2{F0: 32767}
                    p1 = "foo bar"
                    _ = v28
                }
                switch p1 {
                case "":
                    v24 = v24
                }
                _ = v24
            } else {
                fmt.Println(p0, p0)
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(fi29)
                    for fi30 := 0; fi30 < 1; fi30++ {
                        v31 := GenS1{F0: p1, F1: ""}
                        v32 := []int{128}
                        var v33 uint16 = 507
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    var v34 uint8 = 255
                    _ = fi29
                    _ = v34
                }
                for fi35 := 0; fi35 < 4; fi35++ {
                    if (130 != 128) {
                        v36 := GenS1{F0: "a", F1: (p1 + "a")}
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 7; fi37++ {
                        var v38 uint = 0
                        _ = fi37
                        _ = v38
                    }
                    v39 := GenS1{F0: (("" + p1) + ("foo bar" + (p1 + p1))), F1: (("" + (p1 + p1)) + p1)}
                    _ = fi35
                    _ = v39
                }
            }
            p0 += (^(p0))
        }
        if true {
            for fi40 := 0; fi40 < 5; fi40++ {
                v41 := GenS2{F0: 0}
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 8; fi42++ {
                switch "\n\t" {
                case "a":
                    var v43 int32 = 484
                    _ = v43
                }
                _ = fi42
            }
        }
        v44 := []int{127, -411, 2147483647, 4294967295}
        _ = v44
    }
    p0 += (^(p0))
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int64) uint64 {
    if true {
        v45 := []int{}
        p0 = p0
        v46 := GenS1{F0: "hello", F1: "hello"}
        _ = v45
        _ = v46
    } else {
        fmt.Println(p0)
        var v47 uint32 = 42
        _ = v47
    }
    v48 := GenS3{F0: ((-2.25 * ((float32(1.5) * float32(3.14159)) - 0.0)) < 1e10), F1: (false || true), F2: 2}
    fmt.Println(p0)
    v49 := GenS3{F0: (true && true), F1: (405 == 42), F2: 58}
    _ = p0
    _ = v48
    _ = v49
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]bool)
    var v3 string = "a"
    v4 := []int{}
    v5 := make(map[uint32]string)
    var v6 uint16 = 2
    var v50 float64 = ((0.0 * -2.25) / -2.25)
    if ((-110 <= 100) || (v6 >= v6)) {
        var v51 uint8 = (v1 % (((v1 + v1)) | 1))
        switch "\n\t" {
        case "a":
            var v52 int32 = 2147483647
            _ = v52
        case "foo bar":
            fmt.Println(v6, v50, v50)
        default:
            for fi53 := 0; fi53 < 18; fi53++ {
                v1 = v1
                fmt.Println(v6, v50)
                _ = fi53
            }
        }
        _ = v51
    } else {
        fmt.Println(v50, v6)
    }
    v6 = 1
    v54 := GenS2{F0: 1000}
    switch 128 {
    case -2147483648:
        for fi55 := 0; fi55 < 3; fi55++ {
            v6 = v6
            v1 = uint8(v6)
            switch fi55 {
            case -649:
                for fi56 := 0; fi56 < 18; fi56++ {
                    if len(v4) > 0 {
                        v4[0] = fi55
                    }
                    _ = fi56
                }
            case 128:
                v1 += (v1 << ((1) & 63))
            default:
                for fi57 := 0; fi57 < 9; fi57++ {
                    fmt.Println(v50)
                    v50 = 0.1
                    if (!false) {
                        v3 = ((v3 + "foo bar") + "hello")
                        _ = len(v4)
                    }
                    _ = fi57
                }
            }
            _ = fi55
        }
    case 128:
        if false {
            var v58 bool = true
            _ = v58
        } else {
            v1 = (v1 + (^(v1)))
            var v59 int16 = -290
            if false {
                fmt.Println(v6)
                v4 = append(v4, -133)
            }
            _ = v59
        }
    }
    v60 := GenS3{F0: ((v6 / (((v6 ^ (v6 % ((v6) | 1)))) | 1)) != 923), F1: true, F2: (^(v1))}
    v6 = uint16(v1)
    if true {
        v61 := GenS3{F0: ((((v3 + v3) + v3) + (("hello" + v3) + (((v3 + v3) + ("" + "\n\t")) + v3))) >= "hello"), F1: true, F2: (v1 / ((v1) | 1))}
        for fi62 := 0; fi62 < 10; fi62++ {
            if false {
                var v63 bool = true
                if true {
                    fmt.Println(v3)
                }
                for fi64 := 0; fi64 < 9; fi64++ {
                    var v65 int32 = -553
                    if v63 {
                        fmt.Println(v63, v6, v6)
                        v65 -= v65
                    } else {
                        delete(v5, 0)
                        v66 := v2["\n\t"]
                        v50 = v50
                        _ = v66
                    }
                    _ = fi64
                    _ = v65
                }
                _ = v63
            } else {
                if (v1 == 1) {
                    fmt.Println(fi62, fi62, v1)
                    var v67 bool = (true || ((fi62 * fi62) == (fi62 / ((fi62) | 1))))
                    fmt.Println(v3, v67, v67)
                    _ = v67
                } else {
                    if (127 > -1) {
                        v2["foo bar"] = (int64(v6) >= 1)
                        if (true || false) {
                            v3 = "a"
                            var v68 uint = 0
                            if (4294967295 <= 1) {
                                var v69 uint16 = (v6 / ((v6) | 1))
                                _ = v69
                            } else {
                                v70 := GenS4{F0: -554}
                                switch (fi62 * fi62) {
                                case 0:
                                    switch (fi62 - fi62) {
                                    case 128:
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = fi62
                                            }
                                            v3 = (v3 + v3)
                                            for fi71 := 0; fi71 < 8; fi71++ {
                                                v72 := GenS4{F0: -32768}
                                                var v73 bool = (v6 < (v6 >> ((370) & 63)))
                                                _ = fi71
                                                _ = v72
                                                _ = v73
                                            }
                                        } else {
                                            for fi74 := 0; fi74 < 2; fi74++ {
                                                if false {
                                                    var v75 string = (((v3 + "\n\t") + v3) + v3)
                                                    _ = v75
                                                }
                                                var v76 int64 = -113
                                                switch (fi62 % ((fi62) | 1)) {
                                                case 32767:
                                                    v1 = v1
                                                case -986:
                                                    v76 -= (v76 / ((9223372036854775807) | 1))
                                                }
                                                if ("foo bar" > v3) {
                                                    switch v3 {
                                                    case "foo bar":
                                                        if true {
                                                            fmt.Println(fi74)
                                                            var v77 uint32 = 1
                                                            fmt.Println(v3, v3, v1)
                                                            _ = v77
                                                        } else {
                                                            v78 := v5[uint32(v1)]
                                                            if true {
                                                                v79 := GenS4{F0: -431}
                                                                var v80 int64 = (v76 << (((v68 ^ v68)) & 63))
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                var v81 int = fi62
                                                                v5[247] = v3
                                                                _ = v81
                                                            }
                                                            fmt.Println(v3)
                                                            _ = v78
                                                        }
                                                    }
                                                }
                                                var v82 int32 = 0
                                                _ = fi74
                                                _ = v76
                                                _ = v82
                                            }
                                            v50 = float64(0.0)
                                        }
                                    case -9223372036854775808:
                                        _ = len(v2)
                                    default:
                                        v83 := GenS5{F0: fi62}
                                        _ = v83
                                    }
                                case 4294967295:
                                    v3 = v3
                                }
                                _ = v70
                            }
                            _ = v68
                        } else {
                            if true {
                                v6 = 32767
                                for fi84 := 0; fi84 < 18; fi84++ {
                                    switch fi84 {
                                    case -1:
                                        v1 -= uint8(v6)
                                    case 1:
                                        v85 := GenS6{F0: 4294967295}
                                        _ = v85
                                    case 127:
                                        v86 := GenS7{F0: 1000}
                                        _ = v86
                                    default:
                                        _ = len(v2)
                                    }
                                    switch (fi62 / (((fi62 * (fi84 >> ((65535) & 63)))) | 1)) {
                                    case 127:
                                        v87 := GenS8{F0: -32768, F1: 1}
                                        _ = v87
                                    case 331:
                                        var v88 uint8 = (^(v1))
                                        _ = v88
                                    default:
                                        var v89 uint64 = 656
                                        var v90 uint32 = 4294967295
                                        _ = v89
                                        _ = v90
                                    }
                                    v3 = ""
                                    _ = fi84
                                }
                                v6 += (v6 | 65535)
                            }
                            v2[v3] = (((!false) && (false && true)) && true)
                            var v91 int64 = -9223372036854775808
                            _ = v91
                        }
                        fmt.Println(v1)
                    } else {
                        if (!(false || ((^(v6)) > 32767))) {
                            v92 := GenS9{F0: fi62}
                            v50 = v50
                            if (42 > -1) {
                                v93 := GenS10{F0: 170, F1: (((v6 % ((65535) | 1)) > 2) && (float32(1.5) >= ((-2.25 - 3.14159) - -2.25)))}
                                if (v50 == v50) {
                                    switch v3 {
                                    case "hello":
                                        v1 -= 0
                                    case "foo bar":
                                        var v94 uint32 = 42
                                        _ = v94
                                    case "a":
                                        for fi95 := 0; fi95 < 6; fi95++ {
                                            var v96 uint64 = 848
                                            var v97 bool = true
                                            var v98 int16 = -32768
                                            _ = fi95
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                    default:
                                        v99 := GenS4{F0: -32768}
                                        _ = v99
                                    }
                                    var v100 uint64 = 4294967295
                                    for fi101 := 0; fi101 < 7; fi101++ {
                                        switch -9223372036854775808 {
                                        case 0:
                                            _ = len(v2)
                                        case 128:
                                            switch fi101 {
                                            case 42:
                                                switch v3 {
                                                case "a":
                                                    v102 := GenS3{F0: (true || true), F1: false, F2: (v1 / (((v1 * 42)) | 1))}
                                                    _ = v102
                                                }
                                            case 32768:
                                                v103 := GenS4{F0: -1}
                                                _ = v103
                                            case 127:
                                                var v104 uint16 = (v6 | 255)
                                                _ = v104
                                            }
                                        }
                                        v1 -= (v1 & v1)
                                        v50 += (1.5 * v50)
                                        var v105 bool = false
                                        _ = fi101
                                        _ = v105
                                    }
                                    _ = v100
                                } else {
                                    var v106 uint64 = 4294967295
                                    _ = v106
                                }
                                var v107 int16 = 32767
                                _ = v107
                                _ = v93
                            } else {
                                var v108 int32 = 1
                                v3 = ((((v3 + (v3 + (v3 + v3))) + v3) + "foo bar") + "")
                                if (18446744073709551615 < 711) {
                                    v109 := GenS8{F0: 32767, F1: -2147483648}
                                    var v110 uint32 = 1
                                    _ = v109
                                    _ = v110
                                }
                                _ = v108
                            }
                            _ = v92
                        } else {
                            _ = len(v4)
                            v111 := GenS10{F0: 42, F1: false}
                            if true {
                                var v112 int32 = 2147483647
                                fmt.Println(v112, v3)
                                _ = v112
                            } else {
                                v1 += 42
                                v113 := GenS2{F0: 1}
                                for fi114 := 0; fi114 < 4; fi114++ {
                                    if false {
                                        v115 := GenS8{F0: 42, F1: 2147483647}
                                        _ = v115
                                    }
                                    _ = fi114
                                }
                                _ = v113
                            }
                            _ = v111
                        }
                        v116 := GenS10{F0: 0, F1: true}
                        _ = v116
                    }
                }
            }
            _ = fi62
        }
        for fi117 := 0; fi117 < 14; fi117++ {
            fmt.Println(v6)
            _ = fi117
        }
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v54
    _ = v6
    _ = v60
}
