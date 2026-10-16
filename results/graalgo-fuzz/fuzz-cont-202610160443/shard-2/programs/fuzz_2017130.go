package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0(p0 bool, p1 int) uint16 {
    var v4 uint32 = 0
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 173
}

func main() {
    var v1 int16 = -891
    var v2 int16 = (v1 ^ v1)
    v3 := make(map[string]uint32)
    if false {
        var v6 uint64 = 4294967295
        _ = v6
    } else {
        fmt.Println(v1, v2)
        v7 := GenS1{F0: 285, F1: 0}
        _ = v7
    }
    v2 = v2
    switch "hello" {
    case "hello":
        v8 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = v8
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := []int{fi9, (fi9 / (((fi9 ^ -9223372036854775808)) | 1)), (fi9 % (((fi9 - fi9)) | 1))}
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS1{F0: 65535, F1: 1}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint64 = 197
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int16 = v2
        v15 -= (v2 / ((v1) | 1))
        fmt.Println(v1, v2)
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 string = "hello"
        _ = fi16
        _ = v17
    }
    var v18 int8 = 127
    var v19 uint16 = 427
    var v20 int32 = -2147483648
    v1 += (v1 & -1)
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 float32 = 1.5
        v23 := []int{127, 713}
        _ = fi21
        _ = v22
        _ = v23
    }
    v1 = (v1 ^ (-(v2)))
    var v24 float32 = 1e10
    v25 := GenS1{F0: 616, F1: 42}
    var v26 uint32 = 194
    var v27 float32 = float32((float64(3.14159) + ((-2.25 + 1e10) - float64(1e10))))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
