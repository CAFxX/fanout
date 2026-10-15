package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32, p1 int32, p2 string) uint64 {
    p0 = (-(p0))
    v5 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 897
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{818, 32767, 854}
    var v3 bool = false
    var v4 uint = 42
    var v6 uint8 = 1
    var v7 uint64 = 432
    fmt.Println(v3, v4, v4)
    var v8 uint8 = (v6 + v6)
    if v3 {
        fmt.Println(v4, v6, v7)
    } else {
        fmt.Println(v8, v3, v8)
        if true {
            switch "\n\t" {
            case "hello":
                switch "\n\t" {
                case "a":
                    var v9 int16 = 1
                    _ = v9
                default:
                    var v10 bool = (42 <= 379)
                    _ = v10
                }
            case "":
                var v11 string = "foo bar"
                _ = v11
            default:
                var v12 int16 = 48
                _ = v12
            }
        } else {
            var v13 uint32 = 4294967295
            _ = v13
        }
        var v14 int8 = -98
        _ = v14
    }
    fmt.Println(v3)
    var v15 uint16 = 188
    var v16 uint8 = (v8 | 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
