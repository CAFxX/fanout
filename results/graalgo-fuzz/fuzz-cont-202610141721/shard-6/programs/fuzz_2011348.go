package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16) uint32 {
    var v4 int64 = 42
    v5 := GenS1{F0: 806, F1: 281, F2: true}
    v6 := GenS1{F0: 42, F1: 2, F2: true}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1() int16 {
    if false {
        var v7 string = ("foo bar" + "foo bar")
        if false {
            var v8 int8 = 0
            _ = v8
        }
        _ = v7
    } else {
        switch 646 {
        case -1:
            v9 := []int{487, -911, 1, 128}
            _ = v9
        case -547:
            v10 := GenS2{F0: 42}
            _ = v10
        }
        var v11 float64 = 0.0
        _ = v11
    }
    v12 := []int{327, -183, -833, 128}
    _ = v12
    return 1
}

func main() {
    var v1 float32 = 0.0
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]uint32)
    var v13 string = "hello"
    v1 = 0.0
    for fi14 := 0; fi14 < 20; fi14++ {
        v1 += float32((-2.25 - (-2.25 / 1.0)))
        _ = fi14
    }
    v15 := []int{}
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS1{F0: 403, F1: 0, F2: false}
        _ = fi16
        _ = v17
    }
    v2 = v2
    v3[v2] = v2
    v3[v2] = (v2 | 42)
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
