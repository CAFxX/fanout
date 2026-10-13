package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int64 = 42
        v6 := v2[211]
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        var v7 uint = 4294967295
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int8 = 0
            _ = fi8
            _ = v9
        }
        var v10 float32 = float32(0.1)
        _ = v10
        _ = v7
    } else {
        var v11 string = "a"
        fmt.Println(v11)
        v12 := GenS1{F0: v11, F1: 32767}
        _ = v11
        _ = v12
    }
    var v13 float32 = -2.25
    if false {
        var v14 int64 = -1
        switch (("\n\t" + "\n\t") + "") {
        case "hello":
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 string = ("a" + "\n\t")
                var v17 int32 = 0
                _ = fi15
                _ = v16
                _ = v17
            }
        default:
            fmt.Println(v14, v14)
        }
        v13 += v13
        _ = v14
    } else {
        switch -2147483648 {
        case 4294967295:
            v18 := []int{0}
            _ = v18
        case 305:
            var v19 int64 = 42
            _ = v19
        case -352:
            var v20 uint16 = 255
            _ = v20
        }
    }
    v13 = v13
    var v21 uint64 = 61
    v21 = 2
    v13 += 1.5
    v22 := GenS1{F0: ("a" + "hello"), F1: 32768}
    v23 := []int{2147483647}
    v13 = 3.14159
    switch "\n\t" {
    case "a":
        v2[27] = -141
    case "hello":
        var v24 uint = 1
        _ = v24
    default:
        v25 := GenS1{F0: ("" + "\n\t"), F1: 127}
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
