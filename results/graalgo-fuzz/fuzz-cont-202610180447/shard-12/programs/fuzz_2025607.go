package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint16) int {
    var v4 uint16 = (p0 | p0)
    var v5 int8 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 += p0
        _ = fi6
    }
    var v7 string = ""
    p0 += v4
    _ = p0
    _ = v4
    _ = v5
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    if (-12 > -197) {
        fmt.Println(v3)
        v8 := []int{2147483647, 128, 469}
        if (float64(1.5) >= 1.5) {
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 int64 = 450
                _ = fi9
                _ = v10
            }
            var v11 int64 = -749
            if (!true) {
                v3 = v3
            }
            _ = v11
        }
        _ = v8
    }
    fmt.Println(v3)
    var v12 int64 = 0
    var v13 int32 = 1
    v14 := GenS1{F0: 747}
    var v15 uint16 = 0
    delete(v1, 42)
    fmt.Println(v13, v3, v3)
    var v16 string = "\n\t"
    for fi17 := 0; fi17 < 2; fi17++ {
        v13 = v13
        _ = fi17
    }
    v18 := []int{-944, 475}
    v15 = (^(v15))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
