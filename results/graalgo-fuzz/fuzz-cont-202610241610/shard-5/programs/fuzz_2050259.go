package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }

func main() {
    var v1 uint8 = 2
    var v2 int32 = -2147483648
    v3 := []uint32{785, 0, 939, 42}
    v4 := make(map[int]bool)
    var v5 bool = true
    var v6 bool = (!(v2 >= v2))
    v7 := GenS1{F0: 2, F1: 4294967295}
    var v8 int64 = -71
    var v9 int8 = -56
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "\n\t":
        if (65535 < 1) {
            switch "foo bar" {
            case "":
                _ = len(v3)
            default:
                var v10 uint16 = 65535
                _ = v10
            }
            if false {
                v5 = false
                v11 := GenS2{F0: 0, F1: 2147483647, F2: 934}
                _ = v11
            } else {
                v8 += v8
            }
        }
    }
    v3 = append(v3, 493)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
