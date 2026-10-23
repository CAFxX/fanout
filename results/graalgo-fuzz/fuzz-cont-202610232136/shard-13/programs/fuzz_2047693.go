package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int; F2 int16 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int64; F1 bool; F2 uint16 }
func genFn0(p0 uint8) uint32 {
    v4 := GenS1{F0: 31, F1: p0, F2: 32767}
    p0 += 1
    _ = p0
    _ = v4
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-1, 9223372036854775807, 32767, -9223372036854775808}
    var v3 string = ""
    v3 = v3
    v5 := GenS1{F0: 62, F1: 128, F2: 255}
    if true {
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                if true {
                    v3 = v3
                    if (42 == 18446744073709551615) {
                        v3 = "hello"
                    } else {
                        _ = len(v1)
                        v8 := GenS1{F0: -1, F1: 1, F2: 672}
                        switch fi6 {
                        case 2147483647:
                            if (-352 >= 657) {
                                delete(v1, ((("\n\t" + (((v3 + v3) + v3) + v3)) + v3) + (v3 + v3)))
                                var v9 float64 = -2.25
                                _ = v9
                            }
                            fmt.Println(v3)
                            var v10 uint = 18446744073709551615
                            _ = v10
                        case 534:
                            if (!true) {
                                v11 := GenS2{F0: 1}
                                v12 := GenS2{F0: 4294967295}
                                v13 := GenS3{F0: 67}
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                        case 708:
                            var v14 uint16 = 2
                            _ = v14
                        default:
                            v3 = ("foo bar" + (v3 + v3))
                        }
                        _ = v8
                    }
                    if true {
                        v15 := GenS1{F0: 42, F1: 78, F2: 255}
                        if (!true) {
                            var v16 float64 = 1e10
                            switch v3 {
                            case "":
                                v16 = (v16 + 0.0)
                            case "\n\t":
                                v3 = "a"
                            default:
                                v17 := GenS3{F0: 2147483648}
                                _ = v17
                            }
                            _ = v16
                        }
                        _ = v15
                    }
                }
                _ = fi7
            }
            var v18 int32 = 2147483647
            _ = fi6
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            if true {
                if true {
                    v3 = ("\n\t" + ((v3 + ("foo bar" + ((v3 + v3) + v3))) + ((v3 + (v3 + ((v3 + v3) + "foo bar"))) + v3)))
                }
            }
            _ = fi19
        }
        _ = len(v2)
    }
    v20 := GenS1{F0: 100, F1: 202, F2: 65535}
    v1[((("a" + v3) + "") + v3)] = genFn0(0)
    var v21 uint64 = 18446744073709551615
    var v22 uint32 = 1
    if (255 >= 654) {
        var v23 int8 = -32
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := GenS1{F0: (-(v23)), F1: uint8(v21), F2: 255}
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v23, v3, fi26)
                switch 684 {
                case 4294967295:
                    v27 := GenS4{F0: 896}
                    _ = v27
                }
                _ = fi26
            }
            var v28 int8 = (^(v23))
            _ = fi24
            _ = v25
            _ = v28
        }
        _ = v23
    } else {
        fmt.Println(v22, v3)
    }
    if (32767 <= 42) {
        v29 := GenS5{F0: 922}
        switch v3 {
        case "foo bar":
            v22 = v22
        case "\n\t":
            v21 += v21
        case "a":
            v21 = (v21 ^ (v21 % (((v21 << ((252) & 63))) | 1)))
        default:
            v22 += (^(v22))
        }
        switch ("hello" + (v3 + (v3 + ((("hello" + (v3 + "")) + (("foo bar" + (("a" + v3) + v3)) + "a")) + v3)))) {
        case "\n\t":
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 float32 = float32(float64((-2.25 * float32((0.0 / float64((0.1 + (float32(0.0) / 1.0))))))))
                v21 = v21
                _ = fi30
                _ = v31
            }
        case "foo bar":
            delete(v1, "\n\t")
        default:
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(v22, v21)
                if len(v2) > 0 {
                    v2[0] = fi32
                }
                _ = fi32
            }
        }
        _ = v29
    } else {
        var v33 int64 = 1
        _ = v33
    }
    v2 = append(v2, 128)
    switch (v3 + ((v3 + ("foo bar" + v3)) + v3)) {
    case "":
        if (-997 != -2147483648) {
            var v34 int16 = -632
            if false {
                var v35 float64 = (3.14159 - ((0.0 + -2.25) / 3.14159))
                if true {
                    v1["\n\t"] = v22
                    var v36 int32 = 0
                    if false {
                        var v37 float64 = 1.5
                        if (v3 == ("foo bar" + v3)) {
                            var v38 int64 = 275
                            var v39 uint32 = 2
                            v40 := v1[v3]
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        _ = v37
                    } else {
                        var v41 int = int(v34)
                        _ = v41
                    }
                    _ = v36
                }
                var v42 int64 = 9223372036854775807
                _ = v35
                _ = v42
            } else {
                v43 := GenS6{F0: 76, F1: 326, F2: 0}
                _ = v43
            }
            _ = v34
        } else {
            if (((v21 << ((564) & 63)) <= (v21 >> ((95) & 63))) || (127 != 0)) {
                fmt.Println(v22)
            } else {
                if false {
                    var v44 int16 = -754
                    _ = v44
                } else {
                    if true {
                        for fi45 := 0; fi45 < 18; fi45++ {
                            v46 := GenS5{F0: (v21 * 1)}
                            _ = fi45
                            _ = v46
                        }
                    }
                    v47 := GenS6{F0: 0, F1: 127, F2: -69}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v47
                }
            }
            v48 := GenS2{F0: 1}
            _ = v48
        }
    case "hello":
        v49 := GenS7{F0: ((v3 + (("\n\t" + "foo bar") + "")) + ("foo bar" + (v3 + "\n\t"))), F1: 127}
        _ = v49
    default:
        v21 -= 42
    }
    v50 := GenS8{F0: 348, F1: (!false), F2: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v5
    _ = v50
}
