package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 uint16 }
func genFn0(p0 int8, p1 int16) uint16 {
    var v6 string = ("foo bar" + "\n\t")
    v7 := GenS1{F0: (v6 + v6), F1: 618}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 450
}

func main() {
    var v1 uint8 = 0
    var v2 int8 = 127
    var v3 int = 2147483647
    v4 := []string{("" + ("foo bar" + "hello")), "hello"}
    v5 := make(map[uint32]string)
    if true {
        var v8 int = v3
        if (4294967295 <= 42) {
            var v9 float64 = 0.1
            v10 := GenS1{F0: "foo bar", F1: 42}
            _ = v10
            _ = v9
        } else {
            fmt.Println(v3, v1)
            v11 := GenS1{F0: "a", F1: 791}
            var v12 int8 = (v2 + v2)
            _ = v11
            _ = v12
        }
        var v13 int = (v3 ^ (v3 << ((65535) & 63)))
        _ = v13
        _ = v8
    } else {
        v2 = (v2 * (v2 + -1))
        var v14 int16 = 960
        _ = v14
    }
    var v15 uint64 = 375
    delete(v5, 42)
    v2 = (v2 / ((v2) | 1))
    if false {
        v15 = 0
    }
    fmt.Println(v3, v2)
    for fi16 := 0; fi16 < 6; fi16++ {
        if (int8(v1) == (v2 * v2)) {
            v17 := GenS1{F0: (("a" + ("foo bar" + ("hello" + "hello"))) + (("" + "foo bar") + "\n\t")), F1: 42}
            _ = v17
        } else {
            fmt.Println(fi16)
            v15 += (v15 % ((v15) | 1))
        }
        var v18 bool = true
        _ = fi16
        _ = v18
    }
    if (false && (214 == (v1 & (v1 % (((v1 - 0)) | 1))))) {
        v1 -= (v1 & (v1 ^ (^(v1))))
    } else {
        v1 = (v1 | v1)
        v3 = -9223372036854775808
    }
    v2 = (v2 * (v2 >> ((18446744073709551615) & 63)))
    if false {
        v19 := GenS2{F0: 1}
        v3 += (-(v3))
        _ = v19
    } else {
        v20 := GenS1{F0: "hello", F1: 42}
        switch v3 {
        case 421:
            var v21 bool = (!(1e10 >= (((0.1 / -2.25) - 3.14159) - (-2.25 / float32(1.5)))))
            _ = v21
        case -1:
            fmt.Println(v15, v15)
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            if (-1 <= -1) {
                for fi23 := 0; fi23 < 5; fi23++ {
                    if true {
                        v24 := GenS1{F0: ("foo bar" + ("" + (((("" + (("foo bar" + "hello") + "hello")) + ("foo bar" + ("foo bar" + "a"))) + ("" + "foo bar")) + "\n\t"))), F1: 9223372036854775807}
                        _ = v24
                    }
                    if ((true && (((!(437 > 575)) && true) && (v15 > v15))) && false) {
                        switch fi23 {
                        case 267:
                            v5[65535] = ("foo bar" + "")
                        case 1:
                            if true {
                                v2 = 127
                                if false {
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                    }
                                }
                            } else {
                                if false {
                                    switch fi23 {
                                    case 9223372036854775807:
                                        var v25 uint16 = 255
                                        _ = v25
                                    case 1:
                                        if (v15 < uint64(v1)) {
                                            v1 -= v1
                                        } else {
                                            v26 := GenS2{F0: 2147483647}
                                            var v27 int16 = int16(v1)
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v28 := GenS3{F0: v2, F1: "", F2: 32767}
                                    v2 = (v2 - (v2 & v2))
                                    v29 := GenS2{F0: 1}
                                    _ = v28
                                    _ = v29
                                }
                            }
                            v5[2] = ("hello" + "\n\t")
                            if ((!(1 != 0)) && (!(((1.5 * 1e10) <= 0.1) && ((^(v15)) == v15)))) {
                                v2 = -128
                            }
                        default:
                            _ = len(v5)
                        }
                        v30 := GenS3{F0: v2, F1: "", F2: 65535}
                        _ = v30
                    }
                    var v31 int64 = -9223372036854775808
                    _ = fi23
                    _ = v31
                }
            }
            _ = fi22
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
