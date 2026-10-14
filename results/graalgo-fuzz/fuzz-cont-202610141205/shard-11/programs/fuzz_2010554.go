package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
func genFn0() uint8 {
    if true {
        if true {
            var v4 int = 2147483647
            _ = v4
        } else {
            var v5 int64 = 0
            fmt.Println(v5)
            _ = v5
        }
    } else {
        v6 := []int{}
        var v7 float32 = (((3.14159 - -2.25) * (float32((float64(0.0) * 1e10)) + float32(float64(1.5)))) + 3.14159)
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int8 = 1
        _ = fi8
        _ = v9
    }
    switch 529 {
    case 127:
        var v10 int32 = 42
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    var v12 float64 = 1e10
    _ = v12
    return 2
}
func genFn1(p0 int32, p1 uint32, p2 int16) int8 {
    fmt.Println(p2)
    p1 -= p1
    v13 := GenS1{F0: 42}
    switch ("" + (("a" + "\n\t") + "\n\t")) {
    case "\n\t":
        if false {
            v14 := GenS1{F0: 42}
            var v15 uint16 = 1
            _ = v14
            _ = v15
        }
    case "hello":
        var v16 int = 1
        _ = v16
    case "foo bar":
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 bool = (!false)
            _ = fi17
            _ = v18
        }
    }
    var v19 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v19
    return 1
}

func main() {
    var v1 int16 = -79
    var v2 int = 4294967295
    v3 := make(map[uint32]int)
    fmt.Println(v1, v2)
    v1 += v1
    v20 := []int{v2, 32768, 42, (-(v2))}
    fmt.Println(v2, v1)
    var v21 float64 = ((-2.25 * 0.0) * 0.1)
    v3[42] = (v2 ^ 0)
    v22 := GenS2{F0: v2}
    var v23 uint8 = 128
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
