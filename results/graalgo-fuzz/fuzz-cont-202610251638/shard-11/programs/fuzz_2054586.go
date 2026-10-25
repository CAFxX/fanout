package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 int64; F1 int64; F2 string }
type GenS10 struct { F0 int64; F1 int; F2 string }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 uint }

func main() {
    var v1 int8 = 127
    var v2 uint32 = 887
    var v3 float32 = float32(float64(0.0))
    var v4 uint = 1
    v5 := make(map[string]string)
    var v6 float32 = ((v3 / v3) * v3)
    switch 128 {
    case 211:
        switch "\n\t" {
        case "a":
            v7 := GenS1{F0: 9223372036854775807, F1: v4, F2: (v1 ^ (v1 >> ((v4) & 63)))}
            _ = v7
        case "\n\t":
            v8 := []int{int(v4), 32768, -767}
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: 42, F1: v4, F2: (v1 | 1)}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint8 = 42
        v3 = v6
        if ("a" == (("" + "a") + ((("" + "hello") + "hello") + ("a" + ("" + ("hello" + "foo bar")))))) {
            fmt.Println(fi10, fi10)
        } else {
            v2 -= 2147483647
            var v12 int64 = 903
            v13 := GenS1{F0: (v12 & (-(v12))), F1: 1, F2: v1}
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint = v4
    fmt.Println(v4, v1)
    v15 := GenS1{F0: 42, F1: (^(v4)), F2: (-(v1))}
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := []int{fi16, (fi16 & fi16), (fi16 << ((18446744073709551615) & 63))}
        if (true || (-2147483648 > 2147483647)) {
            if true {
                if true {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := GenS2{F0: int32(v2)}
                        _ = fi18
                        _ = v19
                    }
                } else {
                    fmt.Println(fi16)
                    v2 += (v2 | (v2 + v2))
                }
                v20 := GenS2{F0: -1}
                _ = v20
            } else {
                if true {
                    if false {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            switch (fi16 * (-(fi16))) {
                            case -9223372036854775808:
                                fmt.Println(v3, v1, v14)
                            case 32767:
                                var v22 float64 = float64(0.0)
                                v22 -= 1e10
                                v23 := v5["\n\t"]
                                v24 := GenS3{F0: true}
                                v25 := v5["a"]
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            default:
                                fmt.Println(v3, v1)
                            }
                            if true {
                                switch fi21 {
                                case 85:
                                    _ = len(v5)
                                default:
                                    v3 = 1.5
                                }
                                v26 := v5["a"]
                                _ = v26
                            }
                            _ = fi21
                        }
                        _ = len(v17)
                        if (-1 <= -570) {
                            var v27 float32 = 1e10
                            fmt.Println(v1, fi16)
                            _ = v27
                        } else {
                            var v28 uint16 = 255
                            _ = v28
                        }
                    }
                }
                v4 += 255
                for fi29 := 0; fi29 < 9; fi29++ {
                    v17 = append(v17, 377)
                    fmt.Println(v4, fi29)
                    var v30 uint = (v14 >> ((uint(v2)) & 63))
                    v31 := GenS2{F0: 443}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        } else {
            switch "" {
            case "hello":
                var v32 string = ("" + "\n\t")
                _ = v32
            default:
                if false {
                    v1 = 42
                    if false {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = len(v5)
                        delete(v5, "")
                    } else {
                        fmt.Println(v2, v14)
                        switch 296 {
                        case 775:
                            fmt.Println(v3, v1)
                        case 0:
                            for fi33 := 0; fi33 < 11; fi33++ {
                                var v34 bool = (121 < 0)
                                var v35 uint16 = 42
                                if true {
                                    v35 = v35
                                } else {
                                    v35 += 2
                                }
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        case 185:
                            v17 = append(v17, fi16)
                        }
                    }
                } else {
                    v36 := v5[("a" + "foo bar")]
                    _ = v36
                }
            }
            v1 = -128
        }
        _ = fi16
        _ = v17
    }
    v37 := GenS4{F0: 2}
    var v38 bool = (!(438 > 1))
    var v39 float64 = (-2.25 / float64(((v3 + v3) + (((v6 / ((0.0 * v6) / 1.0)) + 0.0) / ((v6 + (float32(0.0) * (v3 + v3))) * v3)))))
    switch -725 {
    case 4294967295:
        for fi40 := 0; fi40 < 11; fi40++ {
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 float32 = (v3 / (v3 - 0.1))
                _ = fi41
                _ = v42
            }
            for fi43 := 0; fi43 < 1; fi43++ {
                v14 = v14
                fmt.Println(v2)
                if ((42 < uint64(v2)) && ((!((1e10 == (v39 - float64(float32(((1e10 - (1.5 + float64(float32(v39)))) + v39))))) && v38)) && (-101 <= (fi43 % (((fi40 << ((v14) & 63))) | 1))))) {
                    switch "hello" {
                    case "foo bar":
                        v44 := GenS1{F0: -1, F1: v4, F2: (v1 * (v1 * (v1 + v1)))}
                        _ = v44
                    case "hello":
                        var v45 uint = (v4 / ((v14) | 1))
                        _ = v45
                    }
                    if (((-1 <= 42) || ((v2 << (((v14 >> ((65535) & 63))) & 63)) != v2)) || (!false)) {
                        var v46 int32 = 42
                        v47 := []int{int(v4)}
                        _ = v46
                        _ = v47
                    }
                }
                _ = fi43
            }
            v48 := GenS2{F0: -955}
            _ = fi40
            _ = v48
        }
    case -365:
        v49 := GenS5{F0: int32(v2), F1: 45}
        _ = v49
    }
    if false {
        v14 -= (v4 << ((v14) & 63))
        fmt.Println(v6, v1, v6)
    } else {
        v6 += -2.25
        fmt.Println(v6)
    }
    for fi50 := 0; fi50 < 7; fi50++ {
        if true {
            if v38 {
                v51 := GenS6{F0: v38, F1: (v14 | 2), F2: 237}
                fmt.Println(v6)
                _ = v51
            }
            var v52 uint32 = (v2 | (v2 * v2))
            _ = v52
        }
        if v38 {
            v53 := GenS1{F0: int64(v1), F1: 255, F2: (v1 | v1)}
            if (0.0 <= (v6 * ((float32(v39) + (3.14159 / -2.25)) * v3))) {
                for fi54 := 0; fi54 < 18; fi54++ {
                    if ((v4 ^ (v4 % ((2) | 1))) >= (v14 & (^(v14)))) {
                        var v55 int64 = 1
                        var v56 uint32 = v2
                        var v57 float64 = (-2.25 - v39)
                        if false {
                            if true {
                                v58 := GenS7{F0: 199, F1: v55}
                                var v59 int64 = (v55 / (((v55 % ((v55) | 1))) | 1))
                                _ = v58
                                _ = v59
                            } else {
                                v55 = v55
                                switch (fi54 ^ fi54) {
                                case 128:
                                    var v60 bool = (!((^(v56)) <= (v2 + v56)))
                                    _ = v60
                                case 32767:
                                    var v61 uint8 = 255
                                    _ = v61
                                default:
                                    v62 := GenS8{F0: 2147483647, F1: 42}
                                    _ = v62
                                }
                                fmt.Println(v14)
                            }
                            v63 := GenS9{F0: 9223372036854775807, F1: v55, F2: ("a" + "a")}
                            var v64 bool = (v55 > (v55 + 0))
                            delete(v5, "a")
                            _ = v63
                            _ = v64
                        }
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                    var v65 uint = 2
                    _ = fi54
                    _ = v65
                }
                v66 := []int{}
                v67 := v5[""]
                _ = v66
                _ = v67
            } else {
                var v68 uint64 = 2
                _ = v68
            }
            _ = v53
        }
        if v38 {
            if false {
                v3 -= ((((v3 - v6) / 1.5) + (v3 - 1.5)) - v6)
                var v69 int8 = v1
                v70 := GenS2{F0: -1}
                _ = v69
                _ = v70
            } else {
                if (1 >= (v14 | (v4 * 190))) {
                    switch ("\n\t" + ("a" + ("hello" + "a"))) {
                    case "foo bar":
                        delete(v5, (("" + "foo bar") + "\n\t"))
                    case "hello":
                        var v71 float32 = (0.1 + v6)
                        _ = v71
                    }
                }
                v72 := GenS9{F0: -1, F1: 42, F2: ((("" + ("" + "\n\t")) + "\n\t") + "foo bar")}
                _ = v72
            }
        } else {
            var v73 float32 = -2.25
            v14 = v14
            _ = v73
        }
        _ = fi50
    }
    switch "foo bar" {
    case "\n\t":
        if (2147483647 <= 9223372036854775807) {
            for fi74 := 0; fi74 < 11; fi74++ {
                for fi75 := 0; fi75 < 16; fi75++ {
                    if ((!v38) && v38) {
                        fmt.Println(v1, v38)
                        var v76 int = (fi75 % ((2147483647) | 1))
                        _ = v76
                    } else {
                        if v38 {
                            v77 := GenS4{F0: 122}
                            _ = v77
                        }
                        if (-1 == 42) {
                            v78 := []int{(fi74 * (fi74 >> (((v4 % ((uint(v2)) | 1))) & 63)))}
                            v3 = (0.1 - ((1e10 + (-2.25 / float32(float64(((((v6 / (1e10 * v3)) * float32(v39)) * 0.0) + (((v6 - -2.25) * (v6 + v6)) + (1e10 * v3))))))) - v3))
                            v79 := GenS10{F0: 0, F1: -9223372036854775808, F2: "\n\t"}
                            var v80 float64 = (float64(((v3 / v3) / v6)) - 1e10)
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            if false {
                                v81 := []int{fi75, 202, (^(fi75)), (^(fi75))}
                                v82 := GenS11{F0: 848}
                                switch -816 {
                                case 128:
                                    if (!v38) {
                                        switch fi75 {
                                        case -2147483648:
                                            switch fi75 {
                                            case 128:
                                                _ = len(v5)
                                            case 779:
                                                v38 = v38
                                                v14 = v14
                                            }
                                        }
                                        var v83 int = (-(fi75))
                                        _ = v83
                                    }
                                default:
                                    if (int16(v1) == -1) {
                                        var v84 int = 42
                                        fmt.Println(v4, v14, v1)
                                        v85 := v5["a"]
                                        _ = v84
                                        _ = v85
                                    } else {
                                        var v86 uint8 = 61
                                        _ = v86
                                    }
                                }
                                _ = v81
                                _ = v82
                            }
                        }
                        if false {
                            delete(v5, (("hello" + "foo bar") + "hello"))
                        }
                        var v87 uint8 = 1
                        if (((3.14159 + v39) / (float64(1.5) * v39)) >= 1e10) {
                            v3 += v3
                        } else {
                            v3 = 1.5
                        }
                        _ = v87
                    }
                    if (v38 && v38) {
                        if (v39 == float64(3.14159)) {
                            var v88 int = fi74
                            _ = v88
                        } else {
                            v89 := []int{fi75, (fi74 / ((32768) | 1))}
                            v90 := GenS1{F0: 9223372036854775807, F1: (v4 * v4), F2: (v1 >> (((v4 << (((^(v4))) & 63))) & 63))}
                            _ = v89
                            _ = v90
                        }
                        switch "hello" {
                        case "":
                            switch "foo bar" {
                            case "foo bar":
                                var v91 uint32 = v2
                                _ = v91
                            case "":
                                switch "\n\t" {
                                case "foo bar":
                                    v38 = (!false)
                                }
                                var v92 float32 = float32(v39)
                                v93 := []int{(^(fi75)), 9223372036854775807, fi75}
                                _ = v92
                                _ = v93
                            case "hello":
                                var v94 uint = (v14 / (((^(v4))) | 1))
                                _ = v94
                            default:
                                v38 = (42 != 793)
                            }
                        case "a":
                            v2 = v2
                        }
                        v95 := GenS12{F0: ("a" + ((("hello" + "a") + (("\n\t" + ("" + "")) + "a")) + "hello")), F1: 1000}
                        var v96 uint32 = (v2 - 42)
                        v97 := GenS10{F0: int64(v96), F1: (-(fi75)), F2: "\n\t"}
                        _ = v95
                        _ = v96
                        _ = v97
                    } else {
                        if (-306 > int32(v14)) {
                            var v98 uint8 = 1
                            v99 := GenS9{F0: 0, F1: 42, F2: ("\n\t" + "\n\t")}
                            _ = v98
                            _ = v99
                        }
                        if (!(2 >= 82)) {
                            v3 -= (3.14159 + v3)
                            var v100 uint = v14
                            var v101 int64 = 9223372036854775807
                            _ = v100
                            _ = v101
                        } else {
                            fmt.Println(fi74, v1, v1)
                            v102 := GenS12{F0: "", F1: 0}
                            _ = v102
                        }
                    }
                    var v103 uint = (v4 >> ((v14) & 63))
                    _ = fi75
                    _ = v103
                }
                _ = fi74
            }
            var v104 uint32 = (^(v2))
            _ = v104
        }
    case "hello":
        v5[""] = (((("hello" + ("\n\t" + ("a" + "hello"))) + ("" + "hello")) + (("" + ("foo bar" + "")) + ("\n\t" + "foo bar"))) + "a")
    case "a":
        fmt.Println(v39, v2)
    default:
        if v38 {
            var v105 int8 = v1
            v106 := GenS13{F0: v4}
            v107 := v5["\n\t"]
            _ = v105
            _ = v106
            _ = v107
        }
    }
    var v108 uint32 = 996
    for fi109 := 0; fi109 < 14; fi109++ {
        fmt.Println(v39)
        var v110 int8 = v1
        _ = fi109
        _ = v110
    }
    v1 = (-(v1))
    v5["foo bar"] = "hello"
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
