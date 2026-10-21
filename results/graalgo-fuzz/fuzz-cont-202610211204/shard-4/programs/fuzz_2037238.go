package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []uint32{2, 42, 2147483647, 184}
    var v2 float32 = float32((3.14159 - 0.1))
    var v3 string = "hello"
    var v4 int8 = 82
    v5 := []uint32{}
    v6 := make(map[string]string)
    fmt.Println(v4, v3)
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 float32 = (((float32(0.1) * 0.1) + v2) - v2)
            fmt.Println(v9)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: ""}
        var v11 bool = (!true)
        _ = fi7
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: "\n\t"}
    if (false || true) {
        fmt.Println(v3, v2)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = 3.14159
    v2 = float32(1.5)
    v3 = "\n\t"
    fmt.Println(v4, v2, v2)
    if true {
        _ = len(v5)
        fmt.Println(v4, v4, v3)
    }
    var v13 int64 = -508
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
