package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []uint32{0}
    v2 := []int{}
    v3 := []uint32{606}
    switch "hello" {
    case "a":
        var v4 uint32 = 98
        _ = v4
    }
    var v5 int64 = 1
    switch 137 {
    case 32767:
        var v6 float64 = -2.25
        _ = v6
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS1{F0: 104}
    fmt.Println(v5)
    if ((float32(((0.1 + 3.14159) - 3.14159)) == 1.5) || (-21 < 100)) {
        if true {
            v5 = (-(v5))
        } else {
            if false {
                v8 := GenS1{F0: 1}
                _ = v8
            }
            var v9 int16 = -1
            var v10 int16 = v9
            _ = v10
            _ = v9
        }
    } else {
        v2 = append(v2, 2147483647)
        _ = len(v3)
    }
    var v11 uint = 2
    var v12 uint64 = 42
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
