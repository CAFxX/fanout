package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 int32) uint {
    v7 := GenS1{F0: 18446744073709551615}
    v8 := GenS2{F0: "hello", F1: p0, F2: 196}
    p0 -= p0
    p0 = 1
    _ = p0
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    var v3 int64 = 209
    var v4 int = -748
    var v5 uint32 = 42
    v6 := make(map[string]uint32)
    if false {
        var v9 int8 = -128
        _ = v9
    } else {
        switch "" {
        case "":
            var v10 int16 = -1
            _ = v10
        }
        if (127 != 1) {
            if false {
                fmt.Println(v3, v4, v3)
                if (!false) {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        switch (v4 << ((65535) & 63)) {
                        case 4294967295:
                            if (false && true) {
                                v2[v4] = (235 <= 127)
                                v4 -= (v4 << ((genFn0(700)) & 63))
                            }
                        case -2147483648:
                            var v12 uint8 = 42
                            _ = v12
                        case 42:
                            var v13 bool = (float32(0.0) <= -2.25)
                            _ = v13
                        }
                        _ = fi11
                    }
                }
            } else {
                if ((-52 != 1) && false) {
                    v14 := GenS1{F0: 1}
                    for fi15 := 0; fi15 < 8; fi15++ {
                        if true {
                            for fi16 := 0; fi16 < 16; fi16++ {
                                v17 := GenS3{F0: 2, F1: 255, F2: (-(fi15))}
                                _ = fi16
                                _ = v17
                            }
                        }
                        if true {
                            v4 = v4
                            fmt.Println(fi15)
                        }
                        _ = fi15
                    }
                    var v18 int64 = v3
                    _ = v14
                    _ = v18
                } else {
                    v5 -= (v5 << ((584) & 63))
                    v19 := v6["\n\t"]
                    v20 := GenS3{F0: 42, F1: 2, F2: v4}
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = 0
                _ = v21
            }
            var v22 uint64 = 4294967295
            v3 = (v3 * v3)
            _ = v22
        } else {
            var v23 int64 = (v3 * v3)
            v24 := GenS3{F0: 2, F1: 32767, F2: (-(v4))}
            var v25 int8 = 1
            _ = v23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v3, v5, v4)
    if (false && false) {
        for fi26 := 0; fi26 < 13; fi26++ {
            _ = len(v2)
            v4 -= v4
            v5 -= (v5 >> ((2) & 63))
            _ = fi26
        }
    }
    _ = len(v6)
    v3 += (v3 | 42)
    v27 := GenS4{F0: 32767}
    v3 = (v3 | v3)
    var v28 string = (("\n\t" + "a") + (("\n\t" + (("hello" + "\n\t") + "hello")) + ""))
    fmt.Println(v5)
    v3 = (^(v3))
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(v28, v3, fi29)
        _ = fi29
    }
    var v30 float64 = 0.1
    _ = len(v1)
    v31 := GenS5{F0: uint64(v5), F1: int64(v4), F2: 42}
    v30 = 0.1
    fmt.Println(v4, v3, v30)
    _ = len(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
