package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int64; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint32; F1 uint16; F2 string }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    var v3 uint32 = 654
    var v4 string = ((("foo bar" + "\n\t") + "\n\t") + (("a" + (((("" + "a") + "") + ("a" + "foo bar")) + "")) + ""))
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    v1 -= 1
    switch 1 {
    case -659:
        var v7 float32 = 3.14159
        _ = v7
    }
    v8 := GenS1{F0: 0}
    var v9 int8 = -128
    v9 += 42
    if true {
        var v10 uint16 = 206
        switch ("foo bar" + ("hello" + v4)) {
        case "\n\t":
            if (false || true) {
                v11 := []int{338, 128}
                for fi12 := 0; fi12 < 4; fi12++ {
                    if (1e10 > 0.0) {
                        var v13 int16 = -956
                        _ = v13
                    }
                    var v14 int64 = -1
                    fmt.Println(fi12, v1, v3)
                    _ = fi12
                    _ = v14
                }
                _ = v11
            } else {
                if false {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        switch (^(fi15)) {
                        case 351:
                            fmt.Println(v9)
                        case -994:
                            v16 := []int{(fi15 - fi15), -653}
                            _ = v16
                        default:
                            switch (fi15 << ((4294967295) & 63)) {
                            case 0:
                                var v17 uint16 = (v10 << ((2) & 63))
                                _ = v17
                            case 127:
                                if ((-912 != 1) && true) {
                                    var v18 uint = 65535
                                    fmt.Println(v9)
                                    _ = v18
                                }
                            }
                        }
                        fmt.Println(v9, fi15)
                        var v19 float64 = -2.25
                        _ = fi15
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if false {
                            var v21 bool = ((!(946 > 2)) || ((!(!false)) || (255 == v10)))
                            switch "foo bar" {
                            case "a":
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    var v23 uint16 = (v10 - v10)
                                    delete(v5, 549)
                                    if v21 {
                                        if v21 {
                                            var v24 uint8 = v1
                                            var v25 uint32 = v3
                                            var v26 uint64 = 42
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        } else {
                                            if v21 {
                                                if (float32((1e10 / (((0.0 + float64(-2.25)) - 0.1) * (-2.25 - 3.14159)))) < (3.14159 * 1.5)) {
                                                    v27 := []int{}
                                                    var v28 float64 = 1.5
                                                    var v29 uint64 = 4294967295
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    switch v4 {
                                                    case "":
                                                        switch v4 {
                                                        case "foo bar":
                                                            v30 := GenS2{F0: -2147483648, F1: -2147483648}
                                                            _ = v30
                                                        case "\n\t":
                                                            v1 += v1
                                                        case "":
                                                            v21 = (0 <= 2147483647)
                                                        }
                                                    case "hello":
                                                        var v31 float64 = 1.5
                                                        _ = v31
                                                    }
                                                    var v32 int32 = 2147483647
                                                    if (!((false && ((v9 * v9) != (v9 ^ (v9 << ((1) & 63))))) || (1 != (fi20 & fi22)))) {
                                                        var v33 int32 = (v32 * (v32 << ((99) & 63)))
                                                        _ = v33
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        var v34 bool = (255 != 4294967295)
                                                        _ = v34
                                                    }
                                                    var v35 int8 = int8(fi20)
                                                    fmt.Println(v32)
                                                    _ = v32
                                                    _ = v35
                                                }
                                                fmt.Println(v21)
                                            } else {
                                                v36 := GenS1{F0: 41}
                                                v10 += (v10 & (v23 << ((18446744073709551615) & 63)))
                                                v10 = uint16(v3)
                                                _ = v36
                                            }
                                        }
                                        v37 := GenS2{F0: 0, F1: 0}
                                        if (818 < 1) {
                                            v23 -= v10
                                        } else {
                                            var v38 uint8 = (v1 / (((v1 - (^(v1)))) | 1))
                                            if ((((v1 << ((4294967295) & 63)) == v38) || false) && v21) {
                                                if v21 {
                                                    if v21 {
                                                        v39 := GenS2{F0: 0, F1: int32(v1)}
                                                        v38 = (v38 - 2)
                                                        _ = v39
                                                    }
                                                    fmt.Println(v1, v23, v10)
                                                } else {
                                                    switch v4 {
                                                    case "a":
                                                        v3 += v3
                                                    case "foo bar":
                                                        v21 = false
                                                    case "hello":
                                                        v4 = v4
                                                    default:
                                                        switch v4 {
                                                        case "a":
                                                            if v21 {
                                                                var v40 float32 = ((1.5 * 3.14159) + (1e10 + (float32(3.14159) - (1.5 * 0.1))))
                                                                v38 += 80
                                                                v41 := GenS1{F0: 0}
                                                                if true {
                                                                    if true {
                                                                        var v42 int8 = v9
                                                                        v1 += (v38 ^ 42)
                                                                        v43 := GenS3{F0: (v42 ^ -128)}
                                                                        v9 = v42
                                                                        _ = v42
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 uint16 = 591
                                                                        v40 -= v40
                                                                        if true {
                                                                            switch v4 {
                                                                            case "a":
                                                                                v10 -= v23
                                                                                v44 = 84
                                                                                v45 := v6[v4]
                                                                                _ = v45
                                                                            default:
                                                                                v4 = v4
                                                                            }
                                                                        }
                                                                        _ = v44
                                                                    }
                                                                    v21 = (!((-80 != 42) && ("hello" > v4)))
                                                                    fmt.Println(v21)
                                                                    var v46 string = (v4 + (v4 + "foo bar"))
                                                                    _ = v46
                                                                } else {
                                                                    v47 := GenS3{F0: (v9 ^ v9)}
                                                                    _ = v47
                                                                }
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        case "":
                                                            var v48 int8 = (v9 % ((v9) | 1))
                                                            _ = v48
                                                        case "\n\t":
                                                            if v21 {
                                                                v6[""] = 42
                                                                v38 = (v1 - (v38 - (v1 << ((2) & 63))))
                                                            }
                                                        default:
                                                            var v49 int16 = -938
                                                            _ = v49
                                                        }
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v3 = (^(v3))
                                                        case "\n\t":
                                                            switch 158 {
                                                            case 127:
                                                                if v21 {
                                                                    if v21 {
                                                                        var v50 uint8 = v38
                                                                        _ = v50
                                                                    } else {
                                                                        v4 = ("" + (v4 + (("\n\t" + (v4 + v4)) + (v4 + v4))))
                                                                    }
                                                                    if false {
                                                                        v3 = 0
                                                                        _ = len(v5)
                                                                    }
                                                                } else {
                                                                    fmt.Println(v3, v4)
                                                                    v51 := []int{4294967295, (^(fi22)), (fi22 ^ 1), fi22}
                                                                    _ = v51
                                                                }
                                                            case 4294967295:
                                                                var v52 int64 = 0
                                                                _ = v52
                                                            case 969:
                                                                fmt.Println(v38)
                                                            default:
                                                                if v21 {
                                                                    v1 = 23
                                                                    if v21 {
                                                                        v53 := GenS1{F0: 42}
                                                                        _ = v53
                                                                    }
                                                                }
                                                                fmt.Println(v21)
                                                                switch (fi22 / (((^(fi20))) | 1)) {
                                                                case -1:
                                                                    var v54 int = 127
                                                                    _ = v54
                                                                default:
                                                                    v10 += uint16(v1)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    switch 32768 {
                                                    case -9223372036854775808:
                                                        fmt.Println(v4)
                                                        var v55 int32 = -409
                                                        if (false && (!((v21 && ((v3 ^ v3) != uint32(v1))) && (float32(-2.25) <= 0.1)))) {
                                                            v38 += v1
                                                            switch fi20 {
                                                            case -9223372036854775808:
                                                                v4 = v4
                                                            case -1:
                                                                v56 := []int{799, 32768, (fi20 ^ -9223372036854775808)}
                                                                _ = v56
                                                            default:
                                                                v1 = 0
                                                            }
                                                        } else {
                                                            fmt.Println(v9, v1)
                                                            fmt.Println(v4)
                                                        }
                                                        _ = v55
                                                    case 4294967295:
                                                        v57 := []int{fi22, 4294967295, fi20, (fi22 | (fi22 % ((fi22) | 1)))}
                                                        _ = v57
                                                    default:
                                                        if v21 {
                                                            v58 := GenS1{F0: 2147483647}
                                                            _ = v58
                                                        } else {
                                                            if (v9 > v9) {
                                                                v59 := []int{(fi22 ^ (fi22 ^ 2147483647)), 128, (fi22 & 113)}
                                                                v1 = uint8(v10)
                                                                var v60 uint16 = 255
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            if true {
                                                                v38 -= v38
                                                                v4 = ("hello" + v4)
                                                            }
                                                            var v61 uint8 = 42
                                                            _ = v61
                                                        }
                                                    }
                                                }
                                                if false {
                                                    v62 := GenS2{F0: -46, F1: -2147483648}
                                                    _ = v62
                                                }
                                            }
                                            _ = v38
                                        }
                                        _ = v37
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                            case "":
                                fmt.Println(v10, v21)
                            default:
                                for fi63 := 0; fi63 < 13; fi63++ {
                                    var v64 uint = 4294967295
                                    var v65 string = "hello"
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                }
                            }
                            _ = v21
                        } else {
                            var v66 float64 = 0.0
                            var v67 uint64 = 232
                            _ = v66
                            _ = v67
                        }
                        var v68 uint32 = v3
                        v9 -= v9
                        _ = fi20
                        _ = v68
                    }
                } else {
                    var v69 uint8 = (^(v1))
                    _ = v69
                }
                fmt.Println(v1, v4)
            }
        case "hello":
            v70 := []int{-1}
            _ = v70
        case "a":
            v3 -= v3
        }
        v71 := []int{-132}
        _ = v10
        _ = v71
    } else {
        fmt.Println(v4)
        v1 = v1
        if (32767 <= 65535) {
            v72 := GenS2{F0: 412, F1: 0}
            v2[983] = v3
            v73 := []int{-1}
            _ = v72
            _ = v73
        }
    }
    v4 = "hello"
    v74 := GenS4{F0: -1, F1: (v9 | v9)}
    if (false || (false || ((v9 >> ((152) & 63)) < 100))) {
        switch v4 {
        case "\n\t":
            for fi75 := 0; fi75 < 16; fi75++ {
                if false {
                    if (!true) {
                        for fi76 := 0; fi76 < 13; fi76++ {
                            v3 += (v3 ^ (v3 % ((v3) | 1)))
                            switch v4 {
                            case "foo bar":
                                if (2147483648 != 285) {
                                    if (((v1 << ((782) & 63)) > v1) || (32767 > 1)) {
                                        v9 = (v9 & (v9 << ((4294967295) & 63)))
                                        v77 := []int{fi76, -9223372036854775808}
                                        if ((v9 / (((v9 - (v9 | -46))) | 1)) < v9) {
                                            var v78 int64 = 9223372036854775807
                                            v3 -= v3
                                            v4 = ""
                                            _ = v78
                                        }
                                        _ = v77
                                    } else {
                                        v79 := GenS4{F0: -632, F1: 125}
                                        if false {
                                            fmt.Println(v1)
                                            var v80 string = "\n\t"
                                            var v81 float64 = 0.1
                                            _ = v80
                                            _ = v81
                                        }
                                        var v82 int = fi75
                                        _ = v79
                                        _ = v82
                                    }
                                    v83 := GenS5{F0: v3, F1: 581}
                                    _ = v83
                                } else {
                                    switch -229 {
                                    case 4294967295:
                                        v84 := GenS5{F0: (v3 - 211), F1: 1}
                                        _ = v84
                                    case 42:
                                        v85 := GenS6{F0: 0, F1: false, F2: 128}
                                        if ((v4 + (((v4 + "") + v4) + v4)) != "a") {
                                            if false {
                                                v86 := GenS1{F0: 639}
                                                fmt.Println(v4)
                                                v4 = "a"
                                                v87 := GenS7{F0: (v3 ^ v3), F1: 255, F2: v4}
                                                _ = v86
                                                _ = v87
                                            }
                                            fmt.Println(v4, v1, v1)
                                        }
                                        _ = v85
                                    }
                                    _ = len(v5)
                                }
                            case "":
                                v6[(("foo bar" + v4) + v4)] = v3
                            case "hello":
                                if ((v4 + v4) == "foo bar") {
                                    var v88 int32 = -2147483648
                                    var v89 string = "\n\t"
                                    _ = v88
                                    _ = v89
                                } else {
                                    if ((v3 - v3) == v3) {
                                        fmt.Println(v4, fi76, fi76)
                                        delete(v2, (^(fi76)))
                                        v90 := []int{fi76, fi75}
                                        var v91 int32 = -865
                                        var v92 uint64 = 18446744073709551615
                                        switch v4 {
                                        case "foo bar":
                                            if (false || true) {
                                                var v93 uint64 = 18446744073709551615
                                                _ = v93
                                            }
                                        case "\n\t":
                                            v91 -= (-(v91))
                                        default:
                                            v94 := GenS6{F0: 0, F1: false, F2: 2}
                                            _ = v94
                                        }
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    }
                                    v4 = "hello"
                                }
                            }
                            v95 := []int{(fi75 % (((-(fi76))) | 1)), fi75, fi75, fi75}
                            _ = fi76
                            _ = v95
                        }
                        v96 := []int{}
                        if true {
                            delete(v5, (fi75 | fi75))
                            var v97 uint64 = 898
                            _ = len(v96)
                            _ = v97
                        } else {
                            v1 = (v1 << ((42) & 63))
                        }
                        _ = v96
                    }
                    if ((1000 <= 42) || true) {
                        v4 = ((v4 + v4) + v4)
                    } else {
                        v3 += uint32(v1)
                        _ = len(v5)
                    }
                }
                _ = fi75
            }
        case "hello":
            for fi98 := 0; fi98 < 9; fi98++ {
                v3 -= 1
                for fi99 := 0; fi99 < 7; fi99++ {
                    v9 -= int8(v3)
                    if (2 < 992) {
                        v9 = -71
                    } else {
                        v1 = 95
                    }
                    fmt.Println(v9)
                    if (((1.5 + 1.5) >= 0.0) || (!false)) {
                        fmt.Println(v1, v3)
                        fmt.Println(fi99)
                        v100 := v5[fi98]
                        _ = v100
                    } else {
                        v101 := GenS4{F0: 1000, F1: 127}
                        _ = v101
                    }
                    delete(v6, v4)
                    v102 := []int{42, (fi99 >> ((1) & 63)), int(v9)}
                    _ = fi99
                    _ = v102
                }
                _ = fi98
            }
        default:
            _ = len(v5)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v8
    _ = v9
}
