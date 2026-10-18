package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int; F2 uint16 }
func genFn0(p0 int16) int8 {
    var v4 float64 = 3.14159
    var v5 uint16 = 937
    fmt.Println(v5, v4, v5)
    fmt.Println(p0)
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 bool = ((fi6 | fi6) < 2147483647)
            _ = fi6
            _ = v7
        }
        if (42 > -772) {
            if ("a" >= "\n\t") {
                p0 += 156
            }
            v8 := GenS1{F0: (^(v5)), F1: -40}
            _ = v8
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}
func genFn1() uint64 {
    var v9 int16 = 42
    v10 := []int{int(v9)}
    if true {
        for fi11 := 0; fi11 < 12; fi11++ {
            v9 = v9
            _ = fi11
        }
        v12 := GenS2{F0: 127, F1: 335}
        _ = v12
    }
    switch -418 {
    case 42:
        for fi13 := 0; fi13 < 11; fi13++ {
            switch fi13 {
            case 1:
                if (!true) {
                    v9 = v9
                    v9 += 32767
                    if false {
                        for fi14 := 0; fi14 < 13; fi14++ {
                            switch fi13 {
                            case 8:
                                if len(v10) > 0 {
                                    v10[0] = fi14
                                }
                            case 128:
                                switch ("foo bar" + (("\n\t" + "a") + (("\n\t" + ("hello" + "\n\t")) + "hello"))) {
                                case "\n\t":
                                    if (fi14 < (fi13 / ((261) | 1))) {
                                        v9 = v9
                                        v15 := GenS1{F0: 32767, F1: 1}
                                        if false {
                                            fmt.Println(v9, fi14)
                                            v16 := GenS1{F0: 0, F1: -1}
                                            if true {
                                                var v17 float64 = (1e10 / 1e10)
                                                var v18 string = "\n\t"
                                                v17 = ((((((float64(((3.14159 - 1e10) / 1.0)) - 3.14159) - v17) / v17) * (v17 + v17)) + (0.0 - v17)) / (float64(1.5) / v17))
                                                v19 := GenS3{F0: 0}
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            }
                                            if false {
                                                fmt.Println(v9, fi13, fi13)
                                                var v20 int16 = -360
                                                v21 := GenS2{F0: int8(fi14), F1: 567}
                                                _ = v20
                                                _ = v21
                                            }
                                            fmt.Println(fi14, fi13, v9)
                                            _ = v16
                                        }
                                        v22 := GenS2{F0: 0, F1: 42}
                                        _ = v15
                                        _ = v22
                                    }
                                }
                            default:
                                fmt.Println(fi14)
                            }
                            _ = fi14
                        }
                    }
                }
            case -202:
                var v23 float32 = float32((0.1 + 0.1))
                _ = v23
            default:
                for fi24 := 0; fi24 < 14; fi24++ {
                    fmt.Println(v9, fi24)
                    var v25 uint32 = 2147483647
                    switch (fi24 * (-(fi24))) {
                    case 1:
                        v9 = (v9 & (^(v9)))
                    case -1:
                        v26 := GenS3{F0: v9}
                        v27 := GenS2{F0: 127, F1: 4294967295}
                        var v28 int32 = 1
                        v9 = v9
                        switch ((((("hello" + "\n\t") + "a") + (("foo bar" + "") + "a")) + "foo bar") + ("\n\t" + "")) {
                        case "\n\t":
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        default:
                            var v29 string = "hello"
                            _ = v29
                        }
                        _ = v26
                        _ = v27
                        _ = v28
                    case 127:
                        var v30 uint16 = 2
                        _ = v30
                    }
                    switch "" {
                    case "a":
                        fmt.Println(fi13, v9)
                    case "\n\t":
                        if ((65535 != v25) && true) {
                            v25 -= 4294967295
                            v25 -= (v25 ^ 4294967295)
                            v31 := GenS4{F0: genFn0(v9)}
                            _ = v31
                        }
                    case "foo bar":
                        switch (("\n\t" + "hello") + (("\n\t" + ("hello" + "\n\t")) + "foo bar")) {
                        case "hello":
                            fmt.Println(v9)
                        case "":
                            if len(v10) > 0 {
                                v10[0] = fi13
                            }
                        case "foo bar":
                            v25 -= (v25 % ((v25) | 1))
                            v32 := GenS3{F0: (v9 * 1000)}
                            _ = v32
                        }
                    }
                    var v33 uint16 = 121
                    switch "" {
                    case "hello":
                        v9 -= -1
                    }
                    _ = len(v10)
                    _ = fi24
                    _ = v25
                    _ = v33
                }
            }
            var v34 int8 = 1
            _ = fi13
            _ = v34
        }
    case -626:
        v35 := GenS4{F0: -24}
        _ = v35
    case -2147483648:
        fmt.Println(v9)
    default:
        v9 += (v9 * 773)
    }
    _ = v10
    _ = v9
    return 647
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", "\n\t", ("foo bar" + "foo bar"), ("foo bar" + "foo bar")}
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 int32 = 2147483647
        var v38 uint16 = 32767
        v39 := GenS3{F0: 1}
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    var v40 int = 32768
    v40 = v40
    for fi41 := 0; fi41 < 11; fi41++ {
        if false {
            v40 += (v40 - v40)
            v40 -= fi41
        }
        var v42 string = "foo bar"
        _ = fi41
        _ = v42
    }
    var v43 uint = 777
    v44 := GenS1{F0: 32767, F1: 42}
    for fi45 := 0; fi45 < 7; fi45++ {
        if (920 != 1) {
            v43 = v43
        } else {
            var v46 bool = true
            var v47 string = "a"
            for fi48 := 0; fi48 < 6; fi48++ {
                _ = len(v1)
                if (!(!(v46 || ((^(v43)) < v43)))) {
                    if ((526 != 763) || ((v43 < v43) || (-1 >= 42))) {
                        switch v40 {
                        case -377:
                            var v49 uint32 = 482
                            _ = v49
                        }
                        var v50 int32 = -632
                        fmt.Println(fi45, v50)
                        _ = v50
                    } else {
                        switch v47 {
                        case "a":
                            v3 = append(v3, "hello")
                        case "":
                            if (128 > 128) {
                                v40 += int(v43)
                                var v51 int8 = genFn0(1)
                                var v52 int = (v40 << ((v43) & 63))
                                _ = v51
                                _ = v52
                            }
                        case "\n\t":
                            var v53 int16 = -1
                            _ = v53
                        default:
                            switch 9223372036854775807 {
                            case -2147483648:
                                if (!false) {
                                    if (0 >= 128) {
                                        fmt.Println(fi48, v47)
                                        v1[2147483647] = 2
                                    }
                                    switch ("a" + "foo bar") {
                                    case "":
                                        var v54 string = ""
                                        _ = v54
                                    default:
                                        var v55 float64 = -2.25
                                        _ = v55
                                    }
                                }
                            case -1:
                                v47 = v47
                            case -441:
                                if ((v43 & v43) > (v43 + (^(v43)))) {
                                    var v56 int64 = 9223372036854775807
                                    var v57 float64 = 1e10
                                    _ = v56
                                    _ = v57
                                }
                                fmt.Println(fi45)
                            default:
                                if (0.1 == 1e10) {
                                    if v46 {
                                        var v58 int = 32768
                                        if (-2.25 < 0.0) {
                                            v59 := GenS5{F0: 42, F1: (fi45 | -1), F2: 42}
                                            _ = v59
                                        } else {
                                            var v60 uint64 = 0
                                            _ = v60
                                        }
                                        _ = v58
                                    }
                                }
                            }
                        }
                    }
                    v47 = "a"
                }
                _ = fi48
            }
            _ = v46
            _ = v47
        }
        for fi61 := 0; fi61 < 4; fi61++ {
            v43 = v43
            switch ("foo bar" + "a") {
            case "":
                var v62 int32 = -783
                fmt.Println(v62, v43)
                v63 := GenS3{F0: -1}
                v64 := GenS2{F0: genFn0(1), F1: v43}
                _ = v62
                _ = v63
                _ = v64
            case "foo bar":
                switch "" {
                case "":
                    fmt.Println(v43, v43, v43)
                case "hello":
                    v65 := GenS1{F0: 32767, F1: 100}
                    _ = v65
                }
            case "\n\t":
                v1[1] = 42
            }
            var v66 uint64 = 238
            switch ("\n\t" + ("a" + "\n\t")) {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                fmt.Println(fi45, v66)
            default:
                if ((!true) || false) {
                    var v67 uint = v43
                    _ = v67
                }
            }
            v43 += (v43 * v43)
            _ = fi61
            _ = v66
        }
        var v68 float32 = float32(1e10)
        _ = fi45
        _ = v68
    }
    v69 := v1[1]
    var v70 uint = v43
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v43
    _ = v44
    _ = v69
    _ = v70
}
