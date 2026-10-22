package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 string; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS10 struct { F0 bool; F1 int8 }
type GenS11 struct { F0 uint8; F1 int32 }
type GenS12 struct { F0 uint32; F1 int8; F2 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int8; F1 bool }
type GenS15 struct { F0 uint64; F1 int8 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 int16; F1 uint; F2 uint8 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]bool)
    var v3 int = -85
    var v4 int32 = 257
    v3 += 1
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            v4 = v4
            v6 := GenS1{F0: -30}
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := []int{v3, (-(v3)), (fi7 + v3)}
            switch "foo bar" {
            case "a":
                v3 -= fi7
            }
            v4 = (^(v4))
            _ = fi7
            _ = v8
        }
    }
    if true {
        switch (v3 ^ 0) {
        case 0:
            fmt.Println(v1, v4)
        case -250:
            v4 += v4
        }
    } else {
        v9 := []int{v3, (v3 + v3), v3}
        _ = v9
    }
    v1 = (v1 - ((1.5 + 1e10) + 1.5))
    if false {
        switch "a" {
        case "":
            fmt.Println(v4, v4, v3)
        case "\n\t":
            switch v3 {
            case 4294967295:
                for fi10 := 0; fi10 < 12; fi10++ {
                    if false {
                        v11 := GenS1{F0: -128}
                        _ = v11
                    }
                    v12 := []int{-282, 32767, v3}
                    if len(v12) > 0 {
                        v12[0] = fi10
                    }
                    _ = fi10
                    _ = v12
                }
            }
        case "a":
            v4 -= (v4 | 2147483647)
        default:
            v13 := []int{}
            _ = v13
        }
        v14 := GenS2{F0: 48}
        if false {
            _ = len(v2)
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int64 = -9223372036854775808
                if false {
                    v3 = v3
                    v3 = v3
                    switch ("hello" + ("\n\t" + "foo bar")) {
                    case "hello":
                        v17 := GenS1{F0: 127}
                        _ = v17
                    case "":
                        if true {
                            var v18 int32 = (v4 | (v4 / ((-2147483648) | 1)))
                            var v19 int16 = 1000
                            _ = v18
                            _ = v19
                        }
                    case "\n\t":
                        if false {
                            var v20 uint64 = 86
                            _ = v20
                        }
                    default:
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v22 := []int{}
                            _ = fi21
                            _ = v22
                        }
                    }
                } else {
                    delete(v2, 4294967295)
                    if (!(0 == -128)) {
                        if (true && ((((v1 / float32(float64((v1 - 0.0)))) * ((float32(0.1) / v1) + float32(float64(((1.5 / 3.14159) * 1.5))))) / v1) == v1)) {
                            v4 = -1
                            delete(v2, 0)
                        }
                        v1 -= (v1 - (v1 + (float32(((1e10 - ((1.5 - 0.1) - 0.0)) * 3.14159)) + float32(0.0))))
                    } else {
                        v23 := GenS3{F0: 345}
                        var v24 int8 = 100
                        switch "a" {
                        case "a":
                            for fi25 := 0; fi25 < 16; fi25++ {
                                var v26 int8 = (v24 & 30)
                                if (1 >= v26) {
                                    var v27 uint16 = 42
                                    v28 := GenS2{F0: 211}
                                    _ = v27
                                    _ = v28
                                } else {
                                    v4 = 42
                                }
                                _ = fi25
                                _ = v26
                            }
                        case "\n\t":
                            v16 = (v16 | (v16 * v16))
                        default:
                            for fi29 := 0; fi29 < 20; fi29++ {
                                v30 := GenS1{F0: 0}
                                fmt.Println(v4, fi15)
                                _ = fi29
                                _ = v30
                            }
                        }
                        _ = v23
                        _ = v24
                    }
                    var v31 int32 = -1
                    _ = v31
                }
                _ = fi15
                _ = v16
            }
        } else {
            if false {
                var v32 int64 = -41
                _ = v32
            } else {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v3, v1, v1)
                case "hello":
                    for fi33 := 0; fi33 < 1; fi33++ {
                        v34 := GenS4{F0: 2}
                        fmt.Println(v4, fi33)
                        _ = fi33
                        _ = v34
                    }
                case "foo bar":
                    if (42 == 42) {
                        if true {
                            var v35 string = "hello"
                            switch v35 {
                            case "a":
                                v36 := GenS1{F0: 127}
                                _ = v36
                            case "foo bar":
                                fmt.Println(v4, v4, v1)
                            case "":
                                var v37 int64 = 668
                                _ = v37
                            default:
                                v1 -= v1
                            }
                            if (549 < 0) {
                                if false {
                                    var v38 int32 = v4
                                    for fi39 := 0; fi39 < 17; fi39++ {
                                        for fi40 := 0; fi40 < 18; fi40++ {
                                            v41 := GenS5{F0: 255, F1: "", F2: fi39}
                                            _ = fi40
                                            _ = v41
                                        }
                                        v3 = fi39
                                        _ = fi39
                                    }
                                    _ = v38
                                } else {
                                    for fi42 := 0; fi42 < 8; fi42++ {
                                        for fi43 := 0; fi43 < 18; fi43++ {
                                            v1 += v1
                                            if false {
                                                var v44 uint = 65535
                                                v45 := GenS1{F0: 0}
                                                v46 := GenS5{F0: 42, F1: v35, F2: fi42}
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            } else {
                                                switch v3 {
                                                case 32768:
                                                    if (("a" + v35) <= ("\n\t" + v35)) {
                                                        switch "hello" {
                                                        case "hello":
                                                            v3 -= fi42
                                                        }
                                                    }
                                                default:
                                                    if (v1 <= v1) {
                                                        var v47 uint = 97
                                                        v48 := []int{(fi43 / (((fi43 * -419)) | 1)), fi42, v3, (v3 >> ((v47) & 63))}
                                                        var v49 float64 = 0.1
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    v2[871] = true
                                                }
                                            }
                                            if true {
                                                v50 := GenS6{F0: v4}
                                                if (!(9223372036854775807 < 0)) {
                                                    fmt.Println(v3)
                                                    fmt.Println(fi42, v1, fi43)
                                                    var v51 int8 = 127
                                                    _ = v51
                                                }
                                                switch (v35 + "\n\t") {
                                                case "\n\t":
                                                    switch v35 {
                                                    case "\n\t":
                                                        v1 -= (0.0 - (0.1 / v1))
                                                        switch fi42 {
                                                        case -1:
                                                            fmt.Println(v4, fi43)
                                                        case 2147483647:
                                                            var v52 int8 = 1
                                                            v1 = ((0.0 / 0.1) * ((v1 / 1e10) * 0.1))
                                                            _ = v52
                                                        case 4294967295:
                                                            if (true || true) {
                                                                var v53 uint32 = 0
                                                                v1 += (((0.0 + v1) / (((v1 * v1) * v1) + v1)) / (v1 - v1))
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v54 := v2[v53]
                                                                    _ = v54
                                                                }
                                                                _ = v53
                                                            }
                                                        }
                                                        var v55 int = -183
                                                        v56 := GenS7{F0: 2, F1: 4294967295, F2: 42}
                                                        _ = v55
                                                        _ = v56
                                                    case "foo bar":
                                                        _ = len(v2)
                                                    case "":
                                                        v57 := GenS1{F0: 45}
                                                        _ = v57
                                                    default:
                                                        fmt.Println(v1, v4)
                                                    }
                                                case "hello":
                                                    var v58 bool = (!false)
                                                    _ = v58
                                                case "":
                                                    v35 = ("foo bar" + v35)
                                                default:
                                                    v59 := v2[42]
                                                    _ = v59
                                                }
                                                _ = v50
                                            }
                                            _ = fi43
                                        }
                                        v60 := GenS8{F0: -95, F1: 564}
                                        _ = fi42
                                        _ = v60
                                    }
                                }
                                v61 := []int{(v3 >> ((255) & 63))}
                                _ = v61
                            } else {
                                fmt.Println(v4)
                                v62 := GenS3{F0: 32767}
                                var v63 int32 = (v4 << ((255) & 63))
                                _ = v62
                                _ = v63
                            }
                            _ = v35
                        } else {
                            var v64 int32 = v4
                            if false {
                                v65 := GenS7{F0: 2, F1: 305, F2: 136}
                                v66 := GenS9{F0: true, F1: 255, F2: 91}
                                _ = v65
                                _ = v66
                            }
                            _ = v64
                        }
                        v1 -= v1
                        var v67 uint = 255
                        _ = v67
                    } else {
                        if false {
                            switch "hello" {
                            case "foo bar":
                                if true {
                                    for fi68 := 0; fi68 < 1; fi68++ {
                                        v1 += float32((float64(v1) + 0.1))
                                        v69 := GenS10{F0: true, F1: 0}
                                        _ = fi68
                                        _ = v69
                                    }
                                } else {
                                    if true {
                                        var v70 uint8 = 0
                                        if ((((("" + "\n\t") + "") + ("foo bar" + "a")) + "hello") < "\n\t") {
                                            fmt.Println(v70, v3)
                                        }
                                        fmt.Println(v3, v4)
                                        _ = v70
                                    } else {
                                        for fi71 := 0; fi71 < 16; fi71++ {
                                            if (662 <= 2) {
                                                v3 += fi71
                                            } else {
                                                if true {
                                                    v72 := v2[0]
                                                    _ = v72
                                                } else {
                                                    if (1 <= 0) {
                                                        v73 := GenS11{F0: 1, F1: (v4 / ((v4) | 1))}
                                                        v4 = (v4 >> ((0) & 63))
                                                        for fi74 := 0; fi74 < 7; fi74++ {
                                                            v75 := []int{fi74, 32767}
                                                            v76 := GenS12{F0: 560, F1: -128, F2: -128}
                                                            v4 = 0
                                                            var v77 uint8 = 1
                                                            v78 := GenS13{F0: 315}
                                                            _ = fi74
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        _ = v73
                                                    } else {
                                                        if true {
                                                            v79 := []int{}
                                                            v1 += (3.14159 + 0.0)
                                                            _ = v79
                                                        } else {
                                                            delete(v2, 0)
                                                        }
                                                        for fi80 := 0; fi80 < 5; fi80++ {
                                                            var v81 string = "hello"
                                                            var v82 string = ""
                                                            _ = fi80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    }
                                                    var v83 uint = 18446744073709551615
                                                    _ = v83
                                                }
                                            }
                                            _ = fi71
                                        }
                                    }
                                }
                            case "hello":
                                v3 = (v3 / ((v3) | 1))
                            case "a":
                                v84 := GenS14{F0: 0, F1: (!false)}
                                _ = v84
                            }
                        } else {
                            if false {
                                var v85 uint8 = 42
                                var v86 uint8 = (^(v85))
                                _ = v85
                                _ = v86
                            } else {
                                if true {
                                    var v87 uint16 = 733
                                    _ = v87
                                } else {
                                    v1 = v1
                                    v88 := GenS15{F0: 18446744073709551615, F1: -116}
                                    v89 := GenS15{F0: 42, F1: 100}
                                    _ = v88
                                    _ = v89
                                }
                                v90 := GenS16{F0: 429}
                                fmt.Println(v4)
                                _ = v90
                            }
                        }
                        if (!((float64((0.0 - 0.1)) < (3.14159 - float64((v1 / -2.25)))) && (true && (0.1 >= v1)))) {
                            v1 = v1
                        }
                        var v91 uint64 = 1
                        _ = v91
                    }
                default:
                    fmt.Println(v1)
                }
                var v92 uint = 1
                _ = v92
            }
            var v93 int = 1
            v94 := GenS17{F0: 9223372036854775807}
            _ = v93
            _ = v94
        }
        _ = v14
    }
    if false {
        for fi95 := 0; fi95 < 11; fi95++ {
            fmt.Println(v3)
            v3 -= v3
            _ = fi95
        }
        for fi96 := 0; fi96 < 3; fi96++ {
            v3 = (v3 >> ((18446744073709551615) & 63))
            v3 += 128
            v97 := GenS7{F0: 156, F1: 759, F2: 2}
            _ = fi96
            _ = v97
        }
        v3 = v3
    } else {
        fmt.Println(v4)
        var v98 int64 = -1
        v99 := GenS12{F0: 65535, F1: 127, F2: 41}
        _ = v98
        _ = v99
    }
    var v100 string = "\n\t"
    for fi101 := 0; fi101 < 19; fi101++ {
        for fi102 := 0; fi102 < 18; fi102++ {
            if ((!(!(7 < 685))) || (!(62 >= 2))) {
                v103 := v2[1]
                v104 := GenS14{F0: -128, F1: (715 >= 0)}
                _ = v103
                _ = v104
            } else {
                v4 -= (v4 << ((255) & 63))
                v105 := GenS4{F0: 42}
                var v106 uint8 = 128
                if false {
                    v4 = v4
                    v106 -= v106
                    var v107 int64 = 9223372036854775807
                    _ = v107
                } else {
                    v108 := GenS18{F0: 1000, F1: 42, F2: (v106 | 128)}
                    v109 := []int{}
                    v4 = int32(fi102)
                    _ = v108
                    _ = v109
                }
                v110 := []int{13, (v3 + 2147483647), v3}
                _ = v105
                _ = v106
                _ = v110
            }
            _ = fi102
        }
        var v111 int64 = 1
        _ = fi101
        _ = v111
    }
    var v112 int64 = -9223372036854775808
    if false {
        fmt.Println(v100, v112)
    } else {
        v112 -= v112
        for fi113 := 0; fi113 < 14; fi113++ {
            var v114 uint8 = 2
            var v115 float64 = ((3.14159 - 0.1) - -2.25)
            _ = fi113
            _ = v114
            _ = v115
        }
        fmt.Println(v1, v3, v112)
    }
    _ = len(v2)
    v100 = "foo bar"
    v3 = (v3 ^ -761)
    var v116 string = "hello"
    fmt.Println(v1, v112)
    v112 = 630
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v112:", v112)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v112
    _ = v116
    _ = v2
    _ = v3
    _ = v4
}
