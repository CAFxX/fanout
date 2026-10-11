package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 2
    var v3 int32 = 1
    var v4 string = ""
    v5 := []int{472, 9223372036854775807}
    for fi6 := 0; fi6 < 18; fi6++ {
        delete(v1, 2147483648)
        var v7 string = "a"
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint = 1
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint16 = 286
    fmt.Println(v4, v4)
    var v11 string = (v4 + v4)
    v12 := GenS1{F0: -1}
    v3 = v3
    var v13 int32 = (^(v3))
    switch v4 {
    case "":
        var v14 uint16 = (v10 / ((v10) | 1))
        _ = v14
    default:
        fmt.Println(v3, v3)
    }
    v2 += v2
    v15 := GenS2{F0: 32767}
    var v16 uint32 = 2147483647
    var v17 uint = 255
    var v18 uint16 = 908
    v19 := GenS3{F0: 287}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
