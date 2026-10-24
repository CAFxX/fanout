package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    var v1 uint64 = 525
    var v2 int64 = -1
    var v3 uint32 = 65535
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    switch 128 {
    case 97:
        v2 = (v2 & (v2 - 1))
    }
    switch -358 {
    case 1:
        var v6 int32 = 2147483647
        _ = v6
    case 333:
        v7 := []int{387}
        _ = v7
    case -620:
        v8 := GenS1{F0: 0, F1: true}
        _ = v8
    }
    if false {
        v2 = 42
        v9 := v5[""]
        switch "" {
        case "foo bar":
            for fi10 := 0; fi10 < 15; fi10++ {
                v5[("foo bar" + "foo bar")] = fi10
                _ = fi10
            }
        case "hello":
            v2 = -1
        case "a":
            var v11 float32 = (1e10 * 0.1)
            _ = v11
        default:
            var v12 uint8 = 255
            _ = v12
        }
        _ = v9
    }
    var v13 uint64 = v1
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 uint = 237
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v3, fi16, v3)
        v13 = uint64(v3)
        _ = fi16
    }
    v17 := []int{230, -852, 1}
    v18 := GenS1{F0: -1, F1: true}
    var v19 float32 = ((-2.25 - 0.0) + 3.14159)
    v20 := GenS1{F0: 912, F1: (true || true)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
