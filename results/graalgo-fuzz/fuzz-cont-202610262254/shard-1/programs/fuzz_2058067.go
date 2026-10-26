package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
func genFn0(p0 string, p1 uint16) int32 {
    var v6 uint32 = 331
    p0 = p0
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}
func genFn1(p0 bool, p1 int) int64 {
    p0 = false
    p0 = ((1e10 >= (float64(0.1) - ((0.1 / -2.25) + (3.14159 + float64(float32((-2.25 * 3.14159))))))) || p0)
    p0 = ((p1 & 850) >= (^(p1)))
    var v7 float64 = 1e10
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 string = "hello"
    var v2 string = v1
    v3 := []uint32{}
    var v4 uint64 = 1
    var v5 string = ("a" + v1)
    switch (v5 + v2) {
    case "":
        v4 -= (v4 - (v4 % (((v4 % (((v4 ^ (v4 + v4))) | 1))) | 1)))
    }
    if false {
        v8 := GenS1{F0: v4, F1: -1, F2: 9223372036854775807}
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v5)
            _ = fi9
        }
        _ = v8
    }
    var v10 float64 = (1e10 + (0.1 + 0.0))
    v10 += v10
    fmt.Println(v1, v2, v4)
    var v11 float32 = 0.0
    for fi12 := 0; fi12 < 10; fi12++ {
        _ = len(v3)
        if false {
            v13 := GenS2{F0: 2147483647, F1: -1, F2: 0}
            _ = v13
        }
        var v14 uint = 1
        _ = fi12
        _ = v14
    }
    var v15 int = -296
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
