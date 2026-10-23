package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 int; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    var v4 uint16 = 32767
    var v5 float32 = (-2.25 * 1.5)
    v6 := make(map[uint32]string)
    var v7 int16 = -32768
    _ = len(v1)
    v8 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808, F2: v7}
    for fi9 := 0; fi9 < 16; fi9++ {
        switch fi9 {
        case -292:
            if false {
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := v6[374]
                    var v12 uint16 = v4
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v7 -= (^(v7))
                for fi13 := 0; fi13 < 15; fi13++ {
                    v6[0] = ("hello" + ("" + "a"))
                    v14 := GenS2{F0: 968, F1: (v3 % ((255) | 1))}
                    var v15 int16 = v7
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch (fi9 ^ (-(fi16))) {
                    case 4294967295:
                        var v17 int32 = -2147483648
                        _ = v17
                    case -9223372036854775808:
                        var v18 uint32 = 2147483647
                        _ = v18
                    case 731:
                        if true {
                            v3 = uint8(v4)
                            if false {
                                v19 := []int{(^(fi9)), (^(fi9)), (fi9 | (fi9 | (fi9 & fi9))), fi16}
                                var v20 float32 = 3.14159
                                if (-2.25 <= v5) {
                                    if false {
                                        v21 := GenS1{F0: -170, F1: 548, F2: v7}
                                        _ = v21
                                    }
                                    v20 += 0.1
                                    var v22 uint16 = v4
                                    fmt.Println(v20, fi16)
                                    _ = v22
                                } else {
                                    switch (fi16 * fi9) {
                                    case 42:
                                        fmt.Println(v5, fi16, v20)
                                    case 32768:
                                        switch fi16 {
                                        case 0:
                                            v4 = v4
                                        case 32768:
                                            var v23 float32 = v20
                                            _ = v23
                                        case -569:
                                            if (v4 == (^(v4))) {
                                                v24 := GenS3{F0: 2147483647, F1: v7, F2: 128}
                                                _ = v24
                                            }
                                        }
                                    case -1:
                                        var v25 int32 = 42
                                        _ = v25
                                    default:
                                        v1[("" + "hello")] = "foo bar"
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        if false {
                                            switch ("hello" + ((("hello" + (("a" + ((("a" + "\n\t") + "foo bar") + "a")) + "a")) + (("foo bar" + "\n\t") + ("" + ("" + "hello")))) + "hello")) {
                                            case "foo bar":
                                                var v26 uint32 = 175
                                                _ = v26
                                            case "":
                                                _ = len(v19)
                                                v27 := GenS4{F0: (v3 % ((v3) | 1))}
                                                v28 := GenS2{F0: 2, F1: (^(v3))}
                                                v5 = v20
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v3, v5, v7)
                                                case "":
                                                    var v29 int8 = 26
                                                    _ = v29
                                                default:
                                                    v5 = (-2.25 - v5)
                                                }
                                                var v30 uint32 = 42
                                                fmt.Println(fi9, v20, v5)
                                                v31 := v1[(((("a" + "foo bar") + "") + "hello") + "")]
                                                _ = v27
                                                _ = v28
                                                _ = v30
                                                _ = v31
                                            default:
                                                var v32 int8 = 1
                                                _ = len(v19)
                                                _ = v32
                                            }
                                            var v33 uint16 = (v4 & (v4 << ((2) & 63)))
                                            _ = v33
                                        }
                                    default:
                                        v7 += (v7 & 0)
                                    }
                                }
                                _ = v19
                                _ = v20
                            }
                        }
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v4, fi9)
                        default:
                            var v34 float64 = (1e10 - 1.5)
                            _ = v34
                        }
                        v4 -= v4
                        var v35 int = fi9
                        if true {
                            delete(v1, ("foo bar" + "foo bar"))
                            v36 := GenS5{F0: ("hello" + "foo bar"), F1: (v4 ^ v4)}
                            _ = v36
                        } else {
                            if (false || true) {
                                var v37 int64 = 1
                                v38 := GenS6{F0: ("a" + "hello"), F1: (2 <= 42)}
                                if (4294967295 > 42) {
                                    v39 := GenS7{F0: 128}
                                    var v40 uint64 = 970
                                    v37 += v37
                                    _ = v39
                                    _ = v40
                                }
                                _ = v37
                                _ = v38
                            } else {
                                v41 := GenS1{F0: 1, F1: 42, F2: (-(v7))}
                                _ = v41
                            }
                            fmt.Println(v5, v7)
                            var v42 uint64 = 499
                            v3 = (v3 % (((v3 / ((v3) | 1))) | 1))
                            delete(v1, "foo bar")
                            _ = v42
                        }
                        v43 := []int{2147483647}
                        _ = v35
                        _ = v43
                    }
                    var v44 uint32 = 924
                    v4 = 710
                    v4 = 32767
                    _ = fi16
                    _ = v44
                }
                var v45 uint16 = v4
                _ = v45
            }
        case 42:
            var v46 float64 = (-2.25 - (0.0 * 1.5))
            _ = v46
        case -1:
            var v47 uint32 = 0
            _ = v47
        }
        for fi48 := 0; fi48 < 15; fi48++ {
            v4 += (^(v4))
            v49 := GenS6{F0: "foo bar", F1: true}
            fmt.Println(v3, v3)
            _ = fi48
            _ = v49
        }
        for fi50 := 0; fi50 < 1; fi50++ {
            fmt.Println(fi9)
            _ = fi50
        }
        _ = fi9
    }
    v51 := []int{32768}
    var v52 uint32 = 420
    if (false || (false || (false || (false || (608 < 18446744073709551615))))) {
        var v53 int64 = -888
        v54 := GenS2{F0: 65535, F1: (^(v3))}
        v5 -= v5
        _ = v53
        _ = v54
    }
    v55 := GenS8{F0: "\n\t", F1: 32768, F2: 929}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v55
    _ = v6
    _ = v7
    _ = v8
}
