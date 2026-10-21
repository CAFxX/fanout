package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }

func main() {
    var v1 uint32 = 1
    v2 := []string{"foo bar", ("hello" + "hello")}
    var v3 int16 = -32768
    var v4 uint64 = 1
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = (v1 + 4294967295)
        _ = fi6
        _ = v7
    }
    delete(v5, 1)
    for fi8 := 0; fi8 < 3; fi8++ {
        v3 += (v3 / ((v3) | 1))
        _ = fi8
    }
    v4 += v4
    v4 += v4
    v3 = v3
    if ((((-2.25 - (1.5 * ((1.5 + 3.14159) + 1.5))) + (0.0 + (1e10 * float32(float64(0.1))))) / 1e10) >= 0.1) {
        delete(v5, 911)
    } else {
        fmt.Println(v1, v1, v4)
        v9 := GenS1{F0: -2147483648, F1: -2147483648, F2: "hello"}
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: fi10, F1: 341, F2: "a"}
        _ = fi10
        _ = v11
    }
    fmt.Println(v4, v3)
    fmt.Println(v1, v1, v4)
    var v12 int64 = 42
    var v13 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
