package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }

func main() {
    v1 := make(map[int]int)
    v2 := []int{2147483647}
    var v3 uint64 = 740
    var v4 uint = 42
    v5 := []int{532, 617}
    if (!false) {
        switch 2147483647 {
        case -9223372036854775808:
            var v6 int32 = 42
            _ = v6
        case -346:
            v4 -= 627
        case 32767:
            v4 = (^(v4))
        }
        var v7 float32 = 0.0
        _ = v7
    } else {
        v1[4294967295] = 365
    }
    switch "foo bar" {
    case "foo bar":
        var v8 bool = false
        _ = v8
    case "a":
        fmt.Println(v4, v4)
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        if false {
            v10 := v1[fi9]
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 += (v3 & v3)
        _ = fi11
    }
    v5 = append(v5, -9223372036854775808)
    var v12 int8 = -128
    v3 = 2
    switch 453 {
    case -540:
        switch 4294967295 {
        case -1:
            fmt.Println(v12, v12)
        case -567:
            fmt.Println(v4, v3)
        case 4294967295:
            v3 = (v3 << ((v4) & 63))
        }
    case 849:
        v3 = v3
    case 9223372036854775807:
        var v13 uint64 = v3
        _ = v13
    }
    v3 = v3
    v14 := GenS1{F0: 2, F1: -237, F2: true}
    for fi15 := 0; fi15 < 11; fi15++ {
        v3 = 4294967295
        _ = fi15
    }
    if false {
        v4 = (v4 * v4)
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 int8 = v12
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint32 = 4294967295
                var v20 float64 = 0.1
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        var v21 float64 = float64(0.0)
        _ = v21
    }
    var v22 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
