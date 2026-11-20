package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
func genFn0(p0 bool, p1 int, p2 uint) uint {
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 7; fi8++ {
        switch "\n\t" {
        case "hello":
            var v9 bool = true
            _ = v9
        case "\n\t":
            if p0 {
                v10 := GenS1{F0: 612}
                v11 := GenS1{F0: 1}
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 int = p1
                    if p0 {
                        var v14 string = ("" + ("hello" + "\n\t"))
                        switch fi8 {
                        case 31:
                            v15 := []int{0, (v13 & 1), fi8}
                            _ = v15
                        case 177:
                            var v16 bool = p0
                            _ = v16
                        case -2147483648:
                            v17 := []int{fi8, fi12, (v13 ^ (fi12 >> ((p2) & 63)))}
                            _ = v17
                        default:
                            v18 := GenS2{F0: -32768, F1: -1}
                            _ = v18
                        }
                        _ = v14
                    } else {
                        if (p0 || (v7 < (v7 ^ v7))) {
                            if (false || (65535 <= 0)) {
                                v19 := []int{}
                                _ = v19
                            } else {
                                var v20 uint32 = 552
                                var v21 string = ("foo bar" + "hello")
                                var v22 uint64 = 1
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        var v23 uint16 = 845
                        var v24 float32 = 0.1
                        _ = v23
                        _ = v24
                    }
                    p1 -= -555
                    v25 := []int{4294967295, (v13 & (^(p1)))}
                    v26 := GenS2{F0: 32767, F1: 60}
                    _ = fi12
                    _ = v13
                    _ = v25
                    _ = v26
                }
                _ = v10
                _ = v11
            } else {
                var v27 int32 = 42
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 bool = (p0 || p0)
                    var v30 uint = 255
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                switch int(v27) {
                case -9223372036854775808:
                    switch (p1 << ((p2) & 63)) {
                    case -9223372036854775808:
                        if ((810 < int16(p2)) && (848 <= 18446744073709551615)) {
                            switch -1 {
                            case 2147483647:
                                p2 = (p2 / (((p2 >> ((p2) & 63))) | 1))
                            }
                            if true {
                                fmt.Println(v27)
                            } else {
                                if (42 != -32768) {
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        var v32 uint8 = (v7 + 127)
                                        v32 += v7
                                        switch (("" + (("" + "a") + ("a" + ("\n\t" + ("\n\t" + ("" + (("hello" + "a") + "foo bar"))))))) + ("a" + "\n\t")) {
                                        case "":
                                            var v33 uint = (^(p2))
                                            _ = v33
                                        case "a":
                                            v7 -= v7
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                    p2 += (p2 / ((18446744073709551615) | 1))
                                    v34 := []int{p1, fi8}
                                    _ = v34
                                }
                                v35 := []int{515, (^(p1))}
                                _ = v35
                            }
                            v36 := GenS2{F0: -720, F1: 9223372036854775807}
                            _ = v36
                        }
                    case 4294967295:
                        p0 = false
                    case -62:
                        v37 := []int{0, p1, fi8}
                        _ = v37
                    }
                case -27:
                    v38 := []int{(p1 - 1), (fi8 | 151), 842}
                    _ = v38
                case 1:
                    for fi39 := 0; fi39 < 9; fi39++ {
                        var v40 string = ""
                        var v41 uint = 255
                        v42 := []int{}
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                default:
                    if true {
                        v43 := []int{p1, 252}
                        _ = v43
                    } else {
                        for fi44 := 0; fi44 < 6; fi44++ {
                            var v45 int16 = 42
                            var v46 int16 = 0
                            switch (fi44 & (^(fi44))) {
                            case -9223372036854775808:
                                v47 := []int{}
                                _ = v47
                            case 32768:
                                v48 := []int{fi8}
                                _ = v48
                            case 42:
                                v46 = (v46 + v45)
                            }
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        v27 = v27
                        v49 := GenS3{F0: int32(fi8), F1: 1}
                        _ = v49
                    }
                }
                _ = v27
            }
        case "":
            for fi50 := 0; fi50 < 17; fi50++ {
                fmt.Println(fi50, p2, p1)
                _ = fi50
            }
        default:
            for fi51 := 0; fi51 < 14; fi51++ {
                p0 = (32767 < -1)
                _ = fi51
            }
        }
        v52 := []int{(p1 / (((p1 * (p1 * 4294967295))) | 1)), fi8, p1}
        p0 = (0.1 == 3.14159)
        _ = fi8
        _ = v52
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 + p2)
}

func main() {
    var v1 int8 = 100
    var v2 bool = false
    v3 := []int{-2147483648}
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    var v6 uint64 = 1
    v53 := GenS4{F0: 60, F1: -82}
    if (!v2) {
        for fi54 := 0; fi54 < 11; fi54++ {
            v6 -= v6
            v55 := GenS5{F0: -114, F1: 0}
            var v56 int8 = (v1 & -114)
            _ = fi54
            _ = v55
            _ = v56
        }
        switch 42 {
        case 127:
            for fi57 := 0; fi57 < 10; fi57++ {
                for fi58 := 0; fi58 < 10; fi58++ {
                    if v2 {
                        v1 -= -51
                    } else {
                        var v59 float64 = float64(1.5)
                        _ = v59
                    }
                    _ = fi58
                }
                fmt.Println(v1, fi57)
                var v60 uint = 246
                _ = fi57
                _ = v60
            }
        case 2147483647:
            switch "" {
            case "\n\t":
                if (0 == 65535) {
                    v1 = (v1 ^ (v1 | 0))
                    fmt.Println(v6)
                }
            case "a":
                fmt.Println(v2, v2)
            }
        default:
            fmt.Println(v1)
        }
        switch 322 {
        case 1:
            var v61 float32 = (3.14159 - (1e10 - ((-2.25 + 1.5) * 0.1)))
            _ = v61
        case -1:
            for fi62 := 0; fi62 < 10; fi62++ {
                if len(v3) > 0 {
                    v3[0] = (fi62 / (((fi62 + (fi62 * (fi62 ^ 1)))) | 1))
                }
                _ = fi62
            }
        case -620:
            if (false && (false || v2)) {
                var v63 int32 = 1
                if (109 >= 127) {
                    v4[32767] = (42 != v63)
                    fmt.Println(v1)
                }
                _ = v63
            }
        default:
            v1 -= (v1 << ((0) & 63))
        }
    } else {
        if (v2 || (698 < 361)) {
            if (-25 >= -771) {
                _ = len(v3)
                var v64 string = "foo bar"
                _ = v64
            }
            var v65 float64 = 0.1
            v66 := GenS6{F0: -2147483648}
            _ = v65
            _ = v66
        } else {
            fmt.Println(v2, v2, v1)
            var v67 uint64 = v6
            var v68 uint64 = v6
            _ = v67
            _ = v68
        }
        delete(v4, -876)
        var v69 uint32 = 2
        _ = v69
    }
    fmt.Println(v1, v6)
    if true {
        switch (("" + "foo bar") + "a") {
        case "foo bar":
            v70 := v4[-714]
            _ = v70
        case "hello":
            if (128 < 127) {
                if v2 {
                    if false {
                        v71 := GenS1{F0: v6}
                        var v72 uint64 = v6
                        _ = v71
                        _ = v72
                    } else {
                        for fi73 := 0; fi73 < 16; fi73++ {
                            if true {
                                v1 = -128
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v6 += 1
                            }
                            v74 := v4[fi73]
                            _ = fi73
                            _ = v74
                        }
                        v75 := GenS3{F0: 1, F1: 42}
                        delete(v4, 2147483647)
                        _ = v75
                    }
                    if false {
                        v76 := v5[4294967295]
                        _ = v76
                    } else {
                        switch "hello" {
                        case "hello":
                            if v2 {
                                var v77 uint = 839
                                switch "" {
                                case "":
                                    for fi78 := 0; fi78 < 5; fi78++ {
                                        v79 := v4[fi78]
                                        var v80 float32 = (0.0 + 1e10)
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                default:
                                    v1 = v1
                                }
                                var v81 uint32 = 30
                                _ = v77
                                _ = v81
                            } else {
                                _ = len(v5)
                                _ = len(v4)
                                var v82 int16 = -32768
                                _ = v82
                            }
                        }
                    }
                }
                var v83 float32 = 0.0
                _ = v83
            } else {
                var v84 int64 = -1
                _ = v84
            }
        }
    } else {
        if (432 >= (v6 >> ((65535) & 63))) {
            v1 -= (v1 | v1)
            switch 4294967295 {
            case -1:
                v1 = (v1 >> ((129) & 63))
            case -9223372036854775808:
                var v85 int64 = 42
                _ = v85
            default:
                if (42 == 4) {
                    v3 = append(v3, 128)
                } else {
                    if v2 {
                        if true {
                            var v86 uint = 475
                            _ = v86
                        }
                    } else {
                        v2 = v2
                        var v87 int16 = -630
                        _ = v87
                    }
                }
            }
            if (!false) {
                v2 = v2
                if v2 {
                    for fi88 := 0; fi88 < 2; fi88++ {
                        var v89 float32 = 1e10
                        v90 := GenS4{F0: 255, F1: (v1 | v1)}
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                    switch 0 {
                    case 0:
                        v6 = 692
                    }
                } else {
                    var v91 string = "\n\t"
                    _ = v91
                }
            } else {
                v1 = (^(v1))
            }
        }
        for fi92 := 0; fi92 < 13; fi92++ {
            fmt.Println(fi92, v6)
            _ = fi92
        }
        v1 = (v1 / (((v1 - (v1 | v1))) | 1))
    }
    for fi93 := 0; fi93 < 1; fi93++ {
        v2 = (v2 && (!(v2 && (!(v2 && (float64(float32(0.1)) >= 1.5))))))
        _ = fi93
    }
    v3 = append(v3, 1)
    var v94 string = "\n\t"
    v1 -= v1
    v95 := GenS3{F0: -1, F1: 0}
    fmt.Println(v2, v1, v2)
    v96 := GenS7{F0: v2}
    v6 = (v6 ^ 4294967295)
    v97 := GenS2{F0: 334, F1: -510}
    v2 = v2
    v94 = v94
    var v98 int64 = -706
    v98 = (v98 % ((v98) | 1))
    v2 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v94
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
