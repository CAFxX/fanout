package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]bool)
    var v3 bool = true
    v4 := make(map[int]string)
    var v5 float64 = (1.5 * 1.5)
    var v6 uint16 = 42
    v1 = v1
    v7 := []int{241, 128, 128, 427}
    var v8 int32 = 508
    switch ("hello" + "a") {
    case "foo bar":
        switch 0 {
        case 2147483647:
            if (uint32(v6) == 1) {
                v9 := GenS1{F0: "", F1: 625}
                var v10 int8 = -12
                switch 744 {
                case 2147483647:
                    _ = len(v7)
                }
                _ = v10
                _ = v9
            } else {
                var v11 int8 = -36
                _ = v11
            }
        default:
            v12 := GenS2{F0: 154}
            _ = v12
        }
    default:
        var v13 bool = (3.14159 == (1e10 * v1))
        _ = v13
    }
    var v14 uint = 991
    var v15 int32 = (v8 % (((v8 >> ((v14) & 63))) | 1))
    v1 = v1
    switch ("" + (("\n\t" + "") + ("\n\t" + "a"))) {
    case "foo bar":
        fmt.Println(v3, v3, v14)
    }
    var v16 float64 = v5
    var v17 int8 = 0
    if len(v7) > 0 {
        v7[0] = 0
    }
    v18 := GenS2{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
