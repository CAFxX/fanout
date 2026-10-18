package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 int; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{-2147483648, -18, 32767}
    var v3 bool = (0 >= 521)
    v4 := make(map[string]string)
    var v5 int64 = -995
    v3 = v3
    if v3 {
        if v3 {
            var v6 string = ""
            _ = v6
        }
    } else {
        if (!v3) {
            if (42 >= 0) {
                v7 := GenS1{F0: 0, F1: 2, F2: (-32768 > 401)}
                v3 = v3
                _ = v7
            } else {
                _ = len(v2)
            }
        }
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v5, v3, v5)
        var v9 float64 = 1.5
        v10 := GenS1{F0: -2147483648, F1: 1, F2: false}
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = v3
    if v3 {
        if v3 {
            var v11 int32 = 1
            _ = v11
        } else {
            var v12 int64 = (v5 ^ -885)
            v5 -= v5
            _ = v12
        }
        delete(v4, "\n\t")
        v13 := GenS1{F0: -1, F1: 1, F2: v3}
        _ = v13
    } else {
        _ = len(v2)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v3 {
        var v14 uint = 1
        fmt.Println(v3)
        _ = v14
    }
    v3 = (v3 && false)
    if false {
        var v15 int64 = v5
        _ = v15
    }
    _ = len(v2)
    if (v3 && v3) {
        var v16 int32 = -2147483648
        var v17 uint = 523
        _ = v16
        _ = v17
    } else {
        var v18 int32 = 702
        var v19 string = ((("" + "hello") + ("foo bar" + "hello")) + (("" + "hello") + (("a" + "a") + "\n\t")))
        _ = v18
        _ = v19
    }
    var v20 float32 = (3.14159 * 1.5)
    var v21 bool = true
    fmt.Println(v3)
    v21 = (!v21)
    fmt.Println(v5)
    v22 := GenS2{F0: 78, F1: -137, F2: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
