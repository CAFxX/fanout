package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int32 = 2147483647
    var v2 int = 127
    var v3 float32 = -2.25
    var v4 uint = 780
    v5 := make(map[int]string)
    var v6 int32 = (v1 % ((-1) | 1))
    v2 = (v2 - (v2 % (((v2 & v2)) | 1)))
    v7 := []int{(v2 + 32767), (v2 & v2), v2}
    v5[v2] = "hello"
    var v8 float64 = 0.1
    switch v2 {
    case -2147483648:
        if len(v7) > 0 {
            v7[0] = v2
        }
    case 444:
        var v9 int8 = 42
        _ = v9
    case 42:
        v8 += v8
    default:
        if false {
            v5[(v2 >> ((v4) & 63))] = "a"
            switch "\n\t" {
            case "":
                v6 += (v1 >> (((v4 ^ v4)) & 63))
            case "hello":
                var v10 float32 = (v3 - ((v3 / (((v3 + 1.5) + v3) - v3)) - ((v3 * (-2.25 / v3)) - 0.0)))
                _ = v10
            case "foo bar":
                fmt.Println(v3, v3)
            default:
                v2 = v2
            }
            v11 := GenS1{F0: false}
            _ = v11
        }
    }
    switch v2 {
    case -77:
        v3 = v3
    case 2147483647:
        v3 = v3
    }
    var v12 int32 = v1
    var v13 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
