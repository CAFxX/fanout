package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 bool }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16; F1 uint16; F2 uint8 }
func genFn0(p0 int64) string {
    v6 := GenS1{F0: 863, F1: 127, F2: 0}
    var v7 int32 = 0
    var v8 uint8 = 51
    switch 0 {
    case -151:
        switch 0 {
        case 32767:
            v9 := GenS1{F0: 1, F1: 2, F2: 57}
            _ = v9
        case 468:
            fmt.Println(v8, p0, v7)
        default:
            if true {
                var v10 uint64 = 42
                for fi11 := 0; fi11 < 13; fi11++ {
                    v8 += v8
                    _ = fi11
                }
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS1{F0: 18446744073709551615, F1: v8, F2: (fi12 | fi12)}
                    for fi14 := 0; fi14 < 20; fi14++ {
                        if false {
                            var v15 uint8 = uint8(v10)
                            _ = v15
                        } else {
                            v16 := []int{(-(fi12)), fi14}
                            var v17 uint32 = 30
                            var v18 int = (fi12 % ((2147483647) | 1))
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        _ = fi14
                    }
                    v19 := GenS1{F0: v10, F1: v8, F2: (^(fi12))}
                    _ = fi12
                    _ = v13
                    _ = v19
                }
                _ = v10
            } else {
                if (3.14159 >= 0.1) {
                    switch 0 {
                    case 4294967295:
                        if false {
                            var v20 string = "a"
                            if (true && true) {
                                v21 := GenS1{F0: 18446744073709551615, F1: v8, F2: 0}
                                v22 := GenS2{F0: 754}
                                _ = v21
                                _ = v22
                            }
                            var v23 string = ("foo bar" + v20)
                            _ = v20
                            _ = v23
                        }
                    default:
                        for fi24 := 0; fi24 < 7; fi24++ {
                            var v25 uint32 = 14
                            fmt.Println(p0, v7)
                            _ = fi24
                            _ = v25
                        }
                    }
                    if false {
                        var v26 int8 = -127
                        _ = v26
                    } else {
                        v8 += (^(v8))
                        if true {
                            var v27 uint64 = 18446744073709551615
                            var v28 bool = false
                            _ = v27
                            _ = v28
                        }
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v8, v8, v7)
                        case "foo bar":
                            var v29 float32 = (1.5 * (1.5 * (((1e10 - 3.14159) / 1.0) + 1.5)))
                            _ = v29
                        default:
                            fmt.Println(p0)
                        }
                    }
                    v30 := []int{32767, 127, 1, 0}
                    _ = v30
                }
                switch "\n\t" {
                case "foo bar":
                    v31 := GenS1{F0: 18446744073709551615, F1: (v8 * v8), F2: 9223372036854775807}
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 7; fi32++ {
                        if true {
                            var v33 float64 = (3.14159 * 1e10)
                            if (!false) {
                                p0 += 548
                            }
                            _ = v33
                        } else {
                            if true {
                                v34 := []int{(fi32 ^ -700)}
                                _ = v34
                            }
                            v35 := []int{32767, 32768, fi32, (fi32 >> ((42) & 63))}
                            for fi36 := 0; fi36 < 6; fi36++ {
                                var v37 int = (fi32 % (((fi36 / (((fi32 + (fi36 / ((fi32) | 1)))) | 1))) | 1))
                                var v38 float64 = 0.1
                                _ = len(v35)
                                _ = len(v35)
                                v8 = (v8 / ((255) | 1))
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            _ = v35
                        }
                        switch "\n\t" {
                        case "a":
                            if true {
                                fmt.Println(v8, v8)
                                var v39 int32 = v7
                                _ = v39
                            }
                        }
                        _ = fi32
                    }
                }
            }
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return ((("" + "\n\t") + "") + (("\n\t" + "\n\t") + ("hello" + (("a" + "") + "foo bar"))))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -959
    var v3 int8 = 0
    v4 := []string{"foo bar", "hello", "hello", "foo bar"}
    v5 := []int{v2, (v2 % ((-419) | 1)), (v2 * (v2 - (v2 * 42))), v2}
    var v40 int8 = v3
    switch "" {
    case "":
        if false {
            var v41 float64 = (0.1 + (1.5 / 1.0))
            _ = v41
        } else {
            var v42 float32 = 3.14159
            _ = v42
        }
    default:
        var v43 int = (v2 << ((4294967295) & 63))
        _ = v43
    }
    if (v3 < (v40 / (((v40 % ((v40) | 1))) | 1))) {
        var v44 string = "foo bar"
        _ = v44
    } else {
        if false {
            v45 := GenS2{F0: -1}
            var v46 float64 = (float64((1.5 * (float32((-2.25 * 3.14159)) + 0.1))) / 1e10)
            v47 := GenS3{F0: 287, F1: 255, F2: ("hello" >= "\n\t")}
            _ = v45
            _ = v46
            _ = v47
        }
        if true {
            var v48 int = v2
            if ("a" == ("a" + "a")) {
                var v49 float32 = 1.5
                v1["\n\t"] = (!(0 < 555))
                switch v2 {
                case 32767:
                    var v50 int64 = 9223372036854775807
                    _ = v50
                default:
                    var v51 int64 = 182
                    _ = v51
                }
                _ = v49
            }
            _ = v48
        }
    }
    fmt.Println(v2)
    _ = len(v1)
    if (0.1 < float64(float32((-2.25 / float64(((((float32((3.14159 - 3.14159)) + -2.25) + -2.25) - 3.14159) * 1.5)))))) {
        var v52 float64 = 1.5
        var v53 float32 = 0.0
        if true {
            var v54 bool = true
            _ = v54
        } else {
            if true {
                _ = len(v1)
                var v55 int16 = -1
                _ = v55
            }
            v3 -= (v3 ^ 127)
        }
        _ = v52
        _ = v53
    } else {
        if (!true) {
            switch (("hello" + "\n\t") + "\n\t") {
            case "\n\t":
                v56 := GenS4{F0: ((!true) || (182 <= 71)), F1: 4294967295}
                _ = v56
            case "a":
                var v57 float32 = 1.5
                _ = v57
            case "foo bar":
                for fi58 := 0; fi58 < 20; fi58++ {
                    v3 = 42
                    _ = fi58
                }
            default:
                var v59 uint64 = 18446744073709551615
                _ = v59
            }
        } else {
            var v60 uint8 = 2
            _ = v60
        }
        for fi61 := 0; fi61 < 18; fi61++ {
            fmt.Println(fi61, v40, fi61)
            for fi62 := 0; fi62 < 17; fi62++ {
                v63 := GenS5{F0: 278, F1: fi61}
                var v64 float64 = (1.5 + (3.14159 - 0.0))
                _ = fi62
                _ = v63
                _ = v64
            }
            _ = fi61
        }
    }
    v65 := GenS6{F0: (!true)}
    switch ("\n\t" + "a") {
    case "":
        v66 := GenS7{F0: 2, F1: 916}
        _ = v66
    case "hello":
        v3 -= -128
    default:
        var v67 int = (^(v2))
        _ = v67
    }
    var v68 int32 = -1
    for fi69 := 0; fi69 < 15; fi69++ {
        v70 := GenS8{F0: (v68 - (v68 | v68)), F1: 884}
        var v71 int = fi69
        v1["a"] = (!(140 == -1))
        _ = fi69
        _ = v70
        _ = v71
    }
    if true {
        for fi72 := 0; fi72 < 1; fi72++ {
            v73 := GenS8{F0: v68, F1: 18446744073709551615}
            _ = fi72
            _ = v73
        }
        switch 9223372036854775807 {
        case 42:
            v2 -= (v2 ^ v2)
        case 278:
            v74 := GenS9{F0: (v40 >> ((0) & 63))}
            _ = v74
        case 32767:
            v2 = -9223372036854775808
        }
    } else {
        var v75 int64 = 42
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi76 := 0; fi76 < 16; fi76++ {
            if false {
                var v77 int16 = 1
                for fi78 := 0; fi78 < 5; fi78++ {
                    switch 127 {
                    case 9223372036854775807:
                        v77 = 42
                    case 0:
                        v68 += v68
                    case 32767:
                        _ = len(v5)
                    default:
                        var v79 uint = 4294967295
                        _ = v79
                    }
                    if true {
                        switch (fi78 | (fi78 << ((18446744073709551615) & 63))) {
                        case -9223372036854775808:
                            fmt.Println(v75, v77)
                        case 32768:
                            if (true && false) {
                                fmt.Println(v75)
                                switch fi78 {
                                case 9223372036854775807:
                                    if len(v4) > 0 {
                                        v4[0] = ((("foo bar" + "") + ("" + "\n\t")) + "\n\t")
                                    }
                                case -65:
                                    if false {
                                        v75 = -546
                                    } else {
                                        var v80 uint8 = 127
                                        if false {
                                            v81 := GenS5{F0: 930, F1: (fi78 + fi78)}
                                            v2 = (v2 >> ((18446744073709551615) & 63))
                                            _ = v81
                                        }
                                        _ = v80
                                    }
                                case 128:
                                    v82 := GenS10{F0: 595, F1: 518, F2: 0}
                                    _ = v82
                                default:
                                    v68 += int32(v2)
                                }
                            }
                        case 32767:
                            var v83 uint64 = 18446744073709551615
                            _ = v83
                        default:
                            delete(v1, ((("a" + "\n\t") + "foo bar") + (("\n\t" + (("foo bar" + "a") + ("hello" + "foo bar"))) + "\n\t")))
                        }
                        var v84 uint16 = 492
                        _ = v84
                    }
                    if false {
                        _ = len(v1)
                        var v85 float64 = -2.25
                        _ = v85
                    } else {
                        v68 -= v68
                    }
                    _ = fi78
                }
                fmt.Println(fi76)
                _ = v77
            } else {
                var v86 uint32 = 76
                var v87 uint16 = 0
                _ = v86
                _ = v87
            }
            _ = fi76
        }
        _ = v75
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v65
    _ = v68
}
