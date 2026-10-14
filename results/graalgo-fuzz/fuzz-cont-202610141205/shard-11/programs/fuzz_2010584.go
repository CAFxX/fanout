package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 255
    v3 := make(map[string]int)
    var v4 string = "\n\t"
    v5 := []int{(v1 | (^(v1))), v1}
    v4 = ((v4 + v4) + (("\n\t" + v4) + v4))
    switch v4 {
    case "hello":
        if true {
            if (0 >= 65535) {
                var v6 int8 = 42
                if (42 > 0) {
                    v3[("\n\t" + "foo bar")] = -943
                    var v7 int64 = 0
                    v6 = v6
                    _ = v7
                }
                _ = v6
            } else {
                v8 := GenS1{F0: 1}
                _ = v8
            }
            v1 = (v1 >> ((0) & 63))
        }
    case "foo bar":
        var v9 bool = true
        _ = v9
    default:
        v3[("hello" + (v4 + v4))] = (v1 * v1)
    }
    if len(v5) > 0 {
        v5[0] = v1
    }
    fmt.Println(v4, v1)
    fmt.Println(v2, v4, v2)
    fmt.Println(v2, v4, v1)
    v1 = (-(v1))
    v10 := GenS2{F0: 0, F1: 389, F2: -1}
    var v11 uint = 18446744073709551615
    v12 := GenS2{F0: (^(v11)), F1: 727, F2: -859}
    v13 := GenS1{F0: 4294967295}
    v14 := GenS3{F0: -1}
    var v15 uint8 = (v2 / ((39) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
