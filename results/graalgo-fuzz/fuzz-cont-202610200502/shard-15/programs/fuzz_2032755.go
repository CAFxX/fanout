package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint32, p1 int) uint8 {
    var v4 uint = 996
    v5 := []int{(p1 | p1), 1, (p1 >> ((v4) & 63))}
    var v6 int64 = -724
    v7 := GenS1{F0: (p1 + (p1 >> (((v4 | (^(v4)))) & 63)))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := []string{}
    v2 := []uint32{65535, 1, 2147483647}
    var v3 string = "\n\t"
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if len(v2) > 0 {
        v2[0] = 750
    }
    var v8 float64 = 1.5
    var v9 bool = (false || true)
    v10 := GenS1{F0: 32768}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: fi11}
        if len(v1) > 0 {
            v1[0] = ("a" + v3)
        }
        if v9 {
            fmt.Println(fi11)
            fmt.Println(fi11, v9)
        }
        _ = fi11
        _ = v12
    }
    var v13 float64 = float64(0.0)
    v9 = false
    v14 := GenS1{F0: 42}
    var v15 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
