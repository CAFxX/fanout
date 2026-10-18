package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 0.0
    var v3 uint64 = 42
    if true {
        var v4 string = ""
        _ = v4
    } else {
        var v5 uint32 = uint32(v3)
        _ = v5
    }
    var v6 uint16 = 319
    var v7 int = 32767
    v8 := GenS1{F0: v1, F1: v1}
    switch "\n\t" {
    case "":
        var v9 uint = 255
        _ = v9
    case "\n\t":
        v10 := []int{(-(v7)), (v7 ^ v7)}
        _ = v10
    case "foo bar":
        var v11 uint = 0
        _ = v11
    }
    if (!(100 > -1)) {
        v12 := GenS1{F0: (v1 << ((18446744073709551615) & 63)), F1: -740}
        _ = v12
    }
    v13 := []int{4294967295, v7}
    var v14 int16 = -2
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v3 += v3
    v15 := GenS1{F0: -608, F1: (v1 - -1)}
    for fi16 := 0; fi16 < 13; fi16++ {
        if (1 > 255) {
            var v17 uint8 = 1
            _ = v17
        }
        _ = fi16
    }
    var v18 string = "hello"
    v18 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
