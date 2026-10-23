package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 string; F2 bool }
type GenS4 struct { F0 bool; F1 uint64 }
func genFn0(p0 int16) uint64 {
    fmt.Println(p0)
    v4 := GenS1{F0: 0}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint8 = 2
        fmt.Println(v6, p0)
        p0 += (-(p0))
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 364
}

func main() {
    var v1 int16 = 42
    var v2 float64 = float64(((float32(3.14159) - 3.14159) + 0.0))
    var v3 string = "a"
    v7 := []int{}
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 10; fi9++ {
        _ = len(v7)
        _ = fi9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v3 = v3
    var v10 int64 = 1
    switch v3 {
    case "a":
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(fi11, v1, v1)
            _ = fi11
        }
    default:
        var v12 uint32 = 67
        _ = v12
    }
    v7 = append(v7, 32767)
    if true {
        if false {
            v13 := GenS2{F0: 92}
            var v14 float64 = v2
            v3 = (v3 + v3)
            _ = v13
            _ = v14
        } else {
            v10 += v10
            var v15 uint8 = 2
            _ = v15
        }
        v10 -= (v10 - (^(v10)))
        if (987 >= 0) {
            var v16 int32 = 42
            var v17 int = -291
            if false {
                v10 = v10
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            } else {
                var v18 float32 = 1e10
                switch "" {
                case "":
                    switch v3 {
                    case "a":
                        fmt.Println(v1, v18)
                    case "":
                        v17 -= 4294967295
                    }
                case "hello":
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 uint32 = 42
                        var v21 int8 = 35
                        if true {
                            if (v21 > -18) {
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v21 += int8(v20)
                                    v17 += (fi22 / (((-(fi22))) | 1))
                                    v23 := GenS2{F0: (v21 / ((1) | 1))}
                                    fmt.Println(v2)
                                    _ = fi22
                                    _ = v23
                                }
                                fmt.Println(v20)
                            }
                            if true {
                                v20 += 2147483647
                                v17 = -407
                                v17 = fi19
                            } else {
                                fmt.Println(v3, v20, v16)
                                if len(v7) > 0 {
                                    v7[0] = -1
                                }
                                v7 = append(v7, (^(fi19)))
                            }
                            for fi24 := 0; fi24 < 6; fi24++ {
                                if true {
                                    v25 := GenS1{F0: v16}
                                    _ = v25
                                }
                                var v26 uint32 = v20
                                _ = len(v7)
                                _ = fi24
                                _ = v26
                            }
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                var v27 bool = false
                _ = v18
                _ = v27
            }
            _ = v16
            _ = v17
        }
    }
    var v28 float64 = (((0.0 / (v2 - (float64(1.5) * ((float64(float32(float64(3.14159))) / (0.0 - 0.1)) * ((((v2 - ((v2 * v2) - (1e10 * (((((float64(1e10) / 1.5) * (3.14159 / v2)) + (v2 * ((v2 + float64(1.5)) * v2))) / v2) / 1e10)))) + ((v2 + (v2 / ((v2 * (float64(1e10) - (1.5 - v2))) / (v2 / -2.25)))) + 0.0)) - (v2 / float64((1.5 + ((0.1 / -2.25) * 3.14159))))) * float64(1e10)))))) + float64(1.5)) - v2)
    if len(v7) > 0 {
        v7[0] = 622
    }
    switch v3 {
    case "foo bar":
        v29 := GenS3{F0: 1, F1: "\n\t", F2: false}
        _ = v29
    }
    fmt.Println(v10, v28)
    if (!false) {
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS4{F0: (1.5 > v2), F1: 562}
            _ = fi30
            _ = v31
        }
        v3 = v3
        var v32 uint16 = 436
        _ = v32
    } else {
        if ((2 == 42) && false) {
            v3 = ("\n\t" + v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v7
    _ = v8
}
