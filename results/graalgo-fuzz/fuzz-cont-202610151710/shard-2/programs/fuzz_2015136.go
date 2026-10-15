package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 int32 }
type GenS5 struct { F0 string }
func genFn0() bool {
    var v4 uint16 = 0
    if ((530 > 0) && ((528 < 1000) && (!(!true)))) {
        v5 := []int{1, 1, 572}
        _ = v5
    }
    var v6 bool = true
    v6 = true
    var v7 float32 = (float32((1.5 * 1e10)) - 0.1)
    _ = v4
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 uint) uint8 {
    for fi8 := 0; fi8 < 2; fi8++ {
        switch 32768 {
        case 4294967295:
            if false {
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := []int{fi8, (fi9 | -971), (fi9 >> (((p0 & p0)) & 63)), (fi9 + 725)}
                    p0 = 379
                    if false {
                        p0 = 65535
                        fmt.Println(fi9, fi8, p0)
                        switch "" {
                        case "":
                            fmt.Println(p0)
                        default:
                            p0 = p0
                        }
                        switch "" {
                        case "a":
                            var v11 uint64 = 4294967295
                            _ = v11
                        case "hello":
                            fmt.Println(fi8, p0, p0)
                        default:
                            if true {
                                var v12 string = "\n\t"
                                if (v12 <= (v12 + v12)) {
                                    var v13 float64 = 0.1
                                    _ = v13
                                } else {
                                    p0 = p0
                                    if false {
                                        if false {
                                            fmt.Println(v12)
                                        } else {
                                            fmt.Println(v12, p0, p0)
                                            v14 := GenS1{F0: 42, F1: 766}
                                            _ = v14
                                        }
                                    }
                                }
                                _ = v12
                            }
                        }
                        var v15 int32 = -975
                        _ = v15
                    }
                    _ = fi9
                    _ = v10
                }
            }
        }
        switch fi8 {
        case 9223372036854775807:
            v16 := GenS2{F0: (fi8 + 662)}
            _ = v16
        case -869:
            p0 = p0
        default:
            v17 := []int{}
            _ = v17
        }
        _ = fi8
    }
    var v18 float64 = 1e10
    var v19 uint = p0
    if false {
        var v20 int32 = -83
        var v21 uint32 = 65535
        if (242 > -9223372036854775808) {
            v21 = v21
        }
        _ = v20
        _ = v21
    } else {
        var v22 int64 = -9223372036854775808
        v23 := GenS3{F0: 42}
        var v24 uint8 = 1
        _ = v22
        _ = v23
        _ = v24
    }
    if false {
        v25 := GenS1{F0: 2, F1: (p0 % ((v19) | 1))}
        _ = v25
    } else {
        p0 = (^(v19))
        v26 := []int{}
        _ = v26
    }
    _ = p0
    _ = v18
    _ = v19
    return 0
}

func main() {
    var v1 int64 = 1
    v2 := []string{"", ("" + (("hello" + "\n\t") + ("hello" + "\n\t"))), "\n\t", ""}
    v3 := []int{32767, 2147483647, 0}
    var v27 bool = false
    var v28 int64 = (-(v1))
    var v29 bool = (-2.25 == float64(float32(1.5)))
    var v30 uint8 = genFn1(4294967295)
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int32 = 1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi31
        _ = v32
    }
    v1 -= v1
    var v33 int64 = v1
    _ = len(v3)
    for fi34 := 0; fi34 < 17; fi34++ {
        if ((false && false) || v27) {
            v35 := GenS4{F0: (v30 << ((2) & 63)), F1: (fi34 * (fi34 % ((fi34) | 1))), F2: 42}
            _ = v35
        }
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := GenS5{F0: "hello"}
            _ = fi36
            _ = v37
        }
        for fi38 := 0; fi38 < 4; fi38++ {
            if (v1 != -1) {
                v29 = (65535 <= 633)
            } else {
                v29 = (3 != 127)
            }
            v28 -= v28
            _ = fi38
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
}
