package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
func genFn0(p0 string, p1 int, p2 string) int32 {
    p2 = (("\n\t" + "") + p2)
    if false {
        var v5 int16 = -226
        _ = v5
    }
    v6 := []int{9223372036854775807, p1, (^(p1))}
    if true {
        var v7 bool = (3.14159 >= 1.5)
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p0)
            _ = fi8
        }
        _ = v7
    } else {
        v6 = append(v6, 32768)
        for fi9 := 0; fi9 < 10; fi9++ {
            _ = len(v6)
            var v10 float32 = (-2.25 + 0.0)
            _ = fi9
            _ = v10
        }
        if (2 >= 0) {
            p0 = p0
            var v11 uint = 158
            v11 = (v11 * v11)
            _ = v11
        } else {
            p2 = "\n\t"
            p2 = p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483647
    var v3 uint = 383
    v4 := []int{714}
    var v12 string = "a"
    if true {
        delete(v1, v12)
        var v13 bool = (-1 < 60)
        _ = v13
    }
    var v14 int = -28
    var v15 uint64 = 1
    if true {
        v3 += (v3 + (^(v3)))
        var v16 float64 = 0.1
        var v17 uint8 = 0
        _ = v16
        _ = v17
    } else {
        var v18 int8 = -1
        v19 := GenS1{F0: v18, F1: 42, F2: 9223372036854775807}
        var v20 uint8 = uint8(v3)
        _ = v18
        _ = v19
        _ = v20
    }
    v1[v12] = (-(v14))
    var v21 float32 = (3.14159 - 3.14159)
    switch (v14 & v14) {
    case -1:
        for fi22 := 0; fi22 < 6; fi22++ {
            v21 += float32(3.14159)
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint = 65535
                v25 := GenS1{F0: -128, F1: -9223372036854775808, F2: 1}
                v1["hello"] = fi22
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 int64 = -9223372036854775808
                var v28 uint8 = 134
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi22
        }
    case 32768:
        v29 := v1[v12]
        _ = v29
    case -9223372036854775808:
        fmt.Println(v12, v21)
    }
    var v30 float64 = 1.5
    v12 = ((v12 + v12) + "hello")
    v14 = (v14 * v14)
    var v31 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
