package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 42
    var v3 int16 = 1
    v4 := GenS1{F0: 4294967295}
    fmt.Println(v2, v3)
    if true {
        var v5 float64 = 1.5
        for fi6 := 0; fi6 < 10; fi6++ {
            delete(v1, int(v2))
            _ = fi6
        }
        _ = v5
    }
    if true {
        var v7 uint32 = 727
        _ = v7
    } else {
        v1[1] = 42
        if true {
            delete(v1, -41)
            v3 -= (v3 + v3)
            v3 = v3
        }
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint16 = (v2 + (^(v2)))
        _ = fi8
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 0}
            var v12 float32 = 0.1
            var v13 string = "\n\t"
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 int32 = -613
        _ = v14
    }
    v15 := []int{}
    var v16 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
