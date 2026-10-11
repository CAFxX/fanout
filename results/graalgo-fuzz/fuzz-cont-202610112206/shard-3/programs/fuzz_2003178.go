package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 int64 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 string = (("" + ("" + "foo bar")) + "foo bar")
    if (255 < 4294967295) {
        v4 := []int{-1, int(v1), 71, -1}
        if true {
            var v5 float32 = 1e10
            _ = v5
        }
        _ = v4
    } else {
        v1 -= 2
        v3 = v3
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        switch fi6 {
        case 703:
            fmt.Println(v1)
        default:
            v2[42] = 288
        }
        _ = fi6
    }
    v7 := []int{-5, 796}
    v8 := GenS1{F0: -44, F1: -935, F2: -2147483648}
    v7 = append(v7, -988)
    switch v3 {
    case "hello":
        if false {
            v9 := GenS2{F0: "a", F1: 0, F2: 389}
            var v10 string = v3
            _ = v10
            _ = v9
        } else {
            switch 32768 {
            case -2147483648:
                var v11 int64 = -9223372036854775808
                _ = v11
            default:
                v3 = (("a" + v3) + "hello")
            }
        }
    }
    var v12 int64 = 0
    v12 = v12
    var v13 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
