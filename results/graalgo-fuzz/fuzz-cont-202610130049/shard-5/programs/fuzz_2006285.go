package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 1
    var v2 uint64 = 42
    v3 := []string{("" + ((("" + ("a" + "")) + "foo bar") + "foo bar"))}
    v4 := make(map[int]uint32)
    var v5 int64 = 728
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v1)
        var v7 uint = 899
        _ = fi6
        _ = v7
    }
    var v8 bool = true
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := v4[32768]
        v10 -= v10
        _ = fi9
        _ = v10
    }
    v2 += v2
    v2 = 18446744073709551615
    v1 = v1
    v3 = append(v3, "hello")
    v2 += 485
    var v11 uint32 = 1
    v2 = (^(v2))
    v12 := v4[393]
    if v8 {
        var v13 bool = (v8 && v8)
        for fi14 := 0; fi14 < 15; fi14++ {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi14
        }
        var v15 int16 = (v1 + v1)
        _ = v13
        _ = v15
    }
    v16 := v4[42]
    v17 := GenS1{F0: 205}
    fmt.Println(v12, v12, v8)
    v12 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
