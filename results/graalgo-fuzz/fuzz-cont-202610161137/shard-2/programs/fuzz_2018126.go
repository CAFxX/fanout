package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32, p1 int8) uint64 {
    v4 := []int{1, 188}
    fmt.Println(p1)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint32 = 65535
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1() int64 {
    v7 := GenS1{F0: -32768, F1: -2147483648, F2: false}
    var v8 string = "foo bar"
    switch "" {
    case "":
        var v9 uint32 = 614
        _ = v9
    case "foo bar":
        v8 = v8
    case "a":
        v8 = "hello"
    default:
        switch 1 {
        case 4294967295:
            v10 := GenS1{F0: -1, F1: -1, F2: (0.0 >= (1.5 - (float32(1.5) + float32(0.1))))}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := []int{741, fi11}
        _ = fi11
        _ = v12
    }
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 814
    var v2 bool = false
    var v3 uint8 = 42
    v13 := []int{}
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: 42, F1: 42, F2: true}
        v16 := GenS2{F0: 2, F1: 32767}
        _ = fi14
        _ = v15
        _ = v16
    }
    v1 = (v1 * (v1 ^ (-(v1))))
    var v17 float32 = ((1e10 / (((0.1 * ((0.0 / 1.5) * float32(3.14159))) + 0.0) + (-2.25 * ((((0.1 / 0.1) / ((1.5 / 0.1) + (1e10 - 0.1))) - 3.14159) / 1.5)))) - 0.0)
    v1 = -188
    v17 -= v17
    var v18 int64 = -1
    var v19 uint8 = 2
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 uint64 = 572
        var v22 uint32 = 941
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS1{F0: 32767, F1: -1, F2: true}
        _ = fi23
        _ = v24
    }
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
