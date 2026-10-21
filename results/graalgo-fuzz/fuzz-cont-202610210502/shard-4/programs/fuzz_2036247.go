package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int = 308
    var v2 int = v1
    var v3 float64 = float64(1.5)
    v4 := []int{v1, v1, (v1 | (^(v2)))}
    v5 := make(map[int]string)
    var v6 int = (v2 % ((2147483647) | 1))
    v7 := v5[(-(v2))]
    var v8 uint64 = 980
    var v9 string = v7
    var v10 int8 = -7
    var v11 uint16 = 1
    var v12 uint8 = 42
    switch v9 {
    case "foo bar":
        v13 := GenS1{F0: (v11 >> ((407) & 63))}
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 5; fi15++ {
                delete(v5, 427)
                _ = fi15
            }
            fmt.Println(v10, v11, v1)
            _ = fi14
        }
    }
    v11 = 0
    v16 := GenS1{F0: 1}
    v12 = (v12 ^ uint8(v8))
    v11 = (^(v11))
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
