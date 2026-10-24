package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
func genFn0(p0 uint32) int32 {
    v6 := GenS1{F0: 521, F1: true, F2: 1}
    for fi7 := 0; fi7 < 13; fi7++ {
        if (0 < 9223372036854775807) {
            v8 := GenS1{F0: -9223372036854775808, F1: true, F2: 4294967295}
            var v9 uint64 = 4294967295
            _ = v8
            _ = v9
        }
        var v10 float64 = 0.1
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int64 = -9223372036854775808
            if (!(!true)) {
                var v13 uint16 = 255
                v14 := []int{}
                _ = v13
                _ = v14
            } else {
                p0 = 1
                if true {
                    v15 := GenS1{F0: v12, F1: true, F2: 2}
                    _ = v15
                }
            }
            if true {
                if false {
                    v16 := []int{fi7, fi7, (-(fi7))}
                    p0 = (^(p0))
                    _ = v16
                } else {
                    switch "a" {
                    case "foo bar":
                        if true {
                            var v17 int32 = -72
                            if true {
                                v17 += 0
                                v18 := []int{0, (fi7 ^ -976), 738, 32768}
                                fmt.Println(fi11)
                                _ = v18
                            } else {
                                fmt.Println(v10)
                            }
                            _ = v17
                        }
                    default:
                        p0 -= (p0 % (((^(p0))) | 1))
                    }
                    var v19 bool = true
                    _ = v19
                }
            } else {
                v20 := GenS1{F0: v12, F1: false, F2: 2}
                if true {
                    var v21 float32 = 0.0
                    v22 := []int{(fi7 + 128), (^(fi7)), (-(fi7))}
                    _ = v21
                    _ = v22
                } else {
                    fmt.Println(v10)
                    v23 := GenS1{F0: -9223372036854775808, F1: false, F2: 806}
                    if true {
                        var v24 uint8 = 0
                        fmt.Println(fi7, v10)
                        if false {
                            var v25 float32 = -2.25
                            fmt.Println(fi7, v24)
                            _ = v25
                        }
                        _ = v24
                    } else {
                        v10 = (float64(1e10) + 3.14159)
                    }
                    fmt.Println(v10, p0, v12)
                    _ = v23
                }
                v10 += ((v10 / 1.0) * (v10 - float64(float32(float64(0.1)))))
                _ = v20
            }
            _ = fi11
            _ = v12
        }
        _ = fi7
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 int32 = -1
    var v2 int = 4294967295
    v3 := make(map[uint32]string)
    var v4 int64 = 42
    var v5 uint64 = 735
    v26 := GenS1{F0: (v4 << ((8) & 63)), F1: true, F2: v5}
    v27 := []int{42, (v2 - v2), (v2 & (v2 % (((v2 - 32768)) | 1))), v2}
    var v28 uint8 = 92
    if true {
        v5 -= v5
    } else {
        var v29 int = (v2 * -676)
        _ = v29
    }
    fmt.Println(v5)
    v4 = (-(v4))
    v4 += (v4 * v4)
    v5 = 436
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
