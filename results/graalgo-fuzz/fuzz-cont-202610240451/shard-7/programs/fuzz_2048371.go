package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint; F2 int64 }
type GenS11 struct { F0 uint; F1 uint }
func genFn0(p0 bool) int64 {
    p0 = true
    p0 = p0
    if p0 {
        v5 := []int{-387, 375}
        switch "foo bar" {
        case "hello":
            if p0 {
                fmt.Println(p0)
            }
        case "foo bar":
            fmt.Println(p0)
        default:
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(fi6, fi6)
                switch "" {
                case "a":
                    var v7 uint32 = 244
                    _ = v7
                }
                _ = fi6
            }
        }
        var v8 float32 = (3.14159 + (-2.25 * (0.1 + float32(3.14159))))
        _ = v5
        _ = v8
    }
    v9 := []int{9223372036854775807, 32767, 2147483647, -895}
    var v10 uint32 = 896
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    var v2 bool = true
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]int)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            delete(v4, fi11)
            v13 := []int{(fi11 - 9223372036854775807), 32767, fi11, 127}
            _ = fi12
            _ = v13
        }
        var v14 string = ""
        var v15 uint = 4294967295
        _ = fi11
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: 147}
    fmt.Println(v1)
    if (42 > 42) {
        switch 32768 {
        case 1:
            var v17 string = "foo bar"
            _ = v17
        default:
            var v18 float32 = -2.25
            _ = v18
        }
        var v19 uint = 326
        v20 := []int{}
        _ = v19
        _ = v20
    } else {
        if true {
            var v21 string = (("a" + ("hello" + "hello")) + "\n\t")
            v2 = (v1 < v1)
            _ = v21
        } else {
            if (int8(v1) <= 100) {
                fmt.Println(v1)
                switch "hello" {
                case "foo bar":
                    if false {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v4[int(v1)] = fi22
                            v2 = v2
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 7; fi23++ {
                            for fi24 := 0; fi24 < 14; fi24++ {
                                if (!(float32(0.0) == (1.5 * ((1.5 + 3.14159) * -2.25)))) {
                                    if true {
                                        v25 := GenS1{F0: (-(fi23)), F1: fi23}
                                        v3 = 42
                                        _ = v25
                                    } else {
                                        if (v1 == (v1 | v1)) {
                                            var v26 float32 = 0.0
                                            if (-812 == -679) {
                                                var v27 uint64 = v3
                                                v1 -= uint32(v27)
                                                if false {
                                                    if v2 {
                                                        v1 += (v1 + 42)
                                                        fmt.Println(v2)
                                                    } else {
                                                        v28 := []int{fi23, (fi24 >> ((65535) & 63)), (fi24 & 0), 128}
                                                        var v29 bool = (fi24 == fi24)
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    v26 -= v26
                                                    v26 -= v26
                                                } else {
                                                    v1 -= (v1 % (((^(v1))) | 1))
                                                    v26 += v26
                                                }
                                                var v30 uint = 42
                                                switch (-(fi23)) {
                                                case 2147483647:
                                                    if (!v2) {
                                                        if v2 {
                                                            var v31 uint8 = 5
                                                            v32 := GenS1{F0: fi23, F1: fi24}
                                                            v4[128] = 498
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            v33 := []int{(fi24 >> ((18446744073709551615) & 63)), int(v1), (fi24 * fi24)}
                                                            fmt.Println(v3)
                                                            v30 = v30
                                                            _ = v33
                                                        }
                                                        v3 = (v3 - (v27 & (v3 + v3)))
                                                    } else {
                                                        var v34 int = fi24
                                                        var v35 uint16 = 32767
                                                        switch ("" + "foo bar") {
                                                        case "hello":
                                                            v1 += v1
                                                        case "foo bar":
                                                            v36 := GenS2{F0: (v34 - (fi24 * v34)), F1: false}
                                                            _ = v36
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                case 42:
                                                    switch (fi23 & (^(fi24))) {
                                                    case -1:
                                                        fmt.Println(v2, fi24, v27)
                                                    case 128:
                                                        var v37 uint = v30
                                                        _ = v37
                                                    default:
                                                        var v38 uint32 = 2147483647
                                                        _ = v38
                                                    }
                                                case 127:
                                                    v2 = v2
                                                }
                                                _ = v27
                                                _ = v30
                                            }
                                            v39 := []int{1}
                                            _ = v26
                                            _ = v39
                                        } else {
                                            v1 = (^(v1))
                                            if true {
                                                v2 = v2
                                            } else {
                                                _ = len(v4)
                                                var v40 uint8 = 255
                                                var v41 uint = 18446744073709551615
                                                _ = v40
                                                _ = v41
                                            }
                                            if (-1 <= 9223372036854775807) {
                                                var v42 float64 = 1.5
                                                _ = v42
                                            }
                                            v1 += v1
                                        }
                                        v43 := GenS3{F0: v3, F1: v3, F2: v1}
                                        var v44 string = "foo bar"
                                        _ = v43
                                        _ = v44
                                    }
                                    if v2 {
                                        _ = len(v4)
                                    } else {
                                        if (true || v2) {
                                            v45 := []int{-1, 32767}
                                            v46 := GenS4{F0: (2 < uint16(v1)), F1: 1}
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                }
                                _ = fi24
                            }
                            _ = fi23
                        }
                        for fi47 := 0; fi47 < 6; fi47++ {
                            switch (fi47 | fi47) {
                            case 4294967295:
                                v2 = (99 != 1)
                            case 32768:
                                v48 := GenS5{F0: 42, F1: -1}
                                _ = v48
                            }
                            var v49 int8 = 91
                            v50 := GenS3{F0: (v3 >> ((2) & 63)), F1: v3, F2: v1}
                            _ = fi47
                            _ = v49
                            _ = v50
                        }
                    } else {
                        fmt.Println(v3)
                        for fi51 := 0; fi51 < 9; fi51++ {
                            v52 := []int{42, fi51, fi51, 2147483647}
                            _ = fi51
                            _ = v52
                        }
                        v53 := GenS1{F0: 32768, F1: -9223372036854775808}
                        _ = v53
                    }
                case "a":
                    var v54 int32 = -1
                    _ = v54
                case "hello":
                    if v2 {
                        if v2 {
                            var v55 uint64 = v3
                            v56 := []int{0, -1, 0}
                            _ = v55
                            _ = v56
                        }
                        v57 := GenS6{F0: 1}
                        var v58 int8 = 42
                        _ = v57
                        _ = v58
                    } else {
                        var v59 int64 = -9223372036854775808
                        _ = v59
                    }
                }
            } else {
                v2 = false
            }
            v1 = (^(v1))
        }
        var v60 uint32 = (^(v1))
        if (false || v2) {
            v60 = v60
            fmt.Println(v60, v3)
        } else {
            var v61 float64 = float64(0.0)
            v3 -= (v3 % ((v3) | 1))
            v62 := GenS2{F0: -2147483648, F1: v2}
            _ = v61
            _ = v62
        }
        _ = v60
    }
    var v63 int8 = 127
    for fi64 := 0; fi64 < 4; fi64++ {
        for fi65 := 0; fi65 < 8; fi65++ {
            switch "hello" {
            case "":
                if true {
                    v3 += uint64(v1)
                    switch (-(fi64)) {
                    case 32768:
                        v66 := GenS6{F0: 1}
                        _ = v66
                    case 57:
                        v63 = (v63 ^ v63)
                    default:
                        if v2 {
                            if false {
                                if false {
                                    v67 := []int{(fi64 + fi65), 9223372036854775807, 42, (fi65 / (((fi64 << ((18446744073709551615) & 63))) | 1))}
                                    v2 = (((0.0 / 1.0) * (((1.5 * 1.5) * 1.5) / -2.25)) == -2.25)
                                    _ = v67
                                } else {
                                    v68 := []int{(fi64 * (fi65 >> ((4294967295) & 63))), fi65}
                                    v68 = append(v68, fi64)
                                    var v69 float64 = -2.25
                                    _ = v68
                                    _ = v69
                                }
                            }
                            v3 = 1
                            v70 := GenS1{F0: 32767, F1: (fi64 - fi64)}
                            _ = v70
                        } else {
                            v71 := GenS7{F0: 866, F1: 42}
                            if v2 {
                                fmt.Println(v2, v1, v2)
                                v63 = (v63 ^ v63)
                                if v2 {
                                    v1 = 2147483648
                                    if (1 < 809) {
                                        v72 := GenS8{F0: v2, F1: 1}
                                        var v73 int8 = v63
                                        if (v2 && v2) {
                                            v74 := GenS5{F0: 65535, F1: (v63 % ((v73) | 1))}
                                            if v2 {
                                                switch fi65 {
                                                case -782:
                                                    v75 := v4[128]
                                                    _ = v75
                                                case 32768:
                                                    v73 += v63
                                                default:
                                                    switch ("\n\t" + ("hello" + "hello")) {
                                                    case "a":
                                                        var v76 float32 = (1e10 / (0.0 + -2.25))
                                                        _ = v76
                                                    case "":
                                                        v77 := v4[(fi65 << ((0) & 63))]
                                                        v78 := []int{int(v1), (-(v77)), int(v1)}
                                                        fmt.Println(v63)
                                                        _ = v77
                                                        _ = v78
                                                    case "foo bar":
                                                        switch fi64 {
                                                        case 2147483647:
                                                            if true {
                                                                v73 += 0
                                                                v2 = (v2 || true)
                                                            }
                                                        case 42:
                                                            _ = len(v4)
                                                        case -884:
                                                            v1 -= 2
                                                        default:
                                                            v79 := GenS9{F0: (v3 - v3)}
                                                            _ = v79
                                                        }
                                                    }
                                                }
                                                v73 = v73
                                            }
                                            v80 := []int{(fi64 << ((118) & 63)), (^(fi65)), -620, (fi65 | (fi65 & 127))}
                                            v4[(fi65 % ((-527) | 1))] = (fi64 | int(v1))
                                            _ = v74
                                            _ = v80
                                        } else {
                                            v81 := GenS1{F0: (fi64 / (((fi65 - fi64)) | 1)), F1: (fi65 >> ((18446744073709551615) & 63))}
                                            fmt.Println(v1, fi65)
                                            _ = v81
                                        }
                                        _ = v72
                                        _ = v73
                                    }
                                    switch "hello" {
                                    case "hello":
                                        v1 = (v1 * (^(v1)))
                                    case "foo bar":
                                        v1 -= (v1 ^ (v1 >> ((2) & 63)))
                                    default:
                                        v82 := []int{-299, (fi64 ^ -784), (^(fi65))}
                                        _ = v82
                                    }
                                }
                            }
                            fmt.Println(v3, v1)
                            _ = v71
                        }
                    }
                } else {
                    v4[fi65] = fi65
                    var v83 int32 = -2147483648
                    _ = v83
                }
            case "a":
                v84 := []int{}
                _ = v84
            default:
                v1 -= (v1 % (((^(v1))) | 1))
            }
            var v85 int64 = 1
            v86 := GenS10{F0: 4294967295, F1: 42, F2: (^(v85))}
            _ = fi65
            _ = v85
            _ = v86
        }
        _ = len(v4)
        for fi87 := 0; fi87 < 13; fi87++ {
            v88 := GenS11{F0: 255, F1: 4294967295}
            _ = fi87
            _ = v88
        }
        _ = fi64
    }
    fmt.Println(v63, v1)
    var v89 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v63:", v63)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v89
}
