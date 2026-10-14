package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int8 {
    var v7 float64 = (0.1 + 1.5)
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int = -9223372036854775808
        v10 := []int{(fi8 | (fi8 * v9)), v9}
        fmt.Println(v7)
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!(-2.25 != v7)) {
        v7 += (v7 / v7)
        v7 += (1e10 * v7)
        fmt.Println(v7)
    }
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    var v3 int16 = 0
    v4 := make(map[uint32]uint32)
    v5 := []int{-290, 42}
    v6 := make(map[string]int)
    var v11 string = "hello"
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int16 = -1
        var v15 uint64 = 4294967295
        for fi16 := 0; fi16 < 9; fi16++ {
            _ = len(v5)
            _ = len(v5)
            fmt.Println(fi13)
            _ = fi16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v17 := GenS1{F0: 18446744073709551615}
    v18 := GenS1{F0: 2}
    var v19 int64 = 575
    fmt.Println(v3, v2)
    var v20 uint8 = v2
    var v21 uint = 1
    if len(v5) > 0 {
        v5[0] = 569
    }
    fmt.Println(v2, v21)
    fmt.Println(v3, v20, v2)
    delete(v1, 2147483648)
    if (!true) {
        v22 := GenS1{F0: 972}
        _ = v22
    }
    var v23 int = 395
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
