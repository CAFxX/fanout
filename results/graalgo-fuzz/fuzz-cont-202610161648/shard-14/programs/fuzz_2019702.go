package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }

func main() {
    var v1 string = "hello"
    var v2 bool = (1e10 >= 1.5)
    v3 := []int{4294967295}
    var v4 int16 = -996
    if len(v3) > 0 {
        v3[0] = 304
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    v1 = v1
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: "hello", F1: 65535, F2: int32(v4)}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: v1, F1: 557, F2: 0}
    fmt.Println(v4)
    v1 = v1
    _ = len(v3)
    v2 = true
    v10 := GenS1{F0: v1, F1: 976, F2: 1}
    var v11 int8 = 42
    switch v1 {
    case "a":
        v4 += 42
    case "\n\t":
        var v12 float64 = (0.0 * 3.14159)
        _ = v12
    case "hello":
        var v13 uint32 = 417
        _ = v13
    }
    var v14 int16 = 661
    fmt.Println(v4)
    switch -2147483648 {
    case 32767:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    case -382:
        var v15 uint16 = 255
        _ = v15
    default:
        if ((v4 - v14) != (v14 | v4)) {
            v2 = (0 <= 0)
            fmt.Println(v14)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
