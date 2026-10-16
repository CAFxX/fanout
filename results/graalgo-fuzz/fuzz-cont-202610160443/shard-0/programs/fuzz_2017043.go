package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint64) uint32 {
    v5 := []int{-1, 32768, 9223372036854775807, 32767}
    var v6 int64 = 754
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 22
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 42
    v3 := make(map[int]int)
    var v4 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 1; fi7++ {
        if false {
            if (-2.25 > 1.5) {
                if (false && (-81 <= 42)) {
                    v2 = (^(v2))
                    var v8 string = "foo bar"
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        if ((((-1 >= 42) && true) && false) || (2 == 2)) {
                            v10 := []int{(fi9 + (fi9 >> ((429) & 63)))}
                            if (32767 != (v2 / ((v2) | 1))) {
                                var v11 int32 = -2147483648
                                v12 := GenS1{F0: genFn0((1e10 < float32(((1.5 * -2.25) * -2.25))), 0), F1: 2, F2: 122}
                                _ = v11
                                _ = v12
                            }
                            _ = len(v1)
                            _ = v10
                        } else {
                            if ((v2 / (((v2 ^ (v2 % ((690) | 1)))) | 1)) <= v2) {
                                if true {
                                    var v13 int = fi9
                                    _ = v13
                                } else {
                                    v2 -= (v2 + v2)
                                    if false {
                                        var v14 uint16 = v2
                                        var v15 int16 = 42
                                        fmt.Println(fi9, v15, v15)
                                        delete(v3, fi9)
                                        _ = v14
                                        _ = v15
                                    } else {
                                        v16 := GenS1{F0: 4294967295, F1: (v2 ^ (^(v2))), F2: -905}
                                        _ = v16
                                    }
                                }
                                v4 = 0
                                var v17 uint32 = 2
                                _ = v17
                            }
                        }
                        if true {
                            var v18 int32 = -1
                            _ = v18
                        }
                        _ = fi9
                    }
                }
            } else {
                v19 := GenS1{F0: 4294967295, F1: 1, F2: 454}
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS2{F0: -947}
                _ = fi20
                _ = v21
            }
        }
        v2 = (v2 % ((v2) | 1))
        v2 = v2
        _ = fi7
    }
    if false {
        var v22 int32 = 789
        _ = v22
    }
    var v23 int64 = (v4 * -434)
    var v24 uint16 = v2
    if (2 < 4294967295) {
        for fi25 := 0; fi25 < 13; fi25++ {
            if (("foo bar" + ("hello" + "\n\t")) < "hello") {
                switch ("foo bar" + "") {
                case "a":
                    if ((false || (!false)) || true) {
                        v26 := v1[(fi25 >> ((4294967295) & 63))]
                        _ = v26
                    }
                default:
                    v27 := v3[(fi25 * -2147483648)]
                    _ = v27
                }
                v28 := GenS1{F0: 652, F1: (v2 | (v2 & (v2 * (v24 >> ((255) & 63))))), F2: 539}
                _ = v28
            } else {
                var v29 string = ("" + "a")
                var v30 uint16 = 2
                _ = v29
                _ = v30
            }
            _ = fi25
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v4, v2)
        default:
            v31 := GenS3{F0: 42, F1: false, F2: ((((v2 >= 1) && false) || (42 < 42)) && true)}
            _ = v31
        }
    }
    switch 327 {
    case -740:
        v4 += (v23 & v4)
    }
    _ = len(v3)
    v32 := []int{-226, 2147483647, 2147483647, 354}
    if ("a" >= ("" + (("" + "\n\t") + "foo bar"))) {
        v1[311] = (!(128 > 127))
        v33 := GenS4{F0: 158}
        var v34 uint64 = 85
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v4
}
