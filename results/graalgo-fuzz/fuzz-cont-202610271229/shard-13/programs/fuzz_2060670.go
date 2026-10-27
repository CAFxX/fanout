package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint = 907
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    var v4 int16 = -1
    var v5 float32 = (0.0 - -2.25)
    v6 := GenS1{F0: (v2 << ((42) & 63)), F1: (v2 ^ v2), F2: 2147483648}
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: v2, F1: v2, F2: 0}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = uint8(v1)
    if false {
        var v10 int32 = 42
        v1 += (v1 | v1)
        _ = v10
    } else {
        v11 := GenS2{F0: 2147483647, F1: v1}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int = (fi12 % ((fi12) | 1))
        if false {
            if ((^(v1)) > v1) {
                var v14 int = fi12
                v15 := GenS3{F0: (v2 & v2)}
                fmt.Println(v13, v5)
                _ = v14
                _ = v15
            }
        }
        if false {
            v9 = (v9 * (v9 ^ (v9 ^ 255)))
            v16 := []int{fi12, (v13 % (((v13 * fi12)) | 1)), fi12, (v13 & (fi12 * (fi12 + 1)))}
            if (103 < 942) {
                fmt.Println(v13, v2, v5)
                var v17 int64 = -9223372036854775808
                for fi18 := 0; fi18 < 13; fi18++ {
                    switch 198 {
                    case 1:
                        var v19 int8 = (^(v2))
                        _ = v19
                    case -9223372036854775808:
                        v9 -= (^(v9))
                    case 128:
                        v13 -= int(v2)
                    default:
                        v3[2147483647] = (fi12 - fi18)
                    }
                    _ = fi18
                }
                _ = v17
            }
            _ = v16
        } else {
            for fi20 := 0; fi20 < 10; fi20++ {
                v21 := []int{128}
                if true {
                    delete(v3, 1)
                    v5 -= (v5 * v5)
                    fmt.Println(v4, v13)
                }
                _ = fi20
                _ = v21
            }
            v22 := GenS2{F0: 2147483647, F1: 1}
            _ = v22
        }
        _ = fi12
        _ = v13
    }
    v23 := GenS2{F0: 1, F1: (^(v1))}
    v5 = ((1e10 * ((float32(1e10) - v5) * 3.14159)) * float32(1.5))
    v24 := GenS3{F0: (^(v2))}
    var v25 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
