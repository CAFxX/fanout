package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 int32, p1 string) int32 {
    var v5 int64 = 9223372036854775807
    var v6 int8 = -128
    v7 := GenS1{F0: 4294967295, F1: 4294967295, F2: v6}
    var v8 float32 = ((0.0 * 0.1) * 0.0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (p0 + p0)
}
func genFn1(p0 int32, p1 string) uint8 {
    p0 -= 2147483647
    p0 -= 0
    var v9 bool = (65535 > 2)
    v10 := []int{42, 42, 830, 4294967295}
    if len(v10) > 0 {
        v10[0] = 128
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    v3 := []int{}
    var v4 uint64 = 2
    v11 := GenS2{F0: 0}
    v12 := GenS2{F0: 255}
    v4 = (v4 ^ v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ("hello" != "foo bar") {
        var v13 int = -740
        if true {
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint32 = 42
                var v16 int = 47
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 42, F1: 1, F2: 112}
            var v18 uint64 = v4
            _ = v17
            _ = v18
        }
        v19 := GenS3{F0: 1, F1: 1, F2: 0}
        _ = v13
        _ = v19
    }
    v4 -= (v4 + 18446744073709551615)
    v20 := GenS2{F0: 255}
    v3 = append(v3, 0)
    if true {
        if true {
            v4 = v4
        } else {
            fmt.Println(v4)
        }
        fmt.Println(v4)
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v4 += (v4 - v4)
        _ = fi21
    }
    v4 = (^(v4))
    var v22 string = "foo bar"
    var v23 int16 = int16(v4)
    v4 -= (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
