package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string }

func main() {
    v1 := []uint32{2147483648, 2147483647}
    var v2 int = 42
    v3 := []int{(v2 % ((v2) | 1)), 42}
    v4 := make(map[string]uint32)
    var v5 float32 = 0.1
    v6 := make(map[int]int)
    v2 += (v2 % (((v2 + v2)) | 1))
    for fi7 := 0; fi7 < 15; fi7++ {
        if len(v3) > 0 {
            v3[0] = (fi7 * -2147483648)
        }
        fmt.Println(v5)
        _ = fi7
    }
    v2 = (v2 >> ((721) & 63))
    for fi8 := 0; fi8 < 2; fi8++ {
        if true {
            switch (v2 + (v2 + 282)) {
            case -9223372036854775808:
                v9 := GenS1{F0: 38, F1: 42}
                _ = v9
            default:
                var v10 float64 = float64(float32(0.0))
                _ = v10
            }
            if false {
                if true {
                    if true {
                        if true {
                            v5 = v5
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        }
                        v2 = -2147483648
                        var v11 float32 = (v5 / v5)
                        _ = v11
                    }
                }
                v5 = (1.5 + (((v5 + (((((0.1 - -2.25) * ((1.5 - ((-2.25 - v5) - (0.1 + 0.1))) / v5)) * ((v5 / (1e10 * v5)) + (v5 - v5))) - v5) * (float32(3.14159) * v5))) * v5) / v5))
            } else {
                fmt.Println(fi8)
                if false {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v5 = v5
                        var v13 int32 = 0
                        if (false && ((!true) || (!true))) {
                            fmt.Println(fi8, v13)
                            var v14 int32 = 2147483647
                            v3 = append(v3, fi8)
                            _ = v14
                        }
                        _ = fi12
                        _ = v13
                    }
                    for fi15 := 0; fi15 < 7; fi15++ {
                        if false {
                            var v16 uint16 = 138
                            _ = v16
                        } else {
                            var v17 float32 = 1.5
                            _ = v17
                        }
                        if (4294967295 <= 42) {
                            v5 -= v5
                        }
                        _ = fi15
                    }
                }
            }
            var v18 float32 = 0.0
            _ = v18
        }
        v2 = 128
        _ = fi8
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(fi19, fi19)
        _ = fi19
    }
    var v20 float64 = 1.5
    if (v20 != (((v20 / float64(v5)) * v20) / float64((0.0 - v5)))) {
        var v21 uint32 = 1
        v20 += 0.1
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(v20, v5)
            _ = fi22
        }
        _ = v21
    } else {
        fmt.Println(v5)
        v20 = (3.14159 + v20)
    }
    var v23 int64 = 9223372036854775807
    var v24 string = "\n\t"
    fmt.Println(v5)
    if true {
        v25 := GenS1{F0: 0, F1: 282}
        v26 := GenS2{F0: (v24 + "a")}
        v23 -= 42
        _ = v25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
