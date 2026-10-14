package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int64; F1 uint8; F2 int }
type GenS7 struct { F0 uint64; F1 uint; F2 int32 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS10 struct { F0 uint8; F1 uint; F2 int32 }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 int16, p1 int32) int32 {
    p0 += (^(p0))
    var v6 float64 = 0.0
    var v7 int8 = 127
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 uint16 = 795
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 int16 = 1000
    var v5 int16 = v4
    var v8 int32 = 2147483647
    if true {
        v9 := GenS1{F0: v5, F1: ((0 != 1) && (640 > 13)), F2: 100}
        v10 := GenS1{F0: 32767, F1: (42 != 42), F2: -128}
        _ = v10
        _ = v9
    } else {
        if (v8 > 1) {
            v11 := []int{128, -935}
            if false {
                v1 = (v1 * v1)
                if false {
                    if true {
                        var v12 float32 = 0.0
                        var v13 uint8 = 2
                        _ = v12
                        _ = v13
                    } else {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    }
                } else {
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        if true {
                            v5 += (^(v5))
                            delete(v2, -1)
                        }
                    case "hello":
                        v14 := GenS1{F0: v4, F1: ((172 < 20) || (float64(3.14159) > float64(0.1))), F2: -28}
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v8 = v8
                        var v16 uint16 = 1
                        v5 = (v5 | (v5 >> ((255) & 63)))
                        _ = fi15
                        _ = v16
                    }
                }
            }
            _ = v11
        }
        v8 -= v8
        switch "\n\t" {
        case "foo bar":
            delete(v3, 9223372036854775807)
        default:
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 10; fi18++ {
                    if true {
                        fmt.Println(v1, fi17, v4)
                        switch (^(fi17)) {
                        case -2147483648:
                            var v19 float32 = 3.14159
                            _ = v19
                        case 127:
                            v1 = 32767
                        default:
                            if (35 > 190) {
                                if true {
                                    switch (fi18 * 42) {
                                    case 0:
                                        v20 := GenS1{F0: (^(v4)), F1: (-63 == 1), F2: 42}
                                        switch fi18 {
                                        case 127:
                                            v21 := GenS2{F0: (v8 | v8)}
                                            _ = v21
                                        case 128:
                                            if (0.1 != (1e10 + (float64(1e10) - ((0.0 + 0.0) * ((0.1 + -2.25) + -2.25))))) {
                                                v22 := []int{fi17, 421, (fi17 << ((2) & 63))}
                                                _ = v22
                                            } else {
                                                _ = len(v2)
                                            }
                                        }
                                        var v23 float32 = (3.14159 / 1.0)
                                        _ = v20
                                        _ = v23
                                    case -1:
                                        v8 += (v8 * v8)
                                    default:
                                        if false {
                                            var v24 uint64 = 18446744073709551615
                                            if true {
                                                if true {
                                                    var v25 int16 = (v4 | v5)
                                                    if true {
                                                        fmt.Println(v24, fi17, v8)
                                                        v26 := GenS1{F0: (-(v25)), F1: false, F2: 127}
                                                        switch 32767 {
                                                        case 784:
                                                            switch "a" {
                                                            case "a":
                                                                if (1 == 255) {
                                                                    var v27 int32 = v8
                                                                    v8 -= v27
                                                                    _ = v27
                                                                }
                                                            case "foo bar":
                                                                var v28 float64 = 1e10
                                                                _ = v28
                                                            default:
                                                                fmt.Println(fi18)
                                                            }
                                                        case 2147483647:
                                                            v29 := GenS2{F0: genFn0(v25, (v8 % ((42) | 1)))}
                                                            _ = v29
                                                        case 32767:
                                                            v30 := []int{(^(fi17))}
                                                            _ = v30
                                                        }
                                                        _ = v26
                                                    } else {
                                                        fmt.Println(v25)
                                                        if (!false) {
                                                            fmt.Println(v25)
                                                        }
                                                        var v31 uint8 = 168
                                                        _ = v31
                                                    }
                                                    if false {
                                                        v32 := GenS3{F0: -23}
                                                        fmt.Println(v1, fi17, v5)
                                                        _ = len(v3)
                                                        _ = v32
                                                    } else {
                                                        v33 := GenS2{F0: int32(v24)}
                                                        if (false && (77 > 100)) {
                                                            var v34 int8 = 14
                                                            _ = v34
                                                        } else {
                                                            var v35 uint64 = v24
                                                            v36 := GenS2{F0: (v8 ^ (v8 >> ((302) & 63)))}
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        switch -9223372036854775808 {
                                                        case 4294967295:
                                                            if (67 == 203) {
                                                                if false {
                                                                    v37 := []int{(fi17 & (fi18 & fi18)), fi18}
                                                                    v38 := GenS4{F0: (v1 % ((v1) | 1))}
                                                                    v39 := v2[fi17]
                                                                    _ = v37
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    var v40 float32 = -2.25
                                                                    fmt.Println(fi18, v25, v4)
                                                                    v41 := v3[fi17]
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                                v42 := []int{fi18, fi17, (fi17 - (^(fi18))), (fi18 | fi17)}
                                                                _ = v42
                                                            } else {
                                                                var v43 uint8 = 42
                                                                v44 := []int{}
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                        case -1:
                                                            var v45 int32 = (v8 >> ((4294967295) & 63))
                                                            _ = v45
                                                        }
                                                        fmt.Println(v1)
                                                        switch (fi17 - (fi18 % ((-2) | 1))) {
                                                        case 32768:
                                                            v46 := []int{638, (fi17 / (((fi18 << ((255) & 63))) | 1)), 128}
                                                            _ = v46
                                                        }
                                                        v25 += (v25 - (v4 / ((-830) | 1)))
                                                        _ = v33
                                                    }
                                                    _ = v25
                                                } else {
                                                    v5 = v5
                                                }
                                            }
                                            if false {
                                                var v47 int = fi17
                                                _ = v47
                                            } else {
                                                if true {
                                                    v1 -= v1
                                                } else {
                                                    delete(v2, fi18)
                                                    v48 := []int{int(v1), fi18}
                                                    var v49 int32 = v8
                                                    var v50 int32 = 2147483647
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                }
                                                if (978 > (v24 << ((255) & 63))) {
                                                    v1 += (v1 << ((0) & 63))
                                                    fmt.Println(v1, fi17, v4)
                                                    var v51 uint64 = v24
                                                    _ = v51
                                                }
                                                var v52 uint64 = v24
                                                _ = v52
                                            }
                                            var v53 int64 = int64(v24)
                                            v1 = v1
                                            _ = v24
                                            _ = v53
                                        }
                                    }
                                    if false {
                                        v54 := GenS5{F0: 4294967295, F1: -468}
                                        _ = v54
                                    }
                                } else {
                                    v1 += (v1 - 2)
                                    if ((!(fi17 < fi18)) || (0 < 18446744073709551615)) {
                                        fmt.Println(v4, v1)
                                    }
                                }
                                fmt.Println(fi18)
                                v55 := GenS6{F0: -1, F1: 174, F2: (fi18 << ((0) & 63))}
                                _ = v55
                            }
                        }
                        switch "" {
                        case "":
                            switch fi17 {
                            case -2147483648:
                                var v56 uint8 = 0
                                _ = v56
                            case -712:
                                var v57 uint = 2
                                _ = v57
                            }
                        }
                    }
                    if (0 != 181) {
                        v58 := GenS4{F0: 32767}
                        var v59 float32 = (3.14159 - (-2.25 / 1.0))
                        _ = v58
                        _ = v59
                    }
                    _ = fi18
                }
                v60 := GenS7{F0: 127, F1: 2, F2: (v8 / ((v8) | 1))}
                for fi61 := 0; fi61 < 11; fi61++ {
                    v62 := GenS8{F0: 0, F1: (int64(v8) > 0)}
                    _ = fi61
                    _ = v62
                }
                _ = fi17
                _ = v60
            }
        }
    }
    for fi63 := 0; fi63 < 20; fi63++ {
        v5 -= 723
        if true {
            v64 := GenS9{F0: (v4 + (v5 & -32768)), F1: 2147483648, F2: 255}
            _ = v64
        } else {
            for fi65 := 0; fi65 < 3; fi65++ {
                fmt.Println(v1, v5, v4)
                if true {
                    switch (^(fi63)) {
                    case 794:
                        var v66 uint16 = v1
                        _ = v66
                    case -1:
                        switch (fi65 / ((fi65) | 1)) {
                        case -399:
                            v8 += v8
                        case 0:
                            if (0.1 <= (3.14159 / -2.25)) {
                                v67 := GenS7{F0: 42, F1: 1, F2: v8}
                                var v68 uint64 = 4294967295
                                _ = v67
                                _ = v68
                            } else {
                                if ((false && false) && ((fi65 != (fi63 & fi63)) || (true || true))) {
                                    fmt.Println(v5)
                                    if true {
                                        var v69 string = "\n\t"
                                        v8 -= -1
                                        _ = v69
                                    } else {
                                        v70 := GenS7{F0: 588, F1: 65535, F2: v8}
                                        var v71 bool = (false || (!(true || (fi63 < fi65))))
                                        v71 = ((v8 / ((0) | 1)) > (v8 >> ((0) & 63)))
                                        _ = v70
                                        _ = v71
                                    }
                                    var v72 float32 = ((((float32(((-2.25 / 3.14159) / 3.14159)) / 0.1) * -2.25) - (float32(0.1) * 0.0)) - 3.14159)
                                    _ = v72
                                }
                            }
                        case -100:
                            var v73 uint64 = 326
                            _ = v73
                        }
                    case 42:
                        v74 := GenS2{F0: genFn0(v4, (v8 >> ((4294967295) & 63)))}
                        _ = v74
                    }
                    v8 = (v8 | v8)
                } else {
                    fmt.Println(v5)
                    v75 := []int{fi63}
                    v76 := GenS4{F0: v1}
                    _ = v75
                    _ = v76
                }
                v4 += (^(v4))
                _ = fi65
            }
        }
        _ = fi63
    }
    v77 := GenS10{F0: 1, F1: 65535, F2: (v8 ^ v8)}
    if true {
        v4 = 202
        for fi78 := 0; fi78 < 17; fi78++ {
            var v79 bool = false
            if (!v79) {
                v80 := GenS2{F0: -2147483648}
                var v81 uint16 = (v1 ^ (v1 / ((v1) | 1)))
                _ = v80
                _ = v81
            } else {
                v5 = -32768
                v79 = v79
            }
            _ = fi78
            _ = v79
        }
    }
    v82 := []int{772, 1}
    if false {
        var v83 uint8 = 42
        _ = v83
    } else {
        var v84 int32 = (v8 ^ v8)
        var v85 int16 = (^(v4))
        v85 -= v4
        _ = v84
        _ = v85
    }
    v86 := GenS11{F0: v1, F1: 325}
    if len(v82) >= 2 {
        v82 = v82[1:len(v82)]
    }
    switch -391 {
    case -1:
        var v87 int64 = -1
        _ = v87
    default:
        switch int(v5) {
        case 128:
            if true {
                if len(v82) >= 2 {
                    v82 = v82[1:len(v82)]
                }
            } else {
                var v88 int8 = 42
                _ = v88
            }
        case -2147483648:
            v1 = v1
        case 9223372036854775807:
            v89 := GenS5{F0: 658, F1: 1}
            _ = v89
        }
    }
    if (2 < 255) {
        for fi90 := 0; fi90 < 8; fi90++ {
            fmt.Println(v1)
            var v91 int8 = 100
            _ = fi90
            _ = v91
        }
    } else {
        if true {
            if (!(!((v5 / (((v5 * (v5 << ((1) & 63)))) | 1)) != (-(v5))))) {
                switch ("a" + "\n\t") {
                case "":
                    var v92 int32 = v8
                    _ = v92
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        if len(v82) > 0 {
                            v82[0] = -1
                        }
                    case "foo bar":
                        v2[-425] = "\n\t"
                    }
                case "foo bar":
                    v93 := GenS11{F0: (^(v1)), F1: 0}
                    _ = v93
                }
            } else {
                var v94 uint8 = 1
                v94 -= (v94 % ((v94) | 1))
                _ = v94
            }
            switch int(v1) {
            case 233:
                var v95 int64 = 98
                _ = v95
            }
        } else {
            var v96 uint32 = 180
            delete(v3, 42)
            _ = v96
        }
        var v97 float64 = float64(-2.25)
        _ = v97
    }
    if (("a" + "a") != (((("" + ("foo bar" + "a")) + "a") + "hello") + ("\n\t" + ((("a" + "hello") + ("" + "")) + "\n\t")))) {
        v5 += (^(v5))
        if true {
            for fi98 := 0; fi98 < 9; fi98++ {
                for fi99 := 0; fi99 < 17; fi99++ {
                    v82 = append(v82, fi98)
                    v8 = 42
                    _ = fi99
                }
                for fi100 := 0; fi100 < 5; fi100++ {
                    v8 += genFn0(v4, v8)
                    fmt.Println(fi98, fi100, v8)
                    if true {
                        if (0 <= 99) {
                            var v101 uint32 = 4294967295
                            _ = v101
                        } else {
                            v1 = v1
                            v1 -= 1
                        }
                    }
                    _ = fi100
                }
                for fi102 := 0; fi102 < 14; fi102++ {
                    switch "hello" {
                    case "hello":
                        var v103 int16 = 1000
                        _ = v103
                    case "":
                        fmt.Println(v1)
                    case "\n\t":
                        v4 = (-(v4))
                    default:
                        v1 = (v1 >> ((4294967295) & 63))
                        if ((4294967295 == 42) || (((!false) || (0.1 >= (0.1 + (((1e10 + (3.14159 * (-2.25 / -2.25))) * 3.14159) * (1e10 / 1.0))))) && ((!true) || true))) {
                            _ = len(v82)
                        } else {
                            var v104 int32 = v8
                            _ = v104
                        }
                        var v105 int32 = v8
                        _ = v105
                    }
                    var v106 uint64 = 18446744073709551615
                    _ = fi102
                    _ = v106
                }
                _ = fi98
            }
        }
        v107 := v3[0]
        _ = v107
    } else {
        for fi108 := 0; fi108 < 18; fi108++ {
            v4 = (v4 + v4)
            _ = fi108
        }
        v109 := GenS12{F0: (^(v1))}
        _ = v109
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v8
    _ = v82
    _ = v86
}
