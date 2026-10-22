package main

import (
    "fmt"
)


func main() {
    var v1 float64 = ((1e10 - (0.1 * 1.5)) * 3.14159)
    v2 := make(map[int]string)
    var v3 bool = true
    var v4 string = ("a" + "")
    var v5 int16 = -1
    v6 := []int{790}
    var v7 int64 = 42
    for fi8 := 0; fi8 < 9; fi8++ {
        _ = len(v2)
        v2[(fi8 / ((fi8) | 1))] = v4
        _ = fi8
    }
    v7 -= (v7 << ((255) & 63))
    if len(v6) > 0 {
        v6[0] = 127
    }
    var v9 uint16 = 65535
    v1 = v1
    v9 -= (v9 - 441)
    if v3 {
        _ = len(v6)
        if v3 {
            v3 = false
            if v3 {
                v3 = true
                if false {
                    delete(v2, -739)
                }
            } else {
                if len(v6) > 0 {
                    v6[0] = 0
                }
            }
        }
        var v10 string = ("hello" + "a")
        _ = v10
    }
    var v11 bool = ((0 <= 42) && (true || (("a" + v4) > v4)))
    v7 = (v7 | v7)
    v1 = 1.5
    for fi12 := 0; fi12 < 13; fi12++ {
        v7 += (-(v7))
        _ = fi12
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
