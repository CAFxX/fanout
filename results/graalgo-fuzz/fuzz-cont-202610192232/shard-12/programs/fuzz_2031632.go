package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 550
    var v3 int32 = -143
    v4 := make(map[int]bool)
    v2 += v2
    var v5 float64 = ((((((float64((0.1 / 1.0)) * (float64(3.14159) / 0.1)) - ((0.1 - 0.0) - 0.1)) + float64(0.1)) * 0.1) * 0.0) / 3.14159)
    v6 := []int{127, 1, 42}
    fmt.Println(v5, v3)
    v3 -= (v3 | (v3 * v3))
    var v7 string = ("foo bar" + ("a" + "foo bar"))
    v7 = (v7 + ("hello" + ("a" + "a")))
    v8 := GenS1{F0: 1}
    if (1 != -93) {
        v9 := v1[1]
        _ = v9
    }
    fmt.Println(v2)
    if (true || (!(v7 < v7))) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 255}
            var v12 int32 = -1
            _ = fi10
            _ = v11
            _ = v12
        }
        if true {
            fmt.Println(v2)
            var v13 uint64 = 1
            _ = v13
        } else {
            v14 := GenS2{F0: 783}
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS1{F0: 255}
                _ = fi15
                _ = v16
            }
            if true {
                fmt.Println(v2)
                v5 += v5
                v1[127] = 42
            } else {
                if (!true) {
                    v17 := GenS3{F0: (v2 / (((-(v2))) | 1)), F1: 42, F2: 127}
                    var v18 int32 = (^(v3))
                    if (0 != 128) {
                        for fi19 := 0; fi19 < 19; fi19++ {
                            for fi20 := 0; fi20 < 7; fi20++ {
                                var v21 uint16 = 42
                                if (false || ((fi19 & 9223372036854775807) < fi19)) {
                                    var v22 uint32 = 2147483647
                                    v23 := GenS4{F0: 69, F1: 27}
                                    _ = v22
                                    _ = v23
                                } else {
                                    if false {
                                        v6 = append(v6, fi20)
                                    }
                                    var v24 int32 = v3
                                    v2 = (v2 % ((int64(v21)) | 1))
                                    _ = v24
                                }
                                v25 := GenS5{F0: v3, F1: (((false && true) && false) && false)}
                                _ = fi20
                                _ = v21
                                _ = v25
                            }
                            v2 += v2
                            _ = fi19
                        }
                    }
                    _ = v17
                    _ = v18
                } else {
                    v7 = "\n\t"
                    var v26 int64 = (v2 % (((v2 | (v2 << ((18446744073709551615) & 63)))) | 1))
                    _ = v26
                }
                var v27 uint8 = 128
                var v28 uint16 = 65535
                _ = v27
                _ = v28
            }
            _ = v14
        }
        v7 = ("hello" + ((v7 + v7) + v7))
    } else {
        fmt.Println(v7)
    }
    v2 -= v2
    v29 := GenS6{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
