package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 bool; F1 int8; F2 int }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 string; F1 uint32; F2 uint8 }
type GenS10 struct { F0 int; F1 uint64; F2 uint16 }
type GenS11 struct { F0 uint16; F1 uint16; F2 uint16 }
func genFn0(p0 uint64) uint32 {
    v4 := GenS1{F0: 1, F1: 65535, F2: "hello"}
    v5 := []int{4294967295, -420}
    p0 += (^(p0))
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int16) int8 {
    switch -2147483648 {
    case 793:
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{(^(fi6)), fi6, fi6, -1}
            _ = fi6
            _ = v7
        }
    case -513:
        v8 := []int{2147483647, 4294967295, 706}
        _ = v8
    }
    switch 127 {
    case -397:
        var v9 uint16 = 0
        _ = v9
    case 128:
        for fi10 := 0; fi10 < 3; fi10++ {
            switch int(p0) {
            case 2147483647:
                switch "" {
                case "\n\t":
                    var v11 uint32 = 785
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := []int{-1, (fi12 + fi10)}
                    p0 -= (p0 >> ((65535) & 63))
                    _ = fi12
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                switch "hello" {
                case "foo bar":
                    v15 := GenS1{F0: -128, F1: 65535, F2: "foo bar"}
                    _ = v15
                case "a":
                    fmt.Println(fi10, fi14)
                default:
                    v16 := []int{(fi10 - fi14)}
                    _ = v16
                }
                v17 := []int{}
                switch 989 {
                case 9223372036854775807:
                    fmt.Println(fi14)
                case -9223372036854775808:
                    if (255 != 0) {
                        p0 = p0
                    } else {
                        p0 = p0
                        p0 = p0
                        var v18 bool = (0 > 85)
                        if false {
                            if len(v17) > 0 {
                                v17[0] = (fi10 + 0)
                            }
                            var v19 int = int(p0)
                            v20 := GenS2{F0: 127}
                            _ = v19
                            _ = v20
                        } else {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        }
                        v18 = true
                        _ = v18
                    }
                case 32767:
                    if (!(0.0 == ((-2.25 / float64(1.5)) - 0.1))) {
                        fmt.Println(fi14)
                        if true {
                            switch fi10 {
                            case 128:
                                var v21 uint32 = 823
                                _ = v21
                            }
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            if len(v17) > 0 {
                                v17[0] = fi10
                            }
                        }
                        if ((p0 >> ((2) & 63)) != 261) {
                            switch fi14 {
                            case 42:
                                p0 = p0
                            case -9223372036854775808:
                                var v22 uint32 = 65535
                                _ = v22
                            case 127:
                                fmt.Println(fi14, p0, p0)
                                p0 = -365
                            default:
                                var v23 float32 = 3.14159
                                _ = v23
                            }
                            var v24 uint = 4294967295
                            v24 -= (v24 & 289)
                            _ = v24
                        } else {
                            var v25 float64 = 1e10
                            p0 = (p0 % ((42) | 1))
                            v26 := GenS3{F0: fi10, F1: 64, F2: (p0 % ((p0) | 1))}
                            _ = v25
                            _ = v26
                        }
                    }
                }
                v27 := GenS2{F0: -22}
                _ = fi14
                _ = v17
                _ = v27
            }
            p0 -= p0
            _ = fi10
        }
    }
    v28 := []int{-905, -2147483648}
    _ = p0
    _ = v28
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -1
    var v3 float64 = 3.14159
    var v29 string = ("foo bar" + "hello")
    delete(v1, 4294967295)
    var v30 int = -232
    var v31 int64 = 1
    v30 -= (v30 + (v30 % (((v30 << ((0) & 63))) | 1)))
    if (-2.25 >= float32(v3)) {
        v30 = 128
    } else {
        v32 := GenS4{F0: (v30 >> ((42) & 63))}
        var v33 int8 = 100
        var v34 uint = 588
        _ = v32
        _ = v33
        _ = v34
    }
    v31 -= v31
    var v35 uint16 = 269
    var v36 int = v30
    for fi37 := 0; fi37 < 12; fi37++ {
        if false {
            for fi38 := 0; fi38 < 2; fi38++ {
                v1[65535] = ((1e10 / (-2.25 / (0.1 / -2.25))) == 0.1)
                v35 = 1
                _ = fi38
            }
            _ = len(v1)
            var v39 string = "foo bar"
            _ = v39
        } else {
            v30 = (-(fi37))
            switch int(v2) {
            case 32767:
                v29 = "hello"
            case 128:
                var v40 int8 = int8(v35)
                _ = v40
            case 734:
                fmt.Println(v36, v35)
            }
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 int16 = 1
                if (v42 > v42) {
                    v30 = (-(fi41))
                } else {
                    fmt.Println(v31, v30)
                    switch v36 {
                    case 772:
                        v43 := v1[2147483648]
                        _ = v43
                    case 128:
                        v29 = ("hello" + "a")
                    case 9223372036854775807:
                        if false {
                            v44 := GenS5{F0: (v31 % (((v31 << ((367) & 63))) | 1)), F1: 18446744073709551615, F2: 59}
                            var v45 int16 = v42
                            _ = v44
                            _ = v45
                        } else {
                            v35 = (^(v35))
                        }
                    default:
                        switch fi37 {
                        case 2147483647:
                            v29 = "\n\t"
                        case 677:
                            v31 += (-(v31))
                        case -9223372036854775808:
                            v46 := []int{(fi37 | v30), int(v42)}
                            _ = v46
                        default:
                            if true {
                                if true {
                                    if true {
                                        v47 := []int{v36, int(v42)}
                                        v36 += 1
                                        v31 -= 502
                                        _ = v47
                                    }
                                    if true {
                                        var v48 uint32 = 2147483648
                                        switch "foo bar" {
                                        case "a":
                                            var v49 uint16 = (v35 << ((42) & 63))
                                            _ = v49
                                        case "\n\t":
                                            var v50 bool = true
                                            _ = v50
                                        }
                                        fmt.Println(v42, v2, v42)
                                        _ = v48
                                    } else {
                                        if false {
                                            v2 = (v2 << ((42) & 63))
                                            switch int(v2) {
                                            case 32768:
                                                v51 := []int{fi37, v30}
                                                _ = v51
                                            default:
                                                switch v29 {
                                                case "foo bar":
                                                    if false {
                                                        var v52 int16 = (v42 << ((2) & 63))
                                                        _ = v52
                                                    }
                                                case "":
                                                    if false {
                                                        if (1 != 2) {
                                                            var v53 int64 = v31
                                                            switch v29 {
                                                            case "":
                                                                v31 += v31
                                                            case "\n\t":
                                                                _ = len(v1)
                                                            }
                                                            v54 := GenS6{F0: 42, F1: (v31 < v53)}
                                                            switch v29 {
                                                            case "a":
                                                                var v55 string = v29
                                                                _ = v55
                                                            case "foo bar":
                                                                var v56 uint32 = 2
                                                                _ = v56
                                                            }
                                                            var v57 float32 = (1e10 / 0.1)
                                                            _ = v53
                                                            _ = v54
                                                            _ = v57
                                                        }
                                                        if true {
                                                            v58 := []int{fi41}
                                                            var v59 int16 = 1
                                                            var v60 float32 = 1e10
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                default:
                                                    v36 = v36
                                                }
                                            }
                                        } else {
                                            switch v29 {
                                            case "\n\t":
                                                fmt.Println(v31, v42, v42)
                                            case "a":
                                                if (0 < v35) {
                                                    var v61 bool = (!true)
                                                    var v62 float64 = 1.5
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    switch v29 {
                                                    case "":
                                                        var v63 float64 = v3
                                                        _ = v63
                                                    case "foo bar":
                                                        if (!((-128 > -128) && true)) {
                                                            v64 := GenS7{F0: (("" + "hello") <= v29), F1: -1, F2: 346}
                                                            _ = v64
                                                        } else {
                                                            switch ("foo bar" + "hello") {
                                                            case "\n\t":
                                                                v42 += v42
                                                            }
                                                            v65 := GenS4{F0: (v36 >> ((65535) & 63))}
                                                            var v66 uint32 = 2147483647
                                                            v67 := []int{(v36 ^ v36)}
                                                            if true {
                                                                v68 := GenS5{F0: (v31 << ((0) & 63)), F1: 42, F2: 212}
                                                                var v69 int = (v30 << ((65535) & 63))
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    default:
                                                        var v70 float64 = v3
                                                        _ = v70
                                                    }
                                                    switch int(v35) {
                                                    case 4294967295:
                                                        _ = len(v1)
                                                    }
                                                    var v71 int32 = v2
                                                    _ = v71
                                                }
                                            }
                                            fmt.Println(v29, v3, v42)
                                            fmt.Println(v35, v42, v3)
                                        }
                                        v72 := []int{fi37, -1}
                                        _ = v72
                                    }
                                    v73 := GenS8{F0: 127, F1: (0.0 >= (3.14159 - (float32(v3) - 0.0)))}
                                    _ = v73
                                }
                                v42 += 42
                            }
                        }
                    }
                    v30 -= (v30 - v30)
                    v29 = (v29 + (v29 + "a"))
                    v74 := []int{}
                    _ = v74
                }
                _ = fi41
                _ = v42
            }
        }
        _ = fi37
    }
    switch (v36 >> ((2) & 63)) {
    case 121:
        v75 := GenS9{F0: (v29 + v29), F1: genFn0(1), F2: 20}
        _ = v75
    case 128:
        switch 42 {
        case 0:
            fmt.Println(v30, v35)
        }
    case 4294967295:
        v36 -= (^(v30))
    }
    if (!((v35 | v35) < 2)) {
        v76 := GenS4{F0: v30}
        _ = v76
    } else {
        for fi77 := 0; fi77 < 6; fi77++ {
            v29 = ""
            _ = fi77
        }
    }
    v30 += v36
    switch (v30 | (v36 >> ((615) & 63))) {
    case -351:
        v78 := []int{42, (v30 * v36), int(v35), v36}
        _ = v78
    case 42:
        for fi79 := 0; fi79 < 10; fi79++ {
            for fi80 := 0; fi80 < 15; fi80++ {
                switch v36 {
                case 127:
                    _ = len(v1)
                case 124:
                    var v81 int64 = 1
                    fmt.Println(v36, v36)
                    _ = v81
                case -9223372036854775808:
                    switch "hello" {
                    case "foo bar":
                        _ = len(v1)
                    case "\n\t":
                        if (42 > v31) {
                            if ((!true) && (genFn0(1) != 970)) {
                                if false {
                                    fmt.Println(fi80)
                                    if false {
                                        fmt.Println(v36, fi80)
                                        fmt.Println(fi80, v2, v31)
                                    } else {
                                        var v82 bool = (679 != 1)
                                        var v83 uint16 = 0
                                        v84 := GenS10{F0: v36, F1: 17, F2: (v83 ^ v83)}
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                } else {
                                    var v85 uint = 215
                                    v31 = v31
                                    _ = v85
                                }
                                switch 70 {
                                case 128:
                                    var v86 int8 = 127
                                    v87 := GenS11{F0: 32767, F1: (v35 + v35), F2: (v35 / ((65535) | 1))}
                                    _ = v86
                                    _ = v87
                                case 32767:
                                    var v88 uint = 139
                                    _ = v88
                                }
                                v36 -= (-(fi80))
                            } else {
                                if (!(!false)) {
                                    v36 = (v30 + (^(fi79)))
                                }
                                var v89 uint8 = 42
                                v1[65535] = false
                                if true {
                                    fmt.Println(v31)
                                }
                                _ = v89
                            }
                            if ((3.14159 - -2.25) == 0.0) {
                                var v90 int32 = (v2 * v2)
                                _ = v90
                            } else {
                                var v91 uint8 = 42
                                _ = v91
                            }
                        } else {
                            v92 := GenS2{F0: 42}
                            if false {
                                v93 := []int{-1, int(v31), 9223372036854775807}
                                _ = v93
                            }
                            v94 := GenS2{F0: 1}
                            v3 = float64(float32(v3))
                            _ = v92
                            _ = v94
                        }
                    case "":
                        if false {
                            fmt.Println(v2)
                        }
                    default:
                        if true {
                            v36 += fi79
                            v95 := []int{32768, 664, 1, 127}
                            if true {
                                var v96 int8 = genFn1(1)
                                v36 = (^(fi79))
                                switch -2147483648 {
                                case 127:
                                    var v97 int32 = -822
                                    _ = v97
                                case 32768:
                                    v98 := GenS8{F0: 0, F1: (true || false)}
                                    _ = v98
                                default:
                                    v31 = (v31 & 1)
                                }
                                _ = v96
                            } else {
                                var v99 uint = 18446744073709551615
                                _ = v99
                            }
                            _ = v95
                        }
                    }
                }
                var v100 int = (v30 << ((42) & 63))
                var v101 int32 = (-(v2))
                if (((((3.14159 + 0.0) * float32(v3)) - (1e10 - 0.0)) + 1.5) <= 0.0) {
                    fmt.Println(v101, v100)
                }
                v101 -= (v2 + (v2 * (v2 ^ v2)))
                _ = fi80
                _ = v100
                _ = v101
            }
            _ = fi79
        }
    case 32768:
        var v102 float32 = 1e10
        _ = v102
    default:
        v30 = (^(v30))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
}
