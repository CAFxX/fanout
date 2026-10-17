package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0(p0 uint16) uint8 {
    switch "\n\t" {
    case "hello":
        switch ("\n\t" + ("\n\t" + "foo bar")) {
        case "a":
            p0 = (^(p0))
        case "\n\t":
            fmt.Println(p0)
        }
    case "a":
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: 0}
            _ = fi7
            _ = v8
        }
    case "foo bar":
        p0 = (^(p0))
    default:
        var v9 uint64 = 1
        _ = v9
    }
    v10 := GenS1{F0: 2147483647}
    v11 := []int{0, 9223372036854775807}
    p0 = p0
    _ = p0
    _ = v10
    _ = v11
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = float64(3.14159)
    v3 := []int{0, 128, 2147483647}
    var v4 uint16 = 65535
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    var v12 uint32 = 2147483647
    var v13 uint64 = 139
    v13 = (v13 ^ 4294967295)
    if (!true) {
        v4 -= v4
        v2 = (3.14159 + (v2 * v2))
        for fi14 := 0; fi14 < 7; fi14++ {
            if (0.0 >= -2.25) {
                fmt.Println(fi14, v2, v2)
                v15 := GenS1{F0: fi14}
                v12 += v12
                _ = v15
            } else {
                switch fi14 {
                case 1:
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 float64 = float64((((((((1e10 * (1e10 + (1.5 + 3.14159))) + (float32((v2 - (((v2 / (0.0 / (float64(0.0) + float64(float32(v2))))) / v2) * 1.5))) + float32(v2))) / 1.0) * float32(v2)) / (1.5 + float32(0.1))) - (1e10 + 1e10)) * (0.1 / (1.5 - 1e10))))
                        _ = fi16
                        _ = v17
                    }
                case 2147483647:
                    var v18 int8 = 0
                    _ = v18
                default:
                    switch fi14 {
                    case 715:
                        v19 := GenS1{F0: fi14}
                        _ = v19
                    case 348:
                        switch fi14 {
                        case 1:
                            if true {
                                v20 := GenS1{F0: fi14}
                                _ = v20
                            }
                        case 345:
                            v2 = (v2 - (float64(1.5) / 1.0))
                        default:
                            _ = len(v1)
                        }
                    default:
                        for fi21 := 0; fi21 < 3; fi21++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi21
                        }
                    }
                }
                _ = len(v1)
                var v22 int64 = 0
                _ = v22
            }
            _ = fi14
        }
    } else {
        v6[32768] = "hello"
        var v23 int8 = -128
        var v24 int64 = -9223372036854775808
        _ = v23
        _ = v24
    }
    if false {
        v25 := GenS2{F0: 32767, F1: 2147483647}
        _ = v25
    } else {
        if false {
            fmt.Println(v2, v13)
            switch "\n\t" {
            case "foo bar":
                var v26 string = "a"
                _ = v26
            case "hello":
                fmt.Println(v12)
            default:
                var v27 float32 = 1e10
                _ = v27
            }
            v2 -= -2.25
        } else {
            v4 += (v4 & uint16(v12))
            if true {
                v3 = append(v3, 42)
            } else {
                var v28 uint32 = v12
                _ = v28
            }
        }
        v5[(^(v12))] = 32767
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = fi29
            }
        }
        v30 := GenS3{F0: 42, F1: (v13 * 1), F2: 0}
        v5[4294967295] = fi29
        _ = fi29
        _ = v30
    }
    if (-10 < -128) {
        var v31 uint32 = (v12 & (^(v12)))
        _ = v31
    } else {
        fmt.Println(v4)
        var v32 int = 127
        _ = v32
    }
    v12 = uint32(v4)
    fmt.Println(v4)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
