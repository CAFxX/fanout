package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0(p0 string, p1 uint, p2 string) uint {
    fmt.Println(p1, p2, p0)
    p0 = (((p0 + (p2 + "")) + p2) + p2)
    p0 = ""
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 int) bool {
    var v6 int16 = 1000
    p0 = (-(p0))
    v7 := GenS1{F0: 0, F1: 42, F2: 2}
    var v8 int8 = -80
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 float32 = 1.5
    var v2 int16 = -32768
    var v3 int16 = 834
    var v4 int16 = v2
    v5 := make(map[int]bool)
    if (!(-784 > 42)) {
        v9 := GenS1{F0: 2, F1: 42, F2: 0}
        v4 += (v2 << ((1) & 63))
        v1 = v1
        _ = v9
    }
    var v10 bool = true
    if ((135 == 2) || (((((((((v1 / (v1 * 3.14159)) / v1) + 3.14159) / -2.25) * 0.1) + ((((v1 + v1) + v1) * 0.1) + v1)) * v1) + v1) > (v1 / (3.14159 / v1)))) {
        if true {
            v11 := GenS1{F0: 556, F1: -1, F2: 42}
            v12 := []int{440, 815}
            _ = v11
            _ = v12
        } else {
            var v13 string = "hello"
            v4 = 1
            v14 := GenS1{F0: 251, F1: 0, F2: 4294967295}
            _ = v13
            _ = v14
        }
        switch "a" {
        case "":
            v15 := []int{42, 0}
            _ = v15
        default:
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 int64 = -9223372036854775808
                v10 = (v2 == v4)
                _ = fi16
                _ = v17
            }
        }
    } else {
        fmt.Println(v3)
        fmt.Println(v3, v4)
        v18 := GenS1{F0: 364, F1: 1, F2: genFn0("\n\t", 65535, "foo bar")}
        _ = v18
    }
    switch 4294967295 {
    case 173:
        switch 1 {
        case 9223372036854775807:
            if (64 > -98) {
                v19 := GenS1{F0: 730, F1: 42, F2: 255}
                _ = len(v5)
                _ = v19
            } else {
                var v20 float64 = ((3.14159 + ((1e10 * ((((float64(v1) / (0.1 * ((0.0 + float64(1e10)) / (1e10 / 1.5)))) / 0.1) + ((((float64(v1) - (-2.25 + (0.0 + 0.1))) * 0.0) + ((0.1 - 1e10) - 1e10)) + 1e10)) + (1e10 * -2.25))) / 1e10)) / 1.0)
                v21 := GenS1{F0: 4294967295, F1: 2147483647, F2: genFn0("foo bar", 0, "\n\t")}
                var v22 uint32 = 65535
                _ = v20
                _ = v21
                _ = v22
            }
        case 1:
            delete(v5, 980)
        case -1:
            switch "foo bar" {
            case "a":
                var v23 float32 = float32((float64((v1 - v1)) + 3.14159))
                _ = v23
            case "":
                if ((v1 * (v1 * (float32(1.5) * v1))) < (0.1 * v1)) {
                    if (0 != -1) {
                        var v24 int = 127
                        var v25 int16 = v4
                        _ = v24
                        _ = v25
                    } else {
                        delete(v5, 128)
                    }
                    if ((213 == 42) || false) {
                        if true {
                            for fi26 := 0; fi26 < 17; fi26++ {
                                var v27 uint64 = 830
                                v28 := v5[-81]
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        } else {
                            v29 := GenS2{F0: 255}
                            v4 += (v2 % ((v4) | 1))
                            v30 := []int{-1, 46}
                            _ = v29
                            _ = v30
                        }
                    }
                    var v31 float32 = 0.1
                    _ = v31
                }
            }
        }
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        if false {
            var v33 uint8 = 128
            var v34 int32 = -2147483648
            _ = v33
            _ = v34
        } else {
            v35 := GenS3{F0: 25, F1: ("\n\t" + "\n\t"), F2: 685}
            var v36 uint8 = 128
            _ = v35
            _ = v36
        }
        if (0 != 128) {
            var v37 int = fi32
            delete(v5, 2147483647)
            _ = v37
        } else {
            var v38 uint16 = 598
            v38 += v38
            _ = v38
        }
        _ = fi32
    }
    if true {
        v39 := GenS2{F0: 0}
        fmt.Println(v2, v10)
        var v40 bool = v10
        _ = v39
        _ = v40
    }
    v41 := []int{}
    v4 = v4
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    v42 := GenS3{F0: 0, F1: "\n\t", F2: 2}
    for fi43 := 0; fi43 < 2; fi43++ {
        v3 = v4
        for fi44 := 0; fi44 < 4; fi44++ {
            var v45 uint8 = 42
            _ = len(v41)
            _ = fi44
            _ = v45
        }
        switch "" {
        case "hello":
            v3 -= (v3 + v3)
        case "":
            if (2 < 2) {
                fmt.Println(v1)
                v46 := GenS2{F0: 42}
                _ = v46
            } else {
                if false {
                    v3 = 957
                    v2 = v3
                }
                if (0.0 < (0.1 + float64(1e10))) {
                    fmt.Println(v10, v2, v1)
                    switch 141 {
                    case -9223372036854775808:
                        var v47 int64 = 1
                        _ = v47
                    case 404:
                        for fi48 := 0; fi48 < 19; fi48++ {
                            switch "hello" {
                            case "\n\t":
                                var v49 float32 = (((v1 / 3.14159) - float32(float64(((v1 - v1) + (v1 + float32(3.14159)))))) / v1)
                                _ = v49
                            case "":
                                var v50 int64 = 1
                                _ = v50
                            case "a":
                                var v51 uint64 = 946
                                _ = v51
                            default:
                                if v10 {
                                    if ((("foo bar" + "\n\t") + "\n\t") <= ("hello" + "\n\t")) {
                                        var v52 uint8 = 1
                                        var v53 int16 = v2
                                        var v54 uint = genFn0(("\n\t" + "a"), 65535, (("a" + "a") + (("a" + "foo bar") + "a")))
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    if len(v41) > 0 {
                                        v41[0] = (fi43 >> ((237) & 63))
                                    }
                                } else {
                                    switch "" {
                                    case "":
                                        _ = len(v41)
                                        switch ("foo bar" + "\n\t") {
                                        case "hello":
                                            switch "\n\t" {
                                            case "hello":
                                                fmt.Println(fi43, v10)
                                                v1 = 0.0
                                            case "foo bar":
                                                v2 += 206
                                            case "a":
                                                v55 := GenS2{F0: 2}
                                                _ = v55
                                            }
                                        case "a":
                                            v56 := GenS1{F0: 868, F1: 42, F2: 992}
                                            _ = v56
                                        default:
                                            v10 = (1 != 0)
                                        }
                                    case "\n\t":
                                        v10 = ((v3 - v4) == v3)
                                    default:
                                        v57 := GenS1{F0: 0, F1: -972, F2: 984}
                                        _ = v57
                                    }
                                    v3 = (v4 ^ 0)
                                }
                            }
                            v3 = 1
                            v2 = v4
                            _ = fi48
                        }
                    }
                    fmt.Println(v3)
                } else {
                    var v58 uint = genFn0("hello", 647, "foo bar")
                    _ = v58
                }
                if v10 {
                    var v59 uint64 = 42
                    v60 := GenS4{F0: -128, F1: -347}
                    switch "foo bar" {
                    case "foo bar":
                        var v61 int32 = -2147483648
                        _ = v61
                    case "a":
                        if (446 > 349) {
                            v62 := GenS3{F0: 543, F1: "hello", F2: v59}
                            _ = v62
                        } else {
                            var v63 uint = 65535
                            _ = v63
                        }
                    default:
                        var v64 uint16 = 973
                        _ = v64
                    }
                    _ = v59
                    _ = v60
                } else {
                    if (1e10 == 1e10) {
                        v65 := v5[648]
                        _ = v65
                    }
                    v10 = v10
                }
            }
        case "foo bar":
            v66 := GenS2{F0: 42}
            _ = v66
        }
        _ = fi43
    }
    var v67 int32 = -318
    if v10 {
        v2 += v4
        for fi68 := 0; fi68 < 6; fi68++ {
            switch "foo bar" {
            case "":
                var v69 float64 = (((float64(float32(3.14159)) * 1.5) - (1.5 - 1e10)) - 0.1)
                _ = v69
            case "\n\t":
                fmt.Println(fi68, v67)
            case "a":
                v70 := GenS3{F0: 718, F1: "", F2: 541}
                _ = v70
            default:
                var v71 float64 = -2.25
                _ = v71
            }
            _ = len(v41)
            v72 := v5[fi68]
            _ = fi68
            _ = v72
        }
    }
    fmt.Println(v1, v10, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v67
}
