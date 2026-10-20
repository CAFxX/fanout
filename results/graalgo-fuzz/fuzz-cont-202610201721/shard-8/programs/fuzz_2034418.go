package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string) int8 {
    p0 = "a"
    v7 := []int{}
    p0 = "hello"
    fmt.Println(p0)
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    _ = p0
    _ = v7
    return -126
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]string)
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 1
    var v5 uint64 = (v3 + 0)
    v6 := make(map[string]string)
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := []int{(^(fi9)), (^(fi9)), (fi9 & -843)}
        _ = fi9
        _ = v10
    }
    var v11 bool = (4294967295 < 0)
    var v12 float64 = (((3.14159 * 0.0) - 0.0) - (float64((((0.0 - (3.14159 + float32(float64(-2.25)))) - -2.25) + (3.14159 + -2.25))) * 0.0))
    if true {
        v13 := GenS1{F0: 2, F1: (v5 | v5), F2: 65535}
        v14 := []int{-9223372036854775808, 2147483647, -536, 127}
        _ = v13
        _ = v14
    }
    var v15 int32 = 33
    fmt.Println(v12)
    switch (("foo bar" + ("" + "hello")) + "a") {
    case "a":
        v16 := GenS1{F0: 32767, F1: 154, F2: 493}
        _ = v16
    case "\n\t":
        v17 := GenS2{F0: v1}
        _ = v17
    }
    var v18 uint8 = 127
    var v19 float32 = (float32((v12 * (3.14159 * float64(float32(((((1e10 / 3.14159) / v12) + 0.0) + v12)))))) * float32(v12))
    var v20 uint64 = (v3 * v5)
    if (!(int8(v18) < 42)) {
        v21 := GenS3{F0: 621}
        v4 -= v4
        _ = v21
    } else {
        v11 = (int8(v5) == 127)
        var v22 uint32 = 4294967295
        _ = v22
    }
    var v23 int64 = (v4 ^ v4)
    v24 := []int{-9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
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
}
