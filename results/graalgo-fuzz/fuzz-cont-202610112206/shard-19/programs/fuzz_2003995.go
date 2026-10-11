package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 42
    var v3 int = 522
    var v4 uint32 = v2
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            if true {
                v7 := []int{fi6}
                _ = v7
            } else {
                v8 := []int{fi5, (fi6 / (((fi5 << ((4294967295) & 63))) | 1))}
                _ = len(v8)
                _ = v8
            }
            v4 = v2
            _ = fi6
        }
        var v9 int = (-(v3))
        _ = fi5
        _ = v9
    }
    v10 := v1["foo bar"]
    switch "\n\t" {
    case "a":
        v11 := GenS1{F0: 528, F1: 532}
        _ = v11
    case "foo bar":
        var v12 float64 = (-2.25 + 1.5)
        _ = v12
    case "\n\t":
        fmt.Println(v2)
    default:
        var v13 uint = 316
        _ = v13
    }
    switch ("hello" + "foo bar") {
    case "a":
        if false {
            v14 := []int{}
            _ = len(v14)
            v3 = (^(v3))
            _ = v14
        }
    case "foo bar":
        var v15 int32 = 2147483647
        _ = v15
    case "":
        v1["hello"] = (v10 << ((251) & 63))
    default:
        for fi16 := 0; fi16 < 9; fi16++ {
            fmt.Println(v2, v4, fi16)
            _ = fi16
        }
    }
    v17 := []int{v10, -931}
    if false {
        delete(v1, "a")
        v10 = 128
        switch "" {
        case "a":
            v18 := GenS1{F0: 1, F1: 9223372036854775807}
            _ = v18
        case "foo bar":
            v10 = (v10 << ((126) & 63))
        default:
            var v19 float32 = -2.25
            _ = v19
        }
    }
    v20 := GenS1{F0: 814, F1: 231}
    v2 -= v4
    if true {
        v4 -= 0
        v1[("" + "foo bar")] = (^(v10))
    } else {
        _ = len(v1)
    }
    v4 += (v2 % ((v4) | 1))
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v4, v2)
        _ = fi21
    }
    v2 += (^(v2))
    v4 = (v2 + (^(v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
