package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = ""
    var v2 int16 = -1
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    fmt.Println(v1, v1)
    _ = len(v3)
    for fi4 := 0; fi4 < 4; fi4++ {
        v1 = "\n\t"
        v5 := []int{454, -982, fi4}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        if false {
            fmt.Println(v2)
            v7 := GenS1{F0: 0}
            _ = v7
        }
        v8 := []int{fi6, fi6, (-(fi6)), 347}
        fmt.Println(v1, v1, v2)
        _ = fi6
        _ = v8
    }
    fmt.Println(v2)
    v2 += v2
    v1 = v1
    v2 += 415
    v2 += 471
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
