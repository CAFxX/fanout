package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }

func main() {
    var v1 int8 = 127
    var v2 int8 = 64
    var v3 int32 = -243
    v4 := make(map[uint32]bool)
    if false {
        v3 += 2147483647
    }
    fmt.Println(v2, v2, v1)
    v5 := []int{127, 127, 127}
    v6 := v4[96]
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 18446744073709551615
        var v9 uint32 = 2147483647
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 -= v3
    fmt.Println(v6, v3)
    var v10 uint16 = 483
    v10 -= 894
    if (v10 == (^(v10))) {
        switch -557 {
        case 2147483647:
            v4[2] = v6
        default:
            if true {
                v11 := GenS1{F0: v1, F1: 1, F2: 0}
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := v4[520]
                    if true {
                        if v6 {
                            switch fi12 {
                            case -1:
                                v6 = false
                            default:
                                _ = len(v5)
                            }
                        } else {
                            var v14 int = (fi12 & (fi12 & (fi12 << ((42) & 63))))
                            _ = v14
                        }
                        if v6 {
                            for fi15 := 0; fi15 < 9; fi15++ {
                                v16 := GenS1{F0: v2, F1: 783, F2: 18446744073709551615}
                                var v17 float32 = 1e10
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        } else {
                            if (v13 && ((v10 % ((v10) | 1)) > (v10 >> ((1) & 63)))) {
                                if (0.0 > ((0.0 + (0.1 + (3.14159 + 1e10))) / 1.0)) {
                                    var v18 string = "hello"
                                    _ = v18
                                }
                                var v19 uint16 = (v10 % (((v10 + 2)) | 1))
                                _ = v19
                            } else {
                                var v20 int8 = int8(fi12)
                                _ = v20
                            }
                        }
                    } else {
                        switch (^(fi12)) {
                        case -989:
                            if (v2 != v2) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                fmt.Println(v1, v10)
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    var v22 int16 = -402
                                    v23 := GenS1{F0: 0, F1: (v22 * v22), F2: 539}
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            }
                        case 450:
                            var v24 float32 = -2.25
                            _ = v24
                        default:
                            v25 := v4[196]
                            _ = v25
                        }
                        _ = len(v4)
                        var v26 int8 = v1
                        _ = v26
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 int32 = (v3 ^ v3)
            v29 := v4[95]
            _ = len(v5)
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
