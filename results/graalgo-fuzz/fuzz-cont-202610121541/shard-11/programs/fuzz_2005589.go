package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint32; F1 uint64; F2 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int, p1 uint16) int {
    v6 := GenS1{F0: -1, F1: -9223372036854775808, F2: p1}
    var v7 float32 = (1.5 * 0.0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p0 + (p0 + p0))
}
func genFn1() int16 {
    var v8 uint64 = 1
    var v9 bool = (42 != 44)
    _ = v8
    _ = v9
    return 1000
}

func main() {
    var v1 int32 = 0
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    var v4 bool = true
    v5 := make(map[uint32]int)
    v4 = ((!((1e10 * (0.0 + 0.0)) > ((0.1 / 1.0) * 3.14159))) || (1.5 < 1e10))
    v1 += (-(v1))
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: -32768, F1: int64(fi10), F2: 0}
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(v2, fi12, fi10)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    if v4 {
        v4 = v4
        if (!v4) {
            switch "" {
            case "\n\t":
                var v13 uint32 = 558
                _ = v13
            default:
                var v14 int = 127
                _ = v14
            }
        } else {
            v1 -= (v1 - v1)
            v15 := []int{42}
            if true {
                var v16 int64 = 682
                _ = v16
            } else {
                v17 := GenS1{F0: -32768, F1: 560, F2: 1}
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 int16 = 32767
                    var v20 int16 = 535
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v21 := GenS2{F0: -452}
                _ = v17
                _ = v21
            }
            _ = v15
        }
    } else {
        if (true && false) {
            var v22 uint32 = 990
            var v23 uint64 = (v2 / (((v2 * 591)) | 1))
            _ = v22
            _ = v23
        }
        if false {
            if (0 == 891) {
                v1 = (-(v1))
            }
            fmt.Println(v1)
            fmt.Println(v2)
        } else {
            v2 = v2
            if v4 {
                var v24 uint8 = 1
                if (true || (0 >= 32768)) {
                    var v25 float32 = -2.25
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v1 -= (^(v1))
                        _ = fi26
                    }
                    var v27 int16 = 0
                    fmt.Println(v24)
                    _ = v27
                }
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 uint16 = 42
                    var v30 uint32 = 42
                    if v4 {
                        var v31 int32 = (v1 * (-(v1)))
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v24
            } else {
                v4 = (((v4 || (!(-9223372036854775808 < -1))) || false) || (("a" + "foo bar") <= "hello"))
                v32 := GenS1{F0: -1, F1: 0, F2: 42}
                _ = v32
            }
            for fi33 := 0; fi33 < 3; fi33++ {
                if ((-32768 < 32767) && true) {
                    switch fi33 {
                    case -175:
                        var v34 int64 = 97
                        _ = v34
                    default:
                        v35 := GenS2{F0: -32768}
                        _ = v35
                    }
                    var v36 int = (fi33 - fi33)
                    if v4 {
                        if v4 {
                            for fi37 := 0; fi37 < 2; fi37++ {
                                switch "" {
                                case "":
                                    var v38 bool = (false || true)
                                    _ = v38
                                case "hello":
                                    var v39 string = ("hello" + "hello")
                                    _ = v39
                                case "foo bar":
                                    if (!(1.5 < 0.0)) {
                                        var v40 int = (v36 / ((genFn0((v36 % (((fi33 | fi33)) | 1)), 0)) | 1))
                                        var v41 int16 = 361
                                        fmt.Println(v40, fi37, v4)
                                        fmt.Println(v41)
                                        _ = v40
                                        _ = v41
                                    }
                                    switch ("" + (("hello" + ("foo bar" + ("foo bar" + "foo bar"))) + "hello")) {
                                    case "a":
                                        if (576 == 255) {
                                            var v42 int8 = 100
                                            switch -485 {
                                            case 1:
                                                if (!v4) {
                                                    fmt.Println(v2)
                                                    v42 = -71
                                                } else {
                                                    _ = len(v3)
                                                    var v43 int = v36
                                                    v44 := GenS3{F0: 42, F1: 1, F2: -2147483648}
                                                    _ = v43
                                                    _ = v44
                                                }
                                                v36 -= (fi37 | (fi33 | genFn0((-(fi37)), 32767)))
                                            }
                                            _ = v42
                                        } else {
                                            v45 := []int{}
                                            _ = v45
                                        }
                                        switch "a" {
                                        case "hello":
                                            if v4 {
                                                switch (("" + ((((("hello" + "\n\t") + "hello") + "") + "a") + "")) + "") {
                                                case "\n\t":
                                                    v36 -= fi37
                                                    switch fi33 {
                                                    case 42:
                                                        var v46 uint = 255
                                                        v1 = 818
                                                        var v47 uint64 = v2
                                                        fmt.Println(v4, v4, v47)
                                                        _ = v46
                                                        _ = v47
                                                    default:
                                                        if true {
                                                            if ((fi33 | (fi33 >> ((1) & 63))) > fi33) {
                                                                v48 := GenS4{F0: 363}
                                                                v49 := GenS5{F0: (!((fi37 + (v36 * (fi37 << ((626) & 63)))) > (fi33 + fi37)))}
                                                                v36 = fi33
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                fmt.Println(v36, v4)
                                                                v50 := []int{}
                                                                var v51 bool = v4
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        } else {
                                                            var v52 float64 = 1.5
                                                            _ = v52
                                                        }
                                                    }
                                                    v53 := GenS4{F0: 655}
                                                    _ = v53
                                                case "":
                                                    fmt.Println(v4)
                                                    if ((-(v36)) < v36) {
                                                        v36 -= genFn0(-172, 255)
                                                        if true {
                                                            fmt.Println(v36, fi37)
                                                            _ = len(v3)
                                                        } else {
                                                            delete(v5, 2)
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v54 float32 = -2.25
                                                                _ = v54
                                                            }
                                                            v4 = (32767 < int16(fi37))
                                                        }
                                                        v55 := GenS1{F0: genFn1(), F1: 0, F2: 32767}
                                                        var v56 uint8 = 2
                                                        v57 := []int{(fi33 / (((-(v36))) | 1)), genFn0(int(v1), 333)}
                                                        if v4 {
                                                            v36 -= (v36 << ((370) & 63))
                                                            if len(v57) >= 2 {
                                                                v57 = v57[1:len(v57)]
                                                            }
                                                            v58 := GenS5{F0: true}
                                                            _ = v58
                                                        } else {
                                                            v36 += fi37
                                                            var v59 uint16 = 255
                                                            _ = len(v57)
                                                            switch ("a" + "a") {
                                                            case "a":
                                                                v5[42] = fi33
                                                                switch fi33 {
                                                                case 979:
                                                                    var v60 uint8 = (v56 % ((2) | 1))
                                                                    v1 = 42
                                                                    _ = v60
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                            }
                                                            _ = len(v3)
                                                            _ = v59
                                                        }
                                                        _ = v55
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        var v61 float64 = (0.1 + -2.25)
                                                        var v62 int32 = -2147483648
                                                        if v4 {
                                                            if v4 {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v1 = 42
                                                                }
                                                                v61 += v61
                                                                var v63 bool = ((v2 | v2) == (v2 + v2))
                                                                _ = v63
                                                            } else {
                                                                var v64 bool = v4
                                                                switch (fi33 - genFn0((-(fi37)), 590)) {
                                                                case 0:
                                                                    v62 += (v1 - (v62 % ((v1) | 1)))
                                                                case 7:
                                                                    var v65 bool = false
                                                                    _ = v65
                                                                }
                                                                var v66 int64 = -9223372036854775808
                                                                _ = v64
                                                                _ = v66
                                                            }
                                                            v67 := []int{}
                                                            v4 = ((true || false) || (!false))
                                                            _ = v67
                                                        }
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    switch (("" + (("foo bar" + "hello") + "foo bar")) + "foo bar") {
                                                    case "":
                                                        v2 = (v2 | (v2 % ((18446744073709551615) | 1)))
                                                    case "\n\t":
                                                        v36 -= -1
                                                    default:
                                                        if (255 == 114) {
                                                            fmt.Println(v4, fi37)
                                                            if v4 {
                                                                if (v4 || (-17 != 42)) {
                                                                    if v4 {
                                                                        if v4 {
                                                                            v4 = false
                                                                            if (791 == 307) {
                                                                                if (!v4) {
                                                                                    v68 := []int{2147483647, fi33}
                                                                                    switch (("\n\t" + "a") + "hello") {
                                                                                    case "\n\t":
                                                                                        var v69 int = genFn0(fi33, 32767)
                                                                                        _ = v69
                                                                                    case "a":
                                                                                        var v70 uint = 65535
                                                                                        _ = v70
                                                                                    default:
                                                                                        var v71 int16 = -820
                                                                                        var v72 uint64 = (v2 | (v2 >> ((42) & 63)))
                                                                                        if (!((v1 << ((42) & 63)) == v1)) {
                                                                                            var v73 int8 = -1
                                                                                            var v74 int64 = 1
                                                                                            v73 -= -128
                                                                                            v4 = (31 == 42)
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        } else {
                                                                                            fmt.Println(v2, v2, fi33)
                                                                                            if v4 {
                                                                                                if len(v68) > 0 {
                                                                                                    v68[0] = (fi33 ^ genFn0(v36, 865))
                                                                                                }
                                                                                                switch fi33 {
                                                                                                case 9223372036854775807:
                                                                                                    v75 := v5[2]
                                                                                                    _ = v75
                                                                                                case 4294967295:
                                                                                                    var v76 bool = (v4 && (v36 >= v36))
                                                                                                    _ = v76
                                                                                                case -656:
                                                                                                    v77 := GenS5{F0: true}
                                                                                                    _ = v77
                                                                                                default:
                                                                                                    var v78 uint = 1
                                                                                                    _ = v78
                                                                                                }
                                                                                                v79 := GenS3{F0: v1, F1: 0, F2: (-(fi37))}
                                                                                                _ = v79
                                                                                            } else {
                                                                                                var v80 int64 = -9223372036854775808
                                                                                                _ = v80
                                                                                            }
                                                                                        }
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                    }
                                                                                    var v81 uint64 = (^(v2))
                                                                                    var v82 int16 = 42
                                                                                    if (!(1 != 65535)) {
                                                                                        v83 := GenS4{F0: 0}
                                                                                        v81 = (v2 >> ((914) & 63))
                                                                                        _ = v83
                                                                                    }
                                                                                    var v84 uint = 917
                                                                                    _ = v68
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                            v85 := []int{fi33, (^(fi37)), fi33}
                                                                            v86 := GenS6{F0: (v2 - (v2 * (^(v2)))), F1: 812}
                                                                            var v87 float32 = 1e10
                                                                            var v88 int = (fi33 / ((v36) | 1))
                                                                            _ = v85
                                                                            _ = v86
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        v89 := GenS7{F0: 42, F1: v2, F2: (0 >= 325)}
                                                                        var v90 float32 = (0.1 - (1.5 - (0.1 / ((float32(1.5) + (-2.25 + 0.1)) * 1e10))))
                                                                        _ = v89
                                                                        _ = v90
                                                                    }
                                                                    var v91 uint16 = 32767
                                                                    v92 := GenS5{F0: false}
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                v93 := v5[2147483647]
                                                                var v94 int32 = (v1 + (v1 << ((387) & 63)))
                                                                _ = v93
                                                                _ = v94
                                                            } else {
                                                                var v95 uint = 18446744073709551615
                                                                var v96 int64 = -9223372036854775808
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                            v1 = 1
                                                        }
                                                    }
                                                    var v97 uint32 = 809
                                                    _ = len(v5)
                                                    _ = v97
                                                default:
                                                    if (!v4) {
                                                        var v98 int16 = -1
                                                        v99 := GenS2{F0: (v98 << ((555) & 63))}
                                                        _ = v98
                                                        _ = v99
                                                    } else {
                                                        if ((v4 && v4) && v4) {
                                                            fmt.Println(v1, v2, v4)
                                                            v2 += (v2 / ((534) | 1))
                                                        }
                                                    }
                                                }
                                                var v100 uint64 = 493
                                                var v101 int32 = v1
                                                _ = v100
                                                _ = v101
                                            } else {
                                                v2 -= v2
                                            }
                                        case "foo bar":
                                            switch fi33 {
                                            case -9223372036854775808:
                                                if (!(v4 && (3.14159 != 3.14159))) {
                                                    var v102 string = "hello"
                                                    var v103 int8 = -62
                                                    v104 := GenS7{F0: 2147483647, F1: v2, F2: ((^(v2)) > v2)}
                                                    v105 := GenS8{F0: v1}
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    var v106 uint16 = 746
                                                    _ = v106
                                                }
                                            case 32767:
                                                var v107 uint32 = 307
                                                _ = v107
                                            }
                                        default:
                                            var v108 uint32 = 1
                                            var v109 uint32 = v108
                                            _ = v108
                                            _ = v109
                                        }
                                    case "hello":
                                        if true {
                                            v4 = (805 != 42)
                                            fmt.Println(fi37, v4)
                                            var v110 uint8 = 56
                                            v111 := []int{v36, (fi33 >> ((2) & 63)), fi37, (v36 >> ((954) & 63))}
                                            _ = v110
                                            _ = v111
                                        }
                                    case "\n\t":
                                        v1 = v1
                                    }
                                }
                                v112 := []int{}
                                _ = fi37
                                _ = v112
                            }
                            switch ("\n\t" + "hello") {
                            case "a":
                                var v113 uint8 = 1
                                _ = v113
                            }
                            v4 = true
                        }
                        v114 := []int{(^(fi33))}
                        _ = v114
                    }
                    _ = v36
                }
                var v115 uint32 = 0
                var v116 bool = (false || v4)
                _ = fi33
                _ = v115
                _ = v116
            }
        }
        v117 := GenS9{F0: (v1 & v1)}
        _ = v117
    }
    switch 2147483647 {
    case 32767:
        if v4 {
            if (2 < 1) {
                fmt.Println(v1, v4)
            }
            for fi118 := 0; fi118 < 8; fi118++ {
                for fi119 := 0; fi119 < 6; fi119++ {
                    var v120 uint64 = v2
                    var v121 uint8 = 255
                    _ = fi119
                    _ = v120
                    _ = v121
                }
                var v122 float64 = 0.0
                var v123 int8 = -26
                _ = fi118
                _ = v122
                _ = v123
            }
        }
    case 0:
        v124 := GenS6{F0: (v2 % ((v2) | 1)), F1: (v1 - v1)}
        _ = v124
    }
    v125 := GenS10{F0: 2147483647, F1: 4294967295}
    var v126 uint = 4294967295
    v2 += (v2 * v2)
    v126 = v126
    if (false || (!v4)) {
        var v127 int = 289
        _ = v127
    } else {
        var v128 uint = v126
        _ = v128
    }
    v129 := GenS10{F0: 2, F1: 2147483647}
    var v130 uint = (^(v126))
    var v131 uint64 = v2
    v130 -= 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v125
    _ = v126
    _ = v129
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
