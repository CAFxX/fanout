package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int16; F2 int8 }

func main() {
    var v1 uint = 2
    var v2 bool = (9223372036854775807 != 0)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    var v5 float32 = ((-2.25 / -2.25) * (float32(0.1) / float32((0.1 * (((1e10 - (((1e10 - 0.1) + (-2.25 / 1e10)) + 3.14159)) - ((3.14159 * ((((0.0 - 1.5) * 3.14159) / 1e10) / 1e10)) / 3.14159)) / float64(((1e10 * -2.25) / 3.14159)))))))
    fmt.Println(v2, v2, v5)
    v6 := GenS1{F0: v2, F1: 2, F2: 1000}
    if ((100 > -120) || false) {
        v7 := v3[686]
        _ = v7
    } else {
        v8 := []int{-373}
        v9 := GenS1{F0: (237 != 622), F1: 1, F2: 32767}
        _ = v8
        _ = v9
    }
    v10 := []int{127, 0, 0, -9223372036854775808}
    for fi11 := 0; fi11 < 17; fi11++ {
        if true {
            v12 := GenS2{F0: 1}
            _ = v12
        }
        v5 = ((v5 * ((v5 * v5) - ((v5 - ((float32(((float64((1.5 - v5)) / -2.25) - (1e10 - 3.14159))) - 0.1) * v5)) - float32(((-2.25 * (-2.25 + 0.1)) - 1.5))))) * (((v5 / v5) * v5) + v5))
        v5 -= v5
        _ = fi11
    }
    fmt.Println(v1, v1)
    if false {
        delete(v3, 65535)
    }
    var v13 int8 = 18
    v1 -= (v1 % (((v1 / (((v1 >> (((v1 % (((v1 - v1)) | 1))) & 63))) | 1))) | 1))
    fmt.Println(v5, v5, v1)
    v14 := GenS3{F0: (v1 ^ 18446744073709551615), F1: -1, F2: v13}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint8 = 255
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
