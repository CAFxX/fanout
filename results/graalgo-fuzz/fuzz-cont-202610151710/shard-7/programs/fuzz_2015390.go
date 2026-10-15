package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 int = 2147483647
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]string)
    var v4 int16 = 938
    v5 := []int{42, (^(v1)), 521}
    v6 := GenS1{F0: 255}
    v7 := GenS1{F0: 1}
    var v8 int64 = 9223372036854775807
    var v9 int32 = -608
    fmt.Println(v2, v4)
    var v10 int = (-(v1))
    v9 += (v9 | (v9 & (^(v9))))
    v11 := GenS2{F0: "a", F1: (-(v9))}
    v12 := GenS1{F0: 0}
    v1 += (v1 * (-(v1)))
    if (!true) {
        if false {
            v13 := GenS1{F0: 58}
            fmt.Println(v10, v4)
            _ = v13
        }
    }
    v4 = v4
    if ((false && (!(0.0 <= float64((3.14159 + -2.25))))) || (true || false)) {
        v2 -= (-(v8))
        var v14 uint16 = 42
        _ = v14
    } else {
        var v15 int8 = 44
        var v16 float32 = -2.25
        _ = v15
        _ = v16
    }
    if (float64(float32(0.0)) > float64(1e10)) {
        switch (^(v1)) {
        case -685:
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS1{F0: 31}
                if ((1.5 / (1e10 * 3.14159)) < -2.25) {
                    v4 += v4
                    if ((((1e10 * 0.1) * float64(1.5)) == 0.0) && (0 != 90)) {
                        var v19 uint = 1
                        _ = v19
                    }
                    v20 := GenS1{F0: 48}
                    _ = v20
                }
                v21 := GenS2{F0: "\n\t", F1: (v9 << ((625) & 63))}
                _ = fi17
                _ = v18
                _ = v21
            }
        case -9223372036854775808:
            fmt.Println(v8, v2)
        }
        var v22 int = v10
        var v23 float64 = 1e10
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
