package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 int64, p1 uint) int {
    v5 := GenS1{F0: 1, F1: 174}
    p1 = 65535
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 bool) int8 {
    p1 = (0 >= 255)
    var v6 string = ""
    switch (v6 + v6) {
    case "hello":
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 float64 = 0.1
                if p1 {
                    v9 := GenS1{F0: 127, F1: 0}
                    v10 := GenS2{F0: (p0 + 526), F1: 0}
                    _ = v10
                    _ = v9
                }
                var v11 int32 = p0
                _ = fi7
                _ = v11
                _ = v8
            }
            var v12 float32 = -2.25
            _ = v12
        }
    default:
        var v13 float32 = 0.1
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 uint32 = 1
    var v2 uint = 638
    var v3 bool = false
    v4 := make(map[int]string)
    fmt.Println(v1, v1, v2)
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint64 = 0
        v16 := []int{fi14, 9223372036854775807, (fi14 ^ fi14)}
        _ = len(v16)
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = v3
    var v17 float64 = float64(1e10)
    if false {
        v18 := []int{4294967295, genFn0(-1, v2), 42, 4294967295}
        fmt.Println(v17, v17, v2)
        delete(v4, -1)
        _ = v18
    }
    v19 := []int{127}
    for fi20 := 0; fi20 < 14; fi20++ {
        v17 = float64(1.5)
        _ = fi20
    }
    _ = len(v19)
    v21 := GenS2{F0: -1, F1: v2}
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
