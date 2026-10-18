package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0() uint8 {
    var v5 int8 = -128
    v5 = (v5 ^ v5)
    v5 -= (^(v5))
    _ = v5
    return 255
}

func main() {
    v1 := []string{"a", ("foo bar" + ("" + "hello"))}
    var v2 uint = 852
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    v6 := GenS1{F0: 0, F1: 0, F2: 45}
    switch "" {
    case "":
        v7 := GenS1{F0: 2, F1: 1, F2: uint8(v2)}
        _ = v7
    case "hello":
        if len(v1) > 0 {
            v1[0] = ("" + ("\n\t" + ""))
        }
    case "a":
        fmt.Println(v2)
    }
    var v8 uint8 = 127
    if false {
        var v9 float64 = (-2.25 - 3.14159)
        if (!false) {
            fmt.Println(v9, v2, v8)
            var v10 int64 = -1
            _ = v10
        }
        _ = v9
    }
    v8 = v8
    v8 = v8
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS2{F0: true}
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int = fi11
            _ = fi13
            _ = v14
        }
        v3["hello"] = 2147483647
        _ = fi11
        _ = v12
    }
    if true {
        var v15 float32 = -2.25
        v2 = 2
        v15 = v15
        _ = v15
    } else {
        v8 += 231
        switch "" {
        case "a":
            for fi16 := 0; fi16 < 3; fi16++ {
                _ = len(v3)
                _ = fi16
            }
        case "":
            v17 := GenS1{F0: 359, F1: 903, F2: 0}
            _ = v17
        }
    }
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
