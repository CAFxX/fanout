package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
func genFn0() uint32 {
    if false {
        if false {
            v5 := []int{42, 762, -109}
            var v6 uint = 552
            for fi7 := 0; fi7 < 20; fi7++ {
                v6 -= 561
                for fi8 := 0; fi8 < 19; fi8++ {
                    v6 = v6
                    _ = fi8
                }
                _ = fi7
            }
            _ = v5
            _ = v6
        }
        if true {
            v9 := GenS1{F0: 293}
            var v10 uint = 237
            _ = v10
            _ = v9
        } else {
            var v11 uint8 = 215
            var v12 uint8 = (v11 + v11)
            var v13 int = -849
            _ = v11
            _ = v12
            _ = v13
        }
    }
    var v14 string = "hello"
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int32 = 390
        _ = fi15
        _ = v16
    }
    _ = v14
    return 2147483648
}
func genFn1() uint32 {
    v17 := GenS2{F0: 0, F1: -9223372036854775808}
    v18 := []int{-9223372036854775808, -9223372036854775808, 4294967295}
    if true {
        for fi19 := 0; fi19 < 2; fi19++ {
            if len(v18) > 0 {
                v18[0] = (fi19 ^ fi19)
            }
            switch "\n\t" {
            case "a":
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 uint32 = 4294967295
                    fmt.Println(v21, v21)
                    var v22 int8 = int8(fi19)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            case "hello":
                var v23 uint64 = 194
                _ = v23
            }
            var v24 uint16 = 32767
            _ = fi19
            _ = v24
        }
    }
    var v25 uint = 460
    _ = v17
    _ = v18
    _ = v25
    return 243
}

func main() {
    v1 := []string{""}
    var v2 bool = (1.5 < (float32((1.5 - 0.1)) - 3.14159))
    v3 := make(map[string]bool)
    var v4 int = 904
    var v26 float32 = 1.5
    v27 := GenS1{F0: 18446744073709551615}
    switch 122 {
    case 127:
        var v28 int16 = 461
        _ = v28
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            if ((v2 || (!v2)) && v2) {
                for fi30 := 0; fi30 < 15; fi30++ {
                    fmt.Println(v26, fi29, fi30)
                    _ = fi30
                }
                for fi31 := 0; fi31 < 16; fi31++ {
                    if v2 {
                        if v2 {
                            var v32 float64 = 3.14159
                            _ = len(v1)
                            switch "\n\t" {
                            case "a":
                                v26 -= v26
                            case "foo bar":
                                var v33 uint32 = 1
                                _ = v33
                            }
                            _ = v32
                        } else {
                            var v34 int64 = -1
                            switch (v4 * (fi29 - fi29)) {
                            case 781:
                                v35 := GenS2{F0: 42, F1: v34}
                                _ = v35
                            case 42:
                                v34 = v34
                            }
                            v26 += (3.14159 - v26)
                            _ = v34
                        }
                        v2 = (206 != 509)
                    } else {
                        var v36 bool = v2
                        _ = v36
                    }
                    v4 -= (fi29 % ((v4) | 1))
                    v37 := GenS2{F0: 0, F1: 42}
                    _ = fi31
                    _ = v37
                }
                v38 := GenS2{F0: 2, F1: -1}
                _ = v38
            }
            for fi39 := 0; fi39 < 11; fi39++ {
                v2 = ((v4 - (fi29 >> ((1) & 63))) == -44)
                switch -1 {
                case 42:
                    var v40 uint8 = 76
                    _ = v40
                default:
                    var v41 int16 = 924
                    _ = v41
                }
                var v42 string = "foo bar"
                _ = fi39
                _ = v42
            }
            _ = fi29
        }
    }
    fmt.Println(v2, v26)
    var v43 uint16 = 42
    v26 = v26
    switch v4 {
    case 73:
        var v44 uint8 = 42
        _ = v44
    case 0:
        var v45 int32 = -671
        _ = v45
    case 905:
        if true {
            switch "" {
            case "foo bar":
                v46 := GenS3{F0: (("" + ("foo bar" + "\n\t")) + "")}
                _ = v46
            }
            if true {
                var v47 float32 = v26
                if false {
                    switch (v4 & (-(v4))) {
                    case -1:
                        v26 += v26
                    case 32767:
                        fmt.Println(v47, v43)
                    case 128:
                        v48 := GenS3{F0: "a"}
                        _ = v48
                    }
                }
                _ = v47
            }
            var v49 int16 = 42
            _ = v49
        } else {
            v4 += v4
            v50 := GenS3{F0: ("hello" + (("" + "hello") + "a"))}
            _ = v50
        }
    default:
        var v51 bool = v2
        _ = v51
    }
    var v52 int32 = 2147483647
    for fi53 := 0; fi53 < 5; fi53++ {
        v52 = v52
        _ = len(v3)
        _ = fi53
    }
    if (("hello" + "\n\t") > "") {
        v52 += v52
    } else {
        var v54 float64 = 1e10
        _ = v54
    }
    v55 := GenS4{F0: 57}
    v43 -= v43
    var v56 bool = v2
    _ = len(v1)
    var v57 float32 = -2.25
    v58 := GenS4{F0: 0}
    if (885 < 932) {
        if v56 {
            for fi59 := 0; fi59 < 18; fi59++ {
                v60 := GenS1{F0: 0}
                _ = fi59
                _ = v60
            }
        }
        if ((-(v4)) < v4) {
            switch v4 {
            case 42:
                v26 = float32(1e10)
            }
            for fi61 := 0; fi61 < 6; fi61++ {
                for fi62 := 0; fi62 < 7; fi62++ {
                    if (255 != 0) {
                        if true {
                            switch 568 {
                            case -372:
                                v26 = v57
                            case 51:
                                v57 -= v57
                            }
                        }
                    }
                    switch "foo bar" {
                    case "":
                        var v63 int = 127
                        _ = v63
                    }
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = fi62
                }
                v2 = ((^(v52)) != v52)
                var v64 int32 = (v52 - v52)
                _ = fi61
                _ = v64
            }
            for fi65 := 0; fi65 < 7; fi65++ {
                v26 = v57
                v2 = v2
                _ = fi65
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v52
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
