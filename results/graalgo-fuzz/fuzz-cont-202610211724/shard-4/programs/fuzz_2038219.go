package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int64, p1 uint32, p2 int8) string {
    if (-74 > 9223372036854775807) {
        var v7 uint32 = p1
        var v8 uint32 = v7
        _ = v7
        _ = v8
    }
    var v9 uint = 42
    v10 := GenS1{F0: v9, F1: -1}
    var v11 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return ("hello" + "\n\t")
}

func main() {
    var v1 float64 = (0.0 + ((((1e10 + 1.5) - -2.25) / ((-2.25 - -2.25) + float64((0.1 * 1.5)))) * (-2.25 + float64(-2.25))))
    var v2 int64 = 9223372036854775807
    var v3 bool = (-294 <= -2147483648)
    v4 := make(map[string]string)
    var v5 uint = 18446744073709551615
    v6 := make(map[string]string)
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v2)
        _ = fi12
    }
    v13 := GenS2{F0: 2147483647, F1: (v2 | (v2 * (^(v2))))}
    v14 := GenS2{F0: 128, F1: (v2 << (((v5 << ((v5) & 63))) & 63))}
    v15 := []int{2147483647, -1, 0, 9223372036854775807}
    var v16 int16 = 1000
    if (128 < 0) {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        fmt.Println(v3, v5, v2)
        fmt.Println(v2)
    } else {
        v16 = (v16 - -32768)
        v17 := GenS3{F0: 350, F1: 992}
        if (1000 >= v16) {
            v5 -= v5
            var v18 int64 = (v2 / ((v2) | 1))
            var v19 uint8 = 233
            _ = v18
            _ = v19
        }
        _ = v17
    }
    fmt.Println(v3, v16, v3)
    if (!v3) {
        fmt.Println(v5)
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            v15 = append(v15, -685)
            var v22 uint = 739
            _ = len(v15)
            _ = fi21
            _ = v22
        }
        v1 = (v1 - (v1 * float64((0.1 + (((0.0 + 3.14159) - (0.1 + 1.5)) - float32(float64((0.1 / 0.1))))))))
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
