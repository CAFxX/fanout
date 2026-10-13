package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint64; F2 int }
func genFn0() int8 {
    if (2 > 2) {
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: 742}
            v7 := GenS1{F0: 320}
            _ = fi5
            _ = v6
            _ = v7
        }
        if true {
            var v8 uint8 = 127
            var v9 float32 = 0.1
            var v10 uint16 = 372
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := []int{-2147483648, -676}
            var v12 string = "foo bar"
            _ = v11
            _ = v12
        }
    } else {
        var v13 bool = true
        var v14 int64 = 0
        _ = v13
        _ = v14
    }
    var v15 float64 = (float64(0.0) + float64(3.14159))
    switch -2147483648 {
    case 32767:
        switch "" {
        case "foo bar":
            fmt.Println(v15)
        default:
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: 532}
                if false {
                    if (9223372036854775807 > -1) {
                        if (1 >= 255) {
                            if false {
                                var v18 int64 = -1
                                v15 += 0.1
                                var v19 int = (-(fi16))
                                _ = v18
                                _ = v19
                            } else {
                                if (v15 >= ((v15 - v15) / v15)) {
                                    v20 := []int{(^(fi16)), fi16, (fi16 / ((fi16) | 1))}
                                    var v21 int = 42
                                    _ = v20
                                    _ = v21
                                }
                            }
                            if false {
                                v15 = ((v15 + v15) * v15)
                                v22 := GenS1{F0: 500}
                                if (!(int32(fi16) >= -2147483648)) {
                                    v15 += (((float64(1.5) + v15) + 1.5) + (v15 + float64(3.14159)))
                                    v23 := GenS1{F0: 4294967295}
                                    _ = v23
                                }
                                _ = v22
                            }
                            v15 -= v15
                        } else {
                            if true {
                                if true {
                                    v15 = (0.0 - v15)
                                    v24 := []int{1}
                                    _ = v24
                                } else {
                                    var v25 float32 = (3.14159 + (float32(v15) * 3.14159))
                                    _ = v25
                                }
                            }
                            v26 := GenS2{F0: (("foo bar" + "foo bar") + "\n\t"), F1: 2, F2: fi16}
                            v27 := GenS1{F0: 42}
                            _ = v26
                            _ = v27
                        }
                    }
                }
                _ = fi16
                _ = v17
            }
        }
    case -2147483648:
        v28 := GenS2{F0: ("hello" + "\n\t"), F1: 814, F2: 9223372036854775807}
        _ = v28
    default:
        fmt.Println(v15)
    }
    v29 := []int{-696, 127, 32767}
    _ = v15
    _ = v29
    return 100
}

func main() {
    var v1 int = 506
    var v2 int = 179
    var v3 uint = 65535
    v4 := make(map[uint32]string)
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 uint16 = 2
        _ = fi30
        _ = v31
    }
    v2 = v1
    var v32 int8 = genFn0()
    var v33 int8 = (v32 ^ 1)
    var v34 int32 = 458
    var v35 float64 = float64(0.0)
    fmt.Println(v34)
    v34 = 795
    v36 := []int{(v2 >> (((v3 * v3)) & 63)), v2, v2, v2}
    var v37 uint64 = 42
    var v38 uint8 = 1
    var v39 float64 = v35
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
