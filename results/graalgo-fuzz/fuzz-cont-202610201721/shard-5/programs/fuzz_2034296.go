package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    v1 := []uint32{42, 4294967295, 2147483647, 4294967295}
    var v2 int8 = 127
    var v3 float32 = 1e10
    var v4 uint16 = 255
    v5 := GenS1{F0: "foo bar", F1: false}
    v4 = 1
    v6 := GenS1{F0: "hello", F1: (1 == 255)}
    var v7 uint32 = 381
    var v8 uint32 = v7
    var v9 uint = 650
    var v10 string = ""
    v11 := GenS2{F0: (v2 * 43), F1: true}
    var v12 uint8 = 1
    v13 := GenS2{F0: v2, F1: false}
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: "", F1: true}
        fmt.Println(v7)
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        switch fi16 {
        case 434:
            var v17 bool = false
            _ = v17
        default:
            v1 = append(v1, v8)
        }
        _ = fi16
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
