package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int64 }
func genFn0(p0 int16) bool {
    v7 := []int{719, -1, -57}
    v8 := GenS1{F0: 255, F1: 42, F2: 4294967295}
    _ = p0
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 int8, p1 uint32, p2 uint8) bool {
    p1 = p1
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := []int{948, 2147483647, fi9}
        var v11 uint32 = (p1 / (((p1 << ((255) & 63))) | 1))
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS2{F0: fi9, F1: (p1 << ((0) & 63))}
            _ = len(v10)
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v14 := []int{115, -1, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 uint64 = 11
    var v4 bool = false
    var v5 uint64 = v3
    var v6 int64 = -9223372036854775808
    if true {
        v5 += v3
        var v15 uint8 = 127
        v3 = 4294967295
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := v1[fi16]
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 4294967295, F1: 2}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := []int{}
        v6 = -1
        fmt.Println(v6, v4)
        _ = fi19
        _ = v20
    }
    v4 = v4
    fmt.Println(v5)
    v21 := GenS1{F0: 246, F1: 1, F2: -1}
    var v22 int = 0
    for fi23 := 0; fi23 < 13; fi23++ {
        if v4 {
            if v4 {
                v24 := GenS1{F0: 127, F1: 85, F2: v22}
                var v25 uint32 = 2147483647
                v26 := []int{v22, (fi23 & (v22 % (((fi23 / ((int(v3)) | 1))) | 1))), (-(v22))}
                _ = v24
                _ = v25
                _ = v26
            } else {
                if true {
                    switch "a" {
                    case "a":
                        for fi27 := 0; fi27 < 2; fi27++ {
                            if (65535 <= 32767) {
                                fmt.Println(v3, fi23)
                            }
                            switch ("\n\t" + "") {
                            case "foo bar":
                                v28 := v2[2147483648]
                                _ = v28
                            case "":
                                v3 -= (v5 >> ((336) & 63))
                                v6 = v6
                            case "\n\t":
                                v6 = int64(fi23)
                            }
                            _ = fi27
                        }
                    case "hello":
                        v22 = 127
                    }
                    var v29 uint8 = 37
                    _ = v29
                } else {
                    v22 = v22
                    var v30 uint = 255
                    _ = v30
                }
                if v4 {
                    v31 := GenS1{F0: 25, F1: 65535, F2: fi23}
                    v32 := v1[-2147483648]
                    _ = v31
                    _ = v32
                }
            }
            var v33 string = "\n\t"
            v6 += 9223372036854775807
            _ = v33
        } else {
            switch "foo bar" {
            case "hello":
                var v34 uint32 = 42
                _ = v34
            case "a":
                v35 := GenS3{F0: v22, F1: 255}
                _ = v35
            default:
                if (!(-947 > -1)) {
                    fmt.Println(v3, v6, v4)
                    var v36 float32 = 1.5
                    if (!false) {
                        var v37 string = ("hello" + "")
                        for fi38 := 0; fi38 < 15; fi38++ {
                            if v4 {
                                var v39 int = fi23
                                var v40 int8 = 127
                                _ = v39
                                _ = v40
                            } else {
                                if (!((90 <= 0) || v4)) {
                                    v4 = true
                                    if ((2 < 65535) || (v4 && ((v37 == "hello") && false))) {
                                        var v41 uint64 = (^(v3))
                                        var v42 uint = 18446744073709551615
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if v4 {
                                            fmt.Println(v6, v3)
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                v5 -= v3
                                            case "":
                                                _ = len(v2)
                                            case "\n\t":
                                                v5 += 18446744073709551615
                                            }
                                        } else {
                                            if v4 {
                                                if v4 {
                                                    var v43 int32 = -2147483648
                                                    switch 42 {
                                                    case -328:
                                                        switch -279 {
                                                        case 2147483647:
                                                            v44 := []int{-2147483648}
                                                            _ = v44
                                                        default:
                                                            if false {
                                                                switch 2147483647 {
                                                                case 4294967295:
                                                                    var v45 uint16 = 162
                                                                    _ = v45
                                                                case 127:
                                                                    v46 := GenS2{F0: (fi38 ^ v22), F1: 0}
                                                                    v22 = (fi23 & fi23)
                                                                    _ = v46
                                                                case 81:
                                                                    v37 = (v37 + v37)
                                                                }
                                                                switch v37 {
                                                                case "\n\t":
                                                                    if v4 {
                                                                        v36 += v36
                                                                    }
                                                                case "a":
                                                                    var v47 int64 = (v6 | 1)
                                                                    _ = v47
                                                                }
                                                            }
                                                        }
                                                    case 42:
                                                        switch v37 {
                                                        case "\n\t":
                                                            var v48 int = (fi23 + -730)
                                                            _ = v48
                                                        case "":
                                                            v49 := GenS4{F0: 1}
                                                            _ = v49
                                                        case "hello":
                                                            v50 := v2[2147483647]
                                                            _ = v50
                                                        default:
                                                            var v51 uint32 = 2
                                                            _ = v51
                                                        }
                                                    case 4294967295:
                                                        if (v4 || (0.1 != 0.1)) {
                                                            v52 := []int{230, 9223372036854775807}
                                                            if v4 {
                                                                switch v37 {
                                                                case "foo bar":
                                                                    var v53 string = ("\n\t" + "foo bar")
                                                                    _ = v53
                                                                }
                                                            }
                                                            v4 = (!(!v4))
                                                            _ = v52
                                                        }
                                                    }
                                                    v54 := []int{}
                                                    _ = v43
                                                    _ = v54
                                                }
                                                v6 = (-(v6))
                                            } else {
                                                var v55 uint64 = 711
                                                v56 := []int{fi38, (fi38 | (v22 % (((fi23 & fi38)) | 1))), 0}
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                        fmt.Println(v6, v4, fi23)
                                        var v57 uint = 331
                                        _ = v57
                                    }
                                    v58 := []int{fi38, fi38, (fi38 + v22)}
                                    _ = v58
                                }
                                v37 = (v37 + "a")
                                v4 = (!v4)
                            }
                            var v59 int8 = 1
                            if (42 > 42) {
                                v6 -= v6
                                v6 = v6
                                var v60 int = 0
                                _ = v60
                            }
                            _ = fi38
                            _ = v59
                        }
                        if false {
                            _ = len(v2)
                            var v61 uint64 = (v3 % ((v3) | 1))
                            _ = v61
                        } else {
                            if true {
                                if v4 {
                                    v36 = (((v36 - v36) + ((3.14159 - v36) - (((v36 / float32((3.14159 - (0.0 + 0.0)))) + v36) + (v36 - ((1e10 + (-2.25 / (((v36 / v36) - (v36 + ((v36 + ((v36 * (1.5 * v36)) * 0.0)) - float32(((1e10 - 3.14159) + 1.5))))) + (v36 / (-2.25 - (v36 + v36)))))) + (1e10 + v36)))))) / v36)
                                }
                                if (((v37 + ("\n\t" + v37)) + v37) <= ("" + v37)) {
                                    v2[0] = v37
                                    for fi62 := 0; fi62 < 9; fi62++ {
                                        switch "" {
                                        case "hello":
                                            v63 := []int{32768, 39}
                                            if len(v63) > 0 {
                                                v63[0] = (^(v22))
                                            }
                                            var v64 uint8 = 253
                                            _ = v63
                                            _ = v64
                                        default:
                                            v65 := GenS2{F0: fi23, F1: 1}
                                            _ = v65
                                        }
                                        v66 := GenS5{F0: 1, F1: (^(v3)), F2: 36}
                                        _ = fi62
                                        _ = v66
                                    }
                                }
                                for fi67 := 0; fi67 < 17; fi67++ {
                                    v68 := GenS2{F0: fi67, F1: 65535}
                                    _ = fi67
                                    _ = v68
                                }
                            }
                            for fi69 := 0; fi69 < 8; fi69++ {
                                v70 := GenS6{F0: 4294967295, F1: v22}
                                _ = fi69
                                _ = v70
                            }
                            v4 = v4
                        }
                        _ = v37
                    } else {
                        switch (v22 << ((0) & 63)) {
                        case 127:
                            var v71 uint = 2
                            _ = v71
                        case -579:
                            v72 := []int{v22, (v22 >> ((18446744073709551615) & 63)), v22, (^(fi23))}
                            _ = v72
                        default:
                            switch "a" {
                            case "hello":
                                var v73 int16 = -269
                                _ = v73
                            case "":
                                if (0 >= 65535) {
                                    v74 := []int{v22, fi23, (v22 - int(v6))}
                                    fmt.Println(v6)
                                    var v75 uint = 0
                                    _ = v74
                                    _ = v75
                                } else {
                                    if (uint32(v3) == 398) {
                                        fmt.Println(v6, v6)
                                        var v76 bool = v4
                                        _ = v76
                                    }
                                    fmt.Println(v3)
                                }
                            case "foo bar":
                                if (255 > 18446744073709551615) {
                                    var v77 uint32 = 65535
                                    _ = v77
                                } else {
                                    var v78 uint = 902
                                    var v79 float32 = v36
                                    _ = v78
                                    _ = v79
                                }
                            }
                        }
                        fmt.Println(fi23, v6)
                        var v80 uint64 = (v3 - v3)
                        _ = v80
                    }
                    _ = v36
                } else {
                    if v4 {
                        var v81 uint32 = 683
                        fmt.Println(v4, v22)
                        _ = v81
                    } else {
                        var v82 uint64 = 95
                        var v83 int8 = int8(v82)
                        if v4 {
                            v84 := GenS2{F0: v22, F1: 4294967295}
                            v85 := []int{}
                            v86 := v2[838]
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                        _ = v82
                        _ = v83
                    }
                    v87 := GenS7{F0: 397, F1: 464}
                    _ = v87
                }
            }
            fmt.Println(v6)
            switch "foo bar" {
            case "":
                if v4 {
                    v3 = v3
                    var v88 bool = (true || v4)
                    v3 -= v3
                    _ = v88
                }
            case "a":
                if (1 > 580) {
                    fmt.Println(fi23, v5, v6)
                    v4 = ((^(fi23)) < (v22 & (v22 ^ (v22 * 161))))
                    v5 = v3
                } else {
                    switch (fi23 * 128) {
                    case -66:
                        v89 := []int{4294967295}
                        _ = v89
                    }
                    var v90 float64 = 1e10
                    _ = v90
                }
            case "foo bar":
                fmt.Println(fi23, v3, fi23)
            }
        }
        v91 := GenS3{F0: 103, F1: 140}
        _ = fi23
        _ = v91
    }
    v92 := GenS8{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
