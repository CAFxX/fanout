package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int; F1 uint64; F2 int8 }
type GenS9 struct { F0 int8; F1 int64 }

func main() {
    var v1 bool = ((32768 >= 107) || false)
    var v2 int = 32768
    var v3 string = ("" + "hello")
    v1 = ((1 == 956) && (-2147483648 != -109))
    var v4 float32 = -2.25
    if v1 {
        if (42 >= 280) {
            fmt.Println(v1, v3, v3)
        } else {
            v5 := []int{(v2 ^ v2), 9223372036854775807, (v2 + -2147483648), (v2 | v2)}
            v1 = (18446744073709551615 <= 69)
            _ = v5
        }
        v6 := GenS1{F0: 983, F1: 1, F2: 255}
        _ = v6
    }
    v7 := GenS2{F0: -459, F1: v3, F2: 42}
    var v8 int8 = 100
    v9 := GenS1{F0: 1, F1: 1, F2: 0}
    var v10 int16 = 42
    var v11 int8 = (v8 & v8)
    v12 := GenS2{F0: 42, F1: ("\n\t" + ""), F2: 42}
    switch "foo bar" {
    case "a":
        for fi13 := 0; fi13 < 6; fi13++ {
            if v1 {
                if (v1 || true) {
                    v14 := []int{v2, fi13, fi13}
                    var v15 float64 = float64(float32(0.1))
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = v14
                    _ = v15
                } else {
                    v3 = "a"
                }
                if v1 {
                    v16 := GenS1{F0: 42, F1: 0, F2: 997}
                    v17 := GenS3{F0: 139, F1: (65535 > 679), F2: 1}
                    for fi18 := 0; fi18 < 6; fi18++ {
                        if ((!v1) || false) {
                            v19 := GenS2{F0: 9223372036854775807, F1: "a", F2: 2147483647}
                            fmt.Println(v4, v4, v4)
                            if ((v10 ^ v10) <= v10) {
                                v20 := []int{-88, -429, (v2 - (-(fi13))), fi13}
                                _ = v20
                            } else {
                                fmt.Println(v11, v3)
                                v3 = v3
                            }
                            _ = v19
                        } else {
                            v21 := []int{}
                            var v22 uint16 = 65535
                            fmt.Println(v10, v3, v4)
                            v23 := GenS1{F0: 2, F1: 42, F2: 761}
                            if v1 {
                                fmt.Println(fi13, v10, v1)
                                var v24 string = "a"
                                _ = v24
                            } else {
                                fmt.Println(v2, v8)
                                v25 := GenS1{F0: 38, F1: 9223372036854775807, F2: (^(v22))}
                                var v26 uint8 = 42
                                _ = v25
                                _ = v26
                            }
                            if len(v21) > 0 {
                                v21[0] = (v2 * (-(v2)))
                            }
                            v27 := GenS2{F0: -1, F1: ("foo bar" + "a"), F2: 24}
                            if (!(42 > 42)) {
                                if (426 != 0) {
                                    var v28 int8 = -62
                                    _ = v28
                                } else {
                                    var v29 float32 = (v4 * (3.14159 * v4))
                                    switch (fi18 + v2) {
                                    case 1:
                                        if false {
                                            if len(v21) > 0 {
                                                v21[0] = (v2 | (v2 << ((414) & 63)))
                                            }
                                        } else {
                                            v11 -= v8
                                        }
                                    case 2147483647:
                                        if v1 {
                                            if (!v1) {
                                                v30 := GenS4{F0: 0}
                                                fmt.Println(v3)
                                                v31 := GenS3{F0: 42, F1: (2147483647 <= 0), F2: 42}
                                                _ = v30
                                                _ = v31
                                            } else {
                                                if v1 {
                                                    var v32 uint = 18446744073709551615
                                                    _ = v32
                                                }
                                            }
                                            v21 = append(v21, -454)
                                        } else {
                                            var v33 uint8 = 44
                                            _ = v33
                                        }
                                    case -607:
                                        fmt.Println(v1, v8, fi18)
                                    }
                                    _ = v29
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v2, v11)
                                case "":
                                    v34 := GenS5{F0: 457, F1: 18446744073709551615, F2: v3}
                                    switch v2 {
                                    case 4294967295:
                                        if true {
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(v10)
                                                fmt.Println(v11, v1, fi18)
                                                var v35 string = ""
                                                _ = v35
                                            }
                                        } else {
                                            var v36 uint64 = uint64(v22)
                                            if len(v21) > 0 {
                                                v21[0] = 42
                                            }
                                            _ = v36
                                        }
                                    default:
                                        fmt.Println(v4)
                                    }
                                    if true {
                                        switch "" {
                                        case "":
                                            fmt.Println(fi13, v2)
                                            fmt.Println(v8, fi18)
                                        case "foo bar":
                                            if ((!v1) && false) {
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                            } else {
                                                fmt.Println(v1, v10, v3)
                                            }
                                        default:
                                            v10 -= (v10 - -188)
                                        }
                                    }
                                    _ = v34
                                }
                            } else {
                                _ = len(v21)
                                if (1 == 591) {
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    var v37 uint = 1
                                    _ = v37
                                } else {
                                    v2 += (fi13 & 127)
                                    v11 += 0
                                    v38 := GenS4{F0: 2147483647}
                                    _ = v38
                                }
                            }
                            v10 -= v10
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v27
                        }
                        _ = fi18
                    }
                    _ = v16
                    _ = v17
                }
                var v39 uint = 2
                _ = v39
            } else {
                v40 := GenS6{F0: -1}
                v41 := []int{4294967295, (fi13 & fi13), v2, fi13}
                _ = v40
                _ = v41
            }
            _ = fi13
        }
    default:
        v42 := GenS2{F0: -258, F1: v3, F2: -2147483648}
        _ = v42
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := []int{v2}
        v8 = v11
        _ = fi43
        _ = v44
    }
    if (!false) {
        for fi45 := 0; fi45 < 9; fi45++ {
            v8 -= v11
            _ = fi45
        }
    } else {
        if ("" > (("foo bar" + "") + "hello")) {
            if v1 {
                switch v3 {
                case "hello":
                    v8 += v11
                case "foo bar":
                    v2 += 1
                default:
                    if v1 {
                        switch v3 {
                        case "hello":
                            v46 := []int{}
                            _ = v46
                        case "":
                            v8 = -1
                        case "foo bar":
                            v47 := []int{}
                            _ = v47
                        }
                        v4 -= 1.5
                    }
                }
                switch v3 {
                case "foo bar":
                    v8 = (^(v8))
                default:
                    if true {
                        for fi48 := 0; fi48 < 8; fi48++ {
                            v49 := GenS7{F0: v1, F1: (("\n\t" + ((("" + (v3 + v3)) + v3) + ("foo bar" + v3))) + v3)}
                            switch (fi48 >> ((2) & 63)) {
                            case 1:
                                switch fi48 {
                                case -1:
                                    v50 := []int{}
                                    _ = v50
                                case 0:
                                    var v51 float32 = v4
                                    _ = v51
                                case 128:
                                    v8 -= v8
                                }
                            case 32768:
                                if (1e10 != 0.0) {
                                    fmt.Println(v11, v11)
                                    v52 := []int{}
                                    _ = v52
                                }
                            case 2147483647:
                                var v53 uint = 0
                                _ = v53
                            }
                            _ = fi48
                            _ = v49
                        }
                        var v54 uint64 = 1
                        v55 := []int{v2, -2147483648, v2, v2}
                        _ = v54
                        _ = v55
                    } else {
                        if false {
                            switch "" {
                            case "hello":
                                for fi56 := 0; fi56 < 11; fi56++ {
                                    for fi57 := 0; fi57 < 3; fi57++ {
                                        v58 := []int{(fi56 ^ 1), v2, (fi56 * (^(fi57))), fi56}
                                        _ = fi57
                                        _ = v58
                                    }
                                    switch "" {
                                    case "hello":
                                        var v59 uint = 0
                                        _ = v59
                                    case "\n\t":
                                        v2 = (fi56 << ((891) & 63))
                                    case "foo bar":
                                        fmt.Println(v4)
                                    }
                                    _ = fi56
                                }
                            case "foo bar":
                                if v1 {
                                    v60 := []int{v2, v2}
                                    _ = v60
                                } else {
                                    if v1 {
                                        v61 := []int{v2, (v2 * (v2 << ((0) & 63)))}
                                        var v62 int8 = -24
                                        _ = v61
                                        _ = v62
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    v3 = "a"
                                }
                            case "":
                                fmt.Println(v2)
                            }
                            v3 = v3
                            var v63 int64 = -250
                            _ = v63
                        } else {
                            fmt.Println(v8, v11)
                            v8 -= (-(v11))
                        }
                        var v64 uint16 = 42
                        _ = v64
                    }
                }
            }
        } else {
            v65 := GenS8{F0: -1, F1: 729, F2: (v8 * v8)}
            v66 := GenS6{F0: (v10 >> ((0) & 63))}
            v11 -= -1
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 15; fi67++ {
            var v68 uint32 = 2147483648
            if v1 {
                if false {
                    fmt.Println(v11)
                    switch v3 {
                    case "foo bar":
                        var v69 string = v3
                        _ = v69
                    case "a":
                        if (v1 || v1) {
                            switch (v2 >> ((724) & 63)) {
                            case 4294967295:
                                var v70 uint8 = 153
                                _ = v70
                            case 128:
                                var v71 int8 = -128
                                _ = v71
                            case 9223372036854775807:
                                v68 += 139
                            default:
                                for fi72 := 0; fi72 < 3; fi72++ {
                                    v8 += v8
                                    if (128 == 2) {
                                        if (v2 >= (fi67 - (-(fi72)))) {
                                            switch fi72 {
                                            case -9223372036854775808:
                                                v73 := GenS8{F0: fi72, F1: 4294967295, F2: 1}
                                                _ = v73
                                            case 9223372036854775807:
                                                v74 := []int{}
                                                _ = v74
                                            case -2147483648:
                                                var v75 int32 = -2147483648
                                                _ = v75
                                            }
                                            var v76 float64 = -2.25
                                            v8 -= v11
                                            if ((fi72 ^ fi72) >= (fi72 + (-(fi67)))) {
                                                fmt.Println(fi67, v11, v3)
                                            }
                                            _ = v76
                                        }
                                        if v1 {
                                            v1 = false
                                        } else {
                                            var v77 string = v3
                                            if v1 {
                                                v11 = (v8 << ((1) & 63))
                                                var v78 int64 = 1
                                                v79 := GenS9{F0: v11, F1: (v78 - v78)}
                                                _ = v78
                                                _ = v79
                                            } else {
                                                fmt.Println(v68)
                                                v80 := GenS1{F0: 1, F1: -1, F2: 32767}
                                                switch int(v68) {
                                                case -803:
                                                    var v81 int8 = 127
                                                    _ = v81
                                                default:
                                                    fmt.Println(v10, fi72)
                                                }
                                                _ = v80
                                            }
                                            v82 := []int{fi72}
                                            _ = v77
                                            _ = v82
                                        }
                                    }
                                    if v1 {
                                        switch v3 {
                                        case "foo bar":
                                            var v83 uint32 = 1
                                            _ = v83
                                        }
                                    }
                                    _ = fi72
                                }
                            }
                            switch (v2 >> ((2) & 63)) {
                            case 2147483647:
                                switch (v2 - (fi67 << ((562) & 63))) {
                                case -20:
                                    v84 := []int{v2, fi67, v2}
                                    _ = v84
                                case 776:
                                    var v85 bool = false
                                    _ = v85
                                }
                            }
                            if v1 {
                                v86 := GenS8{F0: v2, F1: 724, F2: 127}
                                _ = v86
                            } else {
                                v3 = "\n\t"
                                var v87 int8 = 127
                                _ = v87
                            }
                        }
                    case "":
                        v8 = -128
                    }
                }
            }
            _ = fi67
            _ = v68
        }
        for fi88 := 0; fi88 < 7; fi88++ {
            fmt.Println(v3, v11, v8)
            v89 := []int{360}
            _ = fi88
            _ = v89
        }
    }
    fmt.Println(v10)
    v90 := GenS7{F0: v1, F1: v3}
    v91 := []int{}
    switch 4294967295 {
    case 127:
        v8 = (v8 | (^(v8)))
    case 9223372036854775807:
        if (v11 >= 42) {
            if true {
                var v92 uint16 = 763
                fmt.Println(v3)
                var v93 string = v3
                _ = v92
                _ = v93
            }
            v3 = ("" + "hello")
        } else {
            var v94 float32 = -2.25
            _ = v94
        }
    case -2147483648:
        var v95 int = v2
        _ = v95
    }
    var v96 float64 = ((-2.25 / (3.14159 + 1e10)) / 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v91:", len(v91), v91)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
    _ = v90
    _ = v91
    _ = v96
}
