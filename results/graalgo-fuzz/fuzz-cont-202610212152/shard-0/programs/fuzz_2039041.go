package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }

func main() {
    v1 := []int{42}
    v2 := make(map[int]int)
    var v3 int32 = 187
    var v4 int8 = 126
    var v5 float64 = 0.0
    switch (("hello" + ("hello" + "\n\t")) + ("hello" + "foo bar")) {
    case "hello":
        v5 = v5
    case "foo bar":
        fmt.Println(v5, v3)
    case "\n\t":
        v5 -= 0.1
    default:
        fmt.Println(v3, v5)
    }
    v3 -= v3
    v6 := GenS1{F0: 4294967295, F1: 1000}
    v7 := GenS2{F0: true}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            _ = len(v2)
            switch "foo bar" {
            case "hello":
                var v10 int16 = 1000
                _ = v10
            case "":
                v11 := GenS1{F0: -253, F1: 1000}
                _ = v11
            case "\n\t":
                v3 = -2147483648
            }
            _ = fi9
        }
        var v12 uint = 18446744073709551615
        v2[1] = -2147483648
        _ = fi8
        _ = v12
    }
    var v13 uint8 = 42
    if (true && true) {
        if true {
            var v14 string = "foo bar"
            _ = v14
        }
        var v15 int = -2147483648
        for fi16 := 0; fi16 < 1; fi16++ {
            v2[v15] = (fi16 & (fi16 * -938))
            v3 = v3
            _ = fi16
        }
        _ = v15
    }
    if false {
        v17 := GenS1{F0: -1, F1: 32767}
        v5 += -2.25
        _ = v17
    } else {
        _ = len(v1)
        if (416 > -32768) {
            var v18 int32 = -803
            var v19 uint8 = 42
            _ = v18
            _ = v19
        } else {
            var v20 int = 4294967295
            _ = v20
        }
    }
    var v21 uint = 4294967295
    if false {
        var v22 int = 647
        v21 -= (v21 % ((1) | 1))
        _ = v22
    } else {
        var v23 uint32 = 2147483648
        _ = v23
    }
    var v24 int64 = 795
    var v25 int8 = (v4 | 100)
    v26 := v2[int(v24)]
    v27 := GenS3{F0: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
