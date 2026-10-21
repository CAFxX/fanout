package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 string, p1 int8) uint32 {
    fmt.Println(p0)
    v4 := GenS1{F0: 42}
    switch ((p0 + "foo bar") + p0) {
    case "hello":
        v5 := GenS1{F0: 1}
        _ = v5
    case "":
        v6 := []int{32768, 598}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 0
    v3 := make(map[string]bool)
    var v7 int64 = 9223372036854775807
    fmt.Println(v7, v7)
    switch -1 {
    case 32767:
        var v8 uint8 = 127
        _ = v8
    case -1:
        v9 := GenS1{F0: 222}
        _ = v9
    case 32768:
        v10 := []int{-9223372036854775808, 4294967295, 1, 128}
        _ = v10
    default:
        v2 = (v2 % ((v2) | 1))
    }
    v11 := []int{128, -866, -1}
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            var v13 int64 = (^(v7))
            v14 := GenS1{F0: 0}
            _ = v13
            _ = v14
        }
        _ = len(v1)
        switch ("" + "\n\t") {
        case "a":
            for fi15 := 0; fi15 < 18; fi15++ {
                if false {
                    v16 := GenS1{F0: 2}
                    v3[""] = true
                    _ = v16
                }
                fmt.Println(v2)
                v7 += 0
                if false {
                    switch fi15 {
                    case 486:
                        v7 += -1
                    default:
                        v2 += (v2 / ((2) | 1))
                    }
                    if true {
                        v7 = v7
                        if (((((v2 * 878) < 255) && true) && (733 != 162)) || true) {
                            v7 += v7
                        } else {
                            var v17 int = (^(fi12))
                            v18 := GenS1{F0: 352}
                            _ = v17
                            _ = v18
                        }
                        var v19 int8 = 100
                        v3["a"] = false
                        _ = v19
                    } else {
                        delete(v1, "a")
                        var v20 int = (fi12 - (fi12 ^ (fi12 - (-(fi15)))))
                        _ = v20
                    }
                    v2 = (^(v2))
                } else {
                    var v21 float64 = ((0.1 + 0.0) * 0.1)
                    v21 -= (((((v21 * (v21 / 0.1)) - v21) + (1.5 / (float64(1e10) / v21))) / float64(-2.25)) - (v21 * (((0.0 - 1e10) * 1e10) + (float64(float32((((v21 + v21) / (v21 + 0.0)) - (float64(((float32(float64((0.0 - 3.14159))) + (1e10 + 3.14159)) * (((0.1 + -2.25) * -2.25) + 1.5))) - v21)))) - v21))))
                    _ = v21
                }
                v2 += 255
                _ = fi15
            }
        case "\n\t":
            if true {
                v2 = 255
                if len(v11) > 0 {
                    v11[0] = (fi12 * (^(fi12)))
                }
                var v22 int = (^(fi12))
                _ = v22
            } else {
                var v23 float32 = (float32(1.5) / 0.1)
                for fi24 := 0; fi24 < 15; fi24++ {
                    v2 = v2
                    v25 := GenS2{F0: 0, F1: 240}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
        default:
            var v26 float32 = 1e10
            _ = v26
        }
        _ = fi12
    }
    v27 := GenS2{F0: -410, F1: 38}
    fmt.Println(v2, v2)
    _ = len(v1)
    v2 = (v2 >> (((^(v2))) & 63))
    v11 = append(v11, -1)
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v2, v2)
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 string = "foo bar"
            v31 := GenS1{F0: 21}
            if (false || (!(-2.25 >= (0.0 - (-2.25 * 3.14159))))) {
                var v32 uint32 = genFn0(v30, 0)
                _ = v32
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    var v33 uint8 = 128
    v7 = (v7 | -702)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v7
}
