package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint16 = 0
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 uint16 = v1
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 128, F1: 127}
    v6["hello"] = 1
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 12; fi9++ {
        delete(v3, (fi9 >> ((uint(v4)) & 63)))
        fmt.Println(fi9, v4)
        _ = fi9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 float32 = v8
        switch -797 {
        case -2147483648:
            fmt.Println(v1, v8)
        }
        _ = fi10
        _ = v11
    }
    var v12 bool = (4294967295 < 42)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v13 int32 = -1
    var v14 uint8 = 0
    var v15 int64 = -1
    switch 175 {
    case -1:
        var v16 int32 = 0
        _ = v16
    case 42:
        v2 = append(v2, "a")
    case 32768:
        var v17 uint32 = 2147483648
        _ = v17
    default:
        if v12 {
            var v18 int32 = (v13 % ((v13) | 1))
            _ = v18
        } else {
            for fi19 := 0; fi19 < 6; fi19++ {
                v5[4294967295] = (2 < 206)
                _ = fi19
            }
        }
    }
    v12 = v12
    if (!v12) {
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        _ = len(v5)
    }
    v20 := GenS1{F0: 931, F1: 73}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
