package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 float64 = 0.0
    var v2 string = ("a" + "foo bar")
    var v3 float32 = 0.1
    var v4 uint16 = 42
    var v5 int = -9223372036854775808
    switch v2 {
    case "foo bar":
        fmt.Println(v2, v1)
    }
    v6 := GenS1{F0: 4294967295, F1: -2147483648, F2: 0}
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            if true {
                fmt.Println(v4, v5)
                var v9 int8 = 13
                _ = v9
            } else {
                var v10 int = (v5 / (((fi7 ^ fi7)) | 1))
                switch fi8 {
                case 895:
                    v11 := []int{127, 127, -30, fi7}
                    switch v5 {
                    case 32767:
                        var v12 string = ((v2 + v2) + v2)
                        _ = v12
                    default:
                        fmt.Println(v5)
                    }
                    _ = v11
                case -1:
                    v10 += (v5 / ((fi7) | 1))
                case 0:
                    var v13 float32 = v3
                    _ = v13
                default:
                    v14 := GenS2{F0: 664, F1: 826, F2: 1}
                    _ = v14
                }
                _ = v10
            }
            v5 = (-(v5))
            v15 := GenS3{F0: v4, F1: 0, F2: 9223372036854775807}
            v16 := GenS2{F0: -1, F1: -966, F2: -229}
            var v17 int32 = -921
            var v18 uint = 4294967295
            var v19 uint64 = 0
            _ = fi8
            _ = v15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi7
    }
    v20 := []int{}
    if len(v20) > 0 {
        v20[0] = v5
    }
    v4 -= (^(v4))
    if true {
        for fi21 := 0; fi21 < 14; fi21++ {
            if (v3 == (((float32((0.1 * (v1 / (v1 - v1)))) * v3) + 0.0) - v3)) {
                _ = len(v20)
                v22 := GenS4{F0: 1000, F1: -119}
                v1 = 3.14159
                _ = v22
            }
            if false {
                switch v2 {
                case "hello":
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                case "a":
                    v4 = (v4 * (v4 / ((42) | 1)))
                }
            } else {
                var v23 float64 = 3.14159
                _ = v23
            }
            _ = len(v20)
            _ = fi21
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS5{F0: 65535, F1: (0 != 42)}
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 18446744073709551615, F1: -896, F2: 18446744073709551615}
    fmt.Println(v3, v1)
    v27 := GenS6{F0: 42}
    fmt.Println(v4, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
