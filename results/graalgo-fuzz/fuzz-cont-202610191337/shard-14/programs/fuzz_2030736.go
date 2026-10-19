package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 42
    var v3 int = -9223372036854775808
    var v4 float32 = 1.5
    var v5 int64 = -1
    if (v2 >= (v2 % ((42) | 1))) {
        v6 := GenS1{F0: 0}
        fmt.Println(v5)
        _ = v6
    } else {
        var v7 uint32 = 0
        if false {
            v4 = ((float32(0.1) * float32((3.14159 / 1e10))) + (v4 + 3.14159))
        } else {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint = 255
                v7 += (v7 & 376)
                fmt.Println(v7, v5, fi8)
                _ = fi8
                _ = v9
            }
            fmt.Println(v4)
            v3 += (v3 | v3)
        }
        fmt.Println(v7, v3, v4)
        _ = v7
    }
    fmt.Println(v4, v3, v5)
    v10 := GenS1{F0: (v2 | (v2 >> ((2) & 63)))}
    fmt.Println(v4, v3)
    v2 = (v2 ^ (v2 % ((v2) | 1)))
    var v11 bool = false
    fmt.Println(v3)
    v12 := v1[42]
    fmt.Println(v11, v4)
    v12 = (223 < 231)
    var v13 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
