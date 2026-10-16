package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint; F1 uint16; F2 uint }
func genFn0(p0 string, p1 uint, p2 bool) uint32 {
    if p2 {
        var v5 bool = p2
        fmt.Println(v5, p1, p2)
        _ = v5
    }
    var v6 int32 = 0
    switch p0 {
    case "foo bar":
        switch -9223372036854775808 {
        case 112:
            p1 += 4294967295
        case -953:
            if false {
                fmt.Println(p2)
                switch ("foo bar" + p0) {
                case "foo bar":
                    v6 = (-(v6))
                case "\n\t":
                    if (0 <= 2) {
                        fmt.Println(p1)
                        var v7 uint32 = 2147483648
                        p2 = p2
                        _ = v7
                    } else {
                        switch 762 {
                        case 1:
                            switch p0 {
                            case "":
                                var v8 string = "a"
                                _ = v8
                            default:
                                for fi9 := 0; fi9 < 5; fi9++ {
                                    switch (^(fi9)) {
                                    case 32767:
                                        v10 := []int{(fi9 + fi9), 42, -9223372036854775808}
                                        _ = v10
                                    }
                                    _ = fi9
                                }
                            }
                        case 854:
                            var v11 bool = (!true)
                            _ = v11
                        default:
                            p0 = ""
                        }
                    }
                }
                if ((!true) || true) {
                    v12 := []int{1, 2147483647, -204, 32767}
                    _ = v12
                }
            }
        case -523:
            if (-1 < -2147483648) {
                switch "hello" {
                case "a":
                    v13 := []int{}
                    _ = v13
                }
            }
        default:
            fmt.Println(p0)
        }
    case "\n\t":
        v6 += 0
    case "a":
        var v14 string = (("\n\t" + "hello") + p0)
        _ = v14
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}
func genFn1(p0 int64) uint16 {
    if false {
        v15 := []int{55, -9223372036854775808}
        p0 -= (p0 / (((^(p0))) | 1))
        var v16 int16 = -32768
        _ = v15
        _ = v16
    } else {
        var v17 uint32 = 2147483648
        _ = v17
    }
    var v18 int64 = p0
    p0 += (p0 - v18)
    switch 220 {
    case 2147483647:
        v19 := GenS1{F0: -412}
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := []int{}
        for fi22 := 0; fi22 < 6; fi22++ {
            fmt.Println(v18)
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v18
    return 1
}

func main() {
    v1 := []string{"hello", ("\n\t" + (((("hello" + "a") + (("a" + "foo bar") + "foo bar")) + "\n\t") + "\n\t")), ("hello" + "foo bar"), "\n\t"}
    var v2 int = -2147483648
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v23 int = (v2 - v2)
    v23 = 4294967295
    if false {
        v24 := GenS1{F0: v23}
        var v25 int64 = -1
        _ = v24
        _ = v25
    }
    switch "" {
    case "a":
        switch "hello" {
        case "a":
            if true {
                if true {
                    if (-779 <= 9223372036854775807) {
                        for fi26 := 0; fi26 < 12; fi26++ {
                            fmt.Println(fi26)
                            fmt.Println(v2, fi26)
                            v2 += fi26
                            _ = fi26
                        }
                        if (true && (int8(v2) < -128)) {
                            switch "\n\t" {
                            case "\n\t":
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    if (!(994 != genFn1(186))) {
                                        v28 := GenS1{F0: 457}
                                        if false {
                                            if false {
                                                if (807 != 2147483648) {
                                                    if true {
                                                        v29 := GenS2{F0: -1, F1: 28, F2: 42}
                                                        _ = v29
                                                    } else {
                                                        v1 = append(v1, "a")
                                                    }
                                                    fmt.Println(fi27)
                                                } else {
                                                    var v30 uint8 = 104
                                                    fmt.Println(v30, v23)
                                                    _ = v30
                                                }
                                            }
                                        }
                                        _ = v28
                                    } else {
                                        v1 = append(v1, ("foo bar" + ""))
                                        v2 += fi27
                                    }
                                    _ = fi27
                                }
                            case "a":
                                switch v23 {
                                case 673:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -46:
                                    if (!(-848 < 9223372036854775807)) {
                                        v2 -= v2
                                        v23 -= (v23 * v23)
                                        if (830 != -874) {
                                            v31 := GenS2{F0: 100, F1: 2147483648, F2: 32767}
                                            _ = v31
                                        }
                                    } else {
                                        var v32 uint16 = 0
                                        _ = v32
                                    }
                                case -1:
                                    v33 := GenS3{F0: 1000, F1: 65535}
                                    _ = v33
                                }
                            case "foo bar":
                                var v34 float64 = 1.5
                                _ = v34
                            default:
                                v35 := GenS2{F0: 1, F1: genFn0("foo bar", 2, (!(true && false))), F2: -372}
                                _ = v35
                            }
                            switch "hello" {
                            case "":
                                var v36 int32 = -2147483648
                                _ = v36
                            case "a":
                                switch (v23 | v23) {
                                case 2147483647:
                                    fmt.Println(v2)
                                default:
                                    v37 := GenS3{F0: 32767, F1: 0}
                                    _ = v37
                                }
                            case "foo bar":
                                var v38 int = v2
                                _ = v38
                            default:
                                v39 := GenS1{F0: v23}
                                _ = v39
                            }
                        } else {
                            v40 := GenS4{F0: 42, F1: 1, F2: 250}
                            _ = v40
                        }
                    } else {
                        for fi41 := 0; fi41 < 12; fi41++ {
                            switch v23 {
                            case 67:
                                v23 = (v23 / ((82) | 1))
                            default:
                                switch v2 {
                                case -320:
                                    var v42 float64 = ((0.1 * (0.1 * ((1.5 - ((-2.25 + (0.1 - (0.1 - 3.14159))) / (3.14159 * -2.25))) - 1.5))) / (0.0 + 1e10))
                                    _ = v42
                                case 25:
                                    if (v2 >= (v23 & (^(v2)))) {
                                        var v43 float64 = (1.5 - (-2.25 + 1.5))
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        var v44 float64 = 3.14159
                                        _ = v43
                                        _ = v44
                                    } else {
                                        for fi45 := 0; fi45 < 8; fi45++ {
                                            var v46 int64 = -1
                                            var v47 uint16 = 255
                                            fmt.Println(v47, v46, v47)
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                }
                            }
                            for fi48 := 0; fi48 < 6; fi48++ {
                                if true {
                                    v23 -= 32767
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        v2 += (fi48 * v2)
                                    case "foo bar":
                                        if true {
                                            v23 -= (v23 % ((fi48) | 1))
                                            delete(v3, fi48)
                                            v49 := GenS1{F0: fi41}
                                            _ = v49
                                        } else {
                                            if false {
                                                v2 = fi48
                                                v50 := GenS1{F0: (v23 % (((v2 << ((42) & 63))) | 1))}
                                                _ = v50
                                            } else {
                                                if true {
                                                    var v51 uint32 = 2147483648
                                                    _ = v51
                                                }
                                                var v52 uint8 = 248
                                                v23 = 4294967295
                                                _ = v52
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                v1 = append(v1, ("" + "\n\t"))
                                            default:
                                                var v53 uint8 = 128
                                                _ = v53
                                            }
                                            v23 = (v23 / ((v2) | 1))
                                        }
                                    default:
                                        var v54 uint32 = 65535
                                        _ = v54
                                    }
                                }
                                _ = fi48
                            }
                            _ = fi41
                        }
                        _ = len(v1)
                        var v55 int64 = 0
                        _ = v55
                    }
                    v2 += -2147483648
                }
            } else {
                for fi56 := 0; fi56 < 15; fi56++ {
                    var v57 float32 = (float32(0.1) / 1.0)
                    _ = fi56
                    _ = v57
                }
                switch v2 {
                case 0:
                    v2 -= -283
                }
                var v58 int = (^(v23))
                _ = v58
            }
        case "\n\t":
            fmt.Println(v2, v23)
        }
    default:
        var v59 int = (-(v2))
        _ = v59
    }
    fmt.Println(v2, v23)
    if true {
        v60 := GenS1{F0: v23}
        _ = v60
    }
    for fi61 := 0; fi61 < 8; fi61++ {
        var v62 uint = 65535
        _ = fi61
        _ = v62
    }
    v23 = v23
    if false {
        switch "hello" {
        case "foo bar":
            var v63 float32 = -2.25
            _ = v63
        case "\n\t":
            if false {
                v2 = -782
            }
        }
    } else {
        var v64 int16 = 32767
        _ = v64
    }
    var v65 string = "a"
    var v66 uint8 = 245
    if (true || (true || (1.5 <= (0.1 + 1e10)))) {
        _ = len(v1)
    }
    for fi67 := 0; fi67 < 13; fi67++ {
        if true {
            v68 := GenS2{F0: 100, F1: 4294967295, F2: 140}
            _ = v68
        } else {
            var v69 int16 = -107
            var v70 float32 = 0.0
            _ = v69
            _ = v70
        }
        _ = fi67
    }
    v66 -= (v66 - (v66 % (((v66 + 0)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v65
    _ = v66
}
