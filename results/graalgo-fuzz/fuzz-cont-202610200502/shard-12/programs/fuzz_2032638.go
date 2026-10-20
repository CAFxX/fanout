package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 bool = (1.5 < 0.1)
    v2 := []int{-9223372036854775808, 32767, -2147483648}
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    var v5 uint16 = 110
    var v6 float64 = 1e10
    v5 = (v5 & 1)
    for fi7 := 0; fi7 < 14; fi7++ {
        if v1 {
            var v8 uint64 = 4294967295
            if (v5 > (^(v5))) {
                if v1 {
                    v9 := GenS1{F0: int8(v5), F1: false}
                    _ = v9
                }
                if (83 > 42) {
                    v10 := GenS2{F0: (v3 | (v3 - v3))}
                    v1 = true
                    _ = v10
                } else {
                    switch (fi7 << ((42) & 63)) {
                    case 128:
                        _ = len(v4)
                    case 32768:
                        v11 := GenS3{F0: 42}
                        _ = v11
                    default:
                        v5 += 255
                    }
                    for fi12 := 0; fi12 < 11; fi12++ {
                        if ((v8 ^ v8) == (v8 >> ((2) & 63))) {
                            var v13 int16 = -1
                            v14 := GenS3{F0: 2}
                            _ = v13
                            _ = v14
                        } else {
                            v5 -= v5
                            v6 = (v6 * (0.1 / (0.1 + v6)))
                            switch ("a" + "\n\t") {
                            case "":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                fmt.Println(v5)
                            }
                        }
                        var v15 float64 = v6
                        v8 += (v8 ^ 757)
                        _ = fi12
                        _ = v15
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v8
        }
        var v16 uint8 = 128
        v17 := GenS1{F0: 78, F1: v1}
        _ = fi7
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: v3}
    switch "foo bar" {
    case "a":
        if v1 {
            v3 = v3
        } else {
            _ = len(v2)
            v6 -= ((float64(3.14159) * v6) - v6)
        }
    }
    v19 := GenS4{F0: (!((v5 ^ 2) < (v5 << ((936) & 63))))}
    if v1 {
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                v4[fi20] = uint32(v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v22 int = 1
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    } else {
        v6 -= (((v6 + v6) / v6) / (0.1 / float64(3.14159)))
        v3 = -1
        v5 += (^(v5))
    }
    if (2 == 255) {
        _ = len(v2)
    }
    v23 := GenS2{F0: -2147483648}
    delete(v4, 42)
    fmt.Println(v1, v6, v1)
    fmt.Println(v1)
    v5 = v5
    v3 = (v3 % (((v3 << ((1) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
