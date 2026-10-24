package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 int64) bool {
    switch "hello" {
    case "hello":
        switch -533 {
        case 4294967295:
            p0 = 1
        case -622:
            v7 := GenS1{F0: -1}
            _ = v7
        }
    case "":
        v8 := GenS1{F0: -9223372036854775808}
        _ = v8
    }
    v9 := GenS1{F0: 0}
    p0 = p0
    _ = p0
    _ = v9
    return (false || (!true))
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 int = -520
    var v4 float64 = 1e10
    v5 := make(map[uint32]int)
    v6 := make(map[int]uint32)
    if (42 <= 2) {
        for fi10 := 0; fi10 < 8; fi10++ {
            v3 = (fi10 % ((-298) | 1))
            _ = fi10
        }
        if false {
            v4 = ((v4 + v4) / v4)
            fmt.Println(v3, v3)
            fmt.Println(v4)
        } else {
            var v11 bool = false
            var v12 float32 = ((0.1 * ((((1e10 * 1e10) * 0.1) - (1.5 / -2.25)) + (((3.14159 / -2.25) - (float32(v4) - -2.25)) / (1e10 * (1e10 / (1e10 - float32(v4))))))) / -2.25)
            var v13 int16 = -32768
            _ = v11
            _ = v12
            _ = v13
        }
    } else {
        if false {
            if (true || false) {
                if false {
                    var v14 int8 = 0
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v16 := GenS1{F0: (fi15 * fi15)}
                        if false {
                            v14 -= v14
                        } else {
                            switch v3 {
                            case 127:
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    v4 -= (-2.25 / v4)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v4, v14)
                                    _ = fi17
                                }
                            }
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    v4 += v4
                                    v1 = append(v1, fi15)
                                }
                            case "hello":
                                switch fi15 {
                                case -2147483648:
                                    fmt.Println(fi15)
                                case -878:
                                    v4 += v4
                                default:
                                    var v18 float64 = (3.14159 / ((v4 / -2.25) / v4))
                                    _ = v18
                                }
                            default:
                                if true {
                                    if true {
                                        for fi19 := 0; fi19 < 14; fi19++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = fi19
                                        }
                                        var v20 int32 = 2147483647
                                        var v21 int64 = 0
                                        _ = v20
                                        _ = v21
                                    }
                                    v22 := GenS1{F0: v3}
                                    v14 = (v14 | (^(v14)))
                                    _ = v22
                                }
                            }
                        }
                        _ = fi15
                        _ = v16
                    }
                    delete(v2, 928)
                    _ = v14
                }
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                fmt.Println(v4, v3, v4)
                var v24 uint8 = 127
                v24 += 42
                _ = fi23
                _ = v24
            }
        } else {
            v4 = v4
            v25 := GenS1{F0: 526}
            _ = v25
        }
        v4 = (3.14159 - (0.1 * v4))
    }
    v4 = (v4 / v4)
    v3 -= 32767
    v26 := v6[(^(v3))]
    v26 = v26
    v27 := GenS2{F0: -1, F1: -84, F2: 195}
    switch -430 {
    case 42:
        fmt.Println(v3)
    case 1:
        var v28 int8 = 100
        _ = v28
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
