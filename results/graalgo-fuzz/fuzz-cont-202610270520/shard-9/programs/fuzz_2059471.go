package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 int64; F1 int32; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint8; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int; F1 int32; F2 int16 }
type GenS9 struct { F0 uint64; F1 int8 }
func genFn0(p0 int64, p1 int, p2 int64) uint {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: (fi6 * fi6), F1: (!true), F2: 32767}
        _ = fi6
        _ = v7
    }
    var v8 int32 = -84
    var v9 uint16 = 0
    switch (p1 ^ (^(p1))) {
    case -1:
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{(fi10 - (-(fi10))), (fi10 * 128)}
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS2{F0: 1, F1: 1, F2: ((255 <= 1) || (!((-1 >= p0) || false)))}
        _ = v12
    }
    var v13 uint8 = 177
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int8 = 100
    var v4 float64 = 0.1
    v5 := make(map[uint32]bool)
    switch 9223372036854775807 {
    case 2147483647:
        v4 -= (0.1 / v4)
    case 127:
        v14 := []int{311}
        _ = v14
    case -1:
        v15 := GenS3{F0: 0, F1: -263, F2: ((("" + "") + ("\n\t" + "foo bar")) + (((("foo bar" + "\n\t") + "a") + "hello") + "hello"))}
        _ = v15
    }
    v4 += float64(3.14159)
    v3 -= (v3 / ((-22) | 1))
    var v16 uint32 = 0
    if (527 != (v16 ^ (v16 << ((375) & 63)))) {
        switch "\n\t" {
        case "":
            var v17 int8 = v3
            _ = v17
        default:
            v18 := []int{2147483647, 32768, -9223372036854775808}
            _ = v18
        }
        if (3.14159 > float32(v4)) {
            if (!false) {
                v3 = 1
            } else {
                var v19 uint = 0
                _ = v19
            }
            var v20 int64 = 958
            switch "" {
            case "":
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 string = ""
                    v23 := GenS4{F0: -1}
                    var v24 uint8 = 228
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                var v25 uint8 = 154
                _ = v25
            case "foo bar":
                if false {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v27 := GenS5{F0: -344, F1: 42, F2: true}
                        _ = fi26
                        _ = v27
                    }
                    v3 += (v3 & v3)
                    v20 += v20
                } else {
                    v28 := GenS1{F0: 495, F1: false, F2: 0}
                    _ = v28
                }
            }
            _ = v20
        } else {
            if true {
                if true {
                    delete(v2, "a")
                    v4 -= (3.14159 * 0.0)
                    if false {
                        v16 += 4294967295
                        var v29 int16 = -32768
                        _ = v29
                    } else {
                        v30 := []int{127, 4294967295, 686, 32768}
                        _ = v30
                    }
                } else {
                    var v31 uint8 = 42
                    v32 := []int{}
                    if len(v32) > 0 {
                        v32[0] = 963
                    }
                    _ = v31
                    _ = v32
                }
            } else {
                v3 = v3
                var v33 uint8 = 128
                _ = v33
            }
        }
    } else {
        v5[(v16 | 520)] = true
        var v34 uint16 = 148
        _ = v34
    }
    if true {
        if (-952 < -1) {
            v35 := v2["a"]
            _ = v35
        } else {
            if false {
                fmt.Println(v4)
                var v36 bool = true
                var v37 uint32 = (^(v16))
                _ = v36
                _ = v37
            } else {
                fmt.Println(v16, v4, v16)
                if ("\n\t" == "hello") {
                    v38 := GenS4{F0: 1}
                    var v39 int16 = -1
                    if false {
                        for fi40 := 0; fi40 < 11; fi40++ {
                            for fi41 := 0; fi41 < 17; fi41++ {
                                var v42 float64 = v4
                                v43 := []int{656, fi40, (fi41 ^ (fi41 | fi40)), (fi41 % ((-9223372036854775808) | 1))}
                                var v44 bool = ((v16 < 296) || (-2147483648 != 2147483647))
                                _ = fi41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            _ = fi40
                        }
                    }
                    _ = v38
                    _ = v39
                }
            }
            fmt.Println(v4, v3)
            if ((v16 | (v16 >> ((660) & 63))) > v16) {
                var v45 int32 = 2147483647
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 float32 = 3.14159
                    var v48 int16 = 42
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = v45
            }
        }
    } else {
        if false {
            fmt.Println(v3, v3, v16)
            v49 := GenS6{F0: 538}
            fmt.Println(v3, v3)
            _ = v49
        }
        _ = len(v1)
    }
    if (0.1 < 0.1) {
        if false {
            switch "a" {
            case "":
                for fi50 := 0; fi50 < 11; fi50++ {
                    v4 = float64(0.1)
                    v3 -= -15
                    v51 := []int{}
                    _ = fi50
                    _ = v51
                }
            default:
                for fi52 := 0; fi52 < 19; fi52++ {
                    if (42 == 1) {
                        if (57 != genFn0(1, (^(fi52)), -1)) {
                            switch fi52 {
                            case -277:
                                v53 := []int{(fi52 ^ -63), fi52, fi52}
                                _ = v53
                            case -381:
                                for fi54 := 0; fi54 < 14; fi54++ {
                                    fmt.Println(v16, v16)
                                    v16 += v16
                                    if true {
                                        v55 := GenS1{F0: (fi52 % ((int(v16)) | 1)), F1: (v4 >= v4), F2: 1}
                                        switch -977 {
                                        case -328:
                                            var v56 uint32 = (^(v16))
                                            _ = v56
                                        }
                                        fmt.Println(v3, v3)
                                        _ = v55
                                    } else {
                                        if (!true) {
                                            fmt.Println(fi54, v3, fi54)
                                            fmt.Println(v16, v3)
                                            var v57 int = fi52
                                            var v58 uint64 = 2
                                            switch (v57 % (((fi54 + v57)) | 1)) {
                                            case 42:
                                                v58 += v58
                                            default:
                                                var v59 int32 = -2147483648
                                                _ = v59
                                            }
                                            switch fi54 {
                                            case 220:
                                                v5[(v16 ^ v16)] = false
                                                v58 = 42
                                            case 9223372036854775807:
                                                var v60 uint16 = 778
                                                _ = v60
                                            case 32768:
                                                var v61 int8 = (v3 + 1)
                                                _ = v61
                                            }
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v62 := GenS2{F0: -1, F1: (v16 * v16), F2: (false || true)}
                                            switch fi54 {
                                            case -1:
                                                var v63 uint16 = 579
                                                _ = v63
                                            case 0:
                                                v16 -= v16
                                            case 1:
                                                var v64 int64 = 0
                                                _ = v64
                                            default:
                                                delete(v5, 394)
                                                _ = len(v1)
                                                fmt.Println(fi54, fi52, v4)
                                            }
                                            if (0 >= 300) {
                                                if true {
                                                    v65 := GenS3{F0: 42, F1: 0, F2: ("a" + (("foo bar" + "foo bar") + "a"))}
                                                    fmt.Println(v4, v16)
                                                    _ = v65
                                                } else {
                                                    v66 := []int{(fi52 & fi54)}
                                                    _ = v66
                                                }
                                                var v67 int8 = (v3 - (v3 & 42))
                                                var v68 bool = false
                                                _ = v67
                                                _ = v68
                                            } else {
                                                fmt.Println(fi52, v16)
                                                var v69 int = int(v16)
                                                _ = v69
                                            }
                                            _ = v62
                                        }
                                    }
                                    v16 = 65535
                                    _ = fi54
                                }
                            case -644:
                                if (!(((v4 + v4) - (float64(-2.25) / (float64(0.1) * v4))) != 1.5)) {
                                    if (false || true) {
                                        v4 += float64((((-2.25 + 0.1) / (((-2.25 / 3.14159) - (1e10 * 3.14159)) * 1.5)) - (0.1 / 1.0)))
                                        v16 = 706
                                        if true {
                                            v3 = (^(v3))
                                            for fi70 := 0; fi70 < 11; fi70++ {
                                                var v71 uint = 18446744073709551615
                                                var v72 uint16 = 32767
                                                _ = fi70
                                                _ = v71
                                                _ = v72
                                            }
                                        } else {
                                            v3 = (v3 - 0)
                                        }
                                    } else {
                                        for fi73 := 0; fi73 < 1; fi73++ {
                                            delete(v1, (fi52 % (((fi52 ^ fi73)) | 1)))
                                            v74 := []int{(fi52 - fi73), fi52, (-(fi52))}
                                            _ = fi73
                                            _ = v74
                                        }
                                        fmt.Println(v3, fi52)
                                        v1[(fi52 >> ((0) & 63))] = ((-9223372036854775808 > 9223372036854775807) || true)
                                    }
                                } else {
                                    for fi75 := 0; fi75 < 8; fi75++ {
                                        switch (fi52 >> ((2) & 63)) {
                                        case -2147483648:
                                            if false {
                                                switch ("\n\t" + ("foo bar" + (("\n\t" + "hello") + "foo bar"))) {
                                                case "a":
                                                    v16 -= (v16 | 12)
                                                case "":
                                                    delete(v2, ("a" + "a"))
                                                case "hello":
                                                    v76 := []int{fi75, 0, -113, fi75}
                                                    _ = v76
                                                default:
                                                    if (v4 <= float64(float32((v4 / 1.5)))) {
                                                        v16 = (v16 & v16)
                                                        if true {
                                                            fmt.Println(v4, fi75)
                                                            fmt.Println(v4, v16)
                                                            v77 := []int{(fi52 ^ (^(fi75))), (fi75 * (fi52 ^ fi75)), fi52, fi52}
                                                            _ = v77
                                                        }
                                                    } else {
                                                        v16 = v16
                                                        v16 -= (v16 >> ((2) & 63))
                                                    }
                                                }
                                                if (false || true) {
                                                    switch "a" {
                                                    case "":
                                                        v3 += (v3 & v3)
                                                        var v78 int = (^(fi52))
                                                        var v79 int64 = 0
                                                        _ = v78
                                                        _ = v79
                                                    case "a":
                                                        var v80 uint16 = 932
                                                        _ = v80
                                                    default:
                                                        var v81 uint8 = 127
                                                        _ = v81
                                                    }
                                                    v82 := []int{(fi52 & 344), (fi75 - fi52)}
                                                    fmt.Println(v3, v16, fi52)
                                                    _ = v82
                                                } else {
                                                    v3 = (v3 >> ((4294967295) & 63))
                                                    fmt.Println(v4, v3, v4)
                                                    v3 += (v3 >> ((18446744073709551615) & 63))
                                                    if false {
                                                        v4 = 1.5
                                                    } else {
                                                        _ = len(v1)
                                                    }
                                                }
                                                if (!(821 > 0)) {
                                                    if (-53 > 42) {
                                                        delete(v5, (v16 & v16))
                                                    }
                                                    if (!(65535 <= 780)) {
                                                        v83 := []int{fi75, (-(fi75)), 9223372036854775807, (-(fi52))}
                                                        _ = v83
                                                    } else {
                                                        v84 := GenS7{F0: 4294967295, F1: 1}
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if false {
                                                                delete(v1, (fi75 - (^(fi75))))
                                                                v3 += v3
                                                                fmt.Println(fi52)
                                                            }
                                                        case "":
                                                            if ((1.5 * float32(v4)) > ((1.5 + (float32(float64(1e10)) + 1.5)) * 1e10)) {
                                                                if ("" >= (("foo bar" + "a") + "hello")) {
                                                                    _ = len(v1)
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v85 uint32 = v16
                                                                        _ = v85
                                                                    case "\n\t":
                                                                        v3 = (v3 / (((v3 << ((1) & 63))) | 1))
                                                                    }
                                                                    v1[(fi75 & 2147483647)] = (true && false)
                                                                    v16 = 2
                                                                    v86 := GenS8{F0: (fi75 ^ fi75), F1: 1, F2: 32767}
                                                                    _ = v86
                                                                } else {
                                                                    var v87 float32 = ((1.5 * 0.0) + 0.0)
                                                                    v88 := GenS5{F0: -1, F1: 255, F2: true}
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                                v16 = 2
                                                            }
                                                            if true {
                                                                var v89 uint16 = 65535
                                                                v5[v16] = (376 <= 42)
                                                                v3 = int8(v16)
                                                                _ = v89
                                                            } else {
                                                                v90 := GenS6{F0: 32767}
                                                                if true {
                                                                    v91 := []int{fi52, fi52, fi52, fi75}
                                                                    v3 += (v3 << ((2) & 63))
                                                                    var v92 int32 = -2147483648
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                v93 := []int{fi75, fi52, fi52, 854}
                                                                _ = v90
                                                                _ = v93
                                                            }
                                                        case "foo bar":
                                                            v4 = ((float64(float32(3.14159)) / (v4 + 1e10)) + 0.0)
                                                        }
                                                        fmt.Println(fi52, v3, fi75)
                                                        _ = v84
                                                    }
                                                    var v94 uint = 65535
                                                    _ = v94
                                                } else {
                                                    v16 = (v16 & v16)
                                                    _ = len(v5)
                                                }
                                            }
                                        default:
                                            fmt.Println(fi52)
                                        }
                                        v95 := GenS8{F0: (^(fi75)), F1: 0, F2: -348}
                                        _ = fi75
                                        _ = v95
                                    }
                                    for fi96 := 0; fi96 < 15; fi96++ {
                                        if ((v16 >> ((uint(v16)) & 63)) == (v16 - (v16 % (((v16 << ((4294967295) & 63))) | 1)))) {
                                            if (!false) {
                                                var v97 uint8 = 128
                                                v98 := GenS9{F0: 1, F1: (v3 + v3)}
                                                v97 += (v97 + 42)
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                        _ = fi96
                                    }
                                }
                            default:
                                for fi99 := 0; fi99 < 17; fi99++ {
                                    var v100 uint = 18446744073709551615
                                    _ = fi99
                                    _ = v100
                                }
                            }
                        }
                    } else {
                        v101 := []int{(fi52 + fi52), (^(fi52)), fi52, fi52}
                        _ = len(v2)
                        v3 -= 42
                        _ = v101
                    }
                    _ = fi52
                }
            }
        }
        v102 := []int{}
        if (32767 != 65535) {
            v4 -= (1.5 / float64((-2.25 * -2.25)))
        } else {
            var v103 float64 = -2.25
            var v104 int16 = 1000
            _ = v103
            _ = v104
        }
        _ = v102
    } else {
        v105 := []int{-1}
        v16 = v16
        _ = v105
    }
    var v106 int64 = 9223372036854775807
    var v107 string = ("foo bar" + "hello")
    v108 := []int{4294967295, 1, 9223372036854775807, 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v106
    _ = v107
    _ = v108
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
