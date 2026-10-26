package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int) int {
    var v5 int = (p0 - p0)
    var v6 float64 = -2.25
    var v7 int = (-(p0))
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 uint32 = 831
            if true {
                v10 += 0
                if false {
                    v11 := []int{v5, fi9, (v7 * fi8)}
                    _ = v11
                }
                fmt.Println(v6, fi9, v6)
            } else {
                fmt.Println(v6, v7, fi8)
                var v12 uint64 = 18446744073709551615
                v5 = fi8
                _ = v12
            }
            var v13 int16 = 32767
            _ = fi9
            _ = v10
            _ = v13
        }
        var v14 uint64 = 42
        _ = fi8
        _ = v14
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []int{42}
    var v2 uint = 973
    var v3 uint16 = 2
    v4 := make(map[int]uint32)
    var v15 uint16 = v3
    v15 = (v3 << ((v2) & 63))
    if false {
        v16 := GenS1{F0: 2}
        var v17 int8 = -1
        _ = len(v1)
        _ = v16
        _ = v17
    }
    delete(v4, 4294967295)
    var v18 string = ""
    var v19 string = v18
    _ = len(v1)
    var v20 uint64 = 624
    fmt.Println(v3, v19)
    var v21 int = -977
    var v22 string = "\n\t"
    v1 = append(v1, (v21 + (v21 >> (((^(v2))) & 63))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
