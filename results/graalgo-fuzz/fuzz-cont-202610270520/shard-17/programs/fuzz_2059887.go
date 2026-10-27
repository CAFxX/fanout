package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int {
    if true {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint64 = 1
            var v9 bool = true
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 bool = false
        switch 128 {
        case -657:
            if (1 < -806) {
                v10 = (822 > 2)
                v11 := []int{32767, 0}
                _ = v11
            }
        case -166:
            var v12 uint16 = 0
            _ = v12
        }
        _ = v10
    } else {
        var v13 int = 2147483647
        if false {
            fmt.Println(v13)
            v13 -= v13
            var v14 bool = true
            _ = v14
        }
        var v15 int64 = -1
        _ = v13
        _ = v15
    }
    var v16 int64 = 0
    if ((v16 & (v16 >> ((288) & 63))) < 1) {
        v16 += (v16 | v16)
        switch "hello" {
        case "a":
            v17 := []int{}
            _ = v17
        case "\n\t":
            v18 := []int{}
            _ = v18
        default:
            fmt.Println(v16)
        }
        fmt.Println(v16)
    }
    v16 -= (v16 << ((1) & 63))
    _ = v16
    return 32767
}

func main() {
    var v1 uint64 = 385
    var v2 int32 = 0
    var v3 int8 = -1
    var v4 bool = true
    var v5 float64 = 0.0
    v6 := make(map[string]bool)
    switch "a" {
    case "":
        fmt.Println(v5)
    default:
        var v19 uint = 1
        _ = v19
    }
    var v20 uint8 = 1
    v21 := GenS1{F0: (^(v3))}
    v4 = true
    var v22 uint32 = 1
    v23 := GenS1{F0: (v3 << ((255) & 63))}
    var v24 int16 = 42
    v25 := []int{263, genFn0(), genFn0()}
    switch 9223372036854775807 {
    case 200:
        if true {
            if true {
                var v26 float32 = 0.0
                _ = v26
            }
        }
    default:
        var v27 int32 = (-(v2))
        _ = v27
    }
    fmt.Println(v2, v3, v22)
    v28 := GenS1{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
