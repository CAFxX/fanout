package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 2147483647
    var v3 float64 = 1e10
    var v4 uint16 = 885
    var v5 float32 = 1.5
    v6 := []int{}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int8 = 127
        _ = fi7
        _ = v8
    }
    var v9 bool = (v4 < v4)
    switch ("" + "\n\t") {
    case "foo bar":
        var v10 string = ("" + "")
        _ = v10
    case "":
        var v11 uint8 = 128
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 5; fi12++ {
            v3 = v3
            if true {
                v5 = v5
            }
            _ = fi12
        }
    }
    v2 += (v2 * -2147483648)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint64 = 1
        switch ("\n\t" + "\n\t") {
        case "hello":
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := GenS1{F0: "\n\t", F1: "\n\t"}
                _ = fi15
                _ = v16
            }
        }
        v14 += (v14 & 18446744073709551615)
        _ = fi13
        _ = v14
    }
    fmt.Println(v5, v5)
    if false {
        var v17 uint8 = 50
        _ = v17
    } else {
        var v18 int64 = 201
        _ = v18
    }
    var v19 uint16 = v4
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v20 int16 = -32768
    for fi21 := 0; fi21 < 15; fi21++ {
        v1[fi21] = (!false)
        v22 := GenS2{F0: -38}
        fmt.Println(v3, v9)
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: 79}
    switch "\n\t" {
    case "a":
        switch "" {
        case "hello":
            v24 := GenS1{F0: "hello", F1: "foo bar"}
            _ = v24
        case "":
            var v25 int8 = -128
            _ = v25
        case "a":
            var v26 uint32 = 42
            _ = v26
        }
    case "foo bar":
        var v27 float64 = (float64(0.1) + -2.25)
        _ = v27
    case "\n\t":
        var v28 int64 = 186
        _ = v28
    }
    v4 = (v4 ^ v19)
    var v29 int = -319
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
