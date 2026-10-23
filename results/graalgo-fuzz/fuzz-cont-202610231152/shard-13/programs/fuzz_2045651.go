package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int; F1 string }
func genFn0(p0 uint) int {
    p0 = p0
    fmt.Println(p0)
    var v4 uint = 0
    v4 = 309
    if (100 < 127) {
        if (false || false) {
            var v5 bool = true
            v6 := GenS1{F0: 1, F1: -2147483648}
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := []int{-1, (fi7 | (-(fi7)))}
                if true {
                    v5 = v5
                    v9 := GenS1{F0: 65535, F1: fi7}
                    _ = v9
                } else {
                    switch (("a" + "foo bar") + "") {
                    case "":
                        var v10 uint8 = 1
                        _ = v10
                    }
                }
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        }
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]string)
    v3 := []string{"foo bar", "foo bar", "a"}
    switch "hello" {
    case "foo bar":
        if ((!false) && true) {
            if false {
                if (false && (float32((1.5 * float64(0.1))) >= ((-2.25 - -2.25) * 0.1))) {
                    v11 := GenS2{F0: (("" + "\n\t") + ("a" + "hello")), F1: 42, F2: ""}
                    for fi12 := 0; fi12 < 10; fi12++ {
                        fmt.Println(fi12, v1)
                        _ = fi12
                    }
                    _ = v11
                }
                switch (v1 % ((-597) | 1)) {
                case -540:
                    switch (^(v1)) {
                    case 128:
                        var v13 uint16 = 1
                        _ = v13
                    case -1:
                        switch v1 {
                        case -241:
                            for fi14 := 0; fi14 < 3; fi14++ {
                                if true {
                                    var v15 uint = 2
                                    _ = v15
                                }
                                var v16 uint = 352
                                switch v1 {
                                case 4294967295:
                                    var v17 int64 = 9223372036854775807
                                    _ = v17
                                case -1:
                                    var v18 int8 = 13
                                    _ = v18
                                case 2147483647:
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                default:
                                    fmt.Println(v16, v1)
                                }
                                _ = fi14
                                _ = v16
                            }
                        }
                    case -9223372036854775808:
                        var v19 bool = (1 <= 0)
                        _ = v19
                    default:
                        if (0 != 116) {
                            var v20 float32 = (1e10 / 1e10)
                            var v21 uint64 = 1
                            _ = v20
                            _ = v21
                        }
                    }
                case -50:
                    var v22 bool = (-700 != (v1 - v1))
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 int32 = -1
                        if true {
                            fmt.Println(v1, v1, v24)
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        fmt.Println(fi23)
                        _ = fi23
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 5; fi25++ {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = fi25
                }
            } else {
                v26 := GenS2{F0: "", F1: 0, F2: "a"}
                v1 = v1
                _ = v26
            }
            v1 += v1
            for fi27 := 0; fi27 < 1; fi27++ {
                if true {
                    if ((float32(1.5) / (float32(-2.25) + float32(float64((float32((((((-2.25 + 0.0) * -2.25) / 1e10) / float64(1.5)) + 0.1)) - 3.14159))))) <= 3.14159) {
                        switch (fi27 + (v1 >> ((18446744073709551615) & 63))) {
                        case -2147483648:
                            v28 := GenS2{F0: ("a" + ("hello" + "foo bar")), F1: 1, F2: ("a" + "hello")}
                            _ = v28
                        default:
                            for fi29 := 0; fi29 < 20; fi29++ {
                                v1 = 9223372036854775807
                                v1 = (fi29 * (^(fi27)))
                                v1 += (fi29 % ((fi29) | 1))
                                _ = fi29
                            }
                        }
                    }
                }
                for fi30 := 0; fi30 < 17; fi30++ {
                    v1 += (-(fi30))
                    var v31 uint8 = 0
                    _ = fi30
                    _ = v31
                }
                _ = fi27
            }
        } else {
            v1 += (v1 | genFn0(2))
            v1 = 127
        }
    case "hello":
        if (-2.25 <= 1e10) {
            for fi32 := 0; fi32 < 12; fi32++ {
                v1 += v1
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 int16 = 1
                    _ = fi33
                    _ = v34
                }
                v3 = append(v3, ("\n\t" + (("hello" + "hello") + ("a" + "foo bar"))))
                _ = fi32
            }
        } else {
            for fi35 := 0; fi35 < 1; fi35++ {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v1, v1)
                }
                _ = fi35
            }
        }
    case "":
        v36 := GenS3{F0: -1}
        _ = v36
    }
    v1 = genFn0(18446744073709551615)
    for fi37 := 0; fi37 < 18; fi37++ {
        if len(v3) > 0 {
            v3[0] = ("" + "hello")
        }
        var v38 string = "hello"
        _ = fi37
        _ = v38
    }
    var v39 int64 = 823
    if (false && ((((("\n\t" + "a") + "") + ("\n\t" + "\n\t")) > "a") || (false && ((-(v1)) != (v1 - (v1 ^ v1)))))) {
        v40 := v2[v1]
        if ((!(!(v39 < -1))) && (0 > 118)) {
            fmt.Println(v1, v40)
            v40 = v40
            v41 := GenS4{F0: ((42 < -932) || false), F1: -27}
            _ = v41
        }
        _ = v40
    } else {
        v42 := GenS5{F0: 2147483648, F1: 436}
        for fi43 := 0; fi43 < 9; fi43++ {
            if true {
                for fi44 := 0; fi44 < 11; fi44++ {
                    v39 = v39
                    _ = fi44
                }
                var v45 int = (fi43 * fi43)
                _ = v45
            }
            if true {
                delete(v2, v1)
                var v46 int = (fi43 ^ v1)
                v46 = v46
                _ = v46
            }
            _ = fi43
        }
        _ = v42
    }
    if false {
        fmt.Println(v1, v1)
        var v47 string = "hello"
        _ = v47
    }
    v1 -= v1
    for fi48 := 0; fi48 < 14; fi48++ {
        v49 := GenS4{F0: true, F1: 100}
        v39 += (v39 % ((v39) | 1))
        _ = fi48
        _ = v49
    }
    var v50 uint = 286
    if (-110 >= 1) {
        var v51 float64 = 1e10
        switch (("foo bar" + "\n\t") + ("hello" + "a")) {
        case "foo bar":
            v52 := GenS2{F0: "hello", F1: 947, F2: ("" + ("" + "foo bar"))}
            _ = v52
        default:
            if true {
                if (803 < 778) {
                    v53 := GenS2{F0: ("" + (("foo bar" + (("" + ("hello" + "\n\t")) + "hello")) + "foo bar")), F1: 125, F2: "foo bar"}
                    v54 := v2[v1]
                    switch "a" {
                    case "foo bar":
                        v55 := GenS2{F0: ((v54 + ("" + (v54 + ("a" + v54)))) + ("hello" + "foo bar")), F1: 886, F2: v54}
                        _ = v55
                    }
                    _ = v53
                    _ = v54
                } else {
                    v56 := GenS5{F0: 4294967295, F1: 2147483647}
                    _ = v56
                }
            } else {
                v57 := GenS6{F0: "a", F1: (v39 ^ (v39 & (^(v39))))}
                _ = v57
            }
        }
        var v58 uint16 = 65535
        _ = v51
        _ = v58
    }
    for fi59 := 0; fi59 < 9; fi59++ {
        if (258 <= 761) {
            if (42 < 65) {
                switch "hello" {
                case "\n\t":
                    v60 := GenS1{F0: 731, F1: genFn0(v50)}
                    _ = v60
                default:
                    _ = len(v2)
                }
            } else {
                v50 = 2
            }
            v61 := GenS7{F0: (fi59 >> ((0) & 63)), F1: ("foo bar" + "a")}
            _ = v61
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v50
}
