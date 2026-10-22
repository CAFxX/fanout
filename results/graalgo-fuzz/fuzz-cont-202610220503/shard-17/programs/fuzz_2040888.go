package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 bool; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 bool; F2 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 uint8; F2 uint }
func genFn0() int {
    var v7 int64 = 1
    var v8 int64 = v7
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int8 = 100
            fmt.Println(fi10, v7)
            var v12 uint64 = 0
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 bool = false
            var v16 int16 = -12
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint64, p1 uint8) int32 {
    switch (("a" + "hello") + (("a" + "") + ("a" + "foo bar"))) {
    case "hello":
        p1 = p1
    case "a":
        switch 127 {
        case 42:
            v17 := []int{4294967295}
            _ = v17
        default:
            p0 -= 0
        }
    case "foo bar":
        v18 := GenS1{F0: p0, F1: -676, F2: -2147483648}
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS2{F0: 2, F1: 427, F2: 42}
        var v21 uint = 2
        if (-1 < -1) {
            switch (fi19 / (((fi19 >> (((v21 / ((2) | 1))) & 63))) | 1)) {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := []int{(fi22 - fi19)}
                    _ = fi22
                    _ = v23
                }
            case 2147483647:
                var v24 uint8 = p1
                _ = v24
            }
        } else {
            fmt.Println(p0, fi19)
            var v25 int16 = 1
            p1 += 2
            _ = v25
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        p0 = (p0 | 784)
        if true {
            v27 := GenS2{F0: uint16(p0), F1: 636, F2: 42}
            _ = v27
        } else {
            v28 := GenS3{F0: fi26, F1: (p0 | p0), F2: p1}
            if true {
                switch fi26 {
                case 132:
                    if (42 != 890) {
                        if true {
                            p1 = 128
                            v29 := GenS3{F0: fi26, F1: 0, F2: 1}
                            for fi30 := 0; fi30 < 17; fi30++ {
                                var v31 int = (fi26 + 32)
                                var v32 int32 = 1
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            _ = v29
                        }
                        switch (fi26 - (fi26 - 9223372036854775807)) {
                        case -2147483648:
                            p0 -= (^(p0))
                        case 32768:
                            v33 := GenS1{F0: p0, F1: 0, F2: 2147483647}
                            _ = v33
                        case 2147483647:
                            v34 := []int{fi26}
                            _ = v34
                        default:
                            var v35 int16 = -1
                            _ = v35
                        }
                        if false {
                            p1 = p1
                            p0 -= p0
                        } else {
                            p1 -= p1
                            p0 += 0
                            for fi36 := 0; fi36 < 13; fi36++ {
                                if (!false) {
                                    var v37 string = "\n\t"
                                    _ = v37
                                }
                                p1 -= 128
                                p1 = (p1 | 255)
                                _ = fi36
                            }
                        }
                    }
                }
            }
            _ = v28
        }
        _ = fi26
    }
    p1 += 255
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 int8 = -94
    var v3 int64 = 897
    var v4 int16 = -32768
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 127
    v4 = v4
    _ = len(v5)
    v3 += 116
    if false {
        v4 = -1
    } else {
        v38 := GenS4{F0: 18446744073709551615}
        switch "hello" {
        case "":
            var v39 uint64 = 4294967295
            _ = v39
        default:
            if false {
                v5[v1] = (v1 & (v1 << ((1) & 63)))
                fmt.Println(v2, v4)
            } else {
                switch 655 {
                case 0:
                    if (870 <= 2) {
                        switch (("\n\t" + (("\n\t" + "\n\t") + ("hello" + ("hello" + "hello")))) + "") {
                        case "foo bar":
                            v6 = (^(v6))
                        case "\n\t":
                            v4 -= v4
                        case "hello":
                            v40 := []int{-2147483648}
                            _ = v40
                        default:
                            if true {
                                v2 = (v2 & (v2 - v2))
                            }
                        }
                    } else {
                        v41 := v5[v1]
                        v2 += 1
                        fmt.Println(v2, v3)
                        _ = v41
                    }
                case 128:
                    v42 := GenS2{F0: 32767, F1: v3, F2: v4}
                    _ = v42
                case -2147483648:
                    fmt.Println(v3, v4)
                default:
                    var v43 uint32 = v1
                    _ = v43
                }
                v4 -= 132
            }
        }
        _ = v38
    }
    fmt.Println(v6, v1)
    switch 42 {
    case -264:
        if false {
            if ((v2 % ((v2) | 1)) <= v2) {
                for fi44 := 0; fi44 < 10; fi44++ {
                    var v45 int8 = int8(v1)
                    v46 := []int{fi44, fi44, genFn0(), 1}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                v47 := []int{-9223372036854775808, -1}
                _ = v47
            } else {
                v3 += v3
                v48 := []int{32768, 32768, 254}
                _ = v48
            }
            if false {
                v49 := GenS5{F0: v3, F1: (false || true), F2: 4294967295}
                var v50 uint8 = 127
                fmt.Println(v50)
                _ = v49
                _ = v50
            }
            for fi51 := 0; fi51 < 20; fi51++ {
                var v52 uint16 = 37
                _ = fi51
                _ = v52
            }
        } else {
            if (!(!(v6 >= 2))) {
                var v53 uint8 = v6
                _ = v53
            }
            v1 -= v1
        }
    default:
        for fi54 := 0; fi54 < 17; fi54++ {
            for fi55 := 0; fi55 < 4; fi55++ {
                if true {
                    v56 := GenS2{F0: 1, F1: v3, F2: v4}
                    v3 = v3
                    var v57 uint8 = (v6 ^ 128)
                    _ = v56
                    _ = v57
                } else {
                    if true {
                        v3 -= 898
                    }
                }
                v2 += (v2 - 100)
                v58 := v5[v1]
                _ = fi55
                _ = v58
            }
            var v59 int8 = v2
            v3 = int64(v6)
            _ = fi54
            _ = v59
        }
    }
    if (true || false) {
        v60 := GenS6{F0: "foo bar"}
        if ((v1 / (((v1 % (((^(v1))) | 1))) | 1)) != v1) {
            for fi61 := 0; fi61 < 13; fi61++ {
                fmt.Println(v1, fi61, v6)
                _ = fi61
            }
            if (false || ((!false) || (255 < (v6 | v6)))) {
                switch "foo bar" {
                case "\n\t":
                    var v62 float32 = (float32(((1.5 - -2.25) - 0.1)) - 3.14159)
                    _ = v62
                default:
                    if false {
                        v2 = (v2 / ((v2) | 1))
                        v3 = -307
                        if true {
                            for fi63 := 0; fi63 < 2; fi63++ {
                                var v64 float64 = 0.1
                                _ = fi63
                                _ = v64
                            }
                            v65 := []int{167, int(v1), genFn0()}
                            _ = v65
                        } else {
                            var v66 int = 269
                            fmt.Println(v3, v1)
                            v67 := GenS4{F0: 0}
                            _ = v66
                            _ = v67
                        }
                    } else {
                        _ = len(v5)
                        switch -337 {
                        case 32768:
                            fmt.Println(v1, v4)
                        }
                    }
                }
                for fi68 := 0; fi68 < 4; fi68++ {
                    v69 := []int{}
                    var v70 uint64 = 4294967295
                    var v71 int8 = int8(v4)
                    _ = fi68
                    _ = v69
                    _ = v70
                    _ = v71
                }
                v72 := v5[529]
                _ = v72
            }
        }
        _ = v60
    }
    var v73 uint32 = (^(v1))
    v74 := GenS3{F0: 9223372036854775807, F1: 1, F2: v6}
    v6 += (v6 * v6)
    for fi75 := 0; fi75 < 18; fi75++ {
        if false {
            var v76 uint64 = 42
            delete(v5, v1)
            _ = v76
        }
        _ = fi75
    }
    v77 := []int{-940, 4294967295, 988, -847}
    switch 128 {
    case 32767:
        v2 -= v2
    case -468:
        v78 := GenS6{F0: "hello"}
        _ = v78
    }
    v6 -= (v6 & v6)
    if len(v77) > 0 {
        v77[0] = 16
    }
    if true {
        switch 9223372036854775807 {
        case 1:
            v4 -= v4
        case -2147483648:
            for fi79 := 0; fi79 < 18; fi79++ {
                for fi80 := 0; fi80 < 6; fi80++ {
                    v81 := GenS7{F0: (v4 >> ((0) & 63)), F1: ((((v3 | v3) > 9223372036854775807) && false) || (-2.25 < float64(0.1))), F2: v2}
                    var v82 uint32 = (v73 - v73)
                    v83 := GenS3{F0: fi79, F1: 4294967295, F2: (v6 ^ v6)}
                    _ = fi80
                    _ = v81
                    _ = v82
                    _ = v83
                }
                v4 += (^(v4))
                _ = fi79
            }
        }
    } else {
        v6 -= v6
        if ((v4 | (v4 / ((1000) | 1))) != v4) {
            v4 += (v4 << ((2) & 63))
            for fi84 := 0; fi84 < 2; fi84++ {
                if false {
                    if true {
                        v85 := GenS8{F0: ""}
                        if len(v77) >= 2 {
                            v77 = v77[1:len(v77)]
                        }
                        if false {
                            fmt.Println(v73, v4)
                        } else {
                            if len(v77) > 0 {
                                v77[0] = fi84
                            }
                        }
                        _ = v85
                    } else {
                        v1 -= (v73 << ((678) & 63))
                        v73 += v73
                    }
                    for fi86 := 0; fi86 < 16; fi86++ {
                        var v87 int8 = int8(v4)
                        var v88 uint = 4294967295
                        _ = fi86
                        _ = v87
                        _ = v88
                    }
                } else {
                    fmt.Println(v4)
                }
                _ = fi84
            }
        } else {
            var v89 float32 = 0.1
            var v90 int = 32767
            _ = v89
            _ = v90
        }
        for fi91 := 0; fi91 < 8; fi91++ {
            var v92 uint16 = 87
            _ = fi91
            _ = v92
        }
    }
    switch 128 {
    case 415:
        if true {
            fmt.Println(v1)
            if true {
                if (111 != 180) {
                    if (23 == (v2 % ((27) | 1))) {
                        v93 := GenS1{F0: 1, F1: -668, F2: 2147483647}
                        v94 := GenS7{F0: 347, F1: (!true), F2: v2}
                        _ = v93
                        _ = v94
                    }
                } else {
                    v1 += (v73 + (v1 % ((v1) | 1)))
                    v73 = (v73 >> ((65535) & 63))
                }
            }
            if false {
                var v95 bool = (-9223372036854775808 > -844)
                v73 = v1
                var v96 uint16 = 65535
                _ = v95
                _ = v96
            }
        } else {
            fmt.Println(v1, v4)
            v97 := GenS9{F0: v2, F1: v6, F2: 18446744073709551615}
            fmt.Println(v73, v2, v6)
            _ = v97
        }
    case 128:
        if false {
            v98 := GenS5{F0: (^(v3)), F1: false, F2: 0}
            var v99 uint = 42
            _ = v98
            _ = v99
        }
    case 42:
        switch -2147483648 {
        case 32767:
            if false {
                v100 := GenS3{F0: -2147483648, F1: 42, F2: 10}
                fmt.Println(v6, v2, v3)
                _ = v100
            } else {
                for fi101 := 0; fi101 < 17; fi101++ {
                    v6 += uint8(v1)
                    var v102 uint64 = 2
                    var v103 float32 = (0.1 + ((float32(0.0) + 1.5) - 0.1))
                    _ = fi101
                    _ = v102
                    _ = v103
                }
                for fi104 := 0; fi104 < 12; fi104++ {
                    if false {
                        _ = len(v77)
                        if (1 == 127) {
                            var v105 int = int(v4)
                            _ = v105
                        }
                    }
                    fmt.Println(v4, v6, v3)
                    _ = fi104
                }
                v4 = int16(v6)
            }
        case -1:
            _ = len(v77)
        }
    }
    var v106 int8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v77:", len(v77), v77)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v77
}
