package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 234
    var v3 uint = 40
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5, v2, v3)
        _ = fi5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{(fi7 * (fi7 & (fi6 + -9223372036854775808))), (-(fi7)), fi7}
            v4["foo bar"] = false
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    fmt.Println(v2, v2)
    v3 += (^(v3))
    fmt.Println(v2)
    var v9 uint16 = 32767
    v2 -= (^(v2))
    var v10 uint = (^(v3))
    var v11 uint32 = 1
    fmt.Println(v2)
    v3 = 18446744073709551615
    fmt.Println(v10, v9, v3)
    if false {
        var v12 float32 = float32((0.1 + 0.1))
        v3 = 4294967295
        _ = v12
    }
    delete(v1, "")
    v13 := GenS1{F0: 641, F1: (732 >= 161), F2: int64(v11)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
