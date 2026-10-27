package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 string) int32 {
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 = v5
        _ = fi6
    }
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 uint, p1 uint, p2 uint) uint {
    var v7 uint = p2
    v8 := []int{-2147483648, 1, 9223372036854775807}
    fmt.Println(v7, p1, p1)
    _ = len(v8)
    if false {
        if false {
            v9 := GenS1{F0: 106}
            _ = v9
        } else {
            var v10 int16 = 350
            _ = v10
        }
    } else {
        v11 := GenS1{F0: 42}
        var v12 uint64 = 2
        var v13 uint8 = 0
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (v7 | (v7 & p2))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1.5
    v3 := []int{9223372036854775807, -975}
    var v4 int16 = 1000
    for fi14 := 0; fi14 < 16; fi14++ {
        v2 = ((float64(float32(1.5)) * v2) - ((v2 + float64(float32(0.1))) * v2))
        if true {
            fmt.Println(v2)
        }
        _ = fi14
    }
    fmt.Println(v4, v2)
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS1{F0: 42}
        var v17 uint32 = 2147483647
        var v18 uint16 = 871
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 bool = false
    v20 := GenS1{F0: 379}
    v21 := GenS2{F0: 2147483647, F1: -1}
    v19 = v19
    switch "\n\t" {
    case "\n\t":
        delete(v1, ("foo bar" + ("a" + "a")))
    case "":
        v19 = v19
    case "foo bar":
        v2 = (v2 / 3.14159)
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        fmt.Println(v19, v2)
        v4 += v4
        _ = fi22
    }
    fmt.Println(v2)
    switch 0 {
    case 9223372036854775807:
        v4 += v4
    }
    fmt.Println(v4, v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
