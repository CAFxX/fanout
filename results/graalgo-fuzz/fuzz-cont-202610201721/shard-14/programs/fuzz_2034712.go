package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint64; F2 uint }
func genFn0() int64 {
    v5 := GenS1{F0: -615, F1: "foo bar", F2: 100}
    switch 2147483647 {
    case 55:
        var v6 int = 32768
        _ = v6
    case 2147483647:
        var v7 bool = (42 < 42)
        _ = v7
    case -223:
        var v8 uint64 = 4294967295
        _ = v8
    default:
        v9 := GenS2{F0: 63, F1: 1}
        _ = v9
    }
    v10 := []int{-354, -9223372036854775808}
    var v11 int64 = -1
    _ = v10
    _ = v11
    _ = v5
    return v11
}
func genFn1() string {
    v12 := GenS2{F0: 2, F1: -1}
    if false {
        var v13 uint32 = 4294967295
        v13 += v13
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS3{F0: (v13 * 990), F1: -1, F2: 32767}
            v13 += 2147483648
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 uint32 = 558
        v17 := GenS4{F0: -2147483648}
        var v18 uint16 = 681
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 1
    var v20 bool = true
    _ = v12
    _ = v19
    _ = v20
    return "a"
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{(^(v1)), (v1 & v1), (-(v1))}
    var v3 uint16 = 136
    v4 := make(map[string]int)
    v3 = 486
    if len(v2) > 0 {
        v2[0] = v1
    }
    v21 := GenS5{F0: 909, F1: 2, F2: 255}
    v22 := v4[("\n\t" + "hello")]
    var v23 int8 = 100
    var v24 uint8 = 73
    var v25 float32 = float32(0.1)
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int64 = 170
        fmt.Println(v3)
        _ = fi26
        _ = v27
    }
    _ = len(v4)
    fmt.Println(v24, v22)
    v3 -= (v3 * v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
