package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
func genFn0(p0 int) uint8 {
    v6 := []int{(p0 >> ((255) & 63))}
    p0 += p0
    _ = p0
    _ = v6
    return 128
}
func genFn1(p0 uint16, p1 int32) uint64 {
    v7 := []int{127, -183}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 596
}

func main() {
    v1 := []uint32{750, 4294967295}
    v2 := []uint32{}
    var v3 int64 = -9223372036854775808
    var v4 string = "\n\t"
    var v5 float32 = 1e10
    fmt.Println(v4, v4, v4)
    for fi8 := 0; fi8 < 7; fi8++ {
        v5 += float32(1.5)
        _ = fi8
    }
    switch "a" {
    case "foo bar":
        if (v3 != (v3 >> ((422) & 63))) {
            var v9 int8 = 127
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 float32 = v5
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 uint = 658
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            var v14 float64 = 0.1
            _ = v14
            _ = v9
        }
    }
    var v15 float32 = (v5 / v5)
    if (0 < 1) {
        var v16 float64 = 1.5
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v16 = v16
        _ = v16
    } else {
        v15 += v5
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v17 := GenS1{F0: 2147483647, F1: 42, F2: 0}
        _ = v17
    }
    var v18 int8 = 42
    v19 := GenS1{F0: -2147483648, F1: 0, F2: 1}
    v3 += v3
    if false {
        v3 += v3
        fmt.Println(v18)
        v3 += v3
    }
    switch v4 {
    case "hello":
        v2 = append(v2, 2147483647)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
