package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int8; F1 bool; F2 uint16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint8; F1 string; F2 string }
type GenS8 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 uint16) uint {
    v7 := GenS1{F0: 32767, F1: 338}
    var v8 int = 1
    p0 = (p0 / (((^(p0))) | 1))
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]string)
    var v3 bool = false
    v4 := make(map[uint32]int)
    v5 := []int{}
    var v6 uint8 = 127
    if false {
        var v9 int32 = -2147483648
        var v10 int = 34
        _ = v10
        _ = v9
    } else {
        var v11 int64 = 581
        switch "foo bar" {
        case "foo bar":
            v12 := v2[2147483647]
            _ = v12
        case "a":
            for fi13 := 0; fi13 < 18; fi13++ {
                switch ("a" + ("\n\t" + "hello")) {
                case "a":
                    if (65535 != 18446744073709551615) {
                        v14 := GenS2{F0: 877, F1: (("foo bar" + "foo bar") + ((("foo bar" + "\n\t") + "a") + "\n\t"))}
                        v11 -= (v11 | v11)
                        _ = v14
                    }
                case "":
                    switch "\n\t" {
                    case "a":
                        v6 -= (v6 << ((18446744073709551615) & 63))
                    case "foo bar":
                        var v15 uint32 = 65535
                        _ = v15
                    }
                default:
                    if (false && v3) {
                        v5 = append(v5, fi13)
                        v16 := GenS2{F0: 30, F1: "\n\t"}
                        if v3 {
                            switch "\n\t" {
                            case "hello":
                                if (-610 != 2147483647) {
                                    v3 = v3
                                    fmt.Println(v1, v11)
                                    for fi17 := 0; fi17 < 14; fi17++ {
                                        v18 := GenS3{F0: -456, F1: 401}
                                        fmt.Println(fi13, v1)
                                        var v19 int64 = v11
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                    }
                                }
                            case "":
                                v20 := GenS4{F0: 32767, F1: 0}
                                _ = v20
                            }
                            fmt.Println(v11)
                        } else {
                            var v21 int16 = 0
                            _ = v21
                        }
                        _ = v16
                    } else {
                        switch -2147483648 {
                        case 9223372036854775807:
                            if (true && v3) {
                                fmt.Println(v1, v1)
                                var v22 bool = true
                                v6 = (v6 >> ((18446744073709551615) & 63))
                                _ = v22
                            } else {
                                switch (fi13 % ((fi13) | 1)) {
                                case -1:
                                    v6 -= 128
                                case 4294967295:
                                    v23 := GenS3{F0: int16(fi13), F1: 481}
                                    _ = v23
                                }
                                var v24 int32 = -1
                                _ = v24
                            }
                        default:
                            var v25 float32 = -2.25
                            _ = v25
                        }
                        if (v3 || (!(0.0 < 0.0))) {
                            _ = len(v2)
                        } else {
                            v26 := GenS5{F0: (v1 % (((v1 & (v1 + (v1 | -1)))) | 1)), F1: v3, F2: 255}
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v3 = true
                            v28 := GenS6{F0: v1, F1: v1}
                            _ = fi27
                            _ = v28
                        }
                    }
                }
                v6 = v6
                v29 := GenS4{F0: 463, F1: 0}
                _ = fi13
                _ = v29
            }
        case "hello":
            var v30 int64 = -1
            _ = v30
        default:
            switch ("a" + "") {
            case "hello":
                var v31 string = ("a" + "\n\t")
                _ = v31
            }
        }
        v3 = true
        _ = v11
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        delete(v2, 0)
        _ = fi32
    }
    v6 -= 127
    if false {
        var v33 int = 9223372036854775807
        _ = v33
    } else {
        fmt.Println(v1)
        v5 = append(v5, -1)
        if false {
            v34 := GenS7{F0: (^(v6)), F1: "", F2: ("hello" + ("\n\t" + ("" + "")))}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch 127 {
            case 844:
                v3 = (1e10 < 3.14159)
            default:
                var v35 uint64 = 42
                _ = v35
            }
            _ = v34
        } else {
            v3 = (v3 || (((-2.25 / (1.5 + (1e10 + -2.25))) + 3.14159) <= (0.1 * (((1.5 / -2.25) * 0.0) / ((0.1 - 1e10) / 1e10)))))
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 float64 = (0.0 * 0.0)
                _ = fi36
                _ = v37
            }
        }
    }
    fmt.Println(v1, v1, v6)
    var v38 uint64 = 0
    var v39 uint32 = 536
    if true {
        for fi40 := 0; fi40 < 16; fi40++ {
            v6 += v6
            if true {
                if (1 <= v6) {
                    for fi41 := 0; fi41 < 18; fi41++ {
                        if v3 {
                            if v3 {
                                v3 = (211 <= int32(fi40))
                                v39 -= (v39 & (v39 << ((genFn0(0)) & 63)))
                                if v3 {
                                    var v42 uint16 = 25
                                    v39 -= (v39 + v39)
                                    _ = v42
                                }
                                fmt.Println(v1, v39, v1)
                            }
                            v43 := GenS1{F0: -417, F1: 924}
                            _ = v43
                        } else {
                            v1 = int8(v39)
                        }
                        _ = fi41
                    }
                }
                v44 := GenS5{F0: 42, F1: v3, F2: uint16(v39)}
                _ = v44
            }
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := GenS7{F0: v6, F1: ("" + "hello"), F2: ("" + (("foo bar" + "hello") + "a"))}
                v39 -= v39
                _ = fi45
                _ = v46
            }
            _ = fi40
        }
    } else {
        if (v38 > v38) {
            if (v3 || v3) {
                if ((!(!false)) && v3) {
                    switch 128 {
                    case 2147483647:
                        var v47 int = 127
                        _ = v47
                    case -362:
                        var v48 string = ((("" + "\n\t") + ("\n\t" + ((("" + "hello") + "") + (("" + "") + "")))) + (("hello" + ("hello" + "a")) + ("hello" + "foo bar")))
                        _ = v48
                    }
                    var v49 uint16 = 594
                    _ = v49
                }
                v50 := GenS8{F0: v6, F1: 127, F2: v38}
                _ = v50
            } else {
                v51 := v4[v39]
                _ = v51
            }
            switch "foo bar" {
            case "foo bar":
                v39 = (^(v39))
            }
            var v52 int64 = -1
            _ = v52
        }
        var v53 uint8 = (^(v6))
        v39 -= (^(v39))
        _ = v53
    }
    for fi54 := 0; fi54 < 1; fi54++ {
        var v55 uint = 42
        v56 := GenS1{F0: int16(v39), F1: 0}
        _ = fi54
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
