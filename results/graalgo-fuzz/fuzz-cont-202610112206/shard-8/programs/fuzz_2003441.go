package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint64; F2 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8; F1 string; F2 uint32 }
type GenS11 struct { F0 bool; F1 int8 }
func genFn0() int16 {
    v5 := GenS1{F0: 490}
    var v6 bool = true
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string, p1 string, p2 int8) int32 {
    if ((p1 + p1) < ("" + p1)) {
        var v7 float32 = 0.0
        _ = v7
    }
    p0 = p0
    var v8 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 65535
    var v3 int32 = 27
    var v4 float64 = (1e10 - -2.25)
    switch 801 {
    case 127:
        v2 = v2
    }
    var v9 uint32 = v2
    var v10 int8 = 0
    v3 = (^(v3))
    var v11 int16 = -563
    if (!false) {
        var v12 uint16 = 2
        _ = v12
    } else {
        v13 := GenS1{F0: v2}
        _ = v13
    }
    switch "\n\t" {
    case "foo bar":
        var v14 bool = (((v9 >> ((65535) & 63)) == v2) && (688 >= 989))
        _ = v14
    case "":
        v15 := []int{4294967295, -723, 42}
        _ = v15
    }
    v16 := GenS2{F0: 255, F1: (v9 / (((v9 | 2147483648)) | 1)), F2: 23}
    var v17 uint64 = 42
    var v18 float64 = float64(float32(float64(((float32(3.14159) - ((1e10 * float32(v4)) - ((1.5 / -2.25) + 0.1))) + 0.0))))
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint16 = 249
        _ = fi19
        _ = v20
    }
    v11 = v11
    for fi21 := 0; fi21 < 12; fi21++ {
        v11 = (v11 / ((v11) | 1))
        var v22 int = 9223372036854775807
        if (!true) {
            if ((v10 >> ((42) & 63)) < (v10 + (v10 + (-(v10))))) {
                if ((883 != 1) && true) {
                    switch "hello" {
                    case "foo bar":
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                v23 := GenS3{F0: 42, F1: v17}
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    fmt.Println(v18)
                                    _ = fi24
                                }
                                var v25 float64 = 3.14159
                                _ = v23
                                _ = v25
                            } else {
                                var v26 int16 = v11
                                switch 9223372036854775807 {
                                case 4294967295:
                                    v1[v9] = v2
                                case -232:
                                    var v27 string = "hello"
                                    _ = v27
                                default:
                                    for fi28 := 0; fi28 < 20; fi28++ {
                                        var v29 uint32 = 42
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v30 := GenS4{F0: true, F1: v2, F2: v17}
                                    _ = v30
                                case "\n\t":
                                    switch "\n\t" {
                                    case "\n\t":
                                        v17 = (^(v17))
                                    case "foo bar":
                                        fmt.Println(v22, v11, v11)
                                    case "":
                                        for fi31 := 0; fi31 < 14; fi31++ {
                                            v4 = v18
                                            if true {
                                                switch "hello" {
                                                case "foo bar":
                                                    var v32 uint32 = (v9 ^ 799)
                                                    _ = v32
                                                case "\n\t":
                                                    fmt.Println(v3, v22)
                                                case "":
                                                    if (!false) {
                                                        if true {
                                                            v33 := GenS4{F0: false, F1: (v9 % ((754) | 1)), F2: 291}
                                                            v34 := GenS2{F0: 32767, F1: (v2 | 350), F2: 65535}
                                                            v10 -= v10
                                                            v35 := GenS3{F0: (v11 | 1000), F1: v17}
                                                            var v36 int8 = 0
                                                            _ = v33
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        switch fi31 {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi21, v2)
                                                        default:
                                                            v9 = (^(v2))
                                                        }
                                                    } else {
                                                        fmt.Println(v26, v3)
                                                        if false {
                                                            v37 := []int{fi31, fi21}
                                                            _ = v37
                                                        } else {
                                                            switch ("" + (("hello" + "a") + (((("foo bar" + ("hello" + "\n\t")) + "\n\t") + "foo bar") + "hello"))) {
                                                            case "foo bar":
                                                                var v38 float32 = (-2.25 - (1.5 - (-2.25 * (3.14159 / (float32(v18) * 1e10)))))
                                                                fmt.Println(v18, v17)
                                                                var v39 uint = 82
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            v40 := GenS5{F0: (v2 * v9), F1: ""}
                                                            _ = v40
                                                        }
                                                    }
                                                }
                                            }
                                            var v41 float64 = (((v4 / 1.0) * (0.1 / ((v4 - (1.5 / v4)) / (v4 - v18)))) * (v18 - 1.5))
                                            _ = fi31
                                            _ = v41
                                        }
                                    default:
                                        v10 += (v10 | (v10 + -66))
                                    }
                                case "a":
                                    v1[1] = v9
                                }
                                _ = v26
                            }
                        case "a":
                            switch "a" {
                            case "":
                                if false {
                                    fmt.Println(v2, v17, v4)
                                } else {
                                    v42 := v1[v2]
                                    switch "hello" {
                                    case "\n\t":
                                        if (!(-1 > -1)) {
                                            if true {
                                                switch "" {
                                                case "hello":
                                                    v17 += (v17 >> ((1) & 63))
                                                }
                                                fmt.Println(v42, v11, v3)
                                            }
                                        } else {
                                            v43 := []int{fi21, (fi21 << ((1) & 63)), (^(v22))}
                                            fmt.Println(fi21)
                                            _ = v43
                                        }
                                    case "a":
                                        switch 128 {
                                        case 32768:
                                            var v44 uint = 65535
                                            _ = v44
                                        case 2147483647:
                                            if (v2 <= 2) {
                                                switch (v22 ^ v22) {
                                                case 127:
                                                    var v45 int16 = v11
                                                    _ = v45
                                                }
                                                v46 := v1[(v2 ^ (v42 + 730))]
                                                _ = v46
                                            } else {
                                                var v47 bool = false
                                                v48 := []int{(v22 << ((1) & 63)), (v22 + v22), (^(fi21))}
                                                _ = v47
                                                _ = v48
                                            }
                                        case 128:
                                            v17 = 311
                                        }
                                    default:
                                        var v49 float64 = -2.25
                                        _ = v49
                                    }
                                    v50 := GenS6{F0: (^(v3))}
                                    _ = v42
                                    _ = v50
                                }
                            default:
                                var v51 int64 = -853
                                _ = v51
                            }
                        case "":
                            switch v22 {
                            case 351:
                                v3 = v3
                            case 214:
                                if (false || (v10 >= -128)) {
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        v2 = (v9 - 1)
                                    case 32767:
                                        if (genFn0() > (v11 | (v11 / ((-32768) | 1)))) {
                                            if (false && true) {
                                                v52 := GenS5{F0: (v2 ^ v9), F1: ("foo bar" + ("\n\t" + "foo bar"))}
                                                _ = v52
                                            }
                                            for fi53 := 0; fi53 < 6; fi53++ {
                                                switch 1 {
                                                case 128:
                                                    var v54 uint = 2
                                                    v22 -= (fi53 | fi21)
                                                    var v55 int32 = genFn1(((("\n\t" + (("foo bar" + "a") + "foo bar")) + "") + "foo bar"), "hello", v10)
                                                    _ = v54
                                                    _ = v55
                                                }
                                                _ = fi53
                                            }
                                            var v56 int16 = -261
                                            _ = v56
                                        }
                                    case 0:
                                        if ((v2 + (v2 / ((v2) | 1))) > 65535) {
                                            v22 = 42
                                        } else {
                                            for fi57 := 0; fi57 < 6; fi57++ {
                                                fmt.Println(v9, v17, v9)
                                                var v58 uint16 = 0
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v59 int = (fi21 + 201)
                                                    _ = v59
                                                case "":
                                                    if (!(v17 > 2)) {
                                                        switch ("\n\t" + ("hello" + ("a" + ""))) {
                                                        case "\n\t":
                                                            if true {
                                                                fmt.Println(v2, v9)
                                                                v17 = (v17 / (((v17 - 42)) | 1))
                                                            }
                                                            v60 := []int{}
                                                            _ = v60
                                                        case "a":
                                                            if true {
                                                                fmt.Println(v3)
                                                            } else {
                                                                v61 := []int{}
                                                                v4 = v18
                                                                fmt.Println(fi57, v18, fi21)
                                                                _ = v61
                                                            }
                                                        default:
                                                            v58 += (^(v58))
                                                        }
                                                    }
                                                default:
                                                    v17 = (v17 % ((v17) | 1))
                                                }
                                                _ = fi57
                                                _ = v58
                                            }
                                        }
                                    }
                                    var v62 int32 = (v3 + genFn1("\n\t", (("a" + ("a" + "")) + "hello"), (v10 >> ((42) & 63))))
                                    switch (fi21 & (fi21 >> ((1) & 63))) {
                                    case 215:
                                        v9 += 65535
                                    }
                                    _ = v62
                                }
                            }
                        default:
                            for fi63 := 0; fi63 < 17; fi63++ {
                                var v64 int = (v22 & fi21)
                                var v65 uint = 18446744073709551615
                                var v66 int32 = v3
                                _ = fi63
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                        }
                    default:
                        var v67 int16 = 1000
                        _ = v67
                    }
                } else {
                    delete(v1, (v9 / ((v9) | 1)))
                }
            } else {
                switch int(v17) {
                case 127:
                    for fi68 := 0; fi68 < 6; fi68++ {
                        var v69 string = "foo bar"
                        v22 = 1
                        if false {
                            switch (v69 + "") {
                            case "hello":
                                v70 := GenS7{F0: genFn0(), F1: 2, F2: v2}
                                _ = v70
                            case "foo bar":
                                v4 = -2.25
                            default:
                                delete(v1, v9)
                            }
                            if ((fi68 == 0) || true) {
                                fmt.Println(v9, v4)
                                var v71 int64 = -846
                                _ = v71
                            }
                            v72 := GenS8{F0: 2, F1: 403, F2: 1}
                            _ = v72
                        }
                        if (!false) {
                            v73 := GenS1{F0: 4294967295}
                            var v74 uint64 = (v17 / ((v17) | 1))
                            if true {
                                var v75 int64 = -9223372036854775808
                                var v76 int8 = 127
                                var v77 uint32 = (v2 << ((uint(v9)) & 63))
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                            fmt.Println(v69, v4)
                            _ = v73
                            _ = v74
                        }
                        var v78 uint32 = v9
                        _ = fi68
                        _ = v69
                        _ = v78
                    }
                case 128:
                    var v79 uint8 = 203
                    _ = v79
                case 4294967295:
                    v11 = (^(v11))
                }
                if false {
                    v2 = (v2 & v9)
                    v80 := GenS9{F0: 42}
                    v81 := []int{(v22 / ((int(v10)) | 1)), (^(v22)), -620, v22}
                    _ = v80
                    _ = v81
                } else {
                    var v82 int8 = v10
                    if (1 == 128) {
                        if false {
                            var v83 uint16 = 1
                            v4 -= v4
                            if true {
                                var v84 uint8 = 228
                                v82 = (v10 - v10)
                                var v85 int = v22
                                _ = v84
                                _ = v85
                            }
                            _ = v83
                        }
                        _ = len(v1)
                        var v86 bool = (2 == 343)
                        _ = v86
                    }
                    v87 := GenS10{F0: 117, F1: "a", F2: v2}
                    _ = v82
                    _ = v87
                }
            }
            v17 += (v17 >> ((17) & 63))
            var v88 int = fi21
            _ = v88
        } else {
            switch fi21 {
            case 32767:
                var v89 int16 = -829
                _ = v89
            case 32768:
                for fi90 := 0; fi90 < 18; fi90++ {
                    v91 := GenS3{F0: (v11 / (((v11 << ((65535) & 63))) | 1)), F1: (v17 & (^(v17)))}
                    v2 += v9
                    switch "hello" {
                    case "":
                        v3 = 1
                    case "foo bar":
                        if true {
                            v9 = (^(v2))
                        }
                    default:
                        v18 = v4
                    }
                    v3 = (v3 - v3)
                    fmt.Println(fi90, v4)
                    switch int(v10) {
                    case 2147483647:
                        v92 := []int{}
                        _ = v92
                    }
                    if false {
                        fmt.Println(fi21)
                        if false {
                            var v93 bool = true
                            switch 128 {
                            case -1:
                                var v94 int64 = -78
                                _ = v94
                            case 32768:
                                v95 := GenS9{F0: 0}
                                _ = v95
                            case 128:
                                switch "foo bar" {
                                case "":
                                    v18 = v18
                                case "\n\t":
                                    var v96 string = "foo bar"
                                    _ = v96
                                case "foo bar":
                                    fmt.Println(v2, v10)
                                default:
                                    if v93 {
                                        var v97 uint = 997
                                        _ = v97
                                    } else {
                                        var v98 float32 = -2.25
                                        if false {
                                            v93 = (642 == 42)
                                            if v93 {
                                                var v99 uint8 = 71
                                                _ = v99
                                            }
                                            v100 := GenS11{F0: (true || (!((v11 - int16(v10)) <= (-(v11))))), F1: (v10 << ((651) & 63))}
                                            delete(v1, uint32(v17))
                                            _ = v100
                                        }
                                        v101 := GenS4{F0: v93, F1: (v9 + (v2 % (((^(v2))) | 1))), F2: (v17 % ((159) | 1))}
                                        _ = v101
                                        _ = v98
                                    }
                                    fmt.Println(v93, v9, fi90)
                                }
                            default:
                                var v102 string = "a"
                                _ = v102
                            }
                            _ = v93
                        } else {
                            v18 = v4
                        }
                    }
                    _ = fi90
                    _ = v91
                }
            }
            var v103 int8 = 100
            _ = v103
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
