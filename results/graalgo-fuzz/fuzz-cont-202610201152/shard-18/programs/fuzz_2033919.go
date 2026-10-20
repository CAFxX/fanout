package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }

func main() {
    var v1 uint8 = 42
    var v2 string = ("" + (("\n\t" + "foo bar") + ("hello" + "")))
    var v3 float64 = (-2.25 * 1.5)
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    v6 := []int{-334, 594}
    v3 += 3.14159
    var v7 int8 = -1
    var v8 int32 = -1
    var v9 uint8 = v1
    var v10 float64 = v3
    v9 -= v9
    delete(v4, v2)
    fmt.Println(v10)
    v11 := GenS1{F0: 0, F1: (^(v8)), F2: v8}
    v12 := GenS1{F0: 2, F1: (^(v8)), F2: v8}
    switch 32768 {
    case 1:
        var v13 uint = 788
        _ = v13
    default:
        v2 = (("" + (v2 + "a")) + v2)
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v1 = 128
        var v15 uint16 = 32767
        v7 -= 100
        _ = fi14
        _ = v15
    }
    v9 = 2
    v6 = append(v6, int(v9))
    var v16 uint8 = (v1 - (v9 | v1))
    fmt.Println(v10, v8)
    if false {
        for fi17 := 0; fi17 < 11; fi17++ {
            v9 = 127
            var v18 int32 = v8
            v16 += v16
            _ = fi17
            _ = v18
        }
        var v19 int8 = (v7 / (((v7 | (-(v7)))) | 1))
        if false {
            var v20 uint8 = 255
            v21 := GenS1{F0: 861, F1: v8, F2: 0}
            _ = v20
            _ = v21
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
