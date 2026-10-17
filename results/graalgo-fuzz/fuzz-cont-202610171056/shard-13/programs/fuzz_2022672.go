package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int16; F2 int32 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint; F2 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 uint64; F1 bool; F2 bool }
type GenS11 struct { F0 uint16; F1 string; F2 int8 }
func genFn0(p0 uint16, p1 int16, p2 uint32) int {
    if (true || false) {
        var v6 bool = true
        if true {
            var v7 int64 = -1
            _ = v7
        } else {
            v8 := GenS1{F0: "a"}
            _ = v8
        }
        _ = v6
    } else {
        if true {
            var v9 uint8 = 127
            fmt.Println(p1, p1)
            _ = v9
        } else {
            if (0 >= 210) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    var v11 int8 = -43
                    var v12 int8 = (v11 & v11)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 8; fi13++ {
                    if (-2147483648 <= (-(fi13))) {
                        if false {
                            v14 := []int{(fi13 | -368)}
                            _ = v14
                        }
                        fmt.Println(p0, p1, fi13)
                        var v15 int64 = -1
                        _ = v15
                    } else {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            fmt.Println(p2)
                        case "":
                            v16 := []int{fi13, 4294967295, fi13, fi13}
                            _ = v16
                        case "foo bar":
                            var v17 uint64 = 0
                            _ = v17
                        default:
                            var v18 uint64 = 2
                            _ = v18
                        }
                        v19 := []int{fi13, int(p0)}
                        _ = v19
                    }
                    _ = fi13
                }
            } else {
                switch 798 {
                case -20:
                    var v20 int = 42
                    _ = v20
                }
                switch "hello" {
                case "\n\t":
                    switch 9223372036854775807 {
                    case 446:
                        fmt.Println(p2, p0)
                    }
                case "foo bar":
                    v21 := []int{-2147483648, 1, 42}
                    _ = v21
                case "hello":
                    for fi22 := 0; fi22 < 9; fi22++ {
                        var v23 float32 = (3.14159 / 1.5)
                        _ = fi22
                        _ = v23
                    }
                default:
                    p2 = (p2 | 2)
                }
                v24 := GenS2{F0: 92, F1: (p0 / ((42) | 1))}
                _ = v24
            }
        }
        if false {
            fmt.Println(p2, p1)
        } else {
            var v25 bool = true
            p0 = (p0 * 1)
            _ = v25
        }
    }
    switch -502 {
    case -369:
        p2 += (p2 * (p2 - p2))
    default:
        if false {
            p2 -= 42
            switch "hello" {
            case "":
                var v26 float32 = (-2.25 + 1e10)
                _ = v26
            default:
                var v27 int8 = 0
                _ = v27
            }
            fmt.Println(p0, p1)
        }
    }
    var v28 int16 = p1
    var v29 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return 32767
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2
    v3 := []string{"hello"}
    v4 := []uint32{v2}
    var v5 string = ""
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v2 = (v2 / (((^(v2))) | 1))
        if (-2147483648 > -1) {
            v2 = v2
            for fi30 := 0; fi30 < 20; fi30++ {
                if false {
                    var v31 uint32 = 1
                    _ = v31
                }
                _ = fi30
            }
            switch "" {
            case "\n\t":
                fmt.Println(v5)
            case "hello":
                fmt.Println(v2, v2)
            }
        } else {
            fmt.Println(v5, v5)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if false {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint8 = 5
                v5 = (((v5 + "foo bar") + v5) + "a")
                v2 = 42
                _ = fi32
                _ = v33
            }
            v34 := GenS2{F0: 1, F1: 2}
            _ = v34
        } else {
            for fi35 := 0; fi35 < 13; fi35++ {
                if false {
                    v2 -= v2
                }
                _ = fi35
            }
            var v36 float64 = 0.0
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 float64 = 1.5
                var v39 bool = (!false)
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        }
    }
    var v40 float64 = (0.0 / 1.0)
    v5 = v5
    for fi41 := 0; fi41 < 10; fi41++ {
        v5 = v5
        for fi42 := 0; fi42 < 17; fi42++ {
            if true {
                v43 := GenS1{F0: v5}
                fmt.Println(v40)
                _ = len(v3)
                _ = v43
            }
            if true {
                if (42 == 4294967295) {
                    switch ("\n\t" + "") {
                    case "\n\t":
                        var v44 uint16 = 60
                        _ = v44
                    default:
                        v1 = append(v1, (v2 + v2))
                    }
                } else {
                    v40 += (v40 - ((v40 - (3.14159 * v40)) + 0.1))
                    v45 := GenS3{F0: "a"}
                    switch -953 {
                    case -773:
                        if (!(true && true)) {
                            v46 := GenS4{F0: 450, F1: 42, F2: -1}
                            v40 = v40
                            switch "" {
                            case "a":
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                            case "\n\t":
                                var v47 int32 = 42
                                _ = v47
                            }
                            _ = v46
                        } else {
                            v1 = append(v1, 65535)
                        }
                    case 32767:
                        v3 = append(v3, v5)
                    default:
                        switch ((v5 + "\n\t") + v5) {
                        case "\n\t":
                            v5 = v5
                        }
                    }
                    switch "\n\t" {
                    case "":
                        v40 += v40
                    default:
                        if ((^(v2)) != (v2 ^ 42)) {
                            if ((((1e10 - 3.14159) / 1e10) + 0.1) > (1.5 * 3.14159)) {
                                var v48 bool = (18446744073709551615 <= uint64(v2))
                                switch fi42 {
                                case -2147483648:
                                    v5 = v5
                                }
                                _ = v48
                            } else {
                                if false {
                                    v2 = v2
                                    var v49 int64 = 1
                                    _ = v49
                                }
                                switch fi42 {
                                case -528:
                                    fmt.Println(fi41, v5)
                                case 32767:
                                    fmt.Println(fi42)
                                default:
                                    fmt.Println(fi42, v5)
                                    if true {
                                        v50 := GenS2{F0: 100, F1: 32767}
                                        if ("hello" < v5) {
                                            v51 := GenS1{F0: v5}
                                            var v52 int32 = 1
                                            _ = v51
                                            _ = v52
                                        } else {
                                            var v53 bool = (!((42 == 6) || true))
                                            fmt.Println(v40, v53, v40)
                                            _ = len(v3)
                                            v54 := GenS4{F0: 896, F1: -32768, F2: 383}
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = v50
                                    }
                                    if true {
                                        var v55 uint8 = 42
                                        if (593 >= 9223372036854775807) {
                                            v40 += 0.1
                                            var v56 uint16 = 402
                                            _ = v56
                                        }
                                        _ = v55
                                    } else {
                                        if true {
                                            var v57 string = v5
                                            v58 := GenS3{F0: ("a" + v57)}
                                            switch (fi41 + 127) {
                                            case -2147483648:
                                                v59 := GenS2{F0: 127, F1: 380}
                                                switch ("" + ("\n\t" + ("" + (v5 + "")))) {
                                                case "hello":
                                                    v60 := GenS5{F0: 42, F1: 2147483647}
                                                    _ = v60
                                                default:
                                                    var v61 int16 = 0
                                                    _ = v61
                                                }
                                                switch fi41 {
                                                case -1:
                                                    _ = len(v4)
                                                case 200:
                                                    if true {
                                                        var v62 string = ((v57 + "\n\t") + v57)
                                                        v63 := GenS5{F0: 0, F1: -2147483648}
                                                        fmt.Println(fi41)
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        v57 = v5
                                                        v2 += (v2 & (v2 * v2))
                                                        v64 := GenS4{F0: 48, F1: 32767, F2: 1}
                                                        switch "\n\t" {
                                                        case "hello":
                                                            fmt.Println(v2, v57)
                                                            v65 := GenS3{F0: v57}
                                                            var v66 int8 = 82
                                                            _ = v65
                                                            _ = v66
                                                        default:
                                                            var v67 bool = false
                                                            _ = v67
                                                        }
                                                        _ = v64
                                                    }
                                                case 1:
                                                    var v68 string = "\n\t"
                                                    _ = len(v4)
                                                    _ = v68
                                                default:
                                                    fmt.Println(v2, v2, fi42)
                                                }
                                                _ = v59
                                            case 9223372036854775807:
                                                var v69 int16 = 42
                                                _ = v69
                                            }
                                            _ = v57
                                            _ = v58
                                        } else {
                                            var v70 int8 = 46
                                            if (0 >= 191) {
                                                if len(v1) > 0 {
                                                    v1[0] = (^(v2))
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (true && true) {
                                                    switch (v5 + v5) {
                                                    case "a":
                                                        fmt.Println(v40)
                                                    case "":
                                                        v4 = append(v4, (v2 ^ v2))
                                                    case "foo bar":
                                                        switch genFn0(823, 1000, v2) {
                                                        case -902:
                                                            var v71 string = "a"
                                                            _ = v71
                                                        case -542:
                                                            if ((false && false) && true) {
                                                                var v72 int16 = 32767
                                                                v72 += 0
                                                                var v73 uint64 = 68
                                                                var v74 int = (fi41 | (fi41 << ((1) & 63)))
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                    default:
                                                        if true {
                                                            fmt.Println(fi42, v2, v2)
                                                            var v75 uint8 = 42
                                                            _ = v75
                                                        }
                                                    }
                                                    v5 = ("foo bar" + v5)
                                                    v1 = append(v1, (v2 % ((v2) | 1)))
                                                } else {
                                                    v76 := GenS3{F0: v5}
                                                    if (true || false) {
                                                        v40 = float64(-2.25)
                                                        v2 = v2
                                                    } else {
                                                        fmt.Println(fi41, v5, fi42)
                                                        switch v5 {
                                                        case "\n\t":
                                                            v5 = "\n\t"
                                                        case "":
                                                            if false {
                                                                v40 -= v40
                                                                var v77 int64 = 9223372036854775807
                                                                v78 := GenS6{F0: 255, F1: 2, F2: (!true)}
                                                                switch (-(fi42)) {
                                                                case 804:
                                                                    switch fi41 {
                                                                    case 731:
                                                                        if false {
                                                                            v79 := GenS7{F0: v5}
                                                                            var v80 int64 = v77
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            v5 = ""
                                                                        } else {
                                                                            if false {
                                                                                var v81 float64 = v40
                                                                                _ = v81
                                                                            }
                                                                            var v82 bool = (255 != 2)
                                                                            var v83 uint = 4294967295
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                case 32768:
                                                                    switch v5 {
                                                                    case "hello":
                                                                        v84 := GenS8{F0: v2}
                                                                        v85 := GenS9{F0: (false || false), F1: 128}
                                                                        v86 := GenS8{F0: v2}
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                    case "\n\t":
                                                                        fmt.Println(v77, v5)
                                                                    default:
                                                                        v70 -= (v70 / ((0) | 1))
                                                                    }
                                                                case 425:
                                                                    var v87 float64 = v40
                                                                    _ = v87
                                                                default:
                                                                    v88 := GenS8{F0: (v2 & 65535)}
                                                                    _ = v88
                                                                }
                                                                var v89 int = (fi41 + 32767)
                                                                v90 := GenS10{F0: 1, F1: ((v2 & (v2 & (v2 - v2))) <= v2), F2: false}
                                                                _ = v77
                                                                _ = v78
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        default:
                                                            v4 = append(v4, v2)
                                                        }
                                                    }
                                                    _ = len(v3)
                                                    _ = v76
                                                }
                                                switch "" {
                                                case "a":
                                                    _ = len(v4)
                                                }
                                                var v91 uint32 = (v2 << ((65535) & 63))
                                                _ = v91
                                            }
                                            _ = v70
                                        }
                                        v40 -= v40
                                    }
                                }
                            }
                            v1 = append(v1, 4294967295)
                        }
                    }
                    v2 -= v2
                    if true {
                        switch ("a" + v5) {
                        case "\n\t":
                            v5 = v5
                            if len(v3) > 0 {
                                v3[0] = v5
                            }
                            var v92 int64 = -9223372036854775808
                            _ = v92
                        case "hello":
                            if false {
                                v40 -= float64(float32((((0.0 / 1.5) + v40) * v40)))
                                var v93 int8 = 127
                                _ = v93
                            } else {
                                v2 = 2147483647
                                v94 := GenS8{F0: 591}
                                if false {
                                    var v95 uint64 = 1
                                    switch fi41 {
                                    case 2147483647:
                                        fmt.Println(fi41)
                                    case 150:
                                        var v96 int32 = -1
                                        _ = v96
                                    default:
                                        if false {
                                            v97 := GenS7{F0: (v5 + v5)}
                                            _ = v97
                                        }
                                    }
                                    if true {
                                        v98 := GenS11{F0: 255, F1: v5, F2: -1}
                                        var v99 int8 = 42
                                        v40 = v40
                                        _ = v98
                                        _ = v99
                                    } else {
                                        switch fi42 {
                                        case 128:
                                            if true {
                                                fmt.Println(v5, fi41, v2)
                                                v100 := GenS6{F0: 98, F1: 255, F2: (true && false)}
                                                fmt.Println(fi42)
                                                fmt.Println(v95, v95)
                                                fmt.Println(v95, fi42)
                                                v5 = "foo bar"
                                                _ = v100
                                            } else {
                                                v4 = append(v4, (v2 * (v2 & 65535)))
                                            }
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v101 uint64 = v95
                                            _ = v101
                                        }
                                        _ = len(v4)
                                    }
                                    _ = v95
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "")
                                    }
                                }
                                _ = v94
                            }
                        default:
                            var v102 uint64 = 18446744073709551615
                            _ = v102
                        }
                        if false {
                            v103 := GenS5{F0: 9223372036854775807, F1: 53}
                            if len(v1) > 0 {
                                v1[0] = 517
                            }
                            _ = v103
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = len(v1)
                        }
                    }
                    _ = v45
                }
            } else {
                v104 := GenS10{F0: 1, F1: true, F2: true}
                _ = v104
            }
            _ = fi42
        }
        v105 := GenS8{F0: (v2 % ((372) | 1))}
        _ = fi41
        _ = v105
    }
    if len(v1) > 0 {
        v1[0] = (^(v2))
    }
    if true {
        v5 = v5
    }
    v2 = 590
    v5 = (("foo bar" + v5) + ("hello" + v5))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v106 := GenS7{F0: "hello"}
    var v107 int = -511
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
