package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 int8; F1 uint8; F2 int32 }
func genFn0(p0 int16, p1 uint) int32 {
    v6 := GenS1{F0: "hello", F1: 42, F2: 1}
    p1 += p1
    var v7 int8 = 1
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    v3 := []string{v1, (v1 + v1)}
    var v4 float32 = 1e10
    var v5 uint64 = 2
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(fi9, v4)
        _ = fi9
    }
    fmt.Println(v1, v1, v1)
    if true {
        var v10 float32 = float32((1.5 - 1.5))
        v11 := GenS1{F0: (v1 + v1), F1: 122, F2: 127}
        _ = v10
        _ = v11
    } else {
        v4 = v4
    }
    switch 128 {
    case 785:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v12 := GenS2{F0: 127, F1: 170, F2: 706}
    fmt.Println(v5, v1, v1)
    v5 -= v5
    fmt.Println(v1, v5)
    var v13 int8 = 1
    fmt.Println(v1)
    for fi14 := 0; fi14 < 12; fi14++ {
        v4 = v4
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
