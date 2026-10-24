package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint = 4294967295
    var v2 string = "foo bar"
    var v3 int32 = -466
    var v4 uint8 = 255
    v5 := make(map[string]int)
    v6 := make(map[string]uint32)
    var v7 uint64 = 56
    var v8 uint8 = v4
    var v9 bool = true
    if v9 {
        if ((911 > 1) || false) {
            v3 += (v3 | (v3 & (v3 + -2147483648)))
            fmt.Println(v4, v2)
        } else {
            var v10 uint8 = (v4 / ((v4) | 1))
            v3 += (v3 >> ((65535) & 63))
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 int16 = 1
                if true {
                    v4 = (v4 / (((^(v8))) | 1))
                    var v13 uint16 = 2
                    v14 := GenS1{F0: fi11, F1: (v12 + 42)}
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: fi11, F1: v12}
                _ = fi11
                _ = v12
                _ = v15
            }
            _ = v10
        }
        var v16 uint8 = (v4 * v4)
        v6["hello"] = 162
        _ = v16
    }
    var v17 uint = (^(v1))
    v8 = (v4 ^ v8)
    v18 := GenS2{F0: 1}
    v7 = (v7 | v7)
    v19 := GenS1{F0: -9223372036854775808, F1: 815}
    v20 := []int{-1}
    if false {
        var v21 bool = (0 > 42)
        var v22 string = v2
        v20 = append(v20, 42)
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: (((float32((float64(float32(3.14159)) + 3.14159)) == (0.0 * (-2.25 + (-2.25 + 0.1)))) && ((^(v4)) <= (v4 << (((^(v1))) & 63)))) && (v17 >= 0)), F1: (v7 - (v7 / (((v7 >> ((v1) & 63))) | 1)))}
    v4 = 128
    v24 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
