package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 425
    v3 := make(map[int]string)
    if false {
        var v4 int32 = 0
        v2 = v2
        v5 := []int{42}
        _ = v4
        _ = v5
    } else {
        v6 := []int{145}
        fmt.Println(v2)
        v7 := GenS1{F0: 221, F1: "a"}
        _ = v6
        _ = v7
    }
    fmt.Println(v2)
    v8 := GenS1{F0: 65535, F1: "foo bar"}
    v9 := GenS1{F0: 255, F1: ""}
    _ = len(v1)
    v2 = (v2 * (v2 - v2))
    var v10 int = 968
    v11 := []int{32767, 4294967295, (v10 | -896), v10}
    if true {
        var v12 uint8 = 177
        v2 = (v2 | 737)
        _ = v12
    } else {
        var v13 float64 = (0.1 + 1e10)
        for fi14 := 0; fi14 < 15; fi14++ {
            v2 = 615
            _ = fi14
        }
        _ = v13
    }
    fmt.Println(v10)
    v15 := GenS1{F0: 32767, F1: "foo bar"}
    var v16 uint16 = 65535
    for fi17 := 0; fi17 < 16; fi17++ {
        delete(v3, 240)
        _ = fi17
    }
    if true {
        v18 := GenS2{F0: v2}
        v19 := GenS2{F0: v2}
        v2 = (^(v2))
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: (v2 | v2)}
    if false {
        var v21 float64 = ((-2.25 - -2.25) - 3.14159)
        var v22 float32 = (((((1.5 * -2.25) * 1.5) - 0.1) * float32(v21)) + 0.1)
        _ = v21
        _ = v22
    }
    v10 -= 32768
    v2 = (v2 / ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
