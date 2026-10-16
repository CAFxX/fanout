package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint16 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    var v3 int32 = 2147483647
    var v4 bool = (4294967295 == 1)
    v5 := []string{(("hello" + (("" + ("a" + "")) + "foo bar")) + "")}
    var v6 float64 = (v1 / v1)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 float64 = v6
            v10 := GenS1{F0: true, F1: 1}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 int64 = 9223372036854775807
    var v12 string = "\n\t"
    if v4 {
        v12 = v12
        v12 = "a"
        v2[65535] = 4294967295
    }
    delete(v2, 4294967295)
    v1 = v6
    if (127 == 42) {
        if v4 {
            if (3.14159 != float32(((3.14159 - (float64(0.1) / (v1 + (((float64(0.0) * ((((((-2.25 + v6) * float64(0.0)) * v6) / (v6 - float64(0.0))) + 3.14159) + float64(3.14159))) + v6) / v1)))) - (v6 - float64((1e10 / 1.5)))))) {
                if (2 != 54) {
                    var v13 int64 = -9223372036854775808
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := GenS2{F0: fi14}
                        v12 = (v12 + v12)
                        if (42 < 4294967295) {
                            v5 = append(v5, (v12 + (v12 + (((((v12 + v12) + v12) + (("a" + v12) + "\n\t")) + v12) + v12))))
                            v16 := GenS2{F0: fi14}
                            if v4 {
                                var v17 uint64 = 900
                                var v18 string = (v12 + v12)
                                _ = v17
                                _ = v18
                            }
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                    v6 = (v1 * ((0.1 / (((v1 * (3.14159 - 0.1)) + v1) / v1)) - 3.14159))
                }
            } else {
                v6 -= v6
            }
        }
        fmt.Println(v4, v4)
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    v19 := GenS3{F0: 18446744073709551615, F1: 508}
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(v4, v6, v6)
        _ = fi20
    }
    _ = len(v2)
    v4 = (-536 == -1)
    var v21 uint = 255
    v22 := GenS1{F0: v4, F1: 32767}
    v23 := GenS4{F0: 42, F1: v11, F2: 220}
    v24 := GenS5{F0: (v21 % ((18446744073709551615) | 1)), F1: 2, F2: 1}
    var v25 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
