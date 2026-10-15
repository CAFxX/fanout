package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint }
type GenS5 struct { F0 uint; F1 uint32; F2 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 bool; F1 int; F2 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint; F1 int }
func genFn0(p0 string, p1 uint) int {
    switch 128 {
    case 4294967295:
        if (!false) {
            if (false && (false || false)) {
                p1 -= 1
                var v7 int64 = -9223372036854775808
                _ = v7
            }
            switch "" {
            case "foo bar":
                var v8 int8 = -1
                _ = v8
            default:
                v9 := GenS1{F0: -1, F1: 2, F2: "foo bar"}
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    v12 := GenS1{F0: 629, F1: 65535, F2: ("foo bar" + p0)}
                    v13 := []int{fi10}
                    switch (-(fi11)) {
                    case 751:
                        if false {
                            switch (fi11 << ((p1) & 63)) {
                            case 32767:
                                var v14 int16 = 523
                                _ = v14
                            default:
                                fmt.Println(fi10, fi11, fi11)
                            }
                        }
                    case 32767:
                        fmt.Println(p0)
                    case 128:
                        v15 := GenS1{F0: 0, F1: 1, F2: (p0 + p0)}
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                switch fi10 {
                case 0:
                    v16 := GenS1{F0: -1, F1: 382, F2: (((p0 + "") + "hello") + "a")}
                    _ = v16
                case 1:
                    var v17 int8 = -128
                    _ = v17
                case 127:
                    v18 := GenS2{F0: 0}
                    _ = v18
                }
                _ = fi10
            }
        } else {
            var v19 uint = p1
            var v20 uint64 = 709
            _ = v19
            _ = v20
        }
    default:
        var v21 int = 4294967295
        _ = v21
    }
    p1 = 1
    for fi22 := 0; fi22 < 20; fi22++ {
        p0 = (p0 + p0)
        var v23 int = (fi22 / ((128) | 1))
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    return 127
}
func genFn1() int {
    var v24 uint16 = 32767
    v24 = (^(v24))
    switch "\n\t" {
    case "\n\t":
        switch -396 {
        case 32767:
            v25 := GenS3{F0: 1}
            _ = v25
        case 0:
            v26 := GenS4{F0: 0, F1: v24, F2: 18446744073709551615}
            _ = v26
        default:
            if (!false) {
                fmt.Println(v24)
                v24 -= (v24 + v24)
            }
        }
    case "a":
        if (true || (0 > -292)) {
            var v27 uint16 = (v24 / (((v24 % (((v24 | (v24 & 520))) | 1))) | 1))
            _ = v27
        }
    }
    _ = v24
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 float64 = (0.1 - (float64((((float32((0.0 / (float64((1.5 / 1e10)) * 1.5))) / 1.0) + (((3.14159 / -2.25) - 0.1) - 1.5)) - 0.1)) - (1e10 - float64(1e10))))
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    var v6 int64 = 949
    var v28 int64 = (v6 << ((0) & 63))
    var v29 float32 = -2.25
    for fi30 := 0; fi30 < 2; fi30++ {
        v3 = 1e10
        _ = fi30
    }
    v31 := v1[("a" + "foo bar")]
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint = 0
        if true {
            v3 = float64(0.0)
        } else {
            if (1 > 0) {
                var v34 float64 = v3
                var v35 int = 127
                _ = v34
                _ = v35
            } else {
                v2[fi32] = "hello"
                var v36 int32 = -1
                v36 = 36
                _ = v36
            }
        }
        var v37 uint64 = 18446744073709551615
        _ = fi32
        _ = v33
        _ = v37
    }
    v38 := GenS5{F0: 4294967295, F1: 678, F2: -1}
    var v39 uint16 = 0
    v29 -= v29
    v40 := GenS4{F0: 42, F1: (v39 + v39), F2: 144}
    v3 -= (v3 * 1.5)
    for fi41 := 0; fi41 < 9; fi41++ {
        for fi42 := 0; fi42 < 5; fi42++ {
            v4[v31] = (("foo bar" + v31) + v31)
            v6 -= (v6 * 703)
            _ = fi42
        }
        fmt.Println(v39, v39, v28)
        _ = fi41
    }
    v43 := []int{-851, genFn1()}
    for fi44 := 0; fi44 < 15; fi44++ {
        v45 := GenS6{F0: v31}
        if true {
            var v46 float64 = ((0.0 + (v3 * float64(((1.5 - v29) * (3.14159 / v29))))) * ((((float64(((v29 * v29) + v29)) / -2.25) / (v3 * 1.5)) / 0.1) * float64((((3.14159 + 0.1) + v29) * float32(v3)))))
            _ = v46
        }
        var v47 int = 127
        _ = fi44
        _ = v45
        _ = v47
    }
    if (!false) {
        if (42 >= 0) {
            var v48 float32 = 3.14159
            v3 = -2.25
            _ = v48
        }
        v49 := GenS7{F0: 686, F1: 1}
        if ((v39 >= 0) && true) {
            switch 1 {
            case 0:
                var v50 int16 = 1000
                _ = v50
            case 115:
                for fi51 := 0; fi51 < 8; fi51++ {
                    v39 += (v39 | v39)
                    _ = fi51
                }
            case 128:
                for fi52 := 0; fi52 < 5; fi52++ {
                    if true {
                        if (!(true || (!true))) {
                            var v53 int64 = (v28 | 954)
                            var v54 string = v31
                            for fi55 := 0; fi55 < 20; fi55++ {
                                switch "\n\t" {
                                case "\n\t":
                                    if true {
                                        switch fi52 {
                                        case 9223372036854775807:
                                            if (!((fi52 ^ (^(fi52))) >= (^(fi52)))) {
                                                if true {
                                                    if false {
                                                        v56 := GenS3{F0: 547}
                                                        _ = v56
                                                    }
                                                    var v57 int64 = v6
                                                    var v58 uint16 = (v39 * v39)
                                                    var v59 int32 = 0
                                                    if true {
                                                        switch ("\n\t" + v31) {
                                                        case "a":
                                                            if (79 > v6) {
                                                                fmt.Println(fi52, v6, fi52)
                                                                switch 42 {
                                                                case 2147483647:
                                                                    v31 = ((v54 + (v31 + "foo bar")) + v31)
                                                                    v39 = 65535
                                                                default:
                                                                    switch genFn0(v54, 42) {
                                                                    case 2147483647:
                                                                        var v60 string = "\n\t"
                                                                        _ = v60
                                                                    default:
                                                                        v61 := GenS8{F0: 127, F1: 127}
                                                                        if ((-(v59)) > (-(v59))) {
                                                                            var v62 int8 = -83
                                                                            v63 := GenS2{F0: 27}
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            if ((false || (2 > 42)) || (2 < 779)) {
                                                                                v64 := GenS9{F0: (0 != v53), F1: fi52, F2: "\n\t"}
                                                                                _ = v64
                                                                            } else {
                                                                                var v65 int8 = 42
                                                                                v58 += (v39 / ((v39) | 1))
                                                                                _ = v65
                                                                            }
                                                                            var v66 uint8 = 1
                                                                            if true {
                                                                                var v67 string = v54
                                                                                _ = v67
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                }
                                                                fmt.Println(v57)
                                                                if false {
                                                                    v68 := GenS10{F0: 4294967295}
                                                                    _ = len(v4)
                                                                    var v69 uint8 = 42
                                                                    _ = v68
                                                                    _ = v69
                                                                } else {
                                                                    fmt.Println(v57, v28, fi52)
                                                                    v70 := GenS4{F0: -1, F1: (v39 / ((42) | 1)), F2: 302}
                                                                    _ = v70
                                                                }
                                                            } else {
                                                                var v71 int8 = 100
                                                                _ = v71
                                                            }
                                                        }
                                                        v72 := GenS8{F0: -1, F1: 0}
                                                        _ = v72
                                                    } else {
                                                        if true {
                                                            v58 = (^(v39))
                                                            v57 = v53
                                                            var v73 uint = 366
                                                            var v74 uint64 = 9
                                                            if true {
                                                                v29 = (3.14159 * (-2.25 / ((v29 * v29) - (v29 + -2.25))))
                                                            } else {
                                                                if (0.0 == v3) {
                                                                    v75 := GenS11{F0: genFn1()}
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v76 uint32 = 2
                                                                            _ = v76
                                                                        }
                                                                    case "hello":
                                                                        switch v31 {
                                                                        case "\n\t":
                                                                            var v77 string = (v31 + ("a" + "foo bar"))
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                    var v78 int = genFn1()
                                                                    _ = v75
                                                                    _ = v78
                                                                }
                                                                var v79 int8 = 100
                                                                _ = v79
                                                            }
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        var v80 uint64 = 217
                                                        _ = v80
                                                    }
                                                    if true {
                                                        if true {
                                                            v43 = append(v43, fi52)
                                                        }
                                                        var v81 uint = 0
                                                        _ = v81
                                                    } else {
                                                        var v82 int32 = -2147483648
                                                        if false {
                                                            fmt.Println(v82, v29)
                                                            v43 = append(v43, (fi52 ^ (fi55 & genFn0(v54, 65535))))
                                                        }
                                                        _ = v82
                                                    }
                                                    if true {
                                                        var v83 int16 = -285
                                                        switch (fi55 - fi52) {
                                                        case 127:
                                                            fmt.Println(v83)
                                                        case 4294967295:
                                                            fmt.Println(v53)
                                                        default:
                                                            v43 = append(v43, fi55)
                                                        }
                                                        _ = v83
                                                    }
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    fmt.Println(v31, v29, v29)
                                                }
                                                if true {
                                                    fmt.Println(v53, v53)
                                                    _ = len(v43)
                                                } else {
                                                    var v84 int = fi52
                                                    v3 = (1.5 + (1e10 / (float64(-2.25) / (v3 - ((((0.1 + 1e10) * v3) * v3) * 0.1)))))
                                                    _ = v84
                                                }
                                            }
                                        case 994:
                                            if true {
                                                var v85 int = (fi52 + (fi55 % ((genFn0(v54, 135)) | 1)))
                                                _ = v85
                                            } else {
                                                var v86 uint64 = 67
                                                var v87 int16 = 42
                                                _ = v86
                                                _ = v87
                                            }
                                        case 128:
                                            var v88 int64 = (v28 + (v53 * (v6 << ((1) & 63))))
                                            _ = v88
                                        }
                                        switch (fi52 / (((fi52 + -1)) | 1)) {
                                        case 9223372036854775807:
                                            if true {
                                                var v89 int8 = -4
                                                _ = v89
                                            }
                                        case 127:
                                            fmt.Println(v3, v39)
                                        case -9223372036854775808:
                                            var v90 uint = 0
                                            _ = v90
                                        default:
                                            switch v31 {
                                            case "\n\t":
                                                switch fi52 {
                                                case 128:
                                                    var v91 uint = 4294967295
                                                    fmt.Println(v91, fi55)
                                                    _ = v91
                                                }
                                            default:
                                                var v92 uint32 = 2147483647
                                                _ = v92
                                            }
                                        }
                                    }
                                default:
                                    v53 = 1
                                }
                                switch 565 {
                                case 9223372036854775807:
                                    switch (fi55 | (fi55 / ((897) | 1))) {
                                    case 42:
                                        var v93 uint32 = 428
                                        _ = v93
                                    case 32768:
                                        v31 = v31
                                        v94 := GenS10{F0: 42}
                                        _ = v94
                                    case -2147483648:
                                        if false {
                                            v95 := GenS8{F0: 0, F1: -128}
                                            v6 = (v6 & (v53 ^ v6))
                                            v31 = v54
                                            _ = v95
                                        } else {
                                            if (-124 <= 100) {
                                                if (-32768 < 0) {
                                                    var v96 int8 = 127
                                                    fmt.Println(v29)
                                                    fmt.Println(v53, v39, v96)
                                                    _ = v96
                                                }
                                                var v97 int = (fi52 >> ((676) & 63))
                                                _ = v97
                                            }
                                            v98 := GenS12{F0: fi55}
                                            var v99 bool = (v31 > v31)
                                            _ = v98
                                            _ = v99
                                        }
                                    default:
                                        delete(v1, ("hello" + v54))
                                    }
                                case 32768:
                                    v100 := GenS13{F0: 60}
                                    _ = v100
                                case 2147483647:
                                    v101 := GenS14{F0: -283, F1: 100, F2: 0}
                                    _ = v101
                                }
                                v102 := v5[42]
                                _ = fi55
                                _ = v102
                            }
                            _ = v53
                            _ = v54
                        } else {
                            v28 = (v6 ^ 9223372036854775807)
                            var v103 int = (fi52 ^ 9223372036854775807)
                            _ = v103
                        }
                        var v104 uint8 = 2
                        for fi105 := 0; fi105 < 8; fi105++ {
                            var v106 uint16 = 850
                            v5[0] = 127
                            fmt.Println(v6, v106)
                            _ = fi105
                            _ = v106
                        }
                        _ = v104
                    }
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    _ = fi52
                }
            default:
                if (float32((v3 + v3)) < v29) {
                    v107 := GenS15{F0: 123, F1: 4294967295}
                    fmt.Println(v39)
                    _ = v107
                }
            }
            v43 = append(v43, -551)
        }
        _ = v49
    } else {
        var v108 float32 = 0.0
        _ = v108
    }
    v109 := v5[0]
    if ((v28 & v6) == v28) {
        for fi110 := 0; fi110 < 5; fi110++ {
            fmt.Println(v6)
            if len(v43) >= 2 {
                v43 = v43[1:len(v43)]
            }
            switch (v31 + ("foo bar" + v31)) {
            case "hello":
                fmt.Println(v39, v28, fi110)
            default:
                var v111 uint = 26
                _ = v111
            }
            _ = fi110
        }
        fmt.Println(v39)
    } else {
        v112 := GenS14{F0: v6, F1: -45, F2: 215}
        _ = v112
    }
    fmt.Println(v6, v29, v39)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v109
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
