package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 int64, p1 int64, p2 uint) string {
    p1 = -57
    var v4 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}

func main() {
    var v1 int = 0
    v2 := []uint32{0}
    var v3 uint16 = 1
    v5 := GenS1{F0: "foo bar", F1: 748}
    v2 = append(v2, 2147483647)
    fmt.Println(v1, v3)
    if true {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(v3, fi6)
            _ = fi6
        }
        v1 += (v1 / (((v1 | v1)) | 1))
    }
    switch "hello" {
    case "a":
        var v7 int32 = int32(v3)
        _ = v7
    case "foo bar":
        var v8 int16 = -283
        _ = v8
    case "hello":
        v9 := GenS2{F0: 217, F1: 100}
        _ = v9
    }
    var v10 int16 = 0
    var v11 int16 = v10
    if (false || (1e10 == ((-2.25 / -2.25) / (3.14159 * (0.1 + -2.25))))) {
        _ = len(v2)
    }
    var v12 uint = 0
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int32 = -2147483648
        var v15 uint32 = 2147483648
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = append(v2, 2)
    fmt.Println(v10, v11, v10)
    v12 -= v12
    var v16 int = v1
    switch v1 {
    case 1:
        v17 := GenS1{F0: "", F1: (v1 * 657)}
        _ = v17
    default:
        if len(v2) > 0 {
            v2[0] = 109
        }
    }
    var v18 uint32 = 821
    if len(v2) > 0 {
        v2[0] = 65535
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 float64 = 1e10
        fmt.Println(v12)
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v5
}
