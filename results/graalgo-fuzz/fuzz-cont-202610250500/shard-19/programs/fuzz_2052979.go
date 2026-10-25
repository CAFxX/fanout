package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 uint8) uint32 {
    if true {
        v5 := []int{127, 32767}
        _ = v5
    }
    var v6 float32 = -2.25
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    var v3 uint32 = 1
    v4 := make(map[int]int)
    var v7 int16 = 681
    v3 = v3
    var v8 int64 = -359
    v2 = v2
    if true {
        v9 := []int{4294967295}
        fmt.Println(v3)
        v10 := GenS1{F0: 993, F1: genFn0(191), F2: v3}
        _ = v10
        _ = v9
    } else {
        fmt.Println(v2, v7, v3)
        fmt.Println(v7, v8)
    }
    var v11 float32 = -2.25
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint = 4294967295
        fmt.Println(v8)
        fmt.Println(v2, fi12)
        _ = fi12
        _ = v13
    }
    if ((false || ((v2 % ((v2) | 1)) != v2)) || false) {
        fmt.Println(v11, v2, v11)
        v3 = (v3 * (v3 + (^(v3))))
        for fi14 := 0; fi14 < 4; fi14++ {
            v3 = v3
            v15 := []int{-301, 32767}
            _ = len(v15)
            _ = fi14
            _ = v15
        }
    } else {
        v16 := []int{83, 127, -9223372036854775808, 0}
        if ((false && (624 < 4294967295)) || false) {
            v17 := GenS1{F0: 1, F1: v3, F2: genFn0(42)}
            _ = v17
        } else {
            if true {
                var v18 int16 = (v7 << ((65535) & 63))
                var v19 uint8 = 42
                var v20 float32 = (-2.25 / v11)
                _ = v18
                _ = v19
                _ = v20
            }
            if len(v16) > 0 {
                v16[0] = 32767
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                v7 -= v7
                for fi22 := 0; fi22 < 12; fi22++ {
                    v11 = v11
                    if true {
                        var v23 int = fi21
                        fmt.Println(v11, v11)
                        _ = v23
                    } else {
                        if true {
                            v8 += (v8 / (((v8 % ((42) | 1))) | 1))
                            var v24 int64 = 42
                            fmt.Println(v8)
                            _ = v24
                        } else {
                            if len(v16) > 0 {
                                v16[0] = (fi21 % ((fi21) | 1))
                            }
                            if (-128 == 42) {
                                var v25 uint64 = (v2 | v2)
                                _ = v25
                            } else {
                                v8 -= v8
                            }
                            v2 += (v2 * (v2 & v2))
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v8 -= v8
                        default:
                            v1[fi22] = ("a" + "")
                        }
                    }
                    v26 := v1[fi22]
                    _ = fi22
                    _ = v26
                }
                _ = fi21
            }
        }
        var v27 bool = (!false)
        _ = v16
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            v2 += (v2 | v2)
            var v30 uint64 = (v2 ^ v2)
            _ = fi29
            _ = v30
        }
        var v31 uint32 = (v3 * (v3 >> ((65535) & 63)))
        _ = fi28
        _ = v31
    }
    var v32 uint16 = 856
    for fi33 := 0; fi33 < 17; fi33++ {
        fmt.Println(v7, v2, v11)
        v34 := []int{(fi33 & int(v2))}
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v7
    _ = v8
}
