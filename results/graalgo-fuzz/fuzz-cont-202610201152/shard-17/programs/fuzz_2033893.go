package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int8 }
func genFn0() uint32 {
    v7 := []int{-896, 42}
    switch 9223372036854775807 {
    case -795:
        var v8 int32 = -2147483648
        _ = v8
    default:
        var v9 uint = 546
        _ = v9
    }
    var v10 int = 1
    if (!(false || (128 <= 222))) {
        v7 = append(v7, v10)
    }
    _ = v10
    _ = v7
    return 65535
}
func genFn1(p0 string, p1 int) int64 {
    p1 = (p1 & (p1 << ((55) & 63)))
    if (!((float32(1e10) <= 1.5) || false)) {
        v11 := GenS1{F0: 623, F1: (p1 + p1)}
        _ = v11
    } else {
        if true {
            var v12 uint16 = 0
            var v13 int8 = -128
            if (false || false) {
                v14 := GenS1{F0: 2147483647, F1: p1}
                _ = v14
            } else {
                v15 := []int{}
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS1{F0: 4294967295, F1: 363}
                    v18 := GenS1{F0: 4294967295, F1: p1}
                    v12 = v12
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
            _ = v12
            _ = v13
        }
        v19 := GenS2{F0: (((!false) && false) || (9223372036854775807 > 1))}
        _ = v19
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{32768, -1, -906, 32767}
    var v3 uint32 = 4294967295
    v4 := make(map[uint32]bool)
    var v5 float64 = 1e10
    var v6 uint8 = 199
    _ = len(v4)
    v3 -= genFn0()
    switch 128 {
    case -2147483648:
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := v1[""]
            for fi22 := 0; fi22 < 13; fi22++ {
                _ = len(v2)
                fmt.Println(v5, v21, v5)
                _ = fi22
            }
            if true {
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 int32 = -588
                    _ = fi23
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 14; fi25++ {
                    if false {
                        if (!((163 == uint(v3)) || ((false || true) || (1 != 86)))) {
                            switch "\n\t" {
                            case "foo bar":
                                v26 := GenS3{F0: v6}
                                _ = v26
                            case "":
                                if false {
                                    fmt.Println(fi20, fi25)
                                    var v27 uint64 = 224
                                    v28 := GenS4{F0: v3, F1: genFn0()}
                                    var v29 uint = 42
                                    if ((42 > 2147483647) && false) {
                                        fmt.Println(v27)
                                        var v30 uint32 = uint32(v27)
                                        _ = v30
                                    } else {
                                        _ = len(v2)
                                        v3 = (v3 - (v3 ^ 2147483648))
                                        var v31 uint = uint(v3)
                                        var v32 int64 = 42
                                        _ = v31
                                        _ = v32
                                    }
                                    if true {
                                        v33 := GenS4{F0: (v3 | uint32(v29)), F1: v3}
                                        if true {
                                            fmt.Println(v3, v3)
                                            v34 := v1[("hello" + "a")]
                                            v21 -= fi20
                                            var v35 uint32 = 2147483647
                                            _ = v34
                                            _ = v35
                                        }
                                        if ((42 < 42) || ((true && false) || ((v29 - v29) >= (v29 ^ (^(v29)))))) {
                                            _ = len(v2)
                                            v5 -= v5
                                        } else {
                                            var v36 int64 = genFn1("\n\t", (fi25 | fi25))
                                            _ = v36
                                        }
                                        _ = v33
                                    }
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                v5 += 3.14159
                            }
                            fmt.Println(fi25)
                            var v37 int8 = 21
                            _ = v37
                        } else {
                            switch "a" {
                            case "foo bar":
                                switch "hello" {
                                case "":
                                    v3 = v3
                                default:
                                    var v38 int16 = 32767
                                    _ = v38
                                }
                            default:
                                if ((!false) && true) {
                                    var v39 uint8 = (v6 ^ 2)
                                    _ = v39
                                }
                            }
                            v40 := GenS5{F0: 0}
                            fmt.Println(v3, v21, v21)
                            _ = v40
                        }
                        v3 = 2147483648
                    }
                    _ = fi25
                }
                v5 += float64(-2.25)
                v5 = ((v5 + ((-2.25 / 0.1) * (float64(float32((1e10 - v5))) - (v5 + v5)))) + ((1e10 - v5) / ((v5 * (v5 - ((v5 * (v5 - v5)) - -2.25))) * (v5 - v5))))
            }
            _ = fi20
            _ = v21
        }
    }
    switch 32768 {
    case 128:
        var v41 float64 = (v5 * 0.0)
        _ = v41
    }
    v6 += (v6 + (v6 | (v6 << ((65535) & 63))))
    v6 += 2
    var v42 uint = 255
    var v43 float32 = (float32(3.14159) * 1e10)
    v44 := GenS5{F0: 126}
    fmt.Println(v5, v5, v6)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
