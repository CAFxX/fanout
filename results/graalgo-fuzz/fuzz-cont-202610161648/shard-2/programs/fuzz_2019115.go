package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (1e10 * ((float64((3.14159 + (0.0 * (float32(-2.25) * 0.0)))) + 0.1) / 0.1))
    var v3 int32 = -2147483648
    var v4 uint8 = 81
    v5 := []int{824}
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 0
        _ = fi6
        _ = v7
    }
    if true {
        if (42 > 42) {
            var v8 uint32 = 67
            v9 := GenS1{F0: 1, F1: 0, F2: (^(v4))}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint8 = 255
                switch fi10 {
                case 576:
                    v12 := GenS1{F0: 65535, F1: 42, F2: 75}
                    _ = v12
                default:
                    if true {
                        var v13 string = "hello"
                        _ = v13
                    }
                }
                if false {
                    v11 -= (v11 * v4)
                    switch (fi10 >> ((773) & 63)) {
                    case -1:
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v4 = (v4 | v4)
                            v3 = (v3 + v3)
                            _ = fi14
                        }
                    case 42:
                        var v15 bool = true
                        _ = v15
                    case -2147483648:
                        var v16 float64 = float64((1.5 / ((-2.25 / 0.1) * float32(1.5))))
                        _ = v16
                    default:
                        if false {
                            v17 := GenS1{F0: 65535, F1: -1, F2: 0}
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 uint8 = 42
                                v3 = -2147483648
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        } else {
                            v20 := GenS1{F0: 0, F1: 0, F2: (v11 - (v11 & (v11 >> ((2) & 63))))}
                            v21 := v1["hello"]
                            var v22 int16 = 0
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                }
                _ = fi10
                _ = v11
            }
        }
    }
    fmt.Println(v2, v3)
    if len(v5) > 0 {
        v5[0] = 42
    }
    v23 := GenS1{F0: 65535, F1: 32767, F2: 127}
    for fi24 := 0; fi24 < 2; fi24++ {
        v4 = 1
        _ = fi24
    }
    v25 := GenS1{F0: uint16(v4), F1: -32768, F2: (v4 * v4)}
    var v26 uint32 = 899
    v27 := GenS2{F0: 42, F1: 532, F2: 42}
    if false {
        v28 := GenS3{F0: 853, F1: 1, F2: 65535}
        if false {
            for fi29 := 0; fi29 < 16; fi29++ {
                v2 = (0.0 + 3.14159)
                _ = fi29
            }
            var v30 uint16 = 411
            v31 := GenS4{F0: v26, F1: v30, F2: 87}
            _ = v30
            _ = v31
        }
        _ = v28
    }
    var v32 float32 = 0.0
    v33 := GenS5{F0: v3}
    v2 += v2
    v34 := GenS6{F0: uint(v26)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
