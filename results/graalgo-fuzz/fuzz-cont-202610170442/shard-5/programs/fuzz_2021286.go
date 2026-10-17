package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int64; F2 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 string; F1 int64; F2 int64 }
type GenS11 struct { F0 uint; F1 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16; F1 int; F2 int64 }
type GenS14 struct { F0 int64 }
func genFn0(p0 string, p1 uint32, p2 bool) uint8 {
    p2 = ((((((p0 + p0) + p0) + (p0 + "\n\t")) + p0) + p0) > "")
    var v5 uint32 = p1
    var v6 uint8 = 232
    switch 128 {
    case -545:
        var v7 uint = 18446744073709551615
        _ = v7
    case 205:
        switch "\n\t" {
        case "\n\t":
            var v8 int32 = int32(v6)
            _ = v8
        case "a":
            if (!(0.1 >= 0.1)) {
                var v9 int = 32768
                _ = v9
            }
        case "":
            v10 := GenS1{F0: 100}
            _ = v10
        default:
            switch p0 {
            case "a":
                if p2 {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 uint8 = (v6 % ((v6) | 1))
                        var v13 bool = p2
                        v14 := GenS2{F0: 557}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(v5, v5)
                } else {
                    v15 := GenS1{F0: 127}
                    for fi16 := 0; fi16 < 19; fi16++ {
                        p0 = p0
                        p1 -= (^(p1))
                        switch fi16 {
                        case 32768:
                            v17 := []int{fi16}
                            _ = v17
                        case 92:
                            if (-34 != 959) {
                                v18 := GenS3{F0: v6, F1: (42 < (p1 / ((534) | 1))), F2: 649}
                                v5 -= v5
                                _ = v18
                            }
                        case 2147483647:
                            switch (p0 + (p0 + ((("hello" + "\n\t") + (p0 + (p0 + p0))) + "foo bar"))) {
                            case "\n\t":
                                if false {
                                    var v19 int8 = -89
                                    var v20 uint = 255
                                    _ = v19
                                    _ = v20
                                }
                            case "":
                                v21 := GenS4{F0: p0}
                                _ = v21
                            default:
                                p2 = p2
                            }
                        }
                        _ = fi16
                    }
                    _ = v15
                }
            case "":
                fmt.Println(v6, p1, p1)
            case "hello":
                for fi22 := 0; fi22 < 4; fi22++ {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 int = fi22
                        v25 := GenS3{F0: (v6 / (((v6 & v6)) | 1)), F1: true, F2: (^(v5))}
                        var v26 int8 = int8(v6)
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := []int{(fi22 * fi22), (fi22 * (fi22 % (((fi22 | (fi22 & fi22))) | 1)))}
                    _ = fi22
                    _ = v27
                }
            default:
                for fi28 := 0; fi28 < 2; fi28++ {
                    p2 = p2
                    switch -1 {
                    case 32768:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            v30 := []int{}
                            p2 = p2
                            _ = fi29
                            _ = v30
                        }
                    case -2147483648:
                        if p2 {
                            p1 -= (v5 % ((2147483647) | 1))
                            if p2 {
                                v31 := GenS5{F0: fi28}
                                var v32 uint16 = 624
                                _ = v31
                                _ = v32
                            }
                        } else {
                            var v33 bool = (2 > p1)
                            _ = v33
                        }
                    case 177:
                        v34 := []int{(fi28 * -162), 42, fi28, -394}
                        _ = v34
                    }
                    _ = fi28
                }
            }
        }
    case 9223372036854775807:
        var v35 int8 = -128
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v6 * 42)
}
func genFn1(p0 uint8, p1 bool, p2 int16) int8 {
    var v36 float32 = 0.0
    if p1 {
        if p1 {
            for fi37 := 0; fi37 < 3; fi37++ {
                var v38 uint32 = uint32(p0)
                fmt.Println(p0, v36, p2)
                v39 := GenS5{F0: fi37}
                _ = fi37
                _ = v38
                _ = v39
            }
            v40 := []int{}
            _ = v40
        }
        v41 := GenS1{F0: -115}
        if false {
            v42 := GenS5{F0: -9223372036854775808}
            p1 = (2 < 2)
            _ = v42
        }
        _ = v41
    }
    p0 = 0
    var v43 string = ("" + (("hello" + "a") + ""))
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v43
    return 100
}

func main() {
    var v1 uint64 = 57
    var v2 uint32 = 0
    var v3 bool = (uint8(v1) >= 166)
    v4 := []string{"hello"}
    fmt.Println(v2, v3, v2)
    v4 = append(v4, ("a" + "hello"))
    v4 = append(v4, "\n\t")
    v44 := GenS6{F0: 2}
    if (!false) {
        var v45 int = 128
        _ = v45
    } else {
        if true {
            if false {
                var v46 int = 230
                _ = v46
            } else {
                if v3 {
                    var v47 uint8 = 128
                    switch 33 {
                    case -271:
                        v1 = uint64(v2)
                    case -2147483648:
                        var v48 int8 = 42
                        _ = v48
                    case -1:
                        fmt.Println(v47, v3)
                    default:
                        var v49 int8 = 1
                        _ = v49
                    }
                    _ = v47
                } else {
                    var v50 uint16 = 2
                    _ = v50
                }
                if v3 {
                    var v51 uint8 = 0
                    var v52 int16 = 42
                    var v53 uint32 = (v2 + (v2 & (v2 << ((65535) & 63))))
                    _ = v51
                    _ = v52
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 9; fi54++ {
                        switch (fi54 & fi54) {
                        case 2147483647:
                            for fi55 := 0; fi55 < 12; fi55++ {
                                var v56 bool = false
                                v57 := GenS1{F0: 0}
                                _ = len(v4)
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                        case 127:
                            if true {
                                for fi58 := 0; fi58 < 1; fi58++ {
                                    v59 := GenS3{F0: 255, F1: ((9223372036854775807 > 0) && v3), F2: 2147483648}
                                    v60 := GenS7{F0: 362, F1: 1, F2: (1 >= 1)}
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            } else {
                                var v61 uint = 916
                                v3 = v3
                                if v3 {
                                    v61 += v61
                                } else {
                                    var v62 int32 = 1
                                    if (uint8(v2) <= 127) {
                                        if (v3 || (v3 && false)) {
                                            v63 := GenS3{F0: 1, F1: false, F2: 816}
                                            _ = v63
                                        }
                                    } else {
                                        var v64 float32 = 3.14159
                                        v62 -= v62
                                        _ = v64
                                    }
                                    v61 += 42
                                    _ = v62
                                }
                                _ = v61
                            }
                        }
                        v2 -= (v2 / ((993) | 1))
                        for fi65 := 0; fi65 < 5; fi65++ {
                            var v66 float32 = 0.0
                            v67 := GenS4{F0: "\n\t"}
                            fmt.Println(v1, fi54)
                            if (v3 && v3) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v68 := GenS8{F0: 879}
                                _ = v68
                            } else {
                                fmt.Println(fi54)
                                var v69 uint = 782
                                switch fi54 {
                                case 32767:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                _ = v69
                            }
                            fmt.Println(v2, fi54)
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                        _ = fi54
                    }
                }
                if (-32768 <= 1000) {
                    var v70 uint = 200
                    var v71 uint64 = v1
                    _ = v70
                    _ = v71
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v2, v3, v1)
                    v2 -= v2
                }
            }
            if (!v3) {
                for fi72 := 0; fi72 < 7; fi72++ {
                    var v73 uint8 = 223
                    _ = fi72
                    _ = v73
                }
            } else {
                v4 = append(v4, "a")
            }
            fmt.Println(v3, v2)
        } else {
            var v74 int64 = 0
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            _ = v74
        }
        for fi75 := 0; fi75 < 8; fi75++ {
            for fi76 := 0; fi76 < 6; fi76++ {
                v2 -= 2
                var v77 uint64 = 951
                var v78 int = (fi76 | fi76)
                _ = fi76
                _ = v77
                _ = v78
            }
            _ = fi75
        }
    }
    for fi79 := 0; fi79 < 8; fi79++ {
        v3 = (84 < 42)
        switch 610 {
        case -916:
            var v80 int8 = 42
            _ = v80
        case -9223372036854775808:
            v2 = v2
        case -760:
            v4 = append(v4, "\n\t")
        default:
            v81 := GenS9{F0: ("\n\t" + ""), F1: 1000}
            _ = v81
        }
        v82 := GenS1{F0: 127}
        _ = fi79
        _ = v82
    }
    v2 -= (v2 << ((4294967295) & 63))
    v1 = (v1 / (((v1 >> ((178) & 63))) | 1))
    if v3 {
        v83 := GenS7{F0: 775, F1: 0, F2: v3}
        var v84 uint = 255
        fmt.Println(v3, v2, v2)
        _ = v83
        _ = v84
    } else {
        if v3 {
            for fi85 := 0; fi85 < 3; fi85++ {
                fmt.Println(v3, fi85)
                if v3 {
                    v1 = 18446744073709551615
                }
                _ = fi85
            }
        }
        if false {
            var v86 float64 = 1e10
            _ = v86
        } else {
            switch "" {
            case "":
                switch 127 {
                case 2147483647:
                    _ = len(v4)
                case 282:
                    v87 := GenS1{F0: genFn1(42, (!("a" > ((("hello" + "") + "foo bar") + (("hello" + "hello") + "hello")))), -866)}
                    _ = v87
                default:
                    for fi88 := 0; fi88 < 13; fi88++ {
                        switch (fi88 ^ (fi88 ^ (^(fi88)))) {
                        case -910:
                            for fi89 := 0; fi89 < 7; fi89++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v1 = (v1 - 100)
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    switch ("hello" + "\n\t") {
                                    case "a":
                                        v90 := GenS8{F0: -32768}
                                        _ = v90
                                    default:
                                        if v3 {
                                            v3 = (-23 < -128)
                                            _ = len(v4)
                                        } else {
                                            v2 = 2147483648
                                            fmt.Println(fi89)
                                            if (!(-395 > fi88)) {
                                                if v3 {
                                                    var v91 string = "hello"
                                                    var v92 uint16 = 255
                                                    _ = v91
                                                    _ = v92
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi89
                            }
                        default:
                            v93 := GenS10{F0: ("a" + "a"), F1: -1, F2: 0}
                            _ = v93
                        }
                        _ = fi88
                    }
                }
            case "a":
                v94 := GenS11{F0: 1, F1: 127}
                _ = v94
            default:
                v95 := GenS2{F0: 1}
                _ = v95
            }
            v96 := GenS4{F0: ""}
            _ = v96
        }
    }
    fmt.Println(v3)
    var v97 uint64 = (v1 - (v1 + v1))
    var v98 string = (("foo bar" + "") + "")
    switch 127 {
    case -2147483648:
        if (!v3) {
            v97 = 42
        }
    }
    fmt.Println(v1, v2, v1)
    var v99 uint8 = 0
    switch -608 {
    case -371:
        switch 42 {
        case -2147483648:
            var v100 string = (("" + "foo bar") + (v98 + ("foo bar" + (v98 + v98))))
            _ = v100
        case 127:
            switch "\n\t" {
            case "a":
                v97 = v1
            }
        default:
            switch v98 {
            case "a":
                fmt.Println(v99, v2, v98)
            case "":
                if v3 {
                    if false {
                        for fi101 := 0; fi101 < 1; fi101++ {
                            if false {
                                v102 := GenS12{F0: (v98 + ("" + v98))}
                                if (v3 || v3) {
                                    switch fi101 {
                                    case 446:
                                        fmt.Println(v2)
                                    case 4294967295:
                                        var v103 int8 = 50
                                        _ = v103
                                    case 1:
                                        v1 = 18446744073709551615
                                    }
                                }
                                for fi104 := 0; fi104 < 12; fi104++ {
                                    fmt.Println(fi101, v97)
                                    if false {
                                        var v105 int16 = 1000
                                        var v106 float32 = float32(float64(((1.5 + -2.25) / (((0.1 * -2.25) - (((0.0 / ((1.5 / 1.0) + (0.1 * 1.5))) + (1.5 + ((float32(((float64((0.0 / -2.25)) * 3.14159) * (((3.14159 + 1e10) - (3.14159 * 1e10)) - (1.5 / 0.1)))) - ((3.14159 + 1.5) / 1.0)) * 1e10))) * 0.0)) + 0.0))))
                                        v107 := GenS11{F0: 4294967295, F1: genFn0(v98, 2147483648, (((1.5 * 3.14159) > 1e10) && v3))}
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                    } else {
                                        var v108 uint8 = genFn0(v98, 65535, true)
                                        v98 = v98
                                        _ = v108
                                    }
                                    var v109 int32 = -2147483648
                                    _ = fi104
                                    _ = v109
                                }
                                _ = v102
                            }
                            _ = fi101
                        }
                        var v110 uint16 = 0
                        var v111 uint16 = v110
                        _ = v110
                        _ = v111
                    }
                    fmt.Println(v98)
                    switch v98 {
                    case "hello":
                        v4 = append(v4, v98)
                    default:
                        v112 := GenS13{F0: -889, F1: -1, F2: -786}
                        _ = v112
                    }
                }
            case "hello":
                for fi113 := 0; fi113 < 13; fi113++ {
                    if true {
                        var v114 string = (v98 + v98)
                        for fi115 := 0; fi115 < 7; fi115++ {
                            if (!true) {
                                var v116 uint = 924
                                var v117 string = v98
                                var v118 uint32 = v2
                                _ = v116
                                _ = v117
                                _ = v118
                            }
                            _ = fi115
                        }
                        _ = v114
                    } else {
                        fmt.Println(v1)
                    }
                    v4 = append(v4, v98)
                    _ = fi113
                }
            }
        }
    case -2147483648:
        v1 += v1
    }
    v119 := GenS14{F0: 418}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v97
    _ = v98
    _ = v99
}
