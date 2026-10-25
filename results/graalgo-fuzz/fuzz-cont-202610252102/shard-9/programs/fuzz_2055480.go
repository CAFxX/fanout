package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int }
type GenS5 struct { F0 string; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int16; F1 int16; F2 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS11 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 int32, p1 int16) uint16 {
    var v4 int64 = 9223372036854775807
    fmt.Println(v4, v4)
    p0 += p0
    _ = p0
    _ = p1
    _ = v4
    return 65535
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]bool)
    var v3 int32 = 303
    if (true && true) {
        var v5 float64 = 1e10
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{0, (-(fi6)), (fi6 & fi6), (fi6 | fi6)}
        fmt.Println(fi6, v3, v1)
        v8 := GenS1{F0: 127}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (-1 >= 0) {
        var v9 uint16 = 1
        _ = v9
    } else {
        if true {
            var v10 int8 = -1
            v3 += v3
            fmt.Println(v3, v1)
            _ = v10
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int8 = -1
                var v13 uint16 = genFn0((v3 << ((533) & 63)), -692)
                fmt.Println(v1)
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    switch ("" + "\n\t") {
    case "foo bar":
        switch ("a" + ("hello" + "foo bar")) {
        case "hello":
            v1 = v1
        default:
            if true {
                _ = len(v2)
            } else {
                if false {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        var v15 uint64 = 42
                        _ = fi14
                        _ = v15
                    }
                } else {
                    v1 = -1
                }
            }
        }
    default:
        var v16 float64 = 3.14159
        _ = v16
    }
    var v17 uint16 = 733
    var v18 float32 = 1.5
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 int16 = 1000
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS1{F0: 128}
            if true {
                if false {
                    v23 := GenS1{F0: 208}
                    _ = len(v2)
                    _ = v23
                } else {
                    if true {
                        if (!(367 == 1)) {
                            var v24 string = ("\n\t" + ("a" + "a"))
                            fmt.Println(v1, fi19, v24)
                            v20 += (v20 + v20)
                            _ = v24
                        } else {
                            if ((806 > 2) || false) {
                                fmt.Println(v3)
                                if true {
                                    if (-16 == -125) {
                                        if (1 <= 784) {
                                            var v25 bool = ((-(v20)) < (v20 / ((v20) | 1)))
                                            switch ("\n\t" + ("a" + "\n\t")) {
                                            case "hello":
                                                v26 := GenS2{F0: 0, F1: 42}
                                                _ = v26
                                            default:
                                                v27 := []int{(fi21 / ((fi21) | 1)), fi19, 0, (-(fi21))}
                                                var v28 int8 = 1
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := v2["\n\t"]
                                            v20 = (v20 << ((84) & 63))
                                            _ = v25
                                            _ = v29
                                        }
                                        v30 := GenS1{F0: 1}
                                        _ = v30
                                    } else {
                                        fmt.Println(v20, fi21)
                                        _ = len(v2)
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v31 bool = ((3.14159 > v18) && (65535 == 1))
                                        _ = v31
                                    case "foo bar":
                                        fmt.Println(v3, fi21)
                                    case "hello":
                                        v20 = -959
                                    }
                                    v32 := []int{(fi19 * fi19), -1, (-(fi21))}
                                    var v33 int64 = -1
                                    var v34 int16 = v20
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                v2["hello"] = false
                            } else {
                                switch "a" {
                                case "\n\t":
                                    var v35 uint8 = 42
                                    _ = v35
                                default:
                                    if true {
                                        v17 -= (v17 - (^(v17)))
                                        if (float64(1.5) == 0.1) {
                                            v36 := []int{(fi21 + fi19), (fi21 % ((fi21) | 1)), fi21, fi19}
                                            if (2147483648 == 2147483647) {
                                                v36 = append(v36, 42)
                                                v37 := v2["foo bar"]
                                                var v38 uint8 = 1
                                                _ = v37
                                                _ = v38
                                            } else {
                                                v17 -= 255
                                                var v39 uint8 = 2
                                                v36 = append(v36, fi21)
                                                _ = v39
                                            }
                                            v18 = 0.1
                                            _ = v36
                                        }
                                        v18 += 0.1
                                        v40 := []int{(fi21 % ((-1) | 1)), int(v20), (fi21 >> ((433) & 63)), (fi19 >> ((2) & 63))}
                                        _ = v40
                                    }
                                }
                                var v41 float32 = (v18 - v18)
                                var v42 uint32 = 138
                                v43 := v2["a"]
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            if (false || (-1 >= (v3 % (((v1 + (-(v3)))) | 1)))) {
                                v44 := GenS2{F0: 799, F1: 4294967295}
                                v45 := GenS2{F0: 4294967295, F1: 305}
                                _ = v44
                                _ = v45
                            }
                            switch "a" {
                            case "hello":
                                v46 := []int{(fi21 / ((46) | 1))}
                                _ = v46
                            case "a":
                                v18 += float32(1.5)
                                if (false || true) {
                                    switch (("" + "foo bar") + "a") {
                                    case "foo bar":
                                        if (!(false || false)) {
                                            fmt.Println(v3, v20)
                                            if ((v17 - 32767) != v17) {
                                                var v47 float64 = ((-2.25 + 3.14159) + float64(-2.25))
                                                _ = v47
                                            } else {
                                                if (555 > 18446744073709551615) {
                                                    v3 += v3
                                                    if false {
                                                        var v48 int64 = -95
                                                        v49 := GenS1{F0: 0}
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        if false {
                                                            if true {
                                                                var v50 uint32 = 925
                                                                _ = v50
                                                            } else {
                                                                var v51 int = fi19
                                                                v52 := GenS3{F0: 224}
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            var v53 uint16 = 65535
                                                            if true {
                                                                v54 := []int{}
                                                                v18 = v18
                                                                _ = v54
                                                            } else {
                                                                fmt.Println(v17, fi21)
                                                                if (false || ((-(v20)) >= v20)) {
                                                                    delete(v2, "")
                                                                    if false {
                                                                        var v55 uint16 = 2
                                                                        v53 -= v17
                                                                        _ = v55
                                                                    } else {
                                                                        delete(v2, "a")
                                                                        v56 := GenS2{F0: 4294967295, F1: 849}
                                                                        var v57 string = ""
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                    v58 := GenS4{F0: (v17 % ((genFn0(int32(v53), (^(v20)))) | 1)), F1: 1, F2: (fi21 | (fi19 << ((1) & 63)))}
                                                                    _ = v58
                                                                }
                                                                v18 -= (0.1 / v18)
                                                                var v59 int64 = 42
                                                                v60 := GenS5{F0: ("hello" + "foo bar"), F1: (^(v17)), F2: 255}
                                                                fmt.Println(v1, fi21)
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            if (true || true) {
                                                                v1 += -1
                                                                _ = len(v2)
                                                            }
                                                            _ = v53
                                                        }
                                                    }
                                                    v61 := []int{}
                                                    _ = v61
                                                } else {
                                                    if (v20 < (v20 << ((1) & 63))) {
                                                        v3 -= (v3 & v3)
                                                    }
                                                    v20 -= v20
                                                    v62 := v2["a"]
                                                    var v63 uint8 = 147
                                                    v1 += v3
                                                    _ = v62
                                                    _ = v63
                                                }
                                                var v64 uint8 = 180
                                                _ = v64
                                            }
                                            if ((v18 / v18) > ((0.1 * v18) - 3.14159)) {
                                                v1 -= 1
                                                v65 := GenS6{F0: v20, F1: (-(v20)), F2: 127}
                                                _ = v65
                                            }
                                        } else {
                                            if (!(0 >= v20)) {
                                                var v66 int32 = (v1 + (^(v3)))
                                                v67 := GenS6{F0: -32768, F1: (v20 << ((18446744073709551615) & 63)), F2: 1}
                                                _ = len(v2)
                                                var v68 int16 = (v20 + 462)
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            } else {
                                                var v69 uint32 = 2
                                                v70 := GenS3{F0: 1}
                                                _ = v69
                                                _ = v70
                                            }
                                        }
                                    case "a":
                                        if false {
                                            v1 = (v3 / ((v3) | 1))
                                            if ((true || true) || true) {
                                                v71 := GenS7{F0: v17}
                                                delete(v2, "hello")
                                                var v72 string = "foo bar"
                                                _ = len(v2)
                                                var v73 float32 = -2.25
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            }
                                        }
                                    case "hello":
                                        v74 := GenS8{F0: (v20 % ((v20) | 1))}
                                        _ = v74
                                    default:
                                        v75 := GenS4{F0: v17, F1: 59, F2: (fi21 - fi21)}
                                        _ = v75
                                    }
                                    switch fi21 {
                                    case 4294967295:
                                        v3 -= -2147483648
                                        v76 := []int{32767, (fi19 + fi19), fi19, 9223372036854775807}
                                        var v77 int32 = 42
                                        _ = v76
                                        _ = v77
                                    case 32768:
                                        var v78 int = (fi21 - 4294967295)
                                        _ = v78
                                    }
                                    if true {
                                        switch (((("" + "hello") + "foo bar") + ("\n\t" + ((("hello" + "\n\t") + "") + ""))) + ("a" + "hello")) {
                                        case "":
                                            if true {
                                                switch (fi21 >> ((2) & 63)) {
                                                case 9223372036854775807:
                                                    v79 := []int{(fi21 - (fi21 % ((32767) | 1))), (fi19 >> ((42) & 63))}
                                                    _ = v79
                                                case 570:
                                                    switch fi19 {
                                                    case 520:
                                                        v80 := GenS6{F0: v20, F1: (-(v20)), F2: 1}
                                                        _ = v80
                                                    case -287:
                                                        var v81 uint = 2
                                                        _ = v81
                                                    case 2147483647:
                                                        switch 407 {
                                                        case 779:
                                                            fmt.Println(v17, v1, v3)
                                                        case 2147483647:
                                                            v20 -= int16(v3)
                                                        }
                                                    }
                                                case 32767:
                                                    _ = len(v2)
                                                    fmt.Println(v1, fi21)
                                                default:
                                                    var v82 uint8 = 2
                                                    _ = v82
                                                }
                                                if true {
                                                    v18 += ((float32((3.14159 - (0.1 - ((3.14159 - 1.5) * 3.14159)))) / (float32(1e10) / 1.5)) * v18)
                                                    var v83 uint32 = 597
                                                    v84 := GenS3{F0: 128}
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    if false {
                                                        fmt.Println(v3, v3)
                                                        v85 := GenS9{F0: (v20 | v20)}
                                                        _ = v85
                                                    }
                                                }
                                            } else {
                                                var v86 uint8 = 127
                                                v87 := GenS5{F0: ("" + ("a" + ("a" + "foo bar"))), F1: v17, F2: v86}
                                                v88 := v2["\n\t"]
                                                _ = v86
                                                _ = v87
                                                _ = v88
                                            }
                                        case "hello":
                                            if ((false || true) || (float64(0.1) == 0.1)) {
                                                var v89 uint16 = (v17 | 65535)
                                                _ = v89
                                            }
                                        }
                                        if false {
                                            v90 := v2["a"]
                                            fmt.Println(fi19)
                                            _ = v90
                                        } else {
                                            v17 = (v17 & v17)
                                            v20 = (-(v20))
                                            v17 = v17
                                        }
                                        v17 += (^(v17))
                                        v3 = (v3 * 1)
                                    } else {
                                        v91 := []int{}
                                        fmt.Println(v18)
                                        _ = v91
                                    }
                                } else {
                                    v92 := GenS1{F0: 255}
                                    switch "foo bar" {
                                    case "":
                                        var v93 int8 = 33
                                        _ = v93
                                    }
                                    v1 -= (v1 / ((-2147483648) | 1))
                                    _ = v92
                                }
                            case "\n\t":
                                var v94 uint32 = 2147483648
                                _ = v94
                            default:
                                v1 += (v1 + v1)
                            }
                        }
                    } else {
                        fmt.Println(fi19, v18, v18)
                        if false {
                            var v95 string = (("hello" + "") + "\n\t")
                            _ = v95
                        }
                        v20 += (^(v20))
                        if false {
                            v96 := GenS10{F0: 4294967295, F1: false, F2: (v17 & (^(v17)))}
                            var v97 int8 = -47
                            var v98 float64 = (0.1 + 1.5)
                            _ = v96
                            _ = v97
                            _ = v98
                        }
                        v1 += v1
                        fmt.Println(v3)
                        fmt.Println(v3, v18, v20)
                    }
                    var v99 uint32 = 364
                    _ = v99
                }
                if (1 != 2) {
                    v100 := GenS6{F0: 32767, F1: 444, F2: 0}
                    var v101 uint64 = 4294967295
                    if (42 <= -128) {
                        if true {
                            v18 = float32(0.0)
                        }
                    }
                    _ = v100
                    _ = v101
                }
            }
            v102 := []int{fi19}
            _ = fi21
            _ = v102
            _ = v22
        }
        var v103 int8 = 100
        _ = fi19
        _ = v103
        _ = v20
    }
    var v104 int32 = (v1 ^ (v1 << ((18446744073709551615) & 63)))
    for fi105 := 0; fi105 < 7; fi105++ {
        fmt.Println(fi105, v104, fi105)
        _ = fi105
    }
    var v106 int = -750
    if false {
        var v107 int8 = -1
        v108 := []int{}
        _ = v107
        _ = v108
    } else {
        var v109 int64 = 1
        v110 := []int{v106}
        _ = v109
        _ = v110
    }
    v111 := GenS11{F0: "foo bar", F1: 42, F2: 42}
    var v112 int8 = -1
    _ = len(v2)
    var v113 uint64 = 742
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v106:", v106)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v104
    _ = v106
    _ = v111
    _ = v112
    _ = v113
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
