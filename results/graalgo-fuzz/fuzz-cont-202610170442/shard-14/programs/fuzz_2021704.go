package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0() int8 {
    switch 2147483647 {
    case 127:
        var v5 uint64 = 2
        _ = v5
    default:
        var v6 uint8 = 218
        _ = v6
    }
    v7 := GenS1{F0: 18446744073709551615, F1: 1}
    var v8 float32 = 1.5
    var v9 int = 128
    v9 = v9
    _ = v7
    _ = v8
    _ = v9
    return -128
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    var v3 float32 = float32(1.5)
    v4 := make(map[int]string)
    var v10 int8 = 1
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 1, F1: 2147483647}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS1{F0: 4294967295, F1: 1}
    switch 32768 {
    case 0:
        v15 := []int{187, 32767, 9223372036854775807}
        _ = v15
    case 9223372036854775807:
        v3 = -2.25
    case -9223372036854775808:
        v16 := v2[0]
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: 18446744073709551615, F1: 279}
        _ = fi17
        _ = v18
    }
    var v19 string = ("foo bar" + "hello")
    switch ("" + v19) {
    case "hello":
        var v20 uint32 = 87
        _ = v20
    case "a":
        var v21 uint = 886
        _ = v21
    }
    v22 := GenS2{F0: 291}
    switch v19 {
    case "a":
        v23 := GenS1{F0: 42, F1: 1}
        _ = v23
    case "":
        var v24 uint32 = 755
        _ = v24
    case "hello":
        var v25 float64 = 1e10
        _ = v25
    default:
        var v26 float64 = -2.25
        _ = v26
    }
    fmt.Println(v19)
    v10 = (v10 >> ((2) & 63))
    v27 := []int{-2147483648, 2147483647, 983, 466}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
}
