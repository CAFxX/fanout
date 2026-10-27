package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 string = ""
    v2 := []int{-4}
    var v3 uint32 = 42
    var v4 float64 = (1e10 + 1e10)
    v5 := []uint32{(v3 & (v3 & 42)), (v3 ^ (v3 - (^(v3))))}
    v6 := []string{v1, "a", v1}
    v1 = v1
    v7 := GenS1{F0: 1, F1: 16, F2: -713}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            if true {
                var v10 int16 = 874
                if true {
                    v11 := GenS1{F0: (^(fi8)), F1: 706, F2: (-(v10))}
                    if (((v1 + ("" + "\n\t")) + ("" + "hello")) < "hello") {
                        switch "\n\t" {
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = 2147483647
                            }
                        }
                    } else {
                        _ = len(v2)
                        var v12 uint8 = 137
                        var v13 uint16 = 32767
                        switch "a" {
                        case "foo bar":
                            var v14 int64 = 524
                            _ = v14
                        default:
                            v15 := GenS1{F0: fi8, F1: 2, F2: v10}
                            _ = v15
                        }
                        v13 += (v13 - (v13 & v13))
                        if true {
                            v10 -= v10
                        }
                        v16 := GenS1{F0: (fi8 ^ int(v3)), F1: 995, F2: v10}
                        if (42 < 42) {
                            v10 += 1
                        } else {
                            v4 -= v4
                            v17 := GenS2{F0: 42}
                            _ = len(v6)
                            _ = v17
                        }
                        var v18 bool = (255 != 2)
                        _ = v12
                        _ = v13
                        _ = v16
                        _ = v18
                    }
                    var v19 int = fi8
                    v20 := GenS1{F0: (fi9 - (v19 ^ 1)), F1: 2, F2: 801}
                    var v21 float64 = v4
                    _ = v11
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v4 += ((v4 - 0.1) + float64(0.0))
                _ = v10
            } else {
                var v22 uint16 = 42
                _ = v22
            }
            fmt.Println(v4, v4)
            _ = fi9
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS2{F0: 813}
            var v25 uint = 369
            if true {
                v26 := GenS2{F0: -1}
                _ = v26
            }
            v25 -= (^(v25))
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi8
    }
    v4 += (0.0 * (((v4 + (v4 + v4)) + (v4 * float64(-2.25))) * v4))
    if len(v2) > 0 {
        v2[0] = 267
    }
    fmt.Println(v4)
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 int8 = 51
        if false {
            if true {
                fmt.Println(v1)
                v1 = (v1 + "")
            }
            for fi29 := 0; fi29 < 4; fi29++ {
                v28 = v28
                v5 = append(v5, (v3 + 946))
                _ = fi29
            }
            switch "a" {
            case "hello":
                v28 = (v28 >> ((1) & 63))
            case "a":
                for fi30 := 0; fi30 < 18; fi30++ {
                    v4 += (v4 + 1.5)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v31 uint8 = 114
                    _ = fi30
                    _ = v31
                }
            }
        }
        var v32 int8 = -128
        _ = fi27
        _ = v28
        _ = v32
    }
    fmt.Println(v1, v1)
    var v33 int = 488
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
