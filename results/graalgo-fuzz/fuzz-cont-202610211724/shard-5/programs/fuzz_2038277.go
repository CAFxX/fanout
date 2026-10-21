package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 757
    var v3 uint16 = 32767
    var v4 int16 = -704
    v5 := make(map[int]bool)
    switch "\n\t" {
    case "":
        v6 := []int{1}
        _ = v6
    default:
        var v7 float32 = 1e10
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint = 255
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 int8 = 0
                fmt.Println(v9, fi10, v4)
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v2, v3, v3)
    fmt.Println(v3, v3, v2)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int = 2147483647
        _ = fi12
        _ = v13
    }
    fmt.Println(v2)
    v14 := []int{-2147483648, 1}
    var v15 int16 = (v4 & v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
