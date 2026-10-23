package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 int32) uint {
    var v4 uint32 = 859
    var v5 uint = 0
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 128
    v3 := make(map[int]string)
    v6 := GenS1{F0: 1}
    v2 -= v2
    v3[959] = "a"
    if true {
        v7 := GenS2{F0: -19}
        _ = v7
    } else {
        if (42 <= -626) {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS3{F0: (fi8 | -728), F1: 255}
                v2 = (v2 + v2)
                v2 -= 217
                _ = fi8
                _ = v9
            }
            v2 += (v2 * (v2 / ((v2) | 1)))
        }
    }
    var v10 bool = true
    v11 := []int{(v2 - (v2 >> ((1) & 63)))}
    var v12 float64 = -2.25
    for fi13 := 0; fi13 < 2; fi13++ {
        v10 = (32767 != 0)
        _ = fi13
    }
    v14 := GenS4{F0: (42 == 59)}
    fmt.Println(v2, v12)
    v12 += v12
    fmt.Println(v10, v10)
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
