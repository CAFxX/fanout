package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{"a"}
    var v2 uint64 = 107
    var v3 int16 = 42
    v4 := make(map[int]int)
    var v5 int = -2147483648
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 255}
        _ = fi6
        _ = v7
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if (!false) {
        var v8 float32 = ((3.14159 + 3.14159) / 3.14159)
        _ = v8
    }
    switch 32767 {
    case 42:
        var v9 int16 = 318
        _ = v9
    case 1:
        v1 = append(v1, "hello")
    case 4294967295:
        if false {
            v10 := GenS1{F0: 255}
            _ = v10
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += (v2 - (v2 ^ (v2 | v2)))
    var v11 uint = 255
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v3)
        switch "" {
        case "a":
            v13 := GenS1{F0: 1}
            _ = v13
        default:
            var v14 int16 = v3
            _ = v14
        }
        _ = fi12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
