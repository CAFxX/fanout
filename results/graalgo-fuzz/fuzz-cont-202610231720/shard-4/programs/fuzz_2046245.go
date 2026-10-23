package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]string)
    var v3 int = 200
    v4 := make(map[string]uint32)
    var v5 uint64 = 2
    v6 := v2[65535]
    switch (v6 + v6) {
    case "a":
        var v7 uint8 = 128
        _ = v7
    }
    delete(v2, 811)
    v8 := []int{}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = (0.0 / (-2.25 + 1e10))
        if (274 == 42) {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 string = ""
                var v13 int8 = 127
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi9
        _ = v10
    }
    if (!(!(915 < 4294967295))) {
        var v14 uint16 = 595
        v6 = v6
        v15 := GenS1{F0: v6, F1: "hello"}
        _ = v14
        _ = v15
    } else {
        fmt.Println(v1, v6)
        var v16 int = (-(v3))
        fmt.Println(v16)
        _ = v16
    }
    v3 = (-(v3))
    var v17 int64 = 0
    v8 = append(v8, (v3 % ((v3) | 1)))
    v18 := v2[2147483647]
    var v19 int64 = (^(v17))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
