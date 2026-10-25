package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int) int8 {
    var v5 bool = (128 == 42)
    p0 = (p0 / (((p0 * 4294967295)) | 1))
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float64 = (3.14159 + ((0.1 / float64(float32(1.5))) * 3.14159))
        fmt.Println(v5, v7, v5)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    v1 := []int{32768, 127}
    var v2 int64 = -1
    var v3 int64 = (v2 ^ (v2 ^ (-(v2))))
    var v4 uint16 = 2
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int8 = -1
        if (true && false) {
            fmt.Println(fi8, v2)
            switch "foo bar" {
            case "foo bar":
                if false {
                    v1 = append(v1, (fi8 + 594))
                }
            case "a":
                v4 = v4
            default:
                v4 = (v4 << ((255) & 63))
            }
        } else {
            if (127 != 236) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 float32 = 0.0
                    _ = fi10
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: 222, F1: (v9 % ((v9) | 1)), F2: true}
                v9 = v9
                if (fi8 <= (fi8 * fi8)) {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 float64 = ((0.0 / 1e10) / (((0.0 + (0.1 * (0.0 + float64(float32((0.1 + -2.25)))))) / 1e10) * float64(3.14159)))
                        _ = fi13
                        _ = v14
                    }
                }
                _ = v12
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                v1 = append(v1, (fi8 & fi8))
                _ = fi15
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS1{F0: 1, F1: v9, F2: ((1.5 / 1.0) > 1e10)}
                if true {
                    switch fi16 {
                    case 32768:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 677:
                        var v18 uint8 = 128
                        _ = v18
                    case 128:
                        var v19 int64 = -795
                        _ = v19
                    default:
                        var v20 uint16 = (^(v4))
                        _ = v20
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if true {
                        var v21 int16 = 32767
                        _ = v21
                    } else {
                        if ((v9 << ((4294967295) & 63)) >= genFn0((fi8 ^ (fi16 << ((255) & 63))))) {
                            var v22 int = fi16
                            if len(v1) > 0 {
                                v1[0] = (fi8 + v22)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v22
                        }
                        var v23 uint8 = 1
                        _ = v23
                    }
                    v24 := GenS1{F0: 2, F1: -128, F2: (2 > 4294967295)}
                    _ = v24
                }
                _ = fi16
                _ = v17
            }
        }
        _ = fi8
        _ = v9
    }
    v2 -= v2
    v25 := GenS2{F0: 683, F1: -128}
    var v26 int8 = 48
    var v27 uint64 = 4294967295
    var v28 uint = 0
    var v29 uint = (v28 * (v28 - 394))
    for fi30 := 0; fi30 < 5; fi30++ {
        v2 = (v2 >> ((v28) & 63))
        _ = fi30
    }
    v2 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
