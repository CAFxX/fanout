package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64; F2 uint32 }
func genFn0() uint64 {
    var v4 int16 = -278
    if (!true) {
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := []int{fi5}
            v6 = append(v6, 32768)
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 42, F1: 127}
        _ = v7
    }
    _ = v4
    return 4294967295
}

func main() {
    var v1 int16 = 32767
    var v2 int32 = -19
    var v3 int64 = -9223372036854775808
    fmt.Println(v1, v2, v2)
    var v8 int64 = 232
    var v9 bool = true
    v10 := []int{-2147483648}
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 float64 = ((float64(float32((3.14159 + 1e10))) - (1e10 + 1.5)) + 1.5)
        v13 := GenS2{F0: (fi11 << ((18446744073709551615) & 63))}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = float32((3.14159 * 1.5))
    switch "\n\t" {
    case "hello":
        if v9 {
            if (796 > genFn0()) {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS3{F0: v1, F1: 0, F2: 42}
                    _ = fi15
                    _ = v16
                }
                v8 += v8
            }
            var v17 float32 = (((0.0 + v14) - 3.14159) / v14)
            _ = v17
        }
    }
    v18 := GenS2{F0: -682}
    var v19 bool = (32767 <= 255)
    for fi20 := 0; fi20 < 10; fi20++ {
        _ = len(v10)
        _ = fi20
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
