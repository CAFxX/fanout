package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }
func genFn0(p0 int32, p1 uint, p2 int32) string {
    var v7 int = -194
    if true {
        switch v7 {
        case 2147483647:
            var v8 bool = false
            _ = v8
        }
        fmt.Println(p1, p2, p0)
        switch (-(v7)) {
        case -2147483648:
            if true {
                v9 := []int{(-(v7)), v7, v7}
                v10 := GenS1{F0: (v7 - v7)}
                _ = v10
                _ = v9
            }
        }
    } else {
        var v11 float64 = 3.14159
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("\n\t" + "\n\t")
}
func genFn1() int {
    var v12 uint = 308
    v13 := GenS1{F0: 42}
    var v14 bool = true
    _ = v12
    _ = v13
    _ = v14
    return 2147483647
}

func main() {
    var v1 uint16 = 305
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := []uint32{}
    var v5 uint64 = 18446744073709551615
    var v6 float32 = 0.1
    v6 += v6
    var v15 int32 = 0
    for fi16 := 0; fi16 < 8; fi16++ {
        v5 = (v5 - (v5 & (v5 * (^(v5)))))
        var v17 int64 = 10
        v17 += v17
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 42}
    var v19 int8 = 127
    v20 := GenS1{F0: 127}
    for fi21 := 0; fi21 < 15; fi21++ {
        fmt.Println(v15)
        v4 = append(v4, 2147483648)
        _ = fi21
    }
    fmt.Println(v1, v19, v5)
    v22 := GenS2{F0: true, F1: (("hello" + ("foo bar" + "hello")) + "\n\t"), F2: v5}
    fmt.Println(v1, v1, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
