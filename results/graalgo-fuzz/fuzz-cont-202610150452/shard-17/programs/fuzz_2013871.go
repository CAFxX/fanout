package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 127
    v2 := make(map[string]uint32)
    v3 := []string{"a", "hello", "\n\t", "\n\t"}
    var v4 uint64 = 2
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: -562}
    var v7 int8 = 100
    var v8 uint32 = 65535
    var v9 int32 = -478
    switch v1 {
    case 849:
        fmt.Println(v1)
    case 32768:
        fmt.Println(v8)
    default:
        fmt.Println(v9)
    }
    v7 += 51
    var v10 int32 = (v9 >> ((0) & 63))
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 6; fi12++ {
        if (!((v4 - v4) >= (v4 ^ 4294967295))) {
            fmt.Println(v4, v8, v1)
        } else {
            if false {
                fmt.Println(v10)
            } else {
                v13 := GenS1{F0: 42}
                if false {
                    v14 := GenS1{F0: (v11 | (v11 >> ((4294967295) & 63)))}
                    _ = v14
                } else {
                    var v15 uint16 = 65535
                    _ = v15
                }
                var v16 string = "hello"
                _ = v13
                _ = v16
            }
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
