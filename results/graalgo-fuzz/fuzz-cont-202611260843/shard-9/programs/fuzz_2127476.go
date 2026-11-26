package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 bool; F1 int16; F2 bool }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint64; F2 int64 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 bool; F1 int32; F2 bool }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 uint32; F1 bool; F2 uint32 }
func genFn0() string {
    v7 := []int{-738, 128, 212, 32768}
    var v8 uint64 = 2
    var v9 uint64 = v8
    switch 32767 {
    case -9223372036854775808:
        if (-977 <= 1) {
            if false {
                switch 42 {
                case 0:
                    if (!(-2.25 <= 0.0)) {
                        v10 := GenS1{F0: 42, F1: -1, F2: v9}
                        v9 -= v8
                        _ = v10
                    }
                }
                v8 = (^(v9))
                var v11 uint = 966
                _ = v11
            } else {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v9, v8)
                case "hello":
                    v9 += 1
                default:
                    v8 -= v9
                }
            }
            var v12 int32 = 42
            _ = v12
        }
    default:
        fmt.Println(v8, v9)
    }
    var v13 float32 = 1e10
    _ = v13
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1() int8 {
    v14 := GenS1{F0: 88, F1: 421, F2: 42}
    if true {
        var v15 float32 = (-2.25 * float32(1.5))
        v15 = (((v15 / 0.1) + float32(((0.1 - (0.0 / ((1e10 - 1e10) + 3.14159))) + -2.25))) - v15)
        for fi16 := 0; fi16 < 20; fi16++ {
            if false {
                var v17 int16 = -32768
                _ = v17
            }
            _ = fi16
        }
        _ = v15
    }
    var v18 float32 = 0.1
    v18 = (0.1 / v18)
    _ = v14
    _ = v18
    return -127
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = 29
    var v3 bool = true
    var v4 int = -137
    var v5 float64 = (-2.25 + -2.25)
    v6 := make(map[int]int)
    if (128 < 182) {
        switch v4 {
        case 0:
            v4 += v4
        }
    } else {
        if (0 == 2147483648) {
            switch 1 {
            case 0:
                if (v3 || (1000 != -1)) {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        for fi20 := 0; fi20 < 8; fi20++ {
                            v3 = ("\n\t" <= "foo bar")
                            _ = fi20
                        }
                        if (!(2 == 255)) {
                            var v21 int64 = v1
                            v22 := GenS2{F0: 0, F1: (-(v4)), F2: 0}
                            v23 := GenS1{F0: 2147483647, F1: int64(v4), F2: 681}
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS1{F0: fi19, F1: (v1 / (((v1 >> ((4294967295) & 63))) | 1)), F2: 4294967295}
                        _ = fi19
                        _ = v24
                    }
                    fmt.Println(v4, v5)
                }
            }
            fmt.Println(v1, v5)
            fmt.Println(v3)
        } else {
            for fi25 := 0; fi25 < 2; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v3 = v3
                    if false {
                        if (v3 && v3) {
                            v3 = true
                        } else {
                            v27 := GenS3{F0: 807, F1: (1000 != -32768), F2: 208}
                            _ = v27
                        }
                        fmt.Println(v1, v4)
                        fmt.Println(fi25, v1)
                    } else {
                        var v28 float32 = (3.14159 + 1.5)
                        v29 := []int{fi25}
                        if (!v3) {
                            var v30 string = ((("\n\t" + "") + ((("a" + "hello") + "\n\t") + "\n\t")) + ("foo bar" + ""))
                            v31 := GenS1{F0: fi25, F1: -9223372036854775808, F2: 50}
                            _ = v30
                            _ = v31
                        }
                        _ = v28
                        _ = v29
                    }
                    v3 = (v3 || (2 > 255))
                    _ = fi26
                }
                v5 -= float64(0.1)
                v32 := GenS3{F0: 959, F1: (!(!((true && v3) || v3))), F2: 177}
                _ = fi25
                _ = v32
            }
        }
        var v33 uint = 2
        switch (^(v4)) {
        case 660:
            v3 = (1e10 == (0.0 + -2.25))
        case -2147483648:
            var v34 uint32 = 2147483647
            _ = v34
        default:
            switch "\n\t" {
            case "\n\t":
                var v35 int64 = v1
                _ = v35
            case "hello":
                if true {
                    v36 := GenS4{F0: 65535, F1: 42}
                    v37 := GenS3{F0: (^(v33)), F1: v3, F2: 22}
                    if false {
                        v38 := GenS2{F0: 0, F1: 266, F2: (v1 << (((v33 / ((v33) | 1))) & 63))}
                        v39 := []int{900}
                        _ = v38
                        _ = v39
                    } else {
                        var v40 float32 = 1e10
                        _ = v40
                    }
                    _ = v36
                    _ = v37
                } else {
                    var v41 string = ""
                    v2 -= genFn1()
                    _ = v41
                }
            default:
                v3 = true
            }
        }
        _ = v33
    }
    if true {
        if (1 >= 0) {
            var v42 int8 = genFn1()
            var v43 int16 = 32767
            if true {
                var v44 uint64 = 928
                if v3 {
                    var v45 int32 = -1
                    for fi46 := 0; fi46 < 2; fi46++ {
                        v47 := GenS1{F0: (fi46 ^ 4294967295), F1: (-(v1)), F2: (v44 % ((v44) | 1))}
                        var v48 float32 = (0.1 / 1.0)
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    var v49 uint8 = 42
                    _ = v45
                    _ = v49
                }
                v50 := GenS4{F0: 2, F1: v44}
                _ = v44
                _ = v50
            } else {
                v51 := GenS3{F0: 42, F1: (v3 || (v3 || false)), F2: 1}
                v52 := GenS5{F0: v3, F1: (v43 * v43), F2: v3}
                if false {
                    var v53 uint = 4294967295
                    v54 := GenS2{F0: 65535, F1: v4, F2: (v1 * (v1 % (((v1 << (((v53 + v53)) & 63))) | 1)))}
                    v55 := GenS3{F0: 255, F1: (4294967295 >= 42), F2: 1}
                    _ = v53
                    _ = v54
                    _ = v55
                } else {
                    if v3 {
                        fmt.Println(v43, v42)
                        switch "hello" {
                        case "hello":
                            v1 -= v1
                        default:
                            v1 = v1
                        }
                    }
                }
                _ = v51
                _ = v52
            }
            _ = v42
            _ = v43
        } else {
            for fi56 := 0; fi56 < 12; fi56++ {
                v57 := GenS6{F0: 2147483647, F1: 930}
                var v58 uint8 = 255
                v1 = -1
                _ = fi56
                _ = v57
                _ = v58
            }
            v59 := GenS4{F0: 4294967295, F1: 18446744073709551615}
            if v3 {
                switch "\n\t" {
                case "hello":
                    v60 := GenS7{F0: (v2 > (^(v2))), F1: 18446744073709551615, F2: v1}
                    _ = v60
                case "\n\t":
                    var v61 int64 = (v1 | (-(v1)))
                    _ = v61
                default:
                    v2 -= 127
                }
                for fi62 := 0; fi62 < 2; fi62++ {
                    var v63 float64 = (3.14159 * v5)
                    _ = fi62
                    _ = v63
                }
            }
            _ = v59
        }
        var v64 uint64 = 0
        var v65 int8 = (v2 & genFn1())
        _ = v64
        _ = v65
    }
    v4 = 964
    var v66 uint8 = 2
    fmt.Println(v66)
    if v3 {
        v67 := GenS7{F0: true, F1: 42, F2: v1}
        _ = v67
    } else {
        v3 = v3
    }
    v5 -= v5
    if false {
        _ = len(v6)
        var v68 int32 = -2147483648
        for fi69 := 0; fi69 < 2; fi69++ {
            for fi70 := 0; fi70 < 19; fi70++ {
                var v71 uint64 = 42
                var v72 uint16 = uint16(v71)
                var v73 int = 1
                _ = fi70
                _ = v71
                _ = v72
                _ = v73
            }
            if v3 {
                if (2147483648 == 65535) {
                    switch "" {
                    case "":
                        for fi74 := 0; fi74 < 12; fi74++ {
                            if false {
                                fmt.Println(v5)
                                var v75 bool = (fi74 != fi69)
                                var v76 float32 = (1.5 / float32((-2.25 - v5)))
                                fmt.Println(v68, v75, v66)
                                _ = v75
                                _ = v76
                            }
                            _ = fi74
                        }
                    case "foo bar":
                        v4 = 1
                    case "a":
                        v77 := []int{fi69, (fi69 - (-(v4))), fi69}
                        _ = v77
                    }
                }
                fmt.Println(v3, v1)
                if v3 {
                    v78 := GenS8{F0: 188, F1: (v4 | (fi69 >> ((42) & 63)))}
                    _ = v78
                }
            } else {
                v2 = v2
                v68 += v68
                switch "a" {
                case "foo bar":
                    v5 += v5
                case "hello":
                    if v3 {
                        v2 = 127
                        v68 = 0
                    } else {
                        fmt.Println(v68, v4, v4)
                        v79 := GenS6{F0: 2, F1: 339}
                        if (843 == 651) {
                            fmt.Println(v1, fi69)
                            if (18446744073709551615 != 42) {
                                var v80 uint = 879
                                var v81 uint16 = 115
                                v82 := []int{(v4 | (v4 | fi69)), fi69, fi69, fi69}
                                _ = v80
                                _ = v81
                                _ = v82
                            } else {
                                fmt.Println(v3, v3, v68)
                                fmt.Println(v2)
                            }
                        } else {
                            switch "" {
                            case "":
                                for fi83 := 0; fi83 < 18; fi83++ {
                                    if v3 {
                                        var v84 uint = 66
                                        if ((!(348 == 1)) || true) {
                                            fmt.Println(v2)
                                            v85 := GenS7{F0: (v4 > -321), F1: 2, F2: v1}
                                            _ = v85
                                        } else {
                                            var v86 int16 = 374
                                            fmt.Println(v86, v3)
                                            _ = v86
                                        }
                                        v87 := GenS9{F0: true, F1: v68, F2: v3}
                                        _ = v84
                                        _ = v87
                                    } else {
                                        v4 += 1
                                        switch "hello" {
                                        case "hello":
                                            var v88 string = "\n\t"
                                            _ = v88
                                        case "foo bar":
                                            if v3 {
                                                if (v4 <= (v4 << ((255) & 63))) {
                                                    v89 := GenS1{F0: (v4 & v4), F1: v1, F2: 731}
                                                    v90 := GenS4{F0: 2, F1: 10}
                                                    switch (fi83 << ((547) & 63)) {
                                                    case -323:
                                                        v4 = fi69
                                                    default:
                                                        var v91 float32 = 1.5
                                                        _ = v91
                                                    }
                                                    fmt.Println(v2)
                                                    fmt.Println(v2)
                                                    _ = v89
                                                    _ = v90
                                                }
                                            }
                                            v92 := GenS10{F0: 1, F1: 2}
                                            _ = len(v6)
                                            switch fi83 {
                                            case 1:
                                                var v93 uint8 = (v66 | 128)
                                                _ = v93
                                            }
                                            var v94 int32 = (v68 + v68)
                                            v94 += int32(fi69)
                                            _ = v92
                                            _ = v94
                                        case "":
                                            fmt.Println(fi69, v68, v2)
                                        }
                                        var v95 float32 = 0.0
                                        _ = v95
                                    }
                                    fmt.Println(v68, fi69, fi83)
                                    fmt.Println(fi69)
                                    _ = fi83
                                }
                            case "\n\t":
                                v96 := GenS7{F0: (0.0 <= 0.1), F1: 2, F2: 1}
                                _ = v96
                            case "foo bar":
                                for fi97 := 0; fi97 < 11; fi97++ {
                                    v98 := []int{fi97, fi97, v4}
                                    var v99 int8 = genFn1()
                                    switch ("hello" + (("foo bar" + "\n\t") + ("\n\t" + ""))) {
                                    case "a":
                                        v66 -= (v66 * 83)
                                    default:
                                        var v100 uint16 = 1
                                        _ = v100
                                    }
                                    _ = fi97
                                    _ = v98
                                    _ = v99
                                }
                            default:
                                v3 = v3
                            }
                            v101 := GenS11{F0: 33, F1: int8(v4)}
                            switch (fi69 >> ((65535) & 63)) {
                            case 4294967295:
                                v102 := GenS12{F0: 426, F1: ((((v5 - v5) / 1.5) / v5) == -2.25), F2: 2147483648}
                                _ = v102
                            case 32768:
                                var v103 uint = 1
                                _ = v103
                            default:
                                v66 = (v66 - v66)
                            }
                            _ = v101
                        }
                        _ = v79
                    }
                default:
                    for fi104 := 0; fi104 < 14; fi104++ {
                        var v105 int = (fi69 % (((v4 % ((fi104) | 1))) | 1))
                        var v106 uint16 = 255
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                }
            }
            if (!((!(!v3)) && true)) {
                fmt.Println(v68)
                var v107 uint8 = 128
                v4 = fi69
                _ = v107
            }
            _ = fi69
        }
        _ = v68
    }
    v108 := []int{}
    if len(v108) > 0 {
        v108[0] = (v4 - v4)
    }
    v109 := GenS12{F0: 2147483647, F1: v3, F2: 2147483648}
    switch v4 {
    case -1:
        v4 += v4
    case 128:
        var v110 uint16 = 42
        _ = v110
    case 127:
        v2 = (v2 * (v2 ^ v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
