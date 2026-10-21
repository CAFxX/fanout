package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int32) int16 {
    v6 := GenS1{F0: 86, F1: 4294967295, F2: 2}
    var v7 int = int(p1)
    switch (v7 & v7) {
    case -976:
        var v8 float32 = (((-2.25 / 1.0) / 1.5) * 1.5)
        _ = v8
    case 42:
        var v9 int8 = 127
        _ = v9
    }
    if (false || (v7 > (v7 & (v7 % ((9223372036854775807) | 1))))) {
        fmt.Println(p0, v7, v7)
        v10 := []int{(v7 & v7)}
        _ = v10
    } else {
        if false {
            var v11 uint16 = 2
            v12 := GenS2{F0: "", F1: v7}
            fmt.Println(p1, p0)
            _ = v11
            _ = v12
        } else {
            v13 := []int{(v7 << (((p0 ^ 42)) & 63)), v7, v7, (v7 % (((v7 | (v7 * v7))) | 1))}
            if false {
                var v14 uint = p0
                fmt.Println(v7)
                _ = v14
            } else {
                var v15 int64 = -1
                fmt.Println(p0, v7, v7)
                _ = v15
            }
            var v16 uint8 = 127
            _ = v13
            _ = v16
        }
        fmt.Println(v7)
        var v17 uint32 = 998
        _ = v17
    }
    v18 := []int{v7, v7, (v7 / (((v7 / ((v7) | 1))) | 1)), v7}
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 bool, p2 string) uint64 {
    v19 := GenS1{F0: 990, F1: -1, F2: 2}
    v20 := GenS1{F0: 65535, F1: 127, F2: 0}
    v21 := GenS2{F0: "foo bar", F1: -1}
    v22 := GenS2{F0: ((p2 + p2) + (p2 + p2)), F1: -472}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return 336
}

func main() {
    v1 := []uint32{181}
    var v2 string = "\n\t"
    var v3 bool = true
    var v4 int64 = -929
    var v5 bool = v3
    if ((-1 <= 42) || v5) {
        var v23 uint = 469
        var v24 float64 = 1e10
        _ = len(v1)
        _ = v23
        _ = v24
    }
    var v25 int32 = 42
    for fi26 := 0; fi26 < 4; fi26++ {
        v25 += v25
        _ = fi26
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    if true {
        var v27 uint8 = 2
        fmt.Println(v3, v25)
        for fi28 := 0; fi28 < 10; fi28++ {
            v1 = append(v1, uint32(v27))
            if (42 < 201) {
                fmt.Println(v2, v5)
                for fi29 := 0; fi29 < 18; fi29++ {
                    if (100 <= 100) {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        var v30 float64 = ((-2.25 * (float64(1e10) * -2.25)) * (1.5 / 1.0))
                        v31 := GenS3{F0: 210}
                        _ = v30
                        _ = v31
                    }
                    v2 = (v2 + v2)
                    if false {
                        v25 += (^(v25))
                        var v32 int64 = v4
                        _ = v32
                    } else {
                        switch 32767 {
                        case -2147483648:
                            fmt.Println(v3)
                        case 42:
                            v1 = append(v1, 154)
                        case 128:
                            var v33 int = (^(fi29))
                            _ = v33
                        default:
                            if (((((v2 + (v2 + (v2 + v2))) + v2) + "foo bar") + v2) >= "\n\t") {
                                v3 = v3
                                var v34 uint = 255
                                v27 = v27
                                _ = v34
                            } else {
                                v35 := GenS2{F0: "hello", F1: fi28}
                                _ = v35
                            }
                        }
                        v2 = ((((("\n\t" + v2) + v2) + (v2 + ((v2 + v2) + v2))) + "a") + "a")
                    }
                    v27 = 255
                    _ = fi29
                }
            }
            _ = fi28
        }
        _ = v27
    }
    var v36 uint = 4294967295
    if v5 {
        var v37 uint = v36
        _ = v37
    }
    var v38 int = 9223372036854775807
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
