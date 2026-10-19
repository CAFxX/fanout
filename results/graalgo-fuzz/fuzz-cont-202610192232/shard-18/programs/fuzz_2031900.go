package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint64 = 673
    v2 := make(map[int]bool)
    v3 := []uint32{65535, 2, 815, 2147483647}
    var v4 int8 = -24
    v5 := make(map[string]int)
    if ((-68 != -1) || (42 == 310)) {
        var v6 uint = 255
        var v7 int32 = -1
        var v8 float32 = 3.14159
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 65535
        switch "foo bar" {
        case "hello":
            fmt.Println(v1, v9, v9)
        case "foo bar":
            switch 9223372036854775807 {
            case 2147483647:
                var v10 uint = 42
                _ = v10
            default:
                fmt.Println(v1, v1, v9)
            }
        case "a":
            v3 = append(v3, v9)
        }
        v2[241] = (false && true)
        _ = v9
    }
    v11 := GenS1{F0: true}
    v1 = v1
    v1 += 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 -= 127
    v5[""] = 726
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
