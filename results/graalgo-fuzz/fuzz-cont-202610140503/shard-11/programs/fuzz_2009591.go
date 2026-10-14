package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0() uint {
    var v5 uint64 = 18446744073709551615
    if true {
        var v6 int8 = 100
        fmt.Println(v5, v6)
        v7 := GenS1{F0: 4294967295, F1: 42}
        _ = v6
        _ = v7
    }
    _ = v5
    return 827
}
func genFn1(p0 int64, p1 bool) int16 {
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 string = ("a" + ("" + ""))
            v11 := GenS1{F0: 18446744073709551615, F1: -1}
            switch v10 {
            case "hello":
                p1 = (1 != 0)
            case "":
                p1 = p1
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{}
        _ = fi8
        _ = v12
    }
    var v13 int16 = 544
    switch "hello" {
    case "hello":
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(p1, p0)
            for fi15 := 0; fi15 < 1; fi15++ {
                p1 = false
                _ = fi15
            }
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return -521
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    var v3 float32 = -2.25
    var v4 int = -864
    v2 = (v2 % (((v2 * (v2 - v2))) | 1))
    var v16 int8 = 20
    v17 := GenS1{F0: 251, F1: v2}
    if false {
        fmt.Println(v3)
    } else {
        var v18 float64 = 1e10
        _ = v18
    }
    fmt.Println(v16, v3, v2)
    _ = len(v1)
    v19 := v1[v4]
    fmt.Println(v4, v19)
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v16, v2, v3)
        _ = fi20
    }
    switch ("" + "\n\t") {
    case "\n\t":
        v16 += 1
    case "foo bar":
        var v21 uint32 = 2147483648
        _ = v21
    }
    v22 := GenS1{F0: 977, F1: v2}
    v23 := GenS1{F0: 304, F1: (v2 + (v2 ^ v2))}
    var v24 bool = (!((v19 % (((^(v4))) | 1)) >= 958))
    v1[(-(v4))] = int(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
