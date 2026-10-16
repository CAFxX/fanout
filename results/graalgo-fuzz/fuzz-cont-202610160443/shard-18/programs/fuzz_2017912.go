package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []int{127, 42, -869}
    var v2 int32 = 801
    v3 := make(map[int]uint32)
    v4 := []uint32{678, 2147483648}
    var v5 uint16 = 1
    fmt.Println(v2)
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: fi6, F1: 1}
        _ = fi6
        _ = v7
    }
    fmt.Println(v2)
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v5, v5)
    default:
        var v8 uint32 = 537
        _ = v8
    }
    v9 := GenS2{F0: "foo bar"}
    v5 = (v5 / ((1) | 1))
    if false {
        switch 0 {
        case 2147483647:
            var v10 string = ("\n\t" + "")
            _ = v10
        case -674:
            switch "\n\t" {
            case "foo bar":
                var v11 string = ("foo bar" + "foo bar")
                _ = v11
            default:
                var v12 bool = (1e10 > 0.0)
                _ = v12
            }
        }
        fmt.Println(v2, v5)
    }
    v13 := GenS3{F0: 65535}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 float32 = 0.0
    v14 -= 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
