package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int16; F1 uint; F2 uint64 }
type GenS3 struct { F0 string; F1 bool }
func genFn0() int32 {
    v7 := []int{}
    var v8 int = -9223372036854775808
    v8 += v8
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 int64, p2 uint) uint32 {
    if false {
        var v9 bool = true
        switch "" {
        case "\n\t":
            v10 := GenS1{F0: -1, F1: (p1 >> ((554) & 63))}
            _ = v10
        case "hello":
            fmt.Println(v9, p2, v9)
        case "a":
            v11 := GenS2{F0: 1000, F1: (p0 & p2), F2: 18446744073709551615}
            _ = v11
        }
        _ = v9
    }
    var v12 uint64 = 458
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 563
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]bool)
    var v6 uint = 1
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int64 = -9223372036854775808
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v6 = (^(v6))
    var v16 float64 = 1.5
    var v17 uint8 = 2
    if true {
        var v18 string = ("foo bar" + "foo bar")
        if true {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 bool = (v1 < v16)
                _ = fi19
                _ = v20
            }
            var v21 uint = v6
            fmt.Println(v1)
            _ = v21
        }
        _ = v18
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 string = "\n\t"
            v6 = (v6 - v6)
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: "foo bar", F1: false}
        _ = v24
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "hello":
        var v25 uint8 = 31
        _ = v25
    case "\n\t":
        fmt.Println(v1)
    case "foo bar":
        fmt.Println(v6, v4, v4)
    }
    var v26 float32 = (0.0 - 0.0)
    if false {
        v1 += (v16 / 0.1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
