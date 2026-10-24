package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 string; F1 int32; F2 uint64 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 uint16; F1 uint32 }
type GenS14 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS15 struct { F0 int64; F1 int32 }
type GenS16 struct { F0 uint16; F1 string }
type GenS17 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint8) uint8 {
    if true {
        switch 2147483647 {
        case -312:
            var v5 int = 32767
            _ = v5
        }
    } else {
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 -= (^(p0))
            p0 = p1
            p0 = 160
            _ = fi6
        }
        p0 = p0
    }
    var v7 float64 = 1.5
    var v8 uint = 42
    var v9 int32 = -214
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 11
}
func genFn1(p0 int, p1 uint64) int64 {
    p0 += p0
    if true {
        var v10 bool = (-1 > -1)
        fmt.Println(p0, p1)
        fmt.Println(p1, p1, p1)
        _ = v10
    } else {
        var v11 int64 = 1
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(p0)
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := []int{42}
                fmt.Println(p1, v11, p1)
                v15 := GenS1{F0: "foo bar", F1: 255, F2: p1}
                v16 := GenS1{F0: "\n\t", F1: 149, F2: (p1 ^ (p1 * (^(p1))))}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        _ = v11
    }
    v17 := GenS2{F0: p0, F1: 1}
    v18 := GenS3{F0: 65535, F1: 2147483648, F2: 42}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 42
    var v2 uint16 = v1
    var v3 bool = (0 <= -1)
    v4 := make(map[string]int)
    if v3 {
        if v3 {
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 int32 = 42
                _ = fi19
                _ = v20
            }
        } else {
            var v21 uint = 215
            _ = v21
        }
        v22 := GenS4{F0: genFn1(32768, 1), F1: ("hello" + "\n\t")}
        v23 := GenS3{F0: 391, F1: 4294967295, F2: 18446744073709551615}
        _ = v22
        _ = v23
    } else {
        v3 = (v3 || ((3.14159 / ((-2.25 + -2.25) + 3.14159)) > (0.0 + -2.25)))
        for fi24 := 0; fi24 < 8; fi24++ {
            switch fi24 {
            case -896:
                for fi25 := 0; fi25 < 8; fi25++ {
                    fmt.Println(v3, fi24)
                    fmt.Println(v2, v2)
                    switch 32768 {
                    case -172:
                        var v26 bool = v3
                        _ = v26
                    case 32767:
                        v2 += v2
                    case -9223372036854775808:
                        if v3 {
                            fmt.Println(v1, fi25)
                        }
                    default:
                        var v27 uint8 = uint8(v1)
                        _ = v27
                    }
                    switch 2147483647 {
                    case -984:
                        switch fi25 {
                        case 127:
                            v3 = v3
                        case 128:
                            v2 -= v2
                        case 9223372036854775807:
                            delete(v4, "hello")
                        default:
                            var v28 float32 = (0.0 + (float32((1.5 - 1e10)) / 0.1))
                            _ = v28
                        }
                    case -2147483648:
                        var v29 int64 = 1
                        var v30 int32 = int32(v2)
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                }
            case 4294967295:
                v31 := GenS1{F0: "foo bar", F1: genFn0(246, 0), F2: 4294967295}
                _ = v31
            case 42:
                v32 := GenS5{F0: (v2 / ((v2) | 1))}
                _ = v32
            default:
                v1 += (^(v2))
            }
            for fi33 := 0; fi33 < 14; fi33++ {
                v1 += v2
                v3 = false
                v1 += (v1 & v1)
                _ = fi33
            }
            _ = fi24
        }
        v34 := GenS5{F0: (v1 << ((1) & 63))}
        _ = v34
    }
    if v3 {
        if (150 == -20) {
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 int8 = 9
                var v37 int = (fi35 + (fi35 & 0))
                _ = fi35
                _ = v36
                _ = v37
            }
            var v38 string = "\n\t"
            switch v38 {
            case "a":
                v39 := GenS5{F0: 65535}
                _ = v39
            case "\n\t":
                v2 -= 0
            default:
                v40 := []int{127}
                _ = v40
            }
            _ = v38
        }
    }
    v1 = (v1 & (v1 ^ (v1 ^ 987)))
    v2 += v1
    if (-2147483648 == 2147483647) {
        v41 := GenS6{F0: 127, F1: 0}
        var v42 uint = 2
        _ = v41
        _ = v42
    } else {
        v2 = v1
        v2 += v2
        v43 := GenS7{F0: 42, F1: 4294967295, F2: 2}
        _ = v43
    }
    v2 += 687
    var v44 float32 = 0.0
    v45 := GenS7{F0: 2147483647, F1: 529, F2: 1}
    v46 := GenS8{F0: 184, F1: 1}
    v47 := GenS9{F0: "\n\t", F1: 0, F2: 269}
    v48 := GenS10{F0: 42, F1: 37}
    v3 = (!false)
    v49 := GenS9{F0: "a", F1: 0, F2: 2}
    if v3 {
        var v50 uint32 = 2147483648
        if true {
            fmt.Println(v50)
            for fi51 := 0; fi51 < 16; fi51++ {
                var v52 uint = 18446744073709551615
                _ = fi51
                _ = v52
            }
            v50 += (v50 + (v50 & 2147483648))
        } else {
            switch -778 {
            case 42:
                v3 = true
            case 1:
                if (v3 || (true || (!v3))) {
                    v50 -= (v50 - (^(v50)))
                    for fi53 := 0; fi53 < 3; fi53++ {
                        v50 = 1
                        var v54 int = 896
                        if false {
                            v55 := GenS8{F0: 2, F1: 808}
                            var v56 float32 = (v44 / v44)
                            delete(v4, "a")
                            _ = v55
                            _ = v56
                        }
                        _ = fi53
                        _ = v54
                    }
                } else {
                    var v57 int64 = 0
                    var v58 int64 = (^(v57))
                    var v59 uint32 = v50
                    _ = v57
                    _ = v58
                    _ = v59
                }
            }
            fmt.Println(v2, v3)
        }
        _ = v50
    } else {
        var v60 int16 = 0
        if (0 <= 1) {
            switch (("\n\t" + "a") + "hello") {
            case "foo bar":
                v3 = v3
            case "hello":
                var v61 uint8 = 99
                _ = v61
            default:
                switch 128 {
                case 128:
                    for fi62 := 0; fi62 < 6; fi62++ {
                        v44 = (v44 / (v44 * (v44 - (v44 / -2.25))))
                        var v63 int8 = -10
                        var v64 int64 = -556
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                case 870:
                    switch "a" {
                    case "":
                        switch 9223372036854775807 {
                        case -657:
                            for fi65 := 0; fi65 < 17; fi65++ {
                                switch (fi65 * -1) {
                                case 0:
                                    for fi66 := 0; fi66 < 6; fi66++ {
                                        switch fi65 {
                                        case 1:
                                            var v67 int64 = -1
                                            _ = v67
                                        }
                                        v68 := GenS8{F0: 1, F1: 131}
                                        _ = fi66
                                        _ = v68
                                    }
                                }
                                v3 = v3
                                _ = fi65
                            }
                        case -2147483648:
                            switch (("\n\t" + "hello") + "") {
                            case "hello":
                                if (v1 > 0) {
                                    switch 0 {
                                    case -305:
                                        v69 := []int{-412, 1, 321}
                                        _ = v69
                                    case 127:
                                        var v70 int = 127
                                        _ = v70
                                    case 32768:
                                        var v71 bool = v3
                                        _ = v71
                                    }
                                    if false {
                                        v44 = v44
                                        var v72 float64 = 0.0
                                        var v73 uint64 = 18446744073709551615
                                        _ = v72
                                        _ = v73
                                    }
                                    if (4294967295 <= 255) {
                                        var v74 uint8 = 128
                                        var v75 uint64 = uint64(v1)
                                        if ((v74 * v74) == v74) {
                                            fmt.Println(v60)
                                        } else {
                                            v76 := GenS11{F0: 2147483647}
                                            v1 += (v2 / ((v2) | 1))
                                            _ = v76
                                        }
                                        _ = v74
                                        _ = v75
                                    } else {
                                        var v77 int64 = 135
                                        delete(v4, ("a" + "hello"))
                                        _ = v77
                                    }
                                }
                            case "":
                                var v78 bool = (-838 != -668)
                                _ = v78
                            case "\n\t":
                                for fi79 := 0; fi79 < 2; fi79++ {
                                    v80 := GenS12{F0: -914, F1: -346}
                                    v2 += v1
                                    v81 := GenS2{F0: 532, F1: -457}
                                    _ = fi79
                                    _ = v80
                                    _ = v81
                                }
                            }
                        case 2147483647:
                            v82 := []int{460, -2147483648, -892, -382}
                            _ = v82
                        }
                    case "a":
                        if ((!(!v3)) && v3) {
                            v2 = 390
                            var v83 int8 = -1
                            _ = v83
                        } else {
                            if (9223372036854775807 != 9223372036854775807) {
                                if v3 {
                                    fmt.Println(v1, v3, v1)
                                    var v84 float32 = float32(1.5)
                                    if (!(!v3)) {
                                        v85 := GenS13{F0: v2, F1: 4294967295}
                                        var v86 int16 = v60
                                        var v87 uint = 65535
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    } else {
                                        var v88 int16 = -71
                                        if true {
                                            v89 := []int{}
                                            switch ("" + "a") {
                                            case "hello":
                                                if v3 {
                                                    v44 += v44
                                                }
                                            case "\n\t":
                                                var v90 string = "hello"
                                                _ = v90
                                            case "a":
                                                if v3 {
                                                    if (false || v3) {
                                                        v88 += (v88 & (v88 - (v88 * v60)))
                                                        for fi91 := 0; fi91 < 17; fi91++ {
                                                            for fi92 := 0; fi92 < 14; fi92++ {
                                                                v93 := GenS13{F0: v2, F1: 4294967295}
                                                                _ = len(v4)
                                                                _ = fi92
                                                                _ = v93
                                                            }
                                                            _ = fi91
                                                        }
                                                    }
                                                }
                                            default:
                                                v94 := GenS14{F0: (-(v60)), F1: 0, F2: 4294967295}
                                                _ = v94
                                            }
                                            v95 := v4[("\n\t" + "\n\t")]
                                            _ = v89
                                            _ = v95
                                        } else {
                                            var v96 uint = 4294967295
                                            v97 := GenS15{F0: 9223372036854775807, F1: 42}
                                            _ = v96
                                            _ = v97
                                        }
                                        v98 := []int{-705, -84, 2147483647}
                                        _ = v88
                                        _ = v98
                                    }
                                    _ = v84
                                } else {
                                    var v99 float64 = float64(1e10)
                                    _ = v99
                                }
                                for fi100 := 0; fi100 < 11; fi100++ {
                                    fmt.Println(v3, fi100, fi100)
                                    _ = fi100
                                }
                                v101 := []int{0, 32768}
                                _ = v101
                            } else {
                                v60 -= -32768
                                if v3 {
                                    switch ("" + "\n\t") {
                                    case "foo bar":
                                        v102 := GenS16{F0: (^(v1)), F1: ""}
                                        _ = v102
                                    default:
                                        var v103 uint = 255
                                        _ = v103
                                    }
                                    switch "a" {
                                    case "":
                                        v44 += v44
                                    case "\n\t":
                                        var v104 float32 = v44
                                        _ = v104
                                    }
                                    switch 420 {
                                    case 128:
                                        for fi105 := 0; fi105 < 7; fi105++ {
                                            v106 := GenS16{F0: v1, F1: "a"}
                                            fmt.Println(v44, v44, v44)
                                            _ = fi105
                                            _ = v106
                                        }
                                    }
                                } else {
                                    fmt.Println(v60)
                                    var v107 uint32 = 1
                                    v108 := GenS17{F0: -317}
                                    _ = v107
                                    _ = v108
                                }
                                v4[("foo bar" + "hello")] = -1
                            }
                        }
                    case "\n\t":
                        if v3 {
                            var v109 int16 = (^(v60))
                            for fi110 := 0; fi110 < 16; fi110++ {
                                switch fi110 {
                                case -832:
                                    v111 := GenS7{F0: 2147483647, F1: 65535, F2: 18446744073709551615}
                                    _ = v111
                                case 104:
                                    var v112 uint32 = 2147483648
                                    _ = v112
                                case 0:
                                    for fi113 := 0; fi113 < 1; fi113++ {
                                        v114 := []int{fi113, fi110}
                                        _ = fi113
                                        _ = v114
                                    }
                                }
                                _ = fi110
                            }
                            _ = v109
                        } else {
                            v44 = (3.14159 - float32(1.5))
                            v60 = (^(v60))
                        }
                    }
                case 9223372036854775807:
                    v60 -= v60
                default:
                    v2 = (v2 << ((1) & 63))
                }
            }
        }
        v115 := GenS12{F0: 1, F1: -2147483648}
        _ = v115
        _ = v60
    }
    var v116 int16 = -845
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
