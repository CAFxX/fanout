package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }
type GenS4 struct { F0 int8; F1 int; F2 bool }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int8; F2 int32 }
type GenS8 struct { F0 int16; F1 uint8; F2 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64; F1 int64; F2 uint32 }
type GenS11 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 bool, p1 uint16) int16 {
    var v5 uint64 = 42
    v6 := GenS1{F0: ("foo bar" + "\n\t")}
    switch (("foo bar" + "hello") + "") {
    case "\n\t":
        var v7 int32 = -2147483648
        _ = v7
    case "hello":
        v8 := GenS2{F0: (p1 > p1), F1: 127}
        _ = v8
    }
    v9 := GenS3{F0: -128, F1: 1, F2: -832}
    var v10 int64 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v6
    _ = v9
    return 32767
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := []int{fi11, (-(fi11)), fi11, 2147483647}
        v13 := GenS3{F0: -108, F1: 42, F2: -108}
        if true {
            var v14 uint32 = 2
            if (0 == 65535) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    switch "\n\t" {
                    case "hello":
                        var v16 int64 = -678
                        _ = v16
                    case "a":
                        v14 += v14
                    }
                    switch ("" + "") {
                    case "":
                        v17 := GenS3{F0: -128, F1: 104, F2: 42}
                        _ = v17
                    case "hello":
                        v14 = v14
                    case "\n\t":
                        if (!true) {
                            fmt.Println(v14, fi15)
                            var v18 int = (fi15 << ((65535) & 63))
                            if true {
                                var v19 bool = (-32768 < genFn0(true, 362))
                                _ = v19
                            }
                            _ = len(v12)
                            fmt.Println(v14)
                            _ = v18
                        }
                    default:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    }
                    var v20 uint = 382
                    _ = fi15
                    _ = v20
                }
            } else {
                fmt.Println(fi11, fi11)
            }
            _ = v14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    switch "a" {
    case "":
        if ((1e10 / 1e10) >= -2.25) {
            v21 := GenS2{F0: true, F1: 32768}
            _ = v21
        }
    case "foo bar":
        var v22 uint64 = 2
        _ = v22
    case "a":
        var v23 float64 = 0.0
        _ = v23
    default:
        var v24 float32 = (1e10 + -2.25)
        _ = v24
    }
    var v25 int32 = 1
    _ = v25
    return 399
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 0
    var v3 int32 = 444
    v4 := make(map[uint32]bool)
    var v26 int = 4294967295
    v27 := GenS3{F0: 9, F1: 183, F2: 32767}
    v2 = (v2 << ((255) & 63))
    var v28 uint32 = v2
    v4[v2] = false
    if (!true) {
        if (!false) {
            fmt.Println(v2, v28)
        } else {
            var v29 int8 = 100
            v30 := GenS2{F0: ((1 <= 1) && (!true)), F1: (v26 >> ((2) & 63))}
            _ = v29
            _ = v30
        }
        fmt.Println(v26)
    }
    if ((0 != 0) || (!(0 != 1))) {
        if false {
            v31 := GenS4{F0: -1, F1: (^(v26)), F2: false}
            _ = v31
        } else {
            if false {
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 int64 = -9223372036854775808
                    if (!(1e10 < (0.0 * 1e10))) {
                        for fi34 := 0; fi34 < 18; fi34++ {
                            if false {
                                delete(v1, 2147483647)
                            } else {
                                if true {
                                    if false {
                                        fmt.Println(v33)
                                    } else {
                                        fmt.Println(v33, v26, v2)
                                    }
                                }
                                var v35 uint8 = 127
                                v36 := GenS1{F0: "hello"}
                                _ = v35
                                _ = v36
                            }
                            var v37 int = (fi32 / ((608) | 1))
                            if false {
                                var v38 uint16 = 682
                                _ = v38
                            } else {
                                v39 := GenS5{F0: -32768, F1: ((1e10 / 0.1) < -2.25)}
                                if false {
                                    if false {
                                        fmt.Println(v2)
                                        var v40 int16 = 42
                                        v2 = (v28 & (^(v28)))
                                        v2 = (v2 * v2)
                                        _ = v40
                                    }
                                    fmt.Println(v37, v33)
                                }
                                var v41 uint16 = 2
                                _ = v39
                                _ = v41
                            }
                            _ = fi34
                            _ = v37
                        }
                        var v42 string = "foo bar"
                        _ = v42
                    } else {
                        v43 := []int{v26, (v26 / ((fi32) | 1)), (fi32 << ((65535) & 63))}
                        var v44 int16 = 722
                        _ = v43
                        _ = v44
                    }
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v28, v3)
            }
            fmt.Println(v2, v28, v2)
        }
    }
    v45 := GenS6{F0: 42}
    var v46 int8 = int8(v26)
    v47 := GenS7{F0: 64, F1: v46, F2: (v3 / ((1) | 1))}
    _ = len(v4)
    switch "foo bar" {
    case "":
        v48 := []int{v26, -48, (v26 - 9223372036854775807)}
        _ = v48
    case "hello":
        if (!false) {
            v49 := v4[v2]
            _ = v49
        } else {
            var v50 bool = true
            v51 := GenS3{F0: (v46 + v46), F1: 42, F2: int16(v46)}
            _ = v50
            _ = v51
        }
    default:
        if true {
            switch "" {
            case "a":
                v28 = v28
            case "hello":
                _ = len(v4)
            case "":
                v52 := GenS5{F0: -73, F1: (!(int64(v46) == -9223372036854775808))}
                _ = v52
            default:
                var v53 uint = 42
                _ = v53
            }
            if (!(!(v46 == (^(v46))))) {
                v46 = (v46 & 127)
                for fi54 := 0; fi54 < 14; fi54++ {
                    if false {
                        for fi55 := 0; fi55 < 5; fi55++ {
                            switch ("hello" + ("hello" + ("\n\t" + "a"))) {
                            case "a":
                                if ((-2147483648 < (v3 % ((-1) | 1))) || false) {
                                    v28 += (^(v2))
                                    fmt.Println(fi54)
                                }
                            default:
                                fmt.Println(v46, v3, v2)
                            }
                            v46 -= v46
                            var v56 uint8 = 255
                            _ = fi55
                            _ = v56
                        }
                    } else {
                        v57 := []int{v26, int(v3), (fi54 / ((9223372036854775807) | 1)), fi54}
                        fmt.Println(v26, v2)
                        _ = v57
                    }
                    v58 := []int{fi54, v26, (v26 / ((fi54) | 1)), fi54}
                    _ = fi54
                    _ = v58
                }
                if (true || (0 != 2)) {
                    for fi59 := 0; fi59 < 1; fi59++ {
                        v60 := []int{}
                        _ = fi59
                        _ = v60
                    }
                    v2 = (v28 | v28)
                    v61 := GenS1{F0: "a"}
                    _ = v61
                } else {
                    _ = len(v4)
                    switch v26 {
                    case -510:
                        v62 := GenS7{F0: 47, F1: v46, F2: (v3 / ((v3) | 1))}
                        _ = v62
                    default:
                        if false {
                            v2 -= 41
                            v26 += v26
                            var v63 uint32 = (v2 / (((^(v28))) | 1))
                            _ = v63
                        } else {
                            v46 -= (v46 >> ((4294967295) & 63))
                            if false {
                                for fi64 := 0; fi64 < 10; fi64++ {
                                    for fi65 := 0; fi65 < 10; fi65++ {
                                        fmt.Println(fi65, v46)
                                        var v66 bool = false
                                        fmt.Println(v2)
                                        v67 := GenS8{F0: -1, F1: 147, F2: true}
                                        var v68 int64 = -1
                                        _ = fi65
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                    var v69 uint16 = 18
                                    _ = fi64
                                    _ = v69
                                }
                                v3 = v3
                            }
                        }
                    }
                    for fi70 := 0; fi70 < 6; fi70++ {
                        for fi71 := 0; fi71 < 3; fi71++ {
                            var v72 uint16 = 0
                            _ = fi71
                            _ = v72
                        }
                        _ = fi70
                    }
                }
            } else {
                switch "hello" {
                case "hello":
                    if true {
                        var v73 int = 32767
                        var v74 string = "foo bar"
                        _ = v73
                        _ = v74
                    } else {
                        v75 := GenS7{F0: 32767, F1: (v46 | (v46 & 127)), F2: (v3 | (v3 << ((42) & 63)))}
                        v46 -= v46
                        if (!false) {
                            for fi76 := 0; fi76 < 16; fi76++ {
                                v46 = (v46 % (((v46 * (v46 % ((0) | 1)))) | 1))
                                var v77 uint8 = 42
                                if false {
                                    var v78 int8 = v46
                                    v2 = (^(v2))
                                    for fi79 := 0; fi79 < 13; fi79++ {
                                        v80 := GenS2{F0: false, F1: 2147483647}
                                        fmt.Println(v26, v46)
                                        _ = fi79
                                        _ = v80
                                    }
                                    _ = v78
                                }
                                _ = fi76
                                _ = v77
                            }
                            v81 := GenS9{F0: int32(v28)}
                            var v82 string = "foo bar"
                            _ = v81
                            _ = v82
                        }
                        _ = v75
                    }
                case "foo bar":
                    for fi83 := 0; fi83 < 19; fi83++ {
                        var v84 uint32 = v28
                        v26 -= 32768
                        _ = fi83
                        _ = v84
                    }
                case "a":
                    var v85 uint8 = 106
                    _ = v85
                }
                fmt.Println(v2, v26)
                if true {
                    var v86 uint32 = (v28 - 2)
                    if true {
                        for fi87 := 0; fi87 < 2; fi87++ {
                            if false {
                                v88 := GenS10{F0: -9223372036854775808, F1: 42, F2: (v86 + 2147483647)}
                                fmt.Println(v26, v28, v3)
                                var v89 float64 = 1.5
                                _ = v88
                                _ = v89
                            }
                            var v90 int32 = (^(v3))
                            var v91 float64 = 1e10
                            _ = fi87
                            _ = v90
                            _ = v91
                        }
                    } else {
                        for fi92 := 0; fi92 < 3; fi92++ {
                            switch fi92 {
                            case 42:
                                v93 := GenS4{F0: (^(v46)), F1: (-(v26)), F2: (9223372036854775807 >= 9223372036854775807)}
                                _ = v93
                            case 127:
                                if true {
                                    v94 := GenS11{F0: 309, F1: fi92, F2: "foo bar"}
                                    if (true || true) {
                                        var v95 int64 = 0
                                        _ = v95
                                    }
                                    v28 = 0
                                    _ = v94
                                } else {
                                    fmt.Println(fi92)
                                    v86 = v86
                                }
                            default:
                                if (!false) {
                                    fmt.Println(v26)
                                    v96 := GenS5{F0: int16(v46), F1: (42 != v2)}
                                    v97 := []int{42}
                                    _ = v96
                                    _ = v97
                                } else {
                                    v26 -= (fi92 % (((fi92 | (v26 * -702))) | 1))
                                    switch fi92 {
                                    case 9223372036854775807:
                                        for fi98 := 0; fi98 < 1; fi98++ {
                                            if true {
                                                fmt.Println(v46, v2)
                                                v1[v86] = v86
                                            } else {
                                                v99 := GenS8{F0: 1, F1: 2, F2: (!true)}
                                                v100 := GenS7{F0: 65535, F1: (v46 + v46), F2: (v3 % (((^(v3))) | 1))}
                                                v101 := GenS5{F0: -32768, F1: true}
                                                _ = v100
                                                _ = v101
                                                _ = v99
                                            }
                                            _ = fi98
                                        }
                                    case 775:
                                        fmt.Println(v2, v46, fi92)
                                    }
                                    fmt.Println(v26, v2, v86)
                                }
                            }
                            _ = fi92
                        }
                        for fi102 := 0; fi102 < 14; fi102++ {
                            delete(v1, (v28 * (v86 / ((v28) | 1))))
                            fmt.Println(v28, v28)
                            _ = fi102
                        }
                        v26 += -816
                    }
                    fmt.Println(v46, v86)
                    _ = v86
                } else {
                    var v103 uint32 = v28
                    fmt.Println(v26, v26, v26)
                    _ = v103
                }
            }
            for fi104 := 0; fi104 < 7; fi104++ {
                v46 -= v46
                for fi105 := 0; fi105 < 7; fi105++ {
                    var v106 string = "\n\t"
                    _ = fi105
                    _ = v106
                }
                var v107 float64 = float64(((0.1 - 1e10) / 0.1))
                _ = fi104
                _ = v107
            }
        } else {
            fmt.Println(v46)
        }
    }
    v1[(^(v2))] = 2
    for fi108 := 0; fi108 < 6; fi108++ {
        var v109 int64 = -1
        _ = fi108
        _ = v109
    }
    v110 := GenS3{F0: v46, F1: 0, F2: 1000}
    v111 := GenS8{F0: genFn0((0.1 != -2.25), 741), F1: 0, F2: true}
    v112 := []int{(v26 + v26), v26, v26}
    fmt.Println(v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v112:", len(v112), v112)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
