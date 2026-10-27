package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0() uint64 {
    var v7 string = ""
    v7 = (v7 + ("" + (v7 + "hello")))
    var v8 uint64 = 427
    v8 = 369
    _ = v7
    _ = v8
    return (^(v8))
}

func main() {
    var v1 int16 = -1
    var v2 bool = ((4294967295 != 146) || false)
    v3 := make(map[string]uint32)
    v4 := []int{9223372036854775807, 32767}
    var v5 string = ""
    v6 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    switch ((v5 + "a") + ("\n\t" + v5)) {
    case "hello":
        v2 = (!v2)
    default:
        v1 += v1
    }
    var v10 uint64 = 42
    switch 42 {
    case 751:
        if false {
            v6[uint32(v10)] = v2
        } else {
            fmt.Println(v1)
            var v11 uint16 = 261
            _ = v11
        }
    }
    if (204 > 742) {
        if false {
            if v2 {
                v10 = (v10 << ((255) & 63))
            } else {
                var v12 uint64 = (v10 & (v10 & genFn0()))
                v13 := GenS1{F0: -9223372036854775808, F1: 250}
                _ = v12
                _ = v13
            }
        }
    } else {
        v2 = true
        if true {
            var v14 int = -1
            fmt.Println(v10)
            _ = v14
        } else {
            if false {
                fmt.Println(v2, v5, v1)
                var v15 uint = 578
                _ = v15
            }
            v10 += 2
        }
        v16 := GenS1{F0: -1, F1: 466}
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint64 = v10
    v5 = "hello"
    if (v5 < v5) {
        v1 = (v1 / ((v1) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
