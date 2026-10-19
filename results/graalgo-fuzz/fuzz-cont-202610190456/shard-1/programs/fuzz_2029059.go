package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
func genFn0(p0 int8, p1 int8) uint32 {
    switch -9223372036854775808 {
    case -1:
        p0 = (p0 & 18)
    case 32767:
        v7 := GenS1{F0: "foo bar"}
        _ = v7
    case 128:
        v8 := []int{-564}
        _ = v8
    }
    p1 += p1
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := []uint32{}
    var v2 uint = 42
    var v3 bool = true
    var v4 int = 1
    v5 := make(map[string]uint32)
    v6 := []int{(v4 / (((v4 | (v4 % (((v4 & (v4 % ((42) | 1)))) | 1)))) | 1))}
    fmt.Println(v3, v4, v3)
    fmt.Println(v2, v3, v4)
    var v9 float64 = -2.25
    var v10 string = "hello"
    _ = len(v5)
    _ = len(v5)
    v9 = (v9 - (v9 / (3.14159 * 0.1)))
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 20; fi12++ {
        switch (^(fi12)) {
        case 4294967295:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := v5[v10]
                var v15 uint = (^(v2))
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 int32 = 0
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS1{F0: (v10 + v10)}
            v19 := GenS2{F0: v2}
            fmt.Println(fi17)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi12
        _ = v16
    }
    v20 := GenS3{F0: 1, F1: 75, F2: "\n\t"}
    fmt.Println(v11, v11, v2)
    if len(v6) > 0 {
        v6[0] = v4
    }
    v10 = "hello"
    var v21 string = "\n\t"
    v22 := GenS2{F0: (v2 / (((v2 << (((v2 - 318)) & 63))) | 1))}
    if v3 {
        fmt.Println(v4, v11)
        var v23 int64 = v11
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
