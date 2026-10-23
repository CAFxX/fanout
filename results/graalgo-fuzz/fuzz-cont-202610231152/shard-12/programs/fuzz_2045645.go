package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 0
    var v3 int64 = 42
    fmt.Println(v2, v3)
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: fi4, F1: v3}
        var v6 int16 = 0
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 string = "hello"
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 128
        if (65535 == 0) {
            var v10 float64 = float64(((1.5 * (1e10 / float32(-2.25))) - 0.1))
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := []int{(^(fi8)), (^(fi8)), (fi11 << ((65535) & 63)), (fi8 | int(v3))}
                v2 += (v2 ^ 465)
                var v13 string = (v7 + v7)
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int16 = -1
            delete(v1, v7)
            _ = fi14
            _ = v15
        }
        _ = fi8
        _ = v9
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 float32 = (0.1 + (0.1 + -2.25))
        v1[v7] = v7
        _ = fi16
        _ = v17
    }
    var v18 string = ""
    var v19 int16 = -825
    v20 := []int{319}
    v1[v7] = "foo bar"
    v18 = v7
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
