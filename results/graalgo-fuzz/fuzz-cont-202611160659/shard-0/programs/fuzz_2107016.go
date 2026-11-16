package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint16; F1 int; F2 uint8 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 string; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 bool; F2 string }
type GenS11 struct { F0 int16; F1 uint8; F2 int }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int64; F1 int; F2 uint }
func genFn0() string {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 221}
        v8 := GenS2{F0: 324, F1: 0}
        if false {
            var v9 int8 = -128
            v10 := GenS3{F0: -61, F1: (v9 % (((-(v9))) | 1)), F2: -76}
            fmt.Println(fi6, v9)
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -1 {
    case -477:
        var v11 uint16 = 0
        _ = v11
    case 42:
        var v12 uint8 = 42
        _ = v12
    }
    return "foo bar"
}
func genFn1(p0 int64, p1 uint64) int64 {
    v13 := GenS3{F0: 127, F1: -1, F2: -1}
    switch -606 {
    case 32768:
        switch 32768 {
        case 32767:
            v14 := []int{-9223372036854775808, 0, -919, 2147483647}
            _ = v14
        }
    case -700:
        for fi15 := 0; fi15 < 11; fi15++ {
            p0 = -45
            _ = fi15
        }
    case 742:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v13
    return -1
}

func main() {
    var v1 string = ("hello" + "hello")
    var v2 bool = (false || (0.1 >= 0.1))
    var v3 uint = 18446744073709551615
    var v4 int = 2147483647
    v5 := make(map[uint32]bool)
    if v2 {
        v3 = v3
    } else {
        var v16 int64 = 947
        var v17 int64 = genFn1((v16 >> ((v3) & 63)), 980)
        _ = v16
        _ = v17
    }
    var v18 int32 = -2147483648
    switch v4 {
    case -9223372036854775808:
        v3 += v3
    }
    var v19 int = v4
    fmt.Println(v2)
    if true {
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 16; fi21++ {
                v19 -= (^(fi21))
                if (!v2) {
                    switch (^(fi20)) {
                    case 0:
                        if false {
                            v4 = fi20
                            var v22 uint = (v3 ^ v3)
                            _ = v22
                        } else {
                            fmt.Println(fi20)
                            switch "a" {
                            case "\n\t":
                                if (-109 != -53) {
                                    fmt.Println(fi21, v3, v3)
                                    fmt.Println(v3, v18)
                                }
                            }
                            v2 = true
                        }
                    }
                } else {
                    if (9223372036854775807 < 1) {
                        v23 := v5[4294967295]
                        fmt.Println(v1, v18, v2)
                        v24 := []int{}
                        _ = v23
                        _ = v24
                    } else {
                        if ((v4 <= (^(v19))) || false) {
                            switch ((v1 + v1) + (("hello" + (("a" + "hello") + v1)) + "a")) {
                            case "":
                                if ((^(v18)) < v18) {
                                    var v25 int32 = (v18 | (v18 & 600))
                                    var v26 int8 = -1
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v27 uint = v3
                                    var v28 bool = true
                                    var v29 int16 = 753
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            case "hello":
                                var v30 int64 = -9223372036854775808
                                _ = v30
                            case "\n\t":
                                v19 = v19
                            default:
                                var v31 int16 = 32767
                                _ = v31
                            }
                        } else {
                            var v32 uint64 = 1
                            if v2 {
                                if (!v2) {
                                    if false {
                                        v33 := GenS4{F0: 2, F1: false, F2: v32}
                                        v1 = v1
                                        var v34 uint32 = 58
                                        _ = v33
                                        _ = v34
                                    } else {
                                        if false {
                                            var v35 uint8 = 0
                                            v4 -= (v4 ^ (-(v4)))
                                            fmt.Println(fi20)
                                            _ = v35
                                        }
                                        fmt.Println(v2, v1)
                                        if (!(v1 <= ((("\n\t" + "\n\t") + "") + v1))) {
                                            switch -9223372036854775808 {
                                            case 1:
                                                switch v1 {
                                                case "a":
                                                    var v36 uint16 = 681
                                                    _ = v36
                                                default:
                                                    if ((326 < 65535) || (v2 && (v19 == fi21))) {
                                                        v19 += (fi20 - (v19 ^ fi20))
                                                    } else {
                                                        v37 := []int{0, fi20, v19}
                                                        v38 := GenS2{F0: (^(v32)), F1: 178}
                                                        v4 = 0
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                            }
                                            v39 := []int{int(v3)}
                                            _ = v39
                                        } else {
                                            v40 := GenS5{F0: 274, F1: genFn1(-1, 4294967295)}
                                            v41 := GenS6{F0: 65535, F1: 9223372036854775807, F2: 2}
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                }
                                v42 := GenS3{F0: 127, F1: 4, F2: -1}
                                v43 := GenS3{F0: 100, F1: 100, F2: 1}
                                _ = v42
                                _ = v43
                            } else {
                                var v44 uint16 = 1
                                fmt.Println(v44, fi20, v18)
                                _ = v44
                            }
                            _ = v32
                        }
                        v45 := GenS7{F0: fi21, F1: 127}
                        v3 += (v3 * 495)
                        _ = v45
                    }
                    v1 = v1
                    var v46 string = v1
                    _ = v46
                }
                _ = fi21
            }
            v47 := GenS8{F0: v1, F1: 897, F2: 42}
            _ = fi20
            _ = v47
        }
        if (0.1 > 0.0) {
            delete(v5, 42)
            v18 = (v18 & (v18 ^ (v18 << (((v3 * (v3 - (v3 >> (((^(v3))) & 63))))) & 63))))
        }
    } else {
        if v2 {
            v18 -= v18
        } else {
            var v48 int16 = 0
            _ = v48
        }
    }
    if (1 <= 1) {
        var v49 uint16 = 524
        for fi50 := 0; fi50 < 18; fi50++ {
            switch v1 {
            case "foo bar":
                fmt.Println(v3)
            case "\n\t":
                v19 += (fi50 | v4)
            }
            if (1 < 1000) {
                var v51 float32 = 3.14159
                _ = v51
            } else {
                v19 = (v19 % ((v4) | 1))
                for fi52 := 0; fi52 < 5; fi52++ {
                    v53 := []int{v19, (^(fi52))}
                    fmt.Println(v19)
                    _ = fi52
                    _ = v53
                }
            }
            _ = fi50
        }
        v19 += v19
        _ = v49
    } else {
        v54 := GenS9{F0: 811}
        if (v2 && (false || (!(367 == 1)))) {
            for fi55 := 0; fi55 < 20; fi55++ {
                if (2147483647 >= 84) {
                    var v56 int64 = 1
                    _ = v56
                } else {
                    if v2 {
                        switch (v1 + v1) {
                        case "a":
                            if true {
                                var v57 bool = true
                                switch v19 {
                                case 32768:
                                    var v58 uint64 = 30
                                    _ = v58
                                case -9223372036854775808:
                                    v18 += v18
                                case -1:
                                    var v59 bool = (!false)
                                    _ = v59
                                }
                                _ = v57
                            } else {
                                v1 = "hello"
                            }
                        case "":
                            var v60 int16 = 1000
                            _ = v60
                        case "\n\t":
                            v61 := []int{1, -2147483648}
                            _ = v61
                        }
                        switch (v1 + v1) {
                        case "a":
                            switch v1 {
                            case "hello":
                                var v62 uint64 = 465
                                _ = v62
                            case "a":
                                for fi63 := 0; fi63 < 3; fi63++ {
                                    fmt.Println(v19, v19)
                                    _ = fi63
                                }
                            case "foo bar":
                                for fi64 := 0; fi64 < 7; fi64++ {
                                    fmt.Println(v18)
                                    if (0.1 >= 3.14159) {
                                        if (true || (1e10 == 1.5)) {
                                            v3 = (^(v3))
                                            var v65 int = v4
                                            _ = v65
                                        } else {
                                            if v2 {
                                                v66 := GenS9{F0: 0}
                                                if (42 == 714) {
                                                    if (0 > 146) {
                                                        var v67 uint = (v3 & 2)
                                                        _ = v67
                                                    } else {
                                                        delete(v5, 1)
                                                        v19 = v4
                                                    }
                                                    switch (v19 << ((v3) & 63)) {
                                                    case -872:
                                                        v68 := v5[2147483647]
                                                        _ = v68
                                                    case 9223372036854775807:
                                                        v19 = v4
                                                    default:
                                                        fmt.Println(v18, fi64)
                                                    }
                                                    v69 := GenS1{F0: 255}
                                                    _ = v69
                                                } else {
                                                    var v70 float64 = (0.1 * 1e10)
                                                    v70 -= v70
                                                    delete(v5, 0)
                                                    v3 = v3
                                                    _ = v70
                                                }
                                                _ = v66
                                            }
                                            delete(v5, 2147483648)
                                            var v71 int8 = 87
                                            switch v4 {
                                            case 42:
                                                v18 = (v18 >> ((v3) & 63))
                                            }
                                            _ = v71
                                        }
                                        var v72 uint = v3
                                        _ = v72
                                    }
                                    _ = fi64
                                }
                            }
                        case "foo bar":
                            v73 := GenS9{F0: 42}
                            _ = v73
                        default:
                            v1 = v1
                        }
                    } else {
                        v74 := GenS1{F0: 1}
                        v19 = (v19 % (((^(v19))) | 1))
                        _ = v74
                    }
                    v75 := []int{v4, (fi55 ^ -2147483648), (v4 >> ((v3) & 63)), fi55}
                    _ = v75
                }
                v3 -= v3
                _ = fi55
            }
        } else {
            var v76 string = "\n\t"
            if (255 == 65535) {
                v77 := GenS9{F0: 337}
                _ = v77
            }
            _ = v76
        }
        var v78 uint8 = 0
        _ = v54
        _ = v78
    }
    v19 += (-(v4))
    for fi79 := 0; fi79 < 13; fi79++ {
        var v80 int = (v4 / ((fi79) | 1))
        var v81 int64 = genFn1(1, 18446744073709551615)
        v82 := GenS10{F0: 1, F1: (genFn1((v81 - (^(v81))), 815) < (v81 | 708)), F2: v1}
        _ = fi79
        _ = v80
        _ = v81
        _ = v82
    }
    v1 = ("foo bar" + v1)
    v2 = (32768 > (v19 >> ((v3) & 63)))
    var v83 int64 = -1
    _ = len(v5)
    for fi84 := 0; fi84 < 1; fi84++ {
        var v85 int16 = 32767
        v85 = (v85 ^ -1)
        var v86 float32 = 1e10
        _ = fi84
        _ = v85
        _ = v86
    }
    v87 := GenS11{F0: 943, F1: 218, F2: v19}
    if (32767 <= 42) {
        if v2 {
            v2 = (127 != 42)
            v1 = "foo bar"
            if true {
                fmt.Println(v83, v19)
                var v88 float32 = ((-2.25 / float32(float64(((0.1 - 0.0) / 1.5)))) - 0.1)
                _ = v88
            } else {
                switch "a" {
                case "foo bar":
                    v83 -= v83
                case "hello":
                    for fi89 := 0; fi89 < 11; fi89++ {
                        var v90 uint8 = 67
                        var v91 float64 = (0.0 - 1e10)
                        fmt.Println(v4, v2, v83)
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                }
                if (0 != 2) {
                    var v92 int32 = v18
                    v92 = (v18 - v92)
                    _ = v92
                } else {
                    v18 += (v18 & v18)
                    if false {
                        for fi93 := 0; fi93 < 13; fi93++ {
                            var v94 float64 = float64(1e10)
                            if ((v2 || (false && (v2 || ((!((^(v3)) < v3)) && true)))) && (true && (((-(v4)) == (v4 * -352)) && v2))) {
                                fmt.Println(v2, v18, v1)
                                v95 := GenS6{F0: 65535, F1: v4, F2: 35}
                                v83 -= -1
                                _ = v95
                            }
                            _ = fi93
                            _ = v94
                        }
                    }
                }
            }
        }
        v2 = v2
    } else {
        for fi96 := 0; fi96 < 19; fi96++ {
            if (831 >= 139) {
                for fi97 := 0; fi97 < 11; fi97++ {
                    if false {
                        var v98 float64 = -2.25
                        _ = v98
                    }
                    var v99 uint32 = 1
                    v100 := []int{}
                    _ = fi97
                    _ = v100
                    _ = v99
                }
                v1 = v1
                fmt.Println(fi96, v19)
            }
            v2 = (v2 && (v2 && (!((fi96 ^ (v19 / ((32767) | 1))) <= (v19 / ((v19) | 1))))))
            if true {
                v101 := GenS12{F0: 65535}
                if v2 {
                    for fi102 := 0; fi102 < 2; fi102++ {
                        v103 := []int{(^(fi102)), v19}
                        _ = fi102
                        _ = v103
                    }
                    switch fi96 {
                    case -903:
                        switch v1 {
                        case "foo bar":
                            v19 -= (v4 & v19)
                        case "a":
                            fmt.Println(v83, v3, v19)
                        case "hello":
                            v104 := []int{(v4 & v19), v19, fi96}
                            _ = v104
                        }
                    case 2147483647:
                        v83 = genFn1(genFn1(0, 0), 225)
                    default:
                        var v105 int8 = -84
                        _ = v105
                    }
                }
                _ = v101
            } else {
                v106 := v5[1]
                v107 := GenS13{F0: genFn1(-637, 18446744073709551615), F1: 899, F2: 0}
                _ = v106
                _ = v107
            }
            _ = fi96
        }
        for fi108 := 0; fi108 < 15; fi108++ {
            v4 -= -2147483648
            v5[0] = (!(v18 < 0))
            _ = fi108
        }
        v19 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v87
}
