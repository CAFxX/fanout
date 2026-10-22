package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int16; F1 int16; F2 int }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 uint32; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint16; F1 int }
func genFn0(p0 uint64, p1 bool) int16 {
    var v5 int8 = -128
    v6 := GenS1{F0: 4294967295}
    v7 := GenS2{F0: 2, F1: p0}
    var v8 uint = 65535
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint16, p1 bool, p2 int16) int64 {
    switch "" {
    case "\n\t":
        v9 := GenS2{F0: 42, F1: 2}
        _ = v9
    case "foo bar":
        v10 := GenS3{F0: -851, F1: -9223372036854775808, F2: genFn0(442, ((4294967295 >= -1) || p1))}
        _ = v10
    }
    v11 := GenS2{F0: 42, F1: 633}
    p1 = (42 <= 192)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS3{F0: 9223372036854775807, F1: 1, F2: p2}
        v14 := []int{32767, 457}
        p1 = p1
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 int16 = -492
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]uint32)
    var v4 bool = (128 != 1)
    var v15 int16 = v1
    if (0 < -973) {
        v2 = (v2 - (v2 + 42))
        var v16 bool = (1 == 0)
        for fi17 := 0; fi17 < 5; fi17++ {
            v2 += (v2 << ((0) & 63))
            fmt.Println(v4)
            _ = fi17
        }
        _ = v16
    } else {
        v1 += (v15 + v1)
        fmt.Println(v4, v1)
    }
    switch 565 {
    case 128:
        switch "hello" {
        case "a":
            switch ("" + "a") {
            case "a":
                if (128 <= 684) {
                    if false {
                        v1 = v1
                        var v18 uint8 = 113
                        _ = v18
                    } else {
                        v19 := GenS4{F0: -701, F1: 2}
                        _ = v19
                    }
                    v3[4294967295] = 2147483648
                    var v20 int64 = -9223372036854775808
                    _ = v20
                }
            case "\n\t":
                v21 := GenS4{F0: -615, F1: 802}
                _ = v21
            case "foo bar":
                for fi22 := 0; fi22 < 11; fi22++ {
                    delete(v3, 128)
                    _ = fi22
                }
            }
        case "":
            for fi23 := 0; fi23 < 13; fi23++ {
                switch "" {
                case "":
                    switch -894 {
                    case -526:
                        switch "hello" {
                        case "foo bar":
                            for fi24 := 0; fi24 < 2; fi24++ {
                                v25 := GenS3{F0: 4294967295, F1: -169, F2: v1}
                                var v26 int32 = int32(v2)
                                v2 -= v2
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        case "\n\t":
                            for fi27 := 0; fi27 < 2; fi27++ {
                                var v28 int16 = (v15 * (v15 / ((v1) | 1)))
                                v29 := GenS4{F0: int(v1), F1: 4294967295}
                                v30 := GenS5{F0: 1000, F1: (v28 << ((4294967295) & 63)), F2: 903}
                                _ = fi27
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                        case "hello":
                            v31 := GenS6{F0: (v2 / ((v2) | 1)), F1: 9223372036854775807}
                            _ = v31
                        default:
                            var v32 string = "hello"
                            _ = v32
                        }
                    case 1:
                        v4 = v4
                    case -2147483648:
                        var v33 float64 = -2.25
                        _ = v33
                    default:
                        for fi34 := 0; fi34 < 10; fi34++ {
                            if (("\n\t" + "foo bar") == "") {
                                switch (fi23 >> ((116) & 63)) {
                                case 443:
                                    var v35 int16 = (v15 << ((0) & 63))
                                    _ = v35
                                }
                                var v36 float64 = 1.5
                                _ = v36
                            } else {
                                var v37 uint32 = 0
                                _ = v37
                            }
                            if (!(65535 <= 42)) {
                                v4 = (!(2 == 42))
                            } else {
                                var v38 uint8 = 1
                                var v39 bool = ((((1e10 / (1.5 / 1.0)) * 0.1) <= 0.0) || (v15 == 32767))
                                v2 -= v2
                                _ = v38
                                _ = v39
                            }
                            _ = fi34
                        }
                    }
                case "foo bar":
                    if (false || false) {
                        var v40 int32 = -1
                        _ = v40
                    } else {
                        if ((4294967295 < 2147483648) && (395 != 0)) {
                            v41 := GenS1{F0: 1}
                            var v42 uint64 = v2
                            if (0 <= 200) {
                                if (!((true && (1 <= 18446744073709551615)) || v4)) {
                                    v43 := GenS6{F0: 18446744073709551615, F1: 0}
                                    var v44 int8 = 127
                                    _ = v43
                                    _ = v44
                                } else {
                                    for fi45 := 0; fi45 < 2; fi45++ {
                                        if false {
                                            v4 = (3.14159 > (3.14159 * 1.5))
                                        } else {
                                            v1 -= (v15 >> ((255) & 63))
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v46 uint16 = 0
                                                _ = v46
                                            case "":
                                                v3[(fi45 >> ((703) & 63))] = 0
                                            case "\n\t":
                                                switch -1 {
                                                case -1000:
                                                    var v47 uint8 = 0
                                                    _ = v47
                                                case 170:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v15 = 32767
                                                    case "\n\t":
                                                        if v4 {
                                                            switch fi23 {
                                                            case -540:
                                                                if v4 {
                                                                    var v48 int = fi23
                                                                    var v49 float32 = 3.14159
                                                                    var v50 uint8 = 0
                                                                    var v51 int32 = -621
                                                                    var v52 uint16 = 32767
                                                                    switch -580 {
                                                                    case 4294967295:
                                                                        var v53 uint32 = 1
                                                                        _ = v53
                                                                    case 128:
                                                                        switch "hello" {
                                                                        case "":
                                                                            fmt.Println(v4, fi45)
                                                                        case "hello":
                                                                            var v54 int16 = v15
                                                                            _ = v54
                                                                        case "a":
                                                                            v51 = -488
                                                                        }
                                                                    case 42:
                                                                        v55 := GenS1{F0: 875}
                                                                        _ = v55
                                                                    default:
                                                                        v2 += (v2 - v2)
                                                                    }
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            case 102:
                                                                v2 = v42
                                                                if ("foo bar" == "hello") {
                                                                    switch fi23 {
                                                                    case 1:
                                                                        v56 := []int{fi45}
                                                                        _ = v56
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        v4 = (!v4)
                                                                        switch (fi23 - fi45) {
                                                                        case -1:
                                                                            v42 -= 166
                                                                        }
                                                                        v57 := GenS7{F0: 2}
                                                                        _ = v57
                                                                    }
                                                                    v58 := GenS1{F0: 2147483647}
                                                                    v1 = v1
                                                                    _ = v58
                                                                }
                                                            case 127:
                                                                var v59 float64 = float64(1e10)
                                                                _ = v59
                                                            default:
                                                                var v60 int = (fi23 / ((fi45) | 1))
                                                                _ = v60
                                                            }
                                                            v61 := GenS1{F0: 65535}
                                                            fmt.Println(fi45, v4, v1)
                                                            _ = v61
                                                        } else {
                                                            v62 := []int{(fi23 / ((32767) | 1))}
                                                            _ = v62
                                                        }
                                                    }
                                                }
                                            default:
                                                v2 += 2
                                            }
                                            if false {
                                                v63 := GenS2{F0: 2, F1: (v2 & v2)}
                                                if (((2147483647 < 2) && (!true)) && v4) {
                                                    var v64 int32 = 758
                                                    _ = v64
                                                } else {
                                                    v65 := GenS3{F0: 2147483647, F1: genFn1(756, (65535 > 65535), v1), F2: -1}
                                                    v1 = 0
                                                    _ = v65
                                                }
                                                _ = v63
                                            }
                                        }
                                        switch (fi23 | fi23) {
                                        case 0:
                                            var v66 int16 = genFn0((v2 % (((v42 ^ v42)) | 1)), v4)
                                            _ = v66
                                        case 974:
                                            v67 := []int{(fi45 * (-(fi45)))}
                                            _ = v67
                                        default:
                                            v68 := GenS8{F0: -9223372036854775808}
                                            _ = v68
                                        }
                                        var v69 uint8 = 42
                                        var v70 float64 = 0.1
                                        switch "\n\t" {
                                        case "":
                                            switch (fi23 % (((fi45 << ((28) & 63))) | 1)) {
                                            case 128:
                                                var v71 int8 = -101
                                                var v72 int8 = v71
                                                var v73 string = ""
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            default:
                                                switch "a" {
                                                case "\n\t":
                                                    var v74 int32 = -372
                                                    _ = v74
                                                case "":
                                                    if v4 {
                                                        v75 := GenS9{F0: 65535, F1: 42, F2: int8(v1)}
                                                        _ = v75
                                                    } else {
                                                        var v76 int32 = 2147483647
                                                        _ = v76
                                                    }
                                                case "hello":
                                                    v77 := GenS10{F0: fi23}
                                                    v78 := v3[0]
                                                    _ = v77
                                                    _ = v78
                                                }
                                            }
                                        case "a":
                                            switch "hello" {
                                            case "\n\t":
                                                var v79 int16 = (v1 ^ v15)
                                                _ = v79
                                            case "foo bar":
                                                v80 := GenS2{F0: v69, F1: (v42 - (v2 % (((v2 << ((65535) & 63))) | 1)))}
                                                _ = v80
                                            default:
                                                v81 := GenS3{F0: (fi23 & fi45), F1: 42, F2: (v1 % ((32767) | 1))}
                                                _ = v81
                                            }
                                            if v4 {
                                                var v82 float32 = ((-2.25 - 0.1) * (1e10 * (float32(0.1) / 1e10)))
                                                switch 1 {
                                                case 2147483647:
                                                    fmt.Println(v2, fi45, v70)
                                                }
                                                _ = v82
                                            }
                                            v83 := GenS11{F0: (v42 | (^(v2)))}
                                            _ = v83
                                        }
                                        _ = len(v3)
                                        _ = fi45
                                        _ = v69
                                        _ = v70
                                    }
                                    v84 := GenS7{F0: 337}
                                    _ = v84
                                }
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v15, v42)
                                case "\n\t":
                                    for fi85 := 0; fi85 < 11; fi85++ {
                                        v86 := GenS12{F0: 128}
                                        v2 += v2
                                        v87 := []int{fi23, fi23, (^(fi23)), (fi85 % (((fi85 << ((2) & 63))) | 1))}
                                        _ = fi85
                                        _ = v86
                                        _ = v87
                                    }
                                default:
                                    v2 = 1
                                }
                            } else {
                                var v88 uint8 = 129
                                v89 := []int{fi23, (fi23 / (((fi23 / (((fi23 * (fi23 & fi23))) | 1))) | 1))}
                                _ = v88
                                _ = v89
                            }
                            _ = v41
                            _ = v42
                        } else {
                            for fi90 := 0; fi90 < 15; fi90++ {
                                v91 := GenS13{F0: 42, F1: (fi90 / (((fi23 * 1)) | 1))}
                                _ = fi90
                                _ = v91
                            }
                            if false {
                                v2 -= (v2 & 42)
                            } else {
                                v92 := GenS9{F0: 2, F1: uint32(v2), F2: -33}
                                _ = v92
                            }
                        }
                        var v93 uint = 462
                        v94 := []int{}
                        _ = v93
                        _ = v94
                    }
                case "hello":
                    v2 -= v2
                }
                v95 := GenS2{F0: 18, F1: v2}
                v1 += v15
                _ = fi23
                _ = v95
            }
        case "\n\t":
            for fi96 := 0; fi96 < 18; fi96++ {
                v15 = (^(v15))
                _ = fi96
            }
        }
    case 32768:
        v97 := GenS4{F0: 685, F1: 1}
        _ = v97
    case 9223372036854775807:
        var v98 int32 = -2147483648
        _ = v98
    }
    v3[319] = 949
    var v99 uint8 = 1
    var v100 uint64 = v2
    var v101 uint = 4
    var v102 int = 32767
    var v103 int32 = -2147483648
    var v104 int = (v102 >> ((v101) & 63))
    var v105 uint = v101
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v99
}
