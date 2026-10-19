package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() uint32 {
    switch "\n\t" {
    case "a":
        var v4 float32 = ((1e10 / ((-2.25 + 1e10) * (float32(-2.25) + ((0.0 - (1.5 / ((((0.0 + 1.5) - 3.14159) / 1e10) + 1e10))) / 1.0)))) * ((0.0 - ((3.14159 - 1e10) * (3.14159 * 0.1))) / -2.25))
        _ = v4
    case "foo bar":
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 int16 = -32768
            v7 := GenS1{F0: (v6 + (v6 + (v6 % ((398) | 1)))), F1: 4294967295, F2: (v6 - -790)}
            v8 := []int{}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    case "hello":
        var v9 int64 = -9223372036854775808
        _ = v9
    default:
        var v10 uint32 = 2
        _ = v10
    }
    switch "a" {
    case "hello":
        var v11 uint8 = 0
        _ = v11
    case "":
        var v12 uint64 = 18446744073709551615
        _ = v12
    case "a":
        var v13 int32 = -93
        _ = v13
    default:
        var v14 uint16 = 770
        _ = v14
    }
    var v15 float64 = (3.14159 - 3.14159)
    _ = v15
    return 2147483647
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    var v16 string = "\n\t"
    v1 = 1e10
    v1 -= v3
    var v17 uint = 2
    v1 -= (3.14159 * -2.25)
    fmt.Println(v17, v1, v16)
    switch v16 {
    case "":
        var v18 int32 = 2147483647
        _ = v18
    case "\n\t":
        v19 := GenS1{F0: 32767, F1: 18446744073709551615, F2: 0}
        _ = v19
    }
    v3 += ((3.14159 + v3) * (0.0 / ((v3 - float32((0.0 / float64(float32(1e10))))) + v1)))
    if true {
        var v20 int16 = 276
        _ = v20
    }
    fmt.Println(v17)
    var v21 int32 = 584
    v3 = v3
    _ = len(v2)
    v3 = 3.14159
    for fi22 := 0; fi22 < 13; fi22++ {
        switch (v16 + v16) {
        case "hello":
            v1 += 3.14159
        case "a":
            if false {
                v16 = (("foo bar" + "hello") + v16)
            } else {
                for fi23 := 0; fi23 < 11; fi23++ {
                    switch fi22 {
                    case 732:
                        if (0 >= 1000) {
                            v1 += ((v3 / -2.25) / v1)
                            v17 += 42
                            var v24 int32 = 42
                            _ = v24
                        } else {
                            switch v16 {
                            case "\n\t":
                                v1 += float32(0.1)
                            case "hello":
                                v25 := GenS2{F0: 603, F1: v17}
                                _ = v25
                            case "":
                                fmt.Println(v1, v1)
                            }
                            if false {
                                v17 = 255
                                v17 -= (v17 << (((v17 + v17)) & 63))
                                v26 := GenS2{F0: 42, F1: v17}
                                v27 := []int{(fi23 - 311), (fi22 << (((v17 + v17)) & 63)), 205, fi22}
                                var v28 uint8 = 250
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        }
                    case -641:
                        switch fi23 {
                        case 4294967295:
                            fmt.Println(fi22, v17, fi23)
                        case -559:
                            v29 := []int{-2147483648}
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            v16 = v16
                            _ = v29
                        default:
                            fmt.Println(v17, v1)
                        }
                    case 32768:
                        v30 := GenS1{F0: 0, F1: 418, F2: -32768}
                        _ = v30
                    }
                    _ = fi23
                }
                var v31 int32 = (v21 << (((v17 - v17)) & 63))
                _ = v31
            }
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
