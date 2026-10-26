package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    v2 := []string{"foo bar", "hello"}
    v3 := make(map[string]uint32)
    var v4 uint16 = 65535
    v5 := make(map[string]int)
    var v6 int = -9223372036854775808
    var v7 float32 = 1e10
    v4 -= (v4 & 42)
    fmt.Println(v6, v6, v7)
    switch v6 {
    case 811:
        v5["foo bar"] = v6
    case -9223372036854775808:
        fmt.Println(v4, v6)
    default:
        v7 = v7
    }
    var v8 int16 = -566
    v9 := v1["foo bar"]
    for fi10 := 0; fi10 < 4; fi10++ {
        v6 += v9
        _ = fi10
    }
    v8 -= (v8 >> ((554) & 63))
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v4 = v4
            fmt.Println(v9, v6, v7)
            var v13 float32 = v7
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v6, v8)
        v7 = (float32((1.5 / 1.0)) / 1.0)
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
