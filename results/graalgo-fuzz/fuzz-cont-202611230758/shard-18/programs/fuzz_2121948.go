package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint16; F2 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 uint64 }
func genFn0(p0 int16, p1 uint16, p2 uint64) int16 {
    v5 := []int{4294967295, 128, -198, 2147483647}
    if (!true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch "" {
        case "\n\t":
            var v6 uint8 = 52
            _ = v6
        default:
            _ = len(v5)
        }
        var v7 string = "\n\t"
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int16(p1)
}
func genFn1(p0 uint16, p1 string) uint64 {
    for fi8 := 0; fi8 < 16; fi8++ {
        if false {
            for fi9 := 0; fi9 < 20; fi9++ {
                p1 = (((p1 + "a") + ("a" + "")) + p1)
                v10 := []int{}
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi9
                _ = v10
            }
            p1 = ("foo bar" + p1)
            var v11 string = "\n\t"
            _ = v11
        }
        _ = fi8
    }
    p0 += p0
    _ = p0
    _ = p1
    return uint64(p0)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 uint32 = 42
    v4 += v4
    var v12 int = 42
    for fi13 := 0; fi13 < 12; fi13++ {
        switch (("a" + "foo bar") + ("foo bar" + "a")) {
        case "a":
            switch "" {
            case "":
                v12 = fi13
            default:
                if false {
                    switch "foo bar" {
                    case "a":
                        v14 := v2[(^(v12))]
                        _ = v14
                    case "foo bar":
                        v12 = (^(fi13))
                    case "":
                        if true {
                            if ("a" == "foo bar") {
                                switch "foo bar" {
                                case "\n\t":
                                    for fi15 := 0; fi15 < 5; fi15++ {
                                        if true {
                                            var v16 int64 = -1
                                            delete(v1, "a")
                                            if (!(v16 != v16)) {
                                                var v17 int16 = 953
                                                _ = v17
                                            }
                                            var v18 int8 = 0
                                            v16 = -1
                                            _ = v16
                                            _ = v18
                                        }
                                        v19 := []int{(fi15 / (((v12 % ((-1) | 1))) | 1))}
                                        var v20 int8 = int8(v12)
                                        if false {
                                            if (((978 == 42) && true) || (-1 > 42)) {
                                                var v21 uint = 1
                                                _ = v21
                                            } else {
                                                switch int(v4) {
                                                case 9223372036854775807:
                                                    v4 = v4
                                                case 32767:
                                                    v22 := GenS1{F0: 941}
                                                    _ = v22
                                                case 2147483647:
                                                    var v23 int16 = 592
                                                    v20 -= -109
                                                    v4 = 2
                                                    var v24 int = fi15
                                                    v25 := GenS1{F0: 2}
                                                    _ = v23
                                                    _ = v24
                                                    _ = v25
                                                default:
                                                    delete(v1, "hello")
                                                }
                                                fmt.Println(v12, v4)
                                                if false {
                                                    if true {
                                                        var v26 float32 = float32(1.5)
                                                        if true {
                                                            var v27 uint64 = 240
                                                            v4 = (^(v4))
                                                            var v28 uint64 = genFn1(199, ("\n\t" + "foo bar"))
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        var v29 int32 = -207
                                                        _ = v26
                                                        _ = v29
                                                    } else {
                                                        var v30 uint8 = 42
                                                        if true {
                                                            if (true || (!(v20 == 55))) {
                                                                var v31 float32 = 1e10
                                                                _ = v31
                                                            } else {
                                                                switch ("" + "hello") {
                                                                case "a":
                                                                    if (!(!(!((true && (!true)) || true)))) {
                                                                        var v32 string = ""
                                                                        if ((!false) && ((!true) || true)) {
                                                                            v20 = v20
                                                                        } else {
                                                                            v12 = fi15
                                                                            v33 := GenS1{F0: 4294967295}
                                                                            _ = v33
                                                                        }
                                                                        v12 += fi15
                                                                        var v34 int16 = -32768
                                                                        var v35 string = ""
                                                                        _ = v32
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                    var v36 float64 = (-2.25 * 0.0)
                                                                    var v37 uint16 = 905
                                                                    _ = v36
                                                                    _ = v37
                                                                case "\n\t":
                                                                    fmt.Println(fi13, v4, fi15)
                                                                case "foo bar":
                                                                    v38 := GenS1{F0: 772}
                                                                    _ = v38
                                                                }
                                                                v39 := GenS2{F0: 1, F1: 1}
                                                                if len(v19) >= 2 {
                                                                    v19 = v19[1:len(v19)]
                                                                }
                                                                var v40 float32 = (((0.1 + float32(0.1)) + 0.0) / 1.0)
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                            switch -1 {
                                                            case -344:
                                                                v20 = v20
                                                            case -9223372036854775808:
                                                                var v41 float64 = ((0.1 * 1.5) / float64((1.5 + (0.1 * 1.5))))
                                                                _ = v41
                                                            case 32767:
                                                                var v42 uint8 = 0
                                                                _ = v42
                                                            default:
                                                                var v43 float32 = 0.0
                                                                _ = v43
                                                            }
                                                            if (!((v20 << ((18446744073709551615) & 63)) < 72)) {
                                                                if false {
                                                                    v20 = 1
                                                                }
                                                                switch v12 {
                                                                case 127:
                                                                    if false {
                                                                        var v44 int8 = 70
                                                                        _ = v44
                                                                    } else {
                                                                        fmt.Println(fi15)
                                                                        if true {
                                                                            v45 := v1["\n\t"]
                                                                            var v46 uint32 = (^(v4))
                                                                            var v47 int64 = -9223372036854775808
                                                                            var v48 int8 = 127
                                                                            _ = v45
                                                                            _ = v46
                                                                            _ = v47
                                                                            _ = v48
                                                                        } else {
                                                                            v49 := GenS2{F0: 1, F1: 65535}
                                                                            var v50 uint = 4294967295
                                                                            if true {
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    var v51 uint32 = (v4 >> ((65535) & 63))
                                                                                    _ = v51
                                                                                }
                                                                                var v52 int8 = 36
                                                                                if (2 == uint16(v30)) {
                                                                                    v52 = (v20 + v20)
                                                                                } else {
                                                                                    v50 = (v50 + v50)
                                                                                    fmt.Println(v20, v52, fi15)
                                                                                    if (!((true || (v4 <= (v4 & 0))) || ("\n\t" < ("" + "foo bar")))) {
                                                                                        v4 = v4
                                                                                    } else {
                                                                                        var v53 int32 = 862
                                                                                        fmt.Println(fi13)
                                                                                        if false {
                                                                                            v52 += v52
                                                                                        }
                                                                                        v53 += -197
                                                                                        _ = v53
                                                                                    }
                                                                                }
                                                                                v12 += (v12 << ((v50) & 63))
                                                                                fmt.Println(fi13, v50)
                                                                                if true {
                                                                                    if len(v19) >= 2 {
                                                                                        v19 = v19[1:len(v19)]
                                                                                    }
                                                                                } else {
                                                                                    if len(v19) >= 2 {
                                                                                        v19 = v19[1:len(v19)]
                                                                                    }
                                                                                    if (false && ((true || (!false)) || (-545 <= -23))) {
                                                                                        if false {
                                                                                            v50 += (v50 % (((^(v50))) | 1))
                                                                                        }
                                                                                    } else {
                                                                                        delete(v3, 790)
                                                                                    }
                                                                                    v54 := GenS2{F0: -1, F1: v50}
                                                                                    _ = v54
                                                                                }
                                                                                var v55 int = v12
                                                                                _ = v52
                                                                                _ = v55
                                                                            } else {
                                                                                var v56 uint32 = 2
                                                                                v4 -= v56
                                                                                _ = v56
                                                                            }
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                        var v57 uint64 = 542
                                                                        _ = v57
                                                                    }
                                                                case -1:
                                                                    v20 -= v20
                                                                case -585:
                                                                    fmt.Println(fi15, v20, v4)
                                                                }
                                                            }
                                                        }
                                                        v58 := v2[(v12 / (((fi15 % ((-793) | 1))) | 1))]
                                                        _ = v30
                                                        _ = v58
                                                    }
                                                    switch ("a" + "\n\t") {
                                                    case "a":
                                                        v59 := GenS3{F0: v12, F1: v4}
                                                        _ = v59
                                                    case "\n\t":
                                                        var v60 int64 = int64(v4)
                                                        _ = v60
                                                    case "":
                                                        v20 += v20
                                                    }
                                                    var v61 uint16 = 65535
                                                    fmt.Println(v4)
                                                    fmt.Println(fi13, v12, v12)
                                                    _ = v61
                                                }
                                            }
                                            fmt.Println(v12, v4, fi13)
                                        }
                                        var v62 uint8 = 0
                                        if (true && true) {
                                            _ = len(v3)
                                            v63 := GenS3{F0: fi13, F1: (v4 & 4294967295)}
                                            var v64 uint8 = (v62 % ((182) | 1))
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = fi15
                                        _ = v19
                                        _ = v20
                                        _ = v62
                                    }
                                }
                                v65 := GenS4{F0: 9223372036854775807}
                                var v66 int16 = 1000
                                _ = v65
                                _ = v66
                            } else {
                                v67 := GenS3{F0: v12, F1: v4}
                                _ = v67
                            }
                            var v68 int8 = -128
                            v4 = (v4 * v4)
                            _ = v68
                        }
                    default:
                        var v69 uint64 = 4294967295
                        _ = v69
                    }
                    v70 := []int{v12, (v12 + (v12 / (((v12 - fi13)) | 1))), -2147483648}
                    if len(v70) > 0 {
                        v70[0] = (-(v12))
                    }
                    _ = v70
                } else {
                    var v71 uint8 = 128
                    if (true || true) {
                        v72 := []int{}
                        v12 -= (fi13 / ((fi13) | 1))
                        var v73 uint = 42
                        _ = v72
                        _ = v73
                    }
                    _ = v71
                }
            }
        case "":
            if (-1 < -1) {
                v12 = (-(fi13))
                v74 := GenS5{F0: -1, F1: 2}
                v12 = fi13
                _ = v74
            } else {
                v75 := []int{(v12 * fi13), v12, fi13}
                _ = len(v1)
                var v76 bool = (!((-236 > 2147483647) && (-2147483648 != -775)))
                _ = v75
                _ = v76
            }
        case "foo bar":
            fmt.Println(v12)
        default:
            for fi77 := 0; fi77 < 16; fi77++ {
                v12 = (fi13 * (fi13 - (fi77 - (^(fi77)))))
                v78 := GenS4{F0: 0}
                if true {
                    var v79 float32 = (1.5 - 0.0)
                    _ = v79
                } else {
                    var v80 int32 = -874
                    v80 += v80
                    var v81 int64 = -1
                    _ = v80
                    _ = v81
                }
                _ = fi77
                _ = v78
            }
        }
        for fi82 := 0; fi82 < 4; fi82++ {
            if true {
                v83 := GenS1{F0: 89}
                fmt.Println(fi13, v12)
                _ = v83
            }
            _ = fi82
        }
        _ = fi13
    }
    v12 -= (v12 % (((v12 * (v12 | 0))) | 1))
    if (false && false) {
        var v84 uint64 = 947
        var v85 uint64 = genFn1(101, ("foo bar" + "hello"))
        _ = v84
        _ = v85
    } else {
        v86 := []int{-43, v12}
        _ = v86
    }
    for fi87 := 0; fi87 < 6; fi87++ {
        v12 = (-(fi87))
        var v88 uint32 = 169
        v1[("a" + ("a" + "a"))] = (^(v4))
        _ = fi87
        _ = v88
    }
    var v89 string = ("" + "")
    v12 = (v12 % ((127) | 1))
    if (v12 < 32767) {
        v12 = v12
    } else {
        if (1 > 498) {
            v90 := GenS6{F0: false, F1: 54, F2: 42}
            _ = v90
        } else {
            _ = len(v1)
        }
        for fi91 := 0; fi91 < 5; fi91++ {
            v12 += (-(v12))
            var v92 uint32 = v4
            for fi93 := 0; fi93 < 5; fi93++ {
                var v94 int16 = -190
                v95 := GenS7{F0: 21}
                fmt.Println(v94, v92, v4)
                _ = fi93
                _ = v94
                _ = v95
            }
            _ = fi91
            _ = v92
        }
    }
    var v96 int16 = -1
    v97 := v1[v89]
    v96 = (v96 & (v96 ^ (v96 | (v96 + v96))))
    delete(v3, v12)
    fmt.Println(v4, v89, v4)
    var v98 int64 = 1
    var v99 uint8 = 128
    v100 := GenS8{F0: true, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v89
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
