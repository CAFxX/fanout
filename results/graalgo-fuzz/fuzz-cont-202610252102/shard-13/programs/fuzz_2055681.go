package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int64; F1 int8; F2 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 uint; F2 int32 }
type GenS10 struct { F0 int32; F1 uint64; F2 int8 }
type GenS11 struct { F0 int32; F1 int16; F2 bool }
type GenS12 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 string, p1 uint64, p2 int) uint8 {
    if (!((467 != 0) || false)) {
        v7 := GenS1{F0: 1, F1: 0}
        var v8 uint8 = 166
        p0 = p0
        _ = v7
        _ = v8
    } else {
        var v9 float32 = float32(1e10)
        _ = v9
    }
    if false {
        var v10 uint = 437
        for fi11 := 0; fi11 < 20; fi11++ {
            p0 = p0
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int32 = -2147483648
                _ = fi12
                _ = v13
            }
            v10 += v10
            _ = fi11
        }
        v14 := []int{p2, (p2 - p2), p2, (p2 * p2)}
        _ = v10
        _ = v14
    } else {
        p0 = (p0 + "foo bar")
        switch "\n\t" {
        case "a":
            v15 := GenS2{F0: p1}
            _ = v15
        case "\n\t":
            if false {
                if false {
                    if false {
                        v16 := GenS3{F0: (p2 >> ((65535) & 63))}
                        _ = v16
                    } else {
                        switch (p2 & p2) {
                        case 973:
                            v17 := []int{(p2 << ((2) & 63)), -794, 42, p2}
                            _ = v17
                        case -2147483648:
                            var v18 float32 = ((-2.25 - (float32(1e10) / (1e10 * 1.5))) - 3.14159)
                            _ = v18
                        case 128:
                            v19 := []int{p2, p2}
                            _ = v19
                        default:
                            for fi20 := 0; fi20 < 19; fi20++ {
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    v22 := GenS2{F0: p1}
                                    _ = fi21
                                    _ = v22
                                }
                                _ = fi20
                            }
                        }
                        var v23 int16 = 1000
                        p2 += p2
                        _ = v23
                    }
                }
            } else {
                p0 = p0
            }
        case "foo bar":
            p1 += 42
        }
    }
    switch p2 {
    case 127:
        p2 += (p2 ^ 32767)
    case -732:
        v24 := GenS3{F0: p2}
        _ = v24
    case -2147483648:
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(fi25, p0, p2)
            v26 := GenS1{F0: 657, F1: 1}
            _ = fi25
            _ = v26
        }
    }
    p2 = p2
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 bool = (!(18446744073709551615 == 18446744073709551615))
            if (0 <= -730) {
                p1 -= p1
                switch -326 {
                case 42:
                    if false {
                        if v29 {
                            v30 := GenS4{F0: 69}
                            _ = v30
                        }
                        fmt.Println(fi28, v29)
                        if v29 {
                            var v31 float64 = 3.14159
                            var v32 int64 = 9223372036854775807
                            var v33 int = (fi27 * -448)
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            var v34 float64 = 1e10
                            _ = v34
                        }
                        v35 := []int{9223372036854775807, fi27}
                        _ = v35
                    }
                case 32767:
                    var v36 uint8 = 196
                    _ = v36
                case 443:
                    if (v29 && v29) {
                        v37 := []int{-9223372036854775808, 9223372036854775807, fi28, 1}
                        switch "hello" {
                        case "foo bar":
                            var v38 uint8 = 128
                            _ = v38
                        case "a":
                            switch "a" {
                            case "":
                                var v39 string = "foo bar"
                                _ = v39
                            }
                        }
                        switch p0 {
                        case "a":
                            v40 := GenS2{F0: (^(p1))}
                            _ = v40
                        case "hello":
                            if (!v29) {
                                v29 = ((-2.25 * float64(-2.25)) > (float64(-2.25) / 0.1))
                            } else {
                                p1 = p1
                                fmt.Println(p2, p1)
                                v41 := GenS5{F0: 0, F1: 166}
                                var v42 int = 0
                                p1 = p1
                                if len(v37) > 0 {
                                    v37[0] = 1
                                }
                                _ = v41
                                _ = v42
                            }
                        case "foo bar":
                            if (!(!(32767 <= (p2 ^ (fi28 | 912))))) {
                                p2 += 32768
                                p0 = ((((((((p0 + ((p0 + "hello") + (p0 + (p0 + "a")))) + "a") + p0) + "\n\t") + "\n\t") + ((p0 + (p0 + p0)) + (p0 + p0))) + (p0 + (p0 + "a"))) + "foo bar")
                                switch p0 {
                                case "hello":
                                    if (-499 != 1) {
                                        var v43 uint16 = 1
                                        v43 -= v43
                                        v44 := GenS2{F0: (p1 ^ (p1 + p1))}
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if true {
                                            _ = len(v37)
                                        } else {
                                            var v45 uint16 = 2
                                            v37 = append(v37, (p2 << ((42) & 63)))
                                            if (!(!v29)) {
                                                v46 := GenS6{F0: -1, F1: -27}
                                                var v47 uint8 = 154
                                                _ = v46
                                                _ = v47
                                            } else {
                                                if v29 {
                                                    _ = len(v37)
                                                }
                                                if (0 == 2147483647) {
                                                    v45 = (^(v45))
                                                }
                                                v48 := GenS7{F0: -1, F1: 100, F2: v45}
                                                p2 = (p2 & fi28)
                                                _ = v48
                                            }
                                            _ = v45
                                        }
                                    }
                                default:
                                    var v49 uint8 = 128
                                    _ = v49
                                }
                            }
                        default:
                            fmt.Println(p1)
                        }
                        _ = v37
                    } else {
                        p1 = p1
                        v50 := []int{-958, fi27}
                        _ = len(v50)
                        _ = v50
                    }
                default:
                    switch "" {
                    case "hello":
                        v51 := []int{}
                        _ = v51
                    case "":
                        var v52 int32 = 2147483647
                        _ = v52
                    case "a":
                        v53 := []int{(fi27 / ((p2) | 1)), p2, fi27}
                        _ = v53
                    }
                }
                if v29 {
                    if v29 {
                        v54 := GenS5{F0: 811, F1: -500}
                        switch (p2 * (fi27 >> ((159) & 63))) {
                        case 1:
                            if v29 {
                                if v29 {
                                    p0 = p0
                                    v55 := []int{(-(p2)), -579, 109}
                                    _ = v55
                                } else {
                                    v29 = v29
                                }
                            } else {
                                fmt.Println(p1, p0, fi27)
                            }
                        case 4294967295:
                            p2 -= (-(fi27))
                        default:
                            switch 4294967295 {
                            case 940:
                                var v56 float64 = (1.5 / ((((3.14159 - (3.14159 / -2.25)) * 1e10) - ((1e10 / 1.5) / (-2.25 * -2.25))) / ((0.1 * 1e10) + 0.1)))
                                _ = v56
                            case 42:
                                if v29 {
                                    if v29 {
                                        switch p0 {
                                        case "foo bar":
                                            fmt.Println(p0, p2)
                                        case "hello":
                                            var v57 int = p2
                                            _ = v57
                                        default:
                                            if false {
                                                var v58 float64 = 1e10
                                                var v59 float64 = float64(float32((0.1 * 3.14159)))
                                                _ = v58
                                                _ = v59
                                            }
                                        }
                                    }
                                    p2 = 32768
                                    v60 := []int{-2147483648, (^(fi27))}
                                    _ = v60
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if v29 {
                                            v61 := GenS8{F0: ((-2.25 >= 1e10) && (989 != 255))}
                                            switch fi28 {
                                            case -9223372036854775808:
                                                if (1 < -2147483648) {
                                                    switch p2 {
                                                    case 859:
                                                        if (!(v29 && (((p1 << ((255) & 63)) >= 92) && v29))) {
                                                            v62 := []int{}
                                                            var v63 int16 = 0
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            switch p0 {
                                                            case "\n\t":
                                                                v64 := GenS9{F0: 232, F1: 65535, F2: 0}
                                                                _ = v64
                                                            }
                                                            p0 = ""
                                                        }
                                                    case 32768:
                                                        v65 := []int{(fi27 + (fi27 + fi27)), (fi28 * (p2 | fi28)), 32768, 4294967295}
                                                        _ = v65
                                                    }
                                                    if v29 {
                                                        switch (p0 + p0) {
                                                        case "a":
                                                            v66 := GenS10{F0: 42, F1: (^(p1)), F2: 98}
                                                            _ = v66
                                                        default:
                                                            p1 += p1
                                                        }
                                                    }
                                                } else {
                                                    var v67 int8 = -109
                                                    p0 = ""
                                                    v68 := GenS11{F0: -217, F1: 42, F2: (v29 && false)}
                                                    _ = v67
                                                    _ = v68
                                                }
                                            default:
                                                var v69 uint = 65535
                                                _ = v69
                                            }
                                            fmt.Println(p0, fi27, p2)
                                            _ = v61
                                        }
                                    case "":
                                        p1 -= p1
                                    }
                                    p0 = "hello"
                                    switch 127 {
                                    case 9223372036854775807:
                                        var v70 float32 = 3.14159
                                        _ = v70
                                    case -717:
                                        if (!v29) {
                                            if false {
                                                v71 := GenS9{F0: 427, F1: 18446744073709551615, F2: 42}
                                                if (1e10 <= 0.1) {
                                                    v72 := GenS8{F0: (1 != 2147483648)}
                                                    var v73 int16 = 42
                                                    _ = v72
                                                    _ = v73
                                                }
                                                _ = v71
                                            } else {
                                                var v74 bool = false
                                                v29 = ((767 < -1) && v74)
                                                _ = v74
                                            }
                                        } else {
                                            var v75 uint16 = 32767
                                            switch (^(fi27)) {
                                            case 0:
                                                p0 = ""
                                            case 127:
                                                v76 := []int{fi27, (-(p2))}
                                                _ = v76
                                            }
                                            v77 := []int{(^(fi28)), (p2 >> ((4294967295) & 63)), (p2 - (fi27 * fi28))}
                                            _ = v75
                                            _ = v77
                                        }
                                    default:
                                        var v78 int = (p2 & (fi28 / (((fi27 + (p2 / ((int(p1)) | 1)))) | 1)))
                                        _ = v78
                                    }
                                }
                            case -9223372036854775808:
                                switch (^(fi28)) {
                                case 127:
                                    p0 = (p0 + "")
                                default:
                                    v79 := GenS11{F0: 42, F1: 42, F2: (("\n\t" + p0) == (p0 + p0))}
                                    _ = v79
                                }
                                p0 = "\n\t"
                            }
                        }
                        _ = v54
                    } else {
                        v80 := GenS12{F0: 48, F1: -2147483648, F2: (p0 + (p0 + (p0 + "")))}
                        _ = v80
                    }
                }
                if (!(94 > 128)) {
                    var v81 float32 = -2.25
                    _ = v81
                }
            } else {
                p2 -= (p2 - fi28)
                var v82 uint8 = 1
                v83 := GenS9{F0: 32767, F1: 18446744073709551615, F2: -382}
                _ = v82
                _ = v83
            }
            v84 := GenS6{F0: 1, F1: 84}
            _ = fi28
            _ = v29
            _ = v84
        }
        p2 = fi27
        _ = fi27
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int16 = 42
    var v2 int32 = 1
    var v3 string = ("foo bar" + "")
    v4 := make(map[string]int)
    var v5 uint16 = 65535
    var v6 string = v3
    if true {
        switch 32767 {
        case 2147483647:
            var v85 float32 = float32(float64(-2.25))
            _ = v85
        default:
            var v86 int64 = 9223372036854775807
            _ = v86
        }
        if false {
            if true {
                for fi87 := 0; fi87 < 4; fi87++ {
                    v88 := []int{(fi87 % ((-925) | 1))}
                    v89 := GenS5{F0: -1, F1: v1}
                    _ = fi87
                    _ = v88
                    _ = v89
                }
            } else {
                v2 = 0
            }
            fmt.Println(v3)
            switch v3 {
            case "":
                var v90 int = 570
                _ = v90
            }
        }
    }
    fmt.Println(v1, v5, v3)
    var v91 bool = (uint64(v5) == 4294967295)
    var v92 uint8 = 255
    v6 = v6
    _ = len(v4)
    fmt.Println(v5)
    var v93 uint = 876
    if v91 {
        v3 = v3
        fmt.Println(v2, v5)
    } else {
        var v94 uint = (v93 - (v93 >> ((0) & 63)))
        for fi95 := 0; fi95 < 14; fi95++ {
            fmt.Println(v94, v2)
            switch v6 {
            case "":
                for fi96 := 0; fi96 < 18; fi96++ {
                    var v97 uint32 = 4294967295
                    v4[v6] = fi96
                    var v98 uint64 = 18446744073709551615
                    _ = fi96
                    _ = v97
                    _ = v98
                }
            case "hello":
                if true {
                    v94 = v94
                    for fi99 := 0; fi99 < 3; fi99++ {
                        var v100 int = fi95
                        fmt.Println(v93, v93)
                        _ = len(v4)
                        var v101 int8 = -70
                        v6 = "foo bar"
                        _ = fi99
                        _ = v100
                        _ = v101
                    }
                }
            }
            _ = fi95
        }
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v92
    _ = v93
}
