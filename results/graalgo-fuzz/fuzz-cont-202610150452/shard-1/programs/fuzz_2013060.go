package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 int8; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 int; F2 int64 }
type GenS4 struct { F0 int64; F1 int8; F2 uint }
type GenS5 struct { F0 int }
func genFn0() uint32 {
    var v5 uint = 4294967295
    var v6 uint64 = 0
    v5 -= (v5 & v5)
    v7 := []int{1, 128, -803}
    _ = len(v7)
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() int {
    var v8 uint16 = 940
    fmt.Println(v8)
    _ = v8
    return 756
}

func main() {
    var v1 uint16 = 2
    var v2 float32 = 1e10
    var v3 int = 0
    var v4 int32 = -761
    for fi9 := 0; fi9 < 13; fi9++ {
        switch "\n\t" {
        case "a":
            var v10 uint8 = 128
            _ = v10
        case "":
            fmt.Println(v3, fi9, v4)
        }
        var v11 uint8 = uint8(v1)
        var v12 uint = 698
        _ = fi9
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: true, F1: true, F2: 1}
    if (11 >= 100) {
        if (0 == genFn0()) {
            v14 := []int{(v3 | 724), (v3 + v3), genFn1()}
            if (((v3 - (^(v3))) < (v3 / (((v3 * (v3 / ((v3) | 1)))) | 1))) || false) {
                v14 = append(v14, (v3 << ((4294967295) & 63)))
                v14 = append(v14, v3)
            } else {
                _ = len(v14)
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS1{F0: (0 != 873), F1: ((genFn0() <= 2147483647) || true), F2: v4}
                    var v17 int16 = 317
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 float64 = float64(((v2 - v2) / v2))
                _ = v18
            }
            v19 := GenS1{F0: true, F1: true, F2: (v4 - v4)}
            _ = v14
            _ = v19
        }
        v20 := []int{v3, v3, (v3 + (v3 + (v3 + 287)))}
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := []int{fi21, fi21}
        var v23 string = "hello"
        _ = fi21
        _ = v22
        _ = v23
    }
    v4 = (v4 - 928)
    if false {
        if true {
            v24 := GenS1{F0: ((!true) && (!(65535 == 4294967295))), F1: (-128 != -99), F2: v4}
            _ = v24
        } else {
            v25 := GenS2{F0: 100, F1: (("hello" + "\n\t") + ((("foo bar" + (("hello" + "\n\t") + "\n\t")) + (("hello" + "\n\t") + "")) + ("a" + ""))), F2: ("foo bar" + ((("a" + "") + "hello") + ("\n\t" + ("hello" + "foo bar"))))}
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 uint32 = 797
                v28 := []int{genFn1()}
                for fi29 := 0; fi29 < 19; fi29++ {
                    if len(v28) > 0 {
                        v28[0] = 549
                    }
                    var v30 string = ""
                    _ = fi29
                    _ = v30
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            v4 += v4
            _ = v25
        }
        var v31 int8 = 127
        if (!true) {
            v32 := GenS2{F0: 0, F1: "foo bar", F2: "a"}
            switch ("foo bar" + ((("a" + ((("hello" + "") + "hello") + "")) + "") + "foo bar")) {
            case "a":
                var v33 string = "hello"
                _ = v33
            case "":
                for fi34 := 0; fi34 < 18; fi34++ {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 uint16 = (^(v1))
                        v37 := GenS2{F0: (v31 ^ 100), F1: "", F2: ("" + "\n\t")}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    if true {
                        v38 := GenS3{F0: 42, F1: (fi34 << ((42) & 63)), F2: 505}
                        _ = v38
                    } else {
                        var v39 bool = (4294967295 >= 7)
                        v1 = 1
                        var v40 int8 = v31
                        _ = v39
                        _ = v40
                    }
                    _ = fi34
                }
            case "hello":
                var v41 uint8 = 127
                _ = v41
            }
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := []int{fi42, (v3 / (((fi42 / (((^(v3))) | 1))) | 1))}
                var v44 int8 = (v31 % ((-12) | 1))
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = v32
        }
        _ = v31
    }
    v45 := GenS4{F0: 9223372036854775807, F1: -1, F2: 42}
    for fi46 := 0; fi46 < 6; fi46++ {
        var v47 uint = 2
        _ = fi46
        _ = v47
    }
    if true {
        var v48 uint64 = 1
        v49 := []int{386}
        for fi50 := 0; fi50 < 7; fi50++ {
            v48 = v48
            _ = fi50
        }
        _ = v48
        _ = v49
    }
    v51 := []int{-906, (v3 >> ((42) & 63)), (v3 - -861), int(v1)}
    if len(v51) >= 2 {
        v51 = v51[1:len(v51)]
    }
    if ((v1 | (v1 * 495)) >= 433) {
        v4 += v4
        for fi52 := 0; fi52 < 17; fi52++ {
            for fi53 := 0; fi53 < 14; fi53++ {
                var v54 uint = 18446744073709551615
                _ = fi53
                _ = v54
            }
            if len(v51) > 0 {
                v51[0] = v3
            }
            var v55 uint32 = 2147483648
            _ = fi52
            _ = v55
        }
        if (9223372036854775807 < 1) {
            if false {
                var v56 int16 = 32767
                _ = v56
            }
        } else {
            for fi57 := 0; fi57 < 16; fi57++ {
                fmt.Println(v1)
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
                _ = fi57
            }
            for fi58 := 0; fi58 < 16; fi58++ {
                v59 := GenS2{F0: 0, F1: "hello", F2: "hello"}
                fmt.Println(fi58)
                v60 := GenS5{F0: (v3 & 0)}
                _ = fi58
                _ = v59
                _ = v60
            }
        }
    } else {
        if ((-425 < 0) || true) {
            var v61 int8 = -1
            _ = v61
        }
        if false {
            for fi62 := 0; fi62 < 14; fi62++ {
                if false {
                    for fi63 := 0; fi63 < 17; fi63++ {
                        fmt.Println(fi63, v2, v4)
                        var v64 int32 = 42
                        _ = len(v51)
                        switch "" {
                        case "hello":
                            var v65 string = ("foo bar" + "foo bar")
                            _ = v65
                        default:
                            if (-54 <= 65) {
                                fmt.Println(v1, v4)
                            }
                        }
                        _ = fi63
                        _ = v64
                    }
                    v4 = 1
                }
                _ = fi62
            }
            v1 = v1
            if true {
                if ((1.5 + v2) <= 1e10) {
                    for fi66 := 0; fi66 < 13; fi66++ {
                        if (166 == 142) {
                            v2 -= v2
                            v4 -= v4
                        }
                        var v67 int64 = 943
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                        _ = fi66
                        _ = v67
                    }
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                } else {
                    fmt.Println(v3)
                    v51 = append(v51, 127)
                }
                v2 = v2
            }
        } else {
            for fi68 := 0; fi68 < 11; fi68++ {
                v69 := GenS4{F0: -9223372036854775808, F1: 125, F2: 18446744073709551615}
                _ = fi68
                _ = v69
            }
        }
        var v70 int32 = 1
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v51
}
