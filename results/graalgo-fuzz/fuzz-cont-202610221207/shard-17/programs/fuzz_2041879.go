package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint8; F1 int64; F2 bool }
type GenS7 struct { F0 uint; F1 int8; F2 bool }
type GenS8 struct { F0 int8; F1 int32; F2 int64 }
type GenS9 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS10 struct { F0 uint64; F1 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int8; F1 uint16 }
func genFn0() uint8 {
    v7 := GenS1{F0: -583}
    var v8 string = ""
    var v9 int = 9223372036854775807
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint32 = 4294967295
        if (!true) {
            v12 := []int{-9223372036854775808, v9, fi10}
            fmt.Println(v8, v11)
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint64 = 834
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := []int{v9, 42, (fi15 ^ fi15), -2147483648}
                _ = fi15
                _ = v16
            }
            switch (v8 + (("a" + (v8 + ((v8 + "") + v8))) + v8)) {
            case "":
                v11 = 0
            case "foo bar":
                var v17 uint8 = 42
                _ = v17
            default:
                v18 := []int{}
                _ = v18
            }
            v9 += (v9 ^ 2147483647)
        }
        v19 := GenS2{F0: 291, F1: 0}
        _ = fi10
        _ = v11
        _ = v19
    }
    var v20 uint32 = 2147483648
    _ = v20
    _ = v7
    _ = v8
    _ = v9
    return 127
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"a", "\n\t", ("hello" + ""), "foo bar"}
    var v3 int64 = 1
    var v4 float32 = 1.5
    v5 := make(map[string]uint32)
    v6 := []string{("" + "a"), ""}
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if ((("\n\t" + "\n\t") + "\n\t") > "foo bar") {
        fmt.Println(v3)
        var v21 int64 = (v3 * (v3 + -9223372036854775808))
        _ = v21
    }
    if (47 > 25) {
        fmt.Println(v4, v3)
        v3 -= (v3 / ((9223372036854775807) | 1))
    }
    v22 := GenS1{F0: 127}
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(v4, fi23, v3)
        switch ("hello" + "hello") {
        case "":
            switch (((("foo bar" + "a") + "") + "\n\t") + "a") {
            case "foo bar":
                v3 += (v3 + (v3 % ((v3) | 1)))
            }
        }
        v24 := GenS3{F0: 7}
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 0}
    for fi26 := 0; fi26 < 17; fi26++ {
        _ = len(v2)
        _ = fi26
    }
    delete(v1, "hello")
    v27 := GenS1{F0: 177}
    fmt.Println(v4, v3)
    var v28 float32 = v4
    v29 := GenS4{F0: 32767, F1: 1, F2: -83}
    v30 := GenS2{F0: 130, F1: -35}
    v4 -= 3.14159
    for fi31 := 0; fi31 < 15; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := GenS1{F0: (fi31 % ((0) | 1))}
            _ = fi32
            _ = v33
        }
        switch "a" {
        case "hello":
            for fi34 := 0; fi34 < 7; fi34++ {
                v35 := GenS3{F0: 128}
                var v36 float64 = 3.14159
                var v37 float64 = 0.1
                if ((!(1 == 2)) || true) {
                    fmt.Println(v36)
                    switch "a" {
                    case "":
                        switch (("\n\t" + "\n\t") + "foo bar") {
                        case "hello":
                            v4 = 1e10
                        case "":
                            switch "\n\t" {
                            case "hello":
                                v3 = v3
                            case "foo bar":
                                v3 -= 5
                            }
                        case "a":
                            v4 = (1.5 * v4)
                        default:
                            switch "a" {
                            case "hello":
                                if true {
                                    v37 = (-2.25 / v36)
                                } else {
                                    var v38 uint16 = 593
                                    _ = v38
                                }
                            default:
                                v6 = append(v6, "a")
                            }
                        }
                    case "\n\t":
                        v37 = ((float64(((1e10 * v28) * (0.1 - v28))) + v36) - 3.14159)
                    case "hello":
                        if (4294967295 <= 0) {
                            var v39 float32 = (0.1 * float32(v37))
                            _ = v39
                        } else {
                            switch "a" {
                            case "":
                                var v40 uint64 = 42
                                _ = v40
                            case "foo bar":
                                if true {
                                    v37 = ((v37 / (v37 * ((v36 * 0.1) - (v37 - (v36 * (v36 + (v36 - (1.5 / v36)))))))) - ((1.5 - v36) + ((v37 + v37) + (0.0 * v36))))
                                    var v41 uint32 = 491
                                    if (1 < 1) {
                                        v36 -= (1e10 / ((((v36 / v37) + v36) + ((v37 - 0.0) / 0.1)) + ((3.14159 - (v36 * 3.14159)) + 1.5)))
                                    } else {
                                        if (!("hello" > ("\n\t" + ("foo bar" + ("\n\t" + ("hello" + "a")))))) {
                                            fmt.Println(v41, fi34)
                                            v4 += v4
                                        }
                                        v37 = v37
                                    }
                                    v42 := v1["foo bar"]
                                    _ = v41
                                    _ = v42
                                }
                            default:
                                v4 = ((1e10 * 1e10) * 1e10)
                            }
                            switch "hello" {
                            case "foo bar":
                                if true {
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(fi31)
                                    case "a":
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v43 int64 = 790
                                            _ = v43
                                        }
                                    case "foo bar":
                                        v44 := v1[(("" + (("\n\t" + ("" + ("" + ("\n\t" + "\n\t")))) + "\n\t")) + "")]
                                        _ = v44
                                    default:
                                        if (true && (1 != 0)) {
                                            var v45 uint8 = 127
                                            if true {
                                                if (906 >= fi34) {
                                                    v36 = v37
                                                    v36 -= v36
                                                } else {
                                                    if (90 != 85) {
                                                        v28 -= 1.5
                                                    } else {
                                                        v46 := GenS5{F0: -128, F1: 2}
                                                        var v47 int8 = int8(v3)
                                                        v47 -= 58
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    _ = len(v6)
                                                }
                                                if ("hello" <= "foo bar") {
                                                    fmt.Println(v28, v45)
                                                    v2 = append(v2, ("hello" + ("\n\t" + "\n\t")))
                                                    if true {
                                                        delete(v5, ("hello" + ("" + (("foo bar" + ("hello" + "\n\t")) + "foo bar"))))
                                                        var v48 float64 = v37
                                                        _ = v48
                                                    }
                                                }
                                            }
                                            _ = v45
                                        } else {
                                            if (1000 >= 493) {
                                                var v49 int16 = 32767
                                                switch ((("\n\t" + "\n\t") + "a") + ("" + ("hello" + ("" + "foo bar")))) {
                                                case "hello":
                                                    var v50 uint = 4294967295
                                                    _ = v50
                                                }
                                                v5[("" + "\n\t")] = 299
                                                _ = v49
                                            }
                                            if ((1 != 0) && (146 > 168)) {
                                                switch "\n\t" {
                                                case "":
                                                    var v51 int8 = 0
                                                    _ = v51
                                                case "foo bar":
                                                    switch fi34 {
                                                    case -2147483648:
                                                        v52 := GenS3{F0: 1}
                                                        _ = v52
                                                    default:
                                                        v37 = v37
                                                    }
                                                case "hello":
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v53 uint64 = 2
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v54 int16 = 0
                                                                _ = v54
                                                            case "hello":
                                                                var v55 uint8 = 0
                                                                _ = v55
                                                            case "":
                                                                if ((1 > 2) && true) {
                                                                    _ = len(v6)
                                                                }
                                                            }
                                                            _ = v53
                                                        default:
                                                            v56 := GenS4{F0: 713, F1: 550, F2: 92}
                                                            _ = v56
                                                        }
                                                        var v57 uint32 = 65535
                                                        v58 := GenS6{F0: 255, F1: (v3 / (((v3 * (v3 - (v3 & int64(fi34))))) | 1)), F2: ((-2147483648 != 593) || (255 >= 4294967295))}
                                                        _ = v57
                                                        _ = v58
                                                    default:
                                                        fmt.Println(v37)
                                                    }
                                                default:
                                                    v37 = 3.14159
                                                }
                                                var v59 int64 = v3
                                                _ = v59
                                            } else {
                                                delete(v1, "")
                                                var v60 string = (("" + "foo bar") + "hello")
                                                var v61 int = fi34
                                                _ = v60
                                                _ = v61
                                            }
                                        }
                                    }
                                } else {
                                    var v62 string = (("a" + ("a" + "a")) + "foo bar")
                                    switch (v62 + v62) {
                                    case "a":
                                        if ((v62 + (v62 + (v62 + ("a" + v62)))) >= "\n\t") {
                                            if true {
                                                switch fi34 {
                                                case 128:
                                                    var v63 uint16 = 65535
                                                    _ = v63
                                                case 9223372036854775807:
                                                    if (501 > 2) {
                                                        v62 = v62
                                                        fmt.Println(v4, fi34)
                                                    } else {
                                                        switch fi31 {
                                                        case 460:
                                                            var v64 uint64 = 1
                                                            _ = v64
                                                        case -2147483648:
                                                            var v65 uint64 = 4294967295
                                                            _ = v65
                                                        }
                                                        var v66 uint8 = 128
                                                        _ = v66
                                                    }
                                                case 0:
                                                    fmt.Println(v28, v37)
                                                default:
                                                    v4 += v4
                                                }
                                                if false {
                                                    switch (fi34 % ((fi31) | 1)) {
                                                    case -9223372036854775808:
                                                        var v67 int64 = v3
                                                        _ = v67
                                                    case 240:
                                                        delete(v1, ((v62 + ((v62 + v62) + ("\n\t" + "\n\t"))) + (v62 + v62)))
                                                    }
                                                    v6 = append(v6, v62)
                                                    if false {
                                                        if true {
                                                            v68 := GenS7{F0: 2, F1: 0, F2: false}
                                                            _ = v68
                                                        } else {
                                                            var v69 string = "hello"
                                                            switch fi31 {
                                                            case 873:
                                                                v70 := GenS8{F0: -121, F1: 1, F2: v3}
                                                                _ = v70
                                                            case 933:
                                                                v1[v69] = 1
                                                            default:
                                                                v71 := GenS5{F0: 101, F1: 220}
                                                                fmt.Println(fi31)
                                                                v2 = append(v2, ((v62 + "foo bar") + "hello"))
                                                                _ = v71
                                                            }
                                                            v4 += float32(3.14159)
                                                            _ = v69
                                                        }
                                                        v4 += (-2.25 * 0.0)
                                                        v3 = v3
                                                    }
                                                }
                                                var v72 int8 = 1
                                                _ = v72
                                            } else {
                                                _ = len(v5)
                                                fmt.Println(fi31)
                                                if false {
                                                    var v73 int = 32767
                                                    _ = v73
                                                }
                                            }
                                        }
                                    case "hello":
                                        switch (("a" + v62) + v62) {
                                        case "":
                                            if false {
                                                v37 += float64((v28 + (v4 * 0.0)))
                                                var v74 int32 = -2147483648
                                                if (((v28 - v28) / 3.14159) < v4) {
                                                    if (42 >= (fi34 >> ((844) & 63))) {
                                                        v28 += -2.25
                                                    }
                                                    var v75 int16 = 1
                                                    var v76 int64 = 1
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v77 := GenS5{F0: -10, F1: 4294967295}
                                                        _ = v77
                                                    case "":
                                                        v78 := v5[("" + "hello")]
                                                        _ = v78
                                                    }
                                                }
                                                v79 := GenS4{F0: 0, F1: 42, F2: int8(fi31)}
                                                if (!(!false)) {
                                                    v5[v62] = 2147483647
                                                    v80 := GenS7{F0: 886, F1: -40, F2: true}
                                                    _ = v80
                                                } else {
                                                    v74 -= (^(v74))
                                                    fmt.Println(v28)
                                                }
                                                _ = v74
                                                _ = v79
                                            } else {
                                                if (v36 != 3.14159) {
                                                    if (1e10 == v4) {
                                                        v81 := GenS9{F0: 1, F1: 1, F2: 32767}
                                                        var v82 float64 = v36
                                                        var v83 string = ("a" + ("\n\t" + "\n\t"))
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        fmt.Println(v62, fi34, v4)
                                                        if (-802 >= (^(fi31))) {
                                                            v4 = v28
                                                        } else {
                                                            _ = len(v6)
                                                        }
                                                        var v84 float32 = 3.14159
                                                        _ = v84
                                                    }
                                                    v4 = (0.1 + (float32((float64(1.5) + v37)) - 1e10))
                                                }
                                            }
                                        case "hello":
                                            switch fi34 {
                                            case -677:
                                                v3 += v3
                                            case -531:
                                                var v85 bool = false
                                                _ = v85
                                            }
                                        case "a":
                                            v28 = (((v4 / (v28 / 3.14159)) / (1e10 + v4)) * float32((float64((v28 / v28)) * float64(v4))))
                                        default:
                                            v86 := GenS10{F0: 986, F1: -128}
                                            _ = v86
                                        }
                                    case "":
                                        v3 = (v3 - (^(v3)))
                                    default:
                                        var v87 uint8 = 0
                                        _ = v87
                                    }
                                    v37 = (v36 + 0.0)
                                    _ = v62
                                }
                                if true {
                                    switch "foo bar" {
                                    case "a":
                                        v37 += (v37 + (0.0 + v36))
                                    case "":
                                        var v88 bool = (0 >= 22)
                                        _ = v88
                                    case "\n\t":
                                        fmt.Println(v28, v37, v3)
                                    default:
                                        v37 = v36
                                    }
                                } else {
                                    _ = len(v5)
                                    v4 = 3.14159
                                }
                            default:
                                delete(v1, "")
                            }
                        }
                    }
                    var v89 uint16 = 1
                    if true {
                        var v90 uint8 = genFn0()
                        var v91 float32 = v4
                        var v92 uint16 = v89
                        v90 -= (v90 << ((0) & 63))
                        _ = v90
                        _ = v91
                        _ = v92
                    } else {
                        var v93 uint32 = 2147483648
                        v2 = append(v2, ("a" + ""))
                        _ = v93
                    }
                    switch fi34 {
                    case -1:
                        var v94 float64 = v37
                        _ = v94
                    case 2147483647:
                        v37 += 3.14159
                    }
                    v4 += ((0.0 - (v28 * (1.5 - 0.1))) + v28)
                    if (1 >= -1) {
                        v95 := GenS3{F0: 156}
                        _ = v95
                    } else {
                        v37 -= 3.14159
                        v4 += 1.5
                        var v96 int64 = (v3 - v3)
                        v3 -= v3
                        var v97 uint = 65535
                        _ = v96
                        _ = v97
                    }
                    _ = v89
                } else {
                    var v98 int = fi34
                    switch v98 {
                    case 173:
                        fmt.Println(fi34, v36)
                    case 4294967295:
                        var v99 int8 = 127
                        _ = v99
                    case -1:
                        if (1 != 65535) {
                            v36 = v36
                            v28 -= (float32(v36) / (-2.25 * v4))
                            var v100 float32 = (v28 * v28)
                            _ = v100
                        } else {
                            fmt.Println(fi34, v4)
                            if (false || (false || ((v3 % ((316) | 1)) > v3))) {
                                var v101 bool = ((v3 < 9223372036854775807) || (582 > 42))
                                v4 = 1.5
                                v102 := GenS11{F0: genFn0()}
                                _ = v101
                                _ = v102
                            }
                        }
                    }
                    v103 := GenS3{F0: 128}
                    _ = v103
                    _ = v98
                }
                v104 := GenS11{F0: genFn0()}
                _ = fi34
                _ = v104
                _ = v35
                _ = v36
                _ = v37
            }
        case "foo bar":
            v105 := GenS12{F0: -69}
            _ = v105
        default:
            v106 := GenS13{F0: 42, F1: 2}
            _ = v106
        }
        _ = fi31
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
