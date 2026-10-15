package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
func genFn0(p0 int16, p1 uint32) uint8 {
    if (-1 >= -128) {
        p0 += p0
        fmt.Println(p0)
        v6 := []int{-9223372036854775808}
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 0, F1: 0, F2: 4294967295}
        p1 = p1
        p0 += (p0 - p0)
        _ = v9
    }
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 int8, p1 int16, p2 uint8) uint16 {
    v10 := GenS2{F0: 351, F1: 42, F2: 4294967295}
    p2 -= (p2 / (((p2 - 231)) | 1))
    p2 = (p2 / (((^(p2))) | 1))
    p0 += 0
    v11 := GenS2{F0: 973, F1: 0, F2: 732}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 133
}

func main() {
    var v1 int64 = 1
    var v2 uint64 = 1
    v3 := make(map[string]bool)
    var v4 uint16 = 32767
    var v5 float64 = 0.0
    var v12 uint16 = 145
    v12 = (v12 & (^(v4)))
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v12, v2)
        v1 -= v1
        delete(v3, "foo bar")
        _ = fi13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v1 = (v1 & (^(v1)))
        _ = fi14
    }
    if false {
        if (true || ((-332 != -358) || (!(7 <= -1)))) {
            var v15 uint8 = 255
            _ = v15
        }
        var v16 int = 80
        _ = v16
    } else {
        v17 := GenS2{F0: 2147483647, F1: 112, F2: v2}
        v12 += v12
        _ = v17
    }
    fmt.Println(v12, v4, v5)
    v1 = -526
    v18 := []int{42, 128, 0}
    v4 -= (v12 - (v4 ^ (v4 | 32767)))
    var v19 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
