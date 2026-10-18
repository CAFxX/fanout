package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS10 struct { F0 uint32; F1 string; F2 int16 }
type GenS11 struct { F0 uint8; F1 int; F2 bool }
type GenS12 struct { F0 uint16; F1 int8 }
type GenS13 struct { F0 int; F1 int32; F2 int8 }
type GenS14 struct { F0 bool; F1 uint8 }
type GenS15 struct { F0 int64; F1 uint64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int32; F1 int32; F2 uint16 }
type GenS18 struct { F0 string; F1 string; F2 string }
func genFn0(p0 string, p1 int32) uint16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int = (fi4 + fi4)
        if (1 <= 2) {
            var v6 uint32 = 65535
            _ = v6
        }
        var v7 uint64 = 0
        _ = fi4
        _ = v5
        _ = v7
    }
    v8 := GenS1{F0: 531, F1: 42}
    var v9 float32 = float32(1e10)
    p1 += (p1 | p1)
    switch ((p0 + (p0 + "\n\t")) + (p0 + p0)) {
    case "":
        p0 = ("" + (p0 + p0))
    case "foo bar":
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int = -2147483648
    var v2 int32 = -1
    v3 := make(map[string]string)
    if (3.14159 <= 0.1) {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint = 272
            switch fi10 {
            case -9223372036854775808:
                v2 = v2
            }
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 string = "a"
        var v14 uint = 0
        if (true || (2147483648 <= 2147483648)) {
            if false {
                fmt.Println(v13, v2)
            } else {
                v15 := []int{fi12, 343, (v1 / (((v1 % ((fi12) | 1))) | 1))}
                for fi16 := 0; fi16 < 2; fi16++ {
                    if (1000 < 1000) {
                        if true {
                            var v17 uint = v14
                            if false {
                                switch -395 {
                                case 32768:
                                    var v18 uint8 = 1
                                    v17 -= v17
                                    _ = v18
                                case -846:
                                    v1 = (fi12 ^ (fi16 << ((v17) & 63)))
                                default:
                                    v1 = fi16
                                }
                                var v19 int = v1
                                _ = v19
                            }
                            fmt.Println(v1, v13, fi12)
                            _ = v17
                        }
                        if true {
                            v20 := GenS1{F0: v2, F1: 42}
                            var v21 int = 4294967295
                            if true {
                                fmt.Println(fi16)
                                var v22 int64 = 1
                                if (!false) {
                                    fmt.Println(v21, v2, fi16)
                                    var v23 bool = false
                                    _ = v23
                                } else {
                                    v24 := GenS1{F0: (v2 << ((v14) & 63)), F1: 255}
                                    if (v22 >= (^(v22))) {
                                        var v25 float64 = 1.5
                                        var v26 uint32 = 662
                                        v27 := GenS2{F0: 42, F1: -1}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v24
                                }
                                _ = v22
                            }
                            _ = v20
                            _ = v21
                        }
                    } else {
                        v1 -= (fi12 >> (((^(v14))) & 63))
                    }
                    _ = fi16
                }
                v28 := GenS3{F0: (v13 + (v13 + "\n\t")), F1: 0, F2: 2147483648}
                _ = v15
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 int8 = 127
                switch v13 {
                case "a":
                    v2 = 2147483647
                case "":
                    if true {
                        v31 := []int{v1, 4294967295, fi29, 9223372036854775807}
                        v14 = (^(v14))
                        _ = v31
                    } else {
                        if (float64((0.0 / (-2.25 + 1.5))) == float64((1e10 - 3.14159))) {
                            var v32 uint8 = 1
                            switch v13 {
                            case "a":
                                var v33 float64 = 3.14159
                                _ = v33
                            case "hello":
                                v34 := GenS4{F0: v32, F1: (v2 / (((v2 % ((-861) | 1))) | 1)), F2: int16(fi12)}
                                _ = v34
                            case "":
                                fmt.Println(v14)
                            }
                            if (!(255 > 255)) {
                                switch "hello" {
                                case "foo bar":
                                    v2 = (^(v2))
                                case "hello":
                                    v35 := GenS2{F0: v30, F1: int16(v2)}
                                    _ = v35
                                default:
                                    if false {
                                        v36 := GenS5{F0: 18446744073709551615}
                                        var v37 int32 = (v2 << ((42) & 63))
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 float32 = float32(1.5)
                                        v39 := GenS1{F0: (v2 << ((v14) & 63)), F1: v32}
                                        var v40 uint32 = 4294967295
                                        switch -2147483648 {
                                        case 1:
                                            fmt.Println(v13, v40)
                                        case 2147483647:
                                            var v41 int = (fi29 >> ((1) & 63))
                                            _ = v41
                                        default:
                                            v42 := GenS6{F0: 487, F1: v2}
                                            _ = v42
                                        }
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                }
                                v43 := GenS7{F0: 42}
                                _ = v43
                            }
                            switch (^(fi29)) {
                            case -840:
                                if true {
                                    v32 = v32
                                    v44 := GenS8{F0: 685}
                                    if (((((((0.0 * 0.0) + 1.5) - 1e10) - 3.14159) + 1e10) > -2.25) || true) {
                                        if (691 > 37) {
                                            v45 := []int{int(v14), (fi29 << (((v14 + v14)) & 63)), 32768, int(v30)}
                                            _ = v45
                                        } else {
                                            var v46 int64 = 9223372036854775807
                                            v47 := GenS6{F0: 0, F1: (v2 % (((v2 >> (((^(v14))) & 63))) | 1))}
                                            _ = v46
                                            _ = v47
                                        }
                                        var v48 uint = 80
                                        var v49 string = v13
                                        _ = v48
                                        _ = v49
                                    } else {
                                        var v50 int64 = 0
                                        v51 := []int{}
                                        _ = v50
                                        _ = v51
                                    }
                                    switch (fi29 ^ int(v2)) {
                                    case -2147483648:
                                        v52 := GenS9{F0: 42, F1: 220, F2: (v32 << (((v14 >> ((v14) & 63))) & 63))}
                                        _ = v52
                                    default:
                                        var v53 uint = (^(v14))
                                        _ = v53
                                    }
                                    v13 = v13
                                    var v54 float32 = 0.1
                                    switch (fi29 ^ fi29) {
                                    case -1:
                                        var v55 uint = (v14 % ((v14) | 1))
                                        _ = v55
                                    case -892:
                                        v56 := GenS10{F0: 4294967295, F1: v13, F2: 0}
                                        _ = v56
                                    case -59:
                                        if true {
                                            if true {
                                                v1 -= (fi12 & fi12)
                                                switch (v13 + v13) {
                                                case "a":
                                                    v57 := GenS2{F0: v30, F1: 956}
                                                    _ = v57
                                                case "":
                                                    fmt.Println(fi12)
                                                default:
                                                    fmt.Println(v2)
                                                }
                                                if false {
                                                    var v58 int16 = 42
                                                    _ = v58
                                                }
                                            }
                                            var v59 uint8 = v32
                                            if (v32 >= v32) {
                                                delete(v3, v13)
                                                var v60 uint = 42
                                                _ = v60
                                            } else {
                                                var v61 bool = (false || ((!false) || ((v2 * (-(v2))) != 1)))
                                                var v62 float64 = 0.0
                                                if false {
                                                    var v63 float32 = 0.0
                                                    v64 := GenS1{F0: -1, F1: 190}
                                                    _ = v63
                                                    _ = v64
                                                }
                                                v65 := GenS10{F0: 2, F1: v13, F2: 397}
                                                _ = v61
                                                _ = v62
                                                _ = v65
                                            }
                                            var v66 int64 = 1
                                            var v67 int64 = int64(v30)
                                            v68 := []int{(v1 ^ 32767), (fi29 - v1), (^(v1)), (^(fi12))}
                                            _ = v59
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        } else {
                                            if ((v2 | v2) > (v2 ^ v2)) {
                                                var v69 float32 = 0.0
                                                if ((1.5 + (0.1 + v54)) != 1e10) {
                                                    v14 += 1
                                                }
                                                _ = v69
                                            } else {
                                                v70 := GenS7{F0: -1}
                                                var v71 uint64 = 558
                                                if true {
                                                    _ = len(v3)
                                                }
                                                v72 := GenS2{F0: v30, F1: 1}
                                                if (!(false || false)) {
                                                    var v73 int32 = v2
                                                    v1 -= fi12
                                                    delete(v3, (((v13 + "") + v13) + v13))
                                                    if (!false) {
                                                        switch ("\n\t" + "hello") {
                                                        case "\n\t":
                                                            v74 := GenS11{F0: (v32 + 0), F1: int(v30), F2: true}
                                                            _ = v74
                                                        case "":
                                                            var v75 uint32 = uint32(v71)
                                                            _ = v75
                                                        case "hello":
                                                            fmt.Println(v54)
                                                        default:
                                                            if (!(v32 <= v32)) {
                                                                var v76 uint = (v14 * v14)
                                                                var v77 int8 = 5
                                                                var v78 bool = true
                                                                var v79 float64 = (3.14159 + float64(float32(-2.25)))
                                                                _ = v76
                                                                _ = v77
                                                                _ = v78
                                                                _ = v79
                                                            } else {
                                                                if true {
                                                                    v13 = v13
                                                                } else {
                                                                    var v80 uint = 951
                                                                    var v81 bool = (v71 != v71)
                                                                    _ = v80
                                                                    _ = v81
                                                                }
                                                                fmt.Println(v32, v1, fi12)
                                                                if (!false) {
                                                                    v82 := GenS12{F0: 32767, F1: (v30 ^ (v30 ^ v30))}
                                                                    _ = v82
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v71 = (v71 ^ 42)
                                                        fmt.Println(v14, v54, v14)
                                                    }
                                                    _ = v73
                                                }
                                                v83 := GenS13{F0: 9223372036854775807, F1: 2147483647, F2: (v30 % ((v30) | 1))}
                                                delete(v3, (v13 + (((v13 + v13) + "") + "a")))
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                                _ = v83
                                            }
                                            if true {
                                                if (true && (0.0 > v54)) {
                                                    v84 := GenS7{F0: 32767}
                                                    if (false || ((^(v32)) == (v32 ^ v32))) {
                                                        var v85 uint = (^(v14))
                                                        if (!(((!false) && (0.0 < v54)) || ((-(fi29)) > fi29))) {
                                                            v86 := v3[(v13 + "")]
                                                            _ = v86
                                                        } else {
                                                            v87 := GenS14{F0: (false && ((v13 + v13) != "\n\t")), F1: (v32 * (v32 << ((v85) & 63)))}
                                                            var v88 int8 = v30
                                                            v88 = 100
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        _ = v85
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = v84
                                                }
                                                if false {
                                                    v14 = 18446744073709551615
                                                } else {
                                                    v89 := GenS15{F0: -9223372036854775808, F1: 79}
                                                    if false {
                                                        var v90 int64 = -1
                                                        v90 = v90
                                                        v13 = (("" + v13) + v13)
                                                        _ = v90
                                                    } else {
                                                        v91 := GenS1{F0: v2, F1: v32}
                                                        if ((^(v14)) == (v14 / (((v14 / (((v14 * (v14 - (v14 + v14)))) | 1))) | 1))) {
                                                            var v92 uint8 = (v32 - 104)
                                                            v30 += (v30 >> (((^(v14))) & 63))
                                                            switch 1 {
                                                            case 987:
                                                                v2 += -1
                                                            case 4294967295:
                                                                var v93 float64 = 3.14159
                                                                _ = v93
                                                            default:
                                                                v94 := []int{fi29, (-(fi29))}
                                                                _ = v94
                                                            }
                                                            var v95 int8 = (-(v30))
                                                            _ = v92
                                                            _ = v95
                                                        } else {
                                                            var v96 bool = (!(v2 < v2))
                                                            v97 := GenS8{F0: 198}
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        v2 = v2
                                                        _ = v91
                                                    }
                                                    _ = v89
                                                }
                                            }
                                        }
                                    }
                                    _ = v44
                                    _ = v54
                                } else {
                                    v98 := []int{(fi29 << ((0) & 63))}
                                    _ = v98
                                }
                            case 32767:
                                var v99 uint = v14
                                _ = v99
                            }
                            var v100 uint16 = genFn0("foo bar", v2)
                            var v101 uint64 = 0
                            _ = v100
                            _ = v101
                            _ = v32
                        }
                        switch v13 {
                        case "hello":
                            v14 += v14
                        case "\n\t":
                            v102 := GenS16{F0: 0}
                            _ = v102
                        case "":
                            v14 = 612
                        }
                        v1 = 1
                    }
                }
                var v103 uint64 = 0
                _ = fi29
                _ = v103
                _ = v30
            }
            v104 := GenS17{F0: (v2 | int32(v14)), F1: 849, F2: 2}
            _ = v104
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 = (v1 * (v1 * 127))
    v1 = (v1 + (v1 | (v1 << ((2) & 63))))
    v1 = (v1 + (v1 + v1))
    v105 := GenS10{F0: 65535, F1: "\n\t", F2: 1}
    var v106 string = "\n\t"
    v107 := GenS2{F0: 42, F1: -1}
    v1 -= v1
    var v108 int64 = -9223372036854775808
    v109 := GenS14{F0: (1 < 4294967295), F1: 128}
    fmt.Println(v1)
    if true {
        v110 := []int{v1, v1, (v1 + v1), v1}
        for fi111 := 0; fi111 < 18; fi111++ {
            v112 := GenS7{F0: 42}
            fmt.Println(v1)
            _ = fi111
            _ = v112
        }
        v2 += v2
        _ = v110
    }
    v113 := GenS8{F0: 1}
    v114 := GenS18{F0: (v106 + (v106 + ((v106 + v106) + ((v106 + ((v106 + v106) + v106)) + "\n\t")))), F1: v106, F2: (((v106 + (v106 + v106)) + (v106 + v106)) + v106)}
    v1 -= 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v109
    _ = v113
    _ = v114
    _ = v2
    _ = v3
}
