package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 bool; F1 uint }
func genFn0() int16 {
    var v5 int8 = 1
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 4294967295
        if false {
            var v8 int = fi6
            v9 := GenS1{F0: (true && (!(((("" + "a") + "foo bar") + "a") != ("" + "")))), F1: true, F2: v8}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS1{F0: false, F1: false, F2: int(v7)}
                v7 -= v7
                var v12 uint64 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
            if ((false || (0 == 0)) || false) {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int = fi13
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 float64 = 0.1
                    _ = fi15
                    _ = v16
                }
            } else {
                if true {
                    fmt.Println(v7, fi6, v5)
                } else {
                    var v17 float32 = 1e10
                    var v18 int16 = 0
                    var v19 uint64 = 4294967295
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v20 := []int{}
                var v21 int64 = 9223372036854775807
                _ = v20
                _ = v21
            }
        }
        _ = fi6
        _ = v7
    }
    _ = v5
    return -1
}
func genFn1(p0 uint16) string {
    p0 = p0
    var v22 int64 = -171
    var v23 uint64 = 18446744073709551615
    _ = p0
    _ = v22
    _ = v23
    return ((("foo bar" + ("foo bar" + ("" + "a"))) + ("hello" + "a")) + "")
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 1e10
    var v3 float64 = 1.5
    var v4 uint16 = 2
    v2 = float32((((v3 - v3) / (3.14159 * v3)) + ((v3 - 0.0) - (v3 / 0.1))))
    v2 = v2
    switch "\n\t" {
    case "a":
        if (("hello" < "a") && false) {
            if false {
                var v24 int8 = -128
                v2 = v2
                _ = v24
            } else {
                var v25 int8 = 1
                _ = v25
            }
            var v26 int64 = 0
            _ = v26
        }
    case "\n\t":
        switch "a" {
        case "foo bar":
            v27 := []int{-717}
            _ = v27
        case "\n\t":
            var v28 int16 = 501
            _ = v28
        case "hello":
            var v29 bool = true
            _ = v29
        }
    }
    v4 = (v4 * v4)
    var v30 int8 = 42
    v4 = (^(v4))
    v31 := GenS2{F0: 4294967295, F1: 42}
    if (("hello" + ("a" + ("" + ("a" + "foo bar")))) > ("" + "")) {
        for fi32 := 0; fi32 < 12; fi32++ {
            fmt.Println(v3)
            var v33 uint32 = 1
            switch "foo bar" {
            case "":
                v34 := GenS1{F0: ((v30 / ((v30) | 1)) > (v30 % ((-70) | 1))), F1: (!((!true) || false)), F2: fi32}
                _ = v34
            case "hello":
                for fi35 := 0; fi35 < 9; fi35++ {
                    fmt.Println(v2)
                    _ = fi35
                }
            case "\n\t":
                if (!((false && (v33 <= v33)) && false)) {
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v37 := GenS3{F0: 2}
                        var v38 int16 = -32768
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 11; fi39++ {
                        var v40 int64 = 9223372036854775807
                        if ((^(v33)) != 760) {
                            v3 -= v3
                            switch ((("\n\t" + "") + (("hello" + ("\n\t" + "foo bar")) + ("\n\t" + ("a" + "a")))) + "\n\t") {
                            case "\n\t":
                                var v41 string = ""
                                v42 := []int{32767}
                                v3 -= v3
                                _ = v41
                                _ = v42
                            case "foo bar":
                                v43 := GenS2{F0: uint64(v4), F1: 447}
                                _ = v43
                            case "":
                                var v44 uint32 = (v33 ^ v33)
                                _ = v44
                            default:
                                if false {
                                    var v45 int32 = -1
                                    _ = v45
                                } else {
                                    v46 := []int{-28, (fi32 << ((0) & 63)), fi39, 2147483647}
                                    v46 = append(v46, 32768)
                                    _ = v46
                                }
                                fmt.Println(fi32)
                            }
                            v3 = -2.25
                        }
                        v4 -= v4
                        v47 := GenS4{F0: 255}
                        _ = fi39
                        _ = v40
                        _ = v47
                    }
                    v48 := GenS5{F0: (true && false)}
                    _ = v48
                } else {
                    switch "" {
                    case "hello":
                        var v49 uint32 = (^(v33))
                        _ = v49
                    case "foo bar":
                        _ = len(v1)
                    case "\n\t":
                        fmt.Println(v30)
                    default:
                        var v50 int8 = (v30 - -32)
                        _ = v50
                    }
                }
            default:
                fmt.Println(v4, v30)
            }
            _ = fi32
            _ = v33
        }
    } else {
        var v51 bool = (553 < 2147483647)
        fmt.Println(v3, v30, v30)
        var v52 string = ""
        _ = v51
        _ = v52
    }
    v53 := GenS4{F0: 4294967295}
    for fi54 := 0; fi54 < 20; fi54++ {
        if ((fi54 & 128) < fi54) {
            v55 := []int{}
            if ((-1 > 100) && (214 > 2147483648)) {
                switch (-(fi54)) {
                case 128:
                    var v56 bool = true
                    _ = v56
                case 32767:
                    fmt.Println(v3, v3, fi54)
                default:
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v58 := GenS6{F0: 0}
                        v2 = (v2 - 3.14159)
                        _ = fi57
                        _ = v58
                    }
                }
                var v59 uint16 = 452
                v4 = 420
                _ = v59
            }
            _ = v55
        }
        v60 := GenS7{F0: 2}
        v61 := []int{}
        _ = fi54
        _ = v60
        _ = v61
    }
    if false {
        switch 4294967295 {
        case 127:
            for fi62 := 0; fi62 < 7; fi62++ {
                switch 32767 {
                case -1:
                    if (v30 > v30) {
                        var v63 int8 = (v30 & (v30 & v30))
                        var v64 int64 = 1
                        _ = v63
                        _ = v64
                    } else {
                        switch (fi62 + fi62) {
                        case 4294967295:
                            v4 = v4
                        default:
                            var v65 int32 = 1
                            _ = v65
                        }
                    }
                default:
                    for fi66 := 0; fi66 < 4; fi66++ {
                        v67 := GenS8{F0: 42}
                        v30 = v30
                        switch (fi66 | (fi66 ^ fi66)) {
                        case 128:
                            v68 := GenS8{F0: 65535}
                            _ = v68
                        case 430:
                            var v69 int16 = 152
                            _ = v69
                        case 9223372036854775807:
                            v70 := []int{128, fi62, fi62}
                            _ = v70
                        }
                        _ = fi66
                        _ = v67
                    }
                }
                if false {
                    v71 := []int{(fi62 / ((1) | 1)), (fi62 % ((fi62) | 1)), (fi62 ^ fi62)}
                    if true {
                        if false {
                            switch "hello" {
                            case "\n\t":
                                v3 = v3
                            }
                            var v72 int64 = 9223372036854775807
                            _ = v72
                        }
                    } else {
                        for fi73 := 0; fi73 < 12; fi73++ {
                            v4 = (v4 % (((v4 % (((^(v4))) | 1))) | 1))
                            v74 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                            var v75 string = ""
                            _ = fi73
                            _ = v74
                            _ = v75
                        }
                        if true {
                            v2 -= 0.1
                            fmt.Println(v2)
                            v76 := v1["foo bar"]
                            _ = v76
                        } else {
                            v77 := GenS9{F0: 550, F1: 42}
                            _ = v77
                        }
                    }
                    _ = v71
                }
                _ = fi62
            }
        case 9223372036854775807:
            v78 := GenS10{F0: false, F1: 0}
            _ = v78
        default:
            if (65535 <= 65535) {
                fmt.Println(v30, v2)
                v79 := []int{}
                v80 := GenS11{F0: (true || (true || false)), F1: 684}
                _ = v79
                _ = v80
            }
        }
        delete(v1, "hello")
        v3 += (0.1 / v3)
    }
    v81 := []int{}
    delete(v1, "")
    var v82 uint8 = 42
    v2 = (v2 * 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v53
    _ = v81
    _ = v82
}
