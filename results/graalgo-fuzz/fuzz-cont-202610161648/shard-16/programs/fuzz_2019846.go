package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8) int8 {
    p0 -= p0
    var v4 int16 = 42
    v5 := []int{-9223372036854775808, 4294967295, 1, 32768}
    v6 := GenS1{F0: (false || false)}
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int16 = -538
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v1 = v1
    var v7 string = "a"
    if (!(1 > 164)) {
        var v8 uint32 = 42
        v9 := GenS1{F0: false}
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 100, F1: 0}
    v11 := []int{}
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(fi12, fi12, v7)
        _ = fi12
    }
    v13 := GenS3{F0: 127}
    v14 := GenS4{F0: 1, F1: 65535, F2: false}
    var v15 int32 = 42
    delete(v3, 128)
    var v16 float64 = ((((0.1 / 1.0) - 1.5) + 1.5) - (((((3.14159 * -2.25) / (((1e10 - -2.25) / (1.5 - 0.0)) + 1e10)) - 1e10) / (3.14159 + ((float64((-2.25 + 0.1)) - -2.25) + 1e10))) / ((-2.25 - 3.14159) + (1e10 * (3.14159 + 0.0)))))
    _ = len(v11)
    var v17 uint8 = 1
    if (!true) {
        if len(v11) > 0 {
            v11[0] = -868
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint32 = 65535
                if (uint32(v17) == (v20 >> ((18446744073709551615) & 63))) {
                    if false {
                        _ = len(v11)
                    } else {
                        if false {
                            switch v7 {
                            case "\n\t":
                                var v21 float32 = ((-2.25 + 0.1) / 0.1)
                                _ = v21
                            case "foo bar":
                                delete(v2, 42)
                            default:
                                v22 := v2[536]
                                _ = v22
                            }
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        } else {
                            var v23 uint8 = v17
                            _ = v23
                        }
                    }
                    if ((^(v15)) > (v15 << ((42) & 63))) {
                        var v24 uint8 = 42
                        v25 := GenS1{F0: false}
                        if len(v11) > 0 {
                            v11[0] = (fi19 % ((1) | 1))
                        }
                        _ = v24
                        _ = v25
                    } else {
                        v16 = (v16 - v16)
                    }
                }
                _ = fi19
                _ = v20
            }
            var v26 uint16 = 255
            _ = fi18
            _ = v26
        }
    }
    var v27 int16 = v1
    v28 := GenS5{F0: 2147483647}
    v29 := GenS6{F0: 42, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
