package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64, p1 uint16) int64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(p1)
        var v8 int32 = 1
        p1 -= p1
        _ = fi7
        _ = v8
    }
    var v9 uint16 = (p1 + 255)
    fmt.Println(p1, v9)
    p0 -= -9223372036854775808
    _ = p0
    _ = p1
    _ = v9
    return (p0 & p0)
}
func genFn1(p0 uint16, p1 uint16) int {
    var v10 int32 = -2147483648
    p0 = p1
    v11 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -324
}

func main() {
    var v1 int8 = -128
    var v2 uint = 1
    var v3 int32 = -884
    var v4 uint64 = 4294967295
    var v5 string = "hello"
    v6 := make(map[string]string)
    v1 += 53
    fmt.Println(v1)
    v2 = v2
    v12 := []int{2147483647, 128}
    var v13 float64 = (((0.1 / -2.25) + 0.1) / 0.1)
    switch -9223372036854775808 {
    case 32767:
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 float32 = 3.14159
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case -2147483648:
        var v17 int64 = 9223372036854775807
        _ = v17
    }
    var v18 float32 = 0.0
    if false {
        delete(v6, "\n\t")
        _ = len(v6)
    }
    var v19 float32 = 0.0
    v4 += (v4 - 2)
    v20 := GenS2{F0: 2, F1: (v3 << (((^(v2))) & 63))}
    v1 += -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
