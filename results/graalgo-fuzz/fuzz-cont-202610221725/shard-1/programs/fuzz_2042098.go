package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int8, p1 int) int16 {
    var v4 uint8 = 113
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{fi5}
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 bool = true
            switch "a" {
            case "a":
                _ = len(v6)
            }
            v4 -= 128
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: 401, F1: 2, F2: 33}
        v11 := []int{(fi9 / ((fi9) | 1)), (fi9 ^ -2147483648), (p1 + 373), fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    var v12 float64 = 1.5
    var v13 uint64 = 18446744073709551615
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float32 = 1.5
        _ = fi14
        _ = v15
    }
    v3 += v3
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            v3 += 2
            v18 := GenS1{F0: 0, F1: 2, F2: int32(fi16)}
            if (1 <= 124) {
                var v19 uint16 = 0
                if false {
                    var v20 int64 = 1
                    if (!((v13 / (((^(v13))) | 1)) == v13)) {
                        v12 = 0.0
                        if (!false) {
                            if true {
                                delete(v2, 2147483648)
                                if ((!true) && false) {
                                    v21 := GenS1{F0: 2, F1: v13, F2: 1}
                                    var v22 float32 = ((0.0 / 3.14159) - 1e10)
                                    _ = v21
                                    _ = v22
                                } else {
                                    switch fi17 {
                                    case 127:
                                        var v23 int = 828
                                        _ = v23
                                    }
                                }
                            }
                            v19 = (^(v19))
                            var v24 int64 = 42
                            _ = v24
                        }
                        if true {
                            v25 := []int{-720, fi17}
                            _ = v25
                        }
                    } else {
                        v26 := v2[2147483648]
                        _ = v26
                    }
                    fmt.Println(fi16, v12, v1)
                    var v27 int64 = 651
                    var v28 float32 = -2.25
                    _ = v20
                    _ = v27
                    _ = v28
                }
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        if true {
            v29 := []int{(fi16 + (fi16 / (((fi16 >> ((0) & 63))) | 1))), (fi16 << ((42) & 63)), fi16}
            _ = v29
        } else {
            var v30 uint64 = 0
            var v31 float32 = (0.0 - 3.14159)
            v12 += (((((v12 + float64(((v31 * 0.1) - (v31 / 1.5)))) - ((float64(1e10) * (v12 / v12)) + -2.25)) - 0.1) / 0.1) * float64(0.0))
            _ = v30
            _ = v31
        }
        var v32 uint16 = 255
        _ = fi16
        _ = v32
    }
    v1 = v1
    v33 := GenS1{F0: (^(v13)), F1: (v13 / (((^(v13))) | 1)), F2: 603}
    if ((false || false) || false) {
        v1 = v1
    } else {
        v3 -= v3
        fmt.Println(v1)
    }
    v34 := []int{2147483647, -9223372036854775808, 128}
    var v35 uint = 1
    v36 := GenS2{F0: uint(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
