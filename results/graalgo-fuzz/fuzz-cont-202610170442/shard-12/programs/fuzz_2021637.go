package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0() uint {
    v5 := []int{}
    var v6 uint16 = 42
    v7 := GenS1{F0: 853, F1: -897}
    var v8 uint32 = 1
    switch "foo bar" {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v6, v8)
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 bool = ((1e10 - 0.1) == -2.25)
    var v2 int = 32768
    var v3 uint32 = 786
    var v4 float64 = (-2.25 / 1e10)
    var v9 float32 = 0.0
    var v10 uint8 = 150
    var v11 uint16 = 484
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint8 = v10
            v9 = (-2.25 / 1.0)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float64 = float64(v9)
        v17 := []int{}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int = (^(v2))
    fmt.Println(v11, v4, v1)
    if v1 {
        v3 = (v3 - 226)
    } else {
        v1 = ((-9223372036854775808 <= -1) || false)
        v19 := []int{}
        _ = v19
    }
    v20 := []int{}
    v21 := GenS1{F0: (v3 * (^(v3))), F1: v2}
    fmt.Println(v9, v2, v11)
    for fi22 := 0; fi22 < 5; fi22++ {
        _ = len(v20)
        var v23 uint16 = v11
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: v3, F1: (v2 / (((v2 & v18)) | 1))}
    switch ((((("a" + ("foo bar" + "\n\t")) + "foo bar") + ("" + "a")) + ((("foo bar" + "hello") + ("a" + "\n\t")) + ("foo bar" + "a"))) + "a") {
    case "foo bar":
        _ = len(v20)
    case "hello":
        fmt.Println(v1, v11, v9)
    default:
        switch -360 {
        case 1:
            var v25 uint64 = 0
            _ = v25
        }
    }
    switch "" {
    case "hello":
        var v26 uint64 = 18446744073709551615
        _ = v26
    case "\n\t":
        var v27 uint64 = 1
        _ = v27
    default:
        var v28 bool = v1
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 9; fi30++ {
            fmt.Println(v3, fi30, fi29)
            _ = fi30
        }
        v31 := GenS1{F0: v3, F1: v2}
        _ = fi29
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
