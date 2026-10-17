package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 int16; F1 int; F2 int64 }
func genFn0(p0 int32, p1 int16) uint32 {
    v5 := []int{}
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = 3.14159
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 int64) int {
    fmt.Println(p1, p0)
    p1 += p1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint16 = 2
        _ = fi8
        _ = v9
    }
    if (!false) {
        p0 -= (p0 * p0)
        var v10 string = "a"
        _ = v10
    } else {
        switch "\n\t" {
        case "a":
            if true {
                if (!false) {
                    var v11 uint8 = p0
                    _ = v11
                }
                for fi12 := 0; fi12 < 19; fi12++ {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(fi12, p0, p1)
                    default:
                        var v13 bool = false
                        _ = v13
                    }
                    _ = fi12
                }
                v14 := GenS1{F0: -4}
                _ = v14
            } else {
                v15 := []int{42, -9223372036854775808, 1}
                for fi16 := 0; fi16 < 20; fi16++ {
                    p0 = p0
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v18 := GenS1{F0: 1}
                        var v19 int = 42
                        var v20 bool = false
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = fi16
                }
                if len(v15) > 0 {
                    v15[0] = 1
                }
                _ = v15
            }
        default:
            v21 := GenS1{F0: 42}
            _ = v21
        }
        var v22 float32 = 0.0
        var v23 int64 = (p1 >> ((0) & 63))
        _ = v22
        _ = v23
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = -2147483648
    var v3 int8 = -128
    v4 := make(map[int]bool)
    if true {
        if true {
            fmt.Println(v2, v2, v2)
            var v24 uint64 = (^(v1))
            fmt.Println(v1, v2, v24)
            _ = v24
        } else {
            v3 = 127
            var v25 float32 = 1e10
            _ = v25
        }
        var v26 bool = ((-2.25 - 0.1) <= 0.0)
        _ = v26
    } else {
        var v27 float32 = 1.5
        for fi28 := 0; fi28 < 16; fi28++ {
            v27 = v27
            _ = fi28
        }
        _ = v27
    }
    var v29 bool = true
    v29 = false
    if v29 {
        fmt.Println(v29, v29)
        if false {
            v30 := []int{-2147483648}
            for fi31 := 0; fi31 < 12; fi31++ {
                v29 = ((v3 <= (v3 >> ((65535) & 63))) || v29)
                v32 := GenS1{F0: (v2 ^ 165)}
                _ = fi31
                _ = v32
            }
            _ = v30
        }
        delete(v4, 42)
    }
    if ("foo bar" == ("a" + "")) {
        v33 := GenS2{F0: v1, F1: 0, F2: 1}
        _ = v33
    } else {
        if (v29 || v29) {
            switch "a" {
            case "hello":
                for fi34 := 0; fi34 < 18; fi34++ {
                    v3 = (v3 >> ((0) & 63))
                    _ = fi34
                }
            case "a":
                v1 = v1
            }
            v2 += v2
            v35 := GenS2{F0: (v1 << ((708) & 63)), F1: 496, F2: 9223372036854775807}
            _ = v35
        } else {
            for fi36 := 0; fi36 < 5; fi36++ {
                if (!false) {
                    v37 := []int{genFn1(2, -1), (fi36 - (fi36 << ((18446744073709551615) & 63))), 950, genFn1(128, 9223372036854775807)}
                    _ = v37
                }
                v38 := GenS2{F0: (^(v1)), F1: genFn0((v2 + v2), 0), F2: 0}
                if (true || (4294967295 != 255)) {
                    if ((42 != (v2 - (v2 >> ((39) & 63)))) && (false && (1 <= (v1 << ((255) & 63))))) {
                        switch "hello" {
                        case "a":
                            if (2 <= 0) {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    v39 := []int{}
                                    _ = v39
                                }
                                var v40 float64 = 1.5
                                v1 -= (^(v1))
                                _ = v40
                            }
                        case "hello":
                            v2 += -1
                        case "\n\t":
                            switch (fi36 % ((fi36) | 1)) {
                            case 127:
                                v29 = v29
                            case 32767:
                                if true {
                                    v1 += v1
                                }
                            }
                        default:
                            for fi41 := 0; fi41 < 15; fi41++ {
                                v42 := []int{genFn1(0, 162), fi41, (-(fi36)), (fi41 - fi41)}
                                if v29 {
                                    var v43 bool = ((3.14159 - 0.0) != float64(0.1))
                                    switch genFn1(128, -324) {
                                    case 32768:
                                        v4[fi36] = (253 >= 1)
                                    default:
                                        if (1e10 < (1.5 - 0.1)) {
                                            v44 := v4[(fi36 + fi36)]
                                            _ = v44
                                        }
                                    }
                                    _ = v43
                                } else {
                                    if (float64(3.14159) >= (1.5 + 0.0)) {
                                        var v45 string = ("\n\t" + "a")
                                        v45 = (v45 + v45)
                                        fmt.Println(v29, v3, fi36)
                                        _ = v45
                                    } else {
                                        if v29 {
                                            if (v29 && false) {
                                                var v46 string = "\n\t"
                                                fmt.Println(v29, v1)
                                                _ = v46
                                            }
                                            var v47 int8 = 127
                                            _ = v47
                                        }
                                        _ = len(v4)
                                    }
                                    fmt.Println(fi41, fi41)
                                }
                                var v48 int64 = 1
                                _ = fi41
                                _ = v42
                                _ = v48
                            }
                        }
                    } else {
                        var v49 int8 = 1
                        for fi50 := 0; fi50 < 9; fi50++ {
                            if (false || false) {
                                var v51 uint32 = 0
                                delete(v4, fi50)
                                if v29 {
                                    v51 = genFn0(v2, -295)
                                } else {
                                    fmt.Println(v3, fi36, fi36)
                                }
                                var v52 uint32 = 1
                                v53 := []int{}
                                v54 := GenS1{F0: v2}
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                            v55 := []int{2147483647, fi50, fi50, 0}
                            _ = len(v4)
                            _ = fi50
                            _ = v55
                        }
                        _ = v49
                    }
                    switch fi36 {
                    case -336:
                        v56 := []int{(fi36 - (fi36 / (((fi36 | (fi36 + -2147483648))) | 1))), fi36, fi36, genFn1(255, 297)}
                        _ = v56
                    case 7:
                        switch fi36 {
                        case -1:
                            v2 += v2
                        case 1:
                            var v57 string = "\n\t"
                            _ = v57
                        case 4294967295:
                            v58 := []int{-1}
                            _ = v58
                        default:
                            v2 -= v2
                        }
                    case -9223372036854775808:
                        v2 = v2
                    default:
                        var v59 int64 = -1
                        _ = v59
                    }
                }
                _ = fi36
                _ = v38
            }
            v60 := v4[-9223372036854775808]
            _ = v60
        }
        switch 0 {
        case 32767:
            v61 := GenS3{F0: 0, F1: 923, F2: 42}
            _ = v61
        case 1:
            var v62 int64 = 1
            _ = v62
        }
        v2 = (^(v2))
    }
    v1 = (v1 * (v1 * v1))
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    }
    v1 = v1
    var v63 uint32 = 338
    var v64 float32 = -2.25
    v65 := GenS1{F0: v2}
    for fi66 := 0; fi66 < 13; fi66++ {
        v2 += v2
        for fi67 := 0; fi67 < 7; fi67++ {
            switch "" {
            case "":
                var v68 int16 = 0
                _ = v68
            case "foo bar":
                v64 -= v64
            case "a":
                var v69 float64 = 3.14159
                _ = v69
            }
            switch (("foo bar" + (("hello" + "\n\t") + ("foo bar" + "a"))) + "hello") {
            case "a":
                v70 := []int{genFn1(255, 1), (fi66 * fi66), (-(fi66)), (-(fi67))}
                _ = v70
            default:
                var v71 int64 = int64(fi66)
                _ = v71
            }
            v72 := v4[(-(fi67))]
            v63 = (^(v63))
            _ = fi67
            _ = v72
        }
        v73 := []int{int(v1)}
        _ = fi66
        _ = v73
    }
    var v74 int16 = -343
    v75 := []int{genFn1(255, -1), 1, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v74:", v74)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v65
    _ = v74
    _ = v75
}
