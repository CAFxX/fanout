package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 uint32; F1 int32 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint8, p2 uint64) uint {
    var v7 uint16 = 631
    var v8 uint32 = 578
    v9 := GenS1{F0: (p1 & 146), F1: -330}
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: (^(v8)), F1: "hello"}
        p2 = (^(p2))
        if true {
            if (true || (false && (!(32767 <= 32767)))) {
                v7 = v7
                for fi12 := 0; fi12 < 11; fi12++ {
                    v7 = 255
                    _ = fi12
                }
                switch ("a" + "foo bar") {
                case "foo bar":
                    var v13 int32 = -627
                    _ = v13
                }
            } else {
                v14 := GenS3{F0: true, F1: 42, F2: (!false)}
                for fi15 := 0; fi15 < 3; fi15++ {
                    if (!false) {
                        v16 := GenS1{F0: (^(p1)), F1: -400}
                        v8 += 695
                        _ = v16
                    } else {
                        var v17 int16 = -1
                        fmt.Println(v8)
                        if ((!false) || (0.1 > 0.1)) {
                            v8 -= v8
                            switch 32768 {
                            case 128:
                                v8 -= 42
                            case 42:
                                switch (-(fi10)) {
                                case 9223372036854775807:
                                    if true {
                                        var v18 int64 = -1
                                        var v19 string = "a"
                                        _ = v18
                                        _ = v19
                                    }
                                case -33:
                                    v20 := GenS1{F0: (p1 - 128), F1: -9223372036854775808}
                                    _ = v20
                                default:
                                    var v21 float64 = 1e10
                                    _ = v21
                                }
                            }
                        } else {
                            v22 := []int{(fi10 << ((617) & 63)), fi10, fi10, (fi15 * fi10)}
                            var v23 string = "foo bar"
                            _ = v22
                            _ = v23
                        }
                        if false {
                            v24 := []int{}
                            v7 -= uint16(p2)
                            _ = v24
                        } else {
                            switch "" {
                            case "\n\t":
                                var v25 int16 = v17
                                _ = v25
                            case "":
                                var v26 uint32 = uint32(p0)
                                _ = v26
                            case "hello":
                                var v27 uint64 = p2
                                _ = v27
                            }
                            v17 += v17
                            var v28 uint = (p0 >> ((p0) & 63))
                            v29 := []int{(fi15 << (((v28 + (p0 >> ((uint(p2)) & 63)))) & 63)), (fi15 * int(v7))}
                            fmt.Println(p0)
                            _ = v28
                            _ = v29
                        }
                        p0 -= (p0 << (((p0 << (((p0 & p0)) & 63))) & 63))
                        if true {
                            if true {
                                if (1e10 != 3.14159) {
                                    if true {
                                        if true {
                                            p0 += (p0 % (((p0 ^ 255)) | 1))
                                            v30 := []int{}
                                            _ = v30
                                        }
                                    }
                                    var v31 bool = false
                                    if v31 {
                                        v31 = v31
                                        if v31 {
                                            v32 := []int{(fi10 % (((fi10 >> (((p0 + 4294967295)) & 63))) | 1))}
                                            switch fi10 {
                                            case 9223372036854775807:
                                                if v31 {
                                                    v33 := GenS3{F0: v31, F1: 42, F2: (-1 <= -330)}
                                                    _ = v33
                                                }
                                            default:
                                                if len(v32) > 0 {
                                                    v32[0] = fi10
                                                }
                                            }
                                            _ = v32
                                        } else {
                                            p2 -= 1
                                        }
                                    } else {
                                        var v34 uint = p0
                                        if (int(p1) > 2147483647) {
                                            v31 = v31
                                            var v35 uint = v34
                                            v17 = (v17 ^ 42)
                                            _ = v35
                                        } else {
                                            var v36 int16 = (v17 / (((v17 / ((42) | 1))) | 1))
                                            var v37 int8 = -69
                                            v38 := []int{}
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                        var v39 uint8 = p1
                                        _ = v34
                                        _ = v39
                                    }
                                    var v40 bool = false
                                    if v40 {
                                        var v41 bool = v31
                                        _ = v41
                                    }
                                    _ = v31
                                    _ = v40
                                } else {
                                    v17 += 102
                                }
                                fmt.Println(v8, p2)
                            } else {
                                p2 = p2
                                p0 += uint(v8)
                            }
                            v42 := GenS4{F0: p0, F1: v17}
                            _ = v42
                        }
                        _ = v17
                    }
                    if false {
                        var v43 float32 = (0.0 - (float32((-2.25 / ((0.1 * 0.1) - 3.14159))) / (1e10 * 0.1)))
                        v43 = -2.25
                        _ = v43
                    } else {
                        if false {
                            p2 = p2
                            fmt.Println(p2)
                            var v44 uint8 = (p1 % (((^(p1))) | 1))
                            _ = v44
                        } else {
                            var v45 string = "hello"
                            var v46 int32 = 42
                            var v47 int64 = 9223372036854775807
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        fmt.Println(fi10, fi10)
                        var v48 string = ("" + ("foo bar" + "a"))
                        p1 = p1
                        v49 := []int{fi10, -73, int(p1)}
                        switch "hello" {
                        case "a":
                            v50 := GenS5{F0: p1}
                            _ = v50
                        case "":
                            v51 := GenS1{F0: (p1 + (p1 ^ p1)), F1: -531}
                            _ = v51
                        default:
                            v8 += 526
                        }
                        switch int(p0) {
                        case 127:
                            p1 = 201
                        case 11:
                            var v52 uint = p0
                            _ = v52
                        case 32767:
                            if (true && false) {
                                var v53 bool = false
                                _ = v53
                            } else {
                                p1 += (p1 + 245)
                                v48 = v48
                                var v54 uint32 = (v8 / ((v8) | 1))
                                if true {
                                    v49 = append(v49, (fi15 >> ((478) & 63)))
                                    var v55 int8 = 1
                                    v8 = (v54 + (v54 - (v54 & (v8 + (v8 >> (((p0 & (^(p0)))) & 63))))))
                                    if false {
                                        if true {
                                            v56 := GenS5{F0: p1}
                                            _ = v56
                                        }
                                        v54 += uint32(v7)
                                        p1 -= (p1 & uint8(p0))
                                    } else {
                                        p1 += 134
                                        v57 := GenS5{F0: 128}
                                        _ = v57
                                    }
                                    p2 = p2
                                    _ = v55
                                } else {
                                    v54 = 2147483647
                                }
                                if (-128 == -65) {
                                    if len(v49) > 0 {
                                        v49[0] = (fi10 - fi15)
                                    }
                                }
                                _ = v54
                            }
                        default:
                            if ((65535 > p0) || false) {
                                v58 := GenS3{F0: true, F1: -333, F2: (uint64(p1) == (p2 / ((p2) | 1)))}
                                var v59 int16 = -1
                                _ = v58
                                _ = v59
                            } else {
                                switch -133 {
                                case 2147483647:
                                    var v60 uint8 = 0
                                    _ = v60
                                }
                            }
                        }
                        v7 += (v7 >> ((p0) & 63))
                        if true {
                            switch (^(fi15)) {
                            case 32767:
                                v48 = v48
                            case 666:
                                v61 := GenS6{F0: true}
                                _ = v61
                            case -2147483648:
                                if true {
                                    v48 = ("a" + v48)
                                    var v62 int64 = 828
                                    _ = v62
                                } else {
                                    if false {
                                        var v63 uint16 = (v7 - uint16(p1))
                                        if false {
                                            var v64 float32 = 1.5
                                            var v65 uint64 = (^(p2))
                                            _ = v64
                                            _ = v65
                                        }
                                        if len(v49) >= 2 {
                                            v49 = v49[1:len(v49)]
                                        }
                                        _ = v63
                                    } else {
                                        v7 -= v7
                                        fmt.Println(fi10)
                                        v7 -= v7
                                    }
                                    p2 = p2
                                }
                            default:
                                var v66 uint16 = (v7 % ((v7) | 1))
                                _ = v66
                            }
                            v67 := GenS7{F0: (p1 + p1), F1: -547, F2: v7}
                            _ = v67
                        }
                        if ("foo bar" != v48) {
                            p1 += p1
                        } else {
                            v7 -= (v7 + v7)
                            switch fi15 {
                            case 644:
                                v49 = append(v49, (fi10 * fi15))
                            case 32768:
                                v68 := GenS4{F0: p0, F1: 1000}
                                _ = v68
                            }
                        }
                        _ = v48
                        _ = v49
                    }
                    var v69 float32 = float32(0.1)
                    _ = fi15
                    _ = v69
                }
                if true {
                    if true {
                        var v70 uint8 = uint8(p2)
                        var v71 uint8 = (v70 ^ 52)
                        _ = v70
                        _ = v71
                    } else {
                        v72 := GenS5{F0: p1}
                        v73 := GenS6{F0: true}
                        if (4294967295 <= (^(v8))) {
                            if false {
                                v74 := GenS8{F0: 114, F1: (p2 & (^(p2)))}
                                for fi75 := 0; fi75 < 2; fi75++ {
                                    switch (fi10 ^ (^(fi75))) {
                                    case 9223372036854775807:
                                        v76 := GenS2{F0: (^(v8)), F1: "a"}
                                        _ = v76
                                    default:
                                        var v77 float64 = (float64((0.0 - 0.1)) + (0.0 - 0.1))
                                        _ = v77
                                    }
                                    v78 := []int{fi10, (fi10 % (((fi75 & (^(fi10)))) | 1)), fi10}
                                    _ = fi75
                                    _ = v78
                                }
                                _ = v74
                            }
                        }
                        _ = v72
                        _ = v73
                    }
                    switch (fi10 / (((fi10 * int(v8))) | 1)) {
                    case 32768:
                        var v79 int16 = 480
                        _ = v79
                    }
                } else {
                    switch 4294967295 {
                    case -820:
                        p1 = 2
                    default:
                        v80 := []int{(fi10 << ((p0) & 63)), (fi10 & fi10)}
                        _ = v80
                    }
                }
                _ = v14
            }
            var v81 int32 = 0
            fmt.Println(v7, p0)
            _ = v81
        }
        _ = fi10
        _ = v11
    }
    for fi82 := 0; fi82 < 15; fi82++ {
        v83 := []int{fi82}
        var v84 uint8 = 0
        switch fi82 {
        case 127:
            fmt.Println(p1)
        case 724:
            switch "" {
            case "foo bar":
                if true {
                    v84 = (v84 + 2)
                    v84 += v84
                    v7 = (^(v7))
                }
            case "hello":
                v85 := GenS2{F0: uint32(p0), F1: "foo bar"}
                _ = v85
            }
        }
        _ = fi82
        _ = v83
        _ = v84
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 float64 = (0.1 + float64(1.5))
    var v2 uint = 1
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 int8 = 127
    v6 := make(map[int]bool)
    var v86 float32 = (3.14159 - 0.1)
    if (42 >= -1) {
        v5 = v5
        v87 := []int{32767}
        _ = v87
    } else {
        v5 -= v5
        fmt.Println(v2, v1)
        if (false || true) {
            v88 := v6[-9223372036854775808]
            v88 = true
            _ = v88
        }
    }
    var v89 string = (("hello" + (("" + ("a" + "\n\t")) + ("" + "a"))) + "a")
    switch 32768 {
    case -227:
        v90 := GenS9{F0: 1, F1: 631}
        _ = v90
    default:
        switch v89 {
        case "\n\t":
            var v91 int16 = -570
            _ = v91
        case "foo bar":
            fmt.Println(v1, v86, v5)
        case "":
            if (!false) {
                for fi92 := 0; fi92 < 8; fi92++ {
                    v93 := []int{fi92, 128}
                    v1 -= 1.5
                    var v94 uint64 = 18446744073709551615
                    _ = fi92
                    _ = v93
                    _ = v94
                }
                for fi95 := 0; fi95 < 3; fi95++ {
                    v96 := v4[(fi95 & fi95)]
                    _ = fi95
                    _ = v96
                }
            } else {
                var v97 uint8 = 119
                v98 := []int{-296}
                v1 = v1
                _ = v97
                _ = v98
            }
        default:
            fmt.Println(v5, v89, v5)
        }
    }
    v99 := GenS10{F0: 2147483647, F1: -590}
    var v100 uint64 = uint64(v2)
    if false {
        var v101 int32 = -608
        _ = v101
    }
    var v102 uint = v2
    switch 212 {
    case -659:
        fmt.Println(v2, v5)
    case -397:
        var v103 float32 = float32(v1)
        _ = v103
    case 127:
        var v104 int32 = 204
        _ = v104
    }
    switch 861 {
    case -116:
        v100 += v100
    case -516:
        for fi105 := 0; fi105 < 14; fi105++ {
            _ = len(v4)
            switch (fi105 / (((fi105 % (((fi105 / (((fi105 ^ (fi105 % (((fi105 | (fi105 | fi105))) | 1)))) | 1))) | 1))) | 1)) {
            case 127:
                v102 = v2
            case 128:
                for fi106 := 0; fi106 < 17; fi106++ {
                    v107 := v6[(fi106 % ((729) | 1))]
                    switch (fi106 % ((int(v100)) | 1)) {
                    case 4294967295:
                        _ = len(v6)
                    case -9223372036854775808:
                        v1 += v1
                    case -203:
                        var v108 float32 = 3.14159
                        _ = v108
                    }
                    _ = fi106
                    _ = v107
                }
            }
            v2 = 0
            _ = fi105
        }
    case -1:
        switch (("hello" + v89) + "\n\t") {
        case "foo bar":
            delete(v4, -142)
        case "a":
            v109 := GenS11{F0: 0}
            _ = v109
        }
    }
    for fi110 := 0; fi110 < 5; fi110++ {
        var v111 int8 = 127
        _ = fi110
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v86:", v86)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v100
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
    _ = v89
    _ = v99
}
