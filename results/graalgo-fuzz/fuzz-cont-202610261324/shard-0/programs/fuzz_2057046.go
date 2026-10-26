package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int) bool {
    p0 = -567
    var v4 string = (("foo bar" + "\n\t") + "\n\t")
    v5 := []int{p0}
    switch p0 {
    case 1:
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int8 = 42
            fmt.Println(v4)
            _ = fi6
            _ = v7
        }
    default:
        v4 = ("foo bar" + v4)
    }
    _ = p0
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 uint16, p1 uint32) int16 {
    v8 := []int{-9223372036854775808, -627, 4294967295}
    v9 := GenS1{F0: -577}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 383
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 127
    v3 := make(map[uint32]uint32)
    v2 = 42
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float64 = float64(1e10)
        _ = fi10
        _ = v11
    }
    var v12 int = -2147483648
    fmt.Println(v2, v12)
    if (0.1 == (float64(3.14159) / 1.0)) {
        fmt.Println(v2, v2)
    }
    fmt.Println(v2)
    v12 = 128
    for fi13 := 0; fi13 < 5; fi13++ {
        v2 -= v2
        _ = fi13
    }
    var v14 uint16 = 32767
    var v15 float32 = ((-2.25 / 1.5) + 0.0)
    var v16 int8 = -128
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint8 = 42
        if false {
            v19 := v3[2147483648]
            var v20 bool = (!true)
            _ = v19
            _ = v20
        }
        v12 = v12
        _ = fi17
        _ = v18
    }
    if (!false) {
        v21 := GenS2{F0: (v14 | v14), F1: 2}
        var v22 uint16 = v14
        _ = v21
        _ = v22
    }
    v14 -= 605
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
