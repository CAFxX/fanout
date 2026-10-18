package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
func genFn0(p0 int, p1 int, p2 uint64) uint64 {
    var v5 float32 = 3.14159
    switch ("a" + "\n\t") {
    case "hello":
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    v3 := []int{9223372036854775807, 42}
    var v4 int64 = -1
    var v6 uint32 = v2
    _ = len(v1)
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7, v2, v2)
        _ = fi7
    }
    switch "" {
    case "\n\t":
        v2 -= v6
    case "hello":
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 float32 = (-2.25 - -2.25)
            _ = fi8
            _ = v9
        }
    case "a":
        v10 := GenS1{F0: 933}
        _ = v10
    }
    delete(v1, v6)
    var v11 float64 = float64(0.1)
    _ = len(v3)
    v12 := v1[v2]
    for fi13 := 0; fi13 < 6; fi13++ {
        v12 = false
        v11 += (v11 / float64(0.1))
        if v12 {
            fmt.Println(v2, fi13, v12)
            if true {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v12 = ("hello" != "\n\t")
                    v3 = append(v3, fi14)
                    _ = fi14
                }
                v15 := GenS2{F0: (!false)}
                _ = v15
            }
        } else {
            var v16 float32 = 3.14159
            _ = v16
        }
        _ = fi13
    }
    v17 := GenS2{F0: v12}
    var v18 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
