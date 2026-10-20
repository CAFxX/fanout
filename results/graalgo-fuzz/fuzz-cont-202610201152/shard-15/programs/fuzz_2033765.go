package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool) int {
    v7 := GenS1{F0: false, F1: -2147483648}
    if p0 {
        p0 = false
        var v8 int64 = 1
        _ = v8
    } else {
        v9 := GenS2{F0: 4294967295, F1: -1, F2: 2147483647}
        _ = v9
    }
    if p0 {
        if p0 {
            fmt.Println(p0)
        }
        p0 = (!p0)
    }
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint64 = 151
    v2 := []int{-198, -209, -1, 1}
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 float32 = (3.14159 + 1.5)
    var v6 string = ("a" + "\n\t")
    fmt.Println(v5, v6)
    if false {
        var v10 int8 = 100
        v11 := GenS3{F0: uint16(v1), F1: (^(v1))}
        _ = v10
        _ = v11
    } else {
        v6 = v6
        var v12 bool = false
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int64 = 42
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v1, v5, v1)
            var v16 uint = 546
            _ = fi15
            _ = v16
        }
        var v17 uint8 = 47
        _ = fi13
        _ = v14
        _ = v17
    }
    if (false || (true || true)) {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 float64 = (0.1 - 1.5)
            v1 = (v1 + 1)
            _ = fi18
            _ = v19
        }
        var v20 uint32 = 2147483647
        _ = v20
    } else {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint = 543
            if true {
                fmt.Println(fi21, v5, fi21)
                var v23 bool = (-337 < 9223372036854775807)
                var v24 uint64 = (v1 & v1)
                _ = v23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        var v25 int8 = 100
        var v26 int = 826
        _ = v25
        _ = v26
    }
    v5 = -2.25
    fmt.Println(v1)
    var v27 bool = false
    v28 := GenS2{F0: 4294967295, F1: 106, F2: 42}
    v27 = (0 == 42)
    var v29 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
